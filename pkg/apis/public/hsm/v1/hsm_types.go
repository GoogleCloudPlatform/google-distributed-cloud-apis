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

// +gdcloud:manifest:relevant=false,oc=hsm
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Represents a resource on the CipherTrust Manager.
type CTMKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CTMKeySpec   `json:"spec,omitempty"`
	Status CTMKeyStatus `json:"status,omitempty"`
}

// Provides the specification for a CTM key.
type CTMKeySpec struct {
	// The domain a CTM key resides.
	Domain CTMKeyDomain `json:"domain"`

	// The details of a CTM key.
	Details KeyResource `json:"details,omitempty"`
}

// Provides the domain a CTM key resides.
type CTMKeyDomain struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// Provides the details of a CTM key.
type KeyResource struct {
	ID             string   `json:"id"`
	URI            string   `json:"uri,omitempty"`
	CreatedAt      string   `json:"createdAt,omitempty"`
	DestroyDate    string   `json:"destroyDate,omitempty"`
	Name           string   `json:"name"`
	UpdatedAt      string   `json:"updatedAt,omitempty"`
	Meta           *KeyMeta `json:"meta,omitempty"`
	Version        int      `json:"version,omitempty"`
	Algorithm      string   `json:"algorithm,omitempty"`
	Size           int      `json:"size,omitempty"`
	ObjectType     string   `json:"objectType,omitempty"`
	ActivationDate string   `json:"activationDate,omitempty"`
	State          string   `json:"state,omitempty"`
	UUID           string   `json:"uuid,omitempty"`
}
type KeyMeta struct {
	CustomAttributes map[string]string `json:"customAttributes,omitempty"`
}

// Provides the status of a CTM key.
type CTMKeyStatus struct {
	// The status of a CTM key.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The name of a CTM key.
	Name string `json:"name"`
}

// +kubebuilder:object:root=true

// Represents a collection of CTM keys.
type CTMKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []CTMKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&CTMKey{},
		&CTMKeyList{},
	)
}
