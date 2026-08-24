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

	globalv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=siem,component=siem,entities="instances"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:siem-instance-admin"
// +gdcloud:manifest:rbac="describe,list:siem-instance-viewer"

// Represents a managed instance of an organization SIEM service.
// +genclient
type ManagedSIEM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired spec of the ManagedSIEM instance.
	Spec ManagedSIEMSpec `json:"spec,omitempty"`
	// The most recently observed status of the ManagedSIEM instance.
	Status ManagedSIEMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of ManagedSIEM instances.
type ManagedSIEMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedSIEM `json:"items"`
}

// Represents the overall status of an ManagedSIEM instance.
type ManagedSIEMStatus struct {
	// Standard status field of a global resource. Used by the mux controller.
	globalv1alpha1.MuxStatus `json:",inline"`

	// Standard status field of a global resource. Used by the mux controller.
	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []ManagedSIEMZoneStatus `json:"zones,omitempty"`

	// The global siem console fqdn that points to the primary replica.
	Fqdn string `json:"primaryFqdn,omitempty"`

	// The zone ID of the replica currently holding global primary leadership.
	PrimaryZone string `json:"primaryZone,omitempty"`
}

// Represents the zone-specific status of a ManagedSIEM instance.
type ManagedSIEMZoneStatus struct {
	// Standard status field of a global resource. Used by the mux controller.
	globalv1alpha1.ZoneStatus `json:",inline"`

	// Standard status field: the reconciliation status of the replica collected
	// from the zone automatically by the mux controller.
	ReplicaStatus ManagedSIEMReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ManagedSIEM{},
		&ManagedSIEMList{},
	)
}
