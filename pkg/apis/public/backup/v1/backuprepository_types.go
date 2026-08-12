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

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// Defines the options for calling to an S3 endpoint for the backup repository.
type S3Options struct {
	// The bucket within the endpoint to upload backups to.
	// +kubebuilder:validation:Required
	Bucket string `json:"bucket" reflect:"unexport"`

	// The region of a given endpoint. Formatting is storage system dependent.
	// +optional
	Region string `json:"region" reflect:"unexport"`

	// Specifies whether to force path style URLs for objects. For example, `https://s3.amazonaws.com//` instead of `https://.s3.amazonaws.com/`.
	// +optional
	ForcePathStyle bool `json:"forcePathStyle,omitempty" reflect:"unexport"`

	// The server-side encryption algorithm used when storing objects. For example, `AES256 or `aws:kms`.
	// +optional
	ServerSideEncryption string `json:"serverSideEncryption,omitempty" reflect:"unexport"`

	// The KMS key ID to use for object storage if server-side encryption is using KMS.
	// +optional
	KmsKeyID string `json:"kmsKeyId,omitempty" reflect:"unexport"`
}

// Defines the desired state of the backup repository.
type BackupRepositorySpec struct {
	// A reference to an Access Secret which is dependent on your storage system of choice.
	// This Secret is used requests to this endpoint. For example, an S3 Access Secret.
	// +kubebuilder:validation:Required
	SecretReference corev1.SecretReference `json:"secretReference" reflect:"unexport"`

	// The endpoint used to access the backup repository. In the case of Google Private Cloud, this is the S3 endpoint that provides access to the Tenant project.
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint" reflect:"unexport"`

	// The type of the backup repository. For example, S3 or Google Cloud Storage. This tells the agent which storage system or API to use.
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=Unspecified
	Type RepositoryType `json:"type" reflect:"unexport"`

	// Only one of the below fields should be set, corresponding to the type of backup repository that
	// this CR represents.

	// The data used for configuring access to an S3-compatible `BackupRepo` resource.
	// +optional
	S3Options *S3Options `json:"s3Options,omitempty" reflect:"unexport"`

	// The policy that determines whether this backup repository is read-only or read-write.
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=ReadWrite
	ImportPolicy ImportPolicy `json:"importPolicy" reflect:"unexport"`

	// This field determines whether to substitute the backup repository name of imported resources with this backup repository name.
	// +optional
	// +deprecated
	ImportSubstitution *ImportSubstitution `json:"importSubstitution,omitempty" reflect:"unexport"`

	// Specifies the action that a read-write backup repository takes if the
	// storage bucket that it is initialized with has already been claimed by a different
	// backup repository. If `True`, the new backup repository still claims ownership
	// of the storage bucket by replacing the existing sentinel file with its own sentinel file.
	// If `False`, the creation of the new backup repository fails with an error. The
	// default value is `False`. This must only be used if the sentinel file that is
	// overridden no longer has a backup repository, otherwise that backup repository
	// enters an error state which might cause undesired side effects.
	// +kubebuilder:default:=false
	// +optional
	Force bool `json:"force" reflect:"unexport"`
}

// Defines the observed state of a backup repository.
type BackupRepositoryStatus struct {
	// A field that connects a backup repository to the sentinel file that it owns.
	// +optional
	SentinelEtag string `json:"sentinelEtag" reflect:"unexport"`
	// The errors that have occurred during the most recent
	// reconciliation attempt for the backup repository.
	// +kubebuilder:validation:Required
	// +kubebuilder:default:="NoError"
	ReconciliationError ReconciliationError `json:"reconciliationError" reflect:"unexport"`
	// The error messages that might have occurred during reconciliation.
	// +optional
	ReconciliationErrorMessage string `json:"reconciliationErrorMessage" reflect:"unexport"`
	// The most recent errors from reconciliation with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
	// Specifies whether this backup repository has completed the initial import or not.
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=false
	InitialImportDone bool `json:"initialImportDone" reflect:"unexport"`
	// Conditions represents the observations of this backup repository's current state.
	// Known condition types: Ready.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" reflect:"unexport"`
}

// The errors that have occurred during the most recent reconciliation attempt for the backup repository.
// +kubebuilder:validation:Enum=NoError;InternalError;SentinelCreateError
type ReconciliationError string

const (
	NoError             ReconciliationError = "NoError"
	InternalError       ReconciliationError = "InternalError"
	SentinelCreateError ReconciliationError = "SentinelCreateError"
)

const (
	BackupRepositoryConditionReady     = "Ready"
	BackupRepositoryConditionOwnership = "Ownership"
)

const (
	ReasonReady                     = "Ready"
	ReasonNotReady                  = "NotReady"
	ReasonOwnershipVerified         = "OwnershipVerified"
	ReasonOwnershipClaimed          = "OwnershipClaimed"
	ReasonOwnershipLost             = "OwnershipLost"
	ReasonOwnershipClaimFailed      = "OwnershipClaimFailed"
	ReasonOwnershipUnknown          = "OwnershipUnknown"
	ReasonOwnershipRelinquished     = "OwnershipRelinquished"
	ReasonOwnershipRelinquishFailed = "OwnershipRelinquishFailed"
	ReasonOwnershipNotClaimed       = "OwnershipNotClaimed"
)

// The policy that determines whether this backup repository is read-only or read-write.
// +kubebuilder:validation:Enum=ReadOnly;ReadWrite;
type ImportPolicy string

const (
	ReadOnly  ImportPolicy = "ReadOnly"
	ReadWrite ImportPolicy = "ReadWrite"
)

// The type of endpoint being written to.
// +kubebuilder:validation:Enum=Unspecified;S3;
type RepositoryType string

const (
	// An unspecified type. This is the default value.
	RepositoryTypeUnspecified RepositoryType = "Unspecified"
	// An S3 endpoint following the S3 API. For example,
	// https://docs.aws.amazon.com/AmazonS3/latest/userguide/MakingRequests.html.
	RepositoryTypeS3 RepositoryType = "S3"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.type"
// +kubebuilder:printcolumn:name="Policy",type="string",JSONPath=".spec.importPolicy"
// +kubebuilder:printcolumn:name="Error",type="string",JSONPath=".status.reconciliationError"
// +gdcloud:manifest:relevant=false,oc=back
// +genclient
// +genclient:nonNamespaced
// Defines the schema for the `BackupRepository` API.
type BackupRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupRepositorySpec   `json:"spec,omitempty"`
	Status BackupRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of backup repository resources.
type BackupRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupRepository `json:"items"`
}

// Represents whether we substitute the backup repository name on imported resources.
type ImportSubstitution struct {
	SubstituteBRName bool `json:"substituteBRName,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BackupRepository{}, &BackupRepositoryList{})
}

// Custom error BackupRepositoryNotFoundError is returned when a backup repository is not found.
type BackupRepositoryNotFoundError struct {
	Name        string
	ErrorStatus metav1.Status
}

func (e BackupRepositoryNotFoundError) Error() string {
	return fmt.Sprintf("backup repository not found: %s ", e.Name)
}

func (e *BackupRepositoryNotFoundError) Unwrap() error {
	return &errors.StatusError{
		ErrStatus: e.ErrorStatus,
	}
}
