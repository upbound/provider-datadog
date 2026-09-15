package statuspage

import "github.com/crossplane/upjet/v2/pkg/config"

const statusPageDatadog = "statuspage.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_status_page", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Page"
		r.ShortGroup = statusPageDatadog
	})
	p.AddResourceConfigurator("datadog_status_page_component", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "PageComponent"
		r.ShortGroup = statusPageDatadog
		r.References["page_id"] = config.Reference{
			TerraformName: "datadog_status_page",
		}
	})
	p.AddResourceConfigurator("datadog_status_page_degradation_template", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "PageDegradationTemplate"
		r.ShortGroup = statusPageDatadog
		r.References["page_id"] = config.Reference{
			TerraformName: "datadog_status_page",
		}
	})
	p.AddResourceConfigurator("datadog_status_page_maintenance_template", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "PageMaintenanceTemplate"
		r.ShortGroup = statusPageDatadog
		r.References["page_id"] = config.Reference{
			TerraformName: "datadog_status_page",
		}
		r.References["component_ids"] = config.Reference{
			TerraformName: "datadog_status_page_component",
		}
	})
}
