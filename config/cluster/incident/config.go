package incident

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const incidentDatadog = "incident.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_incident_notification_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "NotificationRule"
		r.ShortGroup = incidentDatadog
	})
	p.AddResourceConfigurator("datadog_incident_notification_template", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "NotificationTemplate"
		r.ShortGroup = incidentDatadog
	})
	p.AddResourceConfigurator("datadog_incident_type", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		//
		// Kind must not be "Type": upjet lowercases it for the generated
		// package directory, and "type" is a Go reserved word.
		r.Kind = "IncidentType"
		r.ShortGroup = incidentDatadog
	})
	p.AddResourceConfigurator("datadog_incident_user_defined_field", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "UserDefinedField"
		r.ShortGroup = incidentDatadog
	})
	p.AddResourceConfigurator("datadog_incident_user_defined_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "UserDefinedRole"
		r.ShortGroup = incidentDatadog
	})
}
