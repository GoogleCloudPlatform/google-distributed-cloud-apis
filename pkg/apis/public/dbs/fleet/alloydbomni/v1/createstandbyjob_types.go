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
	eeworkflowapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/workflow/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=`.spec.currentStep`,name="CurrentStep",type="string"
// +kubebuilder:printcolumn:JSONPath=`.spec.attempt`,name="Attempt",type="integer"
// +kubebuilder:printcolumn:JSONPath=`.spec.endTime`,name="EndTime",type="string"
// +kubebuilder:printcolumn:JSONPath=`.spec.cleanup`,name="Cleanup",type="boolean"
// CreateStandbyJob is an internal workflow tracking object. Users should not directly interact with this.
type CreateStandbyJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              eeworkflowapi.WorkflowSpec `json:"spec,omitempty"`
	Status            occoreapi.EntityStatus     `json:"status,omitempty"`
}

func (in *CreateStandbyJob) EntityStatus() *occoreapi.EntityStatus {
	return &in.Status
}

func (in *CreateStandbyJob) WorkflowSpec() *eeworkflowapi.WorkflowSpec {
	return &in.Spec
}

// +kubebuilder:object:root=true
// CreateStandbyJobList contains a list of CreateStandbyJobs
type CreateStandbyJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CreateStandbyJob `json:"items"`
}

func (j *CreateStandbyJobList) CreateStandbyJobListItem() []eeworkflowapi.CreateStandbyJob {
	var items []eeworkflowapi.CreateStandbyJob
	for i := range j.Items {
		items = append(items, &j.Items[i])
	}
	return items
}

func init() {
	SchemeBuilder.Register(&CreateStandbyJob{}, &CreateStandbyJobList{})
}

var (
	_ eeworkflowapi.CreateStandbyJob     = &CreateStandbyJob{}
	_ eeworkflowapi.CreateStandbyJobList = &CreateStandbyJobList{}
)
