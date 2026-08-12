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

package v2

import (
	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// Represents a managed instance of a Keycloak server.
type ManagedKeycloak struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the ManagedKeycloak instance.
	Spec ManagedKeycloakSpec `json:"spec,omitempty"`
	// The most recently observed status of the ManagedKeycloak instance.
	Status ManagedKeycloakStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of ManagedKeycloak instances.
type ManagedKeycloakList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedKeycloak `json:"items"`
}

// Represents the specification or the desired state of a ManagedKeycloak instance.
type ManagedKeycloakSpec struct {
	// Configuration for number of replicas.
	// +optional
	Replicas int `json:"replicas"`

	// Configuration for the database cluster.
	// +optional
	Database *DatabaseSpec `json:"database,omitempty"`
}

// DatabaseSpec defines the configuration for the database cluster.
type DatabaseSpec struct {
	// CPU limit for the database primary instance.
	// +optional
	ResourcesLimitCPU string `json:"resourcesLimitCPU,omitempty"`

	// Memory limit for the database primary instance.
	// +optional
	ResourcesLimitMemory string `json:"resourcesLimitMemory,omitempty"`

	// Volume size for the database cluster.
	// +optional
	VolumeSize string `json:"volumeSize,omitempty"`
}

// Provides the status of an `ManagedKeycloak` resource.
type ManagedKeycloakStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// AdminCredentialSecretRef is the reference to the secret containing the admin credentials.
	// +optional
	AdminCredentialSecretRef string `json:"adminCredentialSecretRef,omitempty"`

	// AdminConsoleURL is the URL for the Keycloak admin console.
	// +optional
	AdminConsoleURL string `json:"adminConsoleURL,omitempty"`

	// ErrorStatus is the error status of the ManagedKeycloak instance.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ManagedKeycloak{},
		&ManagedKeycloakList{},
	)
}
