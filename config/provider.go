/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/upbound/provider-datadog/config/cluster"
	"github.com/upbound/provider-datadog/config/namespaced"
	"github.com/upbound/provider-datadog/config/templates"
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
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns the cluster-scoped provider configuration. sdkProvider
// and fwProvider are the in-process Terraform providers the resources are
// reconciled with. terraformPluginSDKExternalNameConfigs and
// terraformPluginFrameworkExternalNameConfigs assign each resource to one of
// them, and upjet panics when a framework resource is missing from fwProvider
// or, outside code generation, a plugin SDK resource is missing from
// sdkProvider. generationProvider selects the code generation flavour: the
// plugin SDK resources are then described by the JSON schema instead of the
// live Go schema (see codegenSDKProvider) and the schema-mutating
// configurators in codegen.go run.
func GetProvider(sdkProvider *schema.Provider, fwProvider fwprovider.Provider, generationProvider bool) (*ujconfig.Provider, error) {
	return newProvider(rootGroupCluster, sdkProvider, fwProvider, generationProvider, cluster.Configurators)
}

// GetProviderNamespaced returns the namespaced provider configuration. See
// GetProvider for the arguments.
func GetProviderNamespaced(sdkProvider *schema.Provider, fwProvider fwprovider.Provider, generationProvider bool) (*ujconfig.Provider, error) {
	return newProvider(rootGroupNamespaced, sdkProvider, fwProvider, generationProvider, namespaced.Configurators)
}

func newProvider(rootGroup string, sdkProvider *schema.Provider, fwProvider fwprovider.Provider, generationProvider bool, configurators []func(*ujconfig.Provider)) (*ujconfig.Provider, error) {
	if generationProvider {
		var err error
		sdkProvider, err = codegenSDKProvider(sdkProvider)
		if err != nil {
			return nil, err
		}
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup(rootGroup),
		// No resource runs through the Terraform CLI.
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithTerraformPluginFrameworkIncludeList(resourceList(terraformPluginFrameworkExternalNameConfigs)),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(resourceConfigurator()),
		ujconfig.WithControllerTemplate(templates.ControllerTemplate),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
	)
	for _, configure := range configurators {
		configure(pc)
	}
	if generationProvider {
		for name, configure := range codegenConfigurators {
			pc.AddResourceConfigurator(name, configure)
		}
	}
	pc.ConfigureResources()
	registerTerraformConversions(pc)
	return pc, nil
}

// codegenSDKProvider returns the plugin SDK provider used for code
// generation: its resources come from the JSON schema document, which keeps
// the CRD number types stable (the JSON schema does not distinguish integers
// from floats), with the MaxItems constraints synced from the live Go schema
// so that the singleton lists embedded in the CRDs are exactly the ones the
// runtime converts.
func codegenSDKProvider(sdkProvider *schema.Provider) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(providerSchema)); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal the JSON provider schema")
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should be exactly one provider schema but there are %d", len(ps.Schemas))
	}
	var resourceSchemas map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		resourceSchemas = v.ResourceSchemas
	}
	p := &schema.Provider{ResourcesMap: conversiontfjson.GetV2ResourceMap(resourceSchemas)}
	if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
		return nil, errors.Wrap(err, "cannot sync the MaxItems constraints from the Go schema to the JSON schema")
	}
	return p, nil
}

// registerTerraformConversions enables the runtime conversion between the
// embedded objects of the CRDs and the singleton lists Terraform expects on
// every resource the SingletonListEmbedder touched. The embedder only records
// the paths; without this conversion the objects reach the Terraform
// configuration and state unchanged and Terraform rejects them.
func registerTerraformConversions(pc *ujconfig.Provider) {
	for _, r := range pc.Resources {
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}
		r.TerraformConversions = append(r.TerraformConversions, ujconfig.NewTFSingletonConversion())
	}
}
