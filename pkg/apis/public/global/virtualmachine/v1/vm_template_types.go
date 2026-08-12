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

	globalv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
)

const (
	// VirtualMachineTemplateReadyConditionType is the condition type that indicates
	// that the VirtualMachineTemplate is ready for use.
	VirtualMachineTemplateReadyConditionType = "Ready"

	// VirtualMachineTemplateReplicatingConditionType is the condition type that indicates
	// that the VirtualMachineTemplate is replicating.
	VirtualMachineTemplateReplicatingConditionType = "Replicating"

	// VirtualMachineTemplateReplicatingConditionType is the condition type that indicates
	// that the VirtualMachineTemplate is terminating.
	VirtualMachineTemplateTerminatingConditionType = "Terminating"
)

// Represents a template of a Virtual Machine.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmtemplate, vmtemplates}
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
type VirtualMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineTemplateSpec   `json:"spec,omitempty"`
	Status VirtualMachineTemplateStatus `json:"status,omitempty"`
}

// Contains the observed state of the Virtual Machine Template.
type VirtualMachineTemplateStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []VirtualMachineTemplateZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a Virtual Machine Template rolling out to a particular zone.
type VirtualMachineTemplateZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus VirtualMachineTemplateReplicaStatus `json:"replicaStatus,omitempty"`
}

// Defines a list of `VirtualMachineTemplate` resources.
// +kubebuilder:object:root=true
type VirtualMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineTemplate{},
		&VirtualMachineTemplateList{},
	)
}
