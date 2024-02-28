package team

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_team", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_team_link", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamLink"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_team_membership", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamMembership"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_team_permission_setting", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamPermissionSetting"
		r.ShortGroup = "datadog"
	})
}
