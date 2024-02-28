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

type ServiceAccountApplicationKeyInitParameters struct {

	// (String) Name of the application key.
	// Name of the application key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the service account that owns this key.
	// ID of the service account that owns this key.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type ServiceAccountApplicationKeyObservation struct {

	// (String) Creation date of the application key.
	// Creation date of the application key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The last four characters of the application key.
	// The last four characters of the application key.
	Last4 *string `json:"last4,omitempty" tf:"last4,omitempty"`

	// (String) Name of the application key.
	// Name of the application key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the service account that owns this key.
	// ID of the service account that owns this key.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type ServiceAccountApplicationKeyParameters struct {

	// (String) Name of the application key.
	// Name of the application key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) ID of the service account that owns this key.
	// ID of the service account that owns this key.
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

// ServiceAccountApplicationKeySpec defines the desired state of ServiceAccountApplicationKey
type ServiceAccountApplicationKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountApplicationKeyParameters `json:"forProvider"`
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
	InitProvider ServiceAccountApplicationKeyInitParameters `json:"initProvider,omitempty"`
}

// ServiceAccountApplicationKeyStatus defines the observed state of ServiceAccountApplicationKey.
type ServiceAccountApplicationKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountApplicationKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceAccountApplicationKey is the Schema for the ServiceAccountApplicationKeys API. Provides a Datadog service_account_application_key resource. This can be used to create and manage Datadog service account application keys.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type ServiceAccountApplicationKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceAccountId) || (has(self.initProvider) && has(self.initProvider.serviceAccountId))",message="spec.forProvider.serviceAccountId is a required parameter"
	Spec   ServiceAccountApplicationKeySpec   `json:"spec"`
	Status ServiceAccountApplicationKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountApplicationKeyList contains a list of ServiceAccountApplicationKeys
type ServiceAccountApplicationKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountApplicationKey `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountApplicationKey_Kind             = "ServiceAccountApplicationKey"
	ServiceAccountApplicationKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountApplicationKey_Kind}.String()
	ServiceAccountApplicationKey_KindAPIVersion   = ServiceAccountApplicationKey_Kind + "." + CRDGroupVersion.String()
	ServiceAccountApplicationKey_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountApplicationKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountApplicationKey{}, &ServiceAccountApplicationKeyList{})
}