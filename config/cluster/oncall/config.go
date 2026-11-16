package oncall

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const onCallDatadog = "oncall.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_on_call_escalation_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "EscalationPolicy"
		r.ShortGroup = onCallDatadog
	})
	p.AddResourceConfigurator("datadog_on_call_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Schedule"
		r.ShortGroup = onCallDatadog
	})
	p.AddResourceConfigurator("datadog_on_call_team_routing_rules", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamRoutingRules"
		r.ShortGroup = onCallDatadog
		// id is the id of the team the routing rules belong to (see the
		// TemplatedStringAsIdentifier external-name config in
		// config/external_name.go), but upjet drops a spec field literally
		// named "id" from the CRD entirely, so there is no idRef/idSelector
		// to wire a reference onto; the external name can only be set via
		// the crossplane.io/external-name annotation.
		r.References["rule.action.escalation_policy.policy_id"] = config.Reference{
			TerraformName: "datadog_on_call_escalation_policy",
		}
	})
	p.AddResourceConfigurator("datadog_on_call_user_notification_channel", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "UserNotificationChannel"
		r.ShortGroup = onCallDatadog
		r.References["user_id"] = config.Reference{
			TerraformName: "datadog_user",
		}
	})
	p.AddResourceConfigurator("datadog_on_call_user_notification_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "UserNotificationRule"
		r.ShortGroup = onCallDatadog
		r.References["channel_id"] = config.Reference{
			TerraformName: "datadog_on_call_user_notification_channel",
		}
		r.References["user_id"] = config.Reference{
			TerraformName: "datadog_user",
		}
	})
}
