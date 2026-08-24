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
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true

// ModuleInstance correspond to a single deployment of a Saas service component.
//
// In the most common case a SaasInstance will have a single deployment, hence a single Module
// instance. However in more complex architectures a single SaasInstance can be
// composed of multiple modules and have multiple deployments which will correspond with different
// ModuleInstance(s).
//
// Each ModuleInstance upgrades independently of the others. By using this entity, service producers
// can upgrade some components without touching other components. The component type is determined
// by the Module field. All ModuleInstance(s) within the same SaasInstance are considered part of
// the same service.
//
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Namespaced,shortName=mi
// +kubebuilder:printcolumn:JSONPath=`.spec.release`,name="Target Release",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.release`,name="Current Release",type="string"
// +kubebuilder:printcolumn:JSONPath=`.spec.upgradeScheduledAt`,name="Upgrade Scheduled At",type="string"
// +gdcloud:manifest:relevant=false,oc=ez
type ModuleInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ModuleInstanceSpec   `json:"spec,omitempty"`
	Status ModuleInstanceStatus `json:"status,omitempty"`
}

type ModuleInstanceSpec struct {
	// The resource name of the SaasInstance resource this ModuleInstance belongs to.
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	SaasInstance string `json:"saasInstance"`

	// Type of the SaasInstance this ModuleInstance belongs to.
	//
	// This field is immutable.
	//
	// This field is required, but for now it's marked as optional for compatibility reasons.
	SaasType string `json:"saasType,omitempty"`

	// The idenfifier of the type of Module the current ModuleInstance.
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	Module string `json:"module"`

	// The reference to the ConsumerResource of the parent SaasInstance.
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	ConsumerResource ConsumerResource `json:"consumerResource,omitempty"`

	// The identifier of the Release for the current ModuleInstance.
	//
	// A change in the Release will cause the ModuleInstance deployment to be updated with the
	// specified version. This update will occur at the time indicated by UpgradeScheduledAt.
	//
	// +kubebuilder:validation:Required
	Release string `json:"release,omitempty"`

	// The time at which the ModuleInstance will be updated to the new Release.
	UpgradeScheduledAt metav1.Time `json:"upgradeScheduledAt,omitempty"`

	// Additional resources required for the ModuleInstance to be updated.
	//
	// DeploymentResources is a copy of the data contained in the LocalRollout resource of the
	// current Release.
	DeploymentResources []DeploymentResourceReference `json:"deploymentResources,omitempty"`

	// The identifier of the previous Release of the current ModuleInstance.
	//
	// This field will hold the identifier of the previous *successfully deployed* Release of the
	// ModuleInstance and it is used for debugging a change in the Release value.
	PreviousRelease string `json:"previousRelease,omitempty"`

	// Additional resources in the PreviousRelease.
	//
	// This field will hold the DeploymentResources of the the PreviousRelease.
	PreviousDeploymentResources []DeploymentResourceReference `json:"previousDeploymentResources,omitempty"`

	// Variables used by rollout selectors.
	Variables map[string]ModuleInstanceVariable `json:"variables,omitempty"`
}

const (
	ModuleInstanceConditionProvisioned         = "Provisioned"
	ModuleInstanceConditionNeedReschedule      = "NeedReschedule"
	ModuleInstanceConditionUpgraded            = "Upgraded"
	ModuleInstanceConditionUpgrading           = "Upgrading"
	ModuleInstanceConditionUpgradeAcknowledged = "UpgradeAcknowledged"
)

// ModuleInstanceVariable structure containing variables
//
// It's in form of a structure for future extensibility.
type ModuleInstanceVariable struct {
	Value string `json:"value,omitempty"`
}

type ModuleInstanceStatus struct {
	// Release is the identifier of the current release of the Module Instance when the  upgrade is
	// completed.
	Release string `json:"release,omitempty"`

	// Conditions contain conditions for ModuleInstance.
	// The deployment controller sets one of the conditions:
	// - Provisioned:    After successful initial deployment with "True" status.
	//                   The controller can also set "False" status when deployment failed,
	//                   it is expected that it will continue retrying until successful.
	// - NeedReschedule: When the controller detects upgrade cannot be done right now it sets the
	//                   condition with "True" status.
	// - Upgraded:       The controller sets the condition with "False" status when it is
	//                   unable to upgrade successfully (e.g. after multiple attempts or when rollback failed).
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

type ModuleInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ModuleInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModuleInstance{}, &ModuleInstanceList{})
}

func (mi *ModuleInstance) TargetRelease() string  { return mi.Spec.Release }
func (mi *ModuleInstance) CurrentRelease() string { return mi.Status.Release }

func (mi *ModuleInstance) IsProvisioning() bool {
	return mi.CurrentRelease() == ""
}

func (mi *ModuleInstance) SetProvisioned(reason string, message string) {
	mi.Status.Release = mi.Spec.Release
	mi.SetStatusCondition(ModuleInstanceConditionProvisioned, metav1.ConditionTrue, reason, message)
}

func (mi *ModuleInstance) SetConsumerResource(si SaasInstance) {
	mi.Spec.ConsumerResource = si.Spec.ConsumerResource
}

func (mi *ModuleInstance) SetTargetRelease(localRollout *LocalRollout) {
	if mi.TargetRelease() == localRollout.Spec.Release {
		return
	}

	mi.Spec.PreviousRelease = mi.Spec.Release
	mi.Spec.PreviousDeploymentResources = mi.Spec.DeploymentResources

	mi.Spec.Release = localRollout.Spec.Release
	mi.Spec.DeploymentResources = localRollout.Spec.DeploymentResources
}

func (mi *ModuleInstance) SetScheduledForUpgrade(localRollout *LocalRollout, upgradeScheduledAt metav1.Time) {
	mi.SetTargetRelease(localRollout)
	mi.Spec.UpgradeScheduledAt = upgradeScheduledAt
}

func (mi *ModuleInstance) IsScheduledForUpgrade(now time.Time) bool {
	return mi.TargetRelease() != mi.CurrentRelease() && !mi.Spec.UpgradeScheduledAt.IsZero() && now.Before(mi.Spec.UpgradeScheduledAt.Time)
}

func (mi *ModuleInstance) SetNeedReschedule(reason string, message string) {
	mi.SetStatusCondition(ModuleInstanceConditionNeedReschedule, metav1.ConditionTrue, reason, message)
}

func (mi *ModuleInstance) IsNeedingReschedule() bool {
	needReschedule := mi.IsStatusConditionTrue(ModuleInstanceConditionNeedReschedule)
	return mi.TargetRelease() != mi.CurrentRelease() && needReschedule
}

func (mi *ModuleInstance) IsUpgrading(now time.Time) bool {
	return mi.TargetRelease() != mi.CurrentRelease() && !mi.Spec.UpgradeScheduledAt.IsZero() &&
		!now.Before(mi.Spec.UpgradeScheduledAt.Time) && !mi.IsNeedingReschedule() && !mi.IsUpgradeFailed()
}

func (mi *ModuleInstance) SetUpgraded(reason string, message string) {
	mi.Status.Release = mi.Spec.Release
	mi.SetStatusCondition(ModuleInstanceConditionUpgraded, metav1.ConditionTrue, reason, message)
}

func (mi *ModuleInstance) SetUpgradeFailed(reason string, message string) {
	mi.SetStatusCondition(ModuleInstanceConditionUpgraded, metav1.ConditionFalse, reason, message)
}

func (mi *ModuleInstance) IsUpgradeFailed() bool {
	upgradeFailed := mi.IsStatusConditionFalse(ModuleInstanceConditionUpgraded)
	return mi.TargetRelease() != mi.CurrentRelease() && upgradeFailed
}

func (mi *ModuleInstance) SetStatusCondition(conditionType string, status metav1.ConditionStatus, reason string, message string) {
	meta.SetStatusCondition(&mi.Status.Conditions, metav1.Condition{
		Type:               conditionType,
		Status:             status,
		ObservedGeneration: mi.Generation,
		Reason:             reason,
		Message:            message,
	})
}

func (mi *ModuleInstance) IsStatusConditionTrue(conditionType string) bool {
	return mi.IsStatusConditionCurrentAndEqual(conditionType, metav1.ConditionTrue)
}

func (mi *ModuleInstance) IsStatusConditionFalse(conditionType string) bool {
	return mi.IsStatusConditionCurrentAndEqual(conditionType, metav1.ConditionFalse)
}

func (mi *ModuleInstance) IsStatusConditionCurrentAndEqual(conditionType string, status metav1.ConditionStatus) bool {
	condition := meta.FindStatusCondition(mi.Status.Conditions, conditionType)
	return condition != nil && condition.ObservedGeneration == mi.Generation && condition.Status == status
}
