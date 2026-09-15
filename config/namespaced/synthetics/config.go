package synthetics

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_synthetics_concurrency_cap", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
	p.AddResourceConfigurator("datadog_synthetics_global_variable", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
	p.AddResourceConfigurator("datadog_synthetics_private_location", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
	p.AddResourceConfigurator("datadog_synthetics_test", func(r *config.Resource) {
		// rum_settings.client_token_id is a sensitive number upstream, and
		// upjet can only generate string-typed sensitive fields. The value is
		// the numeric id of a RUM client token, not the token itself, so it
		// is exposed as a plain field instead of hiding the whole block.
		r.TerraformResource.
			Schema["options_list"].Elem.(*schema.Resource).
			Schema["rum_settings"].Elem.(*schema.Resource).
			Schema["client_token_id"].Sensitive = false
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "synthetics.datadog"
	})
}
