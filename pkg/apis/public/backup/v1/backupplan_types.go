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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// TODO(b/200285234): This file is heavily patterned off of:
//   cs//depot/google3/blaze-out/genfiles/google/cloud/gkebackup/v1alpha1/backup_plan.pb.go
// Going forward, we should come up with a good process to keep these
// relatively in sync.

// Defines the desired state of a backup plan.
type BackupPlanSpec struct {
	// The name of the cluster containing the data this backup plan backs up.
	// By default, it is the cluster in which this backup plan is created.
	// +optional
	ClusterName string `json:"clusterName,omitempty" reflect:"unexport"`
	// The scheduled backup creation under this backup plan.
	// +kubebuilder:validation:Required
	BackupSchedule *Schedule `json:"backupSchedule" reflect:"unexport"`
	// The backup configuration of this backup plan.
	// +kubebuilder:validation:Required
	BackupConfig BackupConfig `json:"backupConfig" reflect:"unexport"`

	///// P2 Fields, Do Not Need To Be Implemented for V0 /////

	// The lifecycle of backups created under this plan.
	// +optional
	// +kubebuilder:default:={backupDeleteLockDays: 0, backupRetainDays: 0, locked: false}
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy" reflect:"unexport"`
	// A user-specified descriptive string for this backup plan.
	// +optional
	Description string `json:"description,omitempty" reflect:"unexport"`
	// Specifies whether the plan has been deactivated.
	// Setting this field to ‘True’ locks the plan meaning no further updates
	// are allowed, including changes to the deactivated field. It also prevents new
	// backups from being created under this plan, both manually or scheduled.
	// Default to ‘False’.
	// +optional
	// +kubebuilder:default:=false
	Deactivated bool `json:"deactivated" reflect:"unexport"`
}

// Defines a policy that determines when to automatically delete backups created under this backup plan, a
// plan-level minimum number of backup retain days,
// plan-level minimum number of backups to be retained,
// and a lock to disallow any policy updates.
type RetentionPolicy struct {
	///// P2 Fields, Do Not Need To Be Implemented for V0 /////

	// The number of days during for which the deletion of a backup created under this
	// backup plan is blocked.
	// Default to `0`.
	// This field must be an integer value between `0-90`.
	// A backup created under this backup plan is not deletable until it
	// reaches the combined value of the backup's `create_time` and `backup_delete_lock_days`.
	// Updating this backup plan field does not affect existing backups
	// under it. Backups created after a successful update automatically
	// inherit the new value.
	// +optional
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=90
	BackupDeleteLockDays uint `json:"backupDeleteLockDays" reflect:"unexport"`
	// The number of days after which the service deletes a backup.
	// If specified, a backup created under this backup plan are
	// automatically deleted after it reaches the the combined value of `create_time` and
	// `backup_retain_days`.
	// If not specified, backups created under this backup plan are not
	// subject to automatic deletion.
	// Updating this field does not affect existing backups under it. Backups
	// created after a successful update automatically inherit the new
	// value.
	// Note, specifying a value for `backup_retain_days` that is less than
	// `backup_delete_lock_days` at the time of the creation or update is considered
	// invalid, and the request is rejected immediately.
	// +optional
	// +kubebuilder:default:=0
	BackupRetainDays uint `json:"backupRetainDays" reflect:"unexport"`
	// Specifies whether the retention policy of this backup plan is locked.
	// If set to `True`, no further update is allowed on this policy, including
	// the `locked` field itself.
	// Default to `False`.
	// +optional
	// +kubebuilder:default:=false
	Locked bool `json:"locked" reflect:"unexport"`
	// The number of backups to retain.
	// If specified, oldest backups of backup plan which reached its combined value of `create_time` and
	// `backup_delete_lock_days are automatically deleted after the total number of backups under
	// this backup plan reaches the value of `num_backups_to_retain`.
	// If not specified, backups created under this backup plan may or may not
	// subject to automatic deletion based on 'backup_retain_days'.
	// Updating this field does affect existing backups under the backup plan.
	// Note, num_backups_to_retain must be greater than or equal to the number of backups
	// that can be created within backup_delete_lock_days.
	// +optional
	// +kubebuilder:default:=0
	NumBackupsToRetain uint `json:"numBackupsToRetain" reflect:"unexport"`
}

// Represents an inner message type that defines a cron schedule.
type Schedule struct {
	// A cron string schedule on which an operation is executed.
	// +kubebuilder:validation:Required
	CronSchedule string `json:"cronSchedule" reflect:"unexport"`
	// Specifies whether the scheduled operaction is inactive or active.
	// If set to `True`, the scheduled operation will be inactive.
	// Default to `True`.
	// +optional
	// +kubebuilder:default:=true
	Paused bool `json:"paused" reflect:"unexport"`
	// Specifies the maximum number of active scheduled backups that can exist under a backup plan at the same time.
	// If specified, no more than this number of scheduled backups from this backup plan can be in the active state at any given time. If not specified, it defaults to 2.
	// +optional
	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=7
	MaxConcurrency *int `json:"maxConcurrency,omitempty" reflect:"unexport"`
}

// Represents the name and state of a backup in the `RecentBackups` field .
type BackupInfo struct {
	// The name of the backup.
	// +optional
	Name string `json:"name" reflect:"unexport"`
	// The status of the most recently executed backup.
	// +optional
	State BackupState `json:"state,omitempty"`
}

// Defines the observed state of a backup plan.
type BackupPlanStatus struct {
	// The timestamp for the most recently executed backup.
	// +optional
	LastBackupTime metav1.Time `json:"lastBackupTime,omitempty"`
	// The timestamp for the next scheduled backup.
	// +optional
	NextBackupTime metav1.Time `json:"nextBackupTime,omitempty"`
	// TODO(b/205148277): add more status/conditions e.g.: to signal successful scheduling

	// The current state of the backup plan.
	// +optional
	State State `json:"state,omitempty"`
	// The state of the most recently created backup.
	// +optional
	LastBackupState BackupState `json:"lastBackupState,omitempty"`
	// A human-readable description of why the last backup is in the current state.
	// +optional
	LastBackupStateReason string `json:"lastBackupStateReason,omitempty"`
	// The most recent errors with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
	// Conditions represents the observations of this backup plan's current state.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// The most recent backups created by the backup plan.
	// +optional
	RecentBackups []BackupInfo `json:"recentBackups,omitempty"`
	// Counts from the most recent successful backup, used for project-level attribution.
	// LastSuccessfulBackupPodCountByNamespace maps namespace to the number of workload pods in the last successful backup.
	// +optional
	LastSuccessfulBackupPodCountByNamespace map[string]int64 `json:"lastSuccessfulBackupPodCountByNamespace,omitempty"`
	// LastSuccessfulBackupVmCountByNamespace maps namespace to the number of virtual machines in the last successful backup.
	// +optional
	LastSuccessfulBackupVmCountByNamespace map[string]int64 `json:"lastSuccessfulBackupVmCountByNamespace,omitempty"`
	// LastSuccessfulBackupName is the name of the most recent successful backup
	// used as the source of truth for the telemetry maps.
	// +optional
	LastSuccessfulBackupName string `json:"lastSuccessfulBackupName,omitempty"`
}

// A value specifying what state the backup plan is currently in.
// +kubebuilder:validation:Enum=ReadOnlyInactive
type State string

const (
	// A state that shows whether the current backup plan was
	// imported by a read-only backup repository. It will neither create
	// backups or persist itself in object storage.
	ReadOnlyInactive State = "ReadOnlyInactive"
)

const (
	BackupPlanConditionReady = "Ready"
	// Used for indicating whether per-namespace telemetry is populated in a BackupPlan.
	BillingTelemetrySynced = "BackupUsageTelemetrySynchronized"

	// Reasons for BackupUsageTelemetrySynchronized condition.
	ReasonNoSuccessfulBackups   = "NoSuccessfulBackups"
	ReasonTelemetrySynced       = "TelemetrySynced"
	ReasonTelemetryBackfilled   = "TelemetryBackfilled"
	ReasonBackupPlanDeactivated = "BackupPlanDeactivated"

	// Messages for BackupUsageTelemetrySynchronized condition.
	MessageNoSuccessfulBackups       = "No successful backups found in this plan, per-namespace telemetry cleared"
	MessageTelemetrySyncedFormat     = "Per-namespace telemetry synced from backup %s"
	MessageTelemetryBackfilledFormat = "Per-namespace telemetry backfilled from legacy status of backup %s"
	MessageBackupPlanDeactivated     = "BackupPlan is deactivated, per-namespace telemetry cleared"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="LastBackupTime",type="string",JSONPath=".status.lastBackupTime"
// +kubebuilder:printcolumn:name="LastBackupState",type="string",JSONPath=".status.lastBackupState"
// +kubebuilder:printcolumn:name="NextBackupTime",type="string",JSONPath=".status.nextBackupTime"
// +kubebuilder:printcolumn:name="Paused",type="boolean",JSONPath=".spec.backupSchedule.paused"

// +genclient
// +gdcloud:manifest:relevant=false,oc=back
// Defines the schema for the `BackupPlan` API.
type BackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupPlanSpec   `json:"spec,omitempty"`
	Status BackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of backup plans.
type BackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupPlan{}, &BackupPlanList{})
}
