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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RestorePhase is the phase of a restore.
type RestorePhase string

type InstanceBackupRef string

const (
	// Preparing for platform restore.
	RestorePhaseRestorePreparing RestorePhase = "RestorePreparing"
	// Platform restore in progress.
	RestorePhaseRestoreInProgress RestorePhase = "RestoreInProgress"
	// Preparing for database recovery.
	RestorePhaseRecoveryPreparing RestorePhase = "RecoveryPreparing"
	// Database recovery in progress.
	RestorePhaseRecoveryInProgress RestorePhase = "RecoveryInProgress"
	// Preparing for provisioning.
	RestorePhaseProvisionPreparing RestorePhase = "ProvisionPreparing"
	// Provisioning in progress.
	RestorePhaseProvisionInProgress RestorePhase = "ProvisionInProgress"
	// Provisioning succeeeded.
	RestorePhaseProvisionSucceeded RestorePhase = "ProvisionSucceeded"
	// Restore failed.
	RestorePhaseFailed RestorePhase = "Failed"

	// Underlying platform restore phases.
	PlatformRestorePhaseFailed    RestorePhase = "Failed"
	PlatformRestoreInProgress     RestorePhase = "InProgress"
	PlatformRestorePhaseSucceeded RestorePhase = "Succeeded"
)

// TTL defines the time to wait before auto-delete a completed Restore resource.
var TTL = 2 * time.Hour

// +kubebuilder:object:generate=true

// ClonedDBClusterConfig defines the desired config of a cloned DBCluster.
type ClonedDBClusterConfig struct {
	// The name of cloned DBCluster.
	// +kubebuilder:validation:Required
	DBClusterName string `json:"dbclusterName,omitempty"`

	// To be supported: cross-namespace clone
	// InstanceNamespace string
}

// RestoreStrategy specifies the strategy of InstanceRestore.
// +kubebuilder:validation:Enum=Instance;WorkloadOnly
type RestoreStrategy string

const (
	RestoreStrategyInstance     RestoreStrategy = "Instance"
	RestoreStrategyWorkloadOnly RestoreStrategy = "WorkloadOnly"
)

// +kubebuilder:object:generate=true

// InstanceRestoreSpec defines the desired state of InstanceRestore
type InstanceRestoreSpec struct {
	// SourceDBCluster to restore from.
	// +kubebuilder:validation:required
	SourceDBCluster DBClusterRef `json:"sourceDBCluster,omitempty"`

	// Specify one and only one in "PointInTime", "InstanceBackupRef"

	// Previous point-in-time to restore to.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	// +kubebuilder:validation:optional
	PointInTime *metav1.Time `json:"pointInTime,omitempty"`

	// The InstanceBackup to restore from.
	// InstanceRestore and the source InstanceBackup should be in the same namespace.
	// +kubebuilder:validation:optional
	InstanceBackupRef InstanceBackupRef `json:"instanceBackupRef,omitempty"`

	// Settings for the cloned DBCluster.
	// Omit this field will restore to the DBCluster where backup was taken from.
	// +optional
	ClonedDBClusterConfig *ClonedDBClusterConfig `json:"clonedDBClusterConfig,omitempty"`

	// RestoreStrategy specifies the strategy of InstanceRestore.
	// WorkloadOnly cannot be used together with PointInTime or ClonedDBClusterConfig.
	// +optional
	// +kubebuilder:default:=Instance
	RestoreStrategy *RestoreStrategy `json:"restoreStrategy,omitempty"`
}

// +kubebuilder:object:generate=true

// InstanceRestoreStatus defines the observed state of restore.
type InstanceRestoreStatus struct {
	EntityStatus `json:",inline"`
	Phase        RestorePhase `json:"phase,omitempty"`
	PhaseReason  string       `json:"phaseReason,omitempty"`
	// Create time of the underlying Restore.
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`
	// Completion time of the restore
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
	// Actual point-in-time this restore brings the target instance into.
	// Might be different from value specified in spec.PointInTime.
	// +optional
	RestoredPointInTime *metav1.Time `json:"restoredPointInTime,omitempty"`

	// Name of the instance that will be restored
	// +optional
	RestoredInstanceName string `json:"restoredInstanceName,omitempty"`
}

type InstanceRestore interface {
	Entity
	InstanceRestoreSpec() *InstanceRestoreSpec
	InstanceRestoreStatus() *InstanceRestoreStatus
}

func IsWorkloadOnlyRestoreStrategy(r InstanceRestore) bool {
	return r.InstanceRestoreSpec().RestoreStrategy != nil && *(r.InstanceRestoreSpec().RestoreStrategy) == RestoreStrategyWorkloadOnly
}

func LROPerformPITROperationID(r InstanceRestore) string {
	return fmt.Sprintf("PerformPITR_%s", r.GetUID())
}
