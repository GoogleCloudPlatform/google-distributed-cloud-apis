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

	corev1 "k8s.io/api/core/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +gdcloud:manifest:relevant=true,oc=platauth,component=pki,entities="certificate-requests"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="describe,list:certificate-authority-service-operation-manager"
// +gdcloud:manifest:rbac="create,describe,list:certificate-authority-service-certificate-requester"
// +gdcloud:manifest:rbac="create,delete,describe,list,update:certificate-authority-service-admin"
//
// CertificateRequest represents a request to issue certificate from the
// referenced CertificateAuthority.
//
// All fields within the CertificateRequest's `spec` are immutable after creation.
type CertificateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateRequestSpec   `json:"spec"`
	Status CertificateRequestStatus `json:"status,omitempty"`
}

// CertificateRequestSpec defines a request for the issuance of a certificate.
type CertificateRequestSpec struct {
	// Only one of the CSR or CertificateConfig should be set.

	// Certificate Signing Request to sign using CA.
	// +optional
	CSR []byte `json:"csr,omitempty"`

	// certificate config that will be use to create the CSR.
	// +optional
	CertificateConfig *CertificateConfig `json:"certificateConfig,omitempty"`

	// Validity start time of the certificate.
	// If not set, we will use the current time of the request.
	NotBefore *metav1.Time `json:"notBefore,omitempty"`

	// Validity end time of the certificate.
	// If not set, we will use 90 days from the notBefore time as default.
	NotAfter *metav1.Time `json:"notAfter,omitempty"`

	// Name of the secret to store the signed certificate.
	SignedCertificateSecret string `json:"signedCertificateSecret"`

	// A reference to a CertificateAuthority which will sign the certificate.
	// API type:
	//   - Group: pki.security.gdc.goog
	//   - Kind: CertificateAuthority
	CertificateAuthorityRef CAReference `json:"certificateAuthorityRef"`

	// A template used to issue the certificate.
	// +optional
	CertificateTemplate *string `json:"certificateTemplate,omitempty"`

	// SubjectOverride is a raw, ASN.1 DER-encoded X.509 subject.
	// If set, this raw subject will be used instead of the raw subject in the CSR.
	// +optional
	SubjectOverride []byte `json:"subjectOverride,omitempty"`
}

// Supported template types.
const (
	// EndEntityClientAuthCertificate is a template for a client TLS certificate.
	// It has predefined key usages for digital signature and key encipherment,
	// and an extended key usage for client authentication.
	// The basic constraints are set to non-critical, and Subject Alternative Names (SANs) are taken from the CSR.
	EndEntityClientAuthCertificate string = "endEntityClientAuthCertificate"
	// EndEntityServerAuthCertificate is a template for a server TLS certificate.
	// It has predefined key usages for digital signature and key encipherment,
	// and an extended key usage for server authentication.
	// The basic constraints are set to non-critical, and Subject Alternative Names (SANs) are taken from the CSR.
	EndEntityServerAuthCertificate string = "endEntityServerAuthCertificate"
	// BlankSubCACertificate_PathLen0_CSRPassthrough is a template for a subordinate CA certificate with a path length of 0.
	// Key usages, extended key usages, and Subject Alternative Names (SANs) are passed through from the Certificate Signing Request (CSR).
	BlankSubCACertificate_PathLen0_CSRPassthrough string = "blankSubCACertificate_PathLen0_CSRPassthrough"
	// BlankSubCACertificate_PathLen1_CSRPassthrough is a template for a subordinate CA certificate with a path length of 1.
	// Key usages, extended key usages, and Subject Alternative Names (SANs) are passed through from the Certificate Signing Request (CSR).
	BlankSubCACertificate_PathLen1_CSRPassthrough string = "blankSubCACertificate_PathLen1_CSRPassthrough"
	// BlankSubCACertificate_CSRPassthrough is a template for a subordinate CA certificate.
	// The CA capability (IsCA) is always set to true. The path length constraint (MaxPathLen), key usages, extended key usages, and Subject Alternative Names (SANs) are passed through from the Certificate Signing Request (CSR).
	BlankSubCACertificate_CSRPassthrough string = "blankSubCACertificate_CSRPassthrough"
)

// CertificateConfig represents the subject information in an issued certificate.
type CertificateConfig struct {
	// These values are used to create the distinguished name and
	// subject alternative name fields in an X.509 certificate.
	SubjectConfig SubjectConfig `json:"subjectConfig"`

	// Private key options. These include the key algorithm and size.
	// +optional
	PrivateKeyConfig *CertificatePrivateKey `json:"privateKeyConfig,omitempty"`
}

type CertificatePrivateKey struct {
	// Algorithm is the private key algorithm of the corresponding private key
	// for this certificate. If provided, allowed values are either `RSA`,`Ed25519` or `ECDSA`
	// If `algorithm` is specified and `size` is not provided,
	// key size of 384 will be used for `ECDSA` key algorithm and
	// key size of 3072 will be used for `RSA` key algorithm.
	// key size is ignored when using the `Ed25519` key algorithm.
	// See github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1/types_certificate.go for more information.
	// +optional
	Algorithm PrivateKeyAlgorithm `json:"algorithm,omitempty"`

	// Size is the key bit size of the corresponding private key for this certificate.
	// If `algorithm` is set to `RSA`, valid values are `2048`, `3072`, `4096` or `8192`,
	// and will default to `3072` if not specified.
	// If `algorithm` is set to `ECDSA`, valid values are `256`, `384` or `521`,
	// and will default to `384` if not specified.
	// If `algorithm` is set to `Ed25519`, Size is ignored.
	// No other values are allowed.
	// See github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1/types_certificate.go for more information.
	// +optional
	Size int `json:"size,omitempty"`
}

type SubjectConfig struct {
	// The common name of the Certificate.
	// +optional
	CommonName string `json:"commonName"`

	// The organization of the Certificate.
	// +optional
	Organization *string `json:"organization,omitempty"`

	// The locality of the Certificate.
	// +optional
	Locality *string `json:"locality,omitempty"`

	// The state of the Certificate.
	// +optional
	State *string `json:"state,omitempty"`

	// The country of the Certificate.
	// +optional
	Country *string `json:"country,omitempty"`

	// DNSNames is a list of dNSName subjectAltNames to be set on the Certificate.
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`

	// IPAddresses is a list of ipAddress subjectAltNames to be set on the Certificate.
	// +optional
	IPAddresses []string `json:"ipAddresses,omitempty"`

	// RFC822Names is a list of rfc822Name subjectAltNames to be set on the Certificate.
	// +optional
	RFC822Names []string `json:"rfc822Names,omitempty"`

	// URIs is a list of uniformResourceIdentifier subjectAltNames to be set on the Certificate.
	// +optional
	URIs []string `json:"uris,omitempty"`
}

type CertificateRequestStatus struct {
	// List of status conditions to indicate the status of a certificate to be issued.
	// - PENDING: CSR are pending to be signed.
	// - Ready: Indicates that the certificateRequest is fulfilled.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// If no CSR is provided, an auto-generated private key will be used.
	// optional
	AutoGeneratedPrivateKey *corev1.SecretReference `json:"autoGeneratedPrivateKey,omitempty"`
}

// +kubebuilder:object:root=true
// CertificateRequestList represents a collection of certiifcate requests.
type CertificateRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&CertificateRequest{},
		&CertificateRequestList{},
	)
}
