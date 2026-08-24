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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// Defines the desired state of the `ClusterBackupRepositorySpec` resource.
type ClusterBackupRepositorySpec struct {
	// A reference to an Access Secret which is dependent on your storage system of choice.
	// This Secret is used requests to this endpoint. For example, an S3 Access Secret.
	// +kubebuilder:validation:Required
	SecretReference corev1.SecretReference `json:"secretReference"`

	// The endpoint used to access the cluster backup repository.
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`

	// The type of the cluster backup repository. For example, S3 or Google Cloud Storage. This tells the agent which storage system or API to use.
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=Unspecified
	Type RepositoryType `json:"type"`

	// Only one of the below fields should be set, corresponding to the type of backup repository that
	// this CR represents.

	// The data used for configuring access to an S3-compatible `BackupRepo` resource.
	// +optional
	S3Options *S3Options `json:"s3Options,omitempty"`

	// The policy that determines whether this backup repository is read-only or read-write.
	// +kubebuilder:validation:Required
	// +kubebuilder:default:=ReadWrite
	ImportPolicy ImportPolicy `json:"importPolicy"`

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
	Force bool `json:"force"`
}

// Defines the observed state of `ClusterBackupRepository` resource.
type ClusterBackupRepositoryStatus struct {
	// A field that connects a backup repository to the sentinel file that it owns.
	// +optional
	SentinelEtag string `json:"sentinelEtag"`
	// The errors that have occurred during the most recent
	// reconciliation attempt for the backup repository.
	// +kubebuilder:validation:Required
	// +kubebuilder:default:="NoError"
	ReconciliationError ReconciliationError `json:"reconciliationError"`
	// The error messages that might have occurred during reconciliation.
	// +optional
	ReconciliationErrorMessage string `json:"reconciliationErrorMessage"`
	// The most recent errors from reconciliation with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// Specifies the status of the cluster backup repository. Supported conditions include `InitialImportDone`, `Ready`, `Deprecated`.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
//+kubebuilder:resource:path=clusterbackuprepositories,singular=clusterbackuprepository
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.type"
// +kubebuilder:printcolumn:name="Policy",type="string",JSONPath=".spec.importPolicy"
// +kubebuilder:printcolumn:name="Error",type="string",JSONPath=".status.reconciliationError"

// +genclient
// +genclient:nonNamespaced
// +gdcloud:manifest:relevant=false,oc=back
// Defines the schema for the `ClusterBackupRepository` API.
type ClusterBackupRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterBackupRepositorySpec   `json:"spec,omitempty"`
	Status ClusterBackupRepositoryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of `ClusterBackupRepository` resources.
type ClusterBackupRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterBackupRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterBackupRepository{}, &ClusterBackupRepositoryList{})
}
