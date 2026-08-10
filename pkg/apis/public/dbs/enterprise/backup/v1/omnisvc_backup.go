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
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:generate=true

// OmniSvcBackupSpec defines the desired state of Backup that are specific to OmniSvc based product.
type OmniSvcBackupSpec struct {
	// PhysicalBackupSpec contains spec for physical backups, allowing to specify backup type as an enum in {"full","diff","incr"}.
	// This field is optional.
	// Default to full if not specified.
	PhysicalBackupSpec occoreapi.PhysicalBackupSpec `json:"physicalBackupSpec,omitempty"`
	// Backup Source, allowing to specify backup source role as an enum in {"primary", "standby"}
	// This field is optional
	// Default to primary if not specified
	// +optional
	// +kubebuilder:default=primary
	// +kubebuilder:validation:Enum=primary;standby
	BackupSourceRole occoreapi.BackupSourceRole `json:"backupSourceRole,omitempty"`
}

//+kubebuilder:object:generate=true

// OmniSvcBackupStatus defines the observed state of Backup that are specific to OmniSvc based product.
type OmniSvcBackupStatus struct {
	// RetainExpireTime defines the time when the Backup will be
	// automatically deleted. It's an output only field calculated from
	// `create_time` + `retain_days`, and will be updated accordingly when the
	// `retain_days` field of a Backup has been updated.
	// +optional
	RetainExpireTime *metav1.Time `json:"retainExpireTime,omitempty" reflect:"unexport"`

	// PhysicalBackupStatus contains status info that are specific for physical backups.
	PhysicalBackupStatus occoreapi.PhysicalBackupStatus `json:"physicalBackupStatus,omitempty" reflect:"unexport"`
}

type OmniSvcBackup interface {
	Backup
	OmniSvcBackupSpec() *OmniSvcBackupSpec
	OmniSvcBackupStatus() *OmniSvcBackupStatus
}

type OmniSvcBackupList interface {
	BackupList
	OmniSvcBackupListItem() []OmniSvcBackup
}
