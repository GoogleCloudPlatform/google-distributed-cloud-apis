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
// +gdcloud:manifest:relevant=true,oc=dns,component=dns,entities="managed-zones"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:managed-dns-project-admin"
// +gdcloud:manifest:rbac="describe,list:managed-dns-project-viewer"

// ManagedDNSZone represents a managed DNS zone.
type ManagedDNSZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedDNSZoneSpec   `json:"spec,omitempty"`
	Status ManagedDNSZoneStatus `json:"status,omitempty"`
}

var _ v1alpha1.MuxResourceInterface = &ManagedDNSZone{}

// +kubebuilder:object:root=true

// ManagedDNSZoneList represents a list of managed DNS zones.
type ManagedDNSZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedDNSZone `json:"items"`
}

// ManagedDNSZoneStatus provides the overall status of a ManagedDNSZone.
type ManagedDNSZoneStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of rollout statuses for each GDC air-gapped zone that the
	// resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ManagedDNSZoneZoneStatus `json:"zones,omitempty"`
}

// ManagedDNSZoneZoneStatus provides the status of a ManagedDNSZone rolling out
// to a particular GDC air-gapped zone.
type ManagedDNSZoneZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the GDC
	// air-gapped zone. Any condition within the field that has an
	// .observedGeneration less than .rolloutStatus.replicaGeneration is out of
	// date.
	ReplicaStatus ManagedDNSZoneReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ManagedDNSZone{},
		&ManagedDNSZoneList{},
	)
}
