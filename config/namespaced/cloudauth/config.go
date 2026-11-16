/*
Copyright 2026 Upbound Inc.
*/

package cloudauth

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const cloudAuthDatadog = "cloudauth.datadog"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_aws_wif_identity_mapping", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSWIFIdentityMapping"
		r.ShortGroup = cloudAuthDatadog
		r.References["account_identifier"] = config.Reference{
			TerraformName: "datadog_service_account",
		}
	})
	p.AddResourceConfigurator("datadog_aws_wif_intake_mapping", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "AWSWIFIntakeMapping"
		r.ShortGroup = cloudAuthDatadog
	})
	p.AddResourceConfigurator("datadog_github_oidc_identity_mapping", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "GitHubOIDCIdentityMapping"
		r.ShortGroup = cloudAuthDatadog
		r.References["account_identifier"] = config.Reference{
			TerraformName: "datadog_service_account",
		}
	})
	p.AddResourceConfigurator("datadog_github_oidc_intake_mapping", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "datadog"
		r.Kind = "GitHubOIDCIntakeMapping"
		r.ShortGroup = cloudAuthDatadog
	})
}
