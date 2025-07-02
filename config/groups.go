package config

import (
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/name"
)

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "datadog_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

var GroupMap = map[string]GroupKindCalculator{
	"datadog_api_key":                            ReplaceGroupWords("access", 0),
	"datadog_application_key":                    ReplaceGroupWords("access", 0),
	"datadog_authn_mapping":                      ReplaceGroupWords("authentication", 0),
	"datadog_dashboard":                          ReplaceGroupWords("dashboard", 0),
	"datadog_dashboard_json":                     ReplaceGroupWords("dashboard", 0),
	"datadog_dashboard_list":                     ReplaceGroupWords("dashboard", 0),
	"datadog_powerpack":                          ReplaceGroupWords("dashboard", 0),
	"datadog_restriction_policy":                 ReplaceGroupWords("iam", 0),
	"datadog_role":                               ReplaceGroupWords("iam", 0),
	"datadog_user":                               ReplaceGroupWords("iam", 0),
	"datadog_service_account":                    ReplaceGroupWords("iam", 0),
	"datadog_service_account_application_key":    ReplaceGroupWords("iam", 0),
	"datadog_team":                               ReplaceGroupWords("iam", 0),
	"datadog_team_link":                          ReplaceGroupWords("iam", 0),
	"datadog_team_membership":                    ReplaceGroupWords("iam", 0),
	"datadog_team_permission_setting":            ReplaceGroupWords("iam", 0),
	"datadog_spans_metric":                       ReplaceGroupWords("metric", 0),
	"datadog_monitor":                            ReplaceGroupWords("monitor", 0),
	"datadog_monitor_json":                       ReplaceGroupWords("monitor", 0),
	"datadog_downtime":                           ReplaceGroupWords("monitor", 0),
	"datadog_downtime_schedule":                  ReplaceGroupWords("monitor", 0),
	"datadog_webhook":                            ReplaceGroupWords("notification", 0),
	"datadog_webhook_custom_variable":            ReplaceGroupWords("notification", 0),
	"datadog_child_organization":                 ReplaceGroupWords("organization", 0),
	"datadog_security_monitoring_default_rule":   ReplaceGroupWords("securitymonitoring", 2),
	"datadog_security_monitoring_filter":         ReplaceGroupWords("securitymonitoring", 2),
	"datadog_security_monitoring_rule":           ReplaceGroupWords("securitymonitoring", 2),
	"datadog_sensitive_data_scanner_group":       ReplaceGroupWords("sensitivedata", 2),
	"datadog_sensitive_data_scanner_group_order": ReplaceGroupWords("sensitivedata", 2),
	"datadog_sensitive_data_scanner_rule":        ReplaceGroupWords("sensitivedata", 2),
	"datadog_service_level_objective":            ReplaceGroupWords("slo", 0),
}
