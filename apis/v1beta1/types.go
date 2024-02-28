/*
Copyright 2022 Upbound Inc.
*/

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// A ProviderConfigSpec defines the desired state of a ProviderConfig.
type ProviderConfigSpec struct {
	// (String, Sensitive) (Required unless validate is false) Datadog API key.
	// This can also be set via the DD_API_KEY environment variable.
	// +optional
	APIKey string `json:"api_key,omitempty"`

	// (String) The API URL. This can also be set via the DD_HOST environment
	// variable. Note that this URL must not end with the /api/ path. For
	// example, https://api.datadoghq.com/ is a correct value, while
	// https://api.datadoghq.com/api/ is not. And if you're working with
	// "EU" version of Datadog, use https://api.datadoghq.eu/. Other Datadog
	// region examples: https://api.us5.datadoghq.com/,
	// https://api.us3.datadoghq.com/ and https://api.ddog-gov.com/.
	// See https://docs.datadoghq.com/getting_started/site/ for all available
	// regions.
	// +optional
	APIUrl string `json:"api_url,omitempty"`

	// (String, Sensitive) (Required unless validate is false) Datadog APP key.
	// This can also be set via the DD_APP_KEY environment variable.
	// +optional
	AppKey string `json:"app_key,omitempty"`

	// (Number) The HTTP request retry back off base. Defaults to 2.
	// +optional
	HTTPClientRetryBackoffBase int `json:"http_client_retry_backoff_base,omitempty"`

	// (Number) The HTTP request retry back off multiplier. Defaults to 2.
	// +optional
	HTTPClientRetryBackoffMultiplier int `json:"http_client_retry_backoff_multiplier,omitempty"`

	// (String) Enables request retries on HTTP status codes 429 and 5xx.
	// Valid values are [true, false]. Defaults to true.
	// +optional
	HTTPClientRetryEnabled string `json:"http_client_retry_enabled,omitempty"`

	// (Number) The HTTP request maximum retry number. Defaults to 3.
	// +optional
	HTTPClientRetryMaxRetries int `json:"http_client_retry_max_retries,omitempty"`

	// (Number) The HTTP request retry timeout period. Defaults to 60 seconds.
	// +optional
	HTTPClientRetryTimeout int `json:"http_client_retry_timeout,omitempty"`

	// (String) Enables validation of the provided API key during provider
	// initialization. Valid values are [true, false]. Default is true.
	// When false, api_key won't be checked.
	// +optional
	Validate string `json:"validate,omitempty"`

	// Credentials required to authenticate to this provider.
	// +optional
	Credentials ProviderCredentials `json:"credentials"`
}

// ProviderCredentials required to authenticate.
type ProviderCredentials struct {
	// Source of the provider credentials.
	// +kubebuilder:validation:Enum=None;Secret;InjectedIdentity;Environment;Filesystem
	Source xpv1.CredentialsSource `json:"source"`

	xpv1.CommonCredentialSelectors `json:",inline"`
}

// A ProviderConfigStatus reflects the observed state of a ProviderConfig.
type ProviderConfigStatus struct {
	xpv1.ProviderConfigStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A ProviderConfig configures a Datadog provider.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="SECRET-NAME",type="string",JSONPath=".spec.credentials.secretRef.name",priority=1
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,datadog}
type ProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProviderConfigSpec   `json:"spec"`
	Status ProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderConfigList contains a list of ProviderConfig.
type ProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfig `json:"items"`
}

// +kubebuilder:object:root=true

// A ProviderConfigUsage indicates that a resource is using a ProviderConfig.
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="CONFIG-NAME",type="string",JSONPath=".providerConfigRef.name"
// +kubebuilder:printcolumn:name="RESOURCE-KIND",type="string",JSONPath=".resourceRef.kind"
// +kubebuilder:printcolumn:name="RESOURCE-NAME",type="string",JSONPath=".resourceRef.name"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,provider,datadog}
type ProviderConfigUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	xpv1.ProviderConfigUsage `json:",inline"`
}

// +kubebuilder:object:root=true

// ProviderConfigUsageList contains a list of ProviderConfigUsage
type ProviderConfigUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderConfigUsage `json:"items"`
}
