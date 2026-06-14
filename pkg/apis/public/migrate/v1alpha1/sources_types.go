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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MigrationSource CR
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=m4gdc
// +genclient
type MigrationSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MigrationSourceSpec   `json:"spec" valid:"required"`
	Status            MigrationSourceStatus `json:"status,omitempty"`
}

type MigrationSourceSpec struct {
	// Enum of whether the Source is a vCenter, OVA file bucket, etc
	MigrationSourceType string `json:"MigrationSourceType"`
	// URL is the url of the bucket source (or) vCenter URL
	URL string `json:"url"`
	// SecretRef provides the secret reference needed to access the vCenter (or) bucket source
	// +optional
	SecretRef string `json:"secretRef,omitempty"`
	// CertificateRef provides the secret reference for the certificate needed to access the vCenter (or) bucket source
	// +optional
	CertificateRef string `json:"certificateRef,omitempty"`
	// LastRequestedDiscovery is a datetime representing the last time that a discovery was manually requested
	LastRequestedDiscovery metav1.Time `json:"LastRequestedDiscovery,omitempty"`
}

type MigrationSourceStatus struct {
	// Current discovery status of the migration source.
	//+kubebuilder:default:="NOT_STARTED"
	//+optional
	DiscoveryStatus *DiscoveryStatus `json:"discoveryStatus,omitempty"`
	// LastCompletedDiscovery is a datetime representing the last time that a discovery was completed
	LastCompletedDiscovery string `json:"LastCompletedDiscovery,omitempty"`
}

type DiscoveryStatus string

const (
	DiscoveryStatusNotStarted DiscoveryStatus = "NOT_STARTED"
	DiscoveryStatusInProgress DiscoveryStatus = "IN_PROGRESS"
	DiscoveryStatusCompleted  DiscoveryStatus = "COMPLETED"
	DiscoveryStatusFailed     DiscoveryStatus = "FAILED"
	DiscoveryStatusUnknown    DiscoveryStatus = "UNKNOWN"
)

// +kubebuilder:object:root=true
type MigrationSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MigrationSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&MigrationSource{},
		&MigrationSourceList{},
	)
}
