package tag

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

const tagDatadog = "tag.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_tag_pipeline_ruleset", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		//
		// Kind must not be "PipelineRuleset": it pluralizes to the same CRD
		// resource name as "PipelineRulesets" below, and only one of the two
		// CRDs gets written.
		r.Kind = "TagPipelineRuleset"
		r.ShortGroup = tagDatadog
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "rules.mapping", "rules.query", "rules.query.addition", "rules.reference_table")
	})
	p.AddResourceConfigurator("datadog_tag_pipeline_rulesets", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "PipelineRulesets"
		r.ShortGroup = tagDatadog
		r.References["ruleset_ids"] = config.Reference{
			TerraformName: "datadog_tag_pipeline_ruleset",
		}
	})
	p.AddResourceConfigurator("datadog_tag_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Rule"
		r.ShortGroup = tagDatadog
	})
}
