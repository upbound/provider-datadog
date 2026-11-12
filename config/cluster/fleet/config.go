/*
Copyright 2026 Upbound Inc.
*/

package fleet

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_fleet_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Schedule"
		r.ShortGroup = "fleet.datadog"
	})
}
