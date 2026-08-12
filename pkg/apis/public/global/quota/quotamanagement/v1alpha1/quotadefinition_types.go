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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// Defines the quota metrics, quotas and api rules for a service.
// This roughly corresponds to the "quota" section of a OnePlatform
// config without the default quota values.
type QuotaDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuotaDefinitionSpec   `json:"spec"`
	Status QuotaDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a list of QuotaDefinitions
type QuotaDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuotaDefinition `json:"items"`
}

type QuotaDefinitionStatus struct {
	v1alpha1.MuxStatus `json:",inline"` // embed the duck type for global status

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []QuotaDefinitionZoneStatus `json:"zones,omitempty"`
}

type QuotaDefinitionZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"` // embed the duck type for zone status

	// The reconciliation status of the replica collected from the zone.
	ReplicaStatus QuotaDefinitionReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&QuotaDefinition{}, &QuotaDefinitionList{})
}
