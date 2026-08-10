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
	eerestoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/restore/v1"
	fleetapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/fleet/fleet/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true

// RestoreSpec defines the desired state of Restore.
type RestoreSpec struct {
	eerestoreapi.RestoreSpec `json:",inline"`
}

//+kubebuilder:object:generate=true

// RestoreStatus defines the observed state of Restore.
type RestoreStatus struct {
	eerestoreapi.RestoreStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:name="CompleteTime",type="string",JSONPath=".status.completeTime"
// +kubebuilder:printcolumn:name="RestoredPointInTime",type="string",JSONPath=".status.restoredPointInTime"

// Restore is the Schema for the restores API
type Restore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreSpec   `json:"spec,omitempty"`
	Status RestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestoreList contains a list of Restores
type RestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Restore `json:"items"`
}

// EntityStatus returns the core entity status.
func (r *Restore) EntityStatus() *occoreapi.EntityStatus {
	return &r.Status.EntityStatus
}

// RestoreSpec returns the common restore spec.
func (r *Restore) RestoreSpec() *eerestoreapi.RestoreSpec {
	return &r.Spec.RestoreSpec
}

// RestoreStatus returns the common restore status.
func (r *Restore) RestoreStatus() *eerestoreapi.RestoreStatus {
	return &r.Status.RestoreStatus
}

func (r *Restore) DBEngineName() string {
	return string(fleetapi.AlloyDBOmni)
}

func (r *RestoreList) RestoreListItem() []eerestoreapi.Restore {
	var rItems []eerestoreapi.Restore
	for i := range r.Items {
		rItems = append(rItems, &r.Items[i])
	}
	return rItems
}

var (
	_ eerestoreapi.Restore     = &Restore{}
	_ eerestoreapi.RestoreList = &RestoreList{}
)

func init() {
	SchemeBuilder.Register(&Restore{}, &RestoreList{})
}
