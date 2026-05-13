/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/google/uuid"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"datadog_action_connection":                             config.IdentifierFromProvider,
	"datadog_agentless_scanning_aws_scan_options":           config.IdentifierFromProvider,
	"datadog_agentless_scanning_gcp_scan_options":           config.IdentifierFromProvider,
	"datadog_api_key":                                       datadogExternalNameWithInjectedID(),
	"datadog_apm_retention_filter":                          config.IdentifierFromProvider,
	"datadog_apm_retention_filter_order":                    config.IdentifierFromProvider,
	"datadog_app_builder_app":                               config.IdentifierFromProvider,
	"datadog_app_key_registration":                          config.IdentifierFromProvider,
	"datadog_application_key":                               datadogExternalNameWithInjectedID(),
	"datadog_appsec_waf_custom_rule":                        config.IdentifierFromProvider,
	"datadog_appsec_waf_exclusion_filter":                   config.IdentifierFromProvider,
	"datadog_authn_mapping":                                 config.IdentifierFromProvider,
	"datadog_aws_cur_config":                                config.IdentifierFromProvider,
	"datadog_azure_uc_config":                               config.IdentifierFromProvider,
	"datadog_child_organization":                            config.IdentifierFromProvider,
	"datadog_cloud_configuration_rule":                      config.IdentifierFromProvider,
	"datadog_cloud_inventory_sync_config":                   config.IdentifierFromProvider,
	"datadog_cloud_workload_security_agent_rule":            config.IdentifierFromProvider,
	"datadog_compliance_custom_framework":                   config.IdentifierFromProvider,
	"datadog_compliance_resource_evaluation_filter":         config.IdentifierFromProvider,
	"datadog_cost_budget":                                   config.IdentifierFromProvider,
	"datadog_csm_threats_agent_rule":                        config.IdentifierFromProvider,
	"datadog_csm_threats_policy":                            config.IdentifierFromProvider,
	"datadog_custom_allocation_rule":                        config.IdentifierFromProvider,
	"datadog_custom_allocation_rules":                       config.IdentifierFromProvider,
	"datadog_dashboard":                                     config.IdentifierFromProvider,
	"datadog_dashboard_json":                                config.IdentifierFromProvider,
	"datadog_dashboard_list":                                datadogExternalNameWithInjectedID(),
	"datadog_dataset":                                       config.IdentifierFromProvider,
	"datadog_datastore":                                     config.IdentifierFromProvider,
	"datadog_datastore_item":                                config.IdentifierFromProvider,
	"datadog_deployment_gate":                               config.IdentifierFromProvider,
	"datadog_domain_allowlist":                              config.IdentifierFromProvider,
	"datadog_downtime":                                      config.IdentifierFromProvider,
	"datadog_downtime_schedule":                             datadogExternalNameWithInjectedUUID(),
	"datadog_gcp_uc_config":                                 config.IdentifierFromProvider,
	"datadog_incident_notification_rule":                    config.IdentifierFromProvider,
	"datadog_incident_notification_template":                config.IdentifierFromProvider,
	"datadog_incident_type":                                 config.IdentifierFromProvider,
	"datadog_integration_aws":                               config.IdentifierFromProvider,
	"datadog_integration_aws_account":                       config.IdentifierFromProvider,
	"datadog_integration_aws_event_bridge":                  config.IdentifierFromProvider,
	"datadog_integration_aws_external_id":                   config.IdentifierFromProvider,
	"datadog_integration_aws_lambda_arn":                    config.IdentifierFromProvider,
	"datadog_integration_aws_log_collection":                config.IdentifierFromProvider,
	"datadog_integration_aws_tag_filter":                    config.IdentifierFromProvider,
	"datadog_integration_azure":                             config.IdentifierFromProvider,
	"datadog_integration_cloudflare_account":                config.IdentifierFromProvider,
	"datadog_integration_confluent_account":                 config.IdentifierFromProvider,
	"datadog_integration_confluent_resource":                config.IdentifierFromProvider,
	"datadog_integration_fastly_account":                    config.IdentifierFromProvider,
	"datadog_integration_fastly_service":                    config.IdentifierFromProvider,
	"datadog_integration_gcp":                               config.IdentifierFromProvider,
	"datadog_integration_gcp_sts":                           config.IdentifierFromProvider,
	"datadog_integration_ms_teams_tenant_based_handle":      config.IdentifierFromProvider,
	"datadog_integration_ms_teams_workflows_webhook_handle": config.IdentifierFromProvider,
	"datadog_integration_opsgenie_service_object":           config.IdentifierFromProvider,
	"datadog_integration_pagerduty":                         config.IdentifierFromProvider,
	"datadog_integration_pagerduty_service_object":          config.IdentifierFromProvider,
	"datadog_integration_slack_channel":                     config.IdentifierFromProvider,
	"datadog_ip_allowlist":                                  config.IdentifierFromProvider,
	"datadog_logs_archive":                                  config.IdentifierFromProvider,
	"datadog_logs_archive_order":                            config.IdentifierFromProvider,
	"datadog_logs_custom_destination":                       config.IdentifierFromProvider,
	"datadog_logs_custom_pipeline":                          config.IdentifierFromProvider,
	"datadog_logs_index":                                    config.IdentifierFromProvider,
	"datadog_logs_index_order":                              config.IdentifierFromProvider,
	"datadog_logs_integration_pipeline":                     config.IdentifierFromProvider,
	"datadog_logs_metric":                                   config.IdentifierFromProvider,
	"datadog_logs_pipeline_order":                           config.IdentifierFromProvider,
	"datadog_logs_restriction_query":                        config.IdentifierFromProvider,
	"datadog_metric_metadata":                               config.IdentifierFromProvider,
	"datadog_metric_tag_configuration":                      config.IdentifierFromProvider,
	"datadog_monitor":                                       config.IdentifierFromProvider,
	"datadog_monitor_config_policy":                         config.IdentifierFromProvider,
	"datadog_monitor_json":                                  config.IdentifierFromProvider,
	"datadog_monitor_notification_rule":                     config.IdentifierFromProvider,
	"datadog_observability_pipeline":                        config.IdentifierFromProvider,
	"datadog_on_call_escalation_policy":                     config.IdentifierFromProvider,
	"datadog_on_call_schedule":                              config.IdentifierFromProvider,
	"datadog_on_call_team_routing_rules":                    config.IdentifierFromProvider,
	"datadog_on_call_user_notification_channel":             config.IdentifierFromProvider,
	"datadog_on_call_user_notification_rule":                config.IdentifierFromProvider,
	"datadog_openapi_api":                                   config.IdentifierFromProvider,
	"datadog_org_connection":                                config.IdentifierFromProvider,
	"datadog_organization_settings":                         config.IdentifierFromProvider,
	"datadog_powerpack":                                     config.IdentifierFromProvider,
	"datadog_reference_table":                               config.IdentifierFromProvider,
	"datadog_restriction_policy":                            config.IdentifierFromProvider,
	"datadog_role":                                          config.IdentifierFromProvider,
	"datadog_rum_application":                               datadogExternalNameWithInjectedID(),
	"datadog_rum_metric":                                    config.IdentifierFromProvider,
	"datadog_rum_retention_filter":                          config.IdentifierFromProvider,
	"datadog_rum_retention_filters_order":                   config.IdentifierFromProvider,
	"datadog_security_monitoring_critical_asset":            config.IdentifierFromProvider,
	"datadog_security_monitoring_default_rule":              config.IdentifierFromProvider,
	"datadog_security_monitoring_filter":                    config.IdentifierFromProvider,
	"datadog_security_monitoring_rule":                      config.IdentifierFromProvider,
	"datadog_security_monitoring_rule_json":                 config.IdentifierFromProvider,
	"datadog_security_monitoring_suppression":               config.IdentifierFromProvider,
	"datadog_security_notification_rule":                    config.IdentifierFromProvider,
	"datadog_sensitive_data_scanner_group":                  config.IdentifierFromProvider,
	"datadog_sensitive_data_scanner_group_order":            config.IdentifierFromProvider,
	"datadog_sensitive_data_scanner_rule":                   config.IdentifierFromProvider,
	"datadog_service_account":                               config.IdentifierFromProvider,
	"datadog_service_account_application_key":               config.IdentifierFromProvider,
	"datadog_service_definition_yaml":                       config.IdentifierFromProvider,
	"datadog_service_level_objective":                       config.IdentifierFromProvider,
	"datadog_slo_correction":                                config.IdentifierFromProvider,
	"datadog_software_catalog":                              config.IdentifierFromProvider,
	"datadog_spans_metric":                                  datadogExternalNameWithInjectedID(),
	"datadog_synthetics_concurrency_cap":                    config.IdentifierFromProvider,
	"datadog_synthetics_global_variable":                    config.IdentifierFromProvider,
	"datadog_synthetics_private_location":                   config.IdentifierFromProvider,
	"datadog_synthetics_suite":                              config.IdentifierFromProvider,
	"datadog_synthetics_test":                               config.IdentifierFromProvider,
	"datadog_tag_pipeline_ruleset":                          config.IdentifierFromProvider,
	"datadog_tag_pipeline_rulesets":                         config.IdentifierFromProvider,
	"datadog_team":                                          datadogExternalNameWithInjectedID(),
	"datadog_team_hierarchy_links":                          config.IdentifierFromProvider,
	"datadog_team_link":                                     datadogExternalNameWithInjectedID(),
	"datadog_team_membership":                               config.IdentifierFromProvider,
	"datadog_team_notification_rule":                        config.IdentifierFromProvider,
	"datadog_team_permission_setting":                       datadogExternalNameWithInjectedID(),
	"datadog_user":                                          datadogExternalNameWithInjectedID(),
	"datadog_user_role":                                     config.IdentifierFromProvider,
	"datadog_webhook":                                       config.IdentifierFromProvider,
	"datadog_webhook_custom_variable":                       config.IdentifierFromProvider,
	"datadog_workflow_automation":                           config.IdentifierFromProvider,
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

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
