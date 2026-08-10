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

type OrgAPIServer string

const (
	// ManagementPlaneAPIServer is an api server in ArchV2.
	ManagementPlaneAPIServer OrgAPIServer = "Management"
	// InClusterAPIServer is an api server in ArchV2.
	InClusterAPIServer OrgAPIServer = "InCluster"
	// All api servers in the the current architecture.
	AllAPIServers OrgAPIServer = "All"
)

const (
	// ZoneLabel is a label for a global replica resource indicating
	// which zone the replica is located.
	ZoneLabel = "global.private.gdc.goog/zone"
	// GlobalResourceAnnotation is an annotation for a global replica resource indicating
	// the name of global resource.
	GlobalResourceAnnotation = "global.private.gdc.goog/global-resource"
	// StopSyncLabel is set to indicate a replica should stop being rolled out to a zone.
	StopSyncLabel = "rollout.global.private.gdc.goog/stop-sync"
	// This annotation specifies which target api server the global resource is rolled out
	// in an archV2 org, this annotation is only available in archV2 and plus.
	TargetAPIServerAnnotation = "global.private.gdc.goog/target-api-server"
	// DelegateReadyStatusConditionAnnotation indicates that the aggregate status
	// of the global CR should not be populated by the MUX controller.
	DelegateReadyStatusConditionAnnotation = "global.private.gdc.goog/delegate-ready-status-condition"
)

const (
	// ReplicaFinalizerFormat is global.private.gdc.goog/finalizer
	ReplicaFinalizer = "global.private.gdc.goog/finalizer"
)
