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
		// The API returns these blocks zero-valued for tests that do not use
		// them; late-initializing them into the spec makes Terraform demand
		// their required arguments on the next refresh.
		r.LateInitializer = config.LateInitializer{IgnoredFields: []string{
			"mobile_options_list",
			"options_list.rum_settings",
			"request_basicauth",
			"request_client_certificate",
			"request_file",
			"request_proxy",
		}}
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
