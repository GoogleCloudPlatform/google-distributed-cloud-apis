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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true

// OmniSvcInstanceBackupSpec defines the desired state of an InstanceBackup that are specific to OmniSvc based product.
type OmniSvcInstanceBackupSpec struct {
	// PhysicalBackupSpec contains spec for physical backups.
	PhysicalBackupSpec PhysicalBackupSpec `json:"physicalBackupSpec,omitempty"`

	// Role of Backup Source, allowing to specify backup source role as an enum in {"primary", "standby"}
	// This field is optional
	// Default to primary if not specified
	// +optional
	// +kubebuilder:default=primary
	// +kubebuilder:validation:Enum=primary;standby
	BackupSourceRole BackupSourceRole `json:"backupSourceRole,omitempty"`
}

// +kubebuilder:object:generate=true

// OmniSvcInstanceBackupStatus defines the observed state of an InstanceBackup that are specific to OmniSvc based product.
type OmniSvcInstanceBackupStatus struct {
	// PhysicalBackupStatus contains status info that are specific for physical backups.
	PhysicalBackupStatus PhysicalBackupStatus `json:"physicalBackupStatus,omitempty" reflect:"unexport"`
}

// +kubebuilder:object:generate=true

// PhysicalBackupSpec describes the desired state of an InstanceBackup physical backup.
type PhysicalBackupSpec struct {
	// BackupType is the type of backup to be created. It's an enum in {"full","diff","incr"}.
	// Default to full if not specified.
	//
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum={"full","diff","incr"}
	// +kubebuilder:default="full"
	BackupType BackupType `json:"backupType,omitempty"`
}

// +kubebuilder:object:generate=true

// PhysicalBackupStatus describes the status that only applies to physical backups
type PhysicalBackupStatus struct {
	// BackupID is the unique identifier of the physical backup, as tracked by
	// pgBackRest. This ID is used to reference the backup for restore operations.
	// +optional
	BackupID string `json:"backupID,omitempty"`

	// PriorBackup identifies the preceding backup that this backup relies upon.
	// This is particularly relevant for incremental and differential backups.
	// +optional
	PriorBackup string `json:"priorBackup,omitempty"`

	// BackupType indicates the type of the backup performed. It can be "full",
	// "diff" (differential), or "incr" (incremental).
	// +optional
	BackupType BackupType `json:"backupType,omitempty"`

	// StartTime records the exact time when the pgBackRest backup process initiated.
	// +optional
	StartTime metav1.Time `json:"startTime,omitempty"`

	// EndTime records the exact time when the pgBackRest backup process completed.
	// +optional
	EndTime metav1.Time `json:"endTime,omitempty"`

	// WAL contains the names of the starting and ending Write-Ahead Log (WAL)
	// files associated with this backup. The WAL files are essential for
	// ensuring data consistency and enabling point-in-time recovery.
	// +optional
	WAL WAL `json:"wal,omitempty"`

	// LSN contains the starting and ending Log Sequence Numbers (LSN) for the
	// backup. The LSNs provide a precise point in the database's transaction
	// log, defining the exact scope of the data included in the backup and
	// ensuring accurate recovery.
	// +optional
	LSN LSN `json:"lsn,omitempty"`

	// Size provides detailed information about the various sizes related to the
	// backup, such as the total size of the database and the size of the
	// backup data itself.
	// +optional
	Size Size `json:"size,omitempty"`

	// Encrypted specifies whether the backup data is encrypted.
	// +optional
	Encrypted bool `json:"encrypted,omitempty"`

	// Compressed specifies whether the backup data is compressed to save space.
	// +optional
	Compressed bool `json:"compressed,omitempty"`

	// Error indicates if there were errors during the backup.
	// +optional
	Error bool `json:"error,omitempty"`
}

// +kubebuilder:object:generate=true

// WAL defines the start and stop Write-Ahead Log (WAL) file names for a backup.
// The WAL files contain a record of all changes made to the database, which is
// critical for data recovery and ensuring consistency.
type WAL struct {
	// Start specifies the name of the first WAL file required for the backup.
	// +optional
	Start string `json:"start,omitempty"`

	// Stop specifies the name of the last WAL file included in the backup.
	// +optional
	Stop string `json:"stop,omitempty"`
}

// +kubebuilder:object:generate=true

// LSN defines the start and stop Log Sequence Numbers (LSN) for a backup.
// The LSN provides a precise point in the database's transaction log, enabling
// accurate and reliable recovery to a specific moment in time.
type LSN struct {
	// Start specifies the starting LSN, marking the beginning of the data
	// included in the backup.
	// +optional
	Start string `json:"start,omitempty"`

	// Stop specifies the ending LSN, marking the end of the data included in
	// the backup.
	// +optional
	Stop string `json:"stop,omitempty"`
}

// +kubebuilder:object:generate=true

// Size defines the various sizes related to a backup, providing insight into
// the storage footprint and data volume. This information is valuable for
// monitoring and capacity planning.
type Size struct {
	// DatabaseSize is the full uncompressed size of the database to be backed up.
	// +optional
	DatabaseSize resource.Quantity `json:"databaseSize,omitempty"`

	// DatabaseBackupSize is the amount of data in the database to actually back up
	// (which will be the same for full backups).
	// +optional
	DatabaseBackupSize resource.Quantity `json:"databaseBackupSize,omitempty"`

	// BackupSize is the size of only the files in this backup (which will also be the
	// same for full backups).
	// +optional
	BackupSize resource.Quantity `json:"backupSize,omitempty"`

	// BackupSetSize is the size of all the files from this backup and any referenced backups
	// in the repository that are required to restore the database from this backup
	// +optional
	BackupSetSize resource.Quantity `json:"backupSetSize,omitempty"`
}

type OmniSvcInstanceBackup interface {
	InstanceBackup
	OmniSvcInstanceBackupSpec() *OmniSvcInstanceBackupSpec
	OmniSvcInstanceBackupStatus() *OmniSvcInstanceBackupStatus
}

type OmniSvcInstanceBackupList interface {
	InstanceBackupList
	OmniSvcInstanceBackupListItem() []OmniSvcInstanceBackup
}
