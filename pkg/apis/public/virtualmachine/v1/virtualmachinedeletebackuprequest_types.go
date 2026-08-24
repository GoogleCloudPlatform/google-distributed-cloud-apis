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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	backupv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/backup/v1"
)

// Defines the desired state of `VirtualMachineDeleteBackupRequest`.
type VirtualMachineDeleteBackupRequestSpec struct {
	// The name of the `VirtualMachineBackup` to delete. This request is always created in the same `namespace` as the backup.
	VirtualMachineBackupRef corev1.LocalObjectReference `json:"virtualMachineBackupRef"`

	// The set of VMBackupJob Names that need to be deleted. This can be used when entire VM Backup is not
	// required to be deleted but only subset of backups of VMs need to be deleted. List of VMBackupJob CRs for
	// a given VM Backup can be searched using the label {"backup.gdc.goog/backup-name": <VM Backup Name> + "-vb"}
	// All the targeted VMBackupJobs should be part of the same referenced VirtualMachineBackup.
	// If the list is empty, all the VMBackupJobs associated with the given VirtualMachineBackup will be deleted.
	// Once a subset of VM Backups are deleted, complete restore of the VM Backup will not include those VMs.
	// +optional
	TargetedVMBackupJobNames []string `json:"targetedVMBackupJobNames,omitempty"`
}

// Defines the observed state of `VirtualMachineDeleteBackupRequest`.
type VirtualMachineDeleteBackupRequestStatus struct {
	// Defines the time to delete this ephemeral resource.
	TimeToExpire metav1.Time `json:"timeToExpire"`

	// Describes the observed state of the `VirtualMachineDeleteBackupRequest`.
	StatusField backupv1.StatusFields `json:"statusField"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmdeletebackuprequest, vmdeletebackuprequests}
// The Schema for the `VirtualMachineDeleteBackupRequests` API.
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="compute-backups"
// +gdcloud:manifest:verbs=delete
// +gdcloud:manifest:rbac="delete:project-vm-admin;organization-backup-admin"
// +genclient
type VirtualMachineDeleteBackupRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineDeleteBackupRequestSpec   `json:"spec,omitempty"`
	Status VirtualMachineDeleteBackupRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of `VirtualMachineDeleteBackupRequest`.
type VirtualMachineDeleteBackupRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineDeleteBackupRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineDeleteBackupRequest{}, &VirtualMachineDeleteBackupRequestList{})
}
