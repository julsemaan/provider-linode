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

type DomainObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DomainParameters struct {

	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	// +kubebuilder:validation:Optional
	AxfrIps []*string `json:"axfrIps,omitempty" tf:"axfr_ips,omitempty"`

	// A description for this Domain. This is for display purposes only.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// The amount of time in seconds that may pass before this Domain is no longer Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +kubebuilder:validation:Optional
	ExpireSec *float64 `json:"expireSec,omitempty" tf:"expire_sec,omitempty"`

	// The group this Domain belongs to. This is for display purposes only.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// The IP addresses representing the master DNS for this Domain.
	// +kubebuilder:validation:Optional
	MasterIps []*string `json:"masterIps,omitempty" tf:"master_ips,omitempty"`

	// The amount of time in seconds before this Domain should be refreshed. Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +kubebuilder:validation:Optional
	RefreshSec *float64 `json:"refreshSec,omitempty" tf:"refresh_sec,omitempty"`

	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +kubebuilder:validation:Optional
	RetrySec *float64 `json:"retrySec,omitempty" tf:"retry_sec,omitempty"`

	// Start of Authority email address. This is required for master Domains.
	// +kubebuilder:validation:Optional
	SoaEmail *string `json:"soaEmail,omitempty" tf:"soa_email,omitempty"`

	// Used to control whether this Domain is currently being rendered.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +kubebuilder:validation:Optional
	TTLSec *float64 `json:"ttlSec,omitempty" tf:"ttl_sec,omitempty"`

	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	_v1.ResourceSpec `json:",inline"`
	ForProvider      DomainParameters `json:"forProvider"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	_v1.ResourceStatus `json:",inline"`
	AtProvider         DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec"`
	Status            DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
