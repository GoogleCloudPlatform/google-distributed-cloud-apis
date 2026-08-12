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
// +gdcloud:manifest:relevant=false,oc=platauth
// A Certificate represents a managed certificate.
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

type CertificateSpec struct {
	// A reference to the CertificateIssuer that will be used for the issuance of the certificate.
	// If not set, a label named `pki.security.gdc.goog/use-default-issuer: true` needs to be
	// set in order to issue the certificate using the default issuer.
	// API type:
	//   - Group: pki.security.gdc.goog
	//   - Kind: CertificateIssuer
	// +optional
	Issuer *IssuerReference `json:"issuer,omitempty"`

	// Requested common name X509 certificate subject attribute.
	// It should have a length of 64 characters or fewer.
	//
	// For backward compatibility, the behaviour is as follows:
	// If nil, we use the current behavior to set commonName as first DNSName if length is 64 characters or fewer.
	// if empty string, don't set it.
	// if it is set, ensure it is a part of the SANs.
	//
	// +optional
	CommonName *string `json:"commonName,omitempty"`

	// DNSNames is a list of fully-qualified host names to be set on the Certificate.
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`

	// IPAddresses is a list of IPAddress subjectAltNames to be set on the Certificate.
	// +optional
	IPAddresses []string `json:"ipAddresses,omitempty"`

	// The requested 'duration' (i.e. lifetime) of the Certificate.
	// +optional
	Duration *metav1.Duration `json:"duration,omitempty"`

	// RenewBefore implies the rotation time before the certificate expires.
	// +optional
	RenewBefore *metav1.Duration `json:"renewBefore,omitempty"`

	// Configuration of the Certificate secret.
	SecretConfig SecretConfig `json:"secretConfig"`

	// Contains the externally signed certificate
	// +optional
	BYOCertificate *BYOCertificate `json:"byoCertificate,omitempty"`
}

// Externally signed certificate
type BYOCertificate struct {
	// The PEM encoded x509 certificate uploaded by the customer.
	Certificate []byte `json:"certificate"`

	// The PEM encoded x509 certificate of the signer CA used to sign the certificate.
	CA []byte `json:"ca"`
}

type CertificateStatus struct {
	// List of status conditions to indicate the status of the certificate.
	// - Ready: Indicates that the certificate is ready to use.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" `

	// A reference to the CertificateIssuer that is used for the issuance of the certificate.
	// API type:
	//   - Group: pki.security.gdc.goog
	//   - Kind: CertificateIssuer
	// +optional
	IssuedBy *IssuerReference `json:"issuedBy,omitempty"`

	// BYOCertStatus specifies status options for byo-certificates mode.
	// +optional
	BYOCertStatus *BYOCertStatus `json:"byoCertStatus,omitempty"`

	// ErrorStatus contain a list of current errors and the timestamp this field gets updated.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

type BYOCertStatus struct {
	// Certificate Signing Request (CSR) status
	// +optional
	CSRStatus *CSRStatus `json:"csrStatus,omitempty"`

	// Externally signed certificate status
	// +optional
	SignedCertStatus *SignedCertStatus `json:"signedCertStatus,omitempty"`
}

type CSRStatus struct {
	// List of status conditions to indicate the status of a BYO Certificate CSR
	// - WaitingforSigning: Indicates that a new CSR has been generated to be signed by the customer.
	// - Ready: Indicates that the CSR has been signed
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Stores the CSR for the customer to sign.
	// +optional
	CSR []byte `json:"csr,omitempty"`
}

type SignedCertStatus struct {
	// List of status conditions to indicate the status of BYO certificate.
	// - Rejected: Indicates that the certificate does not match with the csr
	// - Ready: Indicates that the certificate is ready to use.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList represents a collection of certificates.
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&Certificate{},
		&CertificateList{},
	)
}
