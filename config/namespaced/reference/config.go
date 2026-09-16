package reference

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-datadog/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_reference_table", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Table"
		r.ShortGroup = "reference.datadog"
		// Plugin-framework single-nested blocks; see common.EmbedSingleNestedBlocks.
		common.EmbedSingleNestedBlocks(r, "file_metadata", "file_metadata.access_details", "file_metadata.access_details.aws_detail", "file_metadata.access_details.azure_detail", "file_metadata.access_details.gcp_detail", "schema")
		// These blocks have required descendants, which upjet's tfjson conversion
		// otherwise pushes into status.atProvider only; they are optional upstream.
		common.MarkSingleNestedBlockConfigurable(r, "file_metadata", true)
		common.MarkSingleNestedBlockConfigurable(r, "schema", true)
	})
}
