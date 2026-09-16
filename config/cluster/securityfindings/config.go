package securityfindings

import "github.com/crossplane/upjet/v2/pkg/config"

const securityFindingsDatadog = "securityfindings.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_security_findings_due_date_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DueDateRule"
		r.ShortGroup = securityFindingsDatadog
	})
	p.AddResourceConfigurator("datadog_security_findings_mute_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "MuteRule"
		r.ShortGroup = securityFindingsDatadog
	})
	p.AddResourceConfigurator("datadog_security_findings_ticket_creation_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TicketCreationRule"
		r.ShortGroup = securityFindingsDatadog
	})
}
