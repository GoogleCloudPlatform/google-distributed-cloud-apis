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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// IAMRoleBinding references a global IAMRole and adds who information via Subject.
// +genclient
type IAMRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMRoleBindingSpec   `json:"spec,omitempty"`
	Status IAMRoleBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of IAMRoleBinding resources.
type IAMRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMRoleBinding `json:"items"`
}

// Provides the specification of the IAMRoleBindingSpec.
type IAMRoleBindingSpec struct {
	// RoleRef contains information that points to the IAMRole being used.
	RoleRef rbacv1.RoleRef `json:"roleRef,omitempty"`

	// The subjects of the global IAMRoleBinding resource.
	Subjects []rbacv1.Subject `json:"subjects,omitempty"`
}

// Provides the status of the IAMRoleBinding.
type IAMRoleBindingStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&IAMRoleBinding{}, &IAMRoleBindingList{})
}
