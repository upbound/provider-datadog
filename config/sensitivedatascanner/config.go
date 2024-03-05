package sensitivedatascanner

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_sensitive_data_scanner_group", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Group"
		r.ShortGroup = "sensitiveDataScanner.datadog"
	})
	p.AddResourceConfigurator("datadog_sensitive_data_scanner_group_order", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "GroupOrder"
		r.ShortGroup = "sensitiveDataScanner.datadog"
	})
	p.AddResourceConfigurator("datadog_sensitive_data_scanner_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Rule"
		r.ShortGroup = "sensitiveDataScanner.datadog"
	})
}
