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
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true

// LocalRollout represents the intent to rollout a new release of a Module to a group or all Module
// Instances in a location.
//
// The location is an abstraction. In the most common case it maps to a single kubernetes cluster.
// However it can have a wider scope in the context of a multi-cluster cloud.
//
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=lr
// +kubebuilder:printcolumn:JSONPath=`.spec.paused`,name="Paused",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="InProgress")].status`,name="In Progress",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.conditions[?(@.type=="Finished")].status`,name="Finished",type="string"
type LocalRollout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocalRolloutSpec   `json:"spec,omitempty"`
	Status LocalRolloutStatus `json:"status,omitempty"`
}

type LocalRolloutSpec struct {
	// The SaasType of the current rollout.
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	SaasType string `json:"saasType"`

	// Whether the current LocalRollout is paused.
	//
	// Paused LocalRollouts will not cause ModuleInstances to be upgraded.
	Paused bool `json:"paused"`

	// The Module identifier of the current LocalRollout.
	//
	// The module identifier is unique within the scope of a single SaasType. This along with the
	// SaasType is used to select the ModuleInstance(s) related to the current LocalRollout.
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	Module string `json:"module"`

	// The identifier of the Release version which is being rolled out.
	//
	// The release identifier is unique within the scope of a single SaasType.
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	Release string `json:"release"`

	// The list of Release versions the Release version can upgrade from.
	UpgradableFrom []string `json:"upgradableFrom"`

	// Additional resources required for the ModuleInstance to be updated.
	//
	// This field is used by SaasProducers to additional data to Saas operator.
	DeploymentResources []DeploymentResourceReference `json:"deploymentResources,omitempty"`
}

type DeploymentResourceReference struct {
	ApiGroup  string `json:"apiGroup"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name"`
}

const (
	LocalRolloutConditionInProgress = "InProgress"
	LocalRolloutConditionFinished   = "Finished"

	LocalRolloutConditionPauseAcknowledged = "PauseAcknowledged"
)

type LocalRolloutStatus struct {
	// Conditions contain conditions for LocalRollout.
	// LocalRollout controller sets following conditions:
	// - InProgress: The controller still waits for ModuleInstances to be upgraded.
	// - Finished:   The controller found no more ModuleInstances to be upgraded.
	//
	// +listType=map
	// +listMapKey=type
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,7,rep,name=conditions"`

	// ErrorStatus implements error status convention as defined in go/gdc-krm-error-signaling.
	// +optional
	ErrorStatus ErrorStatus `json:"errorStatus,omitempty"`
}

// +kubebuilder:object:root=true

type LocalRolloutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []LocalRollout `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LocalRollout{}, &LocalRolloutList{})
}

func (lr *LocalRollout) IsUpgradableFrom(release string) bool {
	for _, r := range lr.Spec.UpgradableFrom {
		if r == release {
			return true
		}
	}
	return false
}

func (lr *LocalRollout) SetInProgress(reason string, message string) {
	lr.SetStatusCondition(LocalRolloutConditionInProgress, metav1.ConditionTrue, reason, message)
}

func (lr *LocalRollout) IsInProgress() bool {
	return lr.IsStatusConditionTrue(LocalRolloutConditionInProgress)
}

func (lr *LocalRollout) SetFinished(reason string, message string) {
	lr.SetStatusCondition(LocalRolloutConditionFinished, metav1.ConditionTrue, reason, message)
}

func (lr *LocalRollout) IsFinished() bool {
	return lr.IsStatusConditionTrue(LocalRolloutConditionFinished)
}

func (lr *LocalRollout) SetStatusCondition(conditionType string, status metav1.ConditionStatus, reason string, message string) {
	meta.SetStatusCondition(&lr.Status.Conditions, metav1.Condition{
		Type:               conditionType,
		Status:             status,
		Reason:             reason,
		Message:            message,
		ObservedGeneration: lr.Generation,
	})
}

func (lr *LocalRollout) IsStatusConditionTrue(conditionType string) bool {
	return lr.IsStatusConditionCurrentAndEqual(conditionType, metav1.ConditionTrue)
}

func (lr *LocalRollout) IsStatusConditionFalse(conditionType string) bool {
	return lr.IsStatusConditionCurrentAndEqual(conditionType, metav1.ConditionFalse)
}

func (lr *LocalRollout) IsStatusConditionCurrentAndEqual(conditionType string, status metav1.ConditionStatus) bool {
	condition := meta.FindStatusCondition(lr.Status.Conditions, conditionType)
	return condition != nil && condition.ObservedGeneration == lr.Generation && condition.Status == status
}
