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

	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

// InstanceRestoreSpec defines the specifications for an AlloyDB InstanceRestore.

// InstanceRestoreSpec defines the desired state of an AlloyDB InstanceRestore.
type InstanceRestoreSpec struct {
	// InstanceRestore specs that are common across all database engines.
	occoreapi.InstanceRestoreSpec `json:",inline"`
}

// InstanceRestoreStatus defines the observed state of an AlloyDB InstanceRestore.
type InstanceRestoreStatus struct {
	// InstanceRestore status that is common across all database engines.
	occoreapi.InstanceRestoreStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:name="CompleteTime",type="string",JSONPath=".status.completeTime"
// +kubebuilder:printcolumn:name="RestoredPointInTime",type="string",JSONPath=".status.restoredPointInTime"

// InstanceRestore is the Schema for the InstanceRestore API
type InstanceRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceRestoreSpec   `json:"spec,omitempty"`
	Status InstanceRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceRestoreList contains a list of AlloyDB InstanceRestore
type InstanceRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceRestore{}, &InstanceRestoreList{})
}
