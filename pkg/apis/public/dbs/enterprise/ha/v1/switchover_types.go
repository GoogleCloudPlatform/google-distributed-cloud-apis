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
	cecoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	LabelSwitchoverDBCluster string = "switchover.dbadmin.goog/dbcluster"
	LabelSwitchoverDBEngine  string = "switchover.dbadmin.goog/dbengine"
	SwitchoverFinalizer      string = "switchovers.dbadmin.goog/finalizer"
)

// +kubebuilder:object:generate=true

// SwitchoverSpec defines the desired state of Switchover
type SwitchoverSpec struct {
	// DBClusterRef is the DBCluster name to initiate a switchover. The `Switchover` object must be created in the same namespace as the DBCluster that it references.
	// This field is required for Switchover.
	// +required
	DBClusterRef common.DBClusterRef `json:"dbclusterRef,omitempty"`
	// NewPrimaryInstance is the standby instance to switch with the current primary.
	// Deprecated: Please use NewPrimary field instead
	// +optional
	// nullon(dbs-fleet)
	NewPrimaryInstance string `json:"newPrimaryInstance,omitempty"`
	// NewPrimary is the standby instance to switch with the current primary.
	NewPrimary string `json:"newPrimary,omitempty"`
}

// +kubebuilder:object:generate=true

// SwitchoverStatus defines the observed state of switchover.
type SwitchoverStatus struct {
	cecoreapi.EntityStatus `json:",inline"`
	// State is the current state of the switchover operation.
	// The values are `InProgress`, `Success`, `Failed_RollbackInProgress`, `Failed_RollbackSuccess`, `Failed_RollbackFailed`
	// `InProgress` means the switchover is still in progress.
	// `Success` means that the switchover has completed.
	// `Failed_RollbackInProgress` means that the operator was unable to promote the new primary instance, and is attempting to restart the old primary instance.
	// `Failed_RollbackSuccess` means that the operator was unable to promote the new primary instance, and successfully restarted the old primary instance.
	// `Failed_RollbackFailed` means that  the operator was unable to promote the new primary instance, and were not able to restart the old primary instance. The DBCluster might need to be manually repaired.
	// +optional
	// +kubebuilder:validation:Enum=InProgress;Success;Failed_RollbackInProgress;Failed_RollbackSuccess;Failed_RollbackFailed
	State cecoreapi.SwitchoverState `json:"state,omitempty"`

	// StartTime is the time that the switchover operation started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CreateTime is the time that the internal switchover workflow mechanism was created.
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`

	// EndTime is the time switchover reached its final state.
	// +optional
	EndTime *metav1.Time `json:"endTime,omitempty"`

	// Internal is used by the system controllers. You should not directly depend on the information in this section.
	Internal SwitchoverInternal `json:"internal,omitempty"`
}

// SwitchoverInternal is used by the DB controllers to manage the switchover.
// These fields should not be used by end users.
type SwitchoverInternal struct {
	// Phase is used to keep track of the current state of the switchover
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=UpdateDbcluster;StopPrimary;PromoteStandby;ValidateNewPrimary;UpdateOldPrimaryResources;UpdateNewPrimaryResources;UpdateOldPrimaryConfigs;Complete;SyncOldPrimary;StartOldPrimary;PreSuccess;RepointStandbys;RollbackPrimary;RollbackStandbys
	Phase string `json:"phase,omitempty"`

	// OldPrimary is the instance that was the primary at the start of the switchover.
	OldPrimary string `json:"oldPrimary"`

	// NewPrimary is the instance that we are attempting to switchover to.
	NewPrimary string `json:"newPrimary"`

	// Attempt is used for retry logic
	// +kubebuilder:default:=0
	Attempt int `json:"attempt"`
}

// Switchover represents the L1 Interface for switchovers
type Switchover interface {
	cecoreapi.Entity
	SwitchoverSpec() *SwitchoverSpec
	SwitchoverStatus() *SwitchoverStatus
}

type SwitchoverList interface {
	ctrlclient.ObjectList
	SwitchoverListItem() []Switchover
}
