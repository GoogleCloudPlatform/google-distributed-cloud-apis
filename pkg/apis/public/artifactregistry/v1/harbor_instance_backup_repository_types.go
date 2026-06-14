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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".spec.endpoint"
// +kubebuilder:printcolumn:name="Region",type="string",JSONPath=".spec.region"
// +kubebuilder:printcolumn:name="Bucket",type="string",JSONPath=".spec.bucket"
// +gdcloud:manifest:relevant=true,oc=haas,component=harbor,entities="backup-repositories"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:harbor-instance-admin"
// Represents an instance of a backup repository for Harbor instance.
type HarborInstanceBackupRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the backup repository.
	// +kubebuilder:validation:Required
	Spec HarborInstanceBackupRepositorySpec `json:"spec"`
	// The most recently observed status of the backup repository.
	Status HarborInstanceBackupRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of backup repositories for Harbor instance.
type HarborInstanceBackupRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HarborInstanceBackupRepository `json:"items"`
}

// Defines the desired state of the Harbor backup repository.
type HarborInstanceBackupRepositorySpec struct {
	// A reference to an Access Secret to access s3 bucket.
	// The secret should contain 2 data from the S3 access granting flow:
	// - access-key-id
	// - access-key
	// +kubebuilder:validation:Required
	SecretReference corev1.SecretReference `json:"secretReference"`

	// The endpoint used to access the Harbor backup repository. In the case of Google Private Cloud, this is the S3 endpoint that provides access to the Tenant project.
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint"`

	// The region of a given endpoint for the bucket.
	// +optional
	Region string `json:"region"`

	// The bucket within the endpoint to upload backups to.
	// +kubebuilder:validation:Required
	Bucket string `json:"bucket"`

	// A user-specified descriptive string for this backup repository.
	// +optional
	Description string `json:"description,omitempty"`
}

// The various states a backup repository can be in.
// +kubebuilder:validation:Enum=Unspecified;NotReady;Ready
type BackupRepositoryState string

const (
	BackupRepositoryStateUnspecified BackupRepositoryState = "Unspecified"
	BackupRepositoryStateReady       BackupRepositoryState = "Ready"
	BackupRepositoryStateNotReady    BackupRepositoryState = "NotReady"
)

// Defines the observed state of a Harbor backup repository.
type HarborInstanceBackupRepositoryStatus struct {
	// Conditions:
	// - Ready: readiness of the backup repository, any error when reconciling embedded object will be surfaced here.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// The current state of the backup repository.
	// +optional
	State BackupRepositoryState `json:"state,omitempty"`
	// A human-readable description of why the backup repository is in the current state.
	// +optional
	StateReason string `json:"reason,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&HarborInstanceBackupRepository{},
		&HarborInstanceBackupRepositoryList{},
	)
}
