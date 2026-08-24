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

	networkingv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/networking/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ftreplica
// `FlowTaggerReplica` represents `FlowTagger` resources that are synced to a particular zonal API server.
type FlowTaggerReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired configuration for `FlowTagger` resource.
	Spec FlowTaggerSpec `json:"spec,omitempty"`

	// The desired configuration for the `FlowTagger` resource.
	Status networkingv1alpha1.FlowTaggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// `FlowTaggerReplicaList` defines a list of `FlowTagger` resources.
type FlowTaggerReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowTaggerReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&FlowTaggerReplica{},
		&FlowTaggerReplicaList{},
	)
}
