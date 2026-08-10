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

const (
	// Indicates that the cluster has been provisioned and is running.
	ClusterStateRunning ClusterState = "Running"
	// Indicates that some work is actively being done on the cluster.
	ClusterStateReconciling ClusterState = "Reconciling"
	// Indicates that the cluster is being provisioned.
	ClusterStateProvisioning ClusterState = "Provisioning"
	// Indicates that the cluster is being deleted.
	ClusterStateDeleting ClusterState = "Deleting"
	// Indicates that an error occurred while reconciling the cluster.
	ClusterStateError ClusterState = "Error"
)
