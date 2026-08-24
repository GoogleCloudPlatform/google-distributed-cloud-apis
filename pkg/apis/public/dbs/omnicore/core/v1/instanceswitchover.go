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
	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SwitchoverStateInProgress                SwitchoverState = "InProgress"
	SwitchoverStateSuccess                   SwitchoverState = "Success"
	SwitchoverStateFailed_RollbackInProgress SwitchoverState = "Failed_RollbackInProgress"
	SwitchoverStateFailed_RollbackSuccess    SwitchoverState = "Failed_RollbackSuccess"
	SwitchoverStateFailed_RollbackFailed     SwitchoverState = "Failed_RollbackFailed"
)

const (
	SwitchoverPhaseInitial                   = ""
	SwitchoverPhaseValidateNewPrimary        = "ValidateNewPrimary"
	SwitchoverPhaseStopPrimary               = "StopPrimary"
	SwitchoverPhaseRollbackPrimary           = "RollbackPrimary"
	SwitchoverPhaseRollbackStandbys          = "RollbackStandbys"
	SwitchoverPhasePromoteStandby            = "PromoteStandby"
	SwitchoverPhaseUpdateNewPrimaryResources = "UpdateNewPrimaryResources"
	SwitchoverPhaseUpdateOldPrimaryResources = "UpdateOldPrimaryResources"
	SwitchoverPhaseUpdateOldPrimaryConfigs   = "UpdateOldPrimaryConfigs"
	SwitchoverPhaseSyncOldPrimary            = "SyncOldPrimary"
	SwitchoverPhaseRepointStandbys           = "RepointStandbys"
	SwitchoverPhasePreSuccess                = "PreSuccess"
	SwitchoverPhaseUpdateDbcluster           = "UpdateDbcluster"
	SwitchoverPhaseComplete                  = "Complete"
)

type SwitchoverState string

// +kubebuilder:object:generate=true

// SwitchoverSpec represents the parameters of a single switchover operation.
type InstanceSwitchoverSpec struct {
	// DBClusterRef is the dbcluster name within the same namespace to initiate a switchover.
	// +required
	DBClusterRef common.DBClusterRef `json:"dbclusterRef,omitempty"`
	// NewPrimary is the standby instance to switch with the current primary.
	// +optional
	NewPrimary string `json:"newPrimary,omitempty"`
}

// +kubebuilder:object:generate=true

// SwitchoverStatus represents the current state of a switchover.
type InstanceSwitchoverStatus struct {
	EntityStatus `json:",inline"`

	// State is the current state of the switchover operation.
	// +optional
	// +kubebuilder:validation:Enum=InProgress;Success;Failed_RollbackInProgress;Failed_RollbackSuccess;Failed_RollbackFailed
	State SwitchoverState `json:"state,omitempty"`

	// StartTime is the time switchover started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CreateTime is the time the underlying switchover was created.
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`

	// EndTime is the time switchover reached its final state.
	// +optional
	EndTime *metav1.Time `json:"endTime,omitempty"`

	// Internal is used by the DBS controllers. Users should not directly depend on the information in this section.
	Internal SwitchoverInternal `json:"internal,omitempty"`
}

func (in *InstanceSwitchoverStatus) GetState() SwitchoverState {
	return in.State
}

func (in *InstanceSwitchoverStatus) SetState(state SwitchoverState) {
	in.State = state
}

func (in *InstanceSwitchoverStatus) GetInternal() *SwitchoverInternal {
	return &in.Internal
}

// SwitchoverInternal is used by the DB controllers to manage the switchover.
// These fields should not be used by end users.
type SwitchoverInternal struct {
	// Phase is used to keep track of the current state of the switchover
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=UpdateDbcluster;StopPrimary;PromoteStandby;ValidateNewPrimary;UpdateOldPrimaryResources;UpdateNewPrimaryResources;UpdateOldPrimaryConfigs;Complete;SyncOldPrimary;StartOldPrimary;PreSuccess;RepointStandbys;RollbackPrimary;RollbackStandbys
	Phase string `json:"phase,omitempty"`

	// OldPrimary is the instance that was the primary at the start of the switchover.
	OldPrimary string `json:"oldPrimary,omitempty"`

	// NewPrimary is the instance that we are attempting to switchover to.
	NewPrimary string `json:"newPrimary,omitempty"`
}

type InstanceSwitchover interface {
	Entity
	InstanceSwitchoverSpec() *InstanceSwitchoverSpec
	InstanceSwitchoverStatus() *InstanceSwitchoverStatus
}
