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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=fre

// Represents a frontend API to create external forwarding rule.
//
// +kubebuilder:printcolumn:name="BackendService",type="string",JSONPath=".spec.backendServiceRef.name"
// +kubebuilder:printcolumn:name="CIDR",type="string",JSONPath=".status.cidr"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +gdcloud:manifest:relevant=true,oc=unet,component=compute,entities="forwarding-rules-external"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:external-load-balancer-admin,load-balancer-admin"
// +gdcloud:manifest:rbac="describe,list:external-load-balancer-viewer"
// +gdcloud:manifest:skipcodegen=true
type ForwardingRuleExternal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   ForwardingRuleExternalSpec   `json:"spec"`
	Status ForwardingRuleExternalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of ForwardingRuleExternal.
type ForwardingRuleExternalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ForwardingRuleExternal `json:"items"`
}

// Describes the attributes that a user expects from a forwarding rule.
type ForwardingRuleExternalSpec struct {
	ForwardingRuleSpecCommon `json:",inline"`
}

// Represents the status of forwarding rule.
type ForwardingRuleExternalStatus struct {
	ForwardingRuleStatusCommon `json:",inline"`
}

func init() {
	SchemeBuilder.Register(&ForwardingRuleExternal{}, &ForwardingRuleExternalList{})
}
