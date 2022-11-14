/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	_v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RdnsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RdnsParameters struct {

	// The Public IPv4 or IPv6 address that will receive the PTR record.  A matching A or AAAA record must exist.
	// The public Linode IPv4 or IPv6 address to operate on.
	// +kubebuilder:validation:Required
	Address *string `json:"address" tf:"address,omitempty"`

	// The name of the RDNS address.
	// The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.
	// +kubebuilder:validation:Required
	Rdns *string `json:"rdns" tf:"rdns,omitempty"`

	// If true, the RDNS assignment will be retried within the operation timeout period.
	// If true, the RDNS assignment will be retried within the operation timeout period.
	// +kubebuilder:validation:Optional
	WaitForAvailable *bool `json:"waitForAvailable,omitempty" tf:"wait_for_available,omitempty"`
}

// RdnsSpec defines the desired state of Rdns
type RdnsSpec struct {
	_v1.ResourceSpec `json:",inline"`
	ForProvider      RdnsParameters `json:"forProvider"`
}

// RdnsStatus defines the observed state of Rdns.
type RdnsStatus struct {
	_v1.ResourceStatus `json:",inline"`
	AtProvider         RdnsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rdns is the Schema for the Rdnss API. Manages the RDNS / PTR record for the IP Address associated with a Linode Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Rdns struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdnsSpec   `json:"spec"`
	Status            RdnsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdnsList contains a list of Rdnss
type RdnsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rdns `json:"items"`
}

// Repository type metadata.
var (
	Rdns_Kind             = "Rdns"
	Rdns_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rdns_Kind}.String()
	Rdns_KindAPIVersion   = Rdns_Kind + "." + CRDGroupVersion.String()
	Rdns_GroupVersionKind = CRDGroupVersion.WithKind(Rdns_Kind)
)

func init() {
	SchemeBuilder.Register(&Rdns{}, &RdnsList{})
}