package cloud

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_cloud_configuration_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "cloud.datadog"
	})
	p.AddResourceConfigurator("datadog_cloud_inventory_sync_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "CloudInventorySyncConfig"
		r.ShortGroup = "cloud.datadog"
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "aws", "azure", "gcp")
	})
	p.AddResourceConfigurator("datadog_cloud_workload_security_agent_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "cloud.datadog"
	})
}
