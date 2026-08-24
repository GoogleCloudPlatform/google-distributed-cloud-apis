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

	eedbrapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/replication/v1"
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.dbcluster.name",name="DBCluster",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="Ready")].status`,name="Ready",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="Ready")].reason`,name="ReadyReason",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="Healthy")].status`,name="Healthy",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="Healthy")].reason`,name="HealthyReason",type="string"

// Replication is the Schema for the Replications API
type Replication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReplicationSpec   `json:"spec,omitempty"`
	Status ReplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:generate=true

// ReplicationSpec defines the desired state of Replication
type ReplicationSpec struct {
	eedbrapi.ReplicationSpec `json:",inline"`
}

// +kubebuilder:object:generate=true

// ReplicationStatus defines the observed state of Replication
type ReplicationStatus struct {
	eedbrapi.ReplicationStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// ReplicationList contains a list of Replication
type ReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Replication `json:"items"`
}

func (r *Replication) ReplicationSpec() *eedbrapi.ReplicationSpec {
	return &r.Spec.ReplicationSpec
}

func (r *Replication) ReplicationStatus() *eedbrapi.ReplicationStatus {
	return &r.Status.ReplicationStatus
}

func (r *Replication) EntityStatus() *occoreapi.EntityStatus {
	return &r.Status.ReplicationStatus.EntityStatus
}

// Replications returns a list of generic eedbcapi.Replication
func (rl *ReplicationList) Replications() (replications []eedbrapi.Replication) {
	for _, pool := range rl.Items {
		replications = append(replications, &pool)
	}
	return replications
}

func init() {
	SchemeBuilder.Register(&Replication{}, &ReplicationList{})
}

var (
	_ eedbrapi.Replication = &Replication{}
	_ occoreapi.Entity     = &Replication{}
)
