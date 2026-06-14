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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=rm
// Represents a namespace that spans across multiple user clusters in an
// organization. It is a namespaced resource, and the controller is expected to
// watch reconcile `Project` objects in a preconfigured namespace.
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// Represents a collection of projects.
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Project `json:"items"`
}

// Provides the specification, or desired state, of a project.
type ProjectSpec struct {
	// ClusterSelector is deprecated. Use ProjectBinding API instead.
	ClusterSelector *ClusterSelector `json:"clusterSelector,omitempty"`
}

// Matches a set of user clusters. An empty selector matches no user clusters. A
// `nil` selector matches all user clusters.
type ClusterSelector struct {
	// MatchNames holds list of namespaced name of user clusters to match.
	MatchNames []corev1alpha1.NamespacedName `json:"matchNames,omitempty"`
}

// Provides the status of a project.
type ProjectStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The name of the propagated namespace.
	PropagatedName string `json:"propagatedName,omitempty"`

	// The propagation statuses and egress NAT IP addresses of all user clusters
	// this project spans across.
	Clusters []ProjectClusterStatus `json:"clusters,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// Contains the propagation status and egress NAT IP address used for a
// specific cluster.
type ProjectClusterStatus struct {
	ClusterStatus      `json:",inline"`
	EgressNATIPAddress *string `json:"egressNATIPAddress,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&Project{},
		&ProjectList{},
	)
}
