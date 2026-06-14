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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Represents a template for a zonal CustomRole replica
// +gdcloud:manifest:relevant=false,oc=iam
// +genclient
type CustomRoleReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomRoleSpec          `json:"spec,omitempty"`
	Status CustomRoleReplicaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of zonal CustomRole replica resources
type CustomRoleReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomRole `json:"items"`
}

// Provides a status of zonal CustomRole replica
type CustomRoleReplicaStatus struct {
	// Conditions represents the observations of this Custom role overall state
	// +listType=map
	// +listMapKey=type
	// + optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Propagated custom role name for all the replicas
	// + optional
	PropagatedCustomRoleName string `json:"propagatedCustomRoleName,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&CustomRoleReplica{}, &CustomRoleReplicaList{})
}
