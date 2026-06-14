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

	backupv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/backup/v1"
)

// Defines the desired state of `VirtualMachineBackup`.
type VirtualMachineBackupSpec struct {

	// The name of the `VirtualMachineBackupPlan` from which this `VirtualMachineBackup`
	// was created. This `VirtualMachineBackupPlan` exists in the same `namespace` as the
	// `VirtualMachineBackupPlan`.
	VirtualMachineBackupPlan string `json:"virtualMachineBackupPlan"`

	// Defines the configuration and scope of the backup.
	BackupConfig VirtualMachineBackupConfig `json:"backupConfig"`
}

// Defines the observed state of a `VirtualMachineBackup`.
type VirtualMachineBackupStatus struct {
	// The name of the underlying backup which this `VirtualMachineBackup` references.
	// The backup must be in the same `namespace` as this VM backup, used in GDC air-gapped Org v1 Architecture.
	Backup string `json:"backup,omitempty"`

	// The status of the underlying backup(s) which this `VirtualMachineBackup` references, used in GDC air-gapped Org v1 Architecture.
	BackupStatus backupv1.BackupStatus `json:"backupStatus,omitempty"`

	// The name of the underlying Config Backup which this `VirtualMachineBackup` references. Config refers to VM and VM Disk.
	// The config backup must be in the same `namespace` as this VM backup, used in GDC air-gapped Org v2 Architecture.
	ConfigBackup string `json:"configBackup,omitempty"`

	// The status of the underlying config backup which this `VirtualMachineBackup` references, used in GDC air-gapped Org v2 Architecture.
	ConfigBackupStatus backupv1.BackupStatus `json:"configBackupStatus,omitempty"`

	// The name of the underlying Volume Backup which this `VirtualMachineBackup` references.
	// The volume backup must be in the same `namespace` as this VM backup, used in GDC air-gapped Org v2 Architecture.
	VolumeBackup string `json:"volumeBackup,omitempty"`

	// The status of the underlying volume backup which this `VirtualMachineBackup` references, used in GDC air-gapped Org v2 Architecture.
	VolumeBackupStatus backupv1.BackupStatus `json:"volumeBackupStatus,omitempty"`

	// Lists the names of all VMs that are included in this backup.
	BackedUpVirtualMachines []string `json:"backedUpVirtualMachines,omitempty"`

	// Lists the names of all the VM disks that are included in this backup.
	// If this backup is a disk snapshot, this is a list of VM disks that you provide.
	// If this is a VM backup, it is the list of the disks which back that VM.
	BackedUpVirtualMachineDisks []string `json:"backedUpVirtualMachineDisks,omitempty"`

	// Count of successful VMBackupJobs
	// it represents the number of VM Resources that have been successfully backed up.
	// This will be populated once the backup ends
	SuccessfulVMBackupJobs int `json:"successfulVMBackupJobs,omitempty"`

	// Count of failed VMBackupJobs
	// it represents the number of VM Resources that have failed to be backed up.
	// This will be populated once the backup ends
	FailedVMBackupJobs int `json:"failedVMBackupJobs,omitempty"`

	// Lists the names of VMs that were successfully backed up.
	SuccessfulVirtualMachines []string `json:"successfulVirtualMachines,omitempty"`

	// Lists the names of VMs that failed to be backed up.
	FailedVirtualMachines []string `json:"failedVirtualMachines,omitempty"`

	// Lists the names of VM disks that were successfully backed up.
	SuccessfulVirtualMachineDisks []string `json:"successfulVirtualMachineDisks,omitempty"`

	// Lists the names of VM disks that failed to be backed up.
	FailedVirtualMachineDisks []string `json:"failedVirtualMachineDisks,omitempty"`

	// Conditions represents the observations of this VM backup's current state.
	// Known condition types: Succeeded
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" reflect:"unexport"`
}

const (
	VMBackupConditionSucceeded = "Succeeded"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmbackup, vmbackups}
// The Schema for the VirtualMachineBackups API.
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="compute-backups"
// +gdcloud:manifest:verbs=describe;list
// +gdcloud:manifest:rbac="describe,list:project-vm-admin;project-vm-viewer;organization-backup-admin"
// +genclient
type VirtualMachineBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineBackupSpec   `json:"spec,omitempty"`
	Status VirtualMachineBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of VirtualMachineBackup.
type VirtualMachineBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineBackup{}, &VirtualMachineBackupList{})
}
