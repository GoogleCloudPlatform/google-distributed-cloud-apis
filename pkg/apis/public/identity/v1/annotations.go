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
	// AnnotationSkipSyncStandardClusterRole skips the sync of the
	// StandardClusterRole to the eventual ClusterRole if the ClusterRole got
	// updated. But it will keep sync/recover the ClusterRole if it's deleted.
	// For example, if the StandardClusterRole has this annotation set to true
	// from the corresponding StandardClusterRoleTemplate, then the IAM
	// controller that remotely watches the propagated ClusterRole will ignore
	// the update event on the ClusterRole modified by another controller and
	// skip the sync to the original StandardClusterRole and template.
	AnnotationSkipSyncStandardClusterRole = Group + "/skip-sync-standardclusterrole"
)
