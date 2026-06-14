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
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="CreateTime",type="string",JSONPath=".status.createTime"
// +kubebuilder:printcolumn:name="CompleteTime",type="string",JSONPath=".status.completeTime"
// +gdcloud:manifest:relevant=true,oc=haas,component=harbor
// +gdcloud:manifest:entities="restores",verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:harbor-instance-admin"
// +gdcloud:manifest:rbac="describe,list:harbor-instance-viewer"
// Defines the schema for the `Restore` API for HarborInstance.
type HarborInstanceRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the Restore.
	// +kubebuilder:validation:Required
	Spec HarborInstanceRestoreSpec `json:"spec"`
	// The most recently observed status of the Restore.
	Status HarborInstanceRestoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of Restore for HaaS instance.
type HarborInstanceRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HarborInstanceRestore `json:"items"`
}

// Defines the desired state of a restore..
type HarborInstanceRestoreSpec struct {
	// The full name of the Harbor instance backup resource that this Restore resource
	// uses to restore from.
	// +kubebuilder:validation:Required
	BackupName string `json:"backupName"`
	// An optional description of the backup. This has no impact on functionality.
	// +optional
	Description string `json:"description,omitempty"`
}

// Defines the observed state of a restore.
type HarborInstanceRestoreStatus struct {
	// TODO(b/353364160): consider also adding standard conditions, and use Kubernetes Job as a reference: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#job-v1-batch
	// The current state of the restore.
	// +optional
	State RestoreState `json:"state,omitempty"`
	// A human-readable description of why the restore is in the current state.
	// +optional
	StateReason string `json:"stateReason,omitempty"`
	// The create time of the restore process.
	// +optional
	CreateTime *metav1.Time `json:"startTime,omitempty"`
	// The end time of the restore process.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
}

// +kubebuilder:validation:Enum=Unspecified;Creating;InProgress;Waiting;Succeeded;Failed;Deleting
type RestoreState string

const (
	// Unspecified, default value
	RestoreStateUnspecified RestoreState = "Unspecified"
	// Restore resource has been validated and processing
	// Harbor Instance has been created and waiting for Instance reconciliation ready
	// in restoration mode
	RestoreStateCreating RestoreState = "Creating"
	// Restore is in progress
	RestoreStateInProgress RestoreState = "InProgress"
	// Restore process has been done, waiting for Instance reconciliation to be ready.
	RestoreStateWaiting RestoreState = "Waiting"
	// Restore succeeded
	RestoreStateSucceeded RestoreState = "Succeeded"
	// Restore failed
	RestoreStateFailed RestoreState = "Failed"
	// Restore is being deleted
	RestoreStateDeleting RestoreState = "Deleting"
)

func init() {
	SchemeBuilder.Register(
		&HarborInstanceRestore{},
		&HarborInstanceRestoreList{},
	)
}
