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
	rbacv1 "k8s.io/api/rbac/v1"

	identityv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/identity/v1"
)

type CustomRoleScopeType = identityv1.CustomRoleScopeType
type StageType = identityv1.StageType
type RoleType = identityv1.RoleType

const (
	OrganizationType = identityv1.OrganizationType
	ProjectType      = identityv1.ProjectType
	Alpha            = identityv1.Alpha
	Beta             = identityv1.Beta
	GA               = identityv1.GA
	Disabled         = identityv1.Disabled
	Role             = identityv1.Role
	ClusterRole      = identityv1.ClusterRole
	ProjectRole      = identityv1.ProjectRole
	OrganizationRole = identityv1.OrganizationRole
)

// Defines the CustomRole data in the `ClusterRoleTemplate` resource
type CustomRoleSpec struct {
	Metadata    CustomRoleMetadata  `json:"metadata,omitempty"`
	ZonalRules  []rbacv1.PolicyRule `json:"zonalRules,omitempty"`
	GlobalRules []rbacv1.PolicyRule `json:"globalRules,omitempty"`
}

func (spec *CustomRoleSpec) GetMetadata() identityv1.ICustomRoleMetadata {
	return &spec.Metadata
}

func (spec *CustomRoleSpec) GetZonalRules() []rbacv1.PolicyRule {
	return spec.ZonalRules
}

func (spec *CustomRoleSpec) GetGlobalRules() []rbacv1.PolicyRule {
	return spec.GlobalRules
}

// Represents the data necessary to create a Custom Role
type CustomRoleMetadata struct {
	// scope of the custom role created which can either be organization or project
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=organization;project
	Scope CustomRoleScopeType `json:"scope,omitempty"`
	// namespace of the role (optional)
	// only required for role deployment if:
	// case 1: when scope is project then it denotes the project namespaces
	// case 2: when scope is project and roleNamespaces contain ['*'] then it denotes all project namespaces
	// case 3: when scope is organization and deployment roleType is role not clusterRole then it denotes literal namespaces
	// +optional
	RoleNamespaces []string `json:"roleNamespaces,omitempty"`
	// title is a friendly title for the role, such as "My Company Admin".
	// +kubebuilder:validation:Required
	Title string `json:"title,omitempty"`
	// description is a short description of the role, such as "My custom role description".
	// +kubebuilder:validation:Required
	Description string `json:"description,omitempty"`
	// id is the name of the role, such as "my-company-admin".
	// +kubebuilder:validation:Required
	Id string `json:"id,omitempty"`
	// stage indicates the stage of a role in the launch lifecycle which can either be [ALPHA, BETA, GA, DISABLED]
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ALPHA;BETA;GA;DISABLED
	Stage StageType `json:"stage,omitempty"`
}

func (metadata *CustomRoleMetadata) GetScope() CustomRoleScopeType {
	return metadata.Scope
}

func (metadata *CustomRoleMetadata) GetRoleNamespaces() []string {
	return metadata.RoleNamespaces
}

func (metadata *CustomRoleMetadata) GetTitle() string {
	return metadata.Title
}

func (metadata *CustomRoleMetadata) GetDescription() string {
	return metadata.Description
}

func (metadata *CustomRoleMetadata) GetId() string {
	return metadata.Id
}

func (metadata *CustomRoleMetadata) GetStage() StageType {
	return metadata.Stage
}

// Provides the information of converted role template
type PropagationInfo struct {
	// name of the role
	// + optional
	RoleName string `json:"roleName,omitempty"`
	// type of the role, it can be [role, clusterRole, projectRole, organizationRole]
	// + optional
	// +kubebuilder:validation:Enum=role;clusterRole;projectRole;organizationRole
	RoleType RoleType `json:"roleType,omitempty"`
	// namespaces of the role where role deployment will occur
	// + optional
	Namespaces []string `json:"namespaces,omitempty"`
}
