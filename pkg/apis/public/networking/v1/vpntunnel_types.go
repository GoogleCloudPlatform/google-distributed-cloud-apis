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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VPNTunnel represents an encrypted IPSec tunnel between an Organization
// network and a remote network. It connects a VPNGateway interface to a
// PeerGateway interface, and uses VPNBGPPeer to exchange routing information
// over the tunnel.
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="VPNGateway",type="string",JSONPath=".spec.vpnInterface.name"
// +kubebuilder:printcolumn:name="VPNGateway Interface",type="string",JSONPath=".spec.vpnInterface.interface"
// +kubebuilder:printcolumn:name="PeerGateway",type="string",JSONPath=".spec.peerInterface.name"
// +kubebuilder:printcolumn:name="PeerGateway Interface",type="string",JSONPath=".spec.peerInterface.interface"
// +kubebuilder:printcolumn:name="VPNBGPPeer",type="string",JSONPath=".spec.vpnBGPPeer.name"
// +kubebuilder:printcolumn:name="PSK",type="string",JSONPath=".spec.ikeKey.name"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +gdcloud:manifest:relevant=true,oc=unet,component=networking,entities="vpn-tunnels"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:vpn-admin"
// +gdcloud:manifest:rbac="describe,list:vpn-viewer"
// +gdcloud:manifest:skipcodegen=true
type VPNTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VPNTunnelSpec   `json:"spec"`
	Status VPNTunnelStatus `json:"status,omitempty"`
}

// VPNTunnelList contains a list of VPNTunnel.
//
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VPNTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of VPNTunnel entries.
	Items []VPNTunnel `json:"items"`
}

// VPNTunnelSpec defines the desired state of VPNTunnel. The tunnel that
// is established supports the IKEv2 protocol with PSK-based authentication.
// Packets going through the tunnel are encrypted using IPSec Tunnel mode,
// where the outer IP header is constructed using a VPNGateway interface
// IP and a PeerGateway interface IP. A VPNTunnel references a VPNGateway
// interface, a PeerGateway interface, a VPNBGPPeer resource, and a secret
// which contains the preshared key for the authentication.
type VPNTunnelSpec struct {
	// The interface on the VPNGateway that is used for the tunnel. The IP
	// from the interface is used as the source IP for packets sent to the
	// remote site over the tunnel.
	VPNInterface GatewayInterfaceRef `json:"vpnInterface"`

	// The interface on the PeerGateway that is used for the tunnel. The IP
	// from the interface is used as the destination IP for packets sent to
	// the remote site over the tunnel.
	PeerInterface GatewayInterfaceRef `json:"peerInterface"`

	// A reference to a VPNBGPPeer which specifies the dynamic routing over
	// the tunnel.
	VPNBGPPeer corev1.ObjectReference `json:"vpnBGPPeer"`

	// The secret that contains the preshared key for initial authentication
	// of the gateways.
	IKEKey corev1.SecretReference `json:"ikeKey"`
}

// VPNTunnelStatus defines the observed state of VPNTunnel.
type VPNTunnelStatus struct {
	// The current status of the tunnel.
	//
	// +optional
	State *TunnelState `json:"state,omitempty"`

	// Indicates the current status of VPNTunnel. Known condition types are:
	//   - "Ready": The VPNTunnel is reconciled and is in an established state.
	//   - "TunnelEstablished": The tunnel is in an established state.
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// The state of the Tunnel, like Established or Connecting.
//
// +kubebuilder:validation:Enum=Established;Connecting;Unknown
type TunnelState string

const (
	// The tunnel is in an established state.
	TunnelStateEstablished TunnelState = "Established"
	// The tunnel is in a connecting state.
	TunnelStateConnecting TunnelState = "Connecting"
	// The tunnel is in an unknown state.
	TunnelStateUnknown TunnelState = "Unknown"
)

// Represents a reference to an interface on a VPNGateway or PeerGateway resource.
type GatewayInterfaceRef struct {
	// The name of the gateway.
	//
	// +kubebuilder:validation:MinLength:=1
	Name string `json:"name"`

	// The namespace of the gateway.
	//
	// +kubebuilder:validation:MinLength:=1
	Namespace string `json:"namespace"`

	// The name of the interface.
	//
	// +kubebuilder:validation:MinLength:=1
	Interface string `json:"interface"`
}

const (
	// Condition type which indicates that the tunnel is in an establised state.
	TunnelEstablishedConditionType = "TunnelEstablished"

	// Condition type which indicates the status of the reconciliation.
	ReconciledConditionType = "Reconciled"
)

const (
	// Condition reason which indicates that getting a resource with controller runtime client failed.
	GetResourceFailure = "GetResourceFailed"

	TunnelInitFailed = "TunnelInitializationFailed"
)

func init() {
	SchemeBuilder.Register(&VPNTunnel{}, &VPNTunnelList{})
}
