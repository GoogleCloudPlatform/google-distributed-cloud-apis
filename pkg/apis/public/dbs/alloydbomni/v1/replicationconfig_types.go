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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=rcg
// +kubebuilder:printcolumn:JSONPath=`.spec.parent.dbnode.name`,name="Parent",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.type",name="Type",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.role",name="Role",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="Ready")].status`,name="Ready",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="Healthy")].status`,name="Healthy",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.physicalUpstream.synchronousEnabled`,name="PUpstream",type="boolean"
// +kubebuilder:printcolumn:JSONPath=`.status.physicalDownstream.synchronousEnabled`,name="PDownstream",type="boolean"
type ReplicationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   occoreapi.ReplicationConfigSpec   `json:"spec,omitempty"`
	Status occoreapi.ReplicationConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type ReplicationConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReplicationConfig{}, &ReplicationConfigList{})
}
