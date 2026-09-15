/*
Copyright 2026 Upbound Inc.
*/

package integration

import "github.com/crossplane/upjet/v2/pkg/config"

const integrationDatadog = "integration.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_integration_aws_event_bridge", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSEventBridge"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_azure", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Azure"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_cloudflare_account", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "CloudflareAccount"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_confluent_account", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ConfluentAccount"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_confluent_resource", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ConfluentResource"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_fastly_account", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "FastlyAccount"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_fastly_service", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "FastlyService"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_gcp", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "GCP"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_gcp_sts", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "GCPSTS"
		r.ShortGroup = integrationDatadog
		// The API fills both sets on its own (Prometheus disabled by default,
		// monitored resources derived from the legacy filters). Copying them
		// into spec makes the next apply fail with "Provider produced
		// inconsistent result" on metric_namespace_configs.
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"metric_namespace_configs", "monitored_resource_configs"},
		}
	})
	p.AddResourceConfigurator("datadog_integration_opsgenie_service_object", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "OpsgenieServiceObject"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_pagerduty", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Pagerduty"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_pagerduty_service_object", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "PagerdutyServiceObject"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_slack_channel", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "SlackChannel"
		r.ShortGroup = integrationDatadog
	})
}
