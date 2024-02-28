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

type ServiceAccountInitParameters struct {

	// (Boolean) Whether the service account is disabled. Defaults to false.
	// Whether the service account is disabled. Defaults to `false`.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// (String) Email of the associated user.
	// Email of the associated user.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Name for the service account.
	// Name for the service account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list a role IDs to assign to the service account.
	// A list a role IDs to assign to the service account.
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type ServiceAccountObservation struct {

	// (Boolean) Whether the service account is disabled. Defaults to false.
	// Whether the service account is disabled. Defaults to `false`.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// (String) Email of the associated user.
	// Email of the associated user.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name for the service account.
	// Name for the service account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list a role IDs to assign to the service account.
	// A list a role IDs to assign to the service account.
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type ServiceAccountParameters struct {

	// (Boolean) Whether the service account is disabled. Defaults to false.
	// Whether the service account is disabled. Defaults to `false`.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// (String) Email of the associated user.
	// Email of the associated user.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Name for the service account.
	// Name for the service account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) A list a role IDs to assign to the service account.
	// A list a role IDs to assign to the service account.
	// +kubebuilder:validation:Optional
	// +listType=set
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

// ServiceAccountSpec defines the desired state of ServiceAccount
type ServiceAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountParameters `json:"forProvider"`
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
	InitProvider ServiceAccountInitParameters `json:"initProvider,omitempty"`
}

// ServiceAccountStatus defines the observed state of ServiceAccount.
type ServiceAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceAccount is the Schema for the ServiceAccounts API. Provides a Datadog service account resource. This can be used to create and manage Datadog service accounts.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type ServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	Spec   ServiceAccountSpec   `json:"spec"`
	Status ServiceAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountList contains a list of ServiceAccounts
type ServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccount `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccount_Kind             = "ServiceAccount"
	ServiceAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccount_Kind}.String()
	ServiceAccount_KindAPIVersion   = ServiceAccount_Kind + "." + CRDGroupVersion.String()
	ServiceAccount_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccount{}, &ServiceAccountList{})
}