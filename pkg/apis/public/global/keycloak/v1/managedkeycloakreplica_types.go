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

// Represents a managed instance of a Keycloak server.
// +genclient
type ManagedKeycloakReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the ManagedKeycloak instance.
	Spec ManagedKeycloakSpec `json:"spec,omitempty"`
	// The most recently observed status of the ManagedKeycloak instance.
	Status ManagedKeycloakReplicaStatus `json:"status,omitempty"`
}

// Represents a collection of ManagedKeycloakReplica instances.
// +kubebuilder:object:root=true
type ManagedKeycloakReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKeycloakReplica `json:"items"`
}

// Represents the specification or the desired state of a ManagedKeycloak instance.
type ManagedKeycloakSpec struct {
	// The name of the Kubernetes Secret containing the initial admin credentials.
	// The secret must contain `username` and `password` keys.
	AdminCredentialSecretName string `json:"adminCredentialSecretName"`

	// Configuration for number of replicas.
	// +optional
	Replicas int `json:"replicas"`
}

// Provides the status of an `ManagedKeycloak` resource.
type ManagedKeycloakReplicaStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ManagedKeycloakReplica{}, &ManagedKeycloakReplicaList{})
}
