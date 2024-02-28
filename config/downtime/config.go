package downtime

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_downtime", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_downtime_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DowntimeSchedule"
		r.ShortGroup = "datadog"
		r.SchemaElementOptions.SetEmbeddedObject("monitor_identifier")
		r.SchemaElementOptions.SetEmbeddedObject("one_time_schedule")
		r.SchemaElementOptions.SetEmbeddedObject("recurring_schedule")
	})
}
