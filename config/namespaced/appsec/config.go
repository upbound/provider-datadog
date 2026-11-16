package appsec

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_appsec_waf_custom_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "WAFCustomRule"
		r.ShortGroup = "appsec.datadog"
	})
	p.AddResourceConfigurator("datadog_appsec_waf_exclusion_filter", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "WAFExclusionFilter"
		r.ShortGroup = "appsec.datadog"
	})
}
