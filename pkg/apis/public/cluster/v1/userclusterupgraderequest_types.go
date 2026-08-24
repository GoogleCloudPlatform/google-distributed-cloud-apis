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

	upgradev1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/upgrade/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Target Version",type="string",JSONPath=".spec.targetVersion"
// +kubebuilder:printcolumn:name="Succeeded",type="string",JSONPath=".status.conditions[?(@.type==\"Succeeded\")].status"

// Represents the configuration of a user cluster upgrade request, such as the
// cluster reference and the target Kubernetes version.
type UserClusterUpgradeRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserClusterUpgradeRequestSpec   `json:"spec,omitempty"`
	Status UserClusterUpgradeRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of `UserClusterUpgradeRequest` resources.
type UserClusterUpgradeRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UserClusterUpgradeRequest `json:"items"`
}

// Provides the specification, such as the desired state, of a
// `UserClusterUpgradeRequest` resource.
type UserClusterUpgradeRequestSpec struct {
	// A reference to the GDC user cluster object to which the upgrade applies.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule=(self == oldSelf),message="Value is immutable"
	ClusterRef corev1.LocalObjectReference `json:"clusterRef"`

	// The target Kubernetes user cluster version.
	// +kubebuilder:validation:Required
	TargetVersion string `json:"targetVersion"`

	// Schema for parallelization of node upgrades
	NodePoolUpgradePhases []UpgradePhaseConfig `json:"nodePoolUpgradePhases,omitempty"`

	// Toggle for concurrent worker node pool ABM upgrades
	// True upgrades all worker node pools together, False upgrades one at a time.
	ConcurrentUpgrades bool `json:"concurrentUpgrades,omitempty"`

	// Disruption defines the disruption level for node upgrades.
	// If this field is not set, patch upgrades default to executing without draining workloads (NoDisruption).
	// Setting this field to "RequiresWorkloadDrain" overrides that default behavior, forcing the system
	// to perform a standard drain process on the nodes.
	// Note: This field is ignored for minor version upgrades, which always require workload drain.
	// +optional
	// +kubebuilder:validation:Enum=RequiresWorkloadDrain;NoDisruption;RequiresNodeReboot
	Disruption *upgradev1.NodeUpgradeDisruption `json:"disruption,omitempty"`
}

type UpgradePhaseConfig struct {
	// Maximum number of elements to upgrade concurrently. -1 to specify all remaining.
	MaxConcurrency int `json:"maxConcurrency"`

	// Percent success for considering this group of node pools as completed.
	CompletionThreshold int `json:"completionThreshold,omitempty"`
}

type UpgradeStepStatus struct {
	// Conditions represents the observed status of the upgrade step.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Phase tracks the index of the current phase.
	Phase *int `json:"phase,omitempty"`

	// PhaseBatchNames lists NodeUpgrade CR names being processed or completed.
	PhaseBatchNames []string `json:"phaseBatchNames,omitempty"`

	// StartTime represents the start time of the upgrade step.
	StartTime *metav1.Time `json:"startTime,omitempty"`
}

type UserClusterUpgradeRequestStatus struct {
	// The current upgrade state.
	// Known condition types: Succeeded.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// The observed start time for the current upgrade.
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// Cluster represents the observed cluster upgrade status.
	Cluster UpgradeStepStatus `json:"cluster,omitempty"`

	// Node represents the observed Node upgrade status
	// for cluster nodes.
	Node UpgradeStepStatus `json:"node,omitempty"`
}

func init() {
	SchemeBuilder.Register(&UserClusterUpgradeRequest{}, &UserClusterUpgradeRequestList{})
}
