package synthetics

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_synthetics_suite", func(r *config.Resource) {
		// Override field names to avoid naming conflicts with
		// datadog_synthetics_global_variable which also has an "options" block
		// that generates the same struct names (OptionsInitParameters, etc.)
		// in the same package.
		r.OverrideFieldNames = map[string]string{
			"OptionsInitParameters": "SuiteOptionsInitParameters",
			"OptionsObservation":    "SuiteOptionsObservation",
			"OptionsParameters":     "SuiteOptionsParameters",
		}
	})
	p.AddResourceConfigurator("datadog_synthetics_test", func(r *config.Resource) {
		// client_token_id in rum_settings is type "number" with sensitive=True,
		// but upjet only supports string types for sensitive fields.
		// We convert it to string to work around this limitation.
		if rs := r.TerraformResource.Schema["options_list"]; rs != nil {
			if res, ok := rs.Elem.(*schema.Resource); ok {
				if rs := res.Schema["rum_settings"]; rs != nil {
					if res, ok := rs.Elem.(*schema.Resource); ok {
						if f := res.Schema["client_token_id"]; f != nil {
							f.Type = schema.TypeString
						}
					}
				}
			}
		}
	})
}
