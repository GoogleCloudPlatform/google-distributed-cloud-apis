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

// Defines the desired state of the `DeleteClusterBackupRequest` resource.
type DeleteClusterBackupRequestSpec struct {
	// The name of the `ClusterBackup` resource to be deleted.
	// Note, the `ClusterBackup` resource must exist in the same namespace as the `DeleteClusterBackupRequest` resource.
	// +kubebuilder:validation:Required
	ClusterBackupRef string `json:"clusterBackupRef"`
}

// Defines the observed state of `DeleteClusterBackupRequest` resource.
type DeleteClusterBackupRequestStatus struct {
	// The time the resource expires.
	TimeToExpire metav1.Time `json:"timeToExpire"`

	// The status of the observed state of `DeleteClusterBackupRequest` resource.
	StatusFields StatusFields `json:"statusField"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-backups",skipcodegen=true
// +gdcloud:manifest:verbs=delete
// +gdcloud:manifest:rbac="create:organization-cluster-backup-admin"
// Defines the schema for the `DeleteClusterBackupRequest` API.
type DeleteClusterBackupRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeleteClusterBackupRequestSpec   `json:"spec,omitempty"`
	Status DeleteClusterBackupRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of `DeleteBackupRequest` resources.
type DeleteClusterBackupRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeleteClusterBackupRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeleteClusterBackupRequest{}, &DeleteClusterBackupRequestList{})
}
