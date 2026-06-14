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

// VPNGateway represents an Organization VPN endpoint. An interface on a
// VPNGateway should be used by a VPNTunnel to establish an encrypted tunnel
// to a remote site.
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Interface0 Name",type="string",JSONPath=".spec.interfaces[0].name"
// +kubebuilder:printcolumn:name="Interface0 IP",type="string",JSONPath=".status.interfaces[0].ip"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
type VPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPNGatewaySpec   `json:"spec"`
	Status VPNGatewayStatus `json:"status,omitempty"`
}

// VPNGatewayList contains a list of VPNGateway.
//
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of VPNGateway entries.
	Items []VPNGateway `json:"items"`
}

// VPNGatewaySpec defines the desired state of VPNGateway. Define the name
// of the interfaces which should be assigned external IPv4 addresses by the
// VPNGateway. Each IPv4 address that is assigned to the interface is
// automatically chosen from an external address pool.
type VPNGatewaySpec struct {
	// The names of each interface. VPNGateway will assign an external IPv4
	// address to each interface in the list.
	//
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:MaxItems:=1
	Interfaces []VPNGatewayInterface `json:"interfaces"`
}

// VPNGatewayStatus defines the observed state of VPNGateway.
type VPNGatewayStatus struct {
	// The list of interfaces on the VPNGateway. Each interface can be
	// used by one VPNTunnel.
	//
	// +optional
	Interfaces []VPNGatewayInterfaceStatus `json:"interfaces,omitempty"`

	// Subnets holds the list of subnets that are used for this VPNGateway.
	//
	// +optional
	Subnets []SubnetReference `json:"subnets,omitempty"`

	// Indicates the current status of VPNGateway. Known condition types are:
	//   - "Ready": The VPNGateway is reconciled and used by a VPNTunnel.
	//   - "IPsAssigned": IPv4 addresses have been assigned to each interface
	//     on the VPNGateway.
	//   - "TunnelsAttached": Each interface on the VPNGateway is used
	//     by a VPNTunnel.
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// Represents an interface on the Gateway. An IP will be allocated
// for each interface and added in the status.
type VPNGatewayInterface struct {
	// The name of the interface.
	//
	// +kubebuilder:validation:MinLength:=1
	Name string `json:"name"`
}

// Represents a Gateway interface on which a tunnel is established.
type VPNGatewayInterfaceStatus struct {
	// The name of the interface.
	Name string `json:"name"`

	// The IPv4 address of the interface.
	IP string `json:"ip"`
}

const (
	// Condition type which indicates that IPv4 addresses have been assigned to
	// each interface on the VPNGateway.
	IPsAssignedConditionType = "IPsAssigned"
	// Condition type which indicates that each interface on the VPNGateway
	// is used by a VPNTunnel.
	TunnelsAttachedConditionType = "TunnelsAttached"
)

const (
	// Reason for IPsAssigned Type when IP's were not allocated.
	IPAllocationFailed = "IPAllocationFailed"
	// Reason for IPsAssigned Type when IP's were not assigned.
	IPAssignmentFailed = "IPAssignmentFailed"
	// Reason for TunnelsEstablished Type when at least one interface is not attached to a VPNTunnel.
	NoTunnelAttached = "NoTunnelAttached"
	// General reason for when a condition can't be resolved.
	Unknown = "Unknown"
	// Reason for Ready Type when a runtime client error occurred.
	ClientError = "ClientError"
)

func init() {
	SchemeBuilder.Register(&VPNGateway{}, &VPNGatewayList{})
}
