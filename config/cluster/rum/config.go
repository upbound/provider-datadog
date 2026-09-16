package rum

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_rum_application", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RUMApplication"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_rum_exclusion_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RUMExclusionFilter"
		r.ShortGroup = "datadog"
		r.References["application_id"] = config.Reference{
			TerraformName: "datadog_rum_application",
		}
	})
	p.AddResourceConfigurator("datadog_rum_metric", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RUMMetric"
		r.ShortGroup = "datadog"
		// compute, filter, and uniqueness are plugin-framework
		// single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "compute", "filter", "uniqueness")
		// compute has a required descendant, which upjet's tfjson conversion
		// otherwise pushes into status.atProvider only; required upstream.
		common.MarkSingleNestedBlockConfigurable(r, "compute", false)
	})
	p.AddResourceConfigurator("datadog_rum_retention_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RUMRetentionFilter"
		r.ShortGroup = "datadog"
		r.References["application_id"] = config.Reference{
			TerraformName: "datadog_rum_application",
		}
	})
	p.AddResourceConfigurator("datadog_rum_retention_quota", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RUMRetentionQuota"
		r.ShortGroup = "datadog"
		r.References["application_id"] = config.Reference{
			TerraformName: "datadog_rum_application",
		}
		// custom has required descendants, which upjet's tfjson conversion
		// otherwise pushes into status.atProvider only; see
		// common.MarkSingleNestedBlockConfigurable. It is also a
		// plugin-framework single-nested block; see
		// common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "custom")
		common.MarkSingleNestedBlockConfigurable(r, "custom", true)
	})
}
