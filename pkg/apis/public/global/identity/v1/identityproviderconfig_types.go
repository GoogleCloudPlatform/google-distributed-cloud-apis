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
	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Represents a configuration for an identity provider that supports OIDC or SAML.
// +genclient
type IdentityProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityProviderConfigSpec   `json:"spec,omitempty"`
	Status IdentityProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `IdentityProviderConfig` resources.
type IdentityProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderConfig `json:"items"`
}

// Provides the status of an `IdentityProviderConfig` resource.
type IdentityProviderConfigStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []IdentityProviderConfigZoneStatus `json:"zones,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// IdentityProviderConfigZoneStatus provides the status of a IdentityProviderConfig rolling
// out to a particular zone.
type IdentityProviderConfigZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus IdentityProviderConfigReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&IdentityProviderConfig{}, &IdentityProviderConfigList{})
}
