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

	eemiapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/migration/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

// +kubebuilder:object:generate=true

// MigrationSpec defines the spec of the migration job.
type MigrationSpec struct {
	eemiapi.MigrationSpec `json:",inline"`
}

// MigrationStatus defines the status of the migration job.
type MigrationStatus struct {
	eemiapi.MigrationStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// Migration is the Schema for the migration API.
type Migration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MigrationSpec   `json:"spec,omitempty"`
	Status MigrationStatus `json:"status,omitempty"`
}

func (in *Migration) EntityStatus() *occoreapi.EntityStatus {
	return &in.Status.EntityStatus
}

func (in *Migration) MigrationSpec() eemiapi.MigrationSpec {
	return in.Spec.MigrationSpec
}

func (in *Migration) MigrationStatus() *eemiapi.MigrationStatus {
	return &in.Status.MigrationStatus
}

//+kubebuilder:object:root=true

type MigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Migration `json:"items"`
}

func (in *MigrationList) MigrationListItem() []eemiapi.Migration {
	var items []eemiapi.Migration
	for i := range in.Items {
		items = append(items, &in.Items[i])
	}
	return items
}

func init() {
	SchemeBuilder.Register(&Migration{}, &MigrationList{})
}

var _ eemiapi.MigrationList = &MigrationList{}
