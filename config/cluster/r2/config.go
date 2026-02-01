package r2

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("cloudflare_r2_bucket", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "Bucket"
	})
	p.AddResourceConfigurator("cloudflare_r2_bucket_cors", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "BucketCORS"
	})
	p.AddResourceConfigurator("cloudflare_r2_bucket_lifecycle", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "BucketLifecycle"
	})
	p.AddResourceConfigurator("cloudflare_r2_bucket_event_notification", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "BucketEventNotification"
	})
	p.AddResourceConfigurator("cloudflare_r2_bucket_lock", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "BucketLock"
	})
	p.AddResourceConfigurator("cloudflare_r2_bucket_sippy", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "BucketSippy"
	})
	p.AddResourceConfigurator("cloudflare_r2_custom_domain", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "CustomDomain"
	})
	p.AddResourceConfigurator("cloudflare_r2_managed_domain", func(r *ujconfig.Resource) {
		r.ShortGroup = "r2"
		r.Kind = "ManagedDomain"
	})
}
