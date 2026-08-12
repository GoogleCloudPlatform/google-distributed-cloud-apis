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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="URL",type="string",JSONPath=".status.url"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".status.version"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +gdcloud:manifest:relevant=true,oc=haas,component=harbor
// +gdcloud:manifest:entities="instances",verbs=create;delete;describe;list
// +gdcloud:manifest:rbac="create,delete,describe,list:harbor-instance-admin"
// +gdcloud:manifest:skipcodegen=true
// Represents an instance of a Harbor container registry.
type HarborInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of the Harbor instance.
	Spec HarborInstanceSpec `json:"spec,omitempty"`
	// The most recently observed status of the Harbor instance.
	Status HarborInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of Harbor container registry instances.
type HarborInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HarborInstance `json:"items"`
}

// Represents the specification or the desired state of a Harbor instance.
type HarborInstanceSpec struct {
	// The reference of which restore resource this instance is restored by.
	// The field is filled by restore reconciliation and immutable once created.
	// +optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="RestoreRef is immutable"
	RestoreRef *HarborInstanceRestoreReference `json:"restoreRef,omitempty"`

	// The reference of which credentials provide administrative access to this instance.
	// +optional
	AdminAccessRef *HarborInstanceAdminAccessReference `json:"adminAccessRef,omitempty"`
}

// Represents the current status of a Harbor instance.
type HarborInstanceStatus struct {
	// The current state of the HarborInstance.
	// A state of 'Ready` indicates that the HarborInstance is ready.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The version of the Harbor instance.
	Version string `json:"version,omitempty"`

	// The URL of Harbor instance's web UI.
	URL string `json:"url,omitempty"`

	// Restore related status will be empty if the instance is not created by restore flow.
	// The timestamp of when this instance is restored.
	// +optional
	RestoreTime *metav1.Time `json:"restoreTime,omitempty"`

	// The reference of which backup resource this instance is restored from.
	// +optional
	RestoreBackupReference *HarborInstanceBackupReference `json:"restoreBackupReference,omitempty"`

	// ErrorStatus holds the most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// TimeSinceLastReady holds the time since the last time HarborInstance was ready.
	TimeSinceLastReady *metav1.Time `json:"timeSinceLastReady,omitempty"`
}

// HarborInstance Conditions
const (
	// The status condition that indicates the Harbor container registry
	// represented by the Harbor instance object is ready.
	HarborInstanceReady = "Ready"
	// The status condition that indicates the helm deployment
	// represented by the HarborInstance object is ready.
	HarborInstanceHelmReady = "HelmReady"
	// The status condition that indicates the admin access
	// represented by the HarborInstance object is ready.
	HarborInstanceAdminAccessReady = "AdminAccessReady"
	// The status condition that indicates the logging target
	// represented by the LoggingTarget object is ready.
	LoggingTargetReady = "LoggingTargetReady"
	// The status condition that indicates the audit logging target
	// represented by the AuditLoggingTarget object is ready.
	AuditLoggingTargetReady = "AuditLoggingTargetReady"
)

// HarborInstance Condition Reason
const (
	// A condition reason for the condition type `Ready`, indicating that a reconciliation is in progress.
	ReasonReconciling = "Reconciling"
	// A condition reason for the condition type `Ready`, indicating that a reconciliation is completed successfully.
	ReasonReconciled = "Reconciled"
	// A condition reason for the condition type `Ready`, indicating that a temporary error has occurred
	// during a reconciliation, and another reconciliation has been scheduled.
	ReasonReconcileBackoff = "ReconcileBackoff"
	// A condition reason for the condition type `Ready`, indicating that shadow project is not ready.
	// Shadow projects are special projects that make it possible to deploy service instances as blackboxes
	// and hide the implementation details from the service consumers.
	ReasonShadowProjectNotReady = "ShadowProjectNotReady"
	// A condition reason for the condition type `Ready` with false condition, indicating that instance has finished restoration mode reconciliation.
	ReasonReconcileRestorationMode = "ReconcileRestorationMode"
	// A condition reason for the condition type `LoggingTargetReady`, indicating that the logging target is not ready.
	ReasonLoggingTargetNotReady = "LoggingTargetNotReady"
	// A condition reason for the condition type `LoggingTargetReady`, indicating that the logging target is ready.
	ReasonLoggingTargetReady = "LoggingTargetReady"
	// A condition reason for the condition type `AuditLoggingTargetReady`, indicating that the audit logging target is not ready.
	ReasonAuditLoggingTargetNotReady = "AuditLoggingTargetNotReady"
	// A condition reason for the condition type `AuditLoggingTargetReady`, indicating that the audit logging target is ready.
	ReasonAuditLoggingTargetReady = "AuditLoggingTargetReady"
)

// HelmReconciliation Condition Reason
const (
	// A condition reason for condition type `HelmReady` True,
	// indicating that a helm deployment is in ready state.
	ReasonHelmReady = "Ready"
	// A condition reason for condition type `HelmReady` False,
	// indicating that a helm deployment is in fatal state from auto recovery and requires manual fix.
	ReasonHelmFatal = "Fatal"
	// RA condition reason for condition type `HelmReady` False,
	// indicating that controller failed to prepare the helm operation with errors in preconditions.
	// Currently this includes failure when getting current helm release revision or existing values from objects in the cluster.
	ReasonHelmPreconditionError = "PreconditionError"
	// A condition reason for condition type `HelmReady` False,
	// indicating that a helm installation has failed and a retry is in progress.
	ReasonHelmProvisionFailed = "ProvisionFailed"
	// A condition reason for condition type `HelmReady` False,
	// indicating that a helm uninstall has failed and a retry is in progress.
	ReasonHelmUninstallFailed = "UninstallFailed"
	// A condition reason for condition type `HelmReady` False,
	// indicating that a helm deployment has been uninstalled.
	ReasonHelmUninstalled = "Uninstalled"
	// A condition reason for condition type `HelmReady` False,
	// indicating that a helm rollback has failed and a retry is in progress.
	// Some rollback can happen and finish during atomic helm operation without setting to this state.
	ReasonHelmRollBackFailed = "RollBackFailed"
	// A condition reason for condition type `HelmReady` False,
	// indicating that a helm explicit rollback by reconciler has succeed and a
	// reconcile is in progress.
	ReasonHelmRolledBack = "RolledBack"
	// A condition reason for condition type `HelmReady` False,
	// indicating that a helm upgrade has failed and a retry is in progress.
	ReasonHelmUpgradeFailed = "UpgradeFailed"
)

// AdminAccess Condition Reason
const (
	// A condition reason for condition type `AdminAccessReady` True,
	// indicating that admin access is in ready state.
	ReasonAdminAccessReady = "Ready"
	// A condition reason for condition type `AdminAccessReady` False,
	// indicating that a temporary error has occurred during a reconciliation, and another reconciliation has been scheduled.
	ReasonAdminAccessReconcileBackoff = "ReconcileBackoff"
	// A condition reason for condition type `AdminAccessReady` False,
	// indicating that admin access has been disabled.
	ReasonAdminAccessDisabled = "Disabled"
)

// HarborInstanceBackupReference represents a Backup Reference to HarborInstanceBackup.
type HarborInstanceBackupReference struct {
	// name is unique within a namespace to reference a HarborInstanceBackup resource.
	// +optional
	Name string `json:"name,omitempty"`
	// namespace defines the space within which the HarborInstanceBackup name must be unique.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// HarborInstanceRestoreReference represents a Restore Reference to HarborInstanceRestore.
type HarborInstanceRestoreReference struct {
	// name is unique within a namespace to reference a HarborInstanceRestore resource.
	// +optional
	Name string `json:"name,omitempty"`
	// namespace defines the space within which the HarborInstanceRestore name must be unique.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// HarborInstanceAdminAccessReference represents an Admin Access Reference to HarborInstance.
type HarborInstanceAdminAccessReference struct {
	// AdminRobotSecretRef is the reference of the secret which contains the credentials for
	// system-level Harbor robot account.
	// The secret must contain the following key-value pair:
	// - `secret`: the secret of the Harbor robot account.
	// +optional
	AdminRobotSecretRef *corev1.LocalObjectReference `json:"adminRobotSecretRef,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&HarborInstance{},
		&HarborInstanceList{},
	)
}
