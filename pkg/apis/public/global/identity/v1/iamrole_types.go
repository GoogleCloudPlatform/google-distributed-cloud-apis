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

// IAMRole references a predefined role in the global and/or zonal APIServer.
// +genclient
type IAMRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMRoleSpec   `json:"spec,omitempty"`
	Status IAMRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of IAMRole resources.
type IAMRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMRole `json:"items"`
}

// Provides the specification for IAMRole.
type IAMRoleSpec struct {
	// RoleRef contains information that points to the role in the global APIServer being used.
	RoleRef *rbacv1.RoleRef `json:"roleRef,omitempty"`

	// ZonalRoleRef contains information that points to the zonal role being used.
	ZonalRoleRef *rbacv1.RoleRef `json:"zonalRoleRef,omitempty"`
}

// Provides the status of the IAMRole.
type IAMRoleStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&IAMRole{}, &IAMRoleList{})
}
