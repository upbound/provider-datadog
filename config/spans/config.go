package spans

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_spans_metric", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "SpansMetric"
		r.ShortGroup = "datadog"
		r.SchemaElementOptions.SetEmbeddedObject("compute")
		r.SchemaElementOptions.SetEmbeddedObject("filter")
	})
}
