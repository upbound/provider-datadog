package synthetics

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_synthetics_concurrency_cap", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
	p.AddResourceConfigurator("datadog_synthetics_global_variable", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
	p.AddResourceConfigurator("datadog_synthetics_private_location", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
	p.AddResourceConfigurator("datadog_synthetics_test", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
}
