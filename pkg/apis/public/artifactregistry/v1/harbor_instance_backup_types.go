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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="CreateTime",type="string",JSONPath=".status.createTime"
// +kubebuilder:printcolumn:name="CompleteTime",type="string",JSONPath=".status.completeTime"
// +kubebuilder:printcolumn:name="ExpireTime",type="string",JSONPath=".status.retainExpireTime"
// +kubebuilder:printcolumn:name="BackupPlan",type="string",JSONPath=".spec.backupPlanName"
// +gdcloud:manifest:relevant=true,oc=haas,component=harbor,entities="backups"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:harbor-instance-admin"
// +gdcloud:manifest:rbac="describe,list:harbor-instance-viewer"
// Defines the schema for the `Backup` API for HarborInstance.
type HarborInstanceBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the backup.
	// +kubebuilder:validation:Required
	Spec HarborInstanceBackupSpec `json:"spec"`
	// The most recently observed status of the backup.
	Status HarborInstanceBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of backup for Harbor instance.
type HarborInstanceBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HarborInstanceBackup `json:"items"`
}

// Defines the desired state of the backup.
type HarborInstanceBackupSpec struct {
	// The name of the backup plan from which this backup was created.
	BackupPlanName string `json:"backupPlanName,omitempty"`
	// Configuration for a backup
	// This field is expected:
	//   1. to be set by the user for manual customized backup explicitly.
	//   2. to be unset with backupPlanName filled for automated scheduled backup or manual backup.
	//      The Backup reconciler will get the config from the Backup Plan.
	// +kubebuilder:validation:Optional
	BackupConfig *BackupConfig `json:"backupConfig,omitempty"`

	// An optional string description of the Backup. This field has no impact on functionality.
	Description string `json:"description,omitempty"`

	// Specifies whether the backup resource was
	// created manually. If True, this backup has been created manually,
	// If False, this backup has been created automatically from the
	// backup plan schedule.
	// +optional
	// +kubebuilder:default:=false
	Manual bool `json:"manual,omitempty"`
	// The number of days to keep
	// this backup for, after which it is automatically deleted.
	// This is calculated from the create_time of the backup.
	// If this field is not specified or set to 0, it means the backup is not automatically
	// deleted.
	// For backups automatically created from a a backup schedule, this field is assigned
	// the value of BackupPlan.RetentionPolicy.backup_retain_days.
	// For backups created manually that leave this field unspecified, the service
	// assigns the value of BackupPlan.RetentionPolicy.backup_retain_days.
	// Default to 0.
	// +optional
	// +kubebuilder:default:=0
	RetainDays uint `json:"retainDays,omitempty"`
}

// The various states a backup can be in.
// +kubebuilder:validation:Enum=Unspecified;Creating;InProgress;Succeeded;Failed;Deleting;DeleteFailed
type BackupState string

// TODO(b/200285235): Should this be merged with BackupJob Phase?
const (
	BackupStateUnspecified   BackupState = "Unspecified"
	BackupStateCreating      BackupState = "Creating"
	BackupStateInProgress    BackupState = "InProgress"
	BackupStateSucceeded     BackupState = "Succeeded"
	BackupStateFailed        BackupState = "Failed"
	BackupStateDeleting      BackupState = "Deleting"
	BackupStateDeletedFailed BackupState = "DeleteFailed"
)

// Defines the observed state of a backup.
type HarborInstanceBackupStatus struct {
	// TODO(b/353364160): consider also adding standard conditions, and use Kubernetes Job as a reference: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#job-v1-batch
	// The current state of the backup.
	// +optional
	State BackupState `json:"state,omitempty"`
	// A human-readable description of why the backup is in the current state.
	// +optional
	StateReason string `json:"reason,omitempty"`
	// The timestamp of when this backup resource was created. This can
	// be converted to and from [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt).
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`
	// The completion time of the backup.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
	// The time when the backup is
	// automatically deleted. It's an output only field calculated from the combined value
	// of create_time and retain_days, and is updated accordingly when the
	// retain_days field of a backup has been updated.
	// +optional
	RetainExpireTime *metav1.Time `json:"retainExpireTime,omitempty"`
	// The total size for backup measured in bytes.
	// +optional
	TotalSizeBytes int64 `json:"totalSizeBytes,omitempty"`
	// The total size for registry backup handled by data transfer service, measured in bytes.
	// +optional
	RegistrySizeBytes int64 `json:"registrySizeBytes,omitempty"`
	// The total size for database backup handled by Dbcluster export, measured in bytes.
	// +optional
	DatabaseSizeBytes int64 `json:"databaseSizeBytes,omitempty"`

	// Subdirectory appended to the Database export location. This will be used when running
	// import during restoration.
	// +optional
	DatabaseExportSubDirectory *string `json:"exportSubDirectory,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&HarborInstanceBackup{},
		&HarborInstanceBackupList{},
	)
}
