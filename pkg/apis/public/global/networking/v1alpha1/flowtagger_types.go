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

	globalv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
	networkingv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/networking/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ft
// `FlowTagger` contains the API schema.
type FlowTagger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// The desired configuration for `FlowTagger` resource.
	Spec FlowTaggerSpec `json:"spec,omitempty"`

	// The observed state for `FlowTagger` resource.
	Status FlowTaggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// `FlowTaggerList` defines a list of `FlowTagger` resources.
type FlowTaggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlowTagger `json:"items"`
}

// `FlowTaggerStatus` defines the observed state of a `FlowTagger` resource.
type FlowTaggerStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []FlowTaggerZoneStatus `json:"zones,omitempty"`
}

// `FlowTaggerZoneStatus` provides the status of an flow tagger rolling out to a particular zone.
type FlowTaggerZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus networkingv1alpha1.FlowTaggerStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&FlowTagger{},
		&FlowTaggerList{},
	)
}
