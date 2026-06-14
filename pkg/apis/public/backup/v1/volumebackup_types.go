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

// TODO(b/200285234): This file is heavily patterned off of:
//   cs//depot/google3/blaze-out/genfiles/google/cloud/gkebackup/v1/volume_backup.pb.go
// Going forward, we should come up with a good process to keep these
// relatively in sync.

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// Defines the desired state of a volume backup.
type VolumeBackupSpec struct {
	// The name of the cluster that this volume backup maintains data for.
	// By default, it is the cluster in which this volume backup is created.
	// +optional
	ClusterName string `json:"clusterName,omitempty" reflect:"unexport"`

	// The name of the associated backup. Each volume backup must be associated with a backup.
	// +kubebuilder:validation:Required
	BackupName string `json:"backupName" reflect:"unexport"`
	// The name of the backup plan from which the asssociated backup was created.
	// +kubebuilder:validation:Required
	BackupPlanName string `json:"backupPlanName" reflect:"unexport"`
	// The source persistent volume claim from which the volume backup is taken from.
	// +kubebuilder:validation:Required
	SourcePvc NamespacedName `json:"sourcePVC" reflect:"unexport"`
	// An underlying volume backup handle which uniquely
	// identifies a volume backup inside of a volume backup repository. This handle
	// doesn't have a unified format and is treated as an opaque string.
	// DEPRECATED: Use the status field instead.
	// +optional
	VolumeBackupHandle string `json:"volumeBackupHandle,omitempty" reflect:"unexport"`
}

// An enum that specifies the format of the volume backup.
// +kubebuilder:validation:Enum=Unspecified;SNAPSHOT_ONLY;PORTABLE;BACKUP_FORMAT_2;
type VolumeBackupFormat string

// TODO(b/200285313): Add VolumeBackup Formats as we onboard Volume Backup methods.
const (
	// Default value, not specified.
	VolumeBackupFormatUnspecified VolumeBackupFormat = "Unspecified"

	// Value for snapshot only backups
	VolumeBackupFormatSnapshotOnly VolumeBackupFormat = "SNAPSHOT_ONLY"

	// Value for portable volume backup
	VolumeBackupFormatPortable VolumeBackupFormat = "PORTABLE"

	// Value for NetApp SM-C (native volume backup)
	VolumeBackupFormatBackupFormat2 VolumeBackupFormat = "BACKUP_FORMAT_2"
)

// An enum that suggests the current state of a volume backup.
// +kubebuilder:validation:Enum=Unspecified;Creating;Snapshotting;Uploading;Succeeded;Failed;Deleting;
type VolumeBackupState string

const (
	// Default value, not specified
	VolumeBackupStateUnspecified VolumeBackupState = "Unspecified"
	// A volume for the backup was identified and backup process is about to
	// start.
	VolumeBackupStateCreating VolumeBackupState = "Creating"
	// A volume is being snapshotted.
	VolumeBackupStateSnapshotting VolumeBackupState = "Snapshotting"
	// A volume snapshot is being uploaded.
	VolumeBackupStateUploading VolumeBackupState = "Uploading"
	// A volume backup ready.
	VolumeBackupStateSucceeded VolumeBackupState = "Succeeded"
	// A volume backup failed.
	VolumeBackupStateFailed VolumeBackupState = "Failed"
	// A volume backup is being deleted.
	VolumeBackupStateDeleting VolumeBackupState = "Deleting"
)

const (
	VolumeBackupConditionSucceeded = "Succeeded"
)

// Defines the observed state of a volume backup.
type VolumeBackupStatus struct {
	// An underlying volume backup handle, which uniquely
	// identifies a volume backup inside of a volume backup repository. This handle
	// doesn't have a unified format and is treated as an opaque string.
	// +optional
	VolumeBackupHandle string `json:"volumeBackupHandle,omitempty" reflect:"unexport"`
	// A volume backup format. For example, `PD`, `Portable`, etc.
	// +optional
	Format VolumeBackupFormat `json:"format,omitempty" reflect:"unexport"`
	// The size of the volume backup in the backup storage. For incremental
	// backups this value may dynamically change if one of the previous volume
	// backups was deleted.
	// +optional
	StorageBytes int64 `json:"storageBytes,omitempty" reflect:"unexport"`
	// The minimum size of the disk to which this volume backup can be restored.
	// +optional
	DiskSizeBytes int64 `json:"diskSizeBytes,omitempty" reflect:"unexport"`
	// The current state of the volume backup.
	// +optional
	State VolumeBackupState `json:"state,omitempty" reflect:"unexport"`
	// A human-readable message indicating details about why the backup is in this state.
	// +optional
	StateMessage string `json:"message,omitempty"`
	// The timestamp when this `VolumeBackup` resource was completed
	// in the text format of [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt).
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty" reflect:"unexport"`
	// The physical upload duration.
	// +optional
	UploadDuration *metav1.Duration `json:"uploadDuration,omitempty" reflect:"unexport"`
	// The most recent errors with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
	// Conditions represents the observations of this VolumeBackup's current state.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +gdcloud:manifest:relevant=false,oc=back
// +genclient
// Defines the schema for the `VolumeBackup` API.
type VolumeBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeBackupSpec   `json:"spec,omitempty"`
	Status VolumeBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `VolumeBackup` resources.
type VolumeBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeBackup{}, &VolumeBackupList{})
}
