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

// Defines the desired state of a cluster volume restore.
type ClusterVolumeRestoreSpec struct {
	// The name of the cluster where volume will be restored.
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf)
	// +kubebuilder:validation:Required
	TargetCluster TargetCluster `json:"targetCluster"`

	// The name of the cluster restore resource that created this cluster volume restore.
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf)
	// +kubebuilder:validation:Required
	ClusterRestoreName string `json:"clusterRestoreName"`

	// The name of the cluster volume backup resource that we are restoring.
	// +kubebuilder:validation:Required
	ClusterVolumeBackupName string `json:"clusterVolumeBackupName"`

	// The target `PersistentVolumeClaim` resource to be restored.
	// +kubebuilder:validation:Required
	TargetPvc NamespacedName `json:"targetPvc"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// Represents an API that wraps around the VolumeRestore custom resource.
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-volume-restores"
// +gdcloud:manifest:verbs=describe;list
// +gdcloud:manifest:rbac="describe,list:organization-cluster-backup-admin"
// +gdcloud:manifest:skipcodegen=true
// Defines the schema for the `ClusterVolumeRestore` API.
type ClusterVolumeRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterVolumeRestoreSpec `json:"spec,omitempty"`
	Status VolumeRestoreStatus      `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a list of `ClusterVolumeRestore` resources.
type ClusterVolumeRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterVolumeRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterVolumeRestore{}, &ClusterVolumeRestoreList{})
}
