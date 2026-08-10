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

// Defines the desired state of a `ManualRestoreRequest` resource.
type ManualRestoreRequestSpec struct {
	// The name of the manual restore to be created.
	// +kubebuilder:validation:Required
	RestoreName string `json:"restoreName"`

	// The name of the restore plan to pull the `RestoreConfig` resource from.
	// +kubebuilder:validation:Required
	RestorePlanName string `json:"restorePlanName"`

	// The name of the backup that is being restored.
	// +kubebuilder:validation:Required
	BackupName string `json:"backupName"`

	// Note, it is expected that the restore plan and backup exist in the same
	// namespace as the `ManualRestoreRequest` resource.

	// A user-specified descriptive string for the restore created by this `ManualRestoreRequest` resource.
	// +optional
	Description string `json:"description"`

	// Filter which can be used to further refine the resource selection of the Restore beyond the coarse-grained scope defined in the RestorePlan.
	// +optional
	Filter *Filter `json:"filter,omitempty"`
}

// Defines the observed state of `ManualRestoreRequest` resource.
type ManualRestoreRequestStatus struct {
	// The time when the resource should expire.
	TimeToExpire metav1.Time `json:"timeToExpire"`

	// The status of the observed state of a `ManualRestoreRequest` resource.
	StatusField StatusFields `json:"statusField"`
}

// +gdcloud:manifest:relevant=false,oc=back
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
//
// +genclient
// Defines the schema for the `ManualRestoreRequest` API.
type ManualRestoreRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManualRestoreRequestSpec   `json:"spec,omitempty"`
	Status ManualRestoreRequestStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of `ManualRestoreRequest` resources.
type ManualRestoreRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManualRestoreRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ManualRestoreRequest{}, &ManualRestoreRequestList{})
}
