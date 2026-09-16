package cloudcost

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

const cloudCostDatadog = "cloudcost.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_aws_cur_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSCURConfig"
		r.ShortGroup = cloudCostDatadog
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "account_filters")
	})
	p.AddResourceConfigurator("datadog_azure_uc_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AzureUCConfig"
		r.ShortGroup = cloudCostDatadog
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "actual_bill_config", "amortized_bill_config")
		// These blocks have required descendants, which upjet's tfjson conversion
		// otherwise pushes into status.atProvider only; they are optional upstream.
		common.MarkSingleNestedBlockConfigurable(r, "actual_bill_config", true)
		common.MarkSingleNestedBlockConfigurable(r, "amortized_bill_config", true)
	})
	p.AddResourceConfigurator("datadog_cost_budget", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "CostBudget"
		r.ShortGroup = cloudCostDatadog
	})
	p.AddResourceConfigurator("datadog_cost_custom_forecast", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "CostCustomForecast"
		r.ShortGroup = cloudCostDatadog
	})
	p.AddResourceConfigurator("datadog_custom_allocation_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		//
		// Kind must not be "CustomAllocationRule": it pluralizes to the same
		// CRD resource name as "CustomAllocationRules" below, and only one of
		// the two CRDs gets written.
		r.Kind = "AllocationRule"
		r.ShortGroup = cloudCostDatadog
		// order_id is computed (API-assigned), not a settable input, so it is
		// not part of forProvider and cannot carry a reference.
		// strategy and strategy.based_on_timeseries are plugin-framework
		// single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "strategy", "strategy.based_on_timeseries")
	})
	p.AddResourceConfigurator("datadog_custom_allocation_rules", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "CustomAllocationRules"
		r.ShortGroup = cloudCostDatadog
		r.References["rule_ids"] = config.Reference{
			TerraformName: "datadog_custom_allocation_rule",
		}
	})
	p.AddResourceConfigurator("datadog_gcp_uc_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "GCPUCConfig"
		r.ShortGroup = cloudCostDatadog
	})
}
