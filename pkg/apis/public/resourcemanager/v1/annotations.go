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
	// RBACScopeSelector matches a set of preset roles whose RBACScopeLabel in the list.
	// For example, if the cluster's RBACScopeSelector is ["user", "system"], and the
	// role's RBACScopeLabel is "user", then the RBACScopeSelector matches this role.
	// An empty selector matches all preset roles which do not have RBACScopeLabel.
	RBACScopeSelector = Group + "/rbac-scope-selector"

	// IsSystemClusterAnnotation is set by Controller to identify system clusters.
	// Users should not change this label's value.
	// This annotation is only populated on system clusters.
	IsSystemClusterAnnotation = Group + "/is-system-cluster"

	// ClusterSelectorAnnotation is set on a v1 Project object to keep track of the
	// ClusterSelector from a v1alpha1 Project. This annotation is for conversions
	// Project versions and users should not update the value.
	ClusterSelectorAnnotation = Group + "/ClusterSelector"

	// ShadowProjectSourceAnnotation is set on a Project to indicate that it was created from a Shadow Project.
	ShadowProjectSourceAnnotation = Group + "/shadow-project-source"
)
