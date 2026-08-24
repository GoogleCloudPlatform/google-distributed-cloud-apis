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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// Defines the specification of the `ProjectRoleBinding` resource. It is the
// same definition as a native `RoleBinding` definition.
type ProjectRoleBindingSpec struct {
	// The subjects of the `RoleBinding` resource created in the cluster.
	Subjects []rbacv1.Subject `json:"subjects,omitempty"`

	// The `RoleRef` resource of the `RoleBinding` object to create in the
	// cluster.
	RoleRef rbacv1.RoleRef `json:"roleRef,omitempty"`
}

// Defines the observed state of the `ProjectRoleBinding` resource.
type ProjectRoleBindingStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The list of propagation statuses for the clusters.
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
// +kubebuilder:resource:shortName=prb
// +gdcloud:manifest:relevant=false,oc=iam
// +genclient
// Represents a project resource that propagates the `RoleBinding` resource
// configuration to all user clusters the project spans across. The namespace
// for the `ProjectRoleBinding` resource corresponds to the project.
type ProjectRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectRoleBindingSpec   `json:"spec,omitempty"`
	Status ProjectRoleBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// Contains a list of `ProjectRoleBinding` resources.
type ProjectRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectRoleBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProjectRoleBinding{}, &ProjectRoleBindingList{})
}
