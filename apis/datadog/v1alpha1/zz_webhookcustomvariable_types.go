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

type WebhookCustomVariableInitParameters struct {

	// (Boolean) Whether the custom variable is secret or not.
	// Whether the custom variable is secret or not.
	IsSecret *bool `json:"isSecret,omitempty" tf:"is_secret,omitempty"`

	// (String) The name of the variable. It corresponds with <CUSTOM_VARIABLE_NAME>.
	// The name of the variable. It corresponds with `<CUSTOM_VARIABLE_NAME>`.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type WebhookCustomVariableObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Whether the custom variable is secret or not.
	// Whether the custom variable is secret or not.
	IsSecret *bool `json:"isSecret,omitempty" tf:"is_secret,omitempty"`

	// (String) The name of the variable. It corresponds with <CUSTOM_VARIABLE_NAME>.
	// The name of the variable. It corresponds with `<CUSTOM_VARIABLE_NAME>`.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type WebhookCustomVariableParameters struct {

	// (Boolean) Whether the custom variable is secret or not.
	// Whether the custom variable is secret or not.
	// +kubebuilder:validation:Optional
	IsSecret *bool `json:"isSecret,omitempty" tf:"is_secret,omitempty"`

	// (String) The name of the variable. It corresponds with <CUSTOM_VARIABLE_NAME>.
	// The name of the variable. It corresponds with `<CUSTOM_VARIABLE_NAME>`.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String, Sensitive) The value of the custom variable.
	// The value of the custom variable.
	// +kubebuilder:validation:Optional
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

// WebhookCustomVariableSpec defines the desired state of WebhookCustomVariable
type WebhookCustomVariableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebhookCustomVariableParameters `json:"forProvider"`
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
	InitProvider WebhookCustomVariableInitParameters `json:"initProvider,omitempty"`
}

// WebhookCustomVariableStatus defines the observed state of WebhookCustomVariable.
type WebhookCustomVariableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebhookCustomVariableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WebhookCustomVariable is the Schema for the WebhookCustomVariables API. Provides a Datadog webhooks custom variable resource. This can be used to create and manage Datadog webhooks custom variables.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type WebhookCustomVariable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.isSecret) || (has(self.initProvider) && has(self.initProvider.isSecret))",message="spec.forProvider.isSecret is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.valueSecretRef)",message="spec.forProvider.valueSecretRef is a required parameter"
	Spec   WebhookCustomVariableSpec   `json:"spec"`
	Status WebhookCustomVariableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookCustomVariableList contains a list of WebhookCustomVariables
type WebhookCustomVariableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebhookCustomVariable `json:"items"`
}

// Repository type metadata.
var (
	WebhookCustomVariable_Kind             = "WebhookCustomVariable"
	WebhookCustomVariable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebhookCustomVariable_Kind}.String()
	WebhookCustomVariable_KindAPIVersion   = WebhookCustomVariable_Kind + "." + CRDGroupVersion.String()
	WebhookCustomVariable_GroupVersionKind = CRDGroupVersion.WithKind(WebhookCustomVariable_Kind)
)

func init() {
	SchemeBuilder.Register(&WebhookCustomVariable{}, &WebhookCustomVariableList{})
}