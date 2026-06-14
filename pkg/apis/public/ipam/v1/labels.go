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
	// A label which describes the usage purpose of a IPAM object.
	UsageLabel = "ipam.gdc.goog/usage"

	// A label which identifies the cluster which the IPAM object will be used for.
	ClusterLabel = "ipam.gdc.goog/cluster"

	// A label which identifies the VPC which the IPAM object is assigned to.
	VPCLabel = "ipam.gdc.goog/vpc"

	// A label which identifies the network segment which the IPAM object belongs to.
	NetworkSegmentLabel = "ipam.gdc.goog/network-segment"

	// A label which identifies the allocation preference of the IPAM object.
	AllocationPreferenceLabel = "ipam.gdc.goog/allocation-preference"

	// A label which identifies the SubnetGroup which the labelled Subnet belongs to.
	SubnetGroupLabel = "ipam.gdc.goog/subnet-group"
)

// Represents the usage of an IPAM object, which are the IPAM defined values of the UsageLabel.
// More values can be defined by the IPAM system users.
type UsageType string

type NetworkSegment string

type AllocationPreference string
