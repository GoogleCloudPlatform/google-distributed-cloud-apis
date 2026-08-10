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

const (
	// An annotation that can be applied to the cluster objects.
	// Cluster controllers will not reconcile objects with this annotation.
	PausedAnnotation = "cluster.gdc.goog/paused"

	// An annotation that can be applied to the cluster objects so they use the
	// specified machine type for the cluster control plane nodes.
	ControlPlaneMachineTypeAnnotation = "cluster.gdc.goog/control-plane-machine-type"

	// An annotation that can be applied to the cluster objects so they use the
	// specified number for the cluster control plane node count.
	ControlPlaneNodeCountAnnotation = "cluster.gdc.goog/control-plane-node-count"

	// An annotation to add the GDC cluster namespace on the ABM cluster
	// to have a mapping from ABM cluster back to the parent GDC cluster
	ClusterNamespaceAnnotation = "cluster.gdc.goog/gdc-cluster-namespace"
)

// Returns `true` if the object has the `paused` annotation.
func HasPaused(o metav1.ObjectMeta) bool {
	if metav1.HasAnnotation(o, PausedAnnotation) {
		return o.Annotations[PausedAnnotation] == "true"
	}
	return false
}
