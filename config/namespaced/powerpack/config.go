package powerpack

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_powerpack", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "Powerpack"
		r.ShortGroup = "datadog"
		// The widget list is generated as a JSON string; see codegen.go.
	})
}
