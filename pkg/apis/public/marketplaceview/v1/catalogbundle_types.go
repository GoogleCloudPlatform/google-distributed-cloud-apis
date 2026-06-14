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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	globalmarketplacev1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/global/marketplace/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// +genclient
// CatalogBundle represents a read-only view of Service Catalogs, Descriptions, and Versions available in a project.
type CatalogBundle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CatalogBundleSpec `json:"spec,omitempty"`
}

// +k8s:openapi-gen=true
// CatalogBundleSpec defined the CatalogBundle.
type CatalogBundleSpec struct {
	ServiceCatalog      globalmarketplacev1.ServiceCatalogReplica       `json:"serviceCatalog,omitempty"`
	ServiceDescriptions []globalmarketplacev1.ServiceDescriptionReplica `json:"serviceDescriptions,omitempty"`
	ServiceVersions     []globalmarketplacev1.ServiceVersionReplica     `json:"serviceVersions,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// CatalogBundleList represents a collection of `CatalogBundle` resource.
type CatalogBundleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogBundle `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&CatalogBundle{},
		&CatalogBundleList{},
	)
}
