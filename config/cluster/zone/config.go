package zone

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *ujconfig.Resource) {
		r.ShortGroup = "zone"
		r.Kind = "Zone"
	})
	p.AddResourceConfigurator("cloudflare_zone_dnssec", func(r *ujconfig.Resource) {
		r.ShortGroup = "zone"
		r.Kind = "DNSSEC"
	})
	p.AddResourceConfigurator("cloudflare_zone_setting", func(r *ujconfig.Resource) {
		r.ShortGroup = "zone"
		r.Kind = "Setting"
	})
}
