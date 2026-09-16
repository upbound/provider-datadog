/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	"context"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
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
	// The widget block is also nested inside every "tab" block, so the
	// JSON-string collapse that works for datadog_powerpack does not fully
	// solve the CRD size problem here; see datadog_dashboard above.
	// "datadog_dashboard_v2":                        config.IdentifierFromProvider,
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
	"datadog_powerpack_v2":                 config.IdentifierFromProvider,
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
	"datadog_action_connection":                     datadogExternalNameWithInjectedUUID(),
	"datadog_agentless_scanning_aws_scan_options":   datadogExternalNameWithInjectedUUID(),
	"datadog_agentless_scanning_azure_scan_options": datadogExternalNameWithInjectedUUID(),
	"datadog_agentless_scanning_gcp_scan_options":   datadogExternalNameWithInjectedUUID(),
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"datadog_api_key":                    datadogExternalNameWithInjectedUUID(),
	"datadog_apm_retention_filter":       datadogExternalNameWithInjectedUUID(),
	"datadog_apm_retention_filter_order": config.IdentifierFromProvider,
	"datadog_app_builder_app":            datadogExternalNameWithInjectedUUID(),
	// id is the id of an existing Application Key; see the reference on this
	// field in the resource configurator.
	"datadog_app_key_registration":                  config.TemplatedStringAsIdentifier("id", "{{ .parameters.id }}"),
	"datadog_application_key":                       datadogExternalNameWithInjectedID(),
	"datadog_appsec_waf_custom_rule":                datadogExternalNameWithInjectedUUID(),
	"datadog_appsec_waf_exclusion_filter":           datadogExternalNameWithInjectedUUID(),
	"datadog_aws_cur_config":                        datadogExternalNameWithInjectedUUID(),
	"datadog_azure_uc_config":                       datadogExternalNameWithInjectedUUID(),
	"datadog_cloud_inventory_sync_config":           datadogExternalNameWithInjectedUUID(),
	"datadog_compliance_custom_framework":           datadogExternalNameWithInjectedUUID(),
	"datadog_compliance_resource_evaluation_filter": datadogExternalNameWithInjectedUUID(),
	// id is optional and computed upstream (a caller may supply one); an
	// omitted id is sent as unset, not empty, so the plain stub works.
	"datadog_cost_budget":             datadogExternalNameWithInjectedUUIDAndNotFoundDiagnostics("404 Not Found"),
	"datadog_cost_custom_forecast":    datadogExternalNameWithInjectedUUID(),
	"datadog_csm_threats_agent_rule":  datadogExternalNameWithInjectedUUID(),
	"datadog_csm_threats_policy":      datadogExternalNameWithInjectedUUID(),
	"datadog_custom_allocation_rule":  datadogExternalNameWithInjectedID(),
	"datadog_custom_allocation_rules": datadogExternalNameWithInjectedUUID(),
	"datadog_dashboard_list":          datadogExternalNameWithInjectedID(),
	"datadog_dataset":                 datadogExternalNameWithInjectedUUID(),
	"datadog_datastore":               datadogExternalNameWithInjectedUUID(),
	// The Terraform id is "<datastore id>:<item key>" and the external name is
	// the item key alone. The template must reference .external_name: with
	// .parameters.item_key instead, upjet takes the whole id as the external
	// name after Create, feeds it back as item_key on the next reconcile and
	// the update fails with "failed to find item with key <id>:<key>". See
	// the reference on datastore_id in the resource configurator.
	"datadog_datastore_item":                                config.TemplatedStringAsIdentifier("item_key", "{{ .parameters.datastore_id }}:{{ .external_name }}"),
	"datadog_domain_allowlist":                              datadogExternalNameWithInjectedUUID(),
	"datadog_downtime_schedule":                             datadogExternalNameWithInjectedUUID(),
	"datadog_gcp_uc_config":                                 datadogExternalNameWithInjectedUUID(),
	"datadog_governance_control":                            datadogExternalNameWithInjectedUUID(),
	"datadog_incident_notification_rule":                    datadogExternalNameWithInjectedUUID(),
	"datadog_incident_notification_template":                datadogExternalNameWithInjectedUUID(),
	"datadog_incident_type":                                 datadogExternalNameWithInjectedUUID(),
	"datadog_incident_user_defined_field":                   datadogExternalNameWithInjectedUUID(),
	"datadog_incident_user_defined_role":                    datadogExternalNameWithInjectedUUID(),
	"datadog_integration_aws_account":                       datadogExternalNameWithInjectedUUID(),
	"datadog_integration_aws_account_ccm_config":            datadogExternalNameWithInjectedUUID(),
	"datadog_integration_aws_event_bridge":                  config.IdentifierFromProvider,
	"datadog_integration_aws_external_id":                   datadogExternalNameWithInjectedUUID(),
	"datadog_integration_azure":                             config.IdentifierFromProvider,
	"datadog_integration_cloudflare_account":                datadogExternalNameWithInjectedUUID(),
	"datadog_integration_confluent_account":                 datadogExternalNameWithInjectedUUID(),
	"datadog_integration_confluent_resource":                datadogExternalNameWithInjectedUUIDPair(),
	"datadog_integration_fastly_account":                    datadogExternalNameWithInjectedUUID(),
	"datadog_integration_fastly_service":                    datadogExternalNameWithInjectedUUIDPair(),
	"datadog_integration_gcp":                               config.IdentifierFromProvider,
	"datadog_integration_gcp_sts":                           config.IdentifierFromProvider,
	"datadog_integration_ms_teams_tenant_based_handle":      datadogExternalNameWithInjectedUUID(),
	"datadog_integration_ms_teams_workflows_webhook_handle": datadogExternalNameWithInjectedUUID(),
	"datadog_ip_allowlist":                                  config.IdentifierFromProvider,
	"datadog_logs_custom_destination":                       datadogExternalNameWithInjectedUUID(),
	"datadog_logs_restriction_query":                        datadogExternalNameWithInjectedUUID(),
	"datadog_monitor_notification_rule":                     datadogExternalNameWithInjectedUUID(),
	"datadog_observability_pipeline":                        datadogExternalNameWithInjectedUUID(),
	"datadog_on_call_escalation_policy":                     datadogExternalNameWithInjectedUUID(),
	"datadog_on_call_schedule":                              datadogExternalNameWithInjectedUUID(),
	// id is the id of the team the routing rules belong to; see the reference
	// on this field in the resource configurator.
	"datadog_on_call_team_routing_rules":             config.TemplatedStringAsIdentifier("id", "{{ .parameters.id }}"),
	"datadog_on_call_user_notification_channel":      datadogExternalNameWithInjectedUUID(),
	"datadog_on_call_user_notification_rule":         datadogExternalNameWithInjectedUUID(),
	"datadog_org_connection":                         datadogExternalNameWithInjectedUUID(),
	"datadog_org_group":                              datadogExternalNameWithInjectedUUID(),
	"datadog_org_group_membership":                   datadogExternalNameWithInjectedUUID(),
	"datadog_org_group_policy":                       datadogExternalNameWithInjectedUUID(),
	"datadog_org_group_policy_override":              datadogExternalNameWithInjectedUUID(),
	"datadog_reference_table":                        datadogExternalNameWithInjectedUUID(),
	"datadog_restriction_policy":                     config.IdentifierFromProvider,
	"datadog_rum_application":                        datadogExternalNameWithInjectedID(),
	"datadog_rum_exclusion_filter":                   datadogExternalNameWithInjectedUUID(),
	"datadog_rum_metric":                             datadogExternalNameWithInjectedUUID(),
	"datadog_rum_retention_filter":                   datadogExternalNameWithInjectedUUID(),
	"datadog_rum_retention_quota":                    datadogExternalNameWithInjectedUUID(),
	"datadog_saml_idp_metadata":                      datadogExternalNameWithInjectedUUID(),
	"datadog_secure_embed_dashboard":                 datadogExternalNameWithInjectedUUID(),
	"datadog_security_findings_due_date_rule":        datadogExternalNameWithInjectedUUID(),
	"datadog_security_findings_mute_rule":            datadogExternalNameWithInjectedUUID(),
	"datadog_security_findings_ticket_creation_rule": datadogExternalNameWithInjectedUUID(),
	"datadog_security_monitoring_critical_asset":     datadogExternalNameWithInjectedUUID(),
	"datadog_security_monitoring_default_rule":       config.IdentifierFromProvider,
	"datadog_security_monitoring_filter":             datadogExternalNameWithInjectedUUID(),
	"datadog_security_monitoring_rule":               datadogExternalNameWithInjectedUUID(),
	"datadog_security_monitoring_rule_json":          datadogExternalNameWithInjectedUUID(),
	"datadog_security_monitoring_suppression":        datadogExternalNameWithInjectedUUID(),
	// Read reports the 404 for the UUID stub as an error instead of removing
	// the resource, so the not-found diagnostic lets Create run.
	"datadog_security_notification_rule":         datadogExternalNameWithInjectedUUIDAndNotFoundDiagnostics("404 Not Found"),
	"datadog_sensitive_data_scanner_group_order": config.IdentifierFromProvider,
	"datadog_service_access_token":               datadogExternalNameWithInjectedUUID(),
	"datadog_service_account":                    datadogExternalNameWithInjectedUUID(),
	"datadog_service_account_application_key":    config.IdentifierFromProvider,
	// The Terraform id is the entity reference "<kind>:<name>"; the API
	// answers 400 to any other shape, so the stub keeps it.
	"datadog_software_catalog":                 datadogExternalNameWithInjectedRef("service"),
	"datadog_spans_metric":                     datadogExternalNameWithInjectedID(),
	"datadog_status_page":                      datadogExternalNameWithInjectedUUID(),
	"datadog_status_page_component":            datadogExternalNameWithInjectedUUID(),
	"datadog_status_page_degradation_template": datadogExternalNameWithInjectedUUID(),
	"datadog_status_page_maintenance_template": datadogExternalNameWithInjectedUUID(),
	"datadog_synthetics_concurrency_cap":       config.IdentifierFromProvider,
	"datadog_synthetics_global_variable":       datadogExternalNameWithInjectedUUID(),
	"datadog_synthetics_private_location":      datadogExternalNameWithInjectedUUID(),
	"datadog_synthetics_suite":                 datadogExternalNameWithInjectedUUID(),
	"datadog_tag_pipeline_ruleset":             datadogExternalNameWithInjectedUUID(),
	"datadog_tag_pipeline_rulesets":            datadogExternalNameWithInjectedUUID(),
	"datadog_tag_rule":                         datadogExternalNameWithInjectedUUIDAndNotFoundDiagnostics("invalid policy_id"),
	"datadog_team":                             datadogExternalNameWithInjectedID(),
	"datadog_team_connection":                  datadogExternalNameWithInjectedUUID(),
	"datadog_team_hierarchy_links":             datadogExternalNameWithInjectedUUID(),
	"datadog_team_link":                        datadogExternalNameWithInjectedID(),
	"datadog_team_membership":                  config.IdentifierFromProvider,
	"datadog_team_notification_rule":           datadogExternalNameWithInjectedUUID(),
	// The framework resource resolves an empty id from team_id and action and
	// stores the API identifier of the setting, which exists as soon as its
	// team does; a stub id only produces a not-found error.
	"datadog_team_permission_setting":           config.IdentifierFromProvider,
	"datadog_team_sync":                         datadogExternalNameWithInjectedUUID(),
	"datadog_user_role":                         datadogExternalNameWithInjectedUUID(),
	"datadog_webhook":                           config.IdentifierFromProvider,
	"datadog_webhook_custom_variable":           config.IdentifierFromProvider,
	"datadog_webhook_oauth2_client_credentials": datadogExternalNameWithInjectedUUID(),
	"datadog_workflow_automation":               datadogExternalNameWithInjectedUUID(),
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

// datadogExternalNameWithInjectedRef is for resources whose Terraform id is
// an entity reference "<kind>:<name>", such as datadog_software_catalog. The
// API rejects a bare UUID as a reference with 400, while an unknown reference
// of the right shape answers an empty list, which the provider's Read turns
// into "does not exist" so that Create can run.
func datadogExternalNameWithInjectedRef(kind string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
		if len(externalName) == 0 {
			return kind + ":" + uuid.New().String(), nil
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

// datadogExternalNameWithInjectedUUIDAndNotFoundDiagnostics is
// datadogExternalNameWithInjectedUUID for a plugin framework resource whose
// Read reports a missing object as an error diagnostic instead of removing it
// from state: an error diagnostic whose summary or detail contains any of the
// substrings is treated as "not found", so that Create runs.
func datadogExternalNameWithInjectedUUIDAndNotFoundDiagnostics(substrings ...string) config.ExternalName {
	e := datadogExternalNameWithInjectedUUID()
	e.IsNotFoundDiagnosticFn = func(diags []*tfprotov6.Diagnostic) bool {
		for _, d := range diags {
			if d == nil || d.Severity != tfprotov6.DiagnosticSeverityError {
				continue
			}
			for _, s := range substrings {
				if strings.Contains(d.Summary, s) || strings.Contains(d.Detail, s) {
					return true
				}
			}
		}
		return false
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
