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

package v2

const (
	// ReadyStatusType is the condition type that indicates the overall
	// readiness of a resource.
	ReadyStatusType = "Ready"
	// PropagatedStatusType is the condition type that indicates the resource
	// is fully propagated to all user clusters.
	PropagatedStatusType = "Propagated"
)

const (
	// ReadyCondition is the condition type that indicates the overall
	// readiness of a resource.
	ReadyCondition = "Ready"
	// InvalidReason indicates the resource is invalid.
	InvalidReason = "Invalid"
	// ReconciledReason indicates the resource is reconciled successfully.
	ReconciledReason = "Reconciled"
	// ReconcileSkippedReason indicates the resource is skipped reconciled successfully.
	ReconcileSkippedReason = "ReconcileSkipped"
	// KubernetesClientErrorReason indicates an error is returned when the k8s client
	// performs read or write action.
	KubernetesClientErrorReason = "KubernetesClientError"
	// InternalErrorReason indicates an error from the controller itself.
	InternalErrorReason = "InternalError"
)

const (
	// PropagatedCondition is the condition type that indicates the resource is fully propagated
	// to all user clusters.
	PropagatedCondition = "Propagated"
	// PropagatedReason indicates that the resource was successfully propagated.
	PropagatedReason = "Propagated"
	// PropagatingReason indicates that the resource was not successfully
	// propagated to all clusters.
	PropagatingReason = "Propagating"
	// PropagationFailedReason indicates that the resource was not successfully
	// propagated.
	PropagationFailedReason = "PropagationFailed"
	// PropagationUnitFailedReason indicates that the resource units reconciliation failed.
	PropagationUnitFailedReason = "PropagationUnitFailed"
)

// Following group contains terminating condition types and reasons.
const (
	// TerminatingCondition is the condition type that indicates that the resource is being
	// deleted.
	TerminatingCondition = "Terminating"
	// TerminatingFailedReason indicates the resource terminating process failed.
	TerminatingFailedReason = "TerminatingFailed"
	// TerminatingReason indicates the resource is being deleted,
	// the condition message explains what blocks resource deletion.
	TerminatingReason = "Terminating"
)
