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

	"gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/common"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

// BackupPlanPhase is the phase of a BackupPlan.
type BackupPlanPhase string

const (
	BackupPlanPhaseUnspecified BackupPlanPhase = "Unspecified"
	BackupPlanPhaseReady       BackupPlanPhase = "Ready"
	BackupRepositoryAnnotation string          = "fleet.dbadmin.goog/backup-repository"
	BucketAnnotation           string          = "fleet.dbadmin.goog/bucket"
)

//+kubebuilder:object:generate=true

// BackupPlanSpec defines the desired state of BackupPlan
type BackupPlanSpec struct {
	// The DBCluster name this backupplan configures. This field is required and immutable.
	//
	// +kubebuilder:validation:Required
	DBClusterRef common.DBClusterRef `json:"dbclusterRef,omitempty"`

	// Number of days after which the service will delete a Backup.
	// If specified, a Backup created under this BackupPlan will be
	// automatically deleted after its age reaches create_time +
	// backup_retain_days.
	// The valid values are from 1 to 90 days.
	// Default to 14 retain days if not specified.
	//
	// +optional
	// +kubebuilder:default:=14
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=90
	BackupRetainDays int `json:"backupRetainDays"`

	// A flag to indicate if the backup creation under this BackupPlan is paused.
	// If set to true, the service will pause the scheduling of new Backups under
	// this BackupPlan.
	// Default to False.
	//
	// +optional
	// +kubebuilder:default:=false
	Paused bool `json:"paused"`

	// A flag to indicate whether logs replication is enabled to support point-in-time recovery.
	// Default to False.
	//
	// nullon(samwise-fleet)
	// +optional
	// +kubebuilder:default:=false
	PITREnabled bool `json:"PITREnabled,omitempty"`
}

//+kubebuilder:object:generate=true

// BackupPlanStatus defines the observed state of BackupPlan.
type BackupPlanStatus struct {
	occoreapi.EntityStatus `json:",inline"`
	Phase                  BackupPlanPhase `json:"phase,omitempty"`

	// LastBackupTime is the timestamp for the most recently executed backup.
	// +optional
	LastBackupTime metav1.Time `json:"lastBackupTime,omitempty"`

	// NextBackupTime is the timestamp for the next scheduled backup.
	// +optional
	NextBackupTime metav1.Time `json:"nextBackupTime,omitempty"`

	// RecoveryWindow is the currently available recovery window.
	// +optional
	RecoveryWindow *occoreapi.TimeWindow `json:"recoveryWindow,omitempty"`

	// ReadWriteBackupRepositoryZone is the currently availability zone of
	// the backup repository with ReadWrite import policy.
	// nullon(samwise-fleet)
	// +optional
	ReadWriteBackupRepositoryZone string `json:"readWriteBackupRepositoryZone,omitempty"`
}

type BackupPlan interface {
	occoreapi.Entity
	BackupPlanSpec() *BackupPlanSpec
	BackupPlanStatus() *BackupPlanStatus
}

type BackupPlanList interface {
	ctrlclient.ObjectList
	BackupPlanListItems() []BackupPlan
}
