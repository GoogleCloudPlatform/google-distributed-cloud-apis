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

	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

// InstanceBackupPlanSpec defines the desired state of a Postgres InstanceBackupPlan.
type InstanceBackupPlanSpec struct {
	// InstanceBackupPlan specs that are common across all database engines.
	occoreapi.InstanceBackupPlanSpec `json:",inline"`
}

// InstanceBackupPlanStatus defines the observed state of a Postgres InstanceBackupPlan.
type InstanceBackupPlanStatus struct {
	// InstanceBackupPlan status that is common across all database engines.
	occoreapi.InstanceBackupPlanStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.lastBackupTime",name="LastBackupTime",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.nextBackupTime",name="NextBackupTime",type="string"

// InstanceBackupPlan is the Schema for the InstanceBackupPlan API
type InstanceBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceBackupPlanSpec   `json:"spec,omitempty"`
	Status InstanceBackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceBackupPlanList contains a list of Postgres InstanceBackupPlan
type InstanceBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceBackupPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceBackupPlan{}, &InstanceBackupPlanList{})
}
