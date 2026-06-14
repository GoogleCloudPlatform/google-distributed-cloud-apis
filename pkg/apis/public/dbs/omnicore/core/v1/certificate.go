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
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DefaultPrivateKeySize = map[PrivateKeyAlgorithm]int{
	RSAKeyAlgorithm:     4096,
	ECDSAKeyAlgorithm:   384,
	Ed25519KeyAlgorithm: 0,
}

// +kubebuilder:validation:Enum=RSA;ECDSA;Ed25519
type PrivateKeyAlgorithm string

const (
	// RSA private key algorithm.
	RSAKeyAlgorithm PrivateKeyAlgorithm = "RSA"

	// ECDSA private key algorithm.
	ECDSAKeyAlgorithm PrivateKeyAlgorithm = "ECDSA"

	// Ed25519 private key algorithm.
	Ed25519KeyAlgorithm PrivateKeyAlgorithm = "Ed25519"
)

// Denotes how private keys should be generated or sourced when a Certificate
// is being issued.
// +kubebuilder:validation:Enum=Never;Always
type PrivateKeyRotationPolicy string

const (
	// RotationPolicyNever means a private key will only be generated if one
	// does not already exist in the target `spec.secretName`.
	RotationPolicyNever PrivateKeyRotationPolicy = "Never"

	// RotationPolicyAlways means a private key matching the specified
	// requirements will be generated whenever a re-issuance occurs.
	RotationPolicyAlways PrivateKeyRotationPolicy = "Always"
)

// CertificateRef contains a reference to a certificate secret and a key for the
// CA certificate.
type CertificateRef struct {
	// SecretRef is a reference to the secret that contains the database server
	// certificate.
	// +optional
	SecretRef corev1.SecretReference `json:"secretRef,omitempty"`
	// CertificateKey is the key within the secret that corresponds to the
	// CA certificate, which client connections can use to the database securely.
	// +optional
	CertificateKey string `json:"certificateKey,omitempty"`
}

// IssuerReference is a reference to a certificate issuer with a given name,
// namespace, kind, and group.
type IssuerReference struct {
	// Name of the issuer being referred to.
	Name string `json:"name"`
	// Namespace of the issuer being referred to.
	// +optional
	// nullon(samwise-fleet,samwise-local)
	Namespace string `json:"namespace,omitempty"`
	// Kind of the cert-manager.io issuer being referred to.
	// +kubebuilder:validation:Enum:=Issuer;ClusterIssuer
	// +kubebuilder:default:=Issuer
	// +optional
	// nullon(dbs-fleet,dbs-local)
	Kind string `json:"kind,omitempty"`
}

// +kubebuilder:object:generate=true

// CertificateRequest contains attributes common to both data plane and control
// plane agent certificates. These attributes are propagated to cert-manager
// Certificate resources created by the controller. Updating these fields after
// initial creation will result in certificate rotation.
type CertificateRequest struct {
	// Duration is the total time duration for which the certificate is valid.
	// RenewBefore and Duration must both be specified together or omitted.
	// +optional
	Duration *metav1.Duration `json:"duration,omitempty"`
	// RenewBefore is the time prior to expiry when cert-manager will begin to try
	// re-issuance. RenewBefore and Duration must both be specified together or
	// omitted.
	// +optional
	RenewBefore *metav1.Duration `json:"renewBefore,omitempty"`
	// PrivateKey contains additional configuration for the private key
	// associated with the certificates.
	// +optional
	PrivateKey *CertificatePrivateKey `json:"privateKey,omitempty"`
}

// +kubebuilder:object:generate=true

// DataPlaneCertificateRequest contains attributes for data plane certificates.
// This includes certificates for the database server, pgBackRest server, and
// pgBouncer auth query client. Updating these fields after initial creation
// will result in certificate rotation.
type DataPlaneCertificateRequest struct {
	CertificateRequest `json:",inline"`

	// DNSNames contains the list of DNS subject alternative names that should be
	// included in the database server's certificate. These names will not be
	// included in the pgBackRest or pgBouncer auth query client certificates.
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`

	// CommonName is the common name that should be used for the database server's
	// certificate. This common name will not be included in the pgBackRest or
	// pgBouncer auth query client certificates. If omitted, the common name will
	// only be populated on the certificate if `fillCommonName` is true, in which
	// case the value will match the first SAN on the certificate.
	// +optional
	CommonName string `json:"commonName,omitempty"`
}

// +kubebuilder:object:generate=true

// CertificatePrivateKey contains configuration options for private keys
// configured by cert-manager.
type CertificatePrivateKey struct {
	// RotationPolicy controls how private keys should be regenerated when a
	// re-issuance is being processed.
	// If set to `Never`, cert-manager will only generate a new private key if
	// one does not already exist.
	// If set to `Always`, cert-manager will generate a new private key matching
	// the specified requirements whenever a re-issuance occurs.
	// +optional
	RotationPolicy PrivateKeyRotationPolicy `json:"rotationPolicy,omitempty"`

	// Algorithm is the private key algorithm of the corresponding private key
	// for this certificate.
	//
	// If provided, allowed values are either `RSA`, `ECDSA` or `Ed25519`.
	// If `algorithm` is specified and `size` is not provided,
	// key size of 4096 will be used for `RSA` key algorithm and
	// key size of 384 will be used for `ECDSA` key algorithm.
	// +optional
	Algorithm PrivateKeyAlgorithm `json:"algorithm,omitempty"`

	// Size is the key bit size of the corresponding private key that is
	// propagated to cert-manager Certificate resources.
	//
	// If `algorithm` is set to `RSA`, valid values are `4096` or `8192`,
	// and will default to `4096` if not specified.
	// If `algorithm` is set to `ECDSA`, valid values are `384` or `521`,
	// and will default to `384` if not specified.
	// If `algorithm` is set to `Ed25519`, Size is ignored.
	// No other values are allowed.
	// +optional
	Size int `json:"size,omitempty"`
}

// ValidateKeySize checks that a given key size is an allowed value for the
// given private key algorithm.
func ValidateKeySize(algorithm PrivateKeyAlgorithm, size int) error {
	errMsg := "private key size for algorithm %s must be one of the following: %v"
	switch algorithm {
	case RSAKeyAlgorithm:
		if size != 4096 && size != 8192 {
			return fmt.Errorf(errMsg, RSAKeyAlgorithm, []int{4096, 8192})
		}
	case ECDSAKeyAlgorithm:
		if size != 384 && size != 521 {
			return fmt.Errorf(errMsg, ECDSAKeyAlgorithm, []int{384, 521})
		}
	case Ed25519KeyAlgorithm:
		break
	default:
		return fmt.Errorf("private key algorithm must be one of the following: %v", []PrivateKeyAlgorithm{RSAKeyAlgorithm, ECDSAKeyAlgorithm, Ed25519KeyAlgorithm})
	}
	return nil
}
