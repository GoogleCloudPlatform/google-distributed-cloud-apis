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

	eebackupapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/backup/v1"
	fleetapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/fleet/fleet/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

const BackupDBClusterLabel = "backups.oracle.dbadmin.goog/dbcluster"

// +kubebuilder:object:generate=true

// BackupSpec defines the desired state of Backup.
type BackupSpec struct {
	eebackupapi.BackupSpec `json:",inline"`
}

//+kubebuilder:object:generate=true

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	eebackupapi.BackupStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="CompleteTime",type="string",JSONPath=".status.completeTime"

// Backup is the Schema for the backups API
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupSpec   `json:"spec,omitempty"`
	Status BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backup
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

func (in *BackupList) BackupListItem() []eebackupapi.Backup {
	var items []eebackupapi.Backup
	for i := range in.Items {
		items = append(items, &in.Items[i])
	}
	return items
}

func (b *Backup) EntityStatus() *occoreapi.EntityStatus {
	return &b.Status.EntityStatus
}

func (b *Backup) BackupSpec() *eebackupapi.BackupSpec {
	return &b.Spec.BackupSpec
}

func (b *Backup) BackupStatus() *eebackupapi.BackupStatus {
	return &b.Status.BackupStatus
}

func (b *Backup) DBEngineName() string {
	return string(fleetapi.Oracle)
}

var (
	_ eebackupapi.Backup     = &Backup{}
	_ eebackupapi.BackupList = &BackupList{}
)

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
