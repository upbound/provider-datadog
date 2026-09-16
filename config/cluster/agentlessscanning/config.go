package agentlessscanning

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_agentless_scanning_aws_scan_options", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSScanOptions"
		r.ShortGroup = "agentlessscanning.datadog"
	})
	p.AddResourceConfigurator("datadog_agentless_scanning_azure_scan_options", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AzureScanOptions"
		r.ShortGroup = "agentlessscanning.datadog"
	})
	p.AddResourceConfigurator("datadog_agentless_scanning_gcp_scan_options", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "GCPScanOptions"
		r.ShortGroup = "agentlessscanning.datadog"
	})
}
