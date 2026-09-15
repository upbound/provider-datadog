/*
Copyright 2026 Upbound Inc.
*/

package spans

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_spans_metric", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "SpansMetric"
		r.ShortGroup = "datadog"
		common.EmbedSingleNestedBlocks(r, "compute", "filter")
		common.MarkSingleNestedBlockConfigurable(r, "compute", false)
	})
}
