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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConditionTypeVirtualMachineTypeReady is the Ready condition type for VirtualMachineType.
	// It indicates whether the VirtualMachineType is supported by the underlying hardware.
	ConditionTypeVirtualMachineTypeReady = "Ready"

	// Reasons for ConditionTypeVirtualMachineTypeReady.

	// ReasonHardwareCompatible indicates that the machine type is compatible with the hardware.
	ReasonHardwareCompatible = "HardwareCompatible"
	// ReasonHardwareIncompatible indicates that the machine type is incompatible with the hardware.
	ReasonHardwareIncompatible = "HardwareIncompatible"
)

// Specifies the CPU and memory attributes of a VM.
// You must specify _either_ `vcpus` and `memory` exclusively, _or_ specify only
// `VirtualMachineType`.
//
// Specifying `vcpus` without `memory`, or vice versa, creates an invalid
// combination.
// Specifying `virtualMachineType` while specifying either `vcpus` or `memory`, or
// both also creates an invalid combination.
type Compute struct {
	// Specifies the name of the referenced `VirtualMachineType`.
	// The reference requires a predefined, or golden `VirtualMachineType` name.
	VirtualMachineType string `json:"virtualMachineType,omitempty"`
	// Specifies the number of VCPUs that are available to the instance.
	// Specify `vcpus` as an integer. This value must be a
	// multiple of 2, with 2 as the minimum and 128 as the maximum allowed.
	VCPUs uint32 `json:"vcpus,omitempty"`
	// Specifies the amount of physical memory available to the instance.
	// memory must have a minimum value of `1Gi`, and can be up to (including) `400Gi`.
	Memory *resource.Quantity `json:"memory,omitempty"`
}

// Defines the CPU and Memory resource of a VM.
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmtype, vmtypes}
// +kubebuilder:printcolumn:name="Supported",type=boolean,JSONPath=".status.supported"
// +gdcloud:manifest:relevant=false,oc=vmm
type VirtualMachineType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineTypeSpec   `json:"spec"`
	Status            VirtualMachineTypeStatus `json:"status,omitempty"`
}

// Defines the configurations of a `VirtualMachineType`.
type VirtualMachineTypeSpec struct {
	// Specifies the number of VCPUs that are available to the instance.
	// Specify `vcpus` as an integer that is a
	// multiple of 2, between 2 and 128, inclusive.
	VCPUs uint32 `json:"vcpus"`
	// Specifies the amount of physical memory available to the instance.
	// `memory` must have a value that is between 1G and 400G, inclusive.
	Memory resource.Quantity `json:"memory"`
}

// Defines the observed state of VirtualMachineType.
type VirtualMachineTypeStatus struct {
	// Specifies if given vmtype is supported or not by the underlying hardware
	// +optional
	Supported *bool `json:"supported,omitempty"`

	// The latest observations of the `VirtualMachineType` state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func (v VirtualMachineTypeStatus) IsSupported() bool {
	return v.Supported != nil && *v.Supported
}

// +kubebuilder:object:root=true
// VirtualMachineTypeList contains a list of VirtualMachineType.
type VirtualMachineTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineType `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineType{},
		&VirtualMachineTypeList{},
	)
}
