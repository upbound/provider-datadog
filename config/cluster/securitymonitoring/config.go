package securitymonitoring

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_security_monitoring_critical_asset", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "CriticalAsset"
		r.ShortGroup = "securitymonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_monitoring_default_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DefaultRule"
		r.ShortGroup = "securitymonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_monitoring_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Filter"
		r.ShortGroup = "securitymonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_monitoring_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Rule"
		r.ShortGroup = "securitymonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_monitoring_rule_json", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "RuleJSON"
		r.ShortGroup = "securitymonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_monitoring_suppression", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Suppression"
		r.ShortGroup = "securitymonitoring.datadog"
	})
	p.AddResourceConfigurator("datadog_security_notification_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "NotificationRule"
		r.ShortGroup = "securitymonitoring.datadog"
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "selectors")
		// selectors has a required descendant, which upjet's tfjson conversion
		// otherwise pushes into status.atProvider only; required upstream (objectvalidator.IsRequired).
		common.MarkSingleNestedBlockConfigurable(r, "selectors", false)
	})
}
