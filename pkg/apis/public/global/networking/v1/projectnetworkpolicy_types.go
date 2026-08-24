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

	globalv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
	networkingv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/networking/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pnp
// `ProjectNetworkPolicy` contains the API schema.
type ProjectNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// The desired configuration for `ProjectNetworkPolicy` resource.
	Spec networkingv1.ProjectNetworkPolicySpec `json:"spec,omitempty"`

	// The observed state for `ProjectNetworkPolicy` resource.
	Status ProjectNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// `ProjectNetworkPolicyList` defines a list of `ProjectNetworkPolicy` resources.
type ProjectNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectNetworkPolicy `json:"items"`
}

// `ProjectNetworkPolicyStatus` defines the observed state of a `ProjectNetworkPolicy` resource.
type ProjectNetworkPolicyStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ProjectNetworkPolicyZoneStatus `json:"zones,omitempty"`
}

// `ProjectNetworkPolicyZoneStatus` provides the status of an project network policy rolling out to a particular zone.
type ProjectNetworkPolicyZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus networkingv1.ProjectNetworkPolicyStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ProjectNetworkPolicy{},
		&ProjectNetworkPolicyList{},
	)
}
