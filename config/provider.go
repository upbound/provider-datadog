/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/cluster"
	"github.com/upbound/provider-datadog/config/namespaced"
)

const (
	resourcePrefix = "datadog"
	modulePath     = "github.com/upbound/provider-datadog"

	// Every resource configurator sets a ShortGroup such as "datadog",
	// "apm.datadog" or "synthetics.datadog" that upjet appends to the root
	// group, so the legacy cluster-scoped tree lands under *.datadog.upbound.io
	// and the namespaced tree under *.datadog.m.upbound.io.
	rootGroupCluster    = "upbound.io"
	rootGroupNamespaced = "m.upbound.io"

	// exampleManifestNamespace is the namespace set on the generated example
	// manifests of namespaced resources.
	exampleManifestNamespace = "crossplane-system"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns the cluster-scoped provider configuration.
func GetProvider() *ujconfig.Provider {
	pc := newProvider(rootGroupCluster)
	for _, configure := range cluster.Configurators {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration.
func GetProviderNamespaced() *ujconfig.Provider {
	pc := newProvider(rootGroupNamespaced,
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: exampleManifestNamespace,
		}))
	for _, configure := range namespaced.Configurators {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}

func newProvider(rootGroup string, opts ...ujconfig.ProviderOption) *ujconfig.Provider {
	options := append([]ujconfig.ProviderOption{
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(ExternalNameConfigurations()),
	}, opts...)
	return ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata), options...)
}
