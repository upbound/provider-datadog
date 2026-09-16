package team

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_team", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_team_connection", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamConnection"
		r.ShortGroup = "datadog"
		// team and connected_team both have required descendants, which
		// upjet's tfjson conversion otherwise pushes into status.atProvider
		// only; see common.MarkSingleNestedBlockConfigurable.
		common.EmbedSingleNestedBlocks(r, "team", "connected_team")
		common.MarkSingleNestedBlockConfigurable(r, "team", false)
		common.MarkSingleNestedBlockConfigurable(r, "connected_team", false)
	})
	p.AddResourceConfigurator("datadog_team_hierarchy_links", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamHierarchyLinks"
		r.ShortGroup = "datadog"
		r.References["parent_team_id"] = config.Reference{
			TerraformName: "datadog_team",
		}
		r.References["sub_team_id"] = config.Reference{
			TerraformName: "datadog_team",
		}
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
	p.AddResourceConfigurator("datadog_team_notification_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamNotificationRule"
		r.ShortGroup = "datadog"
		r.References["team_id"] = config.Reference{
			TerraformName: "datadog_team",
		}
		// email, ms_teams, pagerduty, and slack are plugin-framework
		// single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "email", "ms_teams", "pagerduty", "slack")
	})
	p.AddResourceConfigurator("datadog_team_permission_setting", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamPermissionSetting"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_team_sync", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "TeamSync"
		r.ShortGroup = "datadog"
		// selection_state.external_id has a required descendant, which
		// upjet's tfjson conversion otherwise pushes into status.atProvider
		// only; see common.MarkSingleNestedBlockConfigurable.
		common.EmbedSingleNestedBlocks(r, "selection_state.external_id")
		common.MarkSingleNestedBlockConfigurable(r, "selection_state.external_id", false)
	})
}
