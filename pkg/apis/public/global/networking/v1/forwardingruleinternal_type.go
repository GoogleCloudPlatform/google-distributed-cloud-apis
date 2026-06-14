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

	globalv1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/common/global/v1alpha1"
	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
	networkingv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/networking/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=fri

// Represents a frontend API to create internal forwarding rule.
//
// +kubebuilder:printcolumn:name="BackendService",type="string",JSONPath=".spec.backendServiceRef.name"
// +kubebuilder:printcolumn:name="CIDR",type="string",JSONPath=".metadata.annotations['networking\\.gke\\.io/forwardingRuleCIDR']"
type ForwardingRuleInternal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   networkingv1.ForwardingRuleInternalSpec `json:"spec"`
	Status ForwardingRuleInternalStatus            `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of ForwardingRuleInternal.
type ForwardingRuleInternalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ForwardingRuleInternal `json:"items"`
}

// Represents the status of forwarding rule.
type ForwardingRuleInternalStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"` // embed the duck type for global status

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ForwardingRuleInternalZoneStatus `json:"zones,omitempty"`

	// ErrorStatus holds the most recent errors with last seen time.
	//
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// Subnets holds the list of subnets that are used for this forwarding rule.
	//
	// +optional
	Subnets []networkingv1.SubnetReference `json:"subnets,omitempty"`
}

// Represents the status of a given zone's internal forwarding rule.
type ForwardingRuleInternalZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"` // embed the duck type for zone status

	// The reconciliation status of the replica collected from the zone.
	ReplicaStatus networkingv1.ForwardingRuleInternalStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ForwardingRuleInternal{}, &ForwardingRuleInternalList{})
}
