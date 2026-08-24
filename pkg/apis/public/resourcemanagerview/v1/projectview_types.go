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

	rmv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1"
)

// +gdcloud:manifest:relevant=true,oc=rm,entities="projects"
// +gdcloud:manifest:verbs=describe;list
// +gdcloud:manifest:rbac="list:custom-role-org-admin,custom-role-org-viewer"
// +gdcloud:manifest:rbac="describe,list:marketplace-catalog-publisher,project-creator,project-editor,user-cluster-admin,marketplace-catalog-viewer,standard-cluster-admin,standard-cluster-viewer"
// +gdcloud:manifest:skipcodegen=true
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +genclient
// Represents a project the user has access to.
type ProjectView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status rmv1.ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a list of `ProjectView` objects.
type ProjectViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectView `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ProjectView{},
		&ProjectViewList{},
	)
}
