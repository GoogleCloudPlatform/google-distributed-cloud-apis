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

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Defines the desired state of the restore plan.
type RestorePlanSpec struct {
	// The name of the backup plan from which backups
	// may be used as the source for restores created using this restore plan. This field is required and immutable.
	// +kubebuilder:validation:Required
	BackupPlanName string `json:"backupPlanName" reflect:"unexport"`

	// The restore configuration of this restore plan.
	// +kubebuilder:validation:Required
	RestoreConfig RestoreConfig `json:"restoreConfig" reflect:"unexport"`

	// The name of the cluster that will contain the data restored by this restore plan.
	// By default, it is the cluster in which the restore plan is created.
	// +optional
	ClusterName string `json:"clusterName,omitempty" reflect:"unexport"`

	// A user-specified descriptive string for this restore plan.
	// +optional
	Description string `json:"description,omitempty" reflect:"unexport"`
}

// Defines the observed state of a restore plan.
type RestorePlanStatus struct {
	// The timestamp for the most recently executed restore.
	// +optional
	LastRestoreTime metav1.Time `json:"lastRestoreTime,omitempty"`

	// The current state of the restore plan.
	// +optional
	StatusField StatusFields `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="LastRestoreTime",type="string",JSONPath=".status.lastRestoreTime"
// +gdcloud:manifest:relevant=false,oc=back
// +genclient
// Defines the schema for the `RestorePlan` API.
type RestorePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestorePlanSpec   `json:"spec,omitempty"`
	Status RestorePlanStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of `RestorePlan` resources.
type RestorePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestorePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestorePlan{}, &RestorePlanList{})
}
