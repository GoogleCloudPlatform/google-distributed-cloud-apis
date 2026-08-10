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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=rm,entities="projects"
// +gdcloud:manifest:verbs=create;delete;update
// +gdcloud:manifest:rbac="create:project-creator"
// +gdcloud:manifest:rbac="delete,update:project-editor"
// +gdcloud:manifest:skipcodegen=true
// Represents a namespace that spans across multiple user clusters in an
// organization. It is a namespaced resource, and the controller is expected to
// watch reconcile `Project` objects in a preconfigured namespace.
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of projects.
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Project `json:"items"`
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

// Hub marks this Project version as a conversion Hub.
func (*Project) Hub() {}
