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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=rm
// Represents the value of a tag key to separate the control between resource ownership and policy enforcement.
type TagValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TagValueSpec   `json:"spec,omitempty"`
	Status TagValueStatus `json:"status,omitempty"`
}

// Represents the spec of the `TagValue`.
type TagValueSpec struct {
}

// Represents the status of the `TagValue`.
type TagValueStatus struct {
	// The tag key this tag value refers to.
	TagKeyRef corev1alpha1.NamespacedName `json:"tagKeyRef"`

	// The generated namespace of the tag value where tag value related policies like access control are stored.
	PolicyNamespace string `json:"policyNamespace,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of `TagValue`.
type TagValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []TagValue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TagValue{}, &TagValueList{})
}
