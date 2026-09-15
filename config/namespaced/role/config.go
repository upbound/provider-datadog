package role

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Role"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_user_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "UserRole"
		r.ShortGroup = "datadog"
		r.References["role_id"] = config.Reference{
			TerraformName: "datadog_role",
		}
		r.References["user_id"] = config.Reference{
			TerraformName: "datadog_user",
		}
	})
}
