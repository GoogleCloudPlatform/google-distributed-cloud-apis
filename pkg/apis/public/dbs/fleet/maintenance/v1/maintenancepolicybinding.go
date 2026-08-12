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
)

const (
	BindingValidCondition = "Valid"

	BindingDuplicatedReason        = "BindingDuplicated"
	MaintenancePolicyMissingReason = "MaintenancePolicyMissing"
	ResourceMissingReason          = "ResourceMissing"
	ValidReason                    = "Valid"
)

const (
	MaintenancePolicyBindingResourceType = "maintenancepolicybindings"
)

type MaintenancePolicyRef struct {
	ApiGroup string `json:"apiGroup"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
}

type MaintenancePolicyResourceRef struct {
	ApiGroup string `json:"apiGroup"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
}

type MaintenancePolicyBindingSpec struct {
	MaintenancePolicy MaintenancePolicyRef         `json:"maintenancePolicy"`
	Resource          MaintenancePolicyResourceRef `json:"resource"`
}

type MaintenancePolicyBindingStatus struct {
	// Conditions contain conditions for MaintenancePolicyBindings.
	// The maintenance policy binding controller sets the Valid condition.
	//
	// The following are known values for the reason field:
	// - BindingDuplicated: 		The controller found more than one binding assigned to a given resource,
	//								the status of the condition will be False.
	// - MaintenancePolicyMissing:  The maintenance policy the binding is pointing to is missing, t
	//								the status of the condition will be False.
	// - ResourceMissing:  			The resource the binding is pointing to is missing,
	//								the status of the condition will be False.
	// - Valid: 					There is a single binding for given resource, which both ends are referencing existing entities,
	//								the status of the condition will be True.
	//
	// +listType=map
	// +listMapKey=type
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,7,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +gdcloud:manifest:relevant=false,oc=ez

type MaintenancePolicyBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MaintenancePolicyBindingSpec   `json:"spec,omitempty"`
	Status MaintenancePolicyBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type MaintenancePolicyBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MaintenancePolicyBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MaintenancePolicyBinding{}, &MaintenancePolicyBindingList{})
}
