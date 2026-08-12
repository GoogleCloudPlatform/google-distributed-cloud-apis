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

	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

// InstanceSpec defines the desired state of PostgresInstance
type InstanceSpec struct {
	// Instance specs that are common across all database engines.
	occoreapi.InstanceSpec `json:",inline"`
}

// InstanceStatus defines the observed state of PostgresInstance
type InstanceStatus struct {
	// Instance status that is common across all database engines.
	occoreapi.InstanceStatus `json:",inline"`

	// PrimaryPodIP indicates the IP of postgres primary pod.
	PrimaryPodIP string `json:"primaryPodIP,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".metadata.labels['dbs\\.internal\\.dbadmin\\.goog/ha-role']",name="Role",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.phase`,name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.description`,name="Message",type="string"

// Instance is the Schema for the instances API
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceSpec   `json:"spec,omitempty"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Postgres Instance
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
