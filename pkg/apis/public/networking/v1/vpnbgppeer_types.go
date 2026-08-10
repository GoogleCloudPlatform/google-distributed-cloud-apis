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

// VPNBGPPeer represents a BGP session over a VPN tunnel. A VPNBGPPeer
// establishes a BGP session between a BGP peer in an Organization and BGP peer
// of a remote site across a single VPNTunnel. A VPNBGPPeer should be used by a
// VPNTunnel.
//
// +genclient
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Local IP",type="string",JSONPath=".spec.local.ip"
// +kubebuilder:printcolumn:name="Local ASN",type="string",JSONPath=".spec.local.asn"
// +kubebuilder:printcolumn:name="Remote IP",type="string",JSONPath=".spec.remote.ip"
// +kubebuilder:printcolumn:name="Remote ASN",type="string",JSONPath=".spec.remote.asn"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
type VPNBGPPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPNBGPPeerSpec   `json:"spec"`
	Status VPNBGPPeerStatus `json:"status,omitempty"`
}

// VPNBGPPeerList contains a list of VPNBGPPeer.
//
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VPNBGPPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of VPNBGPPeer entries.
	Items []VPNBGPPeer `json:"items"`
}

// VPNBGPPeerSpec defines the desired state of VPNBGPPeer. The IP for both the
// Remote and Local BGP peer must be in the same "/30" block from the
// "169.254.0.0/16" range.
type VPNBGPPeerSpec struct {
	// Represents the remote BGP peer.
	Remote BGPPeerInterface `json:"remote"`

	// Represents the local BGP peer.
	Local BGPPeerInterface `json:"local"`
}

// VPNBGPPeerStatus defines the observed state of VPNBGPPeer. The Organization
// will create a BGP session with the remote site over the VPNTunnel that
// references this VPNBGPPeer. The local BGP peer will advertise all internal
// CIDR's of the Organization to the remote BGP peer. If the remote BGP peer
// advertises a CIDR that conflicts with the internal CIDR's of the Organization,
// the VPNBGPPeer "Ready" condition will be false.
type VPNBGPPeerStatus struct {
	// Represents the state of the BGP session between the local BGP peer
	// and the remote BGP peer.
	//
	// +optional
	State *SessionState `json:"state,omitempty"`

	// Represents routes advertised to the remote site on the BGP session.
	//
	// +optional
	Advertised []Route `json:"advertised,omitempty"`

	// Represents routes received from the remote site on the BGP session.
	//
	// +optional
	Received []Route `json:"received,omitempty"`

	// Indicates the current status of VPNBGPPeer. Known condition types are:
	//   - "Ready": The Peer Gateway is reconciled and used by a VPNTunnel.
	//   - "BGPSessionEstablished": Each interface on the Peer Gateway is used
	//     by a VPNTunnel.
	//   - "ReceivedRoutesReady": The routes received from the remote BGP peer
	//     do not interfere with routes in the internal CIDR of the Organization.
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// Represents a route advertised or received over the tunnel.
type Route struct {
	// The network prefix of the route.
	//
	// +optional
	Prefix string `json:"prefix,omitempty"`
}

// Represents a reference to a BGP peer.
type BGPPeerInterface struct {
	// The name of the BGP peer.
	//
	// +kubebuilder:validation:MinLength:=1
	Name string `json:"name"`

	// The IP address of the BGP peer.
	//
	// +kubebuilder:validation:MinLength:=1
	IP string `json:"ip"`

	// The Autonomous System Number.
	ASN uint32 `json:"asn"`
}

// The state of the BGP session, like Established or NotEstablished.
//
// +kubebuilder:validation:Enum=Established;NotEstablished;Unknown
type SessionState string

const (
	// The session is in an established state.
	SessionStateEstablished SessionState = "Established"
	// The session is not in an established state.
	SessionStateNotEstablished SessionState = "NotEstablished"
	// The session is in an unknown state.
	SessionStateUnknown SessionState = "Unknown"
)

const (
	// Condition type which indicates that the BGP session is in an establised state.
	BGPSessionEstablishedConditionType = "BGPSessionEstablished"
	// Condition type which indicates that the received routes are valid.
	ReceivedRoutesValidConditionType = "ReceivedRoutesValid"
	// Condition type which indicates that the received routes are in a ready state.
	ReceivedRoutesReadyConditionType = "ReceivedRoutesReady"
	// Condition type which indicates that the IP for both the local and remote BGP IP are in a ready (valid) state.
	ValidIPsConditionType = "ValidIPs"
	// Condition type which indicates that route advertisement is in a ready state.
	AdvertisedRoutesReadyConditionType = "AdvertisedRoutesReady"
)

const (
	// Reason for ReceivedRoutesValid Condition if there is a prefix conflict.
	PrefixConflict = "PrefixConflict"
)

func init() {
	SchemeBuilder.Register(&VPNBGPPeer{}, &VPNBGPPeerList{})
}
