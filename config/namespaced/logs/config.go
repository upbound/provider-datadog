package logs

import "github.com/crossplane/upjet/v2/pkg/config"

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
		// The logs index API serves a new index about a second after the
		// create call returns. An async create keeps the new index name only
		// in memory until the first successful observation, so the not-found
		// read in that window drops it and the index is created again under a
		// name Datadog never releases. A sync create persists the external
		// name before the next observation.
		r.UseAsync = false
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
