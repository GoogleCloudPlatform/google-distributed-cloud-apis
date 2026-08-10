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

// ExportSpec defines the desired state of alloydbomni Export.
type ExportSpec struct {
	// ExportSpec includes export specs common across all database engines.
	occoreapi.ExportSpec `json:",inline"`
}

// ExportStatus defines the observed state of alloydbomni Export.
type ExportStatus struct {
	// Export status that is common across all database engines.
	occoreapi.ExportStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".spec.instanceRef",name="Instance Name",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.exportLocation.type",name="ObjectStore",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.phase`,name="Phase",type="string"

// Export is the Schema for the export API.
type Export struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExportSpec   `json:"spec,omitempty"`
	Status ExportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExportList contains a list of alloydbomni Exports.
type ExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Export `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Export{}, &ExportList{})
}
