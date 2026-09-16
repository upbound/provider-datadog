/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

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
	errNoProviderConfig      = "no providerConfigRef provided"
	errGetProviderConfig     = "cannot get referenced ProviderConfig"
	errResolveProviderConfig = "cannot resolve provider config"
	errTrackUsage            = "cannot track ProviderConfig usage"
	errExtractCredentials    = "cannot extract credentials"
	errUnmarshalCredentials  = "cannot unmarshal datadog credentials as JSON"
	errMarshalConfiguration  = "cannot marshal the provider configuration"
	errConfigureSDKProvider  = "cannot configure the plugin SDK provider"
)

// credentials is the JSON document stored in the secret a ProviderConfig
// points at. Its fields are the arguments of the Terraform provider block
// (https://registry.terraform.io/providers/DataDog/datadog/latest/docs);
// unknown keys are ignored and unset fields are omitted from the provider
// configuration, so the providers apply their own defaults for them.
// bearer_token is an alternative to the api_key/app_key pair and takes
// precedence over it.
type credentials struct {
	APIKey                           *string  `json:"api_key,omitempty"`
	APIURL                           *string  `json:"api_url,omitempty"`
	AppKey                           *string  `json:"app_key,omitempty"`
	BearerToken                      *string  `json:"bearer_token,omitempty"`
	Validate                         *flag    `json:"validate,omitempty"`
	HTTPClientRetryEnabled           *flag    `json:"http_client_retry_enabled,omitempty"`
	HTTPClientRetryBackoffBase       *integer `json:"http_client_retry_backoff_base,omitempty"`
	HTTPClientRetryBackoffMultiplier *integer `json:"http_client_retry_backoff_multiplier,omitempty"`
	HTTPClientRetryJitter            *integer `json:"http_client_retry_jitter,omitempty"`
	HTTPClientRetryMaxRetries        *integer `json:"http_client_retry_max_retries,omitempty"`
	HTTPClientRetryTimeout           *integer `json:"http_client_retry_timeout,omitempty"`
	IgnoreTagKeys                    []string `json:"ignore_tag_keys,omitempty"`
}

// integer is a provider argument typed as a number on both provider schemas.
// It also accepts a numeric string, the form every value had in secrets
// written for earlier versions of this provider.
type integer int

func (i *integer) UnmarshalJSON(b []byte) error {
	n, err := strconv.Atoi(strings.Trim(string(b), `"`))
	if err != nil {
		return errors.Errorf("%s is not an integer", string(b))
	}
	*i = integer(n)
	return nil
}

// flag is a provider argument that the provider schemas type as the strings
// "true" and "false". It accepts a JSON boolean as well and always marshals
// as the string.
type flag bool

func (f *flag) UnmarshalJSON(b []byte) error {
	v, err := strconv.ParseBool(strings.Trim(string(b), `"`))
	if err != nil {
		return errors.Errorf("%s is not a boolean", string(b))
	}
	*f = flag(v)
	return nil
}

func (f flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatBool(bool(f)))
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
		ps.Configuration, err = providerConfiguration(data)
		if err != nil {
			return ps, err
		}
		return ps, configureProviders(ctx, &ps, *sdkProvider)
	}
}

// providerConfiguration translates the credentials secret into the Terraform
// provider configuration shared by both in-process providers. The marshal
// round trip applies the omitempty tags of credentials, so only the arguments
// present in the secret reach the providers.
func providerConfiguration(data []byte) (map[string]any, error) {
	var c credentials
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}
	raw, err := json.Marshal(c) //nolint:gosec // the document is decoded straight back into the in-memory provider configuration, never logged or stored
	if err != nil {
		return nil, errors.Wrap(err, errMarshalConfiguration)
	}
	cfg := map[string]any{}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return nil, errors.Wrap(err, errMarshalConfiguration)
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
