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

/*
Copyright 2022.
*/
package v1

import (
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DatabaseVersionSpec defines the desired state of DatabaseVersion
type DatabaseVersionSpec struct {
	// Specify a user-friendly display name
	VersionDisplayName string `json:"versionDisplayName"`

	// The database engine name of the database version, should be one of [PostgreSQL, Oracle, AlloyDBOmni]
	// +kubebuilder:validation:Enum=PostgreSQL;Oracle;AlloyDBOmni
	DatabaseEngine EngineType `json:"databaseEngine"`

	// The major database version
	MajorVersion string `json:"majorVersion"`

	// The default minor version of this major version for new database cluster creation
	DefaultMinorVersion string `json:"defaultMinorVersion,omitempty"`

	// The list of available minor versions
	MinorDatabaseVersions []MinorDatabaseVersion `json:"minorDatabaseVersions"`
}

type EngineType string

var AllEngineTypes = []EngineType{PostgreSQL, Oracle, AlloyDBOmni}

const (
	PostgreSQL EngineType = "PostgreSQL"
	// TODO(b/290635347): remove PostgreSQLShort if postgresql totally replaces postgres in the code base
	PostgreSQLShort EngineType = "Postgres"
	Oracle          EngineType = "Oracle"
	AlloyDBOmni     EngineType = "AlloyDBOmni"
)

type MinorDatabaseVersion struct {
	// Specify a user-friendly display name
	VersionDisplayName string `json:"versionDisplayName"`

	// The minor database version
	Version string `json:"version"`

	// Edition of a database.
	// +optional
	Edition string `json:"edition,omitempty"`

	// ODS build number for database images provided by Google
	ODSBuildNumber string `json:"ODSBuildNumber,omitempty"`

	// Image for each data plane components
	// +optional
	Components []DataPlaneComponent `json:"components"`
}

type DataPlaneComponent struct {
	// Name of this component. For example database, memoryagent
	// +required
	Name string `json:"name,omitempty"`

	// Image path of this component.
	// +required
	Uri string `json:"uri,omitempty"`
}

type DatabaseVersionStatus struct {
	occoreapi.EntityStatus `json:",inline"`
	// the database image paths of each minor version
	Images map[string]string `json:"Images,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster

// DatabaseVersion is the Schema for the databaseversion API
type DatabaseVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseVersionSpec   `json:"spec,omitempty"`
	Status DatabaseVersionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DatabaseVersionList contains a list of DatabaseVersion
type DatabaseVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DatabaseVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseVersion{}, &DatabaseVersionList{})
}

func (mdv *DatabaseVersion) EntityStatus() *occoreapi.EntityStatus {
	return &mdv.Status.EntityStatus
}
