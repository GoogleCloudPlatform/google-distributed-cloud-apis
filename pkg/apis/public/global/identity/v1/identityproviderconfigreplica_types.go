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
	identityv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/identity/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Represents a configuration for an identity provider that supports OIDC or SAML.
// +gdcloud:manifest:relevant=false,oc=iam
// +genclient
type IdentityProviderConfigReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityProviderConfigSpec          `json:"spec,omitempty"`
	Status IdentityProviderConfigReplicaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `IdentityProviderConfigReplica` resources.
type IdentityProviderConfigReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderConfigReplica `json:"items"`
}

// Provides the specification, or desired state, of an `IdentityProviderConfig` resource.
// Either OIDCConfig or SAMLConfig has to be provided but not both.
type IdentityProviderConfigSpec struct {
	// OIDC specific configuration.
	// +optional
	OIDCConfig *identityv1.OIDCProviderConfig `json:"oidc,omitempty"`

	// SAML specific configuration.
	// +optional
	SAMLConfig *identityv1.SAMLProviderConfig `json:"saml,omitempty"`
}

// Provides the status of an `IdentityProviderConfig` resource.
type IdentityProviderConfigReplicaStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&IdentityProviderConfigReplica{}, &IdentityProviderConfigReplicaList{})
}
