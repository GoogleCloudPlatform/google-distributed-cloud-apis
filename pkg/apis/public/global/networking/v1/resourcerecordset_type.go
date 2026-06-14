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

	"gke-internal.googlesource.com/private-cloud/pkg/apis/common/global/v1alpha1"
)

const (
	// LabelDNSZone is the label key used to associate a ResourceRecordSet with its ManagedDNSZone.
	LabelDNSZone = "clouddns.private.gdc.goog/dnszone"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=dns,component=dns,entities="record-sets"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:managed-dns-project-admin"
// +gdcloud:manifest:rbac="describe,list:managed-dns-project-viewer"

// ResourceRecordSet represents a resource record set.
type ResourceRecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceRecordSetSpec   `json:"spec,omitempty"`
	Status ResourceRecordSetStatus `json:"status,omitempty"`
}

var _ v1alpha1.MuxResourceInterface = &ResourceRecordSet{}

// +kubebuilder:object:root=true

// ResourceRecordSetList represents a list of resource record sets.
type ResourceRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ResourceRecordSet `json:"items"`
}

// ResourceRecordSetStatus provides the overall status of a ResourceRecordSet.
type ResourceRecordSetStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of rollout statuses for each GDC air-gapped zone that the object
	// is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ResourceRecordSetZoneStatus `json:"zones,omitempty"`
}

// ResourceRecordSetZoneStatus provides the status of a ResourceRecordSet
// rolling out to a particular GDC air-gapped zone.
type ResourceRecordSetZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the GDC
	// air-gapped zone. Any condition within the field that has an
	// .observedGeneration less than .rolloutStatus.replicaGeneration is out of
	// date.
	ReplicaStatus ResourceRecordSetReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ResourceRecordSet{},
		&ResourceRecordSetList{},
	)
}
