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
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +gdcloud:manifest:relevant=true,oc=platauth,component=pki,entities="ca-pools"
// +gdcloud:manifest:verbs=create;update;delete;list;describe
// +gdcloud:manifest:skipcodegen=true

// CaPool represents a pool of Certificate Authorities.
type CaPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CaPoolSpec   `json:"spec,omitempty"`
	Status CaPoolStatus `json:"status,omitempty"`
}

type CaPoolSpec struct {
	// Description of the CA Pool.
	// +optional
	Description string `json:"description,omitempty"`
}

type CAStatus struct {
	// Name of the CA.
	Name string `json:"name"`

	// List of status conditions to indicate the status of the CA.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type CaPoolStatus struct {
	// CAs populated dynamically by the CaPoolReconciler.
	// +optional
	CAs []CAStatus `json:"cas,omitempty"`

	// List of status conditions to indicate the status of the CaPool.
	// - Ready: Indicates that the CaPool is ready to use.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// CaPoolList represents a collection of CA pools.
type CaPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CaPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&CaPool{},
		&CaPoolList{},
	)
}
