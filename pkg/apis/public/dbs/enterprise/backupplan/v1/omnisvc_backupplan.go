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

import occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"

//+kubebuilder:object:generate=true

// OmniSvcBackupPlanSpec defines the desired state of BackupPlan that are specific to OmniSvc based product.
type OmniSvcBackupPlanSpec struct {
	// TODO(b/481815226): This fields is removed from the Samwise API doc after the refactor. Invetigate the
	// root cause and add it back.

	// Defines the schedules for different types of backups.
	// Full, Differential and Incremental backup types are supported:
	// https://pgbackrest.org/user-guide.html#concept/backup
	// This field is optional.
	// By default, take one full backup every day at midnight.
	//
	// +kubebuilder:validation:Optional
	BackupSchedules *BackupSchedules `json:"backupSchedules,omitempty"`

	// BackupLocation specifies the remote object storage location to store backups.
	// For example, specs to a GCS buckets.
	// This field is optional.
	// By default, backups are stored in the backup disk.
	//
	// +kubebuilder:validation:Optional
	BackupLocation *occoreapi.StorageSpec `json:"backupLocation,omitempty"`

	// BackupSourceStrategy defines which strategy database instance(s) should use to perform the backup.
	// If not specified, defaults to primary.
	// +optional
	// +kubebuilder:default=primary
	// +kubebuilder:validation:Enum=primary;standby
	BackupSourceStrategy occoreapi.BackupSourceStrategy `json:"backupSourceStrategy,omitempty"`
}

//+kubebuilder:object:generate=true

// BackupSchedules defines the schedules for different types of backups.
type BackupSchedules struct {
	// Defines the Cron schedule for a full pgBackRest backup.
	// Follows the standard Cron schedule syntax:
	// https://k8s.io/docs/concepts/workloads/controllers/cron-jobs/#cron-schedule-syntax
	// +optional
	// +kubebuilder:validation:MinLength=6
	Full string `json:"full,omitempty"`

	// Defines the Cron schedule for a differential pgBackRest backup.
	// Follows the standard Cron schedule syntax:
	// https://k8s.io/docs/concepts/workloads/controllers/cron-jobs/#cron-schedule-syntax
	// +optional
	// +kubebuilder:validation:MinLength=6
	Differential string `json:"differential,omitempty"`

	// Defines the Cron schedule for an incremental pgBackRest backup.
	// Follows the standard Cron schedule syntax:
	// https://k8s.io/docs/concepts/workloads/controllers/cron-jobs/#cron-schedule-syntax
	// +optional
	// +kubebuilder:validation:MinLength=6
	Incremental string `json:"incremental,omitempty"`
}

type OmniSvcBackupPlan interface {
	BackupPlan
	OmniSvcBackupPlanSpec() *OmniSvcBackupPlanSpec
}

type OmniSvcBackupPlanList interface {
	BackupPlanList
	OmniSvcBackupPlanListItem() []OmniSvcBackupPlan
}
