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

// InstanceBackupPhase is the phase of a backup.
type InstanceBackupPhase string

type DBClusterRef string

type InstanceBackupPlanRef string

const (
	InstanceBackupPhaseUnspecified InstanceBackupPhase = "Unspecified"
	InstanceBackupPhaseInProgress  InstanceBackupPhase = "InProgress"
	InstanceBackupPhaseCreating    InstanceBackupPhase = "Creating"
	InstanceBackupPhaseSucceeded   InstanceBackupPhase = "Succeeded"
	InstanceBackupPhaseFailed      InstanceBackupPhase = "Failed"
	InstanceBackupPhaseDeleting    InstanceBackupPhase = "Deleting"
)

type BackupType string

const (
	BackupTypeFull BackupType = "full"
	BackupTypeDiff BackupType = "diff"
	BackupTypeIncr BackupType = "incr"

	// This annotation is used on instancebackup resources to
	// indicate whether to skip uploading backup label
	BackupLabelUploadAnnotation     = "backup.internal.dbadmin.goog/bypass-backup-label-uploading"
	BackupLabelUploadAnnotationSkip = "true"
)

const (
	// This annotation is used on instancebackup/instancebackupplan resources to
	// indicate how to handle underlying backup data when resource is deleted.
	BackupDataAnnotation = "backup.internal.dbadmin.goog/backup-data"
	BackupDataActionKeep = "keep"
	// Backup Source Role
	BackupSourceRolePrimary BackupSourceRole = "primary"
	BackupSourceRoleStandby BackupSourceRole = "standby"
	// Annotation for selected instance during the remote backup
	SelectedBackupSourceInstanceAnnotation = "instancebackup.internal.dbadmin.goog/selected-backup-source-instance"
)

// BackupSourceRole defines which database instance(s) should perform the backup.
type BackupSourceRole string

// +kubebuilder:object:generate=true

// InstanceBackupSpec defines the desired state of InstanceBackup
type InstanceBackupSpec struct {
	// The DBCluster this backup belongs to
	// +kubebuilder:validation:Required
	DBClusterRef DBClusterRef `json:"dbclusterRef,omitempty"`

	// Name of the InstanceBackupPlan from which this backup was created.
	// +kubebuilder:validation:Required
	InstanceBackupPlanRef InstanceBackupPlanRef `json:"instanceBackupPlanRef,omitempty"`

	// Indicate whether this backup is a scheduled or manual backup
	// +optional
	// +kubebuilder:default:=false
	Manual bool `json:"manual,omitempty"`
}

// +kubebuilder:object:generate=true

// InstanceBackupStatus defines the observed state of Backup.
type InstanceBackupStatus struct {
	EntityStatus `json:",inline"`
	Phase        InstanceBackupPhase `json:"phase,omitempty"`
	// Completion time of the Backup
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
	// Create time of the underlying Backup.
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`
}

type InstanceBackup interface {
	client.Object
	EntityStatus() *EntityStatus
	InstanceBackupSpec() *InstanceBackupSpec
	InstanceBackupStatus() *InstanceBackupStatus
}

type InstanceBackupList interface {
	client.ObjectList
	InstanceBackupListItem() []InstanceBackup
}

func KeepBackupData(obj metav1.Object) bool {
	return obj.GetAnnotations() != nil && obj.GetAnnotations()[BackupDataAnnotation] == BackupDataActionKeep
}
