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

	eeexportapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/export/v1"
	fleetapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/fleet/fleet/v1"
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

// +kubebuilder:object:generate=true

// ExportSpec defines the desired state of Export.
type ExportSpec struct {
	eeexportapi.ExportSpec `json:",inline"`
}

//+kubebuilder:object:generate=true

// ExportStatus defines the observed state of Export.
type ExportStatus struct {
	eeexportapi.ExportStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion
//+kubebuilder:printcolumn:JSONPath=".spec.dbclusterRef",name="DBCluster Name",type="string"
//+kubebuilder:printcolumn:JSONPath=".spec.exportLocation.type",name="ObjectStore",type="string"
//+kubebuilder:printcolumn:JSONPath=`.status.phase`,name="Phase",type="string"

// Export is the Schema for the export API.
type Export struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExportSpec   `json:"spec,omitempty"`
	Status ExportStatus `json:"status,omitempty"`
}

// EntityStatus returns the export entity status.
func (in *Export) EntityStatus() *occoreapi.EntityStatus {
	return &in.Status.EntityStatus
}

// ExportSpec returns the common export spec.
func (in *Export) ExportSpec() eeexportapi.ExportSpec {
	return in.Spec.ExportSpec
}

// ExportStatus returns the common export status.
func (in *Export) ExportStatus() *eeexportapi.ExportStatus {
	return &in.Status.ExportStatus
}

// DBEngineName returns the database engine type string.
func (in *Export) DBEngineName() string {
	return string(fleetapi.AlloyDBOmni)
}

//+kubebuilder:object:root=true

type ExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Export `json:"items"`
}

func (in *ExportList) ExportListItem() []eeexportapi.Export {
	var expItems []eeexportapi.Export
	for i := range in.Items {
		expItems = append(expItems, &in.Items[i])
	}
	return expItems
}

func init() {
	SchemeBuilder.Register(&Export{}, &ExportList{})
}

var (
	_ eeexportapi.Export     = &Export{}
	_ eeexportapi.ExportList = &ExportList{}
)
