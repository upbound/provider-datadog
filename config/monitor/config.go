package monitor

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_monitor", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Monitor"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_monitor_config_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "MonitorConfigPolicy"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_monitor_json", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "MonitorJSON"
		r.ShortGroup = "datadog"
	})
}
