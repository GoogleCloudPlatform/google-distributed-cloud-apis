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

// CloudNATMapping is the Schema for the CloudNATMapping API.
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cloudnatmappings,singular=cloudnatmapping
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CloudNATMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Status defines the list of the Cloud NAT mappings.
	// This captures the current mapping data for the environment.
	Status CloudNATMappingStatus `json:"status,omitempty"`
}

// CloudNATMappingStatus defines the list of CloudNATMapping.
type CloudNATMappingStatus struct {
	// EndpointEgressNATMappings is a list of individual mappings between
	// endpoints and their assigned egress IPs.
	// Each object contains a maximum of 100 endpoints,
	// larger lists are split across multiple objects.
	// +optional
	// +listType=atomic
	EndpointEgressNATMappings []EndpointEgressNATMapping `json:"endpointEgressNATMappings,omitempty"`

	// LastUpdateTime represents the last time the EndpointEgressNATMappings
	// was successfully synchronized.
	// This field is updated on every successful refresh of the mapping data.
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`

	// A list of conditions describing the current state of the CloudNATMappings.
	// Known condition types are:
	// * "Ready"
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// EndpointType defines the category of the source endpoint.
// +kubebuilder:validation:Enum=Pod;VM;SystemEP
type EndpointType string

const (
	// EndpointTypePod represents a standard Kubernetes Pod.
	EndpointTypePod EndpointType = "Pod"
	// EndpointTypeVM represents a Virtual Machine endpoint.
	EndpointTypeVM EndpointType = "VM"
	// EndpointTypeSystemEP represents a 1-P system endpoint.
	EndpointTypeSystemEP EndpointType = "SystemEP"
)

// EndpointEgressNATMapping defines the mapping details for a single endpoint.
type EndpointEgressNATMapping struct {
	// EndpointName is the name of the source Pod or VM.
	// +kubebuilder:validation:Required
	EndpointName string `json:"endpointName"`

	// EndpointIP is the internal IPv4 address of the pod or VM.
	// +kubebuilder:validation:Required
	EndpointIP string `json:"endpointIP"`

	// Namespace is the Kubernetes namespace of the source endpoint within the cluster
	// specified by ClusterName.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// EgressIP is the external IP used for outbound traffic.
	// +kubebuilder:validation:Required
	EgressIP string `json:"egressIP"`

	// +optional
	// ClusterName specifies the cluster where the endpoint resides.
	// It is used to disambiguate endpoints with the same name
	// across different clusters and namespaces.
	ClusterName string `json:"clusterName,omitempty"`

	// +optional
	// EndpointType identifies the workload category.
	// Valid values: "Pod", "VM", "SystemEP".
	EndpointType EndpointType `json:"endpointType,omitempty"`
}

// CloudNATMappingList contains a list of CloudNATMapping resources.
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CloudNATMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudNATMapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudNATMapping{}, &CloudNATMappingList{})
}
