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
// +kubebuilder:subresource:status

// Specifies the value of a quota in a specific quota dimension combination.
type QuotaValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuotaValueSpec   `json:"spec,omitempty"`
	Status QuotaValueStatus `json:"status,omitempty"`
}
// +kubebuilder:validation:XValidation:rule="(has(self.dimensionMap) ? self.dimensionMap : []) == (has(oldSelf.dimensionMap) ? oldSelf.dimensionMap : [])", message="DimensionMap is immutable"
type QuotaValueSpec struct {
	// The identifier of the quota that this quota value applies to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Quota identifier is immutable"
	Quota QuotaIdentifier `json:"quota"`

	// The dimension combination that this quota value resource applies to. The key of the map is the name of a dimension, such as region, zone, network_id, and the value of the map is the dimension value.
	// If a dimension in the quota’s dimension does not appear in the dimensions map, the quota value applies to all the dimension values except for those that have another quota value configured for the specific value. In other words, a QuotaValue with a more specific dimension map overrides a QuotaValue with a less specific dimension map. For example, one with dimension map {region: gcd-region1} overrides the one with an empty dimension map.
	// +listType=map
	// +listMapKey=name
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MaxItems=64
	DimensionMap []Dimension `json:"dimensionMap,omitempty"`

	// The desired quota value.
	Value int64 `json:"value"`
}

type QuotaValueStatus struct {
	// +optional
	GrantedValue int64 `json:"grantedValue,omitempty"`

	// +optional
	State StateType `json:"state,omitempty"`

	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The most recent generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

type QuotaIdentifier struct {
	// The API group for which the quota is defined. It corresponds to the service name of a OnePlatform config.
	Group string `json:"group"`

	// The name of the quota this quota value resource is for.
	Name string `json:"name"`
}

// Represents a dimension of a quota.
type Dimension struct {
	// The name of a dimension, such as region, zone, network_id
	Name string `json:"name"`

	// The value of the dimension.
	Value string `json:"value"`
}

// +kubebuilder:validation:Enum=Pending;Granted;PartiallyGranted;Denied
type StateType string

const (
	StateTypePending          StateType = "Pending"
	StateTypeGranted          StateType = "Granted"
	StateTypePartiallyGranted StateType = "PartiallyGranted"
	StateTypeDenied           StateType = "Denied"
)

// +kubebuilder:object:root=true
// Represents a list of QuotaValues
type QuotaValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuotaValue `json:"items"`
}

func init() {
	SchemeBuilder.Register(&QuotaValue{}, &QuotaValueList{})
}
