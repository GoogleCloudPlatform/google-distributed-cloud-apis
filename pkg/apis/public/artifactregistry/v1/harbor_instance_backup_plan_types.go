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
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LastBackupTime",type="string",JSONPath=".status.lastBackupTime"
// +kubebuilder:printcolumn:name="NextBackupTime",type="string",JSONPath=".status.nextBackupTime"
// +kubebuilder:printcolumn:name="Paused",type="boolean",JSONPath=".spec.backupSchedule.paused"
// +kubebuilder:printcolumn:name="Instance",type="string",JSONPath=".spec.backupConfig.backupScope.harborInstance"
// +gdcloud:manifest:relevant=true,oc=haas,component=harbor
// +gdcloud:manifest:entities="backup-plans",verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:harbor-instance-admin"
// +gdcloud:manifest:rbac="describe,list:harbor-instance-viewer"
// Defines the schema for the `BackupPlan` API for HarborInstance.
type HarborInstanceBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the backup plan.
	// +kubebuilder:validation:Required
	Spec HarborInstanceBackupPlanSpec `json:"spec"`
	// The most recently observed status of the backup plan.
	Status HarborInstanceBackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of backup plans for Harbor instance.
type HarborInstanceBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HarborInstanceBackupPlan `json:"items"`
}

// Defines the desired state of a backup plan.
type HarborInstanceBackupPlanSpec struct {
	// The scheduled backup creation under this backup plan.
	// +kubebuilder:validation:Required
	BackupSchedule *Schedule `json:"backupSchedule"`
	// The backup configuration of this backup plan.
	// +kubebuilder:validation:Required
	BackupConfig *BackupConfig `json:"backupConfig"`

	// The lifecycle of backups created under this plan.
	// +optional
	// +kubebuilder:default:={backupRetainDays: 0}
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
	// A user-specified descriptive string for this backup plan.
	// +optional
	Description string `json:"description,omitempty"`
}

// Represents an inner message type that defines the configuration of creating
// a backup from this backup plan.
type BackupConfig struct {
	// The name of the BackupRepository resource identifying the secondary storage for this BackupPlan resource.
	BackupRepository string `json:"backupRepository,omitempty"`

	// The resource selection scope of a backup.
	// +kubebuilder:validation:Required
	BackupScope BackupScope `json:"backupScope"`
}

// Defines the Harbor Instance to back up.
type BackupScope struct {
	// Harbor Instance name to backup in the same namespace.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Instance in BackupScope is immutable"
	HarborInstance string `json:"harborInstance,omitempty"`
}

// Defines a policy that determines when to automatically delete backups created under this backup
// plan, a plan-level minimum number of backup retain days, and a lock to disallow any policy
// updates.
type RetentionPolicy struct {
	// The number of days after which the service deletes a backup.
	// If specified, a backup created under this backup plan is
	// automatically deleted when it reaches the backup_retain_days after the
	// create_time.
	// If not specified, backups created under this backup plan is not
	// subject to automatic deletion.
	// Updating this field does not affect existing backups under it. Backups
	// created after a successful update automatically inherit the new
	// value.
	// +optional
	// +kubebuilder:default:=0
	BackupRetainDays uint `json:"backupRetainDays"`
}

// Represents an inner message type that defines a cron schedule.
type Schedule struct {
	// A cron string schedule on which an operation is executed.
	// +kubebuilder:validation:Required
	CronSchedule string `json:"cronSchedule"`
	// Specifies whether the scheduled operation is paused or unpaused.
	// If set to True, the scheduled operation will be paused and no automated
	// backup will be created.
	// Default to False.
	// +optional
	// +kubebuilder:default:=false
	Paused bool `json:"paused"`
}

// The various states a backup plan can be in.
// +kubebuilder:validation:Enum=Unspecified;NotReady;Ready;Paused
type BackupPlanState string

const (
	BackupPlanStateUnspecified BackupPlanState = "Unspecified"
	BackupPlanStateReady       BackupPlanState = "Ready"
	BackupPlanStateNotReady    BackupPlanState = "NotReady"
	BackupPlanStatePaused      BackupPlanState = "Paused"
)

// Defines the observed state of a backup plan.
type HarborInstanceBackupPlanStatus struct {
	// The timestamp for the most recently executed backup.
	// +optional
	LastBackupTime metav1.Time `json:"lastBackupTime,omitempty"`
	// The timestamp for the next scheduled backup.
	// +optional
	NextBackupTime metav1.Time `json:"nextBackupTime,omitempty"`
	// Conditions:
	// - Ready: readiness of the backup plan, any error when reconciling embedded object will be surfaced here.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// The current state of the backup plan.
	// +optional
	State BackupPlanState `json:"state,omitempty"`
	// A human-readable description of why the backup plan is in the current state.
	// +optional
	StateReason string `json:"reason,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&HarborInstanceBackupPlan{},
		&HarborInstanceBackupPlanList{},
	)
}
