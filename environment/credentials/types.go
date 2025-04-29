// +groupName=credentials.nutanix.com
package credentials

// This file defines single-key secret format by encoding entire secret
// as JSON object.

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CredentialType describes authentication mechanism like basic auth.
type CredentialType string

const (
	// BasicAuthCredentialType is username/password based authentication.
	BasicAuthCredentialType CredentialType = "basic_auth"

	// KeyName is secret
	KeyName = "credentials"
)

// +kubebuilder:object:generate=true
type Credential struct {
	Type CredentialType  `json:"type"`
	Data json.RawMessage `json:"data"`
}

// NutanixCredentials is list of credentials to be embedded in other objects like
// Kubernetes secrets.
// +kubebuilder:object:generate=true
type NutanixCredentials struct {
	Credentials []Credential `json:"credentials"`
}

// BasicAuthCredential is payload in Credential.Data for type of BasicAuthCredentialType
type BasicAuthCredential struct {
	// The Basic Auth (username, password) for the Prism Central
	PrismCentral PrismCentralBasicAuth `json:"prismCentral"`
	// The Basic Auth (username, password) for the Prism Elements (clusters).
	PrismElements []PrismElementBasicAuth `json:"prismElements"`
}

// +kubebuilder:object:generate=true
type BasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// +kubebuilder:object:generate=true
type PrismCentralBasicAuth struct {
	BasicAuth `json:",inline"`
}

// +kubebuilder:object:generate=true
type PrismElementBasicAuth struct {
	BasicAuth `json:",inline"`
	// Name is the unique resource name of the Prism Element (cluster) in the Prism Central's domain
	Name string `json:"name"`
}

type NutanixCredentialKind string

const (
	// Secret kind is enum value
	SecretKind = NutanixCredentialKind("Secret")
)

// +kubebuilder:object:generate=true
type NutanixCredentialReference struct {
	// Kind of the Nutanix credential
	// +kubebuilder:validation:Enum=Secret
	Kind NutanixCredentialKind `json:"kind"`
	// Name of the credential.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// namespace of the credential.
	// +optional
	Namespace string `json:"namespace"`
}

type NutanixTrustBundleKind string

const (
	NutanixTrustBundleKindString    = NutanixTrustBundleKind("String")
	NutanixTrustBundleKindConfigMap = NutanixTrustBundleKind("ConfigMap")
)

// NutanixTrustBundleReference is a reference to a Nutanix trust bundle.
// +kubebuilder:object:generate=true
type NutanixTrustBundleReference struct {
	// Kind of the Nutanix trust bundle
	// +kubebuilder:validation:Enum=String;ConfigMap
	Kind NutanixTrustBundleKind `json:"kind"`
	// Data of the trust bundle if Kind is String.
	// +optional
	Data string `json:"data"`
	// Name of the credential.
	// +optional
	Name string `json:"name"`
	// namespace of the credential.
	// +optional
	Namespace string `json:"namespace"`
}

// NutanixPrismEndpoint defines a Nutanix API endpoint with reference to credentials.
// Credentials are stored in Kubernetes secrets.
// +kubebuilder:object:generate=true
type NutanixPrismEndpoint struct {
	// address is the endpoint address (DNS name or IP address) of the Nutanix Prism Central or Element (cluster)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=256
	Address string `json:"address"`
	// port is the port number to access the Nutanix Prism Central or Element (cluster)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default=9440
	Port int32 `json:"port"`
	// use insecure connection to Prism endpoint
	// +kubebuilder:default=false
	// +optional
	Insecure bool `json:"insecure,omitempty"`
	// AdditionalTrustBundle is a PEM encoded x509 cert for the RootCA that was used to create the certificate
	// for a Prism Central that uses certificates that were issued by a non-publicly trusted RootCA. The trust
	// bundle is added to the cert pool used to authenticate the TLS connection to the Prism Central.
	// +optional
	AdditionalTrustBundle *NutanixTrustBundleReference `json:"additionalTrustBundle,omitempty"`
	// Pass credential information for the target Prism instance
	// +optional
	CredentialRef *NutanixCredentialReference `json:"credentialRef,omitempty"`
}

// NamespaceSelector defines which namespaces can use an identity
// +kubebuilder:object:generate=true
type NamespaceSelector struct {
	// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions,
	// whose key field is "key", the operator is "In", and the values array contains only "value".
	// The requirements are ANDed.
	// +optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// AllowedNamespaces defines namespaces that are allowed to use the identity
// +kubebuilder:object:generate=true
type AllowedNamespaces struct {
	// Selector is a selector of namespaces that are allowed to use the identity
	// If not specified, no namespaces will be allowed
	// +optional
	Selector NamespaceSelector `json:"selector,omitempty"`
}

// NutanixPrismIdentity defines a Nutanix API endpoint with reference to credentials and namespace selectors.
// This type extends NutanixPrismEndpoint to add namespace selection for access control.
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nutanixprismidentities,scope=Cluster,shortName=npi,singular=nutanixprismidentity,categories=nutanix
// +kubebuilder:printcolumn:name="Address",type="string",JSONPath=".spec.address",description="API endpoint address"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type NutanixPrismIdentity struct {
	// Standard Kubernetes type metadata
	// +kubebuilder:validation:Required
	metav1.TypeMeta `json:",inline"`

	// Standard Kubernetes object metadata
	// +kubebuilder:validation:Required
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of NutanixPrismIdentity
	// +kubebuilder:validation:Required
	Spec NutanixPrismIdentitySpec `json:"spec"`

	// Status defines the observed state of NutanixPrismIdentity
	// +optional
	Status NutanixPrismIdentityStatus `json:"status,omitempty"`
}

// NutanixPrismIdentitySpec defines the desired state of NutanixPrismIdentity
// +kubebuilder:object:generate=true
type NutanixPrismIdentitySpec struct {
	// NutanixPrismEndpoint embedded for endpoint configuration
	NutanixPrismEndpoint `json:",inline"`
	// AllowedNamespaces defines which namespaces are allowed to use this identity
	// If empty, defaults to all namespaces
	// +optional
	AllowedNamespaces AllowedNamespaces `json:"allowedNamespaces,omitempty"`
}

// NutanixPrismIdentityStatus defines the observed state of NutanixPrismIdentity
// +kubebuilder:object:generate=true
type NutanixPrismIdentityStatus struct {
	// Conditions represents the latest available observations of the identity's state
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// LastVerified is the last time the credentials were verified to work
	// +optional
	LastVerified *metav1.Time `json:"lastVerified,omitempty"`
}

// NutanixPrismIdentityList contains a list of NutanixPrismIdentity
// +kubebuilder:object:root=true
type NutanixPrismIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NutanixPrismIdentity `json:"items"`
}
