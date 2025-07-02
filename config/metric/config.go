package metric

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_spans_metric", func(r *config.Resource) {
		r.RemoveSingletonListConversion("compute")
		r.RemoveSingletonListConversion("filter")
		r.SchemaElementOptions.SetEmbeddedObject("compute")
		r.SchemaElementOptions.SetEmbeddedObject("filter")
	})
}
