/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"github.com/terraform-providers/terraform-provider-datadog/datadog/fwprovider"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/upbound/provider-datadog/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-datadog/apis/namespaced/v1beta1"
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
	keyHTTPClientRetryMaxRetries        = "http_client_retry_max_retries"
	keyHTTPClientRetryTimeout           = "http_client_retry_timeout"
	keyValidate                         = "validate"

	// error messages
	errNoProviderConfig      = "no providerConfigRef provided"
	errGetProviderConfig     = "cannot get referenced ProviderConfig"
	errResolveProviderConfig = "cannot resolve provider config"
	errTrackUsage            = "cannot track ProviderConfig usage"
	errExtractCredentials    = "cannot extract credentials"
	errUnmarshalCredentials  = "cannot unmarshal datadog credentials as JSON"
	errConfigureSDKProvider  = "cannot configure the plugin SDK provider"
)

// providerStringKeys lists the string-typed Terraform provider arguments that
// can be supplied through the credentials secret.
var providerStringKeys = []string{
	keyAPIKey,
	keyAPIURL,
	keyAppKey,
	keyHTTPClientRetryEnabled,
	keyValidate,
}

// providerIntKeys lists the integer-typed Terraform provider arguments that
// can be supplied through the credentials secret. They are integers on both
// provider schemas, and the plugin framework rejects a string for them.
var providerIntKeys = []string{
	keyHTTPClientRetryBackoffBase,
	keyHTTPClientRetryBackoffMultiplier,
	keyHTTPClientRetryMaxRetries,
	keyHTTPClientRetryTimeout,
}

// TerraformSetupBuilder builds the terraform.SetupFn that configures the
// in-process Datadog Terraform providers for a managed resource. sdkProvider
// is the plugin SDK provider the resources are reconciled with; the plugin
// framework provider is instantiated per call, the way the upstream provider
// server does it.
func TerraformSetupBuilder(sdkProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return ps, errors.Wrap(err, errResolveProviderConfig)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration, err = providerConfiguration(creds)
		if err != nil {
			return ps, err
		}
		return ps, configureProviders(ctx, &ps, *sdkProvider)
	}
}

// providerConfiguration translates the credentials secret into the Terraform
// provider configuration shared by both in-process providers.
func providerConfiguration(creds map[string]string) (map[string]any, error) {
	cfg := map[string]any{}
	for _, key := range providerStringKeys {
		if v, ok := creds[key]; ok {
			cfg[key] = v
		}
	}
	for _, key := range providerIntKeys {
		v, ok := creds[key]
		if !ok {
			continue
		}
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.Wrapf(err, "credential %q must be an integer", key)
		}
		cfg[key] = n
	}
	return cfg, nil
}

// configureProviders configures the plugin SDK provider with the setup's
// configuration and stores the resulting meta, which the SDK resources
// receive on every CRUD call, next to a fresh plugin framework provider that
// upjet configures itself. The SDK provider is passed by value on purpose:
// Configure stores the meta on the provider, and a shared pointer would race
// between resources that reference different ProviderConfigs.
func configureProviders(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	if diags := p.Configure(ctx, tfsdk.NewResourceConfigRaw(ps.Configuration)); diags.HasError() {
		return errors.Errorf("%s: %v", errConfigureSDKProvider, diags)
	}
	ps.Meta = p.Meta()
	ps.FrameworkProvider = fwprovider.New()
	return nil
}

// toSharedPCSpec converts a legacy cluster-scoped ProviderConfig spec into the
// namespaced spec type so the rest of the setup operates on a single type.
func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged: //nolint:staticcheck // still handling cluster-scoped behavior
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed resource")
	}
}

func resolveLegacy(ctx context.Context, crClient client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) { //nolint:staticcheck // still handling cluster-scoped behavior
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(crClient, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(resource.ProviderConfig)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.Errorf("referenced kind %s is not a ProviderConfig", configRef.Kind)
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		// Secret references of a namespaced ProviderConfig are local to the
		// namespace of the managed resource.
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
	default:
		return nil, errors.New("unknown provider config type")
	}
	t := resource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}
