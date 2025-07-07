/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/google/uuid"
)

var terraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	"datadog_authn_mapping":                        config.IdentifierFromProvider,
	"datadog_child_organization":                   config.IdentifierFromProvider,
	"datadog_cloud_configuration_rule":             config.IdentifierFromProvider,
	"datadog_dashboard_json":                       config.IdentifierFromProvider,
	"datadog_downtime":                             config.IdentifierFromProvider,
	"datadog_integration_aws":                      config.IdentifierFromProvider,
	"datadog_integration_aws_lambda_arn":           config.IdentifierFromProvider,
	"datadog_integration_aws_log_collection":       config.IdentifierFromProvider,
	"datadog_integration_aws_tag_filter":           config.IdentifierFromProvider,
	"datadog_integration_opsgenie_service_object":  config.IdentifierFromProvider,
	"datadog_integration_pagerduty":                config.IdentifierFromProvider,
	"datadog_integration_pagerduty_service_object": config.IdentifierFromProvider,
	"datadog_integration_slack_channel":            config.IdentifierFromProvider,
	"datadog_logs_archive":                         config.IdentifierFromProvider,
	"datadog_logs_archive_order":                   config.IdentifierFromProvider,
	"datadog_logs_custom_pipeline":                 config.IdentifierFromProvider,
	"datadog_logs_index":                           config.NameAsIdentifier,
	"datadog_logs_index_order":                     config.NameAsIdentifier,
	"datadog_logs_integration_pipeline":            config.IdentifierFromProvider,
	"datadog_logs_metric":                          config.IdentifierFromProvider,
	"datadog_logs_pipeline_order":                  config.NameAsIdentifier,
	"datadog_metric_metadata":                      config.IdentifierFromProvider,
	"datadog_metric_tag_configuration":             config.IdentifierFromProvider,
	"datadog_monitor":                              config.IdentifierFromProvider,
	"datadog_monitor_config_policy":                config.IdentifierFromProvider,
	"datadog_monitor_json":                         config.IdentifierFromProvider,
	"datadog_organization_settings":                config.IdentifierFromProvider,
	"datadog_powerpack":                            config.IdentifierFromProvider,
	"datadog_role":                                 config.IdentifierFromProvider,
	"datadog_security_monitoring_default_rule":     config.IdentifierFromProvider,
	"datadog_security_monitoring_rule":             config.IdentifierFromProvider,
	"datadog_security_monitoring_filter":           config.IdentifierFromProvider,
	"datadog_sensitive_data_scanner_group":         config.IdentifierFromProvider,
	"datadog_sensitive_data_scanner_rule":          config.IdentifierFromProvider,
	"datadog_service_definition_yaml":              config.IdentifierFromProvider,
	"datadog_service_level_objective":              config.IdentifierFromProvider,
	"datadog_slo_correction":                       config.IdentifierFromProvider,
	"datadog_synthetics_test":                      config.IdentifierFromProvider,
	"datadog_user":                                 datadogExternalNameWithInjectedID(),
}

var terraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"datadog_api_key":                            datadogExternalNameWithInjectedID(),
	"datadog_apm_retention_filter":               datadogExternalNameWithInjectedID(),
	"datadog_apm_retention_filter_order":         config.IdentifierFromProvider,
	"datadog_application_key":                    datadogExternalNameWithInjectedID(),
	"datadog_dashboard_list":                     datadogExternalNameWithInjectedID(),
	"datadog_downtime_schedule":                  datadogExternalNameWithInjectedUUID(),
	"datadog_integration_aws_event_bridge":       config.IdentifierFromProvider,
	"datadog_integration_azure":                  config.IdentifierFromProvider,
	"datadog_integration_cloudflare_account":     config.IdentifierFromProvider,
	"datadog_integration_confluent_account":      config.IdentifierFromProvider,
	"datadog_integration_confluent_resource":     config.IdentifierFromProvider,
	"datadog_integration_fastly_account":         config.IdentifierFromProvider,
	"datadog_integration_fastly_service":         config.IdentifierFromProvider,
	"datadog_integration_gcp":                    config.IdentifierFromProvider,
	"datadog_integration_gcp_sts":                config.IdentifierFromProvider,
	"datadog_ip_allowlist":                       config.IdentifierFromProvider,
	"datadog_restriction_policy":                 config.IdentifierFromProvider,
	"datadog_rum_application":                    datadogExternalNameWithInjectedID(),
	"datadog_sensitive_data_scanner_group_order": config.IdentifierFromProvider,
	"datadog_service_account_application_key":    datadogExternalNameWithInjectedUUID(),
	"datadog_service_account":                    datadogExternalNameWithInjectedID(),
	"datadog_spans_metric":                       datadogExternalNameWithInjectedID(),
	"datadog_synthetics_concurrency_cap":         config.IdentifierFromProvider,
	"datadog_synthetics_global_variable":         datadogExternalNameWithInjectedID(),
	"datadog_synthetics_private_location":        datadogExternalNameWithInjectedID(),
	"datadog_team":                               datadogExternalNameWithInjectedID(),
	"datadog_team_link":                          datadogExternalNameWithInjectedID(),
	"datadog_team_membership":                    config.IdentifierFromProvider,
	"datadog_team_permission_setting":            config.TemplatedStringAsIdentifier("", "TeamPermission-{{ .parameters.team_id }}-{{ .parameters.action }}"),
	"datadog_webhook":                            config.IdentifierFromProvider,
	"datadog_webhook_custom_variable":            config.IdentifierFromProvider,
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var cliReconciledExternalNameConfigs = map[string]config.ExternalName{}

func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// If an external name is configured for multiple architectures,
		// Terraform Plugin Framework takes precedence over Terraform
		// Plugin SDKv2, which takes precedence over CLI architecture.
		e, configured := terraformPluginFrameworkExternalNameConfigs[r.Name]
		if !configured {
			e, configured = terraformPluginSDKExternalNameConfigs[r.Name]
			if !configured {
				e, configured = cliReconciledExternalNameConfigs[r.Name]
			}
		}
		if !configured {
			return
		}
		r.Version = "v1beta1"
		r.ExternalName = e
	}
}

// datadogExternalNameWithInjectedID injects an id when there is none.
// It is a slight modification of IdentifierFromProvider.
func datadogExternalNameWithInjectedID() config.ExternalName {
	// Terraform does not allow team id to be empty.
	// Using a stub value to pass validation.
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			// Some temporary id's need to be numeric
			return "0", nil
		}
		return externalName, nil
	}
	return e
}

// datadogExternalNameWithInjectedUUID injects an id when there is none.
// It is a slight modification of IdentifierFromProvider.
func datadogExternalNameWithInjectedUUID() config.ExternalName {
	// Terraform does not allow team id to be empty.
	// Using a stub value to pass validation.
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			// Some temporary id's need to be in UUID format
			return uuid.New().String(), nil
		}
		return externalName, nil
	}
	return e
}
