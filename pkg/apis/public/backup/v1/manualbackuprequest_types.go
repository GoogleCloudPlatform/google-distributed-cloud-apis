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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Defines the desired state of a `ManualBackupRequest` resource.
type ManualBackupRequestSpec struct {
	// The name of the manual backup to be created.
	// +kubebuilder:validation:Required
	BackupName string `json:"backupName"`

	// The name of the backup plan from which the `BackupConfig` resource is pulled.
	// +kubebuilder:validation:Required
	BackupPlanName string `json:"backupPlanName"`

	// Note, it is expected that the backup plan and associated backup exist in the same
	// namespace as the `ManualBackupRequest` resource.

	// A user-specified descriptive string for the backup created by this `ManualBackupRequest` resource.
	// +optional
	Description string `json:"description"`

	// The number of days from the `create_time` of this backup for which deletion is blocked.
	// For backups automatically created from a schedule, this field is given the value of
	// `BackupPlan.RetentionPolicy.backup_delete_block_days`. If a backup is created with this field
	// unspecified, it is given the value of
	// `BackupPlan.RetentionPolicy.backup_delete_block_days`. If this Backup is created with this field set
	// to a value less than the value of `BackupPlan.RetentionPolicy.backup_delete_block_days`, an
	// invalid response is returned from the agent.
	// This field must be a value within `0-90`.
	// This field must only be increased by an update request, or an invalid
	// response is returned by the agent.
	// Note that this field only applies to backups with a `Succeeded` state.
	// Default to `0`.
	// +optional
	// +kubebuilder:default:=0
	DeleteLockDays uint `json:"deleteLockDays"`

	// The number of days to keep this backup for, after which it is
	// automatically deleted. If this field is not specified or set to `0`, it means the backup is not
	// automatically deleted. For backups automatically created from a backup schedule, this field is assigned
	// the value of `BackupPlan.RetentionPolicy.backup_default_retain_days`. For created backups that leave this
	// field unspecified, the agent to uses the value of
	// `BackupPlan.RetentionPolicy.backup_default_retain_days`. The creation of a backup with this field set
	// to a value less than `delete_lock_days` results in an invalid response from the agent. This
	// field must only be increased in an update request, or an invalid response is returned by
	// the agent immediately.
	// Default to `0`.
	// +optional
	// +kubebuilder:default:=0
	RetainDays uint `json:"retainDays"`
}

// Defines the observed state of a `ManualBackupRequest` resource.
type ManualBackupRequestStatus struct {
	// The time the resource should expire.
	TimeToExpire metav1.Time `json:"timeToExpire"`

	// The status of the observed state of `ManualBackupRequest` resource.
	StatusField StatusFields `json:"statusField"`

	// The statusMessage is any error message for `ManualBackupRequest` resource.
	// +optional
	StatusMessage string `json:"statusMessage"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +gdcloud:manifest:relevant=false,oc=back
// +genclient
// Defines the schema for the `ManualBackupRequest` API.
type ManualBackupRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManualBackupRequestSpec   `json:"spec,omitempty"`
	Status ManualBackupRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of `ManualBackupRequest` resources.
type ManualBackupRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManualBackupRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManualBackupRequest{}, &ManualBackupRequestList{})
}
