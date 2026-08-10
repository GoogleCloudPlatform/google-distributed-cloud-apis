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

// +kubebuilder:validation:Enum:=AES_256_GCM
type AEADAlgorithm string

// +kubebuilder:validation:Enum:=EC_SIGN_P384_SHA384;RSA_SIGN_PKCS1_4096_SHA256;RSA_SIGN_PKCS1_4096_SHA384;RSA_SIGN_PKCS1_4096_SHA512
type SigningAlgorithm string

// +kubebuilder:validation:Enum:=ECDH_P521_AES256
type KeySharingMechanism string

const (
	AEADAlgorithmAES256GCM             = "AES_256_GCM"
	SigningAlgorithmECDSAP384SHA384    = "EC_SIGN_P384_SHA384"
	SigningAlgorithmRSAPKCS14096SHA256 = "RSA_SIGN_PKCS1_4096_SHA256"
	SigningAlgorithmRSAPKCS14096SHA384 = "RSA_SIGN_PKCS1_4096_SHA384"
	SigningAlgorithmRSAPKCS14096SHA512 = "RSA_SIGN_PKCS1_4096_SHA512"
	KeySharingMechanismECDHP521AES256  = "ECDH_P521_AES256"
)

// +kubebuilder:validation:Enum:=PASSTHROUGH_TO_CRYPTOBACKEND;WRAPPED_BY_CRYPTOBACKEND;WRAPPED_BY_ROOTKEY
type ProtectionMode string

const (
	ProtectionModePassthroughToCryptoBackend ProtectionMode = "PASSTHROUGH_TO_CRYPTOBACKEND"
	ProtectionModeWrappedByCryptoBackend     ProtectionMode = "WRAPPED_BY_CRYPTOBACKEND"
	ProtectionModeWrappedByRootKey           ProtectionMode = "WRAPPED_BY_ROOTKEY"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +gdcloud:manifest:relevant=true,oc=kms,component=kms,entities="keys"
// +gdcloud:manifest:verbs=create;describe;list;delete
// +gdcloud:manifest:rbac="create,describe,list,delete:kms-admin"
// +gdcloud:manifest:rbac="create,describe,list:kms-creator"
// +gdcloud:manifest:rbac="describe,list:kms-developer;kms-monitor;kms-viewer"

// Represents a cryptographic key to use for Authenticated
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
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Key material encrypted by the KMS root key.
	EncryptedKeyMaterial []byte `json:"encryptedKeyMaterial,omitempty"`
	// Identifier to the root key that wrapped the key material.
	// Follows the format `root-key-namespace/root-key-type/root-key-name/version`
	// Example - `kms-system/ctm/org1-root-key/1` or `kms-system/local/org-1-root-key/1`
	RootKeyID string `json:"rootKeyID,omitempty"`
}

func (a *AEADKey) Algorithm() AEADAlgorithm {
	return a.Spec.Algorithm
}

func (a *AEADKey) SetAlgorithm(algorithm AEADAlgorithm) {
	a.Spec.Algorithm = algorithm
}

func (a *AEADKey) Conditions() []metav1.Condition {
	return a.Status.Conditions
}

func (a *AEADKey) SetConditions(conditions []metav1.Condition) {
	a.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
// Represents a collection of AEADKeys.
type AEADKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AEADKey `json:"items"`
}

// +genclient
// +gdcloud:manifest:relevant=true,oc=kms,component=kms,entities="keys"
// +gdcloud:manifest:verbs=create;describe;list;delete
// +gdcloud:manifest:rbac="create,describe,list,delete:kms-admin"
// +gdcloud:manifest:rbac="create,describe,list:kms-creator"
// +gdcloud:manifest:rbac="describe,list:kms-developer;kms-monitor;kms-viewer"
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// Represents a cryptographic key to use for creating digital
// signatures.
type SigningKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SigningKeySpec   `json:"spec,omitempty"`
	Status            SigningKeyStatus `json:"status,omitempty"`
}

// Provides the specification for a SigningKey resource.
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.existingKeyID) || self.existingKeyID == oldSelf.existingKeyID",message="existingKeyID is immutable once set"
type SigningKeySpec struct {
	// Algorithm specifies the cryptographic signing algorithm to use for the key.
	// This field is immutable and cannot be changed after creation.
	// +kubebuilder:default:=EC_SIGN_P384_SHA384
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Algorithm is immutable"
	// +optional
	Algorithm SigningAlgorithm `json:"algorithm,omitempty"`

	// ProtectionLevel configures how and where the key material is managed and routed.
	// If this entire block is omitted, the KMS defaults to WRAPPED_BY_ROOTKEY ProtectionMode.
	//
	// Limited access: This field might not be available as it may not be
	// accredited for use in your deployment. You can access it when it's approved.
	//
	// +optional
	ProtectionLevel *ProtectionLevel `json:"protectionLevel,omitempty"`

	// ExistingKeyID specifies the identifier of a pre-existing HSM key to onboard.
	// The KMS uses this value to verify the key's existence during initial
	// creation, which supports use cases like Multi-Zone.
	// To bind to an existing key from another zone, populate this field using the
	// KeyIdentifier found in the Status of the original SigningKey.
	// If left empty, the KMS will generate a new key.
	// This field is immutable and cannot be changed in spec after the SigningKey is created.
	//
	// Limited access: This field might not be available as it may not be
	// accredited for use in your deployment. You can access it when it's approved.
	//
	// +kubebuilder:default=""
	// +optional
	ExistingKeyID string `json:"existingKeyID,omitempty"`
}

// ProtectionLevel bundles the cryptographic execution mode and its targeted routing backend.
// +kubebuilder:validation:XValidation:rule="self.mode == 'WRAPPED_BY_ROOTKEY' ? !has(self.cryptoBackendRef) : true",message="cryptoBackendRef must be omitted when mode is WRAPPED_BY_ROOTKEY"
type ProtectionLevel struct {
	// Mode dictates the boundary and mechanism of cryptographic operations for the key.
	// +kubebuilder:default:="WRAPPED_BY_ROOTKEY"
	// +optional
	Mode ProtectionMode `json:"mode,omitempty"`

	// CryptoBackendRef is a cross-namespace reference to the target CryptoBackend Custom Resource.
	// This field must be omitted when Mode is WRAPPED_BY_ROOTKEY.
	// For all other modes, if this field is omitted, the KMS will default to routing to a backend named "default".
	// +optional
	CryptoBackendRef *NamespacedReference `json:"cryptoBackendRef,omitempty"`
}

// Provides the status for a SigningKey resource.
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.keyIdentifier) || self.keyIdentifier == oldSelf.keyIdentifier",message="keyIdentifier is immutable once set and cannot be removed"
type SigningKeyStatus struct {
	// Conditions represent the latest available observations of the SigningKey's state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// EncryptedKeyMaterial is the key material encrypted by the KMS root key.
	// +optional
	EncryptedKeyMaterial []byte `json:"encryptedKeyMaterial,omitempty"`
	// PublicKey is the public key of the asymmetric signing key pair.
	// +optional
	PublicKey []byte `json:"publicKey,omitempty"`
	// RootKeyID is the identifier of the root key that wrapped the private key material.
	// Follows the format `root-key-namespace/root-key-type/root-key-name/version`
	// Example - `kms-system/ctm/org1-root-key/1` or `kms-system/local/org-1-root-key/1`
	// +optional
	RootKeyID string `json:"rootKeyID,omitempty"`
	// KeyIdentifier is the primary identifier for private and public key objects.
	// This is the serialized goog.gdc.kms.v1.KeyIdentifier protobuf message schema.
	// The field is immutable once set and cannot be removed.
	//
	// Limited access: This field might not be available as it may not be
	// accredited for use in your deployment. You can access it when it's approved.
	//
	// +optional
	KeyIdentifier string `json:"keyIdentifier,omitempty"`
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
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Key Kind",type="string",JSONPath=".status.importedKeyRef.kind"
// +gdcloud:manifest:relevant=true,oc=kms,component=kms,entities="key-imports"
// +gdcloud:manifest:verbs=list;delete;describe
// +gdcloud:manifest:rbac="list,describe:kms-viewer;kms-admin;kms-keyimport-admin"
// +gdcloud:manifest:rbac="delete:kms-keyimport-admin"
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
	// Identifier to the root key that wrapped the key material.
	// Follows the format `root-key-namespace/root-key-type/root-key-name/version`
	// Example - `kms-system/ctm/org1-root-key/1` or `kms-system/local/org-1-root-key/1`
	RootKeyID string `json:"rootKeyID,omitempty"`
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
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +gdcloud:manifest:relevant=true,oc=kms,component=kms,entities="key-exports"
// +gdcloud:manifest:verbs=create;delete;describe;list
// +gdcloud:manifest:rbac="create,delete,describe,list:kms-keyexport-admin"
// +gdcloud:manifest:rbac="describe,list:kms-viewer;kms-admin;kms-monitor"
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
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Completed",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Completed')].reason"
// +gdcloud:manifest:relevant=true,oc=kms,component=kms
// +gdcloud:manifest:entities="rotation-jobs"
// +gdcloud:manifest:verbs=create;describe;list
// +gdcloud:manifest:rbac="create,describe,list:kms-rotationjob-admin"
// +gdcloud:manifest:rbac="describe,list:kms-rotationjob-monitor;kms-org-rotationjob-monitor"
// Represents a cluster level resource that runs root key rotation, and re-encryption of all KMS keys
// in the cluster.
type RotationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RotationJobSpec   `json:"spec,omitempty"`
	Status RotationJobStatus `json:"status,omitempty"`
}

// Provides the specification for a RotationJob resource.
type RotationJobSpec struct {
	// The root key name specified in the form namespaces/<namespace>/secrets/<rootkeyname>
	RootKeyResourceName string `json:"rootKeyResourceName,omitempty"`
	// TTLSecondsAfterCompletion specifies how long a RotationJob should persist after it's completed.
	// If this field is not set, it will default to 24h (86400s)
	// +kubebuilder:default=86400
	// +kubebuilder:validation:Minimum=0
	TTLSecondsAfterCompletion *int32 `json:"ttlSecondsAfterCompletion,omitempty"`
}

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
