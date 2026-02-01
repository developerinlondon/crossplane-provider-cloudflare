// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	rule "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/access/rule"
	dnssettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/account/dnssettings"
	dnssettingsinternalview "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/account/dnssettingsinternalview"
	member "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/account/member"
	subscription "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/account/subscription"
	token "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/account/token"
	shield "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/api/shield"
	shielddiscoveryoperation "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/api/shielddiscoveryoperation"
	shieldoperation "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/api/shieldoperation"
	shieldoperationschemavalidationsettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/api/shieldoperationschemavalidationsettings"
	shieldschema "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/api/shieldschema"
	shieldschemavalidationsettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/api/shieldschemavalidationsettings"
	tokenapi "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/api/token"
	smartrouting "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/argo/smartrouting"
	tieredcaching "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/argo/tieredcaching"
	originpulls "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/authenticated/originpulls"
	originpullscertificate "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/authenticated/originpullscertificate"
	originpullssettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/authenticated/originpullssettings"
	management "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/bot/management"
	ipprefix "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/byo/ipprefix"
	sfuapp "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/calls/sfuapp"
	turnapp "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/calls/turnapp"
	pack "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/certificate/pack"
	connectorrules "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloud/connectorrules"
	account "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/account"
	filter "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/filter"
	healthcheck "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/healthcheck"
	image "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/image"
	list "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/list"
	organization "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/organization"
	queue "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/queue"
	ruleset "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/ruleset"
	snippet "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/snippet"
	snippets "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/snippets"
	stream "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/stream"
	user "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/user"
	worker "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/worker"
	workflow "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/workflow"
	zone "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudflare/zone"
	onerequest "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudforce/onerequest"
	onerequestasset "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudforce/onerequestasset"
	onerequestmessage "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudforce/onerequestmessage"
	onerequestpriority "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/cloudforce/onerequestpriority"
	directoryservice "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/connectivity/directoryservice"
	scanning "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/content/scanning"
	scanningexpression "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/content/scanningexpression"
	hostname "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/custom/hostname"
	hostnamefallbackorigin "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/custom/hostnamefallbackorigin"
	pages "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/custom/pages"
	ssl "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/custom/ssl"
	database "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/d1/database"
	firewall "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/dns/firewall"
	record "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/dns/record"
	zonetransfersacl "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/dns/zonetransfersacl"
	zonetransfersincoming "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/dns/zonetransfersincoming"
	zonetransfersoutgoing "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/dns/zonetransfersoutgoing"
	zonetransferspeer "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/dns/zonetransferspeer"
	zonetransferstsig "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/dns/zonetransferstsig"
	routingaddress "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/routingaddress"
	routingcatchall "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/routingcatchall"
	routingdns "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/routingdns"
	routingrule "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/routingrule"
	routingsettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/routingsettings"
	securityblocksender "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/securityblocksender"
	securityimpersonationregistry "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/securityimpersonationregistry"
	securitytrusteddomains "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/email/securitytrusteddomains"
	rulefirewall "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/firewall/rule"
	tlssetting "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/hostname/tlssetting"
	config "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/hyperdrive/config"
	variant "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/image/variant"
	certificate "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/keyless/certificate"
	credentialcheck "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/leaked/credentialcheck"
	credentialcheckrule "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/leaked/credentialcheckrule"
	item "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/list/item"
	balancer "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/load/balancer"
	balancermonitor "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/load/balancermonitor"
	balancerpool "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/load/balancerpool"
	retention "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/logpull/retention"
	job "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/logpush/job"
	ownershipchallenge "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/logpush/ownershipchallenge"
	networkmonitoringconfiguration "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/networkmonitoringconfiguration"
	networkmonitoringrule "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/networkmonitoringrule"
	transitconnector "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/transitconnector"
	transitsite "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/transitsite"
	transitsiteacl "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/transitsiteacl"
	transitsitelan "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/transitsitelan"
	transitsitewan "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/transitsitewan"
	wangretunnel "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/wangretunnel"
	wanipsectunnel "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/wanipsectunnel"
	wanstaticroute "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/magic/wanstaticroute"
	transforms "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/managed/transforms"
	certificatemtls "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/mtls/certificate"
	policy "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/notification/policy"
	policywebhooks "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/notification/policywebhooks"
	scheduledtest "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/observatory/scheduledtest"
	profile "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/organization/profile"
	cacertificate "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/origin/cacertificate"
	rulepage "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/page/rule"
	shieldpolicy "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/page/shieldpolicy"
	domain "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/pages/domain"
	project "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/pages/project"
	providerconfig "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/providerconfig"
	consumer "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/queue/consumer"
	bucket "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucket"
	bucketcors "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketcors"
	bucketeventnotification "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketeventnotification"
	bucketlifecycle "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketlifecycle"
	bucketlock "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketlock"
	bucketsippy "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/bucketsippy"
	customdomain "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/customdomain"
	manageddomain "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/r2/manageddomain"
	limit "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/rate/limit"
	hostnameregional "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/regional/hostname"
	tieredcache "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/regional/tieredcache"
	domainregistrar "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/registrar/domain"
	validationoperationsettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/schema/validationoperationsettings"
	validationschemas "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/schema/validationschemas"
	validationsettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/schema/validationsettings"
	rules "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/snippet/rules"
	application "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/spectrum/application"
	connector "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/sso/connector"
	audiotrack "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/stream/audiotrack"
	captionlanguage "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/stream/captionlanguage"
	download "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/stream/download"
	key "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/stream/key"
	liveinput "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/stream/liveinput"
	watermark "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/stream/watermark"
	webhook "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/stream/webhook"
	cache "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/tiered/cache"
	validationconfig "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/token/validationconfig"
	validationrules "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/token/validationrules"
	tls "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/total/tls"
	widget "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/turnstile/widget"
	sslsetting "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/universal/sslsetting"
	normalizationsettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/url/normalizationsettings"
	agentblockingrule "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/user/agentblockingrule"
	room "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/waiting/room"
	roomevent "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/waiting/roomevent"
	roomrules "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/waiting/roomrules"
	roomsettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/waiting/roomsettings"
	analyticsrule "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/web/analyticsrule"
	analyticssite "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/web/analyticssite"
	hostnameweb3 "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/web3/hostname"
	version "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/worker/version"
	crontrigger "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/crontrigger"
	customdomainworkers "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/customdomain"
	deployment "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/deployment"
	forplatformsdispatchnamespace "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/forplatformsdispatchnamespace"
	kv "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/kv"
	kvnamespace "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/kvnamespace"
	route "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/route"
	script "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/script"
	scriptsubdomain "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/workers/scriptsubdomain"
	trustaccessaicontrolsmcpportal "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessaicontrolsmcpportal"
	trustaccessaicontrolsmcpserver "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessaicontrolsmcpserver"
	trustaccessapplication "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessapplication"
	trustaccesscustompage "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccesscustompage"
	trustaccessgroup "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessgroup"
	trustaccessidentityprovider "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessidentityprovider"
	trustaccessinfrastructuretarget "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessinfrastructuretarget"
	trustaccesskeyconfiguration "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccesskeyconfiguration"
	trustaccessmtlscertificate "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessmtlscertificate"
	trustaccessmtlshostnamesettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessmtlshostnamesettings"
	trustaccesspolicy "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccesspolicy"
	trustaccessservicetoken "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessservicetoken"
	trustaccessshortlivedcertificate "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccessshortlivedcertificate"
	trustaccesstag "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustaccesstag"
	trustdevicecustomprofile "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicecustomprofile"
	trustdevicecustomprofilelocaldomainfallback "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicecustomprofilelocaldomainfallback"
	trustdevicedefaultprofile "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicedefaultprofile"
	trustdevicedefaultprofilecertificates "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicedefaultprofilecertificates"
	trustdevicedefaultprofilelocaldomainfallback "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicedefaultprofilelocaldomainfallback"
	trustdevicemanagednetworks "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicemanagednetworks"
	trustdevicepostureintegration "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicepostureintegration"
	trustdeviceposturerule "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdeviceposturerule"
	trustdevicesettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdevicesettings"
	trustdextest "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdextest"
	trustdlpcustomentry "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdlpcustomentry"
	trustdlpcustomprofile "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdlpcustomprofile"
	trustdlpdataset "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdlpdataset"
	trustdlpentry "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdlpentry"
	trustdlpintegrationentry "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdlpintegrationentry"
	trustdlppredefinedentry "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdlppredefinedentry"
	trustdlppredefinedprofile "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdlppredefinedprofile"
	trustdnslocation "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustdnslocation"
	trustgatewaycertificate "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustgatewaycertificate"
	trustgatewaylogging "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustgatewaylogging"
	trustgatewaypolicy "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustgatewaypolicy"
	trustgatewayproxyendpoint "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustgatewayproxyendpoint"
	trustgatewaysettings "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustgatewaysettings"
	trustlist "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustlist"
	trustnetworkhostnameroute "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustnetworkhostnameroute"
	trustorganization "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustorganization"
	trustriskbehavior "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustriskbehavior"
	trustriskscoringintegration "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trustriskscoringintegration"
	trusttunnelcloudflared "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trusttunnelcloudflared"
	trusttunnelcloudflaredconfig "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trusttunnelcloudflaredconfig"
	trusttunnelcloudflaredroute "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trusttunnelcloudflaredroute"
	trusttunnelcloudflaredvirtualnetwork "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trusttunnelcloudflaredvirtualnetwork"
	trusttunnelwarpconnector "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zero/trusttunnelwarpconnector"
	cachereserve "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/cachereserve"
	cachevariants "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/cachevariants"
	dnssec "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/dnssec"
	dnssettingszone "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/dnssettings"
	hold "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/hold"
	lockdown "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/lockdown"
	setting "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/setting"
	subscriptionzone "gitlab.com/jarvisai.run/provider-cloudflare/internal/controller/zone/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rule.Setup,
		dnssettings.Setup,
		dnssettingsinternalview.Setup,
		member.Setup,
		subscription.Setup,
		token.Setup,
		shield.Setup,
		shielddiscoveryoperation.Setup,
		shieldoperation.Setup,
		shieldoperationschemavalidationsettings.Setup,
		shieldschema.Setup,
		shieldschemavalidationsettings.Setup,
		tokenapi.Setup,
		smartrouting.Setup,
		tieredcaching.Setup,
		originpulls.Setup,
		originpullscertificate.Setup,
		originpullssettings.Setup,
		management.Setup,
		ipprefix.Setup,
		sfuapp.Setup,
		turnapp.Setup,
		pack.Setup,
		connectorrules.Setup,
		account.Setup,
		filter.Setup,
		healthcheck.Setup,
		image.Setup,
		list.Setup,
		organization.Setup,
		queue.Setup,
		ruleset.Setup,
		snippet.Setup,
		snippets.Setup,
		stream.Setup,
		user.Setup,
		worker.Setup,
		workflow.Setup,
		zone.Setup,
		onerequest.Setup,
		onerequestasset.Setup,
		onerequestmessage.Setup,
		onerequestpriority.Setup,
		directoryservice.Setup,
		scanning.Setup,
		scanningexpression.Setup,
		hostname.Setup,
		hostnamefallbackorigin.Setup,
		pages.Setup,
		ssl.Setup,
		database.Setup,
		firewall.Setup,
		record.Setup,
		zonetransfersacl.Setup,
		zonetransfersincoming.Setup,
		zonetransfersoutgoing.Setup,
		zonetransferspeer.Setup,
		zonetransferstsig.Setup,
		routingaddress.Setup,
		routingcatchall.Setup,
		routingdns.Setup,
		routingrule.Setup,
		routingsettings.Setup,
		securityblocksender.Setup,
		securityimpersonationregistry.Setup,
		securitytrusteddomains.Setup,
		rulefirewall.Setup,
		tlssetting.Setup,
		config.Setup,
		variant.Setup,
		certificate.Setup,
		credentialcheck.Setup,
		credentialcheckrule.Setup,
		item.Setup,
		balancer.Setup,
		balancermonitor.Setup,
		balancerpool.Setup,
		retention.Setup,
		job.Setup,
		ownershipchallenge.Setup,
		networkmonitoringconfiguration.Setup,
		networkmonitoringrule.Setup,
		transitconnector.Setup,
		transitsite.Setup,
		transitsiteacl.Setup,
		transitsitelan.Setup,
		transitsitewan.Setup,
		wangretunnel.Setup,
		wanipsectunnel.Setup,
		wanstaticroute.Setup,
		transforms.Setup,
		certificatemtls.Setup,
		policy.Setup,
		policywebhooks.Setup,
		scheduledtest.Setup,
		profile.Setup,
		cacertificate.Setup,
		rulepage.Setup,
		shieldpolicy.Setup,
		domain.Setup,
		project.Setup,
		providerconfig.Setup,
		consumer.Setup,
		bucket.Setup,
		bucketcors.Setup,
		bucketeventnotification.Setup,
		bucketlifecycle.Setup,
		bucketlock.Setup,
		bucketsippy.Setup,
		customdomain.Setup,
		manageddomain.Setup,
		limit.Setup,
		hostnameregional.Setup,
		tieredcache.Setup,
		domainregistrar.Setup,
		validationoperationsettings.Setup,
		validationschemas.Setup,
		validationsettings.Setup,
		rules.Setup,
		application.Setup,
		connector.Setup,
		audiotrack.Setup,
		captionlanguage.Setup,
		download.Setup,
		key.Setup,
		liveinput.Setup,
		watermark.Setup,
		webhook.Setup,
		cache.Setup,
		validationconfig.Setup,
		validationrules.Setup,
		tls.Setup,
		widget.Setup,
		sslsetting.Setup,
		normalizationsettings.Setup,
		agentblockingrule.Setup,
		room.Setup,
		roomevent.Setup,
		roomrules.Setup,
		roomsettings.Setup,
		analyticsrule.Setup,
		analyticssite.Setup,
		hostnameweb3.Setup,
		version.Setup,
		crontrigger.Setup,
		customdomainworkers.Setup,
		deployment.Setup,
		forplatformsdispatchnamespace.Setup,
		kv.Setup,
		kvnamespace.Setup,
		route.Setup,
		script.Setup,
		scriptsubdomain.Setup,
		trustaccessaicontrolsmcpportal.Setup,
		trustaccessaicontrolsmcpserver.Setup,
		trustaccessapplication.Setup,
		trustaccesscustompage.Setup,
		trustaccessgroup.Setup,
		trustaccessidentityprovider.Setup,
		trustaccessinfrastructuretarget.Setup,
		trustaccesskeyconfiguration.Setup,
		trustaccessmtlscertificate.Setup,
		trustaccessmtlshostnamesettings.Setup,
		trustaccesspolicy.Setup,
		trustaccessservicetoken.Setup,
		trustaccessshortlivedcertificate.Setup,
		trustaccesstag.Setup,
		trustdevicecustomprofile.Setup,
		trustdevicecustomprofilelocaldomainfallback.Setup,
		trustdevicedefaultprofile.Setup,
		trustdevicedefaultprofilecertificates.Setup,
		trustdevicedefaultprofilelocaldomainfallback.Setup,
		trustdevicemanagednetworks.Setup,
		trustdevicepostureintegration.Setup,
		trustdeviceposturerule.Setup,
		trustdevicesettings.Setup,
		trustdextest.Setup,
		trustdlpcustomentry.Setup,
		trustdlpcustomprofile.Setup,
		trustdlpdataset.Setup,
		trustdlpentry.Setup,
		trustdlpintegrationentry.Setup,
		trustdlppredefinedentry.Setup,
		trustdlppredefinedprofile.Setup,
		trustdnslocation.Setup,
		trustgatewaycertificate.Setup,
		trustgatewaylogging.Setup,
		trustgatewaypolicy.Setup,
		trustgatewayproxyendpoint.Setup,
		trustgatewaysettings.Setup,
		trustlist.Setup,
		trustnetworkhostnameroute.Setup,
		trustorganization.Setup,
		trustriskbehavior.Setup,
		trustriskscoringintegration.Setup,
		trusttunnelcloudflared.Setup,
		trusttunnelcloudflaredconfig.Setup,
		trusttunnelcloudflaredroute.Setup,
		trusttunnelcloudflaredvirtualnetwork.Setup,
		trusttunnelwarpconnector.Setup,
		cachereserve.Setup,
		cachevariants.Setup,
		dnssec.Setup,
		dnssettingszone.Setup,
		hold.Setup,
		lockdown.Setup,
		setting.Setup,
		subscriptionzone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rule.SetupGated,
		dnssettings.SetupGated,
		dnssettingsinternalview.SetupGated,
		member.SetupGated,
		subscription.SetupGated,
		token.SetupGated,
		shield.SetupGated,
		shielddiscoveryoperation.SetupGated,
		shieldoperation.SetupGated,
		shieldoperationschemavalidationsettings.SetupGated,
		shieldschema.SetupGated,
		shieldschemavalidationsettings.SetupGated,
		tokenapi.SetupGated,
		smartrouting.SetupGated,
		tieredcaching.SetupGated,
		originpulls.SetupGated,
		originpullscertificate.SetupGated,
		originpullssettings.SetupGated,
		management.SetupGated,
		ipprefix.SetupGated,
		sfuapp.SetupGated,
		turnapp.SetupGated,
		pack.SetupGated,
		connectorrules.SetupGated,
		account.SetupGated,
		filter.SetupGated,
		healthcheck.SetupGated,
		image.SetupGated,
		list.SetupGated,
		organization.SetupGated,
		queue.SetupGated,
		ruleset.SetupGated,
		snippet.SetupGated,
		snippets.SetupGated,
		stream.SetupGated,
		user.SetupGated,
		worker.SetupGated,
		workflow.SetupGated,
		zone.SetupGated,
		onerequest.SetupGated,
		onerequestasset.SetupGated,
		onerequestmessage.SetupGated,
		onerequestpriority.SetupGated,
		directoryservice.SetupGated,
		scanning.SetupGated,
		scanningexpression.SetupGated,
		hostname.SetupGated,
		hostnamefallbackorigin.SetupGated,
		pages.SetupGated,
		ssl.SetupGated,
		database.SetupGated,
		firewall.SetupGated,
		record.SetupGated,
		zonetransfersacl.SetupGated,
		zonetransfersincoming.SetupGated,
		zonetransfersoutgoing.SetupGated,
		zonetransferspeer.SetupGated,
		zonetransferstsig.SetupGated,
		routingaddress.SetupGated,
		routingcatchall.SetupGated,
		routingdns.SetupGated,
		routingrule.SetupGated,
		routingsettings.SetupGated,
		securityblocksender.SetupGated,
		securityimpersonationregistry.SetupGated,
		securitytrusteddomains.SetupGated,
		rulefirewall.SetupGated,
		tlssetting.SetupGated,
		config.SetupGated,
		variant.SetupGated,
		certificate.SetupGated,
		credentialcheck.SetupGated,
		credentialcheckrule.SetupGated,
		item.SetupGated,
		balancer.SetupGated,
		balancermonitor.SetupGated,
		balancerpool.SetupGated,
		retention.SetupGated,
		job.SetupGated,
		ownershipchallenge.SetupGated,
		networkmonitoringconfiguration.SetupGated,
		networkmonitoringrule.SetupGated,
		transitconnector.SetupGated,
		transitsite.SetupGated,
		transitsiteacl.SetupGated,
		transitsitelan.SetupGated,
		transitsitewan.SetupGated,
		wangretunnel.SetupGated,
		wanipsectunnel.SetupGated,
		wanstaticroute.SetupGated,
		transforms.SetupGated,
		certificatemtls.SetupGated,
		policy.SetupGated,
		policywebhooks.SetupGated,
		scheduledtest.SetupGated,
		profile.SetupGated,
		cacertificate.SetupGated,
		rulepage.SetupGated,
		shieldpolicy.SetupGated,
		domain.SetupGated,
		project.SetupGated,
		providerconfig.SetupGated,
		consumer.SetupGated,
		bucket.SetupGated,
		bucketcors.SetupGated,
		bucketeventnotification.SetupGated,
		bucketlifecycle.SetupGated,
		bucketlock.SetupGated,
		bucketsippy.SetupGated,
		customdomain.SetupGated,
		manageddomain.SetupGated,
		limit.SetupGated,
		hostnameregional.SetupGated,
		tieredcache.SetupGated,
		domainregistrar.SetupGated,
		validationoperationsettings.SetupGated,
		validationschemas.SetupGated,
		validationsettings.SetupGated,
		rules.SetupGated,
		application.SetupGated,
		connector.SetupGated,
		audiotrack.SetupGated,
		captionlanguage.SetupGated,
		download.SetupGated,
		key.SetupGated,
		liveinput.SetupGated,
		watermark.SetupGated,
		webhook.SetupGated,
		cache.SetupGated,
		validationconfig.SetupGated,
		validationrules.SetupGated,
		tls.SetupGated,
		widget.SetupGated,
		sslsetting.SetupGated,
		normalizationsettings.SetupGated,
		agentblockingrule.SetupGated,
		room.SetupGated,
		roomevent.SetupGated,
		roomrules.SetupGated,
		roomsettings.SetupGated,
		analyticsrule.SetupGated,
		analyticssite.SetupGated,
		hostnameweb3.SetupGated,
		version.SetupGated,
		crontrigger.SetupGated,
		customdomainworkers.SetupGated,
		deployment.SetupGated,
		forplatformsdispatchnamespace.SetupGated,
		kv.SetupGated,
		kvnamespace.SetupGated,
		route.SetupGated,
		script.SetupGated,
		scriptsubdomain.SetupGated,
		trustaccessaicontrolsmcpportal.SetupGated,
		trustaccessaicontrolsmcpserver.SetupGated,
		trustaccessapplication.SetupGated,
		trustaccesscustompage.SetupGated,
		trustaccessgroup.SetupGated,
		trustaccessidentityprovider.SetupGated,
		trustaccessinfrastructuretarget.SetupGated,
		trustaccesskeyconfiguration.SetupGated,
		trustaccessmtlscertificate.SetupGated,
		trustaccessmtlshostnamesettings.SetupGated,
		trustaccesspolicy.SetupGated,
		trustaccessservicetoken.SetupGated,
		trustaccessshortlivedcertificate.SetupGated,
		trustaccesstag.SetupGated,
		trustdevicecustomprofile.SetupGated,
		trustdevicecustomprofilelocaldomainfallback.SetupGated,
		trustdevicedefaultprofile.SetupGated,
		trustdevicedefaultprofilecertificates.SetupGated,
		trustdevicedefaultprofilelocaldomainfallback.SetupGated,
		trustdevicemanagednetworks.SetupGated,
		trustdevicepostureintegration.SetupGated,
		trustdeviceposturerule.SetupGated,
		trustdevicesettings.SetupGated,
		trustdextest.SetupGated,
		trustdlpcustomentry.SetupGated,
		trustdlpcustomprofile.SetupGated,
		trustdlpdataset.SetupGated,
		trustdlpentry.SetupGated,
		trustdlpintegrationentry.SetupGated,
		trustdlppredefinedentry.SetupGated,
		trustdlppredefinedprofile.SetupGated,
		trustdnslocation.SetupGated,
		trustgatewaycertificate.SetupGated,
		trustgatewaylogging.SetupGated,
		trustgatewaypolicy.SetupGated,
		trustgatewayproxyendpoint.SetupGated,
		trustgatewaysettings.SetupGated,
		trustlist.SetupGated,
		trustnetworkhostnameroute.SetupGated,
		trustorganization.SetupGated,
		trustriskbehavior.SetupGated,
		trustriskscoringintegration.SetupGated,
		trusttunnelcloudflared.SetupGated,
		trusttunnelcloudflaredconfig.SetupGated,
		trusttunnelcloudflaredroute.SetupGated,
		trusttunnelcloudflaredvirtualnetwork.SetupGated,
		trusttunnelwarpconnector.SetupGated,
		cachereserve.SetupGated,
		cachevariants.SetupGated,
		dnssec.SetupGated,
		dnssettingszone.SetupGated,
		hold.SetupGated,
		lockdown.SetupGated,
		setting.SetupGated,
		subscriptionzone.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
