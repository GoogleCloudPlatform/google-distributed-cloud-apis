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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=pb
// +kubebuilder:storageversion
// +gdcloud:manifest:relevant=true,oc=rm,component=resource-manager,entities="project-bindings"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create:project-creator,user-cluster-admin,standard-cluster-admin"
// +gdcloud:manifest:rbac="update:project-editor"
// +gdcloud:manifest:rbac="delete,describe,list:user-cluster-admin,project-editor,standard-cluster-admin"
// +gdcloud:manifest:rbac="describe,list:standard-cluster-viewer"
// +gdcloud:manifest:skipcodegen=true
// Represents a cluster resource that maintains the mapping relations between
// clusters and projects. The namespace of the `ProjectBinding` object
// corresponds to the cluster.
// +genclient
type ProjectBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectBindingSpec `json:"spec,omitempty"`
}

// Provides the specification, or desired state, of a `ProjectBinding` resource.
type ProjectBindingSpec struct {
	ClusterRef ProjectBindingClusterRef `json:"clusterRef"`
	// The Selector is used to specify a set of rules to match Projects.
	Selector ProjectBindingSelector `json:"selector,omitempty"`
}

// Represents the cluster that projects propagate to.
type ProjectBindingClusterRef struct {
	// The cluster name.
	Name string `json:"name"`
}

// Provides a set of rules to match Projects.
// Must choose exactly 0 or 1 of the selectors.
// 0 selector matches all Projects.
// +kubebuilder:validation:MaxProperties=1
type ProjectBindingSelector struct {
	NameSelector  *NameSelector         `json:"nameSelector,omitempty"`
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
}

// Provides a list of Project Name For ProjectBinding to match with.
type NameSelector struct {
	MatchNames []string `json:"matchNames,omitempty"`
}

// Matches checks if the given Project matches the ProjectBindingSelector.
func (s *ProjectBindingSelector) Matches(project *Project) (bool, error) {
	switch {
	case s.NameSelector != nil:
		for _, name := range s.NameSelector.MatchNames {
			if name == project.Name {
				return true, nil
			}
		}
		return false, nil

	case s.LabelSelector != nil:
		labelSelector, err := metav1.LabelSelectorAsSelector(s.LabelSelector)
		if err != nil {
			return false, fmt.Errorf("failed to parse LabelSelector: %w", err)
		}

		return labelSelector.Matches(labels.Set(project.GetLabels())), nil

	default: // matches all Projects if non selector is set.
		return true, nil
	}
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
