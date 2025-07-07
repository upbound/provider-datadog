package dashboard

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/types/comments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_dashboard", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "widget")

		desc, _ := comments.New("JSON widget to add to a dashboard",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["widget"] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: desc.String(),
		}
	})
	p.AddResourceConfigurator("datadog_powerpack", func(r *config.Resource) {
		delete(r.TerraformResource.Schema, "widget")

		desc, _ := comments.New("(String) The JSON formatted definition of the list of widgets to display in the powerpack.",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["widget"] = &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: desc.String(),
		}
	})
}
