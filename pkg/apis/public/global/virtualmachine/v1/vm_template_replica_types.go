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

	vmv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/virtualmachine/v1"
)

// Represents a template of a Virtual Machine Template.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmtemplatereplica, vmtemplatereplicas}
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
type VirtualMachineTemplateReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineTemplateSpec          `json:"spec,omitempty"`
	Status VirtualMachineTemplateReplicaStatus `json:"status,omitempty"`
}

// Defines the specification of a Virtual Machine Template.
type VirtualMachineTemplateSpec struct {
	// Specifies the CPU and Memory of the VM.
	// CPU and Memory can be defined directly or through the VirtualMachineType.
	// Changes to Compute require a reboot to take effect.
	// Compute is immutable when the VM is in Unknown state.
	// +optional
	Compute vmv1.Compute `json:"compute"`

	// Disks specifies the list of disks to create and attach.
	Disks []VirtualMachineDiskTemplate `json:"virtualMachineDiskTemplates"`

	// Specifies the list of startup scripts for the VM.
	// StartupScripts only take effect on VMs that have cloud-init installed.
	// They are executed in alphabetical order, based on the name of each startup script.
	// +optional
	StartupScripts []vmv1.StartupScript `json:"startupScripts"`

	// Specifies the VM's guest environment configuration.
	// If the field is nil the enable field in AccessManagement is true by default.
	// Otherwise, the non-nil configuration for each sub-feature inside the structure
	// overrides the default configuration of the sub-feature.
	// +optional
	GuestEnvironment *vmv1.GuestEnvironment `json:"guestEnvironment,omitempty"`

	// Specifies the VM's security-related configurations.
	// +optional
	ShieldConfig *vmv1.ShieldConfig `json:"shieldConfig,omitempty"`

	// Specifies the network configuration.
	Network *vmv1.NetworkSpec `json:"network,omitempty"`

	// GracefulShutdown configures graceful shutdown for the VM.
	// +optional
	GracefulShutdown *vmv1.GracefulShutdown `json:"gracefulShutdown,omitempty"`
}

// VirtualMachineTemplateDisk defines the properties of a disk to be created and
// attached to a VM instantiated from a template.
type VirtualMachineDiskTemplate struct {
	// Specifies whether this disk will be the boot device.
	// Exactly one disk in the template must be marked as the boot disk.
	Boot *bool `json:"boot,omitempty"`

	// Specifies whether the disk should be automatically deleted when the
	// VirtualMachine is deleted.
	AutoDelete *bool `json:"autoDelete,omitempty"`

	// Spec defines the properties of the disk itself (e.g., size, source image).
	Spec vmv1.VirtualMachineDiskSpec `json:"spec"`
}

// Contains the observed state of the Virtual Machine Template Replica.
type VirtualMachineTemplateReplicaStatus struct {
	// Conditions of the Virtual Machine Template Replica.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Represents a collection of `VirtualMachineTemplateReplica` resources.
// +kubebuilder:object:root=true
type VirtualMachineTemplateReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineTemplateReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineTemplateReplica{},
		&VirtualMachineTemplateReplicaList{},
	)
}
