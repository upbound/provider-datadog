package compliance

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_compliance_custom_framework", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "CustomFramework"
		r.ShortGroup = "compliance.datadog"
	})
	p.AddResourceConfigurator("datadog_compliance_resource_evaluation_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ResourceEvaluationFilter"
		r.ShortGroup = "compliance.datadog"
	})
}
