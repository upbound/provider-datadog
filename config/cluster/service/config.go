package service

import "github.com/crossplane/upjet/v2/pkg/config"

const datadog = "datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_service_access_token", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ServiceAccessToken"
		r.ShortGroup = datadog
		r.References["service_account_id"] = config.Reference{
			TerraformName: "datadog_service_account",
		}
	})
	p.AddResourceConfigurator("datadog_service_account", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ServiceAccount"
		r.ShortGroup = datadog
	})
	p.AddResourceConfigurator("datadog_service_account_application_key", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ServiceAccountApplicationKey"
		r.ShortGroup = datadog
	})
	p.AddResourceConfigurator("datadog_service_definition_yaml", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ServiceDefinitionYAML"
		r.ShortGroup = datadog
	})
	p.AddResourceConfigurator("datadog_service_level_objective", func(r *config.Resource) {
		// The API returns the SLI specification it derives from a metric
		// query (and vice versa); late-initializing either next to the one
		// the user set makes Terraform reject the pair as conflicting.
		r.LateInitializer = config.LateInitializer{IgnoredFields: []string{"query", "sli_specification"}}
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "ServiceLevelObjective"
		r.ShortGroup = datadog
	})
	p.AddResourceConfigurator("datadog_slo_correction", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "SLOCorrection"
		r.ShortGroup = datadog
	})
}
