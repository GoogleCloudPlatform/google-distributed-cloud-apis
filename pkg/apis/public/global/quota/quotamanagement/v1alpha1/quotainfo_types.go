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
// +kubebuilder:storageversion

// A view of the quota definition and effective quota value for a particular project or organization. The effective quota value is derived from the default quota defined in QuotaDefinition and the grantedValue of the QuotaValue of a project/organization and all its ancestors.
type QuotaInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuotaInfoSpec   `json:"spec"`
	Status QuotaInfoStatus `json:"status,omitempty"`
}

type QuotaInfoSpec struct {
	// The API group for which the quota is defined. It corresponds to the service name of a OnePlatform config.
	Group string `json:"group"`
	// The quota definition
	Definition Quota `json:"definition,omitempty"`
	// Quota values for different dimension combinations.
	Values []QuotaValueInfo `json:"values,omitempty"`
	// Whether quota increase is allowed
	AllowIncrease bool `json:"allowIncrease"`
}

// +kubebuilder:object:root=true

// Represents a list of QuotaInfos
type QuotaInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuotaInfo `json:"items"`
}

type QuotaInfoStatus struct {
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type QuotaValueInfo struct {

	// The dimension combination that this quota value resource applies to. The key of the map is the name of a dimension, such as region, zone, network_id, and the value of the map is the dimension value.
	// If a dimension in the quota’s dimension does not appear in the dimensions map, the quota value applies to all the dimension values except for those that have another quota value configured for the specific value. In other words, a QuotaValue with a more specific dimension map overrides a QuotaValue with a less specific dimension map. For example, one with dimension map {region: gcd-region1} overrides the one with an empty dimension map.

	// +listType=map
	// +listMapKey=name
	// +optional
	DimensionMap []Dimension `json:"dimensionMap,omitempty"`
	// The quota value for the quota dimension combination.
	GrantedValue int64 `json:"grantedValue"`
	// The quota value to be reset to if the quota value  is deleted. Quota preference up to this amount does not require approval.
	ResetValue int64 `json:"resetValue"`
}

func init() {
	SchemeBuilder.Register(&QuotaInfo{}, &QuotaInfoList{})
}
