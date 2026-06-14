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

// Defines the desired state of a cluster volume backup.
type ClusterVolumeBackupSpec struct {
	// The name of the cluster that this cluster volume backup maintains data for.
	// +kubebuilder:validation:Required
	TargetCluster TargetCluster `json:"cluster"`

	// The name of the associated backup. Each cluster volume backup must be associated with a backup. It will always
	// be in the same namespace.
	// +kubebuilder:validation:Required
	ClusterBackupName string `json:"clusterBackupName"`
	// The name of the backup plan from which the asssociated backup was created. It will always
	// be in the same namespace.
	// +kubebuilder:validation:Required
	ClusterBackupPlanName string `json:"clusterBackupPlanName"`
	// The source persistent volume claim from which the cluster volume backup is taken from.
	// +kubebuilder:validation:Required
	SourcePvc corev1.TypedObjectReference `json:"sourcePVC"`
}

// Defines the observed state of a cluster volume backup.
type ClusterVolumeBackupStatus struct {
	// An underlying cluster volume backup handle, which uniquely
	// identifies a cluster volume backup inside of a backup repository. This handle
	// doesn't have a unified format and is treated as an opaque string.
	// +optional
	VolumeBackupHandle string `json:"volumeBackupHandle,omitempty"`
	// A cluster volume backup format. For example, `PD`, `Portable`, etc.
	// +optional
	Format VolumeBackupFormat `json:"format,omitempty"`
	// The size of the cluster volume backup in the backup storage. For incremental
	// backups this value may dynamically change if one of the previous volume
	// backups was deleted.
	// +optional
	StorageBytes int64 `json:"storageBytes,omitempty"`
	// The minimum size of the disk to which this volume backup can be restored.
	// +optional
	DiskSizeBytes int64 `json:"diskSizeBytes,omitempty"`
	// The current state of the volume backup.
	// +optional
	State VolumeBackupState `json:"state,omitempty"`
	// A human-readable message indicating details about why the backup is in this state.
	// +optional
	StateMessage string `json:"message,omitempty"`
	// The timestamp when this `ClusterVolumeBackup` resource was completed
	// in the text format of [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt).
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
	// Specifies the status of the cluster volume backup.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-volume-backups"
// +gdcloud:manifest:verbs=describe;list
// +gdcloud:manifest:rbac="describe,list:organization-cluster-backup-admin"
// +gdcloud:manifest:skipcodegen=true
// +genclient
// Defines the schema for the `ClusterVolumeBackup` API.
type ClusterVolumeBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterVolumeBackupSpec   `json:"spec,omitempty"`
	Status ClusterVolumeBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `ClusterVolumeBackup` resources.
type ClusterVolumeBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterVolumeBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterVolumeBackup{}, &ClusterVolumeBackupList{})
}
