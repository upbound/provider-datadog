/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/types/comments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// codegenConfigurators are the resource configurators that change the
// Terraform schema the CRDs are generated from. They run only for the code
// generation provider: at runtime a plugin SDK resource is described by the
// live schema of the in-process provider, whose CRUD functions rely on the
// original field shapes, and a plugin framework resource does not read this
// schema at all.
var codegenConfigurators = map[string]ujconfig.ResourceConfiguratorFn{
	// The widget schema is hundreds of kilobytes of nested blocks; the CRD
	// keeps it as a JSON string that is never sent to Terraform.
	"datadog_powerpack": widgetAsJSONString("(String) The JSON formatted definition of the list of widgets to display in the powerpack."),
	// rum_settings.client_token_id is a sensitive number upstream, and upjet
	// can only generate string-typed sensitive fields. The value is the
	// numeric id of a RUM client token, not the token itself, so it is
	// exposed as a plain field instead of hiding the whole block.
	"datadog_synthetics_test": func(r *ujconfig.Resource) {
		r.TerraformResource.
			Schema["options_list"].Elem.(*schema.Resource).
			Schema["rum_settings"].Elem.(*schema.Resource).
			Schema["client_token_id"].Sensitive = false
	},
}

func widgetAsJSONString(description string) ujconfig.ResourceConfiguratorFn {
	return func(r *ujconfig.Resource) {
		desc, _ := comments.New(description, comments.WithTFTag("-"))
		r.TerraformResource.Schema["widget"] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: desc.String(),
		}
	}
}
