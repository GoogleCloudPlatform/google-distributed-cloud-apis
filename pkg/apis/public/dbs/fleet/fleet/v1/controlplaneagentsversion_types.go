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

/*
Copyright 2022.
*/

package v1

import (
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced

// ControlPlaneAgentsVersion specifies the new image path for each control plane agent component
// during the DBC provisioning and upgrade.
type ControlPlaneAgentsVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControlPlaneAgentsVersionSpec   `json:"spec,omitempty"`
	Status ControlPlaneAgentsVersionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:generate=true

type ControlPlaneAgentsVersionSpec struct {
	// Database engine type.
	// +required
	// +kubebuilder:validation:Enum=PostgreSQL;Oracle;AlloyDBOmni
	DatabaseEngine EngineType `json:"databaseengine,omitempty"`

	// The version of the current control plane agents.
	// +required
	Version string `json:"version,omitempty"`

	// Image path for each of the control plane components.
	Components []ControlPlaneAgentComponent `json:"components"`
}

type ControlPlaneAgentsVersionStatus struct {
	// the controller issues a Ready condition
	// when all images are ready.
	occoreapi.EntityStatus `json:",inline"`
}

//+kubebuilder:object:generate=true

type ControlPlaneAgentComponent struct {
	// Name of this component. For example dbinit, monitoring
	// +required
	Name string `json:"name,omitempty"`

	// Image path of this component.
	// +required
	Uri string `json:"uri,omitempty"`
}

// +kubebuilder:object:root=true

// ControlPlaneAgentsVersionList contains a list of ControlPlaneAgentsVersion.
type ControlPlaneAgentsVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPlaneAgentsVersion `json:"items"`
}

const (
	ControlPlaneAgentsVersionReadyType   = "Ready"
	ControlPlaneAgentsVersionValidReason = "Valid"
)

func init() {
	SchemeBuilder.Register(&ControlPlaneAgentsVersion{}, &ControlPlaneAgentsVersionList{})
}

func (in *ControlPlaneAgentsVersion) EntityStatus() *occoreapi.EntityStatus {
	return &in.Status.EntityStatus
}
