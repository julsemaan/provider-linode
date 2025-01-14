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

type ConfigObservation struct {

	// What algorithm this NodeBalancer should use for routing traffic to backends. (roundrobin, leastconn, source)
	// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected. (none, connection, http, http_body)
	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.
	Check *string `json:"check,omitempty" tf:"check,omitempty"`

	// How many times to attempt a check before considering a backend to be down. (1-30)
	// How many times to attempt a check before considering a backend to be down. (1-30)
	CheckAttempts *float64 `json:"checkAttempts,omitempty" tf:"check_attempts,omitempty"`

	// This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down
	CheckBody *string `json:"checkBody,omitempty" tf:"check_body,omitempty"`

	// How often, in seconds, to check that backends are up and serving requests.
	// How often, in seconds, to check that backends are up and serving requests.
	CheckInterval *float64 `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`

	// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
	// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
	CheckPassive *bool `json:"checkPassive,omitempty" tf:"check_passive,omitempty"`

	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
	CheckPath *string `json:"checkPath,omitempty" tf:"check_path,omitempty"`

	// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
	// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
	CheckTimeout *float64 `json:"checkTimeout,omitempty" tf:"check_timeout,omitempty"`

	// What ciphers to use for SSL connections served by this NodeBalancer. legacy is considered insecure and should only be used if necessary.
	// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
	CipherSuite *string `json:"cipherSuite,omitempty" tf:"cipher_suite,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A structure containing information about the health of the backends for this port. This information is updated periodically as checks are performed against backends.
	NodeStatus []NodeStatusObservation `json:"nodeStatus,omitempty" tf:"node_status,omitempty"`

	// The ID of the NodeBalancer to access.
	// The ID of the NodeBalancer to access.
	NodebalancerID *float64 `json:"nodebalancerId,omitempty" tf:"nodebalancer_id,omitempty"`

	// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. (Defaults to 80)
	// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (http, https, tcp) (Defaults to http)
	// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be tcp. (none, v1, v2) (Defaults to none)
	// The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be `tcp`. Valid values are `none`, `v1`, and `v2`.
	ProxyProtocol *string `json:"proxyProtocol,omitempty" tf:"proxy_protocol,omitempty"`

	// The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	// The read-only common name automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	SSLCommonname *string `json:"sslCommonname,omitempty" tf:"ssl_commonname,omitempty"`

	// The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	// The read-only fingerprint automatically derived from the SSL certificate assigned to this NodeBalancerConfig. Please refer to this field to verify that the appropriate certificate is assigned to your NodeBalancerConfig.
	SSLFingerprint *string `json:"sslFingerprint,omitempty" tf:"ssl_fingerprint,omitempty"`

	// Controls how session stickiness is handled on this port. (none, table, http_cookie)
	// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
	Stickiness *string `json:"stickiness,omitempty" tf:"stickiness,omitempty"`
}

type ConfigParameters struct {

	// What algorithm this NodeBalancer should use for routing traffic to backends. (roundrobin, leastconn, source)
	// What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected. (none, connection, http, http_body)
	// The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.
	// +kubebuilder:validation:Optional
	Check *string `json:"check,omitempty" tf:"check,omitempty"`

	// How many times to attempt a check before considering a backend to be down. (1-30)
	// How many times to attempt a check before considering a backend to be down. (1-30)
	// +kubebuilder:validation:Optional
	CheckAttempts *float64 `json:"checkAttempts,omitempty" tf:"check_attempts,omitempty"`

	// This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down
	// +kubebuilder:validation:Optional
	CheckBody *string `json:"checkBody,omitempty" tf:"check_body,omitempty"`

	// How often, in seconds, to check that backends are up and serving requests.
	// How often, in seconds, to check that backends are up and serving requests.
	// +kubebuilder:validation:Optional
	CheckInterval *float64 `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`

	// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
	// If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.
	// +kubebuilder:validation:Optional
	CheckPassive *bool `json:"checkPassive,omitempty" tf:"check_passive,omitempty"`

	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
	// The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.
	// +kubebuilder:validation:Optional
	CheckPath *string `json:"checkPath,omitempty" tf:"check_path,omitempty"`

	// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
	// How long, in seconds, to wait for a check attempt before considering it failed. (1-30)
	// +kubebuilder:validation:Optional
	CheckTimeout *float64 `json:"checkTimeout,omitempty" tf:"check_timeout,omitempty"`

	// What ciphers to use for SSL connections served by this NodeBalancer. legacy is considered insecure and should only be used if necessary.
	// What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.
	// +kubebuilder:validation:Optional
	CipherSuite *string `json:"cipherSuite,omitempty" tf:"cipher_suite,omitempty"`

	// The ID of the NodeBalancer to access.
	// The ID of the NodeBalancer to access.
	// +crossplane:generate:reference:type=Nodebalancer
	// +kubebuilder:validation:Optional
	NodebalancerID *float64 `json:"nodebalancerId,omitempty" tf:"nodebalancer_id,omitempty"`

	// Reference to a Nodebalancer to populate nodebalancerId.
	// +kubebuilder:validation:Optional
	NodebalancerIDRef *v1.Reference `json:"nodebalancerIdRef,omitempty" tf:"-"`

	// Selector for a Nodebalancer to populate nodebalancerId.
	// +kubebuilder:validation:Optional
	NodebalancerIDSelector *v1.Selector `json:"nodebalancerIdSelector,omitempty" tf:"-"`

	// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443. (Defaults to 80)
	// The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key. (http, https, tcp) (Defaults to http)
	// The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be tcp. (none, v1, v2) (Defaults to none)
	// The version of ProxyProtocol to use for the underlying NodeBalancer. This requires protocol to be `tcp`. Valid values are `none`, `v1`, and `v2`.
	// +kubebuilder:validation:Optional
	ProxyProtocol *string `json:"proxyProtocol,omitempty" tf:"proxy_protocol,omitempty"`

	// The certificate this port is serving. This is not returned. If set, this field will come back as <REDACTED>. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
	// The certificate this port is serving. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
	// +kubebuilder:validation:Optional
	SSLCertSecretRef *v1.SecretKeySelector `json:"sslCertSecretRef,omitempty" tf:"-"`

	// The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as <REDACTED>. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
	// The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.
	// +kubebuilder:validation:Optional
	SSLKeySecretRef *v1.SecretKeySelector `json:"sslKeySecretRef,omitempty" tf:"-"`

	// Controls how session stickiness is handled on this port. (none, table, http_cookie)
	// Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'
	// +kubebuilder:validation:Optional
	Stickiness *string `json:"stickiness,omitempty" tf:"stickiness,omitempty"`
}

type NodeStatusObservation struct {

	// The number of backends considered to be 'DOWN' and unhealthy. These are not in rotation, and not serving requests.
	Down *float64 `json:"down,omitempty" tf:"down,omitempty"`

	// The number of backends considered to be 'UP' and healthy, and that are serving requests.
	Up *float64 `json:"up,omitempty" tf:"up,omitempty"`
}

type NodeStatusParameters struct {
}

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConfigParameters `json:"forProvider"`
}

// ConfigStatus defines the observed state of Config.
type ConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Config is the Schema for the Configs API. Manages a Linode NodeBalancer Config.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigSpec   `json:"spec"`
	Status            ConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigList contains a list of Configs
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

// Repository type metadata.
var (
	Config_Kind             = "Config"
	Config_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Config_Kind}.String()
	Config_KindAPIVersion   = Config_Kind + "." + CRDGroupVersion.String()
	Config_GroupVersionKind = CRDGroupVersion.WithKind(Config_Kind)
)

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
