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

	networkingv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/networking/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pnpreplica
// `ProjectNetworkPolicyReplica` represents `ProjectNetworkPolicy` resources that are synced to a particular zonal API server.
type ProjectNetworkPolicyReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired configuration for `ProjectNetworkPolicy` resource.
	Spec networkingv1.ProjectNetworkPolicySpec `json:"spec,omitempty"`

	// The desired configuration for the `ProjectNetworkPolicy` resource.
	Status networkingv1.ProjectNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// `ProjectNetworkPolicyReplicaList` defines a list of `ProjectNetworkPolicy` resources.
type ProjectNetworkPolicyReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectNetworkPolicyReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ProjectNetworkPolicyReplica{},
		&ProjectNetworkPolicyReplicaList{},
	)
}
