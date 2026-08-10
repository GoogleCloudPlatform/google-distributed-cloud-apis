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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
	rmv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/resourcemanager/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Represents a project resource that propagates the `ClusterRoleBinding`
// resource configuration to all vanilla clusters in the same project.
// The namespace for the `StandardClusterRoleBinding` resource
// corresponds to the project.
// +genclient
// +gdcloud:manifest:relevant=false,oc=iam
type StandardClusterRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StandardClusterRoleBindingSpec   `json:"spec,omitempty"`
	Status StandardClusterRoleBindingStatus `json:"status,omitempty"`
}

// Defines the specification of the `StandardClusterRoleBinding` resource.
// It is the same definition as a native `ClusterRoleBinding` definition.
type StandardClusterRoleBindingSpec struct {
	// The subjects of the `RoleBinding` resource created in the cluster.
	Subjects []rbacv1.Subject `json:"subjects,omitempty"`

	// The `RoleRef` resource of the `RoleBinding` object to create in the
	// cluster.
	RoleRef rbacv1.RoleRef `json:"roleRef,omitempty"`
}

// Defines the observed state of the `StandardClusterRoleBinding`
// resource.
type StandardClusterRoleBindingStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The list of propagation statuses for the clusters.
	Clusters []rmv1.ClusterStatus `json:"clusters,omitempty"`

	// The name of the propagated `ClusterRoleBinding` resource realized
	// in the vanilla clusters.
	PropagatedName string `json:"propagatedName,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `StandardClusterRoleBinding` resources.
type StandardClusterRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StandardClusterRoleBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StandardClusterRoleBinding{}, &StandardClusterRoleBindingList{})
}
