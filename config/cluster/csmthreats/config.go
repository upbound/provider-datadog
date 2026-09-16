package csmthreats

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_csm_threats_agent_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AgentRule"
		r.ShortGroup = "csmthreats.datadog"
		r.References["policy_id"] = config.Reference{
			TerraformName: "datadog_csm_threats_policy",
		}
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "actions.hash", "actions.set")
	})
	p.AddResourceConfigurator("datadog_csm_threats_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Policy"
		r.ShortGroup = "csmthreats.datadog"
	})
}
