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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="URL",type="string",JSONPath=".status.url"
// +kubebuilder:printcolumn:name="Version",type="string",JSONPath=".status.version"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +gdcloud:manifest:relevant=false,oc=haas
// HarborInstance represents an instance of a Harbor container registry.
type HarborInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of the HarborInstance.
	Spec HarborInstanceSpec `json:"spec,omitempty"`
	// Most recently observed status of the HarborInstance.
	Status HarborInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HarborInstanceList represents a collection of Harbor container registry instances.
type HarborInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HarborInstance `json:"items"`
}

// HarborInstanceSpec represents the specification or the desired state of a HarborInstance.
type HarborInstanceSpec struct {
}

// HarborInstanceStatus represents the current status of a HarborInstance.
type HarborInstanceStatus struct {
	// Current HarborInstance state.
	//
	// - Ready: Indicates that the HarborInstance is ready.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Version of the Harbor instance.
	Version string `json:"version,omitempty"`

	// URL of Harbor instance's web UI.
	URL string `json:"url,omitempty"`
}

// HarborInstance Conditions
const (
	// HarborInstanceReady is the status condition that indicates the Harbor container registry
	// represented by the HarborInstance object is ready.
	HarborInstanceReady = "Ready"
	// HarborInstanceHelmReady is the status condition that indicates the helm deployment
	// represented by the HarborInstance object is ready.
	HarborInstanceHelmReady = "HelmReady"
)

// HarborInstance Condition Reason
const (
	// ReasonReconciling is a condition reason for condition type Ready, indicating that a reconciliation is in progress.
	ReasonReconciling = "Reconciling"
	// ReasonReconciled is a condition reason for condition type Ready, indicating that a reconciliation is completed successfully.
	ReasonReconciled = "Reconciled"
	// ReasonReconcileBackoff is a condition reason for condition type Ready, indicating that a temporary error has occurred
	// during a reconciliation, and another reconciliation has been scheduled.
	ReasonReconcileBackoff = "ReconcileBackoff"
	// ReasonShadowProjectNotReady is a condition reason for condition type Ready, indicating that shadow project is not ready.
	ReasonShadowProjectNotReady = "ShadowProjectNotReady"
)

// HelmReconciliation Condition Reason
const (
	// ReasonHelmReady is a condition reason for condition type HelmReady True,
	// indicating that a helm deployment is in ready state.
	ReasonHelmReady = "Ready"
	// ReasonHelmFatal is a condition reason for condition type HelmReady False,
	// indicating that a helm deployment is in fatal state from auto recovery and requires manual fix.
	ReasonHelmFatal = "Fatal"
	// ReasonHelmProvisionFailed is a condition reason for condition type HelmReady False,
	// indicating that a helm installation has failed and a retry is in progress.
	ReasonHelmProvisionFailed = "ProvisionFailed"
	// ReasonHelmUninstallFailed is a condition reason for condition type HelmReady False,
	// indicating that a helm uninstall has failed and a retry is in progress.
	ReasonHelmUninstallFailed = "UninstallFailed"
	// ReasonHelmUninstalled is a condition reason for condition type HelmReady False,
	// indicating that a helm deployment has been uninstalled.
	ReasonHelmUninstalled = "Uninstalled"
	// ReasonHelmRollBackFailed is a condition reason for condition type HelmReady False,
	// indicating that a helm rollback has failed and a retry is in progress.
	// Some rollback can happen and finish during atomic helm operation without setting to this state.
	ReasonHelmRollBackFailed = "RollBackFailed"
	// ReasonHelmRolledBack is a condition reason for condition type HelmReady False,
	// indicating that a helm explicit rollback by reconciler has succeed and a
	// reconcile is in progress.
	ReasonHelmRolledBack = "RolledBack"
	// ReasonHelmUpgradeFailed is a condition reason for condition type HelmReady False,
	// indicating that a helm upgrade has failed and a retry is in progress.
	ReasonHelmUpgradeFailed = "UpgradeFailed"
)

func init() {
	SchemeBuilder.Register(
		&HarborInstance{},
		&HarborInstanceList{},
	)
}
