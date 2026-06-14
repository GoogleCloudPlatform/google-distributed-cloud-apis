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
)

// Represents an inner message type that defines the configuration of creating
// a backup from this backup plan.
type ClusterBackupConfig struct {
	// The resource selection scope of a backup. Examples include
	// `all_namespaces`, selected namespaces, and selected applications.
	// You must specify a single value for `backup_scope`. The `BackupScope` value must be one of the following types:
	// `BackupConfig_AllNamespaces`, `BackupConfig_SelectedNamespaces`, or `BackupConfig_SelectedApplications`.
	// +kubebuilder:validation:Required
	BackupScope BackupScope `json:"backupScope"`

	// The name of the `BackupRepositoryRef` resource identifying the secondary storage for this `ClusterBackupPlan` resource.
	// +optional
	BackupRepositoryRef string `json:"backupRepositoryName,omitempty"`

	// The name of the `ClusterBackupRepositoryRef` resource identifying the secondary storage for this `ClusterBackupPlan` resource.
	// +optional
	ClusterBackupRepositoryRef string `json:"clusterBackupRepositoryName,omitempty"`

	// Specifies whether volume data is backed up.
	// If unset, the default is `false`.
	// +optional
	IncludeVolumeData bool `json:"includeVolumeData,omitempty"`

	// Specifies whether secrets are backed up.
	// If unset, the default is `false`.
	// +optional
	IncludeSecrets bool `json:"includeSecrets,omitempty"`

	// The type of volume backup to perform.
	// +optional
	VolumeStrategy *VolumeStrategy `json:"volumeStrategy,omitempty"`
}

// Represents an API that wraps around the BackupPlan custom resource.
// They are mostly identical, but there are some fields that are selectively omitted.
type ClusterBackupPlanSpec struct {
	// The cluster that will backed up.
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf)
	// +kubebuilder:validation:Required
	TargetCluster TargetCluster `json:"targetCluster"`

	// The scheduled backup creation under this backup plan.
	// +kubebuilder:validation:Required
	BackupSchedule Schedule `json:"backupSchedule"`
	// The backup configuration of this backup plan.
	// +kubebuilder:validation:Required
	ClusterBackupConfig ClusterBackupConfig `json:"clusterBackupConfig"`

	// The lifecycle of backups created under this plan.
	// +optional
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy"`
	// A user-specified descriptive string for this backup plan.
	// +optional
	Description string `json:"description,omitempty"`
	// Specifies whether the plan has been deactivated.
	// Setting this field to ‘true’ locks the plan meaning no further updates
	// are allowed, including changes to the deactivated field. It also prevents new
	// backups from being created under this plan, both manually or scheduled.
	// Default to ‘false’.
	// +optional
	// +kubebuilder:default:=false
	Deactivated bool `json:"deactivated"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="LastBackupTime",type="string",JSONPath=".status.lastBackupTime"
// +kubebuilder:printcolumn:name="LastBackupState",type="string",JSONPath=".status.lastBackupState"
// +kubebuilder:printcolumn:name="NextBackupTime",type="string",JSONPath=".status.nextBackupTime"
// +kubebuilder:printcolumn:name="Paused",type="boolean",JSONPath=".spec.backupSchedule.paused"
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-backup-plans"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:organization-cluster-backup-admin"
// +gdcloud:manifest:skipcodegen=true
type ClusterBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterBackupPlanSpec `json:"spec,omitempty"`
	Status            BackupPlanStatus      `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Represents a list of ClusterBackups
type ClusterBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterBackupPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ClusterBackupPlan{},
		&ClusterBackupPlanList{},
	)
}
