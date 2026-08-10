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
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +gdcloud:manifest:relevant=true,oc=platauth,component=pki,entities="revoke-certificate-requests"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,describe,list,update,delete:certificate-authority-service-operation-manager,certificate-authority-service-admin"
// +gdcloud:manifest:rbac="describe,list:certificate-authority-service-viewer"
// +gdcloud:manifest:skipcodegen=true
// RevokeCertificateRequest defines a request to revoke a certificate.
type RevokeCertificateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RevokeCertificateRequestSpec   `json:"spec"`
	Status RevokeCertificateRequestStatus `json:"status,omitempty"`
}

// RevokeCertificateRequestSpec defines a request for the revocation of a certificate.
type RevokeCertificateRequestSpec struct {
	// CertificateRequestRef references a GDC managed CertificateRequest object containing the certificate to revoke.
	CertificateRequestRef CertificateRequestReference `json:"certificateRequestRef"`

	// Reason for revocation.
	// keyCompromise;caCompromise,certificateHold,affiliationChanged;
	// superseded;cessationOfOperation;privilegeWithdrawn;aACompromise
	Reason RevocationReason `json:"reason"`
}

type CertificateRequestReference struct {
	// Name of the certificate.
	Name string `json:"name"`

	// Namespace of the certificate.
	Namespace string `json:"namespace"`
}

type RevokeCertificateRequestStatus struct {
	// List of status conditions to indicate the status of a certificate to be revoked.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type RevocationReason string

const (
	// Unspecified reason
	Unspecified RevocationReason = "unspecified"

	// Key material for a certificate may have leaked.
	KeyCompromise RevocationReason = "keycompromise"

	// The key material for a certificate authority in the issuing path may have
	// leaked.
	CertificteAuthorityCompromise RevocationReason = "cacompromise"

	// The subject or other attributes in a certificate have changed.
	AffiliationChanged RevocationReason = "affiliationchanged"

	// Certificate has been superseded.
	Superseded RevocationReason = "superseded"

	// This certificate or entities in the issuing path have ceased to
	// operate.
	CessationOfOperation RevocationReason = "cessationofoperation"

	// Certificate should not be considered valid, it is expected that it
	// may become valid in the future.
	CertificateHold RevocationReason = "certificatehold"

	// Certificate no longer has permission to assert the listed
	// attributes.
	PrivilegeWithdrawn RevocationReason = "privilegewithdrawn"

	// The authority which determines appropriate attributes for a certificate
	// may have been compromised.
	AttributeAuthorityCompromise RevocationReason = "aacompromise"
)

// +kubebuilder:object:root=true

// RevokeCertificateRequestList represents a collection of revoke certificate requests.
type RevokeCertificateRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RevokeCertificateRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&RevokeCertificateRequest{},
		&RevokeCertificateRequestList{},
	)
}
