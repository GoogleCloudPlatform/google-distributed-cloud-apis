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

	eebackupplanapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/backupplan/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

const BackupPlanLabel = "alloydbomni.dbadmin.goog/backupplan"

// +kubebuilder:object:generate=true

// BackupPlanSpec defines the desired state of BackupPlan.
type BackupPlanSpec struct {
	eebackupplanapi.BackupPlanSpec `json:",inline"`
}

//+kubebuilder:object:generate=true

// BackupPlanStatus defines the observed state of BackupPlan.
type BackupPlanStatus struct {
	eebackupplanapi.BackupPlanStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion
//+kubebuilder:printcolumn:JSONPath=`.status.phase`,name="Phase",type="string"
//+kubebuilder:printcolumn:JSONPath=".status.lastBackupTime",name="LastBackupTime",type="string"
//+kubebuilder:printcolumn:JSONPath=".status.nextBackupTime",name="NextBackupTime",type="string"

// BackupPlan is the Schema for the backupplans API
type BackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupPlanSpec   `json:"spec,omitempty"`
	Status BackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPlanList contains a list of BackupPlan
type BackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPlan `json:"items"`
}

func (b *BackupPlan) EntityStatus() *occoreapi.EntityStatus {
	return &b.Status.EntityStatus
}

func (b *BackupPlan) BackupPlanSpec() *eebackupplanapi.BackupPlanSpec {
	return &b.Spec.BackupPlanSpec
}

func (b *BackupPlan) BackupPlanStatus() *eebackupplanapi.BackupPlanStatus {
	return &b.Status.BackupPlanStatus
}

func (b *BackupPlanList) BackupPlanListItems() []eebackupplanapi.BackupPlan {
	var backupItems []eebackupplanapi.BackupPlan
	for i := range b.Items {
		backupItems = append(backupItems, &b.Items[i])
	}
	return backupItems
}

var (
	_ eebackupplanapi.BackupPlan     = &BackupPlan{}
	_ eebackupplanapi.BackupPlanList = &BackupPlanList{}
)

func init() {
	SchemeBuilder.Register(&BackupPlan{}, &BackupPlanList{})
}
