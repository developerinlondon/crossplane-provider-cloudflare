package config

import (
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "gitlab.com/jarvisai.run/provider-cloudflare"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudflare.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	pc.ConfigureResources()
	return pc
}

func GetProviderNamespaced() *ujconfig.Provider {
	return nil
}
