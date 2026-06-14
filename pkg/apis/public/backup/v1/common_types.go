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

import corev1 "k8s.io/api/core/v1"

// +kubebuilder:validation:Enum=UserCluster;ManagementAPI
type TargetClusterType string

// Represents a Cluster whose data will be backed up or restored.
type TargetCluster struct {
	// The type of Cluster
	// +kubebuilder:validation:Required
	TargetClusterType TargetClusterType `json:"targetClusterType"`

	// In case of a UserCluster, the name refers to a GDC Cluster inside of the same namespace under the `clusters.cluster.gdc.goog` Group Kind.
	// For the Managemnet API, this field should be left empty.
	// +optional
	TargetClusterName corev1.TypedLocalObjectReference `json:"targetClusterName,omitempty"`
}
