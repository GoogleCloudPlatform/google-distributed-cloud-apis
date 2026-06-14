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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=behs
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".metadata.labels['networking\\.private\\.gdc\\.goog/cluster-name']"
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".metadata.labels['networking\\.private\\.gdc\\.goog/service-namespace']"
// +kubebuilder:printcolumn:name="Service",type="string",JSONPath=".metadata.labels['networking\\.private\\.gdc\\.goog/service-name']"

// Detailed health status for a group of endpoints belonging to a Backend.
type BackendEndpointHealthSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Defines the observed state of the BackendEndpointHealthSet.
	// +optional
	Status BackendEndpointHealthSetStatus `json:"status,omitempty"`
}

// Defines the observed state of the BackendEndpointHealthSet.
type BackendEndpointHealthSetStatus struct {
	// Contains the health status for a group of endpoints belonging to a Backend.
	// +optional
	// +listType=atomic
	EndpointHealthList []EndpointHealth `json:"endpointHealthList,omitempty"`
}

// Contains the health state constraints of a single endpoint.
type EndpointHealth struct {
	// IP address of the endpoint.
	// +optional
	IPAddress string `json:"ipAddress,omitempty"`

	// Name of the instance or pod (e.g. vm-01, pod-xyz1).
	// +optional
	Instance string `json:"instance,omitempty"`

	// Health state of the endpoint.
	// Valid values are: HEALTHY, UNHEALTHY, DRAINING.
	// +optional
	HealthState string `json:"healthState,omitempty"`

	// Last time the health state changed.
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of BackendEndpointHealthSet.
type BackendEndpointHealthSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendEndpointHealthSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackendEndpointHealthSet{}, &BackendEndpointHealthSetList{})
}
