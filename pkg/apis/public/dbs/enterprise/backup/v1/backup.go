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
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/common"
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

// BackupPhase is the phase of a backup.
type BackupPhase string

const (
	BackupPhaseUnspecified BackupPhase = "Unspecified"
	BackupPhaseCreating    BackupPhase = "Creating"
	BackupPhaseInProgress  BackupPhase = "InProgress"
	BackupPhaseSucceeded   BackupPhase = "Succeeded"
	BackupPhaseFailed      BackupPhase = "Failed"
	BackupPhaseDeleting    BackupPhase = "Deleting"
)

//+kubebuilder:object:generate=true

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	// The DBCluster name this backup belongs to.
	// This field is required.
	//
	// +kubebuilder:validation:Required
	DBClusterRef common.DBClusterRef `json:"dbclusterRef,omitempty"`

	// Name of the BackupPlan from which this backup was created.
	// This field is required.
	//
	// +kubebuilder:validation:Required
	BackupPlanRef common.BackupPlanRef `json:"backupPlanRef,omitempty"`

	// Indicate whether this backup is a scheduled or manual backup.
	// This field is optional.
	// Default to false (scheduled backup) if not specified.
	//
	// +optional
	// +kubebuilder:default:=false
	Manual bool `json:"manual,omitempty"`
}

//+kubebuilder:object:generate=true

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	occoreapi.EntityStatus `json:",inline"`
	Phase                  BackupPhase `json:"phase,omitempty"`

	// Creation time of the Backup
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`

	// Completion time of the Backup
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
}

type Backup interface {
	occoreapi.Entity
	BackupSpec() *BackupSpec
	BackupStatus() *BackupStatus
	DBEngineName() string
}

type BackupList interface {
	ctrlclient.ObjectList
	BackupListItem() []Backup
}
