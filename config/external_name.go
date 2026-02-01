package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

var ExternalNameConfigs = map[string]config.ExternalName{
	// R2 Resources
	"cloudflare_r2_bucket":                    config.IdentifierFromProvider,
	"cloudflare_r2_bucket_cors":               config.IdentifierFromProvider,
	"cloudflare_r2_bucket_lifecycle":          config.IdentifierFromProvider,
	"cloudflare_r2_bucket_event_notification": config.IdentifierFromProvider,
	"cloudflare_r2_bucket_lock":               config.IdentifierFromProvider,
	"cloudflare_r2_bucket_sippy":              config.IdentifierFromProvider,
	"cloudflare_r2_custom_domain":             config.IdentifierFromProvider,
	"cloudflare_r2_managed_domain":            config.IdentifierFromProvider,

	// Zone Resources
	"cloudflare_zone":         config.IdentifierFromProvider,
	"cloudflare_zone_dnssec":  config.IdentifierFromProvider,
	"cloudflare_zone_setting": config.IdentifierFromProvider,
}

func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		l[i] = name + "$"
		i++
	}
	return l
}
