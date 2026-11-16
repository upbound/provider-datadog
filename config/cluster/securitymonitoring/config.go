package securitymonitoring

import (
	"github.com/crossplane/upjet/v2/pkg/config"
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
	})
}
