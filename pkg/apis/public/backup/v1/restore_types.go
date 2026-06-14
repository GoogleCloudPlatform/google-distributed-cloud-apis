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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// Defines the desired state of a restore.
type RestoreSpec struct {
	// The full name of the backup resource that this `Restore` resource uses to restore
	// from.
	// +kubebuilder:validation:Required
	BackupName string `json:"backupName" reflect:"unexport"`
	// The name of the restore plan from which this restore inherited its `RestoreConfig` resource.
	// +optional
	RestorePlanName string `json:"restorePlanName" reflect:"unexport"`

	// The name of the cluster that contains the data restored by this restore.
	// By default, it is the cluster in which the restore was created.
	// +optional
	ClusterName string `json:"clusterName,omitempty" reflect:"unexport"`

	// The configuration of the restore.
	// +kubebuilder:validation:Required
	RestoreConfig RestoreConfig `json:"restoreConfig" reflect:"unexport"`
	// An optional description of the backup. This has no impact on functionality.
	// +optional
	Description string `json:"description,omitempty" reflect:"unexport"`
	// Filter can be used to further refine the resource selection of the Restore beyond the coarse-grained scope defined in the RestorePlan.
	// +optional
	Filter *Filter `json:"filter,omitempty" reflect:"unexport"`
}

// Defines the observed state of a restore.
type RestoreStatus struct {
	// The current state of the restore.
	// +optional
	State RestoreState `json:"state,omitempty" reflect:"unexport"`
	// A human-readable description of why the restore is in the current state.
	// +optional
	StateReason string `json:"stateReason,omitempty" reflect:"unexport"`
	// The most recent errors with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
	// Conditions represents the observations of this restore's current state.
	// Known condition types: Succeeded
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" reflect:"unexport"`
	// Specifies whether a restore job has been created for this restore.
	JobCreated bool `json:"jobCreated,omitempty" reflect:"unexport"`
	// The number of resources restored in this restore action.
	// +optional
	ResourcesRestoredCount int64 `json:"resourcesRestoredCount,omitempty" reflect:"unexport"`
	// The number of resources excluded in this restore action.
	// +optional
	ResourcesExcludedCount int64 `json:"resourcesExcludedCount,omitempty" reflect:"unexport"`
	// The number of resources that failed to be restored in this restore action.
	// +optional
	ResourcesFailedCount int64 `json:"resourcesFailedCount,omitempty" reflect:"unexport"`
	// The number of volumes restored in this restore action.
	// +optional
	RestoredVolumesCount int64 `json:"restoredVolumesCount,omitempty" reflect:"unexport"`

	///// P2 Fields, Do Not Need To Be Implemented for V0 /////

	// The create time of the restore process.
	// +optional
	CreateTime *metav1.Time `json:"startTime,omitempty" reflect:"unexport"`

	// The end time of the restore process.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty" reflect:"unexport"`
}

// +kubebuilder:validation:Enum=Unspecified;Creating;InProgress;Succeeded;Failed;Deleting
type RestoreState string

const (
	// Unspecified, default value
	RestoreStateUnspecified RestoreState = "Unspecified"
	// Restore resource has been created
	RestoreStateCreating RestoreState = "Creating"
	// Restore is in progress
	RestoreStateInProgress RestoreState = "InProgress"
	// Restore succeeded
	RestoreStateSucceeded RestoreState = "Succeeded"
	// Restore failed
	RestoreStateFailed RestoreState = "Failed"
	// Restore is being deleted
	RestoreStateDeleting RestoreState = "Deleting"
)

const (
	RestoreConditionSucceeded = "Succeeded"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +gdcloud:manifest:relevant=false,oc=back
// +genclient
// Defines the schema for the `Restore` API.
type Restore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreSpec   `json:"spec,omitempty"`
	Status RestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `Restore` resources.
type RestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Restore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Restore{}, &RestoreList{})
}
