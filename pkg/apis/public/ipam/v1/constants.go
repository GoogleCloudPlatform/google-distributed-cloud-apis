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

const (
	// Defines the subnet as the root IP address range of all the children/grandchildren in the IPAM architecture. Root type subnets must have a CIDR specified and cannot have a parent.
	Root SubnetType = "Root"

	// Defines the subnet as having both a parent and children. The subnet must have a parent; children are optional.
	Branch SubnetType = "Branch"

	// Defines the subnet as being consumed directly by end consumers of IPAM, such as VMs and load balancers. Such subnets cannot have children and optionally have parents. It is the default subnet `type`.
	Leaf SubnetType = "Leaf"
)

const (
	// Restricts propagation on the Subnet, which is the default value.
	None ZonePropagationStrategy = "None"

	// Allows the subnet to be propagated to the zonal API server of the zone specified in the `zone` field.
	SingleZone ZonePropagationStrategy = "SingleZone"

	// Allows the subnet to be propagated to all the zones, which is usually used by anycast IP addresses.
	AllZones ZonePropagationStrategy = "AllZones"
)

const (
	// Represents the Internet Protocol(IP) version v4.
	IPv4 IPFamily = "IPv4"

	// Represents the Internet Protocol(IP) version v6.
	IPv6 IPFamily = "IPv6"
)

const (
	// Represents the referenced target is a single Subnet.
	SingleSubnet ReferenceType = "SingleSubnet"
	// Represents the referenced target is a SubnetGroup.
	SubnetGroup ReferenceType = "SubnetGroup"
)

const (
	// Represents the zonal API server of the root org.
	ZonalRoot IPAMObjectLocation = "zonal-root"

	// Represents the global API server of the root org.
	GlobalRoot IPAMObjectLocation = "global-root"

	// Represents the corresponding global org API server.
	GlobalOrg IPAMObjectLocation = "global-org"

	// Represents the corresponding zonal org API server.
	ZonalOrg IPAMObjectLocation = "zonal-org"
)

const (
	// Represents the IPAM objects are the root IP range of a VPC/VRF network.
	NetworkRootRange UsageType = "network-root-range"

	// Represents the IPAM objects are the root IP range of a VPC/VRF network for a zone.
	ZoneNetworkRootRange UsageType = "zone-network-root-range"

	// Represents the IPAM objects are the global anycast IP range of an org across all zones.
	GlobalAnycastRootRange UsageType = "global-anycast-root-range"
)

const (
	Admin NetworkSegment = "admin"
	Data  NetworkSegment = "data"
)

const Default AllocationPreference = "default"

const Auto IPAMObjectDelegationRole = "auto"

const (
	// Represents the root Subnet group name of infra VPC.
	InfraVPCRootSubnetGroup = "infra-vpc-root-group"

	// Represents the root Subnet group name of default VPC.
	DefaultVPCRootSubnetGroup = "default-vpc-root-group"

	// Represents the root Subnet group name of admin network segment.
	AdminNetworkSegmentRootSubnetGroup = "admin-network-segment-root-group"

	// Represents the root Subnet group name of data network segment.
	DataNetworkSegmentRootSubnetGroup = "data-network-segment-root-group"
)

const (
	// Represents the anycast Subnet group name of infra VPC.
	InfraVPCAnycastSubnetGroup = "infra-vpc-anycast-group"

	// Represents the global only IP group name of default VPC.
	DefaultVPCGlobalIPGroup = "default-vpc-global-ip-group"

	// Represents the anycast Subnet group name of admin network segment.
	AdminNetworkSegmentAnycastSubnetGroup = "admin-network-segment-anycast-group"

	// Represents the anycast Subnet group name of data network segment.
	DataNetworkSegmentAnycastSubnetGroup = "data-network-segment-anycast-group"
)
