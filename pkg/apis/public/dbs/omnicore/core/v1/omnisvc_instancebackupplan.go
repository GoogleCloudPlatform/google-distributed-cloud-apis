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

// +kubebuilder:object:generate=true

// OmniSvcInstanceBackupPlanSpec defines the desired state of an AlloyDB Omni InstanceBackupPlan that are specific to OmniSvc based product.
type OmniSvcInstanceBackupPlanSpec struct {
	// BackupLocation specifies the remote object storage location to store backups.
	// For example, specs to a GCS buckets.
	// Without specifying this, backups are stored in the backup disk by default.
	// +kubebuilder:validation:Optional
	BackupLocation *StorageSpec `json:"backupLocation,omitempty"`
}

type OmniSvcInstanceBackupPlan interface {
	InstanceBackupPlan
	OmniSvcInstanceBackupPlanSpec() *OmniSvcInstanceBackupPlanSpec
}

type OmniSvcInstanceBackupPlanList interface {
	InstanceBackupPlanList
	OmniSvcInstanceBackupPlanListItem() []OmniSvcInstanceBackupPlan
}
