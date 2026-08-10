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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=rm
// Represents the key of a tag to separate the control between resource ownership and policy enforcement.
type TagKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagKeySpec   `json:"spec,omitempty"`
	Status TagKeyStatus `json:"status,omitempty"`
}

// Represents the spec of the `TagKey`.
type TagKeySpec struct {
}

// Represents the status of the `TagKey`.
type TagKeyStatus struct {
	// The generated namespace of the tag key where the tag values are stored.
	ManagementNamespace string `json:"managementNamespace,omitempty"`

	// The label key generated from the tag key.
	// Each tag key generates a corresponding label key. The label key is attached to the
	// resources.
	// Users can use this label key to group resources with tags.
	LabelKey string `json:"labelKey"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of `TagKey`.
type TagKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TagKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TagKey{}, &TagKeyList{})
}
