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
	// An annotation to identify the original API server where the IPAM object is created before it's pivoted/propageted.
	CreationOriginAnnotation = "ipam.gdc.goog/creation-origin"

	// An annotation to identify the destination API server where the IPAM object is going to be pivoted to.
	PivotDestinationAnnotation = "ipam.gdc.goog/pivot-destination"

	// An annotation to identify if the IPAM object will be auto divided for each zone or not.
	PauseAutoDivisionAnnotation = "ipam.gdc.goog/pause-auto-division"

	// An annotation to identify the name of cluster for which the IPAM object will be used.
	ClusterNameAnnotation = "ipam.gdc.goog/cluster-name"

	// An annotation to identify the namespace of cluster for which the IPAM object will be used.
	ClusterNamespaceAnnotation = "ipam.gdc.goog/cluster-namespace"

	// An annotation to identify if the IPAM object is delegated through a role for access and use.
	SubnetDelegationRoleAnnotation = "ipam.gdc.goog/subnet-delegation-role"

	// An annotation to identify the creator of the IPAM object.
	SubnetCreatorNameAnnotation = "ipam.gdc.goog.internal/subnet-creator-name"

	// An annotation to identify if a Subnet is locked for pivoting.
	LockedForPivotingAnnotation = "ipam.gdc.goog/locked-for-pivoting"

	// An annotation to identify the timestamp when the source subnet became ready.
	SourceReadyTimestampAnnotation = "ipam.gdc.goog/source-ready-timestamp"
)

// Represents the location of the IPAM object, which are the valid values of CreationOriginAnnotation and PivotDestinationAnnotation
type IPAMObjectLocation string

// Represents how the delegation role for the IPAM object is created, by default it's auto created by the controller.
type IPAMObjectDelegationRole string
