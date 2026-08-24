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

const (
	// PropagationConditionType is the condition type that indicates the global
	// virtualmachine image is fully propagated to all zones.
	PropagationConditionType = "Propagated"
)

// Represents the disk image that can be used on virtual machine.
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmimage, vmimages}
type VirtualMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   vmv1.VirtualMachineImageSpec `json:"spec,omitempty"`
	Status VirtualMachineImageStatus    `json:"status,omitempty"`
}

// Defines the status of a Virtual Machine Image.
type VirtualMachineImageStatus struct {
	// Conditions of the virtual machine image.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []vmv1.VMMError `json:"errors,omitempty"`

	// The list of zones where this Virtual Machine Image is stored.
	// +optional
	StorageLocations []string `json:"storageLocations,omitempty"`

	// The list of zone statuses where the resource is propagated to.
	// +listType=map
	// +listMapKey=name
	// +optional
	Zones []VirtualMachineImageZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a zone rolling out to a particular zone.
type VirtualMachineImageZoneStatus struct {
	// The name of the zone where the replica this status represents is in.
	Name string `json:"name"`

	// The observations of the current rollout.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Represents a collection of virtual machine images.
// +kubebuilder:object:root=true
type VirtualMachineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VirtualMachineImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineImage{},
		&VirtualMachineImageList{},
	)
}
