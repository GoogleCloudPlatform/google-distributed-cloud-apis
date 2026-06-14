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

package v1alpha1

const (
	// ReadyStatusType is the condition type that indicates the overall readiness of
	// a resource.
	ReadyStatusType = "Ready"

	// PropagatedStatusType is the condition type that indicates the resource is fully propagated
	// to all user clusters.
	PropagatedStatusType = "Propagated"

	// HRASecretReadyStatusType is the condition type that indicates the readiness of HRA secret.
	// When HRASecretReadyStatusType status is true, it indicates one of the following scenarios:
	// 1) HRA has no permissions.
	// 2) HRA has permissions, HRA secret is created and propagated to all user clusters.
	HRASecretReadyStatusType = "HarborSecretReady"

	// S3SecretReadyStatusType is the condition type that indicates the readiness of s3 secret.
	// When HRASecretReadyStatusType status is true, it indicates one of the following scenarios:
	// 1) There is no associated s3 secrets.
	// 2) All associated s3 secret are propagated to all user clusters.
	S3SecretReadyStatusType = "S3SecretReady"

	// EgressNATType is the condition type that indicates the overall readiness of
	// EgressNAT resources in the cluster.
	EgressNATType = "EgressNATReady"
)

const (
	// ReadyCondition is the condition type that indicates the overall readiness of
	// a resource.
	ReadyCondition = "Ready"

	// InvalidReason indicates the resource is invalid.
	InvalidReason = "Invalid"

	// ReconciledReason indicates the resource is reconciled successfully.
	ReconciledReason = "Reconciled"

	// ReconcilingReason indicates that a reconciliation is in progress.
	ReconcilingReason = "Reconciling"
)

const (
	// PropagatedCondition is the condition type that indicates the resource is fully propagated
	// to all user clusters.
	PropagatedCondition = "Propagated"

	// PropagatedReason indicates that the resource was successfully propagated.
	PropagatedReason = "Propagated"

	// PropagatingReason indicates that the resource was not successfully
	// propagated to all clusters.
	PropagatingReason = "Propagating"

	// PropagationFailedReason indicates that the resource was not successfully
	// propagated.
	PropagationFailedReason = "PropagationFailed"
)

const (
	// EgressNATReadyCondition is the condition type that indicates the overall readiness of
	// EgressNAT resources in the user cluster.
	EgressNATReadyCondition = "EgressNATReady"

	// EgressNATConfiguredReason indicates that the EgressNAT resources are
	// configured on the user clusters.
	EgressNATConfiguredReason = "EgressNATConfigured"

	// EgressNATConfigurationFailedReason indicates that the EgressNAT resources configuration failed
	// on the user clusters.
	EgressNATConfigurationFailedReason = "EgressNATConfigurationFailed"

	// EgressNATDisabledReason indicates that the EgressNAT configuration is disabled on the project.
	EgressNATDisabledReason = "EgressNATDisabled"
)

const (
	// DetachingCondition is the condition type that indicates that the Project is detaching from
	// certain user clusters.
	DetachingCondition = "Detaching"

	// DetachingFailedReason indicates the Project detaching process failed in
	// certain user clusters.
	DetachingFailedReason = "DetachingFailed"

	// DetachingReason indicates the Project is detaching from
	// certain user clusters,
	// the condition message explains what blocks resource detaching.
	DetachingReason = "Detaching"
)

// Following group contains terminating condition types and reasons.
const (
	// TerminatingCondition is the condition type that indicates that the resource is being
	// deleted.
	TerminatingCondition = "Terminating"

	// TerminatingFailedReason indicates the resource terminating process failed.
	TerminatingFailedReason = "TerminatingFailed"

	// TerminatingReason indicates the resource is being deleted,
	// the condition message explains what blocks resource deletion.
	TerminatingReason = "Terminating"

	// RemovingOCResourceReason indicates the current ongoing deletion phase is removing OC Resources.
	RemovingOCResourceReason = "RemovingOCResource"

	// RemovingNamespacesReason indicates the current ongoing deletion phase is removing Project Namespaces.
	RemovingNamespacesReason = "RemovingProjectNamespaces"

	// OCResourcesRemainingCondition is the condition type that indicates if any OC resource are remaining.
	OCResourcesRemainingCondition = "OCResourcesRemaining"

	// NamespacesRemainingCondition is the condition type that indicates if any Project Namespaces are remaining.
	NamespacesRemainingCondition = "ProjectNamespacesRemaining"

	// SomeRemainReason is the condition reason that indicates some resources are remaining.
	SomeRemainReason = "SomeRemain"

	// NothingRemainReason is the condition reason that indicates nothing is remaining.
	NothingRemainReason = "NothingRemains"
)

const (
	// AttachedCondition is the condition type that indicates that whether the Project
	// is attached to user clusters.
	// AttachedCondition is an aggregated condition of PropagatedCondition and EgressNATReadyCondition.
	AttachedCondition = "Attached"
)

const (
	// TagMgmtNamespaceCreatedCondition is set to true when the namespace
	// of a tag key is created.
	TagMgmtNamespaceCreatedCondition = "TagMgmtNamespaceCreated"

	// TagPolicyNamespaceCreatedCondition is set to true when the namespace
	// of a tag value is created.
	TagPolicyNamespaceCreatedCondition = "TagPolicyNamespaceCreated"

	// FailedToCreateTagMgmtNamespaceReason indicates failed to create tagKey's
	// managementNamespace.
	FailedToCreateTagMgmtNamespaceReason = "FailedToCreateTagMgmtNamespace"

	// FailedToCreateTagPolicyNamespaceReason indicates failed to create tagvalue's
	// policyNamespace.
	FailedToCreateTagPolicyNamespaceReason = "FailedToCreateTagPolicyNamespace"

	// InvalidTagValueNamespaceReason indicates an error when tagValue created
	// in a invalid namespace.
	InvalidTagValueNamespaceReason = "InvalidTagValueNamespace"

	// TagkeyRolesCreatedCondition is set to true when the roles for
	// a tagkey are created.
	TagkeyRolesCreatedCondition = "TagkeyRolesCreated"

	// TagvalueRolesCreatedCondition is set to true when the roles for
	// a tagvalue are created.
	TagvalueRolesCreatedCondition = "TagvalueRolesCreated"

	// TagRoleCreationInProgressReason indicates tag role creation in progress.
	TagRoleCreationInProgressReason = "TagRoleCreationInProgress"

	// FailedToCreateTagRoleTemplateReason indicates failed to create tag role template.
	FailedToCreateTagRoleTemplateReason = "FailedToCreateTagRoleTemplate"

	// TagKeyDeletingCondition is the deleting condition for a tagkey.
	TagKeyDeletingCondition = "TagKeyDeleting"

	// TagValueExistsReason indicates there is still tagvalue exists.
	TagValueExistsReason = "TagValueExists"

	// RoleTemplateFailedToDeleteReason indicates the roleTemplate failed to delete.
	RoleTemplateFailedToDeleteReason = "RoleTemplateFailedToDelete"

	// MgmtNamespaceFailedToDeleteReason indicates the mgmtNamespace failed to delete.
	MgmtNamespaceFailedToDeleteReason = "MgmtNamespaceFailedToDelete"

	// TagValueDeletingCondition is the deleting condition for a tagvalue.
	TagValueDeletingCondition = "TagValueDeleting"

	// PolicyNamespaceFailedToDeleteReason indicates the policyNamespace failed to delete.
	PolicyNamespaceFailedToDeleteReason = "PolicyNamespaceFailedToDelete"

	// TagKeyInitPermissionCondition is set to true when the initial permission granted for a tagkey.
	TagKeyInitPermissionCondition = "TagKeyInitPermission"

	// TagRoleBindingCreationInProgressReason indicates tag roleBinding creation in progress.
	TagRoleBindingCreationInProgressReason = "TagRoleBindingCreationInProgress"

	// TagValueInitPermissionCondition is set to true when the initial permission granted for a tagvalue.
	TagValueInitPermissionCondition = "TagValueInitPermission"
)

const (
	// KubernetesClientErrorReason indicates an error is returned when the k8s client
	// performs read or write action.
	KubernetesClientErrorReason = "KubernetesClientError"

	// ClusterNotReachableReason indicates the cluster cannot be reached on the network.
	ClusterNotReachableReason = "ClusterNotReachable"

	// ResourceConfiguredReason indicates that the egress NAT resource is
	// configured on the user cluster.
	ResourceConfiguredReason = "Configured"

	// ResourceNotConfiguredReason indicates that the egress NAT resource is
	// not configured on the user cluster.
	ResourceNotConfiguredReason = "NotConfigured"
)

// following constants are all condition reasons for HarborSecretReady condition type.
const (
	// HRACreationFailReason indicates that the HRA creation fails.
	HRACreationFailReason = "HarborRobotCreationFail"
	// HRACreatingReason indicates that the HRA is being created.
	HRACreatingReason = "CreatingHarborRobot"
	// HRASecretCreationFailReason indicates that the HRA secret creation fails.
	HRASecretCreationFailReason = "HarborRobotSecretCreationFail"
	// HRASecretNoPermissionReason indicates that the HRA has no associated RBAC permissions.
	HRASecretNoPermissionReason = "RobotHasNoPermissions"
	// HRASecretPropagatedReason indicates that the HRA secret is propagated to all user clusters.
	HRASecretPropagatedReason = "RobotSecretPropagated"
)

// following constants are all condition reasons for S3SecretReadyStatusType condition type
const (
	// NoS3SecretReason indicates that there is no associated s3 secret bound to
	// Project Service Account.
	NoS3SecretReason = "NoS3Secrets"

	// S3SecretPropagationFailReason indicates that an error occurs during s3 propagation
	S3SecretPropagationFailReason = "S3SecretsPropagationFail"
)

// ProjectWipeoutConfigIndexKey is the index key build on ProjectWipeoutConfig.Spec.ProjectName.
const ProjectWipeoutConfigIndexKey = "spec.projectName"
