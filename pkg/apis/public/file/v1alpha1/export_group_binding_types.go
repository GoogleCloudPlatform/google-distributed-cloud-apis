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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=file
// +genclient
// Ensures that a given share is accessible to a given export group. If the
// subnet is covered by the export group, the volume should be accessible at the endpoint
// specified in the access binding.
type ExportGroupBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of an export group binding.
	// +optional
	Spec ExportGroupBindingSpec `json:"spec,omitempty"`

	// The observed state of an export group binding.
	// Read-only.
	// +optional
	Status ExportGroupBindingStatus `json:"status,omitempty"`
}

// Represents the desired state of an export group binding.
type ExportGroupBindingSpec struct {
	// An export group.
	ExportGroupRef *corev1.ObjectReference `json:"exportGroupRef"`

	// A file share.
	FileShareRef *corev1.ObjectReference `json:"fileShareRef"`
}

// Represents the observed state of an export group binding.
type ExportGroupBindingStatus struct {
	// A list of observed conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of export group bindings.
type ExportGroupBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ExportGroupBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ExportGroupBinding{},
		&ExportGroupBindingList{},
	)
}
