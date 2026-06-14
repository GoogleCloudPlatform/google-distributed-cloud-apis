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
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=pb

// +genclient
// Represents a cluster resource that maintains the mapping relations between
// clusters and projects. The namespace of the `ProjectBinding` object
// corresponds to the cluster.
type ProjectBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectBindingSpec `json:"spec,omitempty"`
}

// Provides the specification, or desired state, of a `ProjectBinding` resource.
type ProjectBindingSpec struct {
	ClusterRef ClusterRef `json:"clusterRef"`
	Selector   Selector   `json:"selector"`
}

// Represents the cluster that projects propagate to.
type ClusterRef struct {
	// The cluster name.
	Name string `json:"name"`
}

// Matches a set of projects.
type Selector struct {
	// The list of project names to match. An empty slice matches all projects.
	MatchNames []string `json:"matchNames"`
}

// +kubebuilder:object:root=true

// Contains a list of `ProjectBinding` resources.
type ProjectBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ProjectBinding{},
		&ProjectBindingList{},
	)
}
