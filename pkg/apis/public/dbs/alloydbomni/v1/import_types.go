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

	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

// ImportSpec defines the desired state of AlloyDB Import.
type ImportSpec struct {
	// ImportSpec includes import specs common across all database engines.
	occoreapi.ImportSpec `json:",inline"`

	// TableExistAction is the action to take when importing into an existing table.
	// The default is to skip.
	// +optional
	// +kubebuilder:default:=skip
	// +kubebuilder:validation:Enum=skip
	TableExistAction string `json:"tableExistAction,omitempty"`
}

// ImportStatus defines the observed state of AlloyDB Import.
type ImportStatus struct {
	// Import status that is common across all database engines.
	occoreapi.ImportStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".spec.instanceRef",name="Instance Name",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.databaseName",name="Database Name",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.dumpStorage.type",name="ObjectStore",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.phase`,name="Phase",type="string"

// Import is the Schema for the import API.
type Import struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImportSpec   `json:"spec,omitempty"`
	Status ImportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImportList contains a list of AlloyDB Imports.
type ImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Import `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Import{}, &ImportList{})
}
