/*
Copyright 2026 Upbound Inc.
*/

package downtime

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_downtime", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.ShortGroup = "datadog"
	})
	p.AddResourceConfigurator("datadog_downtime_schedule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DowntimeSchedule"
		r.ShortGroup = "datadog"
		common.EmbedSingleNestedBlocks(r, "monitor_identifier", "one_time_schedule", "recurring_schedule")
	})
}
