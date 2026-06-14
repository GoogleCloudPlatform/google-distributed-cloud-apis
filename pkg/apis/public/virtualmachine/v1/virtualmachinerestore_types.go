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

// Defines the desired state of `VirtualMachineRestore`.
type VirtualMachineRestoreSpec struct {

	// The name of the VM backup that this `VirtualMachineRestore` is restoring.
	// The `VirtualMachineBackup` is in the same `namespace` as this `VirtualMachineRestore`.
	VirtualMachineBackup string `json:"virtualMachineBackup"`

	// The list of all VMs.
	// created when the underlying restore is successful.
	TargetVirtualMachines []string `json:"targetVirtualMachines,omitempty"`

	//The list of all VM disks.
	// created when the underlying restore is successful.
	TargetVirtualMachineDisks []string `json:"targetVirtualMachineDisks,omitempty"`
}

// Defines the observed state of `VirtualMachineRestore`.
type VirtualMachineRestoreStatus struct {
	// The name of the underlying restore which this `VirtualMachineRestore` references.
	// The restore must be in the same `namespace` as this VM backup, used in GDC air-gapped Org v1 Architecture.
	Restore string `json:"restore,omitempty"`

	// The status of the underlying restore which this `VirtualMachineRestore` references. Used in GDC air-gapped Org v1 Architecture.
	RestoreStatus backupv1.RestoreStatus `json:"restoreStatus,omitempty"`

	// The name of the underlying Config Restore which this `VirtualMachineRestore` references. Config refers to VM and VM Disk.
	// The config restore must be in the same `namespace` as this VM Restore, used in GDC air-gapped Org v2 Architecture.
	ConfigRestore string `json:"configRestore,omitempty"`

	// The status of the underlying config restore which this `VirtualMachineRestore` references, used in GDC air-gapped Org v2 Architecture.
	ConfigRestoreStatus backupv1.RestoreStatus `json:"configRestoreStatus,omitempty"`

	// The name of the underlying Volume Restore which this `VirtualMachineRestore` references.
	// The volume restore must be in the same `namespace` as this VM Restore, used in GDC air-gapped Org v2 Architecture.
	VolumeRestore string `json:"volumeRestore,omitempty"`

	// The status of the underlying volume restore which this `VirtualMachineRestore` references, used in GDC air-gapped Org v2 Architecture.
	VolumeRestoreStatus backupv1.RestoreStatus `json:"volumeRestoreStatus,omitempty"`

	// Conditions represents the observations of this VM Restore's current state.
	// Known condition types: Succeeded
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" reflect:"unexport"`
}

const (
	VMRestoreConditionSucceeded = "Succeeded"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmrestore, vmrestores}
// The Schema for the VirtualMachineRestores API.
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="compute-restores"
// +gdcloud:manifest:verbs=describe;list;delete
// +gdcloud:manifest:rbac="describe,list,delete:organization-backup-admin,project-vm-admin"
// +gdcloud:manifest:rbac="describe,list:project-vm-viewer"
type VirtualMachineRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineRestoreSpec   `json:"spec,omitempty"`
	Status VirtualMachineRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of VirtualMachineRestore.
type VirtualMachineRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineRestore{}, &VirtualMachineRestoreList{})
}
