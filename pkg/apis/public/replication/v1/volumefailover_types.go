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
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName="vf"
// +kubebuilder:printcolumn:name="VRR Ref",type="string",JSONPath=".spec.volumeReplicationRelationshipRef",description="The VRR reference"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="The state of the VolumeFailover"

// Defines the Schema for the 'VolumeFailover' API.
// +genclient
type VolumeFailover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeFailoverSpec   `json:"spec,omitempty"`
	Status VolumeFailoverStatus `json:"status,omitempty"`
}

// Defines the desired state of `VolumeFailover`.
type VolumeFailoverSpec struct {
	// The unique string referencing a volume replication relationship
	VolumeReplicationRelationshipRef string `json:"volumeReplicationRelationshipRef"`
}

// Defines the state of the volume failover used in the `VolumeFailoverStatus`.
type VolumeFailoverState string

// Supported states for the volume failover.
const (
	VolumeFailoverStateRunning   VolumeFailoverState = "Running"
	VolumeFailoverStateError     VolumeFailoverState = "Error"
	VolumeFailoverStateCompleted VolumeFailoverState = "Completed"
)

// Defines the observed state of the `VolumeFailover`.
type VolumeFailoverStatus struct {
	// The current status of the volume failover.
	// +optional
	State *VolumeFailoverState `json:"state,omitempty"`

	// A human-readable message indicating details about why the volume failover is in this state.
	// +optional
	Message *string `json:"message,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of `VolumeReplicationRelationshipReplica` resources.
type VolumeFailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeFailover `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeFailover{}, &VolumeFailoverList{})
}
