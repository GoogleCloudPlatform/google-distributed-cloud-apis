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
	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// Defines the schema for the `FlowGenerator` API.
type FlowGenerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired configuration for the `FlowGenerator` resource.
	Spec FlowGeneratorSpec `json:"spec,omitempty"`

	// The observed state of the `FlowGenerator` resource.
	Status FlowGeneratorStatus `json:"status,omitempty"`
}

// Defines the configuration of a `FlowGenerator` resource.
//
// Rule 1: Immutability (from previous step).
// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="FlowGenerator specification is immutable and cannot be changed after creation."
//
// Rule 2: If Source is empty or is an IP, Destination must be an IP (cannot be a Resource).
// +kubebuilder:validation:XValidation:rule="(has(self.sourceEndpoint) && has(self.sourceEndpoint.resourceRef)) || !has(self.destinationEndpoint.resourceRef)",message="If sourceEndpoint is empty or is an IP, destinationEndpoint.resourceRef cannot be set (destination must be an IP)."
type FlowGeneratorSpec struct {
	// The source endpoint from where the network flows are generated.
	// +kubebuilder:validation:Optional
	SourceEndpoint FlowEndpointSelector `json:"sourceEndpoint,omitempty"`

	// The destination endpoint that receives the generated network flows.
	// +kubebuilder:validation:Required
	DestinationEndpoint FlowEndpointSelector `json:"destinationEndpoint"`

	// The configuration profile defining the characteristics of the traffic flows generated between the source and destination endpoints.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=1
	FlowProfiles []FlowProfile `json:"flowProfiles"`
}

// Specifies an endpoint for network flows.
// It requires exactly one of `ip` or `resourceRef` to be defined.
// +kubebuilder:validation:XValidation:rule="(has(self.ip) || has(self.resourceRef)) && !(has(self.ip) && has(self.resourceRef))",message="must specify either ip or resourceRef, but not both"
type FlowEndpointSelector struct {

	// An IP address denoting a specific network endpoint.
	// For example, `10.0.0.1`.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty"`

	// A reference to a Kubernetes resource denoting a specific network endpoint.
	// +kubebuilder:validation:Optional
	ResourceRef *ResourceInfo `json:"resourceRef,omitempty"`
}

// Provides detailed information about a specific Kubernetes resource.
// +kubebuilder:validation:XValidation:rule="self.objectRef.kind in ['Pod', 'VirtualMachine', 'Node']", message="Kind must be one of: ['Pod', 'VirtualMachine', 'Node']"
type ResourceInfo struct {
	// The reference identifying a specific Kubernetes object, such as a `Pod`, `VirtualMachine`, or `Node`.
	// +kubebuilder:validation:Required
	ObjectRef corev1.TypedObjectReference `json:"objectRef"`

	// The cluster where the resource is located.
	// If empty, it selects the default cluster.
	// For example, `user-cluster-1`.
	// +kubebuilder:validation:Optional
	ClusterRef *corev1alpha1.NamespacedName `json:"clusterRef,omitempty"`
}

// Defines the network protocol for the flow.
// +kubebuilder:validation:Enum=TCP
type Protocol string

const (
	// Defines the TCP protocol.
	ProtocolTCP Protocol = "TCP"
)

// Defines the characteristics of the network flow.
type FlowProfile struct {
	// The protocol and protocol-specific settings for the network flow.
	// +kubebuilder:validation:Required
	ProtocolConfig ProtocolConfiguration `json:"protocolConfig"`

	// The general flow properties, such as total duration, interval between flows, and total flow count.
	// If empty, default values are applied.
	// +optional
	FlowConfig *FlowConfiguration `json:"flowConfig,omitempty"`
}

// Specifies the protocol and protocol-specific settings for the flow.
// Only one protocol (`tcp`) can be configured at a time.
// +kubebuilder:validation:XValidation:rule="(has(self.tcp) ? 1 : 0) == 1",message="Exactly one protocol configuration (tcp) must be specified."
type ProtocolConfiguration struct {
	// The configuration parameters for TCP traffic flows.
	// +optional
	TCP *TCPConfiguration `json:"tcp,omitempty"`
}

// Defines the configuration for TCP traffic flows.
type TCPConfiguration struct {
	// The source port number for TCP flow generation.
	// The value must be between 1 and 65535, inclusive.
	// For example, `12345`.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	SourcePort *uint16 `json:"sourcePort,omitempty"`

	// The destination port number for receiving TCP flows.
	// The value must be between 1 and 65535, inclusive.
	// If empty, it defaults to 80.
	// For example, `80`.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	DestinationPort *uint16 `json:"destinationPort,omitempty"`
}

// Holds general parameters for generating network flows.
// +kubebuilder:validation:XValidation:rule="!(has(self.durationSeconds) && has(self.flowCount))", message="durationSeconds and flowCount are mutually exclusive; only one can be specified."
type FlowConfiguration struct {
	// The interval between two subsequent flows being generated, in seconds.
	// A value of 0 means flows are generated immediately without an interval.
	// If empty, it defaults to 0.
	// For example, `5`.
	// +kubebuilder:validation:Minimum=0
	// +optional
	IntervalSeconds *uint32 `json:"intervalSeconds,omitempty"`

	// The duration, in seconds, for which flows should be generated.
	// This field cannot be specified if `flowCount` is specified.
	// Since only one of `durationSeconds` and `flowCount` can be specified, if both are unset, the duration defaults to 300 seconds (5 minutes).
	// For example, `60`.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	// +optional
	DurationSeconds *uint32 `json:"durationSeconds,omitempty"`

	// The number of sequential flows to be initiated by the generator.
	// This field cannot be specified if `durationSeconds` is specified.
	// For example, `10`.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=50
	// +optional
	FlowCount *uint32 `json:"flowCount,omitempty"`
}

// Defines a valid condition type for a `FlowGenerator` resource.
type FlowGeneratorConditionType string

const (
	// Indicates that flow generation is pending and has not yet been initiated from the client.
	FlowGeneratorPending FlowGeneratorConditionType = "Pending"

	// Indicates that flows are active and currently being generated.
	FlowGeneratorActive FlowGeneratorConditionType = "Active"

	// Indicates that flows have been generated for the specified duration based on the configuration.
	FlowGeneratorCompleted FlowGeneratorConditionType = "Completed"
)

// Defines the observed state of a `FlowGenerator` resource.
type FlowGeneratorStatus struct {
	// A unique identifier that distinguishes flows generated from a particular `FlowGenerator` resource.
	// The trace ID is injected into an IPv4 Option Header in each generated packet and can be used to trace flows across the network.
	// For example, `12345`.
	// +optional
	TraceID *uint32 `json:"traceID,omitempty"`

	// The latest available observations of the current state of the `FlowGenerator` resource.
	// Known condition types are:
	// * `Pending`: Indicates that flow generation is pending and has not yet been initiated from the client.
	// * `Active`: Indicates that flows are active and currently being generated.
	// * `Completed`: Indicates that flows have been generated for the specified duration based on the configuration.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// Defines a list of `FlowGenerator` resources.
type FlowGeneratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// A list of `FlowGenerator` resources.
	Items []FlowGenerator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlowGenerator{}, &FlowGeneratorList{})
}
