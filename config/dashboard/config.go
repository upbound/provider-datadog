package dashboard

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/comments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_dashboard", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Dashboard"
		r.ShortGroup = "datadog"

		delete(r.TerraformResource.Schema, "widget")

		desc, _ := comments.New("JSON widget to add to a dashboard",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["widget"] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: desc.String(),
		}
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
}
