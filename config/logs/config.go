package logs

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_logs_archive_order", func(r *config.Resource) {
		r.References["archive_ids"] = config.Reference{
			TerraformName: "datadog_logs_archive",
		}
	})
	p.AddResourceConfigurator("datadog_logs_index_order", func(r *config.Resource) {
		r.References["indexes"] = config.Reference{
			TerraformName: "datadog_logs_index",
		}
	})
}
