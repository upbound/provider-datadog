package dashboard

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_dashboard", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Dashboard"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_dashboard_json", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DashboardJSON"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_dashboard_list", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DashboardList"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_secure_embed_dashboard", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "SecureEmbedDashboard"
		r.ShortGroup = "datadog"
		r.References["dashboard_id"] = config.Reference{
			TerraformName: "datadog_dashboard_json",
		}
	})
}
