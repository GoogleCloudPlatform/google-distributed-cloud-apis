// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +gdcloud:manifest:relevant=true,oc=platauth,component=pki,multigroup=true
// +gdcloud:manifest:entities="root-cas",verbs=create;describe;list;update;delete,rbac="create,describe,list,update,delete:certificate-authority-service-admin,certificate-authority-operation-manager"
// +gdcloud:manifest:entities="subordinate-cas",verbs=create;get-csr;describe;list;update;delete
// +gdcloud:manifest:rbac="create,get-csr,describe,list,update,delete:certificate-authority-service-admin,certificate-authority-operation-manager"

// CertificateAuthority represents the individual Certificate Authority that will be used to issue the certificates.
type CertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateAuthoritySpec   `json:"spec"`
	Status CertificateAuthorityStatus `json:"status,omitempty"`
}

// +kubebuilder:validation:XValidation:rule="(has(self.secretConfig) && has(self.secretConfig.secretName) && size(self.secretConfig.secretName) > 0) != has(self.kmsConfig)",message="Only one of secretConfig or kmsConfig must be set at a time"
type CertificateAuthoritySpec struct {
	// The profile of the CertificateAuthority.
	// +optional
	CAProfile *CACertificateProfile `json:"caProfile,omitempty"`

	// The CA Certificate provisioning configuration.
	CACertificate CACertificateConfig `json:"caCertificate"`

	// Configuration of the CA secret.
	// +optional
	SecretConfig SecretConfig `json:"secretConfig,omitempty"`

	// Configuration of the CA private keys stored in KMS and optional pre-existing
	// certificates
	// +optional
	KMSConfig *KMSConfig `json:"kmsConfig,omitempty"`

	// Defines the profile of the certificates that will be issued.
	// +optional
	CertificateProfile *CertificateProfile `json:"certificateProfile,omitempty"`

	// Config related to enable ACME protocol.
	// +optional
	ACME *ACMEConfig `json:"acme,omitempty"`

	// Configuration of the Certificate Revocation List (CRL).
	// +optional
	CRL *CRLConfig `json:"crl,omitempty"`

	// Reference to the CaPool this CA belongs to.
	// +optional
	CaPoolRef *string `json:"caPoolRef,omitempty"`
}

type CRLConfig struct {
	// Whether to enable CRL or not. Defaults to true if not specified.
	// +optional
	// +kubebuilder:default=true
	Enabled *bool `json:"enabled,omitempty"`

	// The validity of the CRL. Defaults to 7 days if not specified.
	// +optional
	Validity *metav1.Duration `json:"validity,omitempty"`

	// RenewBefore defines the period before the CRL expires when a new CRL is generated
	// and published. If unset, this defaults to 1/3 of the CRL's validity period.
	// Also, a new CRL is published immediately upon any certificate revocation.
	// +optional
	RenewBefore *metav1.Duration `json:"renewBefore,omitempty"`
}

type KMSConfig struct {
	// A reference to a SigningKey acting as a CA private key.
	SigningKeyRef corev1alpha1.NamespacedName `json:"signingKeyRef"`

	// Stores preexisting signed certificate signed by key from SigningKeyRef
	// +optional
	SignedCertificate *SignedCertificateConfig `json:"signedCertificate,omitempty"`
}

// CACertificateProfile defines the profile for a CA certificate.
type CACertificateProfile struct {
	// The common name of the CA Certificate.
	CommonName string `json:"commonName"`

	// Organizations to be used on the Certificate.
	// +optional
	Organizations []string `json:"organizations,omitempty"`

	// Countries to be used on the Certificate.
	// +optional
	Countries []string `json:"countries,omitempty"`

	// Organizational Units to be used on the Certificate.
	// +optional
	OrganizationalUnits []string `json:"organizationalUnits,omitempty"`

	// Cities to be used on the Certificate.
	// +optional
	Localities []string `json:"localities,omitempty"`

	// State/Provinces to be used on the Certificate.
	// +optional
	Provinces []string `json:"provinces,omitempty"`

	// Street addresses to be used on the Certificate.
	// +optional
	StreetAddresses []string `json:"streetAddresses,omitempty"`

	// Postal codes to be used on the Certificate.
	// +optional
	PostalCodes []string `json:"postalCodes,omitempty"`

	// The requested 'duration' (i.e. lifetime) of the CA Certificate.
	Duration metav1.Duration `json:"duration"`

	// RenewBefore implies the rotation time before the CA certificate expires.
	// +optional
	RenewBefore *metav1.Duration `json:"renewBefore,omitempty"`

	// The maximum path length of the CA certificate.
	// +optional
	MaxPathLength *int `json:"maxPathLength,omitempty"`
}

// CACertificateConfig defines how the CA certificate is going to be provisioned.
// Only one of them will be set at any point in time.
type CACertificateConfig struct {
	// Get the certificate from an external root CA.
	// If set, a CSR will be generated on the status and signed certificate
	// can be upload using this field.
	// +optional
	ExternalCA *ExternalCAConfig `json:"externalCA,omitempty"`

	// Issue a self-signed certificate. (Root CA)
	// +optional
	SelfSignedCA *SelfSignedCAConfig `json:"selfSignedCA,omitempty"`

	// Issue a SubCA certificate from a GDC-managed CA. (Managed Sub CA)
	// +optional
	ManagedSubCA *ManagedSubCAConfig `json:"managedSubCA,omitempty"`

	// Adopt a pre-existing key and certificate. (Imported CA)
	// +optional
	ImportedCA *ImportedCAConfig `json:"importedCA,omitempty"`
}

type ImportedCAConfig struct{}

type ExternalCAConfig struct {
	// Stores a signed certificate signed by external root CA.
	// +optional
	SignedCertificate *SignedCertificateConfig `json:"signedCertificate,omitempty"`
}

type SignedCertificateConfig struct {
	// The PEM encoded x509 certificate uploaded by the customer.
	Certificate []byte `json:"certificate"`

	// The PEM encoded x509 certificate of the signer CA used to sign the certificate.
	// +optional
	CA []byte `json:"ca,omitempty"`
}

// SelfSignedCAConfig defines the configuration for a Root CA certificate.
type SelfSignedCAConfig struct{}

// ManagedSubCAConfig defines the configuration for a SubCA CA certificate.
type ManagedSubCAConfig struct {
	// A reference to a CertificateAuthority which will sign the SubCA certificate.
	// API type:
	//	- Group: pki.security.gdc.goog
	//	- Kind: CertificateAuthority
	CertificateAuthorityRef CAReference `json:"certificateAuthorityRef"`
}

// CertificateProfile defines the specification of the profile of an issued certificate.
type CertificateProfile struct {
	// Allowed key usages for certificates issued under this profile.
	// +optional
	KeyUsage []KeyUsageBits `json:"keyUsage,omitempty"`
	// Allowed extended key usages for certificates issued under this profile.
	// +optional
	ExtendedKeyUsage []ExtendedKeyUsageBits `json:"extendedKeyUsage,omitempty"`
}

// KeyUsageBits defines the different allowed key usages according to RFC 5280 4.2.1.3. Note that
// many of the key usages below are used for certificates outside the context of TLS, and the
// implementation of setting non-TLS bits can be implemented as a later feature.
type KeyUsageBits string

const (
	KeyUsageDigitalSignature KeyUsageBits = "digitalSignature"
	KeyUsageNonRepudiation   KeyUsageBits = "nonRepudiation"
	KeyUsageKeyEncipherment  KeyUsageBits = "keyEncipherment"
	KeyUsageDataEncipherment KeyUsageBits = "dataEncipherment"
	KeyUsageKeyAgreement     KeyUsageBits = "keyAgreement"
	KeyUsageKeyCertSign      KeyUsageBits = "keyCertSign"
	KeyUsageCRLSign          KeyUsageBits = "crlSign"
	KeyUsageEncipherOnly     KeyUsageBits = "encipherOnly"
	KeyUsageDecipherOnly     KeyUsageBits = "decipherOnly"
)

// ExtendedKeyUsageBits defines the different allowed extended key usages according to RFC 5280 4.2.1.12.
// Many extended key usages have been defined by follow-up RFCs, and can be implemented as a later feature
// if issuance of such certificates is needed, for cases such as certificates used for personal
// authentication, code signing or IPSec.
type ExtendedKeyUsageBits string

const (
	ExtendedKeyUsageServerAuthentication ExtendedKeyUsageBits = "serverAuth"
	ExtendedKeyUsageClientAuthentication ExtendedKeyUsageBits = "clientAuth"
)

// SecretConfig defines the configuration for the certificate secret.
type SecretConfig struct {
	// The name of the Secret that will hold the private key and signed certificate.
	SecretName string `json:"secretName"`

	// Defines annotations and labels to be copied to the Secret.
	// +optional
	SecretTemplate *SecretTemplate `json:"secretTemplate,omitempty"`

	// Options for the certificate private key
	// +optional
	PrivateKeyConfig *PrivateKeyConfig `json:"privateKeyConfig,omitempty"`
}

// PrivateKeyConfig defines the configuration of the certificate private key
type PrivateKeyConfig struct {
	// Algorithm is the private key algorithm of the corresponding private key
	// for this certificate. If provided, allowed values are either `RSA`,`Ed25519` or `ECDSA`
	// If `algorithm` is specified and `size` is not provided,
	// key size of 384 will be used for `ECDSA` key algorithm and
	// key size of 3072 will be used for `RSA` key algorithm.
	// key size is ignored when using the `Ed25519` key algorithm.
	// See go/gdch-horizontals/horizontals/SECURITY/SEC-HZ#SEC-HZ-19 for more information.
	// +optional
	Algorithm PrivateKeyAlgorithm `json:"algorithm,omitempty"`

	// Size is the key bit size of the corresponding private key for this certificate.
	// If `algorithm` is set to `RSA`, valid values are `2048`, `3072`, `4096` or `8192`,
	// and will default to `3072` if not specified.
	// If `algorithm` is set to `ECDSA`, valid values are `256`, `384` or `521`,
	// and will default to `384` if not specified.
	// If `algorithm` is set to `Ed25519`, Size is ignored.
	// No other values are allowed.
	// See go/gdch-horizontals/horizontals/SECURITY/SEC-HZ#SEC-HZ-19 for more information.
	// +optional
	Size int `json:"size,omitempty"`

	// NOTE(hyyh): Future support for the user provided private key in secret or KMS can be added here
}

// SecretTemplate defines the default labels and annotations to be copied
// to the Kubernetes Secret resource named in `SecretConfig.SecretName`.
type SecretTemplate struct {
	// Annotations is a key value map to be copied to the target Kubernetes Secret.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels is a key value map to be copied to the target Kubernetes Secret.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// +kubebuilder:validation:Enum=RSA;ECDSA;Ed25519
type PrivateKeyAlgorithm string

const (
	// Denotes the RSA private key type.
	RSAKeyAlgorithm PrivateKeyAlgorithm = "RSA"

	// Denotes the ECDSA private key type.
	ECDSAKeyAlgorithm PrivateKeyAlgorithm = "ECDSA"

	// Denotes the Ed25519 private key type.
	Ed25519KeyAlgorithm PrivateKeyAlgorithm = "Ed25519"
)

type CertificateAuthorityStatus struct {
	// ExternalCA specifies status options for SunCA signed by External root CA.
	// +optional
	ExternalCA *ExternalCAStatus `json:"externalCA,omitempty"`

	// ErrorStatus contain a list of current errors and the timestamp this field gets updated.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// List of status conditions to indicate the status of a Certification Authority.
	// - Pending: CSR are pending to be signed by the customer.
	// - Ready: Indicates that the certificate authority is ready to use.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ACME specific status options.
	// This field should only be set if the Certificate Authority is configured with ACME enabled.
	// +optional
	ACME *ACMEStatus `json:"acme,omitempty"`

	// The URL of the certificate revocation list (CRL) for this CA.
	// +optional
	CRL *string `json:"crl,omitempty"`

	// The PEM encoded X.509 ca certificate chain consisting of intermediate CA certificates and root CA certificate.
	// +optional
	CertificateChain []byte `json:"certificate,omitempty"`
}

type ExternalCAStatus struct {
	// A certificate signing request waiting to be signed by an external CA.
	// +optional
	CSR []byte `json:"csr,omitempty"`
}

// CAReference represents a CertificateAuthority reference. It has information
// to retrieve a CA in any namespace.
type CAReference struct {
	// Name is unique within a namespace to reference a CA resource.
	Name string `json:"name"`
	// Namespace defines the space within which the CA name must be unique.
	Namespace string `json:"namespace"`
}

type ACMEConfig struct {
	// Whether to deploy and access CA via ACME protocol.
	// +optional
	// +kubebuilder:default=false
	Enabled *bool `json:"enabled,omitempty"`
}

type ACMEStatus struct {
	// URI is the unique account identifier, which can also be used to retrieve
	// account details from the CA
	// +optional
	URI string `json:"uri,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateAuthorityList represents a collection of certiifcate authorities.
type CertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateAuthority `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&CertificateAuthority{},
		&CertificateAuthorityList{},
	)
}
