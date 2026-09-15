/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/uuid"
)

// terraformPluginSDKExternalNameConfigs contains the external name
// configurations of the resources that the Terraform plugin SDK implements in
// the wrapped Datadog provider version. Their controllers use the plugin SDK
// client.
var terraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	"datadog_authn_mapping":                      config.IdentifierFromProvider,
	"datadog_child_organization":                 config.IdentifierFromProvider,
	"datadog_cloud_configuration_rule":           config.IdentifierFromProvider,
	"datadog_cloud_workload_security_agent_rule": config.IdentifierFromProvider,
	// "datadog_dashboard":                         config.IdentifierFromProvider,
	"datadog_dashboard_json":                       config.IdentifierFromProvider,
	"datadog_downtime":                             config.IdentifierFromProvider,
	"datadog_integration_opsgenie_service_object":  config.IdentifierFromProvider,
	"datadog_integration_pagerduty":                config.IdentifierFromProvider,
	"datadog_integration_pagerduty_service_object": config.IdentifierFromProvider,
	"datadog_integration_slack_channel":            config.IdentifierFromProvider,
	"datadog_logs_archive":                         config.IdentifierFromProvider,
	"datadog_logs_archive_order":                   config.IdentifierFromProvider,
	"datadog_logs_custom_pipeline":                 config.IdentifierFromProvider,
	"datadog_logs_index":                           config.IdentifierFromProvider,
	"datadog_logs_index_order":                     config.IdentifierFromProvider,
	"datadog_logs_integration_pipeline":            config.IdentifierFromProvider,
	"datadog_logs_metric":                          config.IdentifierFromProvider,
	"datadog_logs_pipeline_order":                  config.IdentifierFromProvider,
	"datadog_metric_metadata":                      config.IdentifierFromProvider,
	"datadog_metric_tag_configuration":             config.IdentifierFromProvider,
	// Upstream moves datadog_monitor to the plugin framework when the
	// TERRAFORM_MONITOR_FRAMEWORK_PROVIDER environment variable is "true".
	// The generated controller hard-wires the plugin SDK client, so leave the
	// variable unset in the provider container.
	"datadog_monitor":                      config.IdentifierFromProvider,
	"datadog_monitor_config_policy":        config.IdentifierFromProvider,
	"datadog_monitor_json":                 config.IdentifierFromProvider,
	"datadog_organization_settings":        config.IdentifierFromProvider,
	"datadog_powerpack":                    config.IdentifierFromProvider,
	"datadog_role":                         config.IdentifierFromProvider,
	"datadog_sensitive_data_scanner_group": config.IdentifierFromProvider,
	"datadog_sensitive_data_scanner_rule":  config.IdentifierFromProvider,
	"datadog_service_definition_yaml":      config.IdentifierFromProvider,
	"datadog_service_level_objective":      config.IdentifierFromProvider,
	"datadog_slo_correction":               config.IdentifierFromProvider,
	"datadog_synthetics_test":              config.IdentifierFromProvider,
	"datadog_user":                         datadogExternalNameWithInjectedID(),
}

// terraformPluginFrameworkExternalNameConfigs contains the external name
// configurations of the resources that the Terraform plugin framework
// implements in the wrapped Datadog provider version. Their controllers use
// the plugin framework client.
var terraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"datadog_api_key":                            datadogExternalNameWithInjectedUUID(),
	"datadog_apm_retention_filter":               datadogExternalNameWithInjectedUUID(),
	"datadog_apm_retention_filter_order":         config.IdentifierFromProvider,
	"datadog_application_key":                    datadogExternalNameWithInjectedID(),
	"datadog_dashboard_list":                     datadogExternalNameWithInjectedID(),
	"datadog_downtime_schedule":                  datadogExternalNameWithInjectedUUID(),
	"datadog_integration_aws_event_bridge":       config.IdentifierFromProvider,
	"datadog_integration_azure":                  config.IdentifierFromProvider,
	"datadog_integration_cloudflare_account":     datadogExternalNameWithInjectedUUID(),
	"datadog_integration_confluent_account":      datadogExternalNameWithInjectedUUID(),
	"datadog_integration_confluent_resource":     datadogExternalNameWithInjectedUUIDPair(),
	"datadog_integration_fastly_account":         datadogExternalNameWithInjectedUUID(),
	"datadog_integration_fastly_service":         datadogExternalNameWithInjectedUUIDPair(),
	"datadog_integration_gcp":                    config.IdentifierFromProvider,
	"datadog_integration_gcp_sts":                config.IdentifierFromProvider,
	"datadog_ip_allowlist":                       config.IdentifierFromProvider,
	"datadog_restriction_policy":                 config.IdentifierFromProvider,
	"datadog_rum_application":                    datadogExternalNameWithInjectedID(),
	"datadog_security_monitoring_default_rule":   config.IdentifierFromProvider,
	"datadog_security_monitoring_filter":         datadogExternalNameWithInjectedUUID(),
	"datadog_security_monitoring_rule":           datadogExternalNameWithInjectedUUID(),
	"datadog_sensitive_data_scanner_group_order": config.IdentifierFromProvider,
	"datadog_service_account":                    datadogExternalNameWithInjectedUUID(),
	"datadog_service_account_application_key":    config.IdentifierFromProvider,
	"datadog_spans_metric":                       datadogExternalNameWithInjectedID(),
	"datadog_synthetics_concurrency_cap":         config.IdentifierFromProvider,
	"datadog_synthetics_global_variable":         datadogExternalNameWithInjectedUUID(),
	"datadog_synthetics_private_location":        datadogExternalNameWithInjectedUUID(),
	"datadog_team":                               datadogExternalNameWithInjectedID(),
	"datadog_team_link":                          datadogExternalNameWithInjectedID(),
	"datadog_team_membership":                    config.IdentifierFromProvider,
	// The framework resource resolves an empty id from team_id and action and
	// stores the API identifier of the setting, which exists as soon as its
	// team does; a stub id only produces a not-found error.
	"datadog_team_permission_setting": config.IdentifierFromProvider,
	"datadog_webhook":                 config.IdentifierFromProvider,
	"datadog_webhook_custom_variable": config.IdentifierFromProvider,
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

// datadogExternalNameWithInjectedUUIDPair is for resources whose Terraform id
// is "<parent uuid>:<child uuid>", such as the Confluent and Fastly
// integration children. The stub keeps the two-part shape so the provider's
// Read can split it and turn the API's 404 into a plain "does not exist".
func datadogExternalNameWithInjectedUUIDPair() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			return uuid.New().String() + ":" + uuid.New().String(), nil
		}
		return externalName, nil
	}
	return e
}

// resourceConfigurator applies the external name configuration of a resource
// from the table of the Terraform plugin library that implements it.
func resourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := terraformPluginSDKExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
			return
		}
		if e, ok := terraformPluginFrameworkExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// resourceList returns the include list that matches exactly the resources of
// an external name configuration table.
func resourceList(configs map[string]config.ExternalName) []string {
	l := make([]string, 0, len(configs))
	for name := range configs {
		l = append(l, "^"+name+"$")
	}
	return l
}
