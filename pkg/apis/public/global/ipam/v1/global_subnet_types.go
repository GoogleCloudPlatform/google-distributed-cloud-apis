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

	ipamv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/ipam/v1"
)

// Represents the request and allocation information of a global IP address range in CIDR format.
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
	Type ipamv1.SubnetType `json:"type"`

	// The request details for acquiring the IPv4 CIDR.
	// +optional
	IPv4Request *ipamv1.SubnetRequest `json:"ipv4Request,omitempty"`

	// The request details for acquiring the IPv6 CIDR.
	// +optional
	IPv6Request *ipamv1.SubnetRequest `json:"ipv6Request,omitempty"`

	// The name of the zone which this Subnet belongs to. If left empty, it is treated as a global only resource.
	// +optional
	Zone *string `json:"zone,omitempty"`

	// The reference to the parent of this Subnet. This Subnet will get IP allocated from the parent if specified.
	// The parent name must be provided if there is a parent. Namespace of the parent can only be omitted when
	// the parent is in the same namespace as this Subnet.
	// +optional
	ParentReference *ipamv1.SubnetReference `json:"parentReference,omitempty"`

	// The strategy used to propagate the Subnet to zonal API server.
	// +kubebuilder:default:="None"
	PropagationStrategy ipamv1.ZonePropagationStrategy `json:"propagationStrategy,omitempty"`
}

// Represents the status of the `Subnet`.
type SubnetStatus struct {
	// The observations of the overall state of the resource.
	// Known condition types: Ready.
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The information of the single parent Subnet from which the CIDRs are allocated from.
	// Empty if the Subnet doesn't have a parent.
	AllocatedParent *ipamv1.SubnetReference `json:"allocatedParent,omitempty"`

	// The allocation information for the IPv4 IP addresses.
	IPv4Allocation *ipamv1.SubnetAllocation `json:"ipv4Allocation,omitempty"`

	// The allocation information for the IPv6 IP addresses.
	IPv6Allocation *ipamv1.SubnetAllocation `json:"ipv6Allocation,omitempty"`

	// The references to the children which are allocated from this Subnet. This field doesn't apply to Leaf type Subnets.
	ChildrenRefs []ipamv1.SubnetReference `json:"childrenRefs,omitempty"`

	// The propagation status of each zone that the Subnet propagated to. This field doesn't apply to Subnet without PropagationStrategy.
	ZonePropagations []PropagationStatus `json:"zonePropagations,omitempty"`

	// The active consumer of the subnet.
	ConsumerRef *ipamv1.ConsumerRef `json:"consumerRef,omitempty"`
}

type PropagationStatus struct {
	// The name of the zone in which the Subnet is propagated to.
	Zone string `json:"zone,omitempty"`

	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
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
