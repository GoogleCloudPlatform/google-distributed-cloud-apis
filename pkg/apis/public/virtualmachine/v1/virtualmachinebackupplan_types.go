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

	backupv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/backup/v1"
)

const VMBackupPlanTemplateAnnotation = "virtualmachine.gdc.goog/vmbp-template"

// Defines a `VirtualMachineBackupPlan` which provides instructions for creating an
// underlying `BackupPlan` resource and `ProtectedApplication` to perform backups.
type VirtualMachineBackupPlanSpec struct {
	// The `VirtualMachineBackupPlanTemplate` must live in the same `namespace` as this vm backup plan.
	// +optional
	VirtualMachineBackupPlanTemplate string `json:"virtualMachineBackupPlanTemplate,omitempty"`
	// Defines the configuration and scope of the backup.
	BackupConfig VirtualMachineBackupConfig `json:"backupConfig"`
	// The scheduled backup creation under this VM backup plan.
	// Schedule is a mutable field which can be edited after creation
	// By default, VM backup plan will be paused
	// +optional
	// +kubebuilder:default:={paused: true, cronSchedule: "0 * * * *"}
	BackupSchedule *backupv1.Schedule `json:"backupSchedule"`
	// The lifecycle of backups created under this plan.
	// by default, backups can be deleted without any locking period
	// by default, backups are not deleted by automatic schedule and retention policy can be modified
	// +optional
	// +kubebuilder:default:={backupDeleteLockDays: 0, backupRetainDays: 0, locked: false}
	RetentionPolicy *backupv1.RetentionPolicy `json:"retentionPolicy"`
	// Specifies whether the plan has been deactivated.
	// Setting this field to 'true' locks the plan meaning no further updates
	// are allowed, including changes to the deactivated field. It also prevents new
	// backups from being created under this plan, both manually or scheduled.
	// Default to 'false'.
	// +optional
	// +kubebuilder:default:=false
	Deactivated bool `json:"deactivated,omitempty"`
}

// Contains configuration details for executing the backup, including scope, location,
// and volume backup strategy.
type VirtualMachineBackupConfig struct {
	// Identifies the secondary storage location for this `VirtualMachineBackupPlan`.
	// This field is meant for internal use only. Provide Backup repository only through VirtualMachineBackupPlanTemplate
	// +optional
	BackupRepository string `json:"backupRepository"`
	// Specifies the resource(s) covered by this `VirtualMachineBackupPlan`.
	BackupScope VirtualMachineBackupScope `json:"backupScope"`
	// Declares the strategy to use for backing-up volumes; for example, use a local snapshot vs. using remote or provisioner-specific backup.
	VolumeStrategy backupv1.VolumeStrategy `json:"volumeStrategy"`
}

// Defines the scope of resources for the `VirtualMachineBackupPlan` to capture.
type VirtualMachineBackupScope struct {
	// Specifies the VMs for the `VirtualMachineBackupPlan` to capture.
	// +optional
	SelectedVirtualMachines []VirtualMachineResourceConfig `json:"selectedVirtualMachines,omitempty"`
	// Specifies the disks for this `VirtualMachineBackupPlan` to capture.
	// +optional
	SelectedVirtualMachineDisks []VirtualMachineResourceConfig `json:"selectedVirtualMachineDisks,omitempty"`
	// Specifies the label(s) selecting one/multiple VM or VM disk resources in same namespace as backup plan
	// +optional
	VMResourceLabelSelector map[string]string `json:"vmResourceLabelSelector,omitempty"`
}

// Specifies a VM resource and additional parameters for backing up that resource.
type VirtualMachineResourceConfig struct {
	// The name of the resource being backed up. It must exist in the same `namespace` as the plan.
	ResourceName string `json:"resourceName"`
}

// Defines the observed state of `VirtualMachineBackupPlan`.
type VirtualMachineBackupPlanStatus struct {
	// The name of the underlying backup plan managed by this `VirtualMachineBackupPlan`, used in GDC air-gapped Org v1 Architecture.
	BackupPlan *string `json:"backupPlan,omitempty"`
	// The embedded status of the underlying backup plan.
	BackupPlanStatus *backupv1.BackupPlanStatus `json:"backupPlanStatus,omitempty"`
	// The name of the underlying config backup plan managed by this `VirtualMachineBackupPlan`. Config here refers to VM and VM Disk, used in GDC air-gapped Org v2 Architecture.
	ConfigBackupPlan *string `json:"configBackupPlan,omitempty"`
	// The name of the underlying volume backup plan managed by this `VirtualMachineBackupPlan`, used in GDC air-gapped Org v2 Architecture.
	VolumeBackupPlan *string `json:"volumeBackupPlan,omitempty"`
	// The timestamp for the most recently executed backup.
	// This field is used to schedule next backup
	// +optional
	LastBackupTime metav1.Time `json:"lastBackupTime,omitempty"`
	// The timestamp for the next scheduled backup.
	// This field is used to schedule next backup
	// +optional
	NextBackupTime metav1.Time `json:"nextBackupTime,omitempty"`
	// The timestamp of last reconciliation of this resource.
	// +optional
	LastReconcileTime metav1.Time `json:"lastReconcileTime,omitempty"`
	// ErrorMessage holds the error message that occurred during the last reconciliation attempt.
	// This field is optional and will only be populated if an error occurred.
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Conditions represents the observations of this VM backup plan's current state.
	// Known condition types: Ready
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

const (
	VMBackupPlanConditionReady = "Ready"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmbackupplan, vmbackupplans}
// The Schema for the VirtualMachineBackupPlans API.
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="compute-backup-plans"
// +gdcloud:manifest:verbs=create;list;describe;delete;update
// +gdcloud:manifest:rbac="create:organization-backup-admin"
// +gdcloud:manifest:rbac="describe,list:project-vm-admin;project-vm-viewer;organization-backup-admin"
// +gdcloud:manifest:rbac="delete,update:project-vm-admin;organization-backup-admin"
type VirtualMachineBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineBackupPlanSpec   `json:"spec,omitempty"`
	Status VirtualMachineBackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of VirtualMachineBackupPlan.
type VirtualMachineBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineBackupPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineBackupPlan{}, &VirtualMachineBackupPlanList{})
}
