package integration

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_integration_confluent_resource", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			TerraformName: "datadog_integration_confluent_account",
		}
	})
	p.AddResourceConfigurator("datadog_integration_fastly_service", func(r *config.Resource) {
		r.References["account_id"] = config.Reference{
			TerraformName: "datadog_integration_fastly_account",
		}
	})
}
