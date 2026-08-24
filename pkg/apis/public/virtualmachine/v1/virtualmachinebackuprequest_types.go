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

// Defines the desired state of `VirtualMachineBackupRequest`.
type VirtualMachineBackupRequestSpec struct {
	// The name of the `VirtualMachineBackupPlanTemplate` resource used to generate a `VirtualMachineBackupPlan`.
	// Re-uses the pre-existing `VirtualMachineBackupPlan` made from this template, if there is one.
	// The `VirtualMachineBackupPlanTemplate` must live in the same `namespace` as this request.
	// Specify **only one** of either `VirtualMachineBackupPlanTemplate` OR `VirtualMachineBackupPlan`.
	// +optional
	VirtualMachineBackupPlanTemplate string `json:"virtualMachineBackupPlanTemplate,omitempty"`

	// The virtual machine backup plan which this request uses for adhoc request with backup scope
	// VirtualMachine and VirtualMachineDisk fields should be omitted when this field is provided
	// The `VirtualMachineBackupPlan` must live in the same `namespace` as this request.
	// +optional
	VirtualMachineBackupPlan string `json:"virtualMachineBackupPlan,omitempty"`

	// The `VirtualMachine` that is being backed up. This is used with the `VirtualMachineBackupPlanTemplate` to automatically
	// generate a `VirtualMachineBackupPlan` if one does not exist.
	// Specify **only one** of either `VirtualMachine` OR `VirtualMachineDisk`. When `virtualMachine` is specified, the backup strategy is always `ProvisionerSpecific`.
	VirtualMachine string `json:"virtualMachine,omitempty"`

	// The `VirtualMachineDisk` being backed-up. This is used with `VirtualMachineBackupPlanTemplate` to automatically
	// generate a `VirtualMachineBackupPlan` if one does not exist. When `virtualMachineDisk` is specified to be backed-up, the backup strategy is `SnapshotOnly`.
	// Specify **only one** of either `VirtualMachine` OR `VirtualMachineDisk`.
	VirtualMachineDisk string `json:"virtualMachineDisk,omitempty"`

	// The name of the `VirtualMachineBackup` to be created. The backup is always created in the same namespace as the request.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule=`self.matches('^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$')`,message="VirtualMachineBackupName must be a valid DNS-1123 subdomain (lowercase alphanumeric, '-', or '.')."
	// +kubebuilder:validation:MaxLength=253
	VirtualMachineBackupName string `json:"virtualMachineBackupName,omitempty"`
}

// Defines the observed state of VirtualMachineBackupRequest.
type VirtualMachineBackupRequestStatus struct {
	// When this ephemeral resource will be deleted.
	TimeToExpire metav1.Time `json:"timeToExpire"`

	// Describes the observed state of the VirtualMachineBackupRequest
	StatusField backupv1.StatusFields `json:"statusField"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmbackuprequest, vmbackuprequests}
// The Schema for the VirtualMachineBackupRequests API.
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="compute-backups"
// +gdcloud:manifest:verbs=create
// +gdcloud:manifest:rbac="create:project-vm-admin;organization-backup-admin"
// +genclient
type VirtualMachineBackupRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineBackupRequestSpec   `json:"spec,omitempty"`
	Status VirtualMachineBackupRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of VirtualMachineBackupRequest.
type VirtualMachineBackupRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineBackupRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineBackupRequest{}, &VirtualMachineBackupRequestList{})
}
