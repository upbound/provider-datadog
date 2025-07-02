/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"github.com/upbound/provider-datadog/apis/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// provider config
	// source: https://registry.terraform.io/providers/DataDog/datadog/latest/docs

	// arguments
	keyAPIKey                           = "api_key"
	keyAPIURL                           = "api_url"
	keyAppKey                           = "app_key"
	keyHTTPClientRetryBackoffBase       = "http_client_retry_backoff_base"
	keyHTTPClientRetryBackoffMultiplier = "http_client_retry_backoff_multiplier"
	keyHTTPClientRetryEnabled           = "http_client_retry_enabled"
	keyHTTPClientRetryMaxRetries        = "http_client_retry_max_retries "
	keyHTTPClientRetryTimeout           = "http_client_retry_timeout"
	keyValidate                         = "validate"

	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal datadog credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(sdkProvider *schema.Provider, fwProvider provider.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set provider configuration
		ps.Configuration = map[string]any{}
		if v, ok := creds[keyAPIKey]; ok {
			ps.Configuration[keyAPIKey] = v
		}
		if v, ok := creds[keyAPIURL]; ok {
			ps.Configuration[keyAPIURL] = v
		}
		if v, ok := creds[keyAppKey]; ok {
			ps.Configuration[keyAppKey] = v
		}
		if v, ok := creds[keyHTTPClientRetryBackoffBase]; ok {
			ps.Configuration[keyHTTPClientRetryBackoffBase] = v
		}
		if v, ok := creds[keyHTTPClientRetryBackoffMultiplier]; ok {
			ps.Configuration[keyHTTPClientRetryBackoffMultiplier] = v
		}
		if v, ok := creds[keyHTTPClientRetryEnabled]; ok {
			ps.Configuration[keyHTTPClientRetryEnabled] = v
		}
		if v, ok := creds[keyHTTPClientRetryMaxRetries]; ok {
			ps.Configuration[keyHTTPClientRetryMaxRetries] = v
		}
		if v, ok := creds[keyHTTPClientRetryTimeout]; ok {
			ps.Configuration[keyHTTPClientRetryTimeout] = v
		}
		if v, ok := creds[keyValidate]; ok {
			ps.Configuration[keyValidate] = v
		}
		return ps, errors.Wrap(configureNoForkOCIClient(ctx, &ps, sdkProvider, fwProvider), "failed to configure the Terraform datadog provider meta")
	}
}

func configureNoForkOCIClient(ctx context.Context, ps *terraform.Setup, sdkProvier *schema.Provider, fwProvider provider.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := sdkProvier.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = sdkProvier.Meta()
	ps.FrameworkProvider = fwProvider
	return nil
}
