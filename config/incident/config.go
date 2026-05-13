package incident

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_incident_type", func(r *config.Resource) {
		// We need to override the default upjet-generated Kind which would be "Type"
		// since "type" is a Go reserved keyword and cannot be used as a package name.
		r.Kind = "IncidentType"
	})
}
