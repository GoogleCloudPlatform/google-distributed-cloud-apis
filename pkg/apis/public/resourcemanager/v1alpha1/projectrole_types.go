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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// Defines the desired state of a `ProjectRole` resource. It is the same
// definition as a native Kubernetes `Role`.
type ProjectRoleSpec struct {
	Rules []rbacv1.PolicyRule `json:"rules,omitempty"`
}

// Defines the observed state of a `ProjectRole` resource.
type ProjectRoleStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The list of propagation statuses on the clusters.
	Clusters []ClusterStatus `json:"clusters,omitempty"`

	// The name of the propagated `ProjectRole` resource realized in the user
	// clusters.
	PropagatedName string `json:"propagatedName,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pr
// +gdcloud:manifest:relevant=false,oc=iam
// +genclient
// Represents a project resource that propagates the `Role` configuration to all
// user clusters the project spans across. The namespace of the `ProjectRole`
// resource corresponds to the project.
type ProjectRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectRoleSpec   `json:"spec,omitempty"`
	Status ProjectRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// Contains a list of `ProjectRole` resources.
type ProjectRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProjectRole{}, &ProjectRoleList{})
}
