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

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	eeimportapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/import/v1"
	fleetapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/fleet/fleet/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

// +kubebuilder:object:generate=true

// ImportSpec defines the desired state of Import.
type ImportSpec struct {
	eeimportapi.ImportSpec `json:",inline"`

	// TableExistAction is the action to take when importing into an existing table.
	// The default is to skip.
	// +optional
	// +kubebuilder:default:=skip
	// +kubebuilder:validation:Enum=skip;append;truncate;replace
	TableExistAction string `json:"tableExistAction,omitempty"`
}

//+kubebuilder:object:generate=true

// ImportStatus defines the observed state of Import.
type ImportStatus struct {
	eeimportapi.ImportStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion
//+kubebuilder:printcolumn:JSONPath=".spec.dbclusterRef",name="DBCluster Name",type="string"
//+kubebuilder:printcolumn:JSONPath=".spec.databaseName",name="Database Name",type="string"
//+kubebuilder:printcolumn:JSONPath=".spec.dumpStorage.type",name="ObjectStore",type="string"
//+kubebuilder:printcolumn:JSONPath=`.status.phase`,name="Phase",type="string"

// Import is the Schema for the import API.
type Import struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImportSpec   `json:"spec,omitempty"`
	Status ImportStatus `json:"status,omitempty"`
}

// EntityStatus returns the import entity status.
func (in *Import) EntityStatus() *occoreapi.EntityStatus {
	return &in.Status.EntityStatus
}

// ImportSpec returns the common import spec.
func (in *Import) ImportSpec() eeimportapi.ImportSpec {
	return in.Spec.ImportSpec
}

// ImportStatus returns the common import status.
func (in *Import) ImportStatus() *eeimportapi.ImportStatus {
	return &in.Status.ImportStatus
}

// DBEngineName returns the database engine type string.
func (in *Import) DBEngineName() string {
	return string(fleetapi.Oracle)
}

//+kubebuilder:object:root=true

type ImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Import `json:"items"`
}

func (in *ImportList) ImportListItem() []eeimportapi.Import {
	var impItems []eeimportapi.Import
	for i := range in.Items {
		impItems = append(impItems, &in.Items[i])
	}
	return impItems
}

func init() {
	SchemeBuilder.Register(&Import{}, &ImportList{})
}

var (
	_ eeimportapi.Import     = &Import{}
	_ eeimportapi.ImportList = &ImportList{}
)
