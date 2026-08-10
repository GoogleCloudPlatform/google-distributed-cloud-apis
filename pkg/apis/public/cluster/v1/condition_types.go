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
	// Indicates whether the cluster namespace is created.
	//
	// Deprecated: use ClusterNamespaceReady instead.
	ClusterNamespaceCreated = "NamespaceCreated"
	// Indicates whether the cluster namespace is ready.
	ClusterNamespaceReady = "NamespaceReady"
	// Indicates whether the cluster's network resources, such as IP address pools
	// or CIDR blocks, are ready.
	ClusterNetworkResourcesReady = "NetworkResourcesReady"
	// Indicates whether the internal cluster object is created or updated.
	//
	// Deprecated: This condition is obsolete and it's incorporated into the reconciling condition.
	ClusterDeployed = "ClusterDeployed"
	// Indicates whether the cluster's node pools are created or updated.
	//
	// Deprecated: This condition is obsolete and it's incorporated into the reconciling condition.
	ClusterNodePoolsDeployed = "NodePoolsDeployed"
	// Indicates whether the nodes with accelerators have been configured for GPUAllocation reconciling.
	//
	// Deprecated: This condition is obsolete and it's incorporated into the reconciling condition.
	ClusterAcceleratorNodePoolsConfigured = "AcceleratorNodePoolsConfigured"
	// Indicates whether the sshless mode has been enabled or not. If yes, also indicates whether it has been configured or not.
	SSHlessModeEnabled = "SSHlessModeEnabled"
	// Indicates whether all critical system components are ready.
	CriticalComponentReady = "CriticalComponentReady"

	// Provides the condition type that indicates the overall readiness of
	// the cluster.
	ReadyCondition = "Ready"
	// Provides the condition type that indicates whether the cluster is
	// reconciling or not.
	ReconcilingCondition = "Reconciling"
	// Provides the condition type that indicates whether the cluster is
	// in a stalled state or not.
	//
	// Deprecated: This condition is obsolete and it's incorporated into the reconciling condition.
	StalledCondition = "Stalled"
	// Provides the condition type that indicates whether the GPU allocations of the
	// node pool is in a ready state or not.
	GPUAllocationReadyCondition = "GPUAllocationReady"
	// Provides the condition type that indicates whether the virtual machines of the
	// node pool are in a ready state or not.
	VirtualMachineReadyCondition = "VirtualMachineReady"
)
