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

// Defines the configuration of a `ClusterRestore`.
type ClusterRestoreConfig struct {
	// The policy to use for volume data
	// restoration. Provides a default value of `NO_VOLUME_DATA_RESTORATION` if no value is specified.
	// +optional
	// +kubebuilder:default:=NoVolumeDataRestoration
	VolumeDataRestorePolicy *VolumeDataRestorePolicy `json:"volumeDataRestorePolicy,omitempty"`

	// The policy that resolves conflicts
	// when restoring cluster-scoped resources.
	// This request is invalid if this field has a value of `CLUSTER_RESOURCE_CONFLICT_POLICY_UNSPECIFIED` and
	// `cluster_resource_restore_scope` is specified.
	// +kubebuilder:validation:Required
	ClusterResourceConflictPolicy ClusterResourceConflictPolicy `json:"clusterResourceConflictPolicy,omitempty"`

	// The restoration mode to use for
	// namespaced resources.
	// The request is invalid if this field has a value of `NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED` and
	// `namespaced_resource_restore_scope` is specified.
	// +kubebuilder:validation:Required
	NamespacedResourceRestoreMode NamespacedResourceRestoreMode `json:"namespacedResourceRestoreMode,omitempty"`

	// The non-namespaced resources to be
	// restored.
	// If this field is not specified, no cluster resource is restored.
	// Note, even though `PersistentVolume` resources are non-namespaced, they are
	// handled separately. See the `VolumeDataRestorePolicy` resource for details. Specifying
	// a `PersistentVolume` `GroupKind` in this list does not determine whether
	// a `PersistentVolume` is restored.
	// +optional
	ClusterResources *ClusterResourceSelection `json:"clusterResources,omitempty"`

	// The specific namespaced resources to restore.
	// If defined, only the resources defined in this `allowlist` are restored.
	// +optional
	NamespacedResourceAllowlist []metav1.GroupKind `json:"namespacedResourceAllowlist,omitempty"`

	// The selected namespace resources
	// to restore. One of the entries must be specified along with a valid `Type`.
	//
	// The `Type` values that are valid to be assigned to `restoreScope` are
	// `AllNamespaces`, `SelectedNamespaces`, or `SelectedApplications`.
	// +optional
	NamespacedResourceRestoreScope *BackupScope `json:"namespacedResourceRestoreScope"`

	// The rules followed during the substitution of backed-up Kubernetes resources.
	// An empty list means no substitution will occur. Substitution rules are
	// applied sequentially in the order defined. This order matters, as changes
	// made by a rule may impact the matching logic of the subsequent rule.
	// Only one of `SubstitutionRules` or `TransformationRules` can be specified for a given restore operation.
	// +optional
	SubstitutionRules []SubstitutionRule `json:"substitutionRules,omitempty"`

	// The rules followed during the transformation of backed-up Kubernetes resources.
	// An empty list means no transformation will happen. Transformation rules are
	// applied sequentially in the order defined. This order matters, as changes
	// made by a rule may impact the matching logic of a subsequent rule.
	// Only one of `SubstitutionRules` or `TransformationRules` can be specified for a given restore operation.
	// +optional
	TransformationRules []TransformationRule `json:"transformationRules,omitempty"`

	// The name of the cluster backup repository which identifies the repository for the restore resource.
	// This field must be attached in read-write mode.
	// If this field is not supplied then it will be selected using the following logic:
	// 1. If the backup that we are performing the restore on points to a read-write repository in the current
	// cluster, this repository is selected.
	// 2. If the backup that we are performing a restore on points to a read-only repository, any
	// read-write repository from the cluster is selected and used.
	// +optional
	ClusterBackupRepositoryRef string `json:"clusterBackupRepositoryRef,omitempty"`
}

// Represents an API that wraps around the RestorePlan custom resource.
// They are mostly identical, but there are some fields that are selectively omitted. Defines the desired state of the 'ClusterRestorePlan'.
type ClusterRestorePlanSpec struct {
	// The cluster where data will be restored.
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf)
	// +kubebuilder:validation:Required
	TargetCluster TargetCluster `json:"targetCluster"`

	// The name of the cluster backup plan from which cluster backups
	// may be used as the source for cluster restores created using this `ClusterRestorePlan`. This field is required and immutable.
	// +kubebuilder:validation:Required
	ClusterBackupPlanName string `json:"clusterBackupPlanName"`

	// The cluster restore configuration of this cluster restore plan.
	// +kubebuilder:validation:Required
	ClusterRestoreConfig ClusterRestoreConfig `json:"clusterRestoreConfig"`

	// A user-specified descriptive string for this cluster restore plan.
	// +optional
	Description string `json:"description,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="LastRestoreTime",type="string",JSONPath=".status.lastRestoreTime"
// +genclient
// +gdcloud:manifest:relevant=true,oc=back,component=backup,entities="cluster-restore-plans"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:organization-cluster-backup-admin"
// Defines the schema for the `ClusterRestorePlan` API.
type ClusterRestorePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRestorePlanSpec `json:"spec,omitempty"`
	Status RestorePlanStatus      `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a list of `ClusterRestorePlan` resources.
type ClusterRestorePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRestorePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterRestorePlan{}, &ClusterRestorePlanList{})
}
