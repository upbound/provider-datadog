package monitor

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

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
	p.AddResourceConfigurator("datadog_monitor_notification_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "MonitorNotificationRule"
		r.ShortGroup = "datadog"
		// bundle_config, conditional_recipients, and filter are
		// plugin-framework single-nested blocks; see
		// common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "bundle_config", "conditional_recipients", "filter")
		// conditional_recipients has a required descendant, which upjet's tfjson conversion
		// otherwise pushes into status.atProvider only; optional upstream.
		common.MarkSingleNestedBlockConfigurable(r, "conditional_recipients", true)
	})
}
