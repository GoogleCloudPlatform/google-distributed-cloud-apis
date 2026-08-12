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

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1alpha1"
)

// +kubebuilder:object:root=true
// +gdcloud:manifest:relevant=false,oc=rm
// +genclient
// Represents a project the user has access to.
type ProjectView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   v1alpha1.ProjectSpec   `json:"spec,omitempty"`
	Status v1alpha1.ProjectStatus `json:"status,omitempty"`
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
