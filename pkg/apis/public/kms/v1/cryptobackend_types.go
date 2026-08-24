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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CryptoBackendType defines the cryptographic provider type (e.g., "ThalesLunaHSM").
// +kubebuilder:validation:Enum:=ThalesLunaHSM
type CryptoBackendType string

const (
	CryptoBackendTypeThalesLunaHSM CryptoBackendType = "ThalesLunaHSM"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.type"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"

// CryptoBackend represents a logical cryptographic provider within the platform.
// It acts as the bridge between high-level cryptographic requests and the
// underlying physical or virtual security boundaries (e.g., Thales Luna HSMs).
// +gdcloud:manifest:relevant=false,oc=kms
type CryptoBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CryptoBackendSpec   `json:"spec,omitempty"`
	Status CryptoBackendStatus `json:"status,omitempty"`
}

// CryptoBackendSpec defines the desired state of CryptoBackend.
// +kubebuilder:validation:XValidation:rule="self.type == oldSelf.type",message="type is immutable and cannot be changed after creation"
// +kubebuilder:validation:XValidation:rule="has(self.thalesLunaHSM) == (self.type == 'ThalesLunaHSM')",message="thalesLunaHSM configuration must be specified if and only if type is ThalesLunaHSM"
type CryptoBackendSpec struct {
	// Type defines the cryptographic provider type (e.g., "ThalesLunaHSM").
	// +kubebuilder:validation:Required
	Type CryptoBackendType `json:"type"`

	// ThalesLunaHSM contains configuration specific to Thales Luna HSMs.
	// +optional
	ThalesLunaHSM *ThalesLunaHSMConfig `json:"thalesLunaHSM,omitempty"`
}

// ThalesLunaHSMConfig defines the connection and pooling parameters for physical Thales Luna HSM.
type ThalesLunaHSMConfig struct {
	// The Luna Network OrgHSM for the CryptoBackend.
	// These resources are expected to be in the Organization's Control Plane API.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	OrgHSMRefs []NamespacedReference `json:"orgHSMRefs"`

	// ClientName is the identifier for this KMS client. It must exactly match
	// one of the client names defined in the upstream OrgHSM specification.
	// Format:
	//   <zone-name>-<org-name>-kms-<suffix>
	//   (e.g., us-east1-org1-kms-tenant-pki)
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^[a-z0-9-]+-[a-z0-9-]+-kms-[a-z0-9-]+$"
	ClientName string `json:"clientName"`

	// Credentials required to interact with the Luna HSM partitions.
	// +kubebuilder:validation:Required
	Credentials LunaCredentials `json:"credentials"`
}

// LunaCredentials holds references to the HSM passwords.
type LunaCredentials struct {
	// PSOPasswordRef is the reference to the secret holding the PSO (Partition Security Officer) password.
	// The password must be stored in the Secret's data map under the key "password".
	// +kubebuilder:validation:Required
	PSOPasswordRef corev1.SecretReference `json:"psoPasswordRef"`

	// COPasswordRef is the reference to the secret holding the CO (Crypto Officer) password.
	// The password must be stored in the Secret's data map under the key "password".
	// +kubebuilder:validation:Required
	COPasswordRef corev1.SecretReference `json:"coPasswordRef"`

	// CloningDomainRef is the reference to the secret holding the Cloning Domain.
	// The cloning domain must be stored in the Secret's data map under the key "domain".
	// +kubebuilder:validation:Required
	CloningDomainRef corev1.SecretReference `json:"cloningDomainRef"`
}

// CryptoBackendStatus defines the observed state and discovered endpoints of the backend.
type CryptoBackendStatus struct {
	// Conditions represent the latest available observations of the backend's state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Endpoint is the cluster-local URL (e.g., "dns://kms-backend-1.platform.svc.cluster.local:8443")
	// exposed by the KMS controller for this specific backend.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// ThalesLunaHSM contains status details specific to Thales Luna HSMs.
	// +optional
	ThalesLunaHSM *ThalesLunaHSMStatus `json:"thalesLunaHSM,omitempty"`
}

// ThalesLunaHSMStatus contains the operational state of the Thales HA Group.
type ThalesLunaHSMStatus struct {
	// HAGroupID is the logical token ID generated by the Thales client library.
	// +optional
	HAGroupID string `json:"haGroupID,omitempty"`

	// ActiveMembers represents the physical partitions successfully pooled.
	// +optional
	ActiveMembers []LunaHAGroupMember `json:"activeMembers,omitempty"`

	// ConfigMapRef points to the dynamically generated ConfigMap containing
	// 'Chrystoki.conf' file contents.
	// +optional
	ConfigMapRef *NamespacedReference `json:"configMapRef,omitempty"`
}

// LunaHAGroupMember represents a single partition's health within the HA Group.
type LunaHAGroupMember struct {
	// Serial is the physical partition serial number.
	// +kubebuilder:validation:Required
	Serial string `json:"serial"`
}

// +kubebuilder:object:root=true

// CryptoBackendList contains a list of CryptoBackend resources.
// Used by the API server to return paginated lists of backends.
type CryptoBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CryptoBackend `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&CryptoBackend{},
		&CryptoBackendList{},
	)
}
