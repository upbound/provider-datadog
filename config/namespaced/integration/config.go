/*
Copyright 2026 Upbound Inc.
*/

package integration

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

const integrationDatadog = "integration.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_integration_aws_account", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSAccount"
		r.ShortGroup = integrationDatadog
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "auth_config", "auth_config.aws_auth_config_keys", "auth_config.aws_auth_config_role", "aws_regions", "logs_config", "logs_config.lambda_forwarder")
		common.EmbedSingleNestedBlocks(r, "logs_config.lambda_forwarder.log_source_config", "metrics_config", "metrics_config.namespace_filters", "resources_config", "traces_config", "traces_config.xray_services")
		// These blocks have required descendants, which upjet's tfjson conversion
		// otherwise pushes into status.atProvider only; they are optional upstream.
		common.MarkSingleNestedBlockConfigurable(r, "logs_config", true)
		common.MarkSingleNestedBlockConfigurable(r, "logs_config.lambda_forwarder", true)
		common.MarkSingleNestedBlockConfigurable(r, "logs_config.lambda_forwarder.log_source_config", true)
		common.MarkSingleNestedBlockConfigurable(r, "metrics_config", true)
	})
	p.AddResourceConfigurator("datadog_integration_aws_account_ccm_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSAccountCCMConfig"
		r.ShortGroup = integrationDatadog
		r.References["aws_account_config_id"] = config.Reference{
			TerraformName: "datadog_integration_aws_account",
		}
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "ccm_config")
	})
	p.AddResourceConfigurator("datadog_integration_aws_event_bridge", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSEventBridge"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_aws_external_id", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSExternalID"
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
	p.AddResourceConfigurator("datadog_integration_ms_teams_tenant_based_handle", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "MSTeamsTenantBasedHandle"
		r.ShortGroup = integrationDatadog
	})
	p.AddResourceConfigurator("datadog_integration_ms_teams_workflows_webhook_handle", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "MSTeamsWorkflowsWebhookHandle"
		r.ShortGroup = integrationDatadog
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
