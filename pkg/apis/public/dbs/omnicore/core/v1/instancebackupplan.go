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

// InstanceBackupPlanPhase is the phase of an InstanceBackupPlan.
type InstanceBackupPlanPhase string

// BackupSourceStrategy defines which strategy instance(s) should use to schedule the backup.
type BackupSourceStrategy string

const (
	InstanceBackupPlanPhaseUnspecified InstanceBackupPlanPhase = "Unspecified"
	InstanceBackupPlanPhaseReady       InstanceBackupPlanPhase = "Ready"
	InstanceBackupPlanPhaseInProgress  InstanceBackupPlanPhase = "InProgress"
	BackupPlanVisibilityInternal       string                  = "internal"
	BackupPlanVisibilityLabel          string                  = "dbadmin.goog/visibility"
	// PermanentDeletionLabel is to indicate the dbcluster is undergoing permanent deletion.
	PermanentDeletionLabel string = "dbadmin.goog/permanent-deletion"

	// Backup Source Strategy
	BackupSourceStrategyPrimary BackupSourceStrategy = "primary"
	BackupSourceStrategyStandby BackupSourceStrategy = "standby"
)

//+kubebuilder:object:generate=true

// InstanceBackupPlanSpec defines the desired state of InstanceBackupPlan
type InstanceBackupPlanSpec struct {
	// The DBCluster this backup plan configures.
	// +kubebuilder:validation:Required
	DBClusterRef DBClusterRef `json:"dbclusterRef,omitempty"`

	// Number of days after which the service will delete an InstanceBackup.
	// If specified, an InstanceBackup created under this InstanceBackupPlan
	// will be automatically deleted after its age reaches create_time +
	// backup_retain_days. The valid values are from 1 to 90 days.
	// Default to 14 retain days.
	// +optional
	// +kubebuilder:default:=14
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=90
	BackupRetainDays uint `json:"backupRetainDays"`

	// BackupRepository is the name of the GDCH Backup BackupRepository resource identifying the secondary storage for this `InstanceBackupPlan`.
	// If not provided, the default "dbs-backup-repository" will be used.
	// nullon(samwise-fleet,samwise-local)
	// +optional
	BackupRepository string `json:"backupRepository,omitempty"`

	// A flag to indicate if the backup creation under this plan is paused.
	// If set to true, the service will pause the scheduling of new
	// InstanceBackups under this InstanceBackupPlan.
	// Default to False.
	// +optional
	// +kubebuilder:default:=false
	Paused bool `json:"paused"`

	// A flag to indicate whether logs replication is enabled to support point-in-time recovery.
	// Default to False.
	// nullon(samwise-fleet,samwise-local)
	// +optional
	// +kubebuilder:default:=false
	PITREnabled bool `json:"PITREnabled"`

	// A flag to indicate if the instancebackupplan is in import mode.
	// If set to true, the service will passively import the existing platform backup resources
	// but will not creating new ones.
	// Default to False
	// nullon(samwise-fleet,samwise-local)
	// +optional
	// +kubebuilder:default:=false
	ImportMode bool `json:"importMode"`
}

//+kubebuilder:object:generate=true

// InstanceBackupPlanStatus defines the observed state of InstanceBackupPlan.
type InstanceBackupPlanStatus struct {
	EntityStatus `json:",inline"`
	Phase        InstanceBackupPlanPhase `json:"phase,omitempty"`
	// LastBackupTime is the timestamp for the most recently executed backup.
	// +optional
	// +nullable
	LastBackupTime metav1.Time `json:"lastBackupTime,omitempty"`

	// NextBackupTime is the timestamp for the next scheduled backup.
	// +optional
	// +nullable
	NextBackupTime metav1.Time `json:"nextBackupTime,omitempty"`

	// RecoveryWindow is the currently available recovery window.
	RecoveryWindow *TimeWindow `json:"recoveryWindow,omitempty"`
}

type InstanceBackupPlan interface {
	client.Object
	EntityStatus() *EntityStatus
	InstanceBackupPlanSpec() *InstanceBackupPlanSpec
	InstanceBackupPlanStatus() *InstanceBackupPlanStatus
}

type InstanceBackupPlanList interface {
	client.ObjectList
	InstanceBackupListItem() []InstanceBackupPlan
}

func IsInternalBackupPlan(ibp InstanceBackupPlan) bool {
	if ibp.GetLabels() == nil {
		return false
	}
	if visibility, ok := ibp.GetLabels()[BackupPlanVisibilityLabel]; ok && visibility == BackupPlanVisibilityInternal {
		return true
	}
	return false
}
