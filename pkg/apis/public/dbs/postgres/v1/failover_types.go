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
	eehaapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/ha/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// FailoverSpec defines the desired state of postgresql Failover.
type FailoverSpec struct {
	// FailoverSpec includes failover specs common across all database engines.
	eehaapi.FailoverSpec `json:",inline"`
}

// FailoverStatus defines the observed state of postgresql Failover.
type FailoverStatus struct {
	// Failover status that is common across all database engines.
	eehaapi.FailoverStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
//+kubebuilder:printcolumn:JSONPath=`.status.state`,name="state",type="string"
//+kubebuilder:printcolumn:JSONPath=".status.internal.phase",name="phase",type="string"
//+kubebuilder:printcolumn:JSONPath=".spec.dbclusterRef",name="dbcluster",type="string"

// Failover is the Schema for the failover API.
type Failover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FailoverSpec   `json:"spec,omitempty"`
	Status FailoverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FailoverList contains a list of postgresql Failovers.
type FailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Failover `json:"items"`
}

var _ client.ObjectList = &FailoverList{}

func init() {
	SchemeBuilder.Register(&Failover{}, &FailoverList{})
}
