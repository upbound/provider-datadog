package data

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_dataset", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Dataset"
		r.ShortGroup = "data.datadog"
	})
	p.AddResourceConfigurator("datadog_datastore", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "Datastore"
		r.ShortGroup = "data.datadog"
		// creator_user_id is computed (API-assigned), not a settable input,
		// so it is not part of forProvider and cannot carry a reference.
	})
	p.AddResourceConfigurator("datadog_datastore_item", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "DatastoreItem"
		r.ShortGroup = "data.datadog"
		// datastore_id also composes the Terraform id ("<datastore
		// id>:<item key>"); see config/external_name.go.
		r.References["datastore_id"] = config.Reference{
			TerraformName: "datadog_datastore",
		}
	})
}
