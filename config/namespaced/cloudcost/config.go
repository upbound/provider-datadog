package cloudcost

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const cloudCostDatadog = "cloudcost.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_aws_cur_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSCURConfig"
		r.ShortGroup = cloudCostDatadog
	})
	p.AddResourceConfigurator("datadog_azure_uc_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AzureUCConfig"
		r.ShortGroup = cloudCostDatadog
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
