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

// Common condition types used in the rollout status of a global MUX resource.
const (
	// Indicates that the replica of a global MUX resource for a zone has been
	// updated to the current revision in the global API server.
	ReplicatedCondition = "Replicated"

	// Indicates that the replica of a global MUX resource for a zone has been
	// synchronized to the zonal API server.
	SyncedCondition = "Synced"

	// Indicates that the global MUX resource is fully replicated, synchronized,
	// and internally ready across all configured zones.
	ReadyCondition = "Ready"
)

// Common reasons used in the status of a global MUX resource.
const (
	// RolloutFailedReason indicates the global resource rollout failed.
	RolloutFailedReason = "RolloutFailed"
	// RolloutSucceedReason indicates the global resource rollout succeeded.
	RolloutSucceedReason = "RolloutSucceeded"

	// CurrentReason indicates a replica is currently synced with a global MUX resource.
	CurrentReason = "Current"
	// NotCurrentReason indicates a replica is not currently synced with a global MUX resource.
	NotCurrentReason = "NotCurrent"
	// ZonalReplicaDeleted indicates a zonal replica was deleted by something other than MUX. MUX's
	// ability to detect when this happens is best-effort; its purpose is primarily to allow MUX to
	// filter out irrelevant events when measuring its latency SLI.
	ZonalReplicaDeletedReason = "ZonalReplicaDeleted"

	// PausedReason indicates a replica is not synced with a global MUX resource due to syncing with the zone being paused.
	PausedReason = "Paused"

	// DeletingReason indicates the global resource is terminating.
	DeletingReason = "Deleting"
	// ReplicatingReason indicates the global resource is waiting for zonal replicas to sync.
	ReplicatingReason = "Replicating"
	// ReconcilingReason indicates the global resource is synced but waiting for remote controllers to ready.
	ReconcilingReason = "Reconciling"
	// ReadyReason indicates the global resource is fully synced and ready everywhere.
	ReadyReason = "Ready"
)
