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
	v1 "k8s.io/api/core/v1"
)

type ExternalServerPhase string

const (
	PasswordSecretPassword = "password"
	CertSecretCACrt        = "ca.crt"
	CertSecretTLSCrt       = "tls.crt"
	CertSecretTLSKey       = "tls.key"
)

//+kubebuilder:object:generate=true

// ExternalServerSpec contains metadata for an external database server.
type ExternalServerSpec struct {
	// Host is the host ip of the external database server.
	// +kubebuilder:validation:Required
	Host string `json:"host"`

	// Port is the port of the external database server.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`

	// UserName is the name of the user used  to connect to the external database.
	// +kubebuilder:validation:Required
	Username string `json:"username"`

	// Password is the reference to the secret storing external database server password.
	// +kubebuilder:validation:Required
	Password *v1.SecretReference `json:"password"`

	// CertRef is the reference to the secret storing external database server certificate.
	// +kubebuilder:validation:Optional
	CertRef *v1.SecretReference `json:"certRef,omitempty"`
}

//+kubebuilder:object:generate=true

// ExternalServerStatus is the status of an external database server.
type ExternalServerStatus struct {
	EntityStatus `json:",inline"`
}
