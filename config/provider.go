/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/DataDog/terraform-provider-datadog/v3/datadog"
	"github.com/DataDog/terraform-provider-datadog/v3/datadog/fwprovider"
	"github.com/crossplane/crossplane-runtime/pkg/errors"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/provider-datadog/config/iam"
	"github.com/upbound/provider-datadog/config/integration"
	"github.com/upbound/provider-datadog/config/logs"
	"github.com/upbound/provider-datadog/config/sensitivedata"
	"github.com/upbound/provider-datadog/config/slo"

	"github.com/upbound/provider-datadog/config/apm"
	"github.com/upbound/provider-datadog/config/authentication"
	"github.com/upbound/provider-datadog/config/dashboard"
	"github.com/upbound/provider-datadog/config/metric"
	"github.com/upbound/provider-datadog/config/monitor"
	"github.com/upbound/provider-datadog/config/security"
	"github.com/upbound/provider-datadog/config/synthetics"
)

const (
	resourcePrefix = "datadog"
	modulePath     = "github.com/upbound/provider-datadog"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(generationProvider bool) (*ujconfig.Provider, error) {
	sdkProvider := datadog.Provider()
	fwProvider := fwprovider.New()

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("datadog.upbound.io"),
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithTerraformPluginFrameworkIncludeList(resourceList(terraformPluginFrameworkExternalNameConfigs)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
		ujconfig.WithDefaultResourceOptions(
			ResourceConfigurator(),
			GroupKindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apm.Configure,
		dashboard.Configure,
		metric.Configure,
		monitor.Configure,
		security.Configure,
		synthetics.Configure,
		authentication.Configure,
		iam.Configure,
		integration.Configure,
		logs.Configure,
		sensitivedata.Configure,
		slo.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	registerTFConversions(pc)
	return pc, nil
}

func registerTFConversions(pc *ujconfig.Provider) {
	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		// the controller will be reconciling on the CRD API version
		// with the converted API (with embedded objects in place of
		// singleton lists), so we need the appropriate Terraform
		// converter in this case.
		r.TerraformConversions = []ujconfig.TerraformConversion{
			ujconfig.NewTFSingletonConversion(),
		}

		pc.Resources[name] = r
	}
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]ujconfig.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}
