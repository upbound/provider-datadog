package child

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_child_organization", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "ChildOrganization"
		r.ShortGroup = "datadog"
	})
}
