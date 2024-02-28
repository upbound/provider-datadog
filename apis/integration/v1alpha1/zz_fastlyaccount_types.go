// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FastlyAccountInitParameters struct {

	// (String) The API key for the Fastly account.
	// The API key for the Fastly account.
	APIKey *string `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	// (String) The name of the Fastly account.
	// The name of the Fastly account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FastlyAccountObservation struct {

	// (String) The API key for the Fastly account.
	// The API key for the Fastly account.
	APIKey *string `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the Fastly account.
	// The name of the Fastly account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FastlyAccountParameters struct {

	// (String) The API key for the Fastly account.
	// The API key for the Fastly account.
	// +kubebuilder:validation:Optional
	APIKey *string `json:"apiKey,omitempty" tf:"api_key,omitempty"`

	// (String) The name of the Fastly account.
	// The name of the Fastly account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// FastlyAccountSpec defines the desired state of FastlyAccount
type FastlyAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FastlyAccountParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FastlyAccountInitParameters `json:"initProvider,omitempty"`
}

// FastlyAccountStatus defines the observed state of FastlyAccount.
type FastlyAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FastlyAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FastlyAccount is the Schema for the FastlyAccounts API. Provides a Datadog IntegrationFastlyAccount resource. This can be used to create and manage Datadog integrationfastlyaccount.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type FastlyAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.apiKey) || (has(self.initProvider) && has(self.initProvider.apiKey))",message="spec.forProvider.apiKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FastlyAccountSpec   `json:"spec"`
	Status FastlyAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FastlyAccountList contains a list of FastlyAccounts
type FastlyAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FastlyAccount `json:"items"`
}

// Repository type metadata.
var (
	FastlyAccount_Kind             = "FastlyAccount"
	FastlyAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FastlyAccount_Kind}.String()
	FastlyAccount_KindAPIVersion   = FastlyAccount_Kind + "." + CRDGroupVersion.String()
	FastlyAccount_GroupVersionKind = CRDGroupVersion.WithKind(FastlyAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&FastlyAccount{}, &FastlyAccountList{})
}