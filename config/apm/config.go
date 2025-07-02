package apm

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_apm_retention_filter", func(r *config.Resource) {
		r.RemoveSingletonListConversion("filter")
		r.SchemaElementOptions.SetEmbeddedObject("filter")
	})
	p.AddResourceConfigurator("datadog_apm_retention_filter_order", func(r *config.Resource) {
		r.References["filter_ids"] = config.Reference{
			TerraformName: "datadog_apm_retention_filter",
		}
	})
}
