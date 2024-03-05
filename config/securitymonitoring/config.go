package securitymonitoring

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_security_monitoring_default_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DefaultRule"
		r.ShortGroup = "securityMonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_monitoring_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Filter"
		r.ShortGroup = "securityMonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_monitoring_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Rule"
		r.ShortGroup = "securityMonitoring.datadog"
	})
}
