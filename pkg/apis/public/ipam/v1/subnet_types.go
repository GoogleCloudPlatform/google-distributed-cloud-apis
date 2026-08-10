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

// Represents a type which is used to specify the subnet's position in the IP management tree architecture.
// +kubebuilder:validation:Enum=Root;Branch;Leaf
type SubnetType string

// Represents the type of strategies to sync Subnets between
// global/zonal or root/org API servers.
// +kubebuilder:validation:Enum=None;SingleZone;AllZones
type ZonePropagationStrategy string

// Represents the IP address version.
// +kubebuilder:validation:Enum=IPv4;IPv6
type IPFamily string

// Contains the IP addresses claiming information for a single IP address version (IPv4/IPv6).
type SubnetRequest struct {
	// The CIDR to allocate. This field must be specified when acquiring a dedicated CIDR.
	// +optional
	CIDR *string `json:"cidr,omitempty"`

	// The prefix length of the CIDR wanted. This field can be specified when there is no specific requirements on the CIDR to be allocated. If both CIDR and prefix are left empty, the request acquires a /32(IPv4) or /128(IPv6) random CIDR by default.
	// +optional
	PrefixLength *int32 `json:"prefixLength,omitempty"`
}

// Contains the information used to reference a single `Subnet`.
type SubnetReference struct {
	Name string `json:"name"`

	// The namespace of the referenced Subnet. If it's used in a spec, the namespace can be left empty, which means the referenced Subnet is in the same namespace as the object referencing it.
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	// The type of the Subnet(s) this SubnetReference is referencing to, can be SingleSubnet/SubnetGroup.
	// +kubebuilder:default:="SingleSubnet"
	Type ReferenceType `json:"type,omitempty"`
}

// Represents the type of the Subnets a SubnetReference is referencing to.
// +kubebuilder:validation:Enum=SingleSubnet;SubnetGroup
type ReferenceType string

// Represents a CIDR entry with IP version identified.
type CIDREntry struct {
	// The IP address version of the entry.
	Version IPFamily `json:"version"`

	// The CIDR of the entry.
	CIDR string `json:"cidr"`
}

// Contains a single IP version(IPv4/IPv6)'s CIDR allocation result of a `Subnet`.
type SubnetAllocation struct {
	// The CIDR allocated to the Subnet.
	CIDR string `json:"cidr"`

	// The CIDRs left after excluding those consumed by the Subnet's children.
	AvailableCIDRs []string `json:"availableCIDRs,omitempty"`
}

// Represents the request and allocation information of a zonal `Subnet`.
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Parent",type="string",JSONPath=".spec.parentReference"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="IPv4 CIDR",type="string",JSONPath=".status.ipv4Allocation.cidr"
// +kubebuilder:printcolumn:name="IPv6 CIDR",type="string",JSONPath=".status.ipv6Allocation.cidr"
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubnetSpec   `json:"spec,omitempty"`
	Status SubnetStatus `json:"status,omitempty"`
}

// Defines the specification of a `Subnet`.
type SubnetSpec struct {
	// The type of the Subnet in the context of the IPAM tree architecture.
	// +kubebuilder:default:="Leaf"
	Type SubnetType `json:"type"`

	// The request details for acquiring the IPv4 CIDR.
	// +optional
	IPv4Request *SubnetRequest `json:"ipv4Request,omitempty"`

	// The request details for acquiring the IPv6 CIDR.
	// +optional
	IPv6Request *SubnetRequest `json:"ipv6Request,omitempty"`

	// The reference to the parent of this Subnet. This Subnet will get IP allocated from the parent if specified.
	// The parent name must be provided if there is a parent. Namespace of the parent can only be omitted when
	// the parent is in the same namespace as this Subnet.
	// +optional
	ParentReference *SubnetReference `json:"parentReference,omitempty"`

	// The specification needed for setting up network configurations. If the `Subnet` doesn't have network implications, this field should be empty.
	// +optional
	NetworkSpec *NetworkSpec `json:"networkSpec,omitempty"`
}

// Defines the status of the `Subnet`.
type SubnetStatus struct {
	// The observations of the overall state of the resource.
	// Known condition types: Ready.
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The information of the single parent Subnet from which the CIDRs are allocated from.
	// Empty if the Subnet doesn't have a parent.
	AllocatedParent *SubnetReference `json:"allocatedParent,omitempty"`

	// The allocation information for the IPv4 IP addresses.
	IPv4Allocation *SubnetAllocation `json:"ipv4Allocation,omitempty"`

	// The allocation information for the IPv6 IP addresses.
	IPv6Allocation *SubnetAllocation `json:"ipv6Allocation,omitempty"`

	// The references to the children which are allocated from this subnet. This field doesn't apply to `Leaf` type subnets.
	ChildrenRefs []SubnetReference `json:"childrenRefs,omitempty"`

	// The allocated result of network configurations.
	NetworkStatus *NetworkStatus `json:"networkStatus,omitempty"`

	// The active consumer of the subnet.
	ConsumerRef *ConsumerRef `json:"consumerRef,omitempty"`
}

type ConsumerRef struct {
	Version   string `json:"version"`
	Group     string `json:"group"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// Contains the information to set up network features for the `Subnet`.
type NetworkSpec struct {
	// Specifies whether a gateway IP address must be assigned for the subnet.
	EnableGateway bool `json:"enableGateway"`

	// Specifies whether a VLAN ID must be assigned for the subnet.
	EnableVLANID bool `json:"enableVLANID"`

	// The dedicated VLAN ID. If this field is defined, the `VLANID` field must be `true`.
	// +optional
	StaticVLANID *int32 `json:"staticVLANID,omitempty"`
}

// Contains the allocation result for the network configurations.
type NetworkStatus struct {
	// The VLAN ID acquired for the subnet.
	VLANID *int32 `json:"vlanID,omitempty"`

	// The gateway IP addresses acquired for the Subnet. If the Subnet is single-stack, there must be only one IPv4/IPv6 IP address in the list. If the subnet is dual-stack, there must be one IPv4 address and one IPv6 IP addresse in the list.
	Gateways []CIDREntry `json:"gateways,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a collection of `Subnet` resources.
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&Subnet{},
		&SubnetList{},
	)
}
