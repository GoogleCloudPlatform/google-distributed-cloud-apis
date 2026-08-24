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

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +genclient
// Represents a replicated organization resources that will be synced to a
// particular zonal API server.
// An organization resource will have a replica for each zone. Upon an update of
// the organization resource, the replicas will be progressively updated based
// on the resource's rollout strategy.
type ProjectReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec          `json:"spec,omitempty"`
	Status ProjectReplicaStatus `json:"status,omitempty"`
}

var _ v1alpha1.MuxReplicaInterface = &ProjectReplica{}

func (r *ProjectReplica) GetSpec() any {
	return v1alpha1.NewTypedFieldAccessor(&r.Spec).GetField()
}

func (r *ProjectReplica) SetSpec(spec any) error {
	return v1alpha1.NewTypedFieldAccessor(&r.Spec).SetField(spec)
}

func (r *ProjectReplica) GetStatus() any {
	return v1alpha1.NewTypedFieldAccessor(&r.Status).GetField()
}

func (r *ProjectReplica) SetStatus(status any) error {
	return v1alpha1.NewTypedFieldAccessor(&r.Status).SetField(status)
}

// +kubebuilder:object:root=true

// Represents a collection of projects.
type ProjectReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ProjectReplica `json:"items"`
}

// Provides the specification (i.e., desired state) of a project.
type ProjectSpec struct {
}

// Provides the status of a project replica.
type ProjectReplicaStatus struct {
	// Conditions represents the observations of this project's overall
	// state.
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// AvailableClusters represents the amount of available user clusters.
	AvailableClusters int64 `json:"availableClusters,omitempty"`

	// ErrorStatus contain a list of current errors and the timestamp this field
	// gets updated.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ProjectReplica{},
		&ProjectReplicaList{},
	)
}
