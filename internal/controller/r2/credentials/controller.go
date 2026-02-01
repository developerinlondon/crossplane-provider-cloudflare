package credentials

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go"
	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"gitlab.com/jarvisai.run/provider-cloudflare/apis/r2/v1alpha1"
	apisv1beta1 "gitlab.com/jarvisai.run/provider-cloudflare/apis/v1beta1"
)

const (
	errNotCredentials       = "managed resource is not a Credentials custom resource"
	errGetProviderConfig    = "cannot get provider config"
	errGetCredentials       = "cannot get credentials"
	errUnmarshalCredentials = "cannot unmarshal credentials"
	errCreateToken          = "cannot create API token"
	errDeleteToken          = "cannot delete API token"
	errGetToken             = "cannot get API token"
	errLookupPermissions    = "cannot lookup permission groups"
)

const (
	r2ReadPermissionName  = "Workers R2 Storage Bucket Item Read"
	r2WritePermissionName = "Workers R2 Storage Bucket Item Write"
)

func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.Credentials_GroupVersionKind.String())

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.Credentials_GroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:   mgr.GetClient(),
			logger: o.Logger,
		}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithPollInterval(o.PollInterval),
		managed.WithTimeout(3*time.Minute),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.Credentials{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

func SetupGated(mgr ctrl.Manager, o tjcontroller.Options) error {
	return Setup(mgr, o)
}

type connector struct {
	kube   client.Client
	logger logging.Logger
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.Credentials)
	if !ok {
		return nil, errors.New(errNotCredentials)
	}

	apiToken, err := c.getCloudflareCredentials(ctx, cr)
	if err != nil {
		return nil, err
	}

	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create Cloudflare client")
	}

	return &external{
		kube:   c.kube,
		api:    api,
		logger: c.logger,
	}, nil
}

func (c *connector) getCloudflareCredentials(ctx context.Context, cr *v1alpha1.Credentials) (string, error) {
	configRef := cr.GetProviderConfigReference()
	if configRef == nil {
		return "", errors.New("no providerConfigRef provided")
	}

	pc := &apisv1beta1.ProviderConfig{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return "", errors.Wrap(err, errGetProviderConfig)
	}

	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c.kube, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return "", errors.Wrap(err, errGetCredentials)
	}

	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", errors.Wrap(err, errUnmarshalCredentials)
	}

	if token, ok := creds["api_token"]; ok && token != "" {
		return token, nil
	}

	return "", errors.New("no api_token found in credentials")
}

type external struct {
	kube   client.Client
	api    *cloudflare.API
	logger logging.Logger
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.Credentials)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotCredentials)
	}

	externalName := meta.GetExternalName(cr)
	if externalName == "" {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	token, err := e.api.GetAPIToken(ctx, externalName)
	if err != nil {
		if isNotFound(err) {
			return managed.ExternalObservation{ResourceExists: false}, nil
		}
		return managed.ExternalObservation{}, errors.Wrap(err, errGetToken)
	}

	cr.Status.AtProvider.TokenID = token.ID
	cr.Status.AtProvider.Status = token.Status
	if !token.IssuedOn.IsZero() {
		cr.Status.AtProvider.IssuedOn = token.IssuedOn.Format(time.RFC3339)
	}
	cr.Status.AtProvider.Endpoint = fmt.Sprintf("%s.r2.cloudflarestorage.com", cr.Spec.ForProvider.AccountID)

	cr.SetConditions(xpv1.Available())

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.Credentials)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotCredentials)
	}

	permGroups, err := e.lookupR2PermissionGroups(ctx, cr.Spec.ForProvider.Permissions)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errLookupPermissions)
	}

	resources := buildResourceScope(cr.Spec.ForProvider.AccountID, cr.Spec.ForProvider.BucketName)

	tokenRequest := cloudflare.APIToken{
		Name: cr.Spec.ForProvider.Name,
		Policies: []cloudflare.APITokenPolicies{{
			Effect:           "allow",
			Resources:        resources,
			PermissionGroups: permGroups,
		}},
	}

	token, err := e.api.CreateAPIToken(ctx, tokenRequest)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateToken)
	}

	meta.SetExternalName(cr, token.ID)

	secretAccessKey := sha256Hash(token.Value)

	connDetails := managed.ConnectionDetails{
		"access_key_id":     []byte(token.ID),
		"secret_access_key": []byte(secretAccessKey),
		"endpoint":          []byte(fmt.Sprintf("%s.r2.cloudflarestorage.com", cr.Spec.ForProvider.AccountID)),
		"token_value":       []byte(token.Value),
	}

	return managed.ExternalCreation{
		ConnectionDetails: connDetails,
	}, nil
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*v1alpha1.Credentials)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errNotCredentials)
	}

	externalName := meta.GetExternalName(cr)
	if externalName == "" {
		return managed.ExternalDelete{}, nil
	}

	err := e.api.DeleteAPIToken(ctx, externalName)
	if err != nil && !isNotFound(err) {
		return managed.ExternalDelete{}, errors.Wrap(err, errDeleteToken)
	}

	return managed.ExternalDelete{}, nil
}

func (e *external) Disconnect(ctx context.Context) error {
	return nil
}

func (e *external) lookupR2PermissionGroups(ctx context.Context, permissions []string) ([]cloudflare.APITokenPermissionGroups, error) {
	allGroups, err := e.api.ListAPITokensPermissionGroups(ctx)
	if err != nil {
		return nil, err
	}

	wantRead := contains(permissions, "read")
	wantWrite := contains(permissions, "write")

	if len(permissions) == 0 {
		wantRead = true
		wantWrite = true
	}

	var result []cloudflare.APITokenPermissionGroups
	for _, g := range allGroups {
		if wantRead && g.Name == r2ReadPermissionName {
			result = append(result, cloudflare.APITokenPermissionGroups{ID: g.ID})
		}
		if wantWrite && g.Name == r2WritePermissionName {
			result = append(result, cloudflare.APITokenPermissionGroups{ID: g.ID})
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no R2 permission groups found")
	}

	return result, nil
}

func buildResourceScope(accountID string, bucketName *string) map[string]interface{} {
	if bucketName != nil && *bucketName != "" {
		return map[string]interface{}{
			fmt.Sprintf("com.cloudflare.edge.r2.bucket.%s_default_%s", accountID, *bucketName): "*",
		}
	}
	return map[string]interface{}{
		fmt.Sprintf("com.cloudflare.api.account.%s", accountID): "*",
	}
}

func sha256Hash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isNotFound(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return errStr == "HTTP status 404: Could not find token" ||
		errStr == "HTTP status 404" ||
		errStr == "could not find token"
}
