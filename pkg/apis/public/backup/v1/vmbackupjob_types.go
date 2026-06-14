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
	"k8s.io/apimachinery/pkg/types"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// Defines the desired state of the backup.
type VMBackupJobSpec struct {
	// The control plane backup name that the vm backup job is related to.
	// The name is immutable.
	// +kubebuilder:validation:Required
	BackupRef *corev1alpha1.NamespacedName `json:"backupRef,omitempty"`

	// The reference of the VM this CR is backing up
	// This reference could be nil if the backup is a disk backup
	// +optional
	VMRef *corev1alpha1.NamespacedName `json:"vmRef,omitempty"`

	// List of VM Disks this CR is backing up
	// +kubebuilder:validation:Required
	VMDiskRefs []*VMDisk `json:"vmDiskRefs,omitempty"`

	// map of volume backup name to VolumeBackup struct under this backup
	// This data is embedded into VMBackupJob and not present in the cluster
	// +optional
	VolumeBackupRefs map[string]*VolumeBackup `json:"volumeBackupRefs,omitempty"`

	// An optional string description of the Backup. This field has no impact on functionality.
	// +optional
	Description string `json:"description,omitempty"`

	// Specifies whether the backup resource was
	// created manually. If `True`, this backup has been created manually,
	// If `False`, this Backup has been created automatically from the
	// backup plan schedule.
	// +optional
	// +kubebuilder:default:=false
	Manual bool `json:"manual,omitempty"`
	// The number of days from the `create_time` of this
	// backup for which deletion is blocked.
	// For backups created automatically from a backup schedule, this field is set to
	// the value of `BackupPlan.RetentionPolicy.backup_delete_block_days`.
	// For backups created manually that leave this field unspecified, the service
	// assigns the value of `BackupPlan.RetentionPolicy.backup_delete_block_days`.
	// If a backup is created where the value of this field is less than the value of
	// `BackupPlan.RetentionPolicy.backup_delete_block_days`, an invalid
	// response is returned from the service.
	// This field must be an integer value between `0-90`.
	// This field must only be increased by an update request, or an invalid
	// response is returned by the service.
	// Note, this field only applies to backups with a `Succeeded` state.
	// Default to `0`.
	// +optional
	// +kubebuilder:default:=0
	DeleteLockDays uint `json:"deleteLockDays,omitempty"`
	// The number of days to keep
	// this backup for, after which it is automatically deleted.
	// This is calculated from the `create_time` of the backup.
	// If this field is not specified or set to `0`, it means the backup is not automatically
	// deleted.
	// For backups automatically created from a a backup schedule, this field is assigned
	// the value of `BackupPlan.RetentionPolicy.backup_retain_days`.
	// For backups created manually that leave this field unspecified, the service
	// assigns the value of `BackupPlan.RetentionPolicy.backup_retain_days`.
	// If a backup is created where the value of this field is less than
	// the value of `delete_lock_days`, an invalid response is returned by the service.
	// This field must only be increased by an update request, or an invalid
	// response is returned by the service.
	// Default to `0`.
	// +optional
	// +kubebuilder:default:=0
	RetainDays uint `json:"retainDays,omitempty"`
}

type VMDisk struct {
	// Namespace where the VM Disk resides in
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// Name of the VM Disk
	// +optional
	Name string `json:"name"`
	// Name of the underlying PVC
	// +optional
	PvcName string `json:"pvcName"`
	// UID of the PVC
	// this field is used to get this particular VMBackupJob given the UID of the PVC
	// +optional
	PvcUID types.UID `json:"uid"`
	// The total size, measured in bytes.
	// +optional
	SizeBytes int64 `json:"sizeBytes,omitempty"`
}

// The various states a backup can be in.
// +kubebuilder:validation:Enum=Initialized;InProgress;InTransfer;Succeeded;Failed;Deleting;DeleteFailed;Retry
type VMBackupJobState string

const (
	VMBackupJobStateInitialized   VMBackupJobState = "Initialized"
	VMBackupJobStateInProgress    VMBackupJobState = "InProgress"
	VMBackupJobStateInTransfer    VMBackupJobState = "InTransfer"
	VMBackupJobStateSucceeded     VMBackupJobState = "Succeeded"
	VMBackupJobStateFailed        VMBackupJobState = "Failed"
	VMBackupJobStateDeleting      VMBackupJobState = "Deleting"
	VMBackupJobStateDeletedFailed VMBackupJobState = "DeleteFailed"
	VMBackupJobStateRetry         VMBackupJobState = "Retry"
)

const (
	VMBackupJobConditionSucceeded = "Succeeded"
)

// Defines the observed state of a backup.
type VMBackupJobStatus struct {
	// The current state of the backup.
	// +optional
	State VMBackupJobState `json:"state,omitempty" reflect:"unexport"`
	// A human-readable description of why the backup is in the current state.
	// +optional
	StateReason string `json:"reason,omitempty"`
	// The most recent errors with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
	// Conditions represents the observations of this VMBackupJob's current state.
	// Known condition types: Succeeded
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// The timestamp of when this VMBackupJob resource was created. This can
	// be converted to and from [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt).
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`
	// The timestamp of when all disks in VMBackupJob resource are snapshotted. This can
	// be converted to and from [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt).
	// +optional
	SnapshotTime *metav1.Time `json:"snapshotTime,omitempty"`
	// The timestamp of when this VMBackupJob was last updated.
	// This can be converted to and from
	// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt).
	// +optional
	UpdateTime *metav1.Time `json:"updateTime,omitempty"`
	// The total number of volumes backed up.
	// +optional
	VolumeCount int64 `json:"volumeCount,omitempty"`

	// total number of retries required to backup this VM/VM disk
	// +optional
	RetryCount int64 `json:"retryCount,omitempty"`

	// The time when the deletion lock
	// will expire. This is an output only field calculated from the combined value of `create_time` and
	// `delete_lock_days`, and is updated accordingly when the
	// `delete_lock_days` field of a backup is updated.
	// Note, this field only applies to backups with a `Succeeded` state.
	// +optional
	DeleteLockExpireTime *metav1.Time `json:"deleteLockExpireTime,omitempty"`
	// The time when the backup is
	// automatically deleted. It's an output only field calculated from the combined value
	// of `create_time` and `retain_days`, and is updated accordingly when the
	// `retain_days` field of a backup has been updated.
	// +optional
	RetainExpireTime *metav1.Time `json:"retainExpireTime,omitempty"`
	// The completion time of the backup.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="CreateTime",type="string",JSONPath=".status.createTime"
// +gdcloud:manifest:relevant=false,oc=back
// +genclient
// Defines the schema for the `Backup` API.
type VMBackupJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VMBackupJobSpec   `json:"spec,omitempty"`
	Status VMBackupJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of backups.
type VMBackupJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMBackupJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMBackupJob{}, &VMBackupJobList{})
}
