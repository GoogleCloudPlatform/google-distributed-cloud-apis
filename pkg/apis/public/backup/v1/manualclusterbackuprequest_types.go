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

// Defines the desired state of a `ManualClusterBackupRequest` resource.
type ManualClusterBackupRequestSpec struct {
	// The name of the cluster backup to be created. It is created
	// inside of the same namespace as the `ManualClusterBackupRequest` resource.
	// +kubebuilder:validation:Required
	ClusterBackupName string `json:"clusterBackupName"`

	// The name of the cluster backup plan from which the `ClusterBackupConfig` resource is pulled.
	// +kubebuilder:validation:Required
	//
	// Note, it is expected that the cluster backup plan exists in the same
	// namespace as the `ManualClusterBackupRequest` resource.
	ClusterBackupPlanRef string `json:"clusterBackupPlanRef"`

	// A user-specified descriptive string for the cluster backup created by this `ManualClusterBackupRequest` resource.
	// +optional
	Description string `json:"description"`

	// The number of days from the `create_time` of this backup for which deletion is blocked.
	// For backups automatically created from a schedule, this field is given the value of
	// `BackupPlan.RetentionPolicy.backup_delete_block_days`. If a `Backup` is created with this field
	// unspecified, it is given the value of
	// `BackupPlan.RetentionPolicy.backup_delete_block_days`. If this `Backup` is created with this field set
	// to a value less than the value of `ClusterBackupPlan.RetentionPolicy.backup_delete_block_days`, an
	// invalid response is returned from the agent.
	// This field must be a value within `0-90`.
	// This field must only be increased by an update request, or an invalid
	// response is returned by the agent.
	// Note that this field only applies to backups with a `Succeeded` state.
	// +optional
	DeleteLockDays *uint `json:"deleteLockDays,omitempty"`

	// The number of days to keep this backup for, after which it is
	// automatically deleted. If this field is not specified or set to `0`, it means the backup is not
	// automatically deleted. For backups automatically created from a backup schedule, this field is assigned
	// the value of `BackupPlan.RetentionPolicy.backup_default_retain_days`. For created backups that leave this
	// field unspecified, the agent uses the value of
	// `ClusterBackupPlan.RetentionPolicy.backup_default_retain_days`. The creation of a backup with this field set
	// to a value less than `delete_lock_days` results in an invalid response from the agent. This
	// field must only be increased in an update request, or an invalid response is returned by
	// the agent immediately.
	// +optional
	RetainDays *uint `json:"retainDays,omitempty"`
}

// Defines the observed state of a `ManualClusterBackupRequest` resource.
type ManualClusterBackupRequestStatus struct {
	// The time the resource expires.
	TimeToExpire metav1.Time `json:"timeToExpire"`

	// The status of the observed state a of `ManualClusterBackupRequest` resource.
	StatusField StatusFields `json:"statusField"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-backups"
// +gdcloud:manifest:verbs=create
// +gdcloud:manifest:rbac="create:organization-cluster-backup-admin"
// Defines the schema for the `ManualClusterBackupRequest` API.
type ManualClusterBackupRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManualClusterBackupRequestSpec   `json:"spec,omitempty"`
	Status ManualClusterBackupRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of `ManualClusterBackupRequest` resources.
type ManualClusterBackupRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManualClusterBackupRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManualClusterBackupRequest{}, &ManualClusterBackupRequestList{})
}
