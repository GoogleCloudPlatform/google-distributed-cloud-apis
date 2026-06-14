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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// Contains the propagation status for a specific cluster.
type ClusterStatus struct {
	corev1alpha1.NamespacedName `json:",inline"`
	Conditions                  []metav1.Condition `json:"conditions,omitempty"`
}

// Provides the overall propagation status.
type PropagationStatus struct {
	// The observations of this resource's current propagation state.
	// The `Propagated` condition is set to `true` if this resource is synced to
	// all target clusters, and pruned from all non-target clusters.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The list of propagation statuses for all clusters to which the resource has
	// been propagated.
	// The `Propagated` condition is set to `true` in the `ClusterPropagationStatus`
	// resource if this resource is synced to the cluster, and its
	// `ObservedGeneration` is set to the generation of the propagated resource in
	// the target cluster. If this resource is successfully pruned from a cluster,
	// its `ClusterPropagationStatus` is removed from the list.
	Clusters []ClusterPropagationStatus `json:"clusters,omitempty"`
}

// Provides the propagation status of a cluster.
type ClusterPropagationStatus struct {
	ClusterStatus `json:",inline"`

	// The namespace of the propagated resource.
	PropagatedNamespace string `json:"propagatedNamespace,omitempty"`
}

func (cs ClusterPropagationStatus) ClusterName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: cs.Namespace,
		Name:      cs.Name,
	}
}
