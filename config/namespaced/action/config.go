package action

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_action_connection", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Connection"
		r.ShortGroup = "action.datadog"
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "anthropic", "anthropic.api_key", "asana", "asana.access_token", "aws", "aws.assume_role")
		common.EmbedSingleNestedBlocks(r, "azure", "azure.tenant", "circle_ci", "circle_ci.api_key", "clickup", "clickup.api_key")
		common.EmbedSingleNestedBlocks(r, "cloudflare", "cloudflare.api_token", "cloudflare.global_api_token", "config_cat", "config_cat.sdk_key", "datadog")
		common.EmbedSingleNestedBlocks(r, "datadog.api_key", "fastly", "fastly.api_key", "freshservice", "freshservice.api_key", "gcp")
		common.EmbedSingleNestedBlocks(r, "gcp.service_account", "gemini", "gemini.api_key", "gitlab", "gitlab.api_key", "grey_noise")
		common.EmbedSingleNestedBlocks(r, "grey_noise.api_key", "http", "http.token_auth", "http.token_auth.body", "launch_darkly", "launch_darkly.api_key")
		common.EmbedSingleNestedBlocks(r, "notion", "notion.api_key", "okta", "okta.api_token", "openai", "openai.api_key")
		common.EmbedSingleNestedBlocks(r, "service_now", "service_now.basic_auth", "split", "split.api_key", "statsig", "statsig.api_key")
		common.EmbedSingleNestedBlocks(r, "virus_total", "virus_total.api_key")
	})
}
