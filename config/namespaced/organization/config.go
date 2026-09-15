package organization

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_organization_settings", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "OrganizationSettings"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_org_connection", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "OrgConnection"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_org_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "OrgGroup"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_org_group_membership", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "OrgGroupMembership"
		r.ShortGroup = "datadog"
		r.References["org_group_id"] = config.Reference{
			TerraformName: "datadog_org_group",
		}
	})
	p.AddResourceConfigurator("datadog_org_group_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "OrgGroupPolicy"
		r.ShortGroup = "datadog"
		r.References["org_group_id"] = config.Reference{
			TerraformName: "datadog_org_group",
		}
	})
	p.AddResourceConfigurator("datadog_org_group_policy_override", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "OrgGroupPolicyOverride"
		r.ShortGroup = "datadog"
		r.References["org_group_id"] = config.Reference{
			TerraformName: "datadog_org_group",
		}
		r.References["policy_id"] = config.Reference{
			TerraformName: "datadog_org_group_policy",
		}
	})
}
