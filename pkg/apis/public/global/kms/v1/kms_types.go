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

	kmspublicv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/kms/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +gdcloud:manifest:relevant=true,oc=kms,component=kms,entities="mz-aeadkeys"
// +gdcloud:manifest:verbs=delete
// +gdcloud:manifest:rbac="delete:global-mzaeadkey-deleter"
// +gdcloud:manifest:skipcodegen=true

// MZAEADKey represents the key material that needs to be replicated across multiple zones.
type MZAEADKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              kmspublicv1.AEADKeySpec `json:"spec,omitempty"`
	Status            MZAEADKeyStatus         `json:"status,omitempty"`
}

// MZAEADKeyStatus defines the observed state of MZAEADKey.
type MZAEADKeyStatus struct {
	// Provides the status for an MZAEADKey.
	kmspublicv1.AEADKeyStatus `json:",inline"`
	// Zones represent the status and details of key replication in each zone.
	Zones       []MZAEADKeyZoneStatus `json:"zones,omitempty"`
	PrimaryZone string                `json:"primaryZone,omitempty"`
}

// MZAEADKeyZoneStatus represents the replication status of a specific zone.
type MZAEADKeyZoneStatus struct {
	ZoneName string `json:"zoneName,omitempty"`
	// Indicates whether the zone has successfully completed processing the key material or not.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func (a *MZAEADKey) Algorithm() kmspublicv1.AEADAlgorithm {
	return a.Spec.Algorithm
}

func (a *MZAEADKey) SetAlgorithm(algorithm kmspublicv1.AEADAlgorithm) {
	a.Spec.Algorithm = algorithm
}

func (a *MZAEADKey) Conditions() []metav1.Condition {
	return a.Status.Conditions
}

func (a *MZAEADKey) SetConditions(conditions []metav1.Condition) {
	a.Status.Conditions = conditions
}

// +kubebuilder:object:root=true

// MZAEADKeyList contains a list of MZAEADKey
type MZAEADKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MZAEADKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MZAEADKey{}, &MZAEADKeyList{})
}
