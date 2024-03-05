package webhook

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_webhook", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Webhook"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_webhook_custom_variable", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "WebhookCustomVariable"
		r.ShortGroup = "datadog"
	})
}
