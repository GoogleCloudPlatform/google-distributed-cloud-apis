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

// Defines the desired state of a `ManualClusterRestoreRequest` resource.
type ManualClusterRestoreRequestSpec struct {
	// The name of the cluster restore to be created.
	// +kubebuilder:validation:Required
	ClusterRestoreName string `json:"clusterRestoreName"`

	// The name of the cluster restore plan to pull the `ClusterRestoreConfig` resource from.
	// +kubebuilder:validation:Required
	ClusterRestorePlanName string `json:"clusterRestorePlanName"`

	// The name of the cluster backup that is being restored.
	// Note, it is expected that the cluster restore plan and cluster backup exist in the same
	// namespace as the `ManualClusterRestoreRequest` resource.
	// +kubebuilder:validation:Required
	ClusterBackupName string `json:"clusterBackupName"`

	// A user-specified descriptive string for the cluster restore created by this `ManualClusterRestoreRequest` resource.
	// +optional
	Description string `json:"description"`

	// Filter which can be used to further refine the resource selection of the cluster restore beyond the coarse-grained scope defined in the ClusterRestorePlan.
	// +optional
	Filter *Filter `json:"filter,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-restores"
// +gdcloud:manifest:verbs=create
// +gdcloud:manifest:rbac="create:organization-cluster-backup-admin"
// +gdcloud:manifest:skipcodegen=true
// Defines the schema for the `ManualClusterRestoreRequest` API.
type ManualClusterRestoreRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManualClusterRestoreRequestSpec `json:"spec,omitempty"`
	Status ManualRestoreRequestStatus      `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a list of `ManualClusterRestoreRequest` resources.
type ManualClusterRestoreRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManualClusterRestoreRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManualClusterRestoreRequest{}, &ManualClusterRestoreRequestList{})
}
