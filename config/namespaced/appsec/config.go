package appsec

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_appsec_waf_custom_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "WAFCustomRule"
		r.ShortGroup = "appsec.datadog"
		// action, action.parameters, condition.parameters, and
		// condition.parameters.options are plugin-framework single-nested
		// blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "action", "action.parameters", "condition.parameters", "condition.parameters.options")
	})
	p.AddResourceConfigurator("datadog_appsec_waf_exclusion_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "WAFExclusionFilter"
		r.ShortGroup = "appsec.datadog"
		// rules_target.tags is a plugin-framework single-nested block; see
		// common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "rules_target.tags")
	})
}
