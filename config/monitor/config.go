package monitor

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_downtime_schedule", func(r *config.Resource) {
		r.RemoveSingletonListConversion("monitor_identifier")
		r.RemoveSingletonListConversion("one_time_schedule")
		r.RemoveSingletonListConversion("recurring_schedule")
		r.SchemaElementOptions.SetEmbeddedObject("monitor_identifier")
		r.SchemaElementOptions.SetEmbeddedObject("one_time_schedule")
		r.SchemaElementOptions.SetEmbeddedObject("recurring_schedule")
	})

	p.AddResourceConfigurator("datadog_downtime", func(r *config.Resource) {
		r.References["monitor_id"] = config.Reference{
			TerraformName: "datadog_monitor",
		}
	})
}
