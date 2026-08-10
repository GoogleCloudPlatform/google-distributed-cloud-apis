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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DiscoveredVM CR
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=m4gdc
// +genclient
type DiscoveredVM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiscoveredVMSpec   `json:"spec"`
	Status            DiscoveredVMStatus `json:"status,omitempty"`
}

type DiscoveredVMSpec struct {
}

type DiscoveredVMDetails struct {
	VMID string `json:"VMID"` // MORef from the vCenter
	// Specifies the number of VCPUs that are available to the instance.
	// Specify `vcpus` as an integer.
	// +optional
	VCPUs uint32 `json:"vCPU"`
	// Specifies the amount of physical memory available to the instance.
	// +optional
	Memory *resource.Quantity `json:"memory"` // in MB
	// +optional
	DiskCount int32 `json:"diskCount"`
}

type DiscoveredVMStatus struct {
	// +optional
	VMDetails DiscoveredVMDetails `json:"VMDetails,omitempty"`
	// +optional
	MigrationStatus MigrationStatus `json:"migrationStatus,omitempty"`
}

// Enum for MigrationStatus
type MigrationStatus string

const (
	MigrationNotStarted MigrationStatus = "Not Started"
	MigrationInProgress MigrationStatus = "In Progress"
	MigrationCompleted  MigrationStatus = "Completed"
	MigrationFailed     MigrationStatus = "Failed"
)

// +kubebuilder:object:root=true
type DiscoveredVMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DiscoveredVM `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveredVM{}, &DiscoveredVMList{})
}
