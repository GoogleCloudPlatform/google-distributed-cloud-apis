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
	cmacme "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="IsDefault",type="string",JSONPath=".metadata.labels['pki\\.security\\.gdc\\.goog/is-default-issuer']"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +gdcloud:manifest:relevant=false,oc=platauth
// CertificateIssuer represents an issuer for Certificate as a Service.
// You can mark a CertificateIssuer as the default issuer by adding/setting
// the label `pki.security.gdc.goog/is-default-issuer: true`.
type CertificateIssuer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateIssuerSpec   `json:"spec,omitempty"`
	Status CertificateIssuerStatus `json:"status,omitempty"`
}

type CertificateIssuerSpec struct {
	// BYOCertConfig configures this issuer in BYO-Cert mode.
	// +optional
	BYOCertConfig *BYOCertIssuerConfig `json:"byoCertConfig,omitempty"`

	// CAaaSConfig configures this issuer to sign
	// certificates using CA deployed by the CertificateAuthority API.
	// +optional
	CAaaSConfig *CAaaSIssuerConfig `json:"caaasConfig,omitempty"`

	// ACMEConfig configures this issuer to sign certificates using ACME server.
	// +optional
	ACMEConfig *ACMEIssuerConfig `json:"acmeConfig,omitempty"`
}

// BYOCertIssuerConfig defines an issuer based on the BYO-Cert model.
type BYOCertIssuerConfig struct {
	// FallbackCertificateAuthority is the reference to a default CAaaS operated CA.
	// API type:
	//   - Group: pki.security.gdc.goog
	//   - Kind: CertificateAuthority
	FallbackCertificateAuthority *CAReference `json:"fallbackCertificateAuthority"`
}

// CAaaSIssuerConfig defines an issuer that requests certificates from a CA created using the CAaaS service.
type CAaaSIssuerConfig struct {
	// A reference to a CertificationAuthority which will sign the certificate.
	// API type:
	//   - Group: pki.security.gdc.goog
	//   - Kind: CertificateAuthority
	CertificateAuthorityRef *CAReference `json:"certificateAuthorityRef"`
}

type ACMEIssuerConfig struct {
	// This contains the Root CA data of certificates issued by ACME server.
	RootCACertificate []byte `json:"rootCACertificate"`

	// ACME configures this issuer to communicate with a RFC 8555 (ACME) server
	// to obtain signed certificates.
	// ACME is an acme.cert-manager.io/v1 ACMEIssuer.
	ACME *cmacme.ACMEIssuer `json:"acme"`
}

type CertificateIssuerStatus struct {
	// Stores the root CA used by the current certificate issuer.
	CA []byte `json:"ca,omitempty"`

	// List of status conditions to indicate the status of the CertificateIssuer.
	// - Ready: Indicates that the CertificateIssuer is ready to use.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// IssuerReference represents an Issuer Reference. It has information
// to retrieve an issuer in any namespace.
type IssuerReference struct {
	// Name is unique within a namespace to reference an issuer resource.
	Name string `json:"name"`
	// Namespace defines the space within which the issuer name must be unique.
	Namespace string `json:"namespace"`
}

// +kubebuilder:object:root=true

// CertificateIssuerList represents a collection of certiifcate issuers.
type CertificateIssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateIssuer `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&CertificateIssuer{},
		&CertificateIssuerList{},
	)
}
