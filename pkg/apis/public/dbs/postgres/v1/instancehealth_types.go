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

	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

// InstanceHealthStatus defines the observed state of a Postgres InstanceHealth.
type InstanceHealthStatus struct {
	// InstanceHealth status that is common across all database engines.
	occoreapi.InstanceHealthStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=`.status.consecutiveHealthcheckFailures`,name="Failures",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.lastHealthcheckRunTime`,name="Last Healthcheck",type="string"

// InstanceHealth is the Schema for the InstanceHealth API
type InstanceHealth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status InstanceHealthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceHealthList contains a list of Postgres InstanceHealth
type InstanceHealthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceHealth `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceHealth{}, &InstanceHealthList{})
}
