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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:object:root=true
// LROJob is an internal object that helps single-thread LRO jobs.
type LROJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec LROSpec `json:"spec,omitempty"`
}
type LROSpec struct {
	// LROName is the name for the LRO
	// +required
	LROName string `json:"lroName,omitempty"`
}

// +kubebuilder:object:root=true

// LROJobList contains a list of postgresql LROJobs.
type LROJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LROJob `json:"items"`
}

var _ client.ObjectList = &LROJobList{}

func init() {
	SchemeBuilder.Register(&LROJob{}, &LROJobList{})
}
