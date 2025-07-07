package slo

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_slo_correction", func(r *config.Resource) {
		r.References["slo_id"] = config.Reference{
			TerraformName: "datadog_service_level_objective",
		}
	})
}
