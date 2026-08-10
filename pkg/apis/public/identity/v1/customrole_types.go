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
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Represents a template for a zonal CustomRole
// Custom roles provide fine-grained control over user permissions, unlike predefined roles.
// This allows organizations to tailor access rights to their specific needs, balancing operational
// efficiency with security. By adhering to the principle of least privilege, custom roles
// significantly enhance security and protect sensitive data.
// +gdcloud:manifest:relevant=true,oc=iam,component=iam,entities="roles"
// +gdcloud:manifest:verbs=create;delete;describe;list;update;
// +gdcloud:manifest:rbac="create,delete,describe,update,list:organization-iam-admin,custom-role-org-admin,custom-role-project-admin,project-iam-admin"
// +gdcloud:manifest:rbac="describe,list:organization-iam-viewer,custom-role-org-viewer,custom-role-project-viewer,project-iam-viewer"
// +gdcloud:manifest:skipcodegen=true
// +genclient
type CustomRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomRoleSpec   `json:"spec,omitempty"`
	Status CustomRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `CustomRole` resource
type CustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomRole `json:"items"`
}

// Provides a status of CustomRole
type CustomRoleStatus struct {
	// Conditions represents the observations of this Custom role overall state
	// +listType=map
	// +listMapKey=type
	// + optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// propagation information of converted template for zonal role template conversion
	// + optional
	PropagationInfo PropagationInfo `json:"propagationInfo,omitempty"`
}

func init() {
	SchemeBuilder.Register(&CustomRole{}, &CustomRoleList{})
}
