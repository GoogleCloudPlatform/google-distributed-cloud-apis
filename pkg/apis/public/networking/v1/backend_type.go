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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=be

// Identifies endpoints for a load balancer.
type Backend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   BackendSpec   `json:"spec"`
	Status BackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains list of Backends.
type BackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backend `json:"items"`
}

// Describes the attributes that a user expects from backend.
//
// +kubebuilder:validation:XValidation:rule="has(oldSelf.clusterName) == has(self.clusterName)", message="ClusterName is immutable"
type BackendSpec struct {
	// A name of cluster to which the scope of the defined selectors should be
	// limited to. This does not apply to VM workloads. This field is optional.
	// This field is immutable.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ClusterName is immutable"
	// +kubebuilder:validation:MaxLength=250
	ClusterName *string `json:"clusterName,omitempty"`

	// A selector defining which endpoints (Pods or VMs) to use for this backend.
	// This field is required. This field is immutable.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="has(self.matchLabels)", message="MatchLabels is required"
	// +kubebuilder:validation:XValidation:rule="!has(self.matchLabels) || size(self.matchLabels) > 0", message="MatchLabels must have at least 1 label"
	EndpointsLabels metav1.LabelSelector `json:"endpointsLabels"`
}


// Represents the status of backend.
type BackendStatus struct {
	// A list of conditions describing the current state of the backend.
	// Known condition types are:
	// * "Ready"
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// ErrorStatus holds the most recent errors with last seen time.
	//
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// Holds the current configuration parameters applied to the backend.
	// This field is optional.
	//
	// +optional
	// +kubebuilder:validation:Optional
	ActiveConfiguration *ActiveConfiguration `json:"activeConfiguration,omitempty"`
}

// Defines the active configuration parameters applied to the backend.
type ActiveConfiguration struct {

	// The active connection draining value being applied to backends. A value of 0 disables connection draining.
	// This field is optional.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	ConnectionDraining *int32 `json:"connectionDraining,omitempty"`
}

// BackendReadyConditionReason defines the set of reasons that explain
// why a particular Backend Ready condition status is set.
type BackendReadyConditionReason string

const (
	// ClusterNotFound indicates that the referenced user cluster object was not found.
	ClusterNotFound BackendReadyConditionReason = "ClusterNotFound"

	// OperationForbidden indicates that create/update/delete operation is not allowed
	// in the API.
	OperationForbidden BackendReadyConditionReason = "OperationForbidden"
)

func init() {
	SchemeBuilder.Register(&Backend{}, &BackendList{})
}
