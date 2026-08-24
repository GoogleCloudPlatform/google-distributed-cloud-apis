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
	"strings"

	corev1 "k8s.io/api/core/v1"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	cecoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

const MigrationKind = "Migration"

//+kubebuilder:object:generate=true

// MigrationSpec defines the spec of the migration job.
type MigrationSpec struct {
	// Source is a database server that acts as the source for migration.
	// +kubebuilder:validation:Required
	Source SourceDatabaseServer `json:"source"`

	// Target is a database server that acts as the target of migration.
	// +kubebuilder:validation:Required
	Target TargetDatabaseServer `json:"target"`

	// Control is used to control the state of a migration job.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum:=start;stop;promote
	// +kubebuilder:default:=start
	Control cecoreapi.MigrationControl `json:"control,omitempty"`
}

//+kubebuilder:object:generate=true

// SourceDatabaseServer defines the migration source database server.
type SourceDatabaseServer struct {
	// SourceReference is an object reference to an ExternalServer or DBCluster.
	// +kubebuilder:validation:Required
	SourceReference corev1.ObjectReference `json:"reference"`

	// Databases is a list of databases on the source database server to migrate.
	// If empty all databases will be migrated.
	// +kubebuilder:validation:Optional
	Databases []string `json:"databases,omitempty"`
}

//+kubebuilder:object:generate=true

// TargetDatabaseServer defines the migration target database server.
type TargetDatabaseServer struct {
	// TargetReference is an object reference to a database cluster.
	// Support for ExternalServer may be added in the future.
	// +kubebuilder:validation:Required
	TargetReference corev1.ObjectReference `json:"reference"`
}

//+kubebuilder:object:generate=true

// MigrationStatus is the status of a migration job.
type MigrationStatus struct {
	cecoreapi.EntityStatus `json:",inline"`
}

// Migration is a L1 Migration interface.
type Migration interface {
	cecoreapi.Entity
	MigrationSpec() MigrationSpec
	MigrationStatus() *MigrationStatus
}

type MigrationList interface {
	ctrlclient.ObjectList
	MigrationListItem() []Migration
}

const MigrationReplicationProfileNamePrefix = "mi-"

func MigrationProfileName(miName string) string {
	return MigrationReplicationProfileNamePrefix + miName
}

func MigrationNameFromProfile(profileName string) string {
	return strings.TrimPrefix(profileName, MigrationReplicationProfileNamePrefix)
}
