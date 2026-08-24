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
	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Represents an API that wraps around the Restore custom resource.
// Defines the desired state of a ClusterRestore.
type ClusterRestoreSpec struct {
	// The cluster where data will be restored.
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf)
	// +kubebuilder:validation:Required
	TargetCluster TargetCluster `json:"targetCluster"`

	// The name of the cluster backup, which must be in the same namespace as the cluster restore.
	// +kubebuilder:validation:Required
	ClusterBackupName string `json:"clusterBackupName"`

	// The name of the cluster restore plan from which this cluster restore inherited its `ClusterRestoreConfig` resource.
	// +optional
	ClusterRestorePlanName string `json:"clusterRestorePlanName"`

	// The configuration of the cluster restore.
	// +kubebuilder:validation:Required
	ClusterRestoreConfig ClusterRestoreConfig `json:"clusterRestoreConfig"`

	// An optional description of the cluster restore. This has no impact on functionality.
	// +optional
	Description string `json:"description,omitempty"`

	// Filter can be used to further refine the resource selection of the cluster restore beyond the coarse-grained scope defined in the `ClusterRestorePlan`.
	// +optional
	Filter *Filter `json:"filter,omitempty"`
}

// Defines the observed state of a cluster restore.
type ClusterRestoreStatus struct {
	// Specifies the status of the cluster restore. Supported conditions include `JobCreated`, `Succeeded`.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The current state of the cluster restore.
	// +optional
	State RestoreState `json:"state,omitempty"`

	// A human-readable description of why the cluster restore is in the current state.
	// +optional
	StateReason string `json:"stateReason,omitempty"`

	// The most recent errors with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// The number of resources restored in this cluster restore action.
	// +optional
	ResourcesRestoredCount int64 `json:"resourcesRestoredCount,omitempty"`

	// The number of resources excluded in this cluster restore action.
	// +optional
	ResourcesExcludedCount int64 `json:"resourcesExcludedCount,omitempty"`

	// The number of resources that failed to be restored in this cluster restore action.
	// +optional
	ResourcesFailedCount int64 `json:"resourcesFailedCount,omitempty"`

	// The number of volumes restored in this cluster restore action.
	// +optional
	RestoredVolumesCount int64 `json:"restoredVolumesCount,omitempty"`

	// The create time of the cluster restore process.
	// +optional
	CreateTime *metav1.Time `json:"startTime,omitempty"`

	// The end time of the cluster restore process.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-restores"
// +gdcloud:manifest:verbs=create;delete;describe;list
// +gdcloud:manifest:rbac="create,delete,describe,list:organization-cluster-backup-admin"
// +gdcloud:manifest:skipcodegen=true
// Defines the schema for the `ClusterRestore` API.
type ClusterRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRestoreSpec   `json:"spec,omitempty"`
	Status ClusterRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a list of cluster restores.
type ClusterRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterRestore{}, &ClusterRestoreList{})
}
