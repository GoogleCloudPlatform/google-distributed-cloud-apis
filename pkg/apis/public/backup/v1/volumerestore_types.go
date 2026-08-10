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

// Defines the desired state of a volume restore.
type VolumeRestoreSpec struct {
	// The name of the cluster that contains the data from this volume restore.
	// By default, it is the cluster in which this volume restore is created.
	// +optional
	ClusterName string `json:"clusterName,omitempty" reflect:"unexport"`

	// The name of the restore resource that created this volume restore.
	// +kubebuilder:validation:Required
	RestoreName string `json:"restoreName" reflect:"unexport"`
	// The name of the volume backup resource that we are restoring.
	// +kubebuilder:validation:Required
	VolumeBackupName string `json:"volumeBackupName" reflect:"unexport"`
	// The target `PersistentVolumeClaim` resource to be restored.
	// +kubebuilder:validation:Required
	TargetPvc NamespacedName `json:"targetPvc" reflect:"unexport"`
}

// +kubebuilder:validation:Enum=Unspecified;Creating;InProgress;Succeeded;Failed;Deleting
type VolumeRestoreState string

const (
	// Unspecified, default value
	VolumeRestoreStateUnspecified VolumeRestoreState = "Unspecified"
	// A volume for the restore was identified and restore process is about to
	// start.
	VolumeRestoreStateCreating VolumeRestoreState = "Creating"
	// A volume is being restored.
	VolumeRestoreStateInProgress VolumeRestoreState = "InProgress"
	// A volume has been restored.
	VolumeRestoreStateSucceeded VolumeRestoreState = "Succeeded"
	// A volume restoration failed.
	VolumeRestoreStateFailed VolumeRestoreState = "Failed"
	// A volume restoration is being deleted
	VolumeRestoreStateDeleting VolumeRestoreState = "Deleting"
)

const (
	VolumeRestoreConditionSucceeded = "Succeeded"
)

// Defines the observed state of a volume restore.
type VolumeRestoreStatus struct {
	// An underlying volume backup handle, which uniquely
	// identifies a volume backup inside of a volume backup repository. This handle
	// doesn't have a unified format and is treated as an opaque string.
	// +optional
	VolumeBackupHandle string `json:"volumeBackupHandle,omitempty" reflect:"unexport"`
	// The current state of the volume restoration.
	// +optional
	State VolumeRestoreState `json:"state,omitempty" reflect:"unexport"`
	// A human readable description of the current state.
	// +optional
	StateMessage string `json:"stateMessage,omitempty" reflect:"unexport"`
	// Conditions represents the observations of this Volume Restore's current state.
	// Known condition types: Ready.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" reflect:"unexport"`

	///// P2 Fields, Do Not Need To Be Implemented for V0 /////

	// The end time of the volume restore process.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty" reflect:"unexport"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +gdcloud:manifest:relevant=false,oc=back
// +genclient
// Defines the schema for the `VolumeRestore` API.
type VolumeRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeRestoreSpec   `json:"spec,omitempty"`
	Status VolumeRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `VolumeRestore` resources.
type VolumeRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeRestore{}, &VolumeRestoreList{})
}
