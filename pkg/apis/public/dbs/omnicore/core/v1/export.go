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

type ExportPhase string

const (
	ExportUnspecified ExportPhase = ""
	ExportPending     ExportPhase = "ExportPending"
	ExportInProgress  ExportPhase = "ExportInProgress"
	ExportComplete    ExportPhase = "ExportComplete"
	ExportFailed      ExportPhase = "ExportFailed"

	LabelExportDBCluster = "export.dbadmin.goog/dbcluster"
)

const (
	// ConditionTypeExportComplete indicates the export has completed its execution.
	ConditionTypeExportComplete ConditionType = "ExportComplete"

	// ConditionReason
	ConditionReasonFailedToEnsureBucketAccess ConditionReason = "FailedToEnsureBucketAccess"
	ConditionReasonFailureInLockAcquisition   ConditionReason = "FailureInLockAcquisition"
	ConditionReasonOtherOperationInProgress   ConditionReason = "OtherOperationInProgress"
	ConditionReasonInstanceNotReady           ConditionReason = "InstanceNotReady"
	ConditionReasonExportInitializing         ConditionReason = "ExportInitializing"
	ConditionReasonExportInProgress           ConditionReason = "ExportInProgress"
	ConditionReasonExportFailed               ConditionReason = "ExportFailed"
	ConditionReasonExportCompleted            ConditionReason = "ExportCompleted"
)

//+kubebuilder:object:generate=true

// CommonExportSpec defines the desired state of the db agnostic export spec.
type CommonExportSpec struct {
	// ExportLocation specifies a storage location for the export files.
	// A user is to ensure proper write access to the storage bucket from within the Operator.
	// +required
	ExportLocation *StorageSpec `json:"exportLocation"`
}

//+kubebuilder:object:generate=true

// ExportStatus defines the observed state of Export.
type ExportStatus struct {
	EntityStatus `json:",inline"`

	// Phase is a summary of current state of the export.
	// +optional
	Phase ExportPhase `json:"phase,omitempty"`

	// StartTime is the time export started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CompleteTime is the time export completed.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`

	// ExportSubDirectory is the subdirectory appended to ExportLocation to store exported files.
	// +optional
	ExportSubDirectory *string `json:"exportSubDirectory,omitempty"`
}

//+kubebuilder:object:generate=true

// ExportSpec defines the common spec of an instance export.
type ExportSpec struct {
	// CommonExportSpec defines the desired state of the db agnostic export spec.
	// +required
	CommonExportSpec `json:",inline"`

	// InstanceRef is the instance to export into.
	// +required
	InstanceRef string `json:"instanceRef,omitempty"`
}

// Export represents the contract for the Anthos DB Operator compliant
// database Operator providers to abide by.
type Export interface {
	Entity
	ExportSpec() ExportSpec
	ExportStatus() *ExportStatus
}

type ExportList interface {
	client.ObjectList
	ExportListItem() []Export
}
