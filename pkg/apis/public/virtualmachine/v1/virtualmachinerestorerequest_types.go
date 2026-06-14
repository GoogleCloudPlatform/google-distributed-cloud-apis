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

// Defines the desired state of `VirtualMachineRestoreRequest`.
type VirtualMachineRestoreRequestSpec struct {
	// The name of the VM backup that to restore.
	// The `VirtualMachineBackup` resides in the same `namespace` as does this `VirtualMachineRestoreRequest`.
	VirtualMachineBackup string `json:"virtualMachineBackup"`

	// The name given to the `VirtualMachineRestore` resource created.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:XValidation:rule=`self.matches('^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$')`,message="RestoreName must be a valid DNS-1123 subdomain (lowercase alphanumeric, '-', or '.')."
	RestoreName string `json:"restoreName"`

	// The prefix given to the resources which are restored
	// by `VirtualMachineBackup`. The name of the restored resource would be the prefix + name of backed up resource.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:XValidation:rule=`self.matches('^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$')`,message="RestoredResourceName must be a valid DNS-1123 subdomain (lowercase alphanumeric, '-', or '.')."
	RestoredResourceName string `json:"restoredResourceName"`

	// The description given to the newly created
	// resource.
	RestoredResourceDescription string `json:"restoredResourceDescription"`
	// The filters that can be used to refine the VM resource selection during the Restore.
	// +optional
	Filter *FilterSpec `json:"filter,omitempty"`

	// The namespace where the VM will be restored, if not provided the original namespace from the backup will be used.
	// Supported only for cluster node VMs restores.
	// +optional
	TargetNamespace string `json:"targetNamespace"`
}

// Defines the observed state of `VirtualMachineRestoreRequest`.
type VirtualMachineRestoreRequestStatus struct {
	// When this ephemeral resource will be deleted.
	TimeToExpire metav1.Time `json:"timeToExpire"`

	// Describes the observed state of the `VirtualMachineRestoreRequest`.
	StatusField backupv1.StatusFields `json:"statusField"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmrestorerequest, vmrestorerequests}
// The Schema for the VirtualMachineRestoreRequests API.
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="compute-restores"
// +gdcloud:manifest:verbs=create
// +gdcloud:manifest:rbac="create:organization-backup-admin;project-vm-admin"
// +gdcloud:manifest:skipcodegen=true
// +genclient
type VirtualMachineRestoreRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineRestoreRequestSpec   `json:"spec,omitempty"`
	Status VirtualMachineRestoreRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of VirtualMachineRestoreRequest.
type VirtualMachineRestoreRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineRestoreRequest `json:"items"`
}

// Defines the fine-grained restore Filter.
type FilterSpec struct {
	// The set of VMs that needs to be restored.
	// These VMs should be the ones that are part of the VirtualMachineBackup resource that is provided in the same VirtualMachineRestoreRequest.
	// +optional
	TargetedVirtualMachines []VirtualMachineResourceConfig `json:"targetedVirtualMachines,omitempty"`
	// The set of VMDisks that needs to be restored.
	// These VMDisks should be the ones that are part of the VirtualMachineBackup resource that is provided in the same VirtualMachineRestoreRequest.
	// +optional
	TargetedVirtualMachineDisks []VirtualMachineResourceConfig `json:"targetedVirtualMachineDisks,omitempty"`
	// Specifies the filter options for restoring disks.
	// Specifying this field will not restore VirtualMachines, it will only restore VirtualMachineDisks based on selected field in VirtualMachineDiskFilterOptions.
	// This field shouldn't be specified if one of TargetedVirtualMachines or TargetedVirtualMachineDisks is specified.
	// +optional
	VirtualMachineDiskOptions *VirtualMachineDiskFilterOptions `json:"virtualMachineDiskOptions,omitempty"`
}

// Specifies the filter options for restoring disks.
type VirtualMachineDiskFilterOptions struct {
	// Specifies that all disks from VirtualMachineBackup needs to be restored.
	AllDisks bool `json:"allDisks,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineRestoreRequest{}, &VirtualMachineRestoreRequestList{})
}
