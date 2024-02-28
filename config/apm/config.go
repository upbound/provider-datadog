package apm

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_apm_retention_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RetentionFilter"
		r.ShortGroup = "apm.datadog"
		r.SchemaElementOptions.SetEmbeddedObject("filter")
	})
	p.AddResourceConfigurator("datadog_apm_retention_filter_order", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RetentionFilterOrder"
		r.ShortGroup = "apm.datadog"
	})
}
