package authentication

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_authn_mapping", func(r *config.Resource) {
		r.References["role"] = config.Reference{
			TerraformName: "datadog_role",
		}
	})
}
