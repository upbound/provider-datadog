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

type ArchiveOrderInitParameters struct {

	// (List of String) The archive IDs list. The order of archive IDs in this attribute defines the overall archive order for logs. If archive_ids is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	// The archive IDs list. The order of archive IDs in this attribute defines the overall archive order for logs. If `archive_ids` is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	ArchiveIds []*string `json:"archiveIds,omitempty" tf:"archive_ids,omitempty"`
}

type ArchiveOrderObservation struct {

	// (List of String) The archive IDs list. The order of archive IDs in this attribute defines the overall archive order for logs. If archive_ids is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	// The archive IDs list. The order of archive IDs in this attribute defines the overall archive order for logs. If `archive_ids` is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	ArchiveIds []*string `json:"archiveIds,omitempty" tf:"archive_ids,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ArchiveOrderParameters struct {

	// (List of String) The archive IDs list. The order of archive IDs in this attribute defines the overall archive order for logs. If archive_ids is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	// The archive IDs list. The order of archive IDs in this attribute defines the overall archive order for logs. If `archive_ids` is empty or not specified, it will import the actual archive order, and create the resource. Otherwise, it will try to update the order.
	// +kubebuilder:validation:Optional
	ArchiveIds []*string `json:"archiveIds,omitempty" tf:"archive_ids,omitempty"`
}

// ArchiveOrderSpec defines the desired state of ArchiveOrder
type ArchiveOrderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ArchiveOrderParameters `json:"forProvider"`
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
	InitProvider ArchiveOrderInitParameters `json:"initProvider,omitempty"`
}

// ArchiveOrderStatus defines the observed state of ArchiveOrder.
type ArchiveOrderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ArchiveOrderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ArchiveOrder is the Schema for the ArchiveOrders API. Provides a Datadog Logs Archive API https://docs.datadoghq.com/api/v2/logs-archives/ resource, which is used to manage Datadog log archives order.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,datadog}
type ArchiveOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ArchiveOrderSpec   `json:"spec"`
	Status            ArchiveOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ArchiveOrderList contains a list of ArchiveOrders
type ArchiveOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArchiveOrder `json:"items"`
}

// Repository type metadata.
var (
	ArchiveOrder_Kind             = "ArchiveOrder"
	ArchiveOrder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ArchiveOrder_Kind}.String()
	ArchiveOrder_KindAPIVersion   = ArchiveOrder_Kind + "." + CRDGroupVersion.String()
	ArchiveOrder_GroupVersionKind = CRDGroupVersion.WithKind(ArchiveOrder_Kind)
)

func init() {
	SchemeBuilder.Register(&ArchiveOrder{}, &ArchiveOrderList{})
}