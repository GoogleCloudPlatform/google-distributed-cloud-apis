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

// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

// Service is an Oracle Operator provided service.
type Service string

// InstanceSpec defines the desired state of Instance.
type InstanceSpec struct {
	// InstanceSpec represents the database engine agnostic
	// part of the spec describing the desired state of an Instance.
	occoreapi.InstanceSpec `json:",inline"`

	// CDBName is the intended name of the CDB attribute. If the CDBName is
	// different from the original name (with which the CDB was created) the
	// CDB will be renamed.
	// +optional
	CDBName string `json:"cdbName,omitempty"`

	// CharacterSet used to create a database (the default is AL32UTF8).
	// +optional
	CharacterSet string `json:"characterSet,omitempty"`

	// MemoryPercent represents the percentage of memory that should be allocated
	// for Oracle SGA (default is 25%).
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	MemoryPercent int `json:"memoryPercent,omitempty"`

	// EnableTLS enables TLS encrypted connection.
	EnableTLS bool `json:"EnableTLS,omitempty"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	// InstanceStatus represents the database engine agnostic
	// part of the status describing the observed state of an Instance.
	occoreapi.InstanceStatus `json:",inline"`

	// List of database names (e.g. PDBs) hosted in the Instance.
	DatabaseNames []string `json:"databasenames,omitempty"`

	// Last backup ID.
	BackupID string `json:"backupid,omitempty"`

	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	LastRestoreTime *metav1.Time `json:"lastRestoreTime,omitempty"`

	// LastDatabaseIncarnation stores the parent incarnation number
	LastDatabaseIncarnation string `json:"lastDatabaseIncarnation,omitempty"`

	// CurrentDatabaseIncarnation stores the current incarnation number
	CurrentDatabaseIncarnation string `json:"currentDatabaseIncarnation,omitempty"`

	// CurrentActiveStateMachine stores the name of the state machine currently active.
	CurrentActiveStateMachine string `json:"CurrentActiveStateMachine,omitempty"`

	// LockedByController is a shared lock field granting exclusive access
	// to maintenance operations to only one controller.
	// Empty value means unlocked.
	// Non-empty value contains the name of the owning controller.
	// +optional
	LockedByController string `json:"lockedBy,omitempty"`

	// SSLEnabled implies the SSL based connection is enabled.
	SSLEnabled bool `json:"SSLEnabled,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=genericinstances,shortName=ginst
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:JSONPath=".spec.type",name="DB Engine",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.version",name="Version",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.internalConnectivity.url",name="URL",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="DBDaemonReady")].status`,name="DBDReadyStatus",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="DBDaemonReady")].reason`,name="DBDReadyReason",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="DatabaseReady")].status`,name="DBReadyStatus",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="DatabaseReady")].reason`,name="DBReadyReason",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="DatabaseReadyDeprecated")].status`,name="DBReadyDStatus",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="DatabaseReadyDeprecated")].reason`,name="DBReadyDReason",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.isChangeApplied",name="IsChangeApplied",type="string",priority=1

// Instance is the Schema for the instances API.
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceSpec   `json:"spec,omitempty"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instance.
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
