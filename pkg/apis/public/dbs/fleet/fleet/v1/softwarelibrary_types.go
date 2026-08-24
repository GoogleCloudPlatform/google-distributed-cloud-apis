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

// +kubebuilder:object:generate=true
package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +gdcloud:manifest:relevant=false,oc=dbs
type SoftwareLibrary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DatabaseEngines []DatabaseEngine `json:"databaseEngines,omitempty"`
}

//+kubebuilder:object:generate=true

type DatabaseEngine struct {
	Name EngineType `json:"name"`

	// Whether the database engine is disallowed for new cluster creation
	IsDisabled bool `json:"isDisabled"`

	// A list of available database versions supported for the database engine
	DatabaseVersions []Database `json:"databaseVersions"`
}

//+kubebuilder:object:generate=true

type Database struct {
	// Specify a user-friendly display name
	VersionDisplayName string `json:"versionDisplayName"`

	// The major database version
	MajorVersion string `json:"majorVersion"`

	//The minor database version
	MinorVersion string `json:"minorVersion,omitempty"`

	// Edition of a database.
	// +optional
	Edition string `json:"edition,omitempty"`

	// Whether the database major version is disallowed for new cluster creation
	IsDisabled bool `json:"isDisabled"`
}

// +kubebuilder:object:root=true

// SoftwareLibraryList contains a list of SoftwareLibrary
type SoftwareLibraryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SoftwareLibrary `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SoftwareLibrary{}, &SoftwareLibraryList{})
}

const SoftwareLibraryName = "available-database-engines"
