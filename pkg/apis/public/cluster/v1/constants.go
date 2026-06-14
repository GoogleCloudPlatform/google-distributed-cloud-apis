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
	"crypto/sha256"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
)

const (
	// ServiceClusterMaxPodsPerNode indicates the max number of pods per node in a
	// service cluster in Lancer.
	// It is lower than other clusters in order to optimize IP usage in the cluster.
	ServiceClusterMaxPodsPerNode = 32

	// ServiceClusterMaxPodsPerNodeLegacyGCE indicates the max number of pods per node in a
	// GCE/legacy based service cluster.
	ServiceClusterMaxPodsPerNodeLegacyGCE = 64

	// PerimeterClusterMaxPodsPerNode indicates the max number of pods per node
	// in a perimeter cluster.
	PerimeterClusterMaxPodsPerNode = 64

	// PerimeterClusterDataWorkerPoolName indicates the name of the NodePool
	// containing perimeter cluster data vrf nodes.
	PerimeterClusterDataWorkerPoolName = "perimeter-data-node-pool"

	// PerimeterClusterAdminWorkerPoolName indicates the name of the NodePool
	// containing perimeter cluster admin vrf nodes.
	PerimeterClusterAdminWorkerPoolName = "perimeter-admin-node-pool"

	// InternalMKSClusterNamespace is the namespace storing internal MKS cluster
	// API objects.
	InternalMKSClusterNamespace = "mks-system"

	// SharedServiceClusterDefaultWorkerName is the name of the shared service
	// cluster default CPU worker NodePool.
	SharedServiceClusterDefaultWorkerName = "shared-service-default-worker"

	// SharedServiceClusterDefaultGPUWorkerName is the name of the shared service
	// cluster default GPU worker NodePool.
	SharedServiceClusterDefaultGPUWorkerName = "shared-service-default-gpu-worker"

	// SharedServiceClusterDefaultWorkerCount is the shared service cluster
	// default CPU worker NodePool size.
	SharedServiceClusterDefaultWorkerCount = 2

	// SharedServiceClusterPodCIDRMaskSize is pod CIDR mask size of the shared
	// service cluster.
	SharedServiceClusterPodCIDRMaskSize = 17

	// SharedServiceClusterPodCIDRMaskSizeLegacy is pod CIDR mask size of the shared
	// service cluster in a legacy environment.
	SharedServiceClusterPodCIDRMaskSizeLegacy = 19

	// SharedServiceClusterServiceCIDRMaskSize is service CIDR mask size of the
	// shared service cluster.
	SharedServiceClusterServiceCIDRMaskSize = 23

	// GPUWorkerTaintKey is the taint key of the GPU worker nodes.
	GPUWorkerTaintKey = "cluster.gdc.goog/gpu-worker"

	// DedicatedNodeTaintKeyPrefix is the taint key for dedicated nodes.
	DedicatedNodeTaintKeyPrefix = "cluster.gdc.goog/dedicated-node-"
)

// Vanilla cluster constants
const (
	// truncateWidth is the truncate width for the shadow project hash.
	truncateWidth = 8
)

// SharedServiceClusterNameFor is a helper function to generate a namespaced name
// for the shared service cluster object of an org in the infra/org admin cluster
// The shared service cluster may be deleted by IO once an org has
// been initialized, so it is not guaranteed that the cluster object will always
// exist in the infra/org admin cluster.
func SharedServiceClusterNameFor(name string) types.NamespacedName {
	clusterName := "g-" + name + "-shared-service"
	return types.NamespacedName{
		Namespace: InternalMKSClusterNamespace,
		Name:      clusterName,
	}
}

// PerimeterClusterNameFor is a helper function to generate a namespaced name
// for the perimeter cluster object of an org in the infra cluster
func PerimeterClusterNameFor(name string) types.NamespacedName {
	clusterName := "g-" + name + "-perimeter"
	return types.NamespacedName{
		Namespace: InternalMKSClusterNamespace,
		Name:      clusterName,
	}
}

// VanillaClusterNameHashFor returns a hash determined by the given cluster name and the project name.
func VanillaClusterNameHashFor(clusterName, projectName string) string {
	h := sha256.New()
	h.Write([]byte(clusterName))
	h.Write([]byte(projectName))
	bs := h.Sum(nil)
	hash := fmt.Sprintf("%x", bs)
	return hash[:truncateWidth]
}

// VanillaClusterNamespaceFor returns a namespace by appending mks-system
// to the project name and the vanilla cluster name hash.
func VanillaClusterNamespaceFor(clusterName, projectName string) string {
	return VanillaClusterUniqueNameFor(clusterName, projectName) + "-mks-system"
}

// VanillaClusterUniqueNameFor returns a unique identifier for a vanilla
// cluster with its AO project and vanilla cluster name hash.
func VanillaClusterUniqueNameFor(clusterName, projectName string) string {
	return projectName + "-" + VanillaClusterNameHashFor(clusterName, projectName)
}

// Cluster types constants copied locally under public cluster namespace
// to remove dependency on internal package private/lcm/v1.
const (
	UserClusterType      = "user"
	VanillaClusterType   = "vanilla"
	ServiceClusterType   = "service"
	PerimeterClusterType = "perimeter"
)
