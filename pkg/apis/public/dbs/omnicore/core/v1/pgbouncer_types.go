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
	corev1 "k8s.io/api/core/v1"
)

type PgBouncerPhase string

const (
	FetchingIP                PgBouncerPhase = "Acquiring IP"
	PgBouncerReady            PgBouncerPhase = "Ready"
	WaitingForDBReady         PgBouncerPhase = "WaitingForDBReady"
	WaitingForDeploymentReady PgBouncerPhase = "WaitingForDeploymentReady"
	PgBouncerError            PgBouncerPhase = "Error"
	AccessModeReadOnly        string         = "ro"
	AccessModeReadWrite       string         = "rw"
	DefaultFrontendPort       int32          = 6432
)

// +kubebuilder:object:generate=true
type PgBouncerSpec struct {
	// +kubebuilder:validation:Required
	DBClusterRef DBClusterRef `json:"dbclusterRef"`
	// +kubebuilder:default="rw"
	// +kubebuilder:validation:Enum=rw;ro
	AccessMode string `json:"accessMode,omitempty"`
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	ReplicaCount int32             `json:"replicaCount,omitempty"`
	Parameters   map[string]string `json:"parameters,omitempty"`
	// +optional
	AllowSuperUserAccess bool `json:"allowSuperUserAccess,omitempty"`
	// +optional
	ServerTLS *ServerTLSSpec `json:"serverTLS,omitempty"`
}

//+kubebuilder:object:generate=true

// ServerTLSSpec defines the certificate secret for encrypted communication
// used by PgBouncer to connect to database cluster for auth query.
type ServerTLSSpec struct {
	// CertSecret references the certificate secret within the same namespace.
	// The secret must contain entries ca.crt (CA certificate), tls.key (private key),
	// and tls.crt (leaf certificate). The values within this secret are
	// used to populate the server_tls_ca_file, server_tls_cert_file, and
	// server_tls_key_file in pgbouncer.ini. The CA certificate must match the CA
	// that signed the database cluster's leaf certificate. The leaf certificate
	// must contain the CommonName "alloydbpgbouncer".
	CertSecret *corev1.LocalObjectReference `json:"certSecret,omitempty"`
}

// +kubebuilder:object:generate=true
// PgBouncerStatus defines the observed state of PgBouncer.
type PgBouncerStatus struct {
	Endpoints []Endpoint     `json:"endpoints,omitempty"`
	Phase     PgBouncerPhase `json:"phase"`
	Port      int32          `json:"port,omitempty"`
	// +optional
	EntityStatus `json:",inline"`
}
