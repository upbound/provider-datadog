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

type FilterInitParameters struct {

	// following the span search syntax. Defaults to "*".
	// The search query - following the span search syntax. Defaults to `"*"`.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type FilterObservation struct {

	// following the span search syntax. Defaults to "*".
	// The search query - following the span search syntax. Defaults to `"*"`.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type FilterParameters struct {

	// following the span search syntax. Defaults to "*".
	// The search query - following the span search syntax. Defaults to `"*"`.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`
}

type RetentionFilterInitParameters struct {

	// (Boolean) the status of the retention filter.
	// the status of the retention filter.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block, Optional) The spans filter. Spans matching this filter will be indexed and stored. (see below for nested schema)
	// The spans filter. Spans matching this filter will be indexed and stored.
	Filter *FilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// processing-sampling is available. Valid values are spans-sampling-processor.
	// The type of the retention filter, currently only spans-processing-sampling is available. Valid values are `spans-sampling-processor`.
	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`

	// (String) The name of the retention filter.
	// The name of the retention filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Sample rate to apply to spans going through this retention filter as a string, a value of 1.0 keeps all spans matching the query.
	// Sample rate to apply to spans going through this retention filter as a string, a value of 1.0 keeps all spans matching the query.
	Rate *string `json:"rate,omitempty" tf:"rate,omitempty"`
}

type RetentionFilterObservation struct {

	// (Boolean) the status of the retention filter.
	// the status of the retention filter.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block, Optional) The spans filter. Spans matching this filter will be indexed and stored. (see below for nested schema)
	// The spans filter. Spans matching this filter will be indexed and stored.
	Filter *FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// processing-sampling is available. Valid values are spans-sampling-processor.
	// The type of the retention filter, currently only spans-processing-sampling is available. Valid values are `spans-sampling-processor`.
	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the retention filter.
	// The name of the retention filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Sample rate to apply to spans going through this retention filter as a string, a value of 1.0 keeps all spans matching the query.
	// Sample rate to apply to spans going through this retention filter as a string, a value of 1.0 keeps all spans matching the query.
	Rate *string `json:"rate,omitempty" tf:"rate,omitempty"`
}

type RetentionFilterParameters struct {

	// (Boolean) the status of the retention filter.
	// the status of the retention filter.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (Block, Optional) The spans filter. Spans matching this filter will be indexed and stored. (see below for nested schema)
	// The spans filter. Spans matching this filter will be indexed and stored.
	// +kubebuilder:validation:Optional
	Filter *FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// processing-sampling is available. Valid values are spans-sampling-processor.
	// The type of the retention filter, currently only spans-processing-sampling is available. Valid values are `spans-sampling-processor`.
	// +kubebuilder:validation:Optional
	FilterType *string `json:"filterType,omitempty" tf:"filter_type,omitempty"`

	// (String) The name of the retention filter.
	// The name of the retention filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Sample rate to apply to spans going through this retention filter as a string, a value of 1.0 keeps all spans matching the query.
	// Sample rate to apply to spans going through this retention filter as a string, a value of 1.0 keeps all spans matching the query.
	// +kubebuilder:validation:Optional
	Rate *string `json:"rate,omitempty" tf:"rate,omitempty"`
}

// RetentionFilterSpec defines the desired state of RetentionFilter
type RetentionFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RetentionFilterParameters `json:"forProvider"`
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
	InitProvider RetentionFilterInitParameters `json:"initProvider,omitempty"`
}

// RetentionFilterStatus defines the observed state of RetentionFilter.
type RetentionFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RetentionFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RetentionFilter is the Schema for the RetentionFilters API. The object describing the configuration of the retention filter to create/update.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type RetentionFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filterType) || (has(self.initProvider) && has(self.initProvider.filterType))",message="spec.forProvider.filterType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rate) || (has(self.initProvider) && has(self.initProvider.rate))",message="spec.forProvider.rate is a required parameter"
	Spec   RetentionFilterSpec   `json:"spec"`
	Status RetentionFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RetentionFilterList contains a list of RetentionFilters
type RetentionFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RetentionFilter `json:"items"`
}

// Repository type metadata.
var (
	RetentionFilter_Kind             = "RetentionFilter"
	RetentionFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RetentionFilter_Kind}.String()
	RetentionFilter_KindAPIVersion   = RetentionFilter_Kind + "." + CRDGroupVersion.String()
	RetentionFilter_GroupVersionKind = CRDGroupVersion.WithKind(RetentionFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&RetentionFilter{}, &RetentionFilterList{})
}