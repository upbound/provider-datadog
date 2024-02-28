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

type DashboardJSONInitParameters struct {

	// (String) The JSON formatted definition of the Dashboard.
	// The JSON formatted definition of the Dashboard.
	Dashboard *string `json:"dashboard,omitempty" tf:"dashboard,omitempty"`

	// (Set of Number) A list of dashboard lists this dashboard belongs to.
	// A list of dashboard lists this dashboard belongs to.
	// +listType=set
	DashboardLists []*float64 `json:"dashboardLists,omitempty" tf:"dashboard_lists,omitempty"`

	// (String) The URL of the dashboard.
	// The URL of the dashboard.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type DashboardJSONObservation struct {

	// (String) The JSON formatted definition of the Dashboard.
	// The JSON formatted definition of the Dashboard.
	Dashboard *string `json:"dashboard,omitempty" tf:"dashboard,omitempty"`

	// (Set of Number) A list of dashboard lists this dashboard belongs to.
	// A list of dashboard lists this dashboard belongs to.
	// +listType=set
	DashboardLists []*float64 `json:"dashboardLists,omitempty" tf:"dashboard_lists,omitempty"`

	// (Set of Number) The list of dashboard lists this dashboard should be removed from. Internal only.
	// The list of dashboard lists this dashboard should be removed from. Internal only.
	// +listType=set
	DashboardListsRemoved []*float64 `json:"dashboardListsRemoved,omitempty" tf:"dashboard_lists_removed,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The URL of the dashboard.
	// The URL of the dashboard.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type DashboardJSONParameters struct {

	// (String) The JSON formatted definition of the Dashboard.
	// The JSON formatted definition of the Dashboard.
	// +kubebuilder:validation:Optional
	Dashboard *string `json:"dashboard,omitempty" tf:"dashboard,omitempty"`

	// (Set of Number) A list of dashboard lists this dashboard belongs to.
	// A list of dashboard lists this dashboard belongs to.
	// +kubebuilder:validation:Optional
	// +listType=set
	DashboardLists []*float64 `json:"dashboardLists,omitempty" tf:"dashboard_lists,omitempty"`

	// (String) The URL of the dashboard.
	// The URL of the dashboard.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// DashboardJSONSpec defines the desired state of DashboardJSON
type DashboardJSONSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DashboardJSONParameters `json:"forProvider"`
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
	InitProvider DashboardJSONInitParameters `json:"initProvider,omitempty"`
}

// DashboardJSONStatus defines the observed state of DashboardJSON.
type DashboardJSONStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DashboardJSONObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DashboardJSON is the Schema for the DashboardJSONs API. Provides a Datadog dashboard JSON resource. This can be used to create and manage Datadog dashboards using the JSON definition.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type DashboardJSON struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dashboard) || (has(self.initProvider) && has(self.initProvider.dashboard))",message="spec.forProvider.dashboard is a required parameter"
	Spec   DashboardJSONSpec   `json:"spec"`
	Status DashboardJSONStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DashboardJSONList contains a list of DashboardJSONs
type DashboardJSONList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DashboardJSON `json:"items"`
}

// Repository type metadata.
var (
	DashboardJSON_Kind             = "DashboardJSON"
	DashboardJSON_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DashboardJSON_Kind}.String()
	DashboardJSON_KindAPIVersion   = DashboardJSON_Kind + "." + CRDGroupVersion.String()
	DashboardJSON_GroupVersionKind = CRDGroupVersion.WithKind(DashboardJSON_Kind)
)

func init() {
	SchemeBuilder.Register(&DashboardJSON{}, &DashboardJSONList{})
}