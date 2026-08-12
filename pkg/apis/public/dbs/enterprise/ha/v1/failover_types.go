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
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/common"
	cecoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

const (
	FailoverStateInProgress FailoverState = "InProgress"
	FailoverStateSuccess    FailoverState = "Success"
	FailoverStateFailed     FailoverState = "Failed"

	LabelFailoverDBCluster string = "failover.dbadmin.goog/dbcluster"
	LabelFailoverDBEngine  string = "failover.dbadmin.goog/dbengine"
	LabelFailoverHARole    string = "failover.dbadmin.goog/ha-role"

	AnnotationAutoFailover       string = "dbs.internal.dbadmin.googl/auto-failover"
	AnnotationAutoFailoverReason string = "dbs.internal.dbadmin.googl/auto-failover-reason"
	AnnotationOldPrimary         string = "dbs.internal.dbadmin.googl/old-primary"

	AnnotationAutoFailoverByHAManager           string = "ha-manager"
	AnnotationAutoFailoverByDBClusterReconciler string = "dbcluster-reconciler"

	FailoverFinalizer = "failovers.dbadmin.goog/finalizer"
)

// +kubebuilder:object:generate=true

// FailoverSpec represents the parameters of a single failover operation.
type FailoverSpec struct {
	// DBClusterRef is the DBCluster name to initiate a failover. The `Failover` object must be created in the same namespace as the DBCluster that it references.
	// This field is required for Failover.
	// +required
	DBClusterRef common.DBClusterRef `json:"dbclusterRef,omitempty"`
	// NewPrimary is the standby instance to promote as the new primary. If left empty, the system will automatically pick the best one to failover to.
	// +optional
	NewPrimary string `json:"newPrimary,omitempty"`
}

// +kubebuilder:object:generate=true

// FailoverStatus represents the current state of a failover.
type FailoverStatus struct {
	cecoreapi.EntityStatus `json:",inline"`

	// State is the current state of the failover operation.
	// The values are `InProgress`, `Success`, `Failed`
	// `InProgress` means the failover is still in progress.
	// `Success` means that the failover has completed. It is complete when the new primary instance is successfully promoted.
	// `Failed` means that the operator was unable to promote the new primary instance. The DBCluster might need to be manually repaired.
	// +optional
	// +kubebuilder:validation:Enum=InProgress;Success;Failed
	State FailoverState `json:"state,omitempty"`

	// StartTime is the time that the failover operation started.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// CreateTime is the time that the internal failover workflow mechanism was created.
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`

	// EndTime is the time failover reached its final state.
	// +optional
	EndTime *metav1.Time `json:"endTime,omitempty"`

	// Internal is used by the system controllers. You should not directly depend on the information in this section.
	Internal FailoverInternal `json:"internal,omitempty"`
}

type FailoverState string

// +kubebuilder:object:generate=true

// FailoverInternal is used by the DB controllers to manage the failover.
// These fields should not be used by end users.
type FailoverInternal struct {
	// Phase is used to keep track of the current state of the failover
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=StopPrimary;PromoteStandby;ValidateNewPrimary;UpdateOldPrimaryResources;UpdateNewPrimaryResources;Cleanup;Complete;UpdateStandbys;Recreate;PreComplete
	Phase string `json:"phase,omitempty"`

	// OldPrimary is the instance that was the primary at the start of the failover.
	OldPrimary string `json:"oldPrimary"`

	// NewPrimary is the instance that we are attempting to failover to.
	NewPrimary string `json:"newPrimary"`

	// NewPrimaryZone is the zone that we are attempting to failover to.
	// +kubebuilder:validation:Optional
	NewPrimaryZone string `json:"newPrimaryZone"`

	// Attempt is used for retry logic
	// +kubebuilder:default:=0
	Attempt int `json:"attempt"`

	// LastPhaseChangeTime is the last time the L1 failover changed to a new phase.
	// This is used for timeout logic.
	// +kubebuilder:validation:Optional
	LastPhaseChangeTime *metav1.Time `json:"lastPhaseChangeTime,omitempty"`
}

// Failover represents the L1 Interface for failovers
type Failover interface {
	cecoreapi.Entity
	FailoverSpec() *FailoverSpec
	FailoverStatus() *FailoverStatus
}

type FailoverList interface {
	ctrlclient.ObjectList
	FailoverListItem() []Failover
}

// L2Failover represents the failover state for the L2 area
type L2Failover interface {
	cecoreapi.Entity
	DBCName() string
	SetDBCName(string)
	GetL2FailoverStatus() L2FailoverStatus
}

type L2FailoverStatus interface {
	GetState() FailoverState
	SetState(FailoverState)
	GetStartTime() *metav1.Time
	SetStartTime(*metav1.Time)
	GetCreateTime() *metav1.Time
	SetCreateTime(*metav1.Time)
	GetEndTime() *metav1.Time
	SetEndTime(*metav1.Time)
	GetInternal() *FailoverInternal
}

type L2FailoverList interface {
	ctrlclient.ObjectList
	L2FailoverListItem() []L2Failover
}

type Internal interface {
	GetPhase() string
	SetPhase(string)
	GetOldPrimary() string
	SetOldPrimary(string)
	GetNewPrimary() string
	SetNewPrimary(string)
}
