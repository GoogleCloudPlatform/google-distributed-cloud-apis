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

	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

// InstanceBackupSpec defines the desired state of a Postgres InstanceBackup.
type InstanceBackupSpec struct {
	// InstanceBackup specs that are common across all database engines.
	occoreapi.InstanceBackupSpec `json:",inline"`
}

// InstanceBackupStatus defines the observed state of a Postgres InstanceBackup.
type InstanceBackupStatus struct {
	// InstanceBackup status that is common across all database engines.
	occoreapi.InstanceBackupStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.completeTime",name="CompleteTime",type="string"

// InstanceBackup is the Schema for the InstanceBackup API
type InstanceBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceBackupSpec   `json:"spec,omitempty"`
	Status InstanceBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceBackupList contains a list of Postgres InstanceBackup
type InstanceBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceBackup{}, &InstanceBackupList{})
}
