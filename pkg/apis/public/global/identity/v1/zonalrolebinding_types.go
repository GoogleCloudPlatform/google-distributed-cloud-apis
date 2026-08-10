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
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ZonalRoleBinding references a zonal Role and adds who information via Subject.
// +genclient
type ZonalRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZonalRoleBindingSpec   `json:"spec,omitempty"`
	Status ZonalRoleBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of ZonalRoleBinding resources.
type ZonalRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZonalRoleBinding `json:"items"`
}

// Provides the status of the ZonalRoleBinding resource.
type ZonalRoleBindingStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ZonalRoleBindingZoneStatus `json:"zones,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// ZonalRoleBindingZoneStatus provides the status of a ZonalRoleBinding rolling
// out to a particular zone.
type ZonalRoleBindingZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus ZonalRoleBindingReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ZonalRoleBinding{}, &ZonalRoleBindingList{})
}
