/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/linode/provider-linode/apis/domain/v1alpha1"
	v1alpha1domain_record "github.com/linode/provider-linode/apis/domain_record/v1alpha1"
	v1alpha1firewall "github.com/linode/provider-linode/apis/firewall/v1alpha1"
	v1alpha1firewall_device "github.com/linode/provider-linode/apis/firewall_device/v1alpha1"
	v1alpha1lke_cluster "github.com/linode/provider-linode/apis/lke_cluster/v1alpha1"
	v1alpha1stackscript "github.com/linode/provider-linode/apis/stackscript/v1alpha1"
	v1alpha1apis "github.com/linode/provider-linode/apis/v1alpha1"
	v1beta1 "github.com/linode/provider-linode/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1domain_record.SchemeBuilder.AddToScheme,
		v1alpha1firewall.SchemeBuilder.AddToScheme,
		v1alpha1firewall_device.SchemeBuilder.AddToScheme,
		v1alpha1lke_cluster.SchemeBuilder.AddToScheme,
		v1alpha1stackscript.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}