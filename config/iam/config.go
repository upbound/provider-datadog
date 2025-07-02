package iam

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_service_account", func(r *config.Resource) {
		r.References["roles"] = config.Reference{
			TerraformName: "datadog_role",
		}
	})
	p.AddResourceConfigurator("datadog_service_account_application_key", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			TerraformName: "datadog_service_account",
		}
	})
	p.AddResourceConfigurator("datadog_team_link", func(r *config.Resource) {
		r.References["team_id"] = config.Reference{
			TerraformName: "datadog_team",
		}
	})
	p.AddResourceConfigurator("datadog_team_membership", func(r *config.Resource) {
		r.References["team_id"] = config.Reference{
			TerraformName: "datadog_team",
		}
		r.References["user_id"] = config.Reference{
			TerraformName: "datadog_user",
		}
	})
	p.AddResourceConfigurator("datadog_team_permission_setting", func(r *config.Resource) {
		r.References["team_id"] = config.Reference{
			TerraformName: "datadog_team",
		}
	})
	p.AddResourceConfigurator("datadog_user", func(r *config.Resource) {
		r.References["roles"] = config.Reference{
			TerraformName: "datadog_role",
		}
	})
}
