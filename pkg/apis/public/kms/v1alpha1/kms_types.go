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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum:=AES_256_GCM
type AEADAlgorithm string

// +kubebuilder:validation:Enum:=EC_SIGN_P384_SHA384
type SigningAlgorithm string

// +kubebuilder:validation:Enum:=ECDH_P521_AES256
type KeySharingMechanism string

const (
	AEADAlgorithmAES256GCM            = "AES_256_GCM"
	SigningAlgorithmECDSAP384SHA384   = "EC_SIGN_P384_SHA384"
	KeySharingMechanismECDHP521AES256 = "ECDH_P521_AES256"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// Represents a cryptographic key to use for Authenticated
// +gdcloud:manifest:relevant=false,oc=kms
// Encryption with Associated Data (AEAD) operations.
type AEADKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AEADKeySpec   `json:"spec,omitempty"`
	Status            AEADKeyStatus `json:"status,omitempty"`
}

// Provides the specification for an AEADKey.
type AEADKeySpec struct {
	// +kubebuilder:default:=AES_256_GCM
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Algorithm is immutable"
	Algorithm AEADAlgorithm `json:"algorithm,omitempty"`
}

// Provides the status for an AEADKey.
type AEADKeyStatus struct {
	// A report that indicates when an AEADKey creation is complete and ready for use.
	Conditions           []metav1.Condition `json:"conditions,omitempty"`
	EncryptedKeyMaterial []byte             `json:"encryptedKeyMaterial,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a collection of AEADKeys.
type AEADKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AEADKey `json:"items"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +gdcloud:manifest:relevant=false,oc=kms
// Represents a cryptographic key to use for creating digital
// signatures.
type SigningKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SigningKeySpec   `json:"spec,omitempty"`
	Status            SigningKeyStatus `json:"status,omitempty"`
}

// Provides the specification for a SigningKey resource.
type SigningKeySpec struct {
	// +kubebuilder:default:=EC_SIGN_P384_SHA384
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Algorithm is immutable"
	Algorithm SigningAlgorithm `json:"algorithm,omitempty"`
}

// Provides the status for a SigningKey resource.
type SigningKeyStatus struct {
	// A report that indicates when a SigningKey creation is complete and ready for use.
	Conditions           []metav1.Condition `json:"conditions,omitempty"`
	EncryptedKeyMaterial []byte             `json:"encryptedKeyMaterial,omitempty"`
	PublicKey            []byte             `json:"publicKey,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a collection of SigningKey resources.
type SigningKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SigningKey `json:"items"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Key Kind",type="string",JSONPath=".status.importedKeyRef.kind"
// +gdcloud:manifest:relevant=false,oc=kms
// Represents a request to import a key.
type KeyImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyImportSpec   `json:"spec,omitempty"`
	Status            KeyImportStatus `json:"status,omitempty"`
}

// Provides the specification for a KeyImport resource.
type KeyImportSpec struct {
	// The information from the sender to unwrap the key material to import.
	Context KeySharingContext `json:"context"`
	// The wrapped key material to import.
	KeyToImport *WrappedKey `json:"keyToImport,omitempty"`
}

// Provides the status for a KeyImport resource.
type KeyImportStatus struct {
	// The status of the KeyImport resource as awaiting, successful, or failed and a reason for the failure.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// The information required to wrap the key to import.
	PeerContext PeerSharingContext `json:"peerContext,omitempty"`
	// A reference to the imported key.
	ImportedKeyRef corev1.TypedLocalObjectReference `json:"importedKeyRef,omitempty"`
}

// Contains information by the party that initiates intent for key import and export operations.
type KeySharingContext struct {
	// The algorithms to use to wrap keys.
	// +kubebuilder:default:=ECDH_P521_AES256
	Mechanism KeySharingMechanism `json:"mechanism"`
	PublicKey []byte              `json:"publicKey,omitempty"`
}

// Contains information by the party that responds to the intent for key import operations.
type PeerSharingContext struct {
	PublicKey  []byte `json:"publicKey,omitempty"`
	PrivateKey []byte `json:"privateKey,omitempty"`
}

// Contains the customer key wrapped for import or export operations.
type WrappedKey struct {
	// The attributes required to create the KMS key.
	Metadata KeyMetadata `json:"metadata"`
	// The wrapped key material.
	KeyMaterial []byte `json:"keyMaterial"`
}

// Represents the attributes required to create or re-create the customer key.
type KeyMetadata struct {
	// KeyTypeMeta has the kind and version information for the key.
	KeyTypeMeta metav1.TypeMeta `json:"keyTypeMeta"`
	// The algorithm to use with the key.
	Algorithm string `json:"algorithm,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a collection of key imports.
type KeyImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyImport `json:"items"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +gdcloud:manifest:relevant=false,oc=kms
// Represents a request to export a key.
type KeyExport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyExportSpec   `json:"spec,omitempty"`
	Status            KeyExportStatus `json:"status,omitempty"`
}

// Provides the specification for a KeyExport resource.
type KeyExportSpec struct {
	// The information necessary to wrap the key to export.
	Context KeySharingContext `json:"context"`
	// A reference to the key for export.
	KeyToExport corev1.TypedLocalObjectReference `json:"keyToExport"`
}

// Provides the status for a KeyExport resource.
type KeyExportStatus struct {
	// The status on the KeyExport resource as successful or failed, and provides a reason for the failure.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// The exported key material.
	ExportedKey *WrappedKey `json:"exportedKey,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a collection of KeyExport resources.
type KeyExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyExport `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Completed",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].reason"
// +gdcloud:manifest:relevant=false,oc=kms
// Represents a cluster level resource that runs root key rotation, and re-encryption of all KMS keys
// in the cluster.
type RotationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RotationJobSpec   `json:"spec,omitempty"`
	Status RotationJobStatus `json:"status,omitempty"`
}

// Provides the specification for a RotationJob resource.
type RotationJobSpec struct{}

// Provides the status for a RotationJob resource.
type RotationJobStatus struct {
	// The status of the RotationJob resource as successful or failed, and provides a reason for the failure.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of RotationJob resources.
type RotationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RotationJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&AEADKey{},
		&AEADKeyList{},
		&SigningKey{},
		&SigningKeyList{},
		&KeyImport{},
		&KeyExport{},
		&KeyImportList{},
		&KeyExportList{},
		&RotationJob{},
		&RotationJobList{},
	)
}
