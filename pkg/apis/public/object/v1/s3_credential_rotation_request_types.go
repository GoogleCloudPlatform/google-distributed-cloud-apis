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

// S3CredentialRotationRequestConditionType represents a condition type for S3CredentialRotationRequest.
type S3CredentialRotationRequestConditionType string

const (
	// ConditionTypeReady indicates the reconciler is ready and processing the request.
	// True: Actively processing or has finished.
	// False: An error has occurred, check other conditions and reasons.
	// Unknown: Initial state.
	ConditionTypeReady S3CredentialRotationRequestConditionType = "Ready"

	// ConditionTypeRotationCompleted indicates whether the new key has been generated
	// and the local Secret has been updated.
	// True: Rotation phase is complete.
	// False: Rotation phase is not complete or has failed.
	ConditionTypeRotationCompleted S3CredentialRotationRequestConditionType = "RotationCompleted"

	// ConditionTypeRevocationCompleted indicates whether the old key has been permanently revoked
	// from StorageGRID and removed from the Secret.
	// True: Revocation is complete.
	// False: Revocation is not complete or has failed.
	ConditionTypeRevocationCompleted S3CredentialRotationRequestConditionType = "RevocationCompleted"

	// ConditionTypeSucceeded is the primary condition, true when all phases (Rotation and Revocation) are complete.
	ConditionTypeSucceeded S3CredentialRotationRequestConditionType = "Succeeded"
)

// Constants for metav1.Condition.Reason
const (
	// ReasonReconciling indicates that the resource is being reconciled.
	ReasonReconciling string = "Reconciling"
	// ReasonRotationSucceeded indicates that the key rotation phase completed successfully.
	ReasonRotationSucceeded string = "RotationSucceeded"
	// ReasonRevocationSucceeded indicates that the key revocation phase completed successfully.
	ReasonRevocationSucceeded string = "RevocationSucceeded"

	// ReasonSucceeded indicates the entire process is successfully complete. Used with ConditionTypeSucceeded.
	ReasonSucceeded string = "Succeeded"
	// ReasonRotationInProgress indicates that the rotation phase is currently in progress.
	ReasonRotationInProgress string = "RotationInProgress"
	// ReasonRevocationPending indicates that the rotation phase is complete, and the system is waiting for the OldKeyRevocationTimestamp to start revocation.
	ReasonRevocationPending string = "RevocationPending"
	// ReasonRevocationInProgress indicates that the revocation phase is actively being processed.
	ReasonRevocationInProgress string = "RevocationInProgress"

	// ReasonSecretSubjectInvalid indicates that the subject annotation on the secret is missing or invalid.
	ReasonSecretSubjectInvalid string = "SecretSubjectInvalid"
	// ReasonSecretSubjectTypeInvalid indicates that the subject type label on the secret is missing or invalid.
	ReasonSecretSubjectTypeInvalid string = "SecretSubjectTypeInvalid"
	// ReasonTargetSecretPatchFailed indicates a failure to patch the target Secret resource in Kubernetes.
	ReasonTargetSecretPatchFailed string = "TargetSecretPatchFailed"

	// Failure/Rejection Reasons
	ReasonTargetSecretNotFound      string = "TargetSecretNotFound"
	ReasonSecretTypeInvalid         string = "SecretTypeInvalid"
	ReasonStorageGRIDError          string = "StorageGRIDError"
	ReasonGlobalSecretUpdateFailed  string = "GlobalSecretUpdateFailed"
	ReasonInternalError             string = "InternalError"
	ReasonKeyRevocationFailed       string = "KeyRevocationFailed"
	ReasonValidationFailed          string = "ValidationFailed" // e.g., issues with spec
	ReasonRotationAlreadyInProgress string = "RotationAlreadyInProgress"
)

// +gdcloud:manifest:relevant=false,oc=obj
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=s3credentialrotationrequests,singular=s3credentialrotationrequest,scope=Namespaced,shortName=s3rr
// +kubebuilder:printcolumn:name="SecretName",type="string",JSONPath=".spec.secretRef.name"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type=='RotationCompleted')].status"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// S3CredentialRotationRequest is the Schema for the S3CredentialRotationRequests API
// +genclient
// +kubebuilder:storageversion
type S3CredentialRotationRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3CredentialRotationRequestSpec   `json:"spec,omitempty"`
	Status S3CredentialRotationRequestStatus `json:"status,omitempty"`
}

// S3CredentialRotationRequestSpec defines the desired state of a S3CredentialRotationRequest.
// This spec is used to initiate the rotation of S3 credentials stored in a Kubernetes Secret.
type S3CredentialRotationRequestSpec struct {
	// SECURITY: The backend controller MUST perform an authorization check to ensure
	// the entity creating this request has the right to access and modify the
	// referenced Secret.

	// SecretName is deprecated and will be removed in a future release.
	// It is superseded by SecretRef and should not be used.
	// +optional
	// +deprecated
	SecretName string `json:"secretName,omitempty"`

	// SecretRef references the Kubernetes Secret containing the S3 credentials
	// that need to be rotated.
	// +optional
	SecretRef corev1.SecretReference `json:"secretRef,omitempty"`

	// ExpirationDelay specifies the duration for which the old, rotated S3 key
	// will remain valid after the new key has been created. This allows for a graceful
	// transition period for clients and applications to switch to the new credentials.
	// If not specified, this value is defaulted by a defaulting webhook to 72 hours.
	// A value of "0s" indicates that the old key should be revoked immediately
	// after the new key is activated.
	// The format is a duration string that is parsable by Go's time.ParseDuration function.
	// For example, "72h", "90m", "30s". Note that units such as "d" for days are not supported.
	// See https://pkg.go.dev/time#ParseDuration for more details on the format.
	// +optional
	// +kubebuilder:default:="72h"
	ExpirationDelay *metav1.Duration `json:"expirationDelay,omitempty"`

	// RequestTTL defines the time-to-live (TTL) for this S3CredentialRotationRequest
	// object itself. The TTL timer begins after the credential rotation is complete
	// (i.e., when the `RotationCompleted` status condition becomes `True`).
	// After this duration, the request object will be automatically garbage
	// collected by the system. This mechanism preserves the request as an audit
	// record for a configurable period.
	// If not specified, this value is defaulted by a defaulting webhook to 90 days.
	// The format is a duration string that is parsable by Go's time.ParseDuration function.
	// For example, "2160h" (for 90 days), "72h", "30m". Note that units such as "d" for days are not supported.
	// See https://pkg.go.dev/time#ParseDuration for more details on the format.
	// +optional
	// +kubebuilder:default:="2160h"
	RequestTTL *metav1.Duration `json:"requestTTL,omitempty"`
}

// S3CredentialRotationRequestStatus defines the observed state of a S3CredentialRotationRequest.
// It provides feedback on the progress and outcome of the credential rotation process.
type S3CredentialRotationRequestStatus struct {
	// Conditions represents the latest available observations of the rotation
	// request's state. It provides a detailed breakdown of the rotation process,
	// including its current phase and any errors that may have occurred.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// OldKeyRevocationTimestamp is the timestamp when the old, rotated S3 key is
	// scheduled to be permanently revoked from the backend storage system.
	// This time is calculated by adding the ExpirationDelay to the LastTransitionTime
	// of the RotationCompleted condition.
	// +optional
	OldKeyRevocationTimestamp *metav1.Time `json:"oldKeyRevocationTimestamp,omitempty"`
}

// +kubebuilder:object:root=true
// S3CredentialRotationRequestList contains a list of S3CredentialRotationRequest
type S3CredentialRotationRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3CredentialRotationRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&S3CredentialRotationRequest{}, &S3CredentialRotationRequestList{})
}
