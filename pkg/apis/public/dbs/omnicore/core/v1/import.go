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

type ImportPhase string

const (
	ImportUnspecified ImportPhase = ""
	ImportPending     ImportPhase = "ImportPending"
	ImportInProgress  ImportPhase = "ImportInProgress"
	ImportComplete    ImportPhase = "ImportComplete"
	ImportFailed      ImportPhase = "ImportFailed"

	LabelImportDBCluster = "import.dbadmin.goog/dbcluster"
)

const (
	// ConditionTypeImportComplete indicates the import has completed its execution.
	ConditionTypeImportComplete ConditionType = "ImportComplete"

	// ConditionReason
	ConditionReasonImportInitializing ConditionReason = "ImportInitializing"
	ConditionReasonImportInProgress   ConditionReason = "ImportInProgress"
	ConditionReasonImportFailed       ConditionReason = "ImportFailed"
	ConditionReasonImportCompleted    ConditionReason = "ImportCompleted"
)

//+kubebuilder:object:generate=true

// CommonImportSpec defines the desired state of the db agnostic import spec.
type CommonImportSpec struct {
	// DatabaseName is the database resource name within Instance to import into.
	// +required
	DatabaseName string `json:"databaseName,omitempty"`

	// DumpStorage specifies a storage location for the import dump files.
	// A user is to ensure proper read access to the storage bucket from within the Operator.
	// +required
	DumpStorage *StorageSpec `json:"dumpStorage"`

	// LogStorage optionally specifies a storage location to copy import log to.
	// A user is to ensure proper write access to the storage bucket from within the
	// Operator.
	// +optional
	LogStorage *StorageSpec `json:"logStorage,omitempty"`

	// DownloadOnly when set to true means dump file will be downloaded but not imported into DB. Default is false.
	// +optional
	// +kubebuilder:default=false
	DownloadOnly bool `json:"downloadOnly,omitempty"`
}

//+kubebuilder:object:generate=true

// ImportStatus defines the observed state of Import.
type ImportStatus struct {
	EntityStatus `json:",inline"`

	// Phase is a summary of current state of the import.
	// +optional
	Phase ImportPhase `json:"phase,omitempty"`

	// StartTime is the time import started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CompleteTime is the time import completed.
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`

	// DumpPath is the path of the downloaded dump file for download only import.
	// +optional
	DumpPath string `json:"dumpPath,omitempty"`
}

//+kubebuilder:object:generate=true

// ImportSpec defines the common spec of an instance import
type ImportSpec struct {
	// CommonImportSpec defines the desired state of the db agnostic import spec.
	// +required
	CommonImportSpec `json:",inline"`

	// InstanceRef is the instance to import into.
	// +required
	InstanceRef string `json:"instanceRef,omitempty"`
}

// Import represents the contract for the Anthos DB Operator compliant
// database Operator providers to abide by.
type Import interface {
	Entity
	ImportSpec() ImportSpec
	ImportStatus() *ImportStatus
}

type ImportList interface {
	client.ObjectList
	ImportListItem() []Import
}
