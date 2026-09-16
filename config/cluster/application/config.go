package application

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_application_key", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AppKey"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_app_key_registration", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AppKeyRegistration"
		r.ShortGroup = "datadog"
		// id is the identifier of an existing Application Key (see the
		// TemplatedStringAsIdentifier external-name config in
		// config/external_name.go), but upjet drops a spec field literally
		// named "id" from the CRD entirely, so there is no idRef/idSelector
		// to wire a reference onto; the external name can only be set via
		// the crossplane.io/external-name annotation.
	})
}
