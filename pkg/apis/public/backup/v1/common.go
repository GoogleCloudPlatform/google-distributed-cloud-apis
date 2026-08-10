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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	APIVersion                  = Group + "/" + Version
	ImportedLabelName           = Group + "/imported-repository"
	ImportedLabelRepositoryType = Group + "/imported-repository-type"

	// The annotation for object storage backup configs.
	ObjectStorageBackupConfigKey = Group + "/objectstoragebackupconfig"
	// The annotation for object storage restore configs.
	ObjectStorageRestoreConfigKey = Group + "/objectstoragerestoreconfig"
	// The UID of the K8s cluster in which the resource was originally created.
	OriginClusterUIDAnnotation = Group + "/origin-cluster-uid"
	// The name of the Cluster whose workloads (aka API resources)  will be backed up
	WorkloadClusterNameLabel = Group + "/workload-cluster-name"
	// Used to group backups created by a BackupPlan. For VM Backups, this label does not contain the -cb or -vb suffix.
	BackupPlanNameLabel = Group + "/backup-plan-name"
	// Whether a Backup request is created manually or by a schedule
	BackupType                     = Group + "/backup-type"
	BackupRepositoryAnnotationName = "backup.gdc.goog/backup-repository-name"
)

var (
	CurrentClusterUID = types.UID("")
)

// TODO(b/200285234): This file is heavily patterned off of:
//   cs//depot/google3/blaze-out/genfiles/google/cloud/gkebackup/v1alpha1/common.pb.go
// Going forward, we should come up with a good process to keep these
// relatively in sync.

// Represents an inner message type that defines the configuration of creating
// a backup from this backup plan.
type BackupConfig struct {
	// The resource selection scope of a backup. Examples include
	// `all_namespaces`, selected namespaces, and selected applications.
	// You must specify a single value for `backup_scope`. The `BackupScope` value must be one of the following types:
	// `BackupConfig_AllNamespaces`, `BackupConfig_SelectedNamespaces`, or `BackupConfig_SelectedApplications`.
	// +kubebuilder:validation:Required
	BackupScope BackupScope `json:"backupScope" reflect:"unexport"`

	// The name of the `BackupRepository` resource identifying the secondary storage for this `BackupPlan` resource.
	BackupRepository string `json:"backupRepository,omitempty" reflect:"unexport"`

	// Specifies whether volume data is backed up.
	// If unset, the default is `False`.
	// +optional
	IncludeVolumeData bool `json:"includeVolumeData,omitempty" reflect:"unexport"`

	// Specifies whether secrets are backed up.
	// If unset, the default is `False`.
	// +optional
	IncludeSecrets bool `json:"includeSecrets,omitempty" reflect:"unexport"`
	///// P2 Fields, Do Not Need To Be Implemented for V0 /////

	// An encryption key. This field is immutable.
	// +optional
	EncryptionKey *EncryptionKey `json:"encryptionKey,omitempty" reflect:"unexport"`

	// The tradeoffs to use when backing up volumes.
	// +optional
	VolumeStrategy VolumeStrategy `json:"volumeStrategy,omitempty" reflect:"unexport"`
}

// The tradeoffs to use when backing up volumes.
// +kubebuilder:validation:Enum=ProvisionerSpecific;LocalSnapshotOnly;Portable
type VolumeStrategy string

const (
	ProvisionerSpecific VolumeStrategy = "ProvisionerSpecific"
	LocalSnapshotOnly   VolumeStrategy = "LocalSnapshotOnly"
	Portable            VolumeStrategy = "Portable"
)

// NOTE(apsherid): interface not supported, though is default way of expressing one-of's by go_proto, replaced with apiextensions.JSON, see here: https://yaqs.corp.google.com/eng/q/1598751479437459456

// Defines which namespaces and applications to back up.
type BackupScope struct {
	// Specifies whether all namespaces are backed up.
	// If set to `True`, indicates that all namespaces must be backed up.
	// +optional
	AllNamespaces bool `json:"allNamespaces,omitempty" reflect:"unexport"`
	// The list of namespaces to back up.
	// +optional
	SelectedNamespaces *Namespaces `json:"selectedNamespaces,omitempty" reflect:"unexport"`
	// The list of applications to back up.
	// +optional
	SelectedApplications *NamespacedNames `json:"selectedApplications,omitempty" reflect:"unexport"`
}

type Namespaces struct {
	// A list of namespaces.
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`
}

type NamespacedNames struct {
	// A list of namespaced names.
	// +optional
	NamespacedNames []*NamespacedName `json:"namespacedNames,omitempty"`
}

// Defines a name in a specific namespace.
type NamespacedName struct {
	// The namespace of the resource in Kubernetes.
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// The name of the resource in Kubernetes.
	// +optional
	Name string `json:"name,omitempty"`
}

// Defines an encryption key. For preview, supports Google Cloud Platform KMS only.
type EncryptionKey struct {
	// TODO(b/200285412): Swap out for GPC supporting Encryption key, this will not be a gcp encryption key, if any.

	// A Google KMS encryption key in the format:
	// `projects/<project>/locations/<location>/keyRings/<key-ring>/cryptoKeys/<key>`.
	// +optional
	GcpKmsEncryptionKey string `json:"gcpKMSEncryptionKey,omitempty"`
}

// Represents various states a resource can be in.
// +kubebuilder:validation:Enum=Pending;Succeeded;Failed
type StatusFields string

const (
	StatusFieldPending StatusFields = "Pending"
	StatusFieldSuccess StatusFields = "Succeeded"
	StatusFieldFailed  StatusFields = "Failed"
)

// Defines the configuration of a restore.
type RestoreConfig struct {
	// The policy to use for volume data
	// restoration. Provides a default value of `NO_VOLUME_DATA_RESTORATION` if no value is specified.
	// +optional
	// +kubebuilder:default:=NoVolumeDataRestoration
	VolumeDataRestorePolicy *VolumeDataRestorePolicy `json:"volumeDataRestorePolicy,omitempty" reflect:"unexport"`
	// The policy that resolves conflicts
	// when restoring cluster-scoped resources.
	// The request is invalid if this field is not specified.
	// This request is invalid if this field has a value of `CLUSTER_RESOURCE_CONFLICT_POLICY_UNSPECIFIED` and
	// `cluster_resource_restore_scope` is specified.
	// +optional
	ClusterResourceConflictPolicy *ClusterResourceConflictPolicy `json:"clusterResourceConflictPolicy,omitempty" reflect:"unexport"`
	// The restoration mode to use for
	// namespaced resources.
	// The request is invalid if this field is not specified.
	// The request is invalid if this field has a value of `NAMESPACED_RESOURCE_RESTORE_MODE_UNSPECIFIED` and
	// `namespaced_resource_restore_scope` is specified.
	// +optional
	NamespacedResourceRestoreMode *NamespacedResourceRestoreMode `json:"namespacedResourceRestoreMode,omitempty" reflect:"unexport"`
	// The non-namespaced resources to be
	// restored.
	// If this field is not specified, no cluster resource is restored.
	// Note, even though `PersistentVolume` resources are non-namespaced, they are
	// handled separately. See the `VolumeDataRestorePolicy` resource for details. Specifying
	// a `PersistentVolume` `GroupKind` in this list does not determine whether
	// a `PersistentVolume` is restored.
	// +optional
	ClusterResources *ClusterResourceSelection `json:"clusterResources,omitempty" reflect:"unexport"`
	// The specific namespaced resources to restore.
	// If defined, only the resources defined in this `allowlist` are restored.
	// +optional
	NamespacedResourceAllowlist []metav1.GroupKind `json:"namespacedResourceAllowlist,omitempty" reflect:"unexport"`
	// The selected namespace resources
	// to restore. One of the entries must be specified along with a valid `Type`.
	//
	// The `Type` values that are valid to be assigned to `restoreScope` are
	// `AllNamespaces`, `SelectedNamespaces`, or `SelectedApplications`.
	// +optional
	NamespacedResourceRestoreScope *BackupScope `json:"restoreScope" reflect:"unexport"`
	// The rules followed during the substitution of backed-up Kubernetes resources.
	// An empty list means no substitution will occur. Substitution rules are
	// applied sequentially in the order defined. This order matters, as changes
	// made by a rule may impact the matching logic of the subsequent rule.
	// Only one of `SubstitutionRules` or `TransformationRules` can be specified for a given restore operation.
	// +optional
	SubstitutionRules []SubstitutionRule `json:"substitutionRules,omitempty" reflect:"unexport"`
	// The rules followed during the transformation of backed-up Kubernetes resources.
	// An empty list means no transformation will happen. Transformation rules are
	// applied sequentially in the order defined. This order matters, as changes
	// made by a rule may impact the matching logic of a subsequent rule.
	// Only one of `SubstitutionRules` or `TransformationRules` can be specified for a given restore operation.
	// +optional
	TransformationRules []TransformationRule `json:"transformationRules,omitempty" reflect:"unexport"`
	// T name of the backup repository which identifies the repository for the restore resource.
	// This field must be attached in read-write mode.
	// Note, this backup repository might differ from the backup repository for the backup this restore references.
	// If this field is not supplied then it will be selected using the following logic:
	// 1. If the backup that we are performing the restore on points to a read-write repository in the current
	// cluster, this repository is selected.
	// 2. If the backup that we are performing a restore on points to a read-only repository, any
	// read-write repository from the cluster is selected and used.
	// +optional
	BackupRepository string `json:"backupRepository,omitempty" reflect:"unexport"`
	// An encryption key. This field is immutable.
	// +optional
	EncryptionKey *EncryptionKey `json:"encryptionKey,omitempty" reflect:"unexport"`
}

// Defines the fine-grained restore Filter
type Filter struct {
	//Selects resources for restoration.
	//If specified, only resources which match inclusion_filters will be selected for restoration.
	//A resource will be selected if it matches any ResourceSelector of the inclusion_filters.
	// +optional
	InclusionFilters []ResourceSelector `json:"inclusionFilters,omitempty" reflect:"unexport"`
	//Excludes resources from restoration.
	//If specified, a resource will not be restored if it matches any ResourceSelector of the exclusion_filters.
	// +optional
	ExclusionFilters []ResourceSelector `json:"exclusionFilters,omitempty" reflect:"unexport"`
}

// Defines ResourceSelector for fine-grained restore Filter
type ResourceSelector struct {
	//Selects resources using their Kubernetes GroupKinds.
	//If specified, only resources of provided GroupKind will be selected.
	// +optional
	GroupKind *metav1.GroupKind `json:"groupKind,omitempty"`
	//Selects resources using their resource names.
	//If specified, only resources with the provided name will be selected.
	// +optional
	Name string `json:"name,omitempty" reflect:"unexport"`
	//Selects resources using their namespaces.
	//This only applies to namespace scoped resources and cannot be used for selecting cluster scoped resources.
	//If specified, only resources in the provided namespace will be selected.
	//If not specified, the filter will apply to both cluster scoped and namespace scoped resources (e.g. name or label).
	//The Namespace resource itself will be restored if and only if any resources within the namespace are restored.
	// +optional
	Namespace string `json:"namespace,omitempty" reflect:"unexport"`
	//Selects resources using Kubernetes labels.
	//If specified, a resource will be selected if and only if the resource has all of the provided labels and all the label values match.
	// +optional
	Labels map[string]string `json:"labels,omitempty" reflect:"unexport"`
}

// Defines the rules of target candidates for substitution.
type SubstitutionTarget struct {
	// A list of target namespaces a
	// substitution rule is applied to. A resource's original namespace must
	// match one of the names in the list to be considered as a substitution
	// candidate. For any non-namespaced
	// resource to be considered, an empty string must be provided.
	// If empty, all namespaces are considered.
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`

	// A list of target `GroupKind` resources a substitution
	// rule applies to. A resource's original `GroupKind` resource must match one of the
	// `GroupKind` resources specified in the list to be considered as a substitution candidate.
	// If empty, all `GroupKind` resources are considered.
	// Note, a resource needs to satisfy both namespace requirements and the `GroupKind`
	// requirements to be considered as a substitution candidate.
	// +optional
	GroupKinds []metav1.GroupKind `json:"groupKinds,omitempty"`

	// The string representation of the JSON Path which leads to the
	// fields in the target resource, which are in the JSON format for substitution.
	// +kubebuilder:validation:Required
	JSONPath string `json:"jsonPath"`
}

// Defines the rules for substitution.
type SubstitutionRule struct {
	// The matching criteria that checks whether a resource is a substitution
	// candidate and has fields that match the supplied `JSONPath` value.
	// +kubebuilder:validation:Required
	Target SubstitutionTarget `json:"target"`

	// A regular expression pattern string which is applied to
	// the JSON value matched from the supplied `Target` value. If not specified, the original
	// value is always substituted using the specified `NewValue` value.
	// Note, an empty string is a legitimate regex pattern, and is treated like
	// any other string.
	// +optional
	// +kubebuilder:default:=.*
	OriginValuePattern *string `json:"originValuePattern,omitempty"`

	// The desired value in string format to substitute to.
	// +kubebuilder:validation:Required
	NewValue string `json:"newValue"`
}

// Defines the rules to filter down the list of candidate resources
// to provide transformation rules against.  All fields with `ResourceFilter` are optional,
// so not specifying any of the `ResourceFilter` fields results in matching on all resources.
type ResourceFilter struct {
	// A list of target namespaces a
	// transformation rule will apply to. A resource's original namespace must
	// match one of the names in the list to be considered as a transformation
	// candidate. An empty string is expected in the list for any non-namespaced
	// resource to be considered.
	// If empty, all namespaces are considered.
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`

	// A list of target `GroupKind` resources a transformation
	// rule will apply to. A resource's original `GroupKind` resource must match one of the
	// `GroupKind` resources specified in the list to be considered as a transformation candidate.
	// If empty, all `GroupKind` resources are considered.
	// Note, a resource needs to satisfy both the `Namespaces` requirement and `GroupKinds`
	// requirement to be considered as a transformation candidate.
	// +optional
	GroupKinds []metav1.GroupKind `json:"groupKinds,omitempty"`

	// The string representation of the path in the JSON which leads to the
	// fields in the target resource for filtering.
	// `JSONPath` documentation can be viewed here: https://github.com/json-path/JsonPath/blob/master/README.md.
	// `JSONPath` formats can be tested online (e.g. https://jsonpath.com/).
	// +optional
	JSONPath *string `json:"jsonPath,omitempty"`
}

// The intended operation for a transformation rule.
// +kubebuilder:validation:Enum=replace;add;remove;copy;move;test
type TransformationRuleOperation string

// Supported selection types.
const (
	Replace TransformationRuleOperation = "replace"
	Add     TransformationRuleOperation = "add"
	Remove  TransformationRuleOperation = "remove"
	Copy    TransformationRuleOperation = "copy"
	Move    TransformationRuleOperation = "move"
	Test    TransformationRuleOperation = "test"
)

// Represents the action the user wants to perform as part of transforming a resource.
type TransformationRuleAction struct {
	// The type of action the user wants to perform as part of transforming resources
	// upon a restore operation.
	// The list of valid operations can be viewed at https://www.rfc-editor.org/rfc/rfc6902#section-4.4.
	// +kubebuilder:validation:Required
	Op TransformationRuleOperation `json:"op"`

	// The JSON patch representation of a path within the JSON which leads to the
	// fields in the target resource for move and copy operations.
	// +optional
	From *string `json:"from,omitempty"`

	// The JSON patch representation of a path within the JSON which leads to the
	// fields in the target resource for some operations.
	// +kubebuilder:validation:Required
	Path string `json:"path"`

	// The desired value in string format to use for transformation.
	// +optional
	Value *string `json:"value,omitempty"`
}

// Defines rules of transformation utilizing `jsonPatch`.
type TransformationRule struct {
	// A list of operations to take against
	// candidate resources based on the JSON Patch RFC.
	// +kubebuilder:validation:Required
	FieldActions []TransformationRuleAction `json:"fieldActions"`

	// The matching criteria to check that determines whether a resource is a transformation
	// candidate.
	// Note, an empty or blank `resourceFilter` value results in a match on all resources.
	// +kubebuilder:validation:Required
	ResourceFilter ResourceFilter `json:"resourceFilter"`
}

// Defines the selection of namespaced resources for restoration.
// One of the following values must be specified along with a matching `Type`:
//
//	`AllNamespaces`, `SelectedNamespaces`, or `SelectedApplications`.
type NamespacedResourceSelection struct {
	// The type of selection to use.
	// +kubebuilder:validation:Required
	Type NamespacedResourceSelectionType `json:"type"`

	// Specifies whether or not to restore all namespaced
	// resources in the backup.
	// Leaving this field unspecified, or setting it to `False` with the `Type`
	// set to `NamespacedResourceSelectionType.AllNamespaces`
	// is invalid.
	// +optional
	AllNamespaces *bool `json:"allNamespaces,omitempty"`

	// A list of the selected namespaces to restore. If a value is provided along with
	// the `Type` set to `NamespacedResourceSelectionType.SelectedNamespaces`,
	// resources with an original namespace that feature in the specified
	// list are restored.
	// Note, specifying an empty string in this list does
	// not restore non-namespaced cluster resources. To restore
	// cluster resources, see `ClusterResourceSelection`.
	// +optional
	SelectedNamespaces []string `json:"selectedNamespaces,omitempty"`

	// A list of selected `ProtectedApplication` resources to restore. If a value is provided,
	// along with the `Type` set to `NamespacedResourceSelectionType.SelectedApplications`,
	// the resources belonging to one of the listed applications are restored.
	// +optional
	SelectedApplications []NamespacedName `json:"selectedApplications,omitempty"`
}

// Represents the selection of non-namespaced resources for restoration.
type ClusterResourceSelection struct {
	// A list of `GroupKind` resources. A non-namespaced resource must be included as one of the
	// `GroupKind` resources specified in the list.
	// Note, even though `PersistentVolume` resources are non-namespaced, these are
	// handled separately. See `VolumeDataRestorePolicy` for details. Specifying
	// a `PersistentVolume` `GroupKind` in this list does not affect whether or
	// not a `PersistentVolume` is restored.
	// +optional
	SelectedGroupKinds []metav1.GroupKind `json:"selectedGroupKinds,omitempty"`
}

// Defines the selected resources to restore. It consists of
// `NamespacedResources` and `ClusterResources` resources. The `NamespacedResources`
// field specifies the namespaced Kubernetes resources selected for restoration,
// and the `ClusterResources` field specifies selected non-namespaced resources.
type RestoreResourceSelection struct {
	// The selected namespaces Kubernetes resources in the backup to restore.
	// +optional
	NamespacedResources *NamespacedResourceSelection `json:"namespacedResources,omitempty"`

	// The selected non-namespaced Kubernetes resources in the backup to restore.
	// +optional
	ClusterResources *ClusterResourceSelection `json:"clusterResources,omitempty"`
}

// Defines the policies of volumes and cluster resources.
type RestorePolicies struct {
	// The policy to use for volume data restoration.
	// Default to `NoVolumeDataRestoration` if not specified.
	// +optional
	VolumePolicy *VolumeDataRestorePolicy `json:"volumePolicy,omitempty"`

	// The policy used to resolve conflicts when restoring cluster-scoped resources.
	// Default to `UseExistingVersion` if not specified.
	// +optional
	ClusterResourcePolicy *ClusterResourceConflictPolicy `json:"clusterResourcePolicy,omitempty"`

	// The restoration mode for namespaced resources.
	// Default to `FailOnConflict` if not specified.
	// +optional
	NamespacedResourceRestoreMode *NamespacedResourceRestoreMode `json:"namespacedResourceRestoreMode,omitempty"`
}

// The type of namespaced resources selection.
// +kubebuilder:validation:Enum=AllNamespaces;SelectedNamespaces;SelectedApplications
type NamespacedResourceSelectionType string

// The supported selection types.
const (
	AllNamespaces        NamespacedResourceSelectionType = "AllNamespaces"
	SelectedNamespaces   NamespacedResourceSelectionType = "SelectedNamespaces"
	SelectedApplications NamespacedResourceSelectionType = "SelectedApplications"
)

// The restoration behavior
// for non-namespaced cluster resources to follow when a conflict occurs.
// +kubebuilder:validation:Enum=UseExistingVersion;UseBackupVersion
type ClusterResourceConflictPolicy string

// The supported conflict policies for cluster resources.
const (
	// The policy that decides whether to use the existing version of the cluster resource.
	// If used, Nno conflict is reported and the final state of the
	// `RestoreJob` resource is not affected.
	UseExistingVersion ClusterResourceConflictPolicy = "UseExistingVersion"

	// The policy that drives the controller to delete
	// the existing version and replace it with the one in the backup.
	// This is an extremely dangerous option which might cause
	// data loss. For example, deletion of a Custom Resource Definition causes Kubernetes to delete all
	// Custom Resources of that type.
	UseBackupVersion ClusterResourceConflictPolicy = "UseBackupVersion"
)

// The restoration behavior
// for namespaced resources.
// +kubebuilder:validation:Enum=DeleteAndRestore;FailOnConflict;MergeSkipOnConflict;MergeReplaceOnConflict
type NamespacedResourceRestoreMode string

const (
	// The mode where the existing top level resources
	// being restored and the resources underneath them are
	// deleted before restoration. Examples of top level resources are
	// `Namespace` or `ProtectedApplication` resources.
	// Note that this mode could cause data loss as it deletes existing
	// resources from the target cluster. A typical scenario to use
	// this mode in the `RestoreJob` would be rollback from a problematic upgrade
	// to the previously saved state.
	DeleteAndRestore NamespacedResourceRestoreMode = "DeleteAndRestore"

	// The mode where the restoration fails when it
	// encounters a conflict. Existence checks at `Namespace`
	// or `ProtectedApplication` level will happen before restoration.
	// If the application being restored or the namespace exists in the
	// target cluster, the `RestoreJob` fails immediately.
	// If a conflict happened after existence checks, and the
	// workload resource being restored or namespace gets
	// created by another process or manually, a conflict is
	// be reported.
	FailOnConflict NamespacedResourceRestoreMode = "FailOnConflict"

	// The mode where restoration skips any resource collisons. Resources
	// that are not already on the cluster will be applied.
	MergeSkipOnConflict NamespacedResourceRestoreMode = "MergeSkipOnConflict"

	// The mode where restoration replaces the existing version with the backup
	// version in the case they differ.
	// Note that this mode could cause data loss as it replaces the existing
	// resources in the target cluster, and the original PV can be retained or
	// deleted depending on its reclaim policy.
	MergeReplaceOnConflict NamespacedResourceRestoreMode = "MergeReplaceOnConflict"
)

// The different volume data restoration policies.
// +kubebuilder:validation:Enum=NoVolumeDataRestoration;ReuseVolumeHandleFromBackup;RestoreVolumeDataFromBackup
type VolumeDataRestorePolicy string

const (
	// A policy where a new `PersistentVolume` resource is
	// restored using the volume backup data in the backup.
	RestoreVolumeDataFromBackup VolumeDataRestorePolicy = "RestoreVolumeDataFromBackup"

	// A policy where a new `PersistentVolume` resource is
	// pre-provisioned using the volume handle of the original
	// `PersistentVolume` resource in the backup.
	ReuseVolumeHandleFromBackup VolumeDataRestorePolicy = "ReuseVolumeHandleFromBackup"

	// A policy where the `PersistentVolume` resource is not restored.
	// The restoration only restores selected `PersistentVolumeClaim` resources
	// and expects corresponding controllers to either dynamically provision
	// blank `PersistentVolume` resources or bind them to pre-provisioned
	// `PersistentVolume` resources created out-of-bounds.
	NoVolumeDataRestoration VolumeDataRestorePolicy = "NoVolumeDataRestoration"
)

// GetObjectCopy returns a deepcopy of the given controlplane CR.
func GetObjectCopy(object client.Object) (client.Object, error) {
	switch o := object.(type) {
	case *Backup:
		return o.DeepCopy(), nil
	case *BackupPlan:
		return o.DeepCopy(), nil
	case *Restore:
		return o.DeepCopy(), nil
	case *RestorePlan:
		return o.DeepCopy(), nil
	case *VolumeRestore:
		return o.DeepCopy(), nil
	case *VolumeBackup:
		return o.DeepCopy(), nil
	case *VMBackupJob:
		return o.DeepCopy(), nil
	default:
		return nil, fmt.Errorf("unknown type for imported object: %T", o)
	}
}

// GetUIDAnnotation gets the UID annotation for an object.
func GetUIDAnnotation(object client.Object) (string, bool) {
	v, ok := object.GetAnnotations()[OriginClusterUIDAnnotation]
	return v, ok
}

// WorkloadClusterName gets the name of the Cluster containing the Workloads to be backed up.
//
// If the empty string is returned, then there is no external cluster for the job.
func WorkloadClusterName(object client.Object) string {
	// Will return empty string if key is missing.
	return object.GetLabels()[WorkloadClusterNameLabel]
}

func SetWorkloadClusterName(object client.Object, workloadClusterName string) error {
	m := object.GetLabels()

	// Create the map and set it on the object if it is nil.
	if m == nil {
		m = map[string]string{}
		object.SetLabels(m)
	} else if v, found := m[WorkloadClusterNameLabel]; found {
		return fmt.Errorf("workload cluster name is already set to: %v", v)
	}

	m[WorkloadClusterNameLabel] = workloadClusterName
	return nil
}

// HasIdenticalAnnotation return whether two resources contain the same value for the provided annotation
func HasIdenticalAnnotation(a client.Object, b client.Object, annotation string) bool {
	aAnnotation, aFound := a.GetAnnotations()[annotation]
	bAnnotation, bFound := b.GetAnnotations()[annotation]
	if aFound != bFound {
		return false
	}
	return aAnnotation == bAnnotation
}
