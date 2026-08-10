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

// PeerGateway represents a remote VPN endpoint. An interface on a PeerGateway
// should be used by a single VPNTunnel to establish an encrypted tunnel to the
// remote site.
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Interface0 Name",type="string",JSONPath=".spec.interfaces[0].name"
// +kubebuilder:printcolumn:name="Interface0 IP",type="string",JSONPath=".spec.interfaces[0].ip"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
type PeerGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PeerGatewaySpec   `json:"spec"`
	Status PeerGatewayStatus `json:"status,omitempty"`
}

// PeerGatewayList contains a list of PeerGateway.
//
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PeerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of PeerGateway entries.
	Items []PeerGateway `json:"items"`
}

// PeerGatewaySpec defines the desired state of PeerGateway.
type PeerGatewaySpec struct {
	// The list of interfaces on the Peer Gateway which will
	// be used for VPN connections. Each interface should be
	// used by one VPNTunnel.
	//
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:MaxItems:=1
	Interfaces []PeerGatewayInterface `json:"interfaces"`
}

// PeerGatewayStatus defines the observed state of PeerGateway.
type PeerGatewayStatus struct {
	// Indicates the current status of PeerGateway. Known condition types are:
	//   - "Ready": The Peer Gateway is reconciled and used by a VPNTunnel.
	//   - "TunnelsEstablished": Each interface on the Peer Gateway is used
	//     by a VPNTunnel.
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// Represents a Gateway interface on which a tunnel is established.
type PeerGatewayInterface struct {
	// The name of the interface.
	//
	// +kubebuilder:validation:MinLength:=1
	Name string `json:"name"`

	// The IPv4 address of the interface.
	//
	// +kubebuilder:validation:MinLength:=1
	IP string `json:"ip"`
}

const (
	// Condition type which indicates that the IPv4 addresses that have been
	// assigned to each interface on the PeerGateway do not intersect with
	// internal IPs.
	ValidGatewayIPsConditionType = "ValidGatewayIPs"
)

const (
	// Reason for ValidGatewayIP Type when an IP is invalid.
	InvalidIP = "InvalidIP"

	// Reason for ValidGatewayIP Type when IP's conflict with internal Organization CIDRs.
	IPConflictExists = "IPConflictExists"

	// Reason for ValidGatewayIP Type when no IP's conflict with internal Organization CIDRs.
	NoIPConflictExists = "NoIPConflictExists"
)

func init() {
	SchemeBuilder.Register(&PeerGateway{}, &PeerGatewayList{})
}
