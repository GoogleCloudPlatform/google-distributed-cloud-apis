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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=file
// +genclient
// Allows for control of who should have access for a share.
type ExportGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of an export group.
	// +optional
	Spec ExportGroupSpec `json:"spec,omitempty"`

	// The observed state of an export group.
	// Read-only.
	// +optional
	Status ExportGroupStatus `json:"status,omitempty"`
}

// Represents the desired state of an export group.
type ExportGroupSpec struct {
	// A list of subnets whose endpoints constitute the export group.
	// 0.0.0.0/0 is acceptable.
	// +optional
	Subnets []IPSubnetString `json:"subnets,omitempty"`
}

// Represents the observed state of an export group.
type ExportGroupStatus struct {
	// A list of observed conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of export groups.
type ExportGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ExportGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ExportGroup{},
		&ExportGroupList{},
	)
}

// +kubebuilder:validation:Pattern=`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/(?:3[0-2]|[12]?[0-9])$`
type IPSubnetString string

type IPSubnetStringList []IPSubnetString
