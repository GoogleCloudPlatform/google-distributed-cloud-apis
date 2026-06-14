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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Name",JSONPath=".metadata.name",type=string
// +kubebuilder:printcolumn:name="Active",JSONPath=".spec.active",type=boolean
// +kubebuilder:printcolumn:name="Location Type",JSONPath=".spec.locationType",type=string
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type=string
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// Defines the schema for the BucketLocationReplica API.
// +genclient
// +genclient:nonNamespaced
type BucketLocationReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketLocationSpec          `json:"spec,omitempty"`
	Status BucketLocationReplicaStatus `json:"status,omitempty"`
}

type BucketLocationReplicaStatus struct {
	// Conditions represents the observations of this bucket location's overall
	// state.
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Contains a list of BucketLocationReplicas.
// +kubebuilder:object:root=true
type BucketLocationReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLocationReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BucketLocationReplica{}, &BucketLocationReplicaList{})
}
