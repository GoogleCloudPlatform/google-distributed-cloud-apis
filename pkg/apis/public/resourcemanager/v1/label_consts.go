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
	// ProjectNamespaceLabel is set by Controller to identify the namespace of
	// the project for a propagated namespace. Users should not change this
	// label's value.
	ProjectNamespaceLabel = Group + "/project-namespace"

	// ProjectNameLabel is set by Controller to identify the name of the project
	// for a propagated namespace. Users should not change this label's value.
	ProjectNameLabel = Group + "/project-name"

	// ParentProjectLabel is set by a Service Controller to identify a shadow
	// project's parent project.
	ParentProjectLabel = Group + "/parent-project"

	// ServiceShortNameLabel is set by a Service Controller to identify a shadow
	// project's service project.
	ServiceShortNameLabel = Group + "/service-short-name"

	// RBACScopeLabel is set in ProjectRole/OrganizationRole CRD to identify
	// the scope of a preset projectrole or organizationrole.
	RBACScopeLabel = Group + "/rbac-scope"

	// SystemResourceLabel denotes a resource which is a system resource.
	//
	// Labeling an object as a system resource may influence the configuration
	// of the object or dependents of the object.
	//
	// How an object is affected by this label is subject to the object's
	// reconciler.
	SystemResourceLabel = Group + "/system-resource"

	// ProjectBindingForUserProjectLabel is set by UI to identify the ProjectBinding
	// creation is from users' project creation operations.
	ProjectBindingForUserProjectLabel = Group + "/projectbinding-for-user-project"

	// TagLabelPrefix is the prefix for the tag generated label key.
	TagLabelPrefix = "tag." + Group
)
