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

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// ServiceDescription represents a description of a Marketplace service.
type ServiceDescription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDescriptionSpec   `json:"spec,omitempty"`
	Status ServiceDescriptionStatus `json:"status,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.version`
// Represents a single version of a Marketplace service.
type ServiceVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceVersionSpec   `json:"spec"`
	Status ServiceVersionStatus `json:"status,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// Represents a single group of Marketplace services.
type ServiceCatalog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceCatalogSpec   `json:"spec,omitempty"`
	Status ServiceCatalogStatus `json:"status,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// Represents a ServiceCatalog to Project mapping relations.
type ServiceCatalogBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceCatalogBindingSpec   `json:"spec,omitempty"`
	Status ServiceCatalogBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceDescription` custom resources.
type ServiceDescriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceDescription `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceVersion` custom resources.
type ServiceVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceVersion `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceCatalog` custom resources.
type ServiceCatalogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceCatalog `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceCatalogBinding` custom resources.
type ServiceCatalogBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceCatalogBinding `json:"items"`
}

// Provides the overall status of a global service description.
type ServiceDescriptionStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ServiceDescriptionZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a Service Description rolling out to a particular zone.
type ServiceDescriptionZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus ServiceDescriptionReplicaStatus `json:"replicaStatus,omitempty"`
}

// Provides the overall status of a global service version.
type ServiceVersionStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ServiceVersionZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a Service Version rolling out to a particular zone.
type ServiceVersionZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus ServiceVersionReplicaStatus `json:"replicaStatus,omitempty"`
}

// Provides the overall status of a global service catalog.
type ServiceCatalogStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ServiceCatalogZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a Service Catalog rolling out to a particular zone.
type ServiceCatalogZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus ServiceCatalogReplicaStatus `json:"replicaStatus,omitempty"`
}

// Provides the overall status of a global service catalog binding.
type ServiceCatalogBindingStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ServiceCatalogBindingZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a Service Catalog rolling out to a particular zone.
type ServiceCatalogBindingZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus ServiceCatalogBindingReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ServiceDescription{},
		&ServiceVersion{},
		&ServiceCatalog{},
		&ServiceCatalogBinding{},
		&ServiceDescriptionList{},
		&ServiceVersionList{},
		&ServiceCatalogList{},
		&ServiceCatalogBindingList{},
	)
}
