package logs

import "github.com/crossplane/upjet/pkg/config"

const logsDatadog = "logs.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_logs_archive", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
	p.AddResourceConfigurator("datadog_logs_archive_order", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
	p.AddResourceConfigurator("datadog_logs_custom_pipeline", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
	p.AddResourceConfigurator("datadog_logs_index", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
	p.AddResourceConfigurator("datadog_logs_index_order", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
	p.AddResourceConfigurator("datadog_logs_integration_pipeline", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
	p.AddResourceConfigurator("datadog_logs_metric", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
	p.AddResourceConfigurator("datadog_logs_pipeline_order", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = logsDatadog
	})
}
