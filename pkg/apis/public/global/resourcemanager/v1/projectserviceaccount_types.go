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

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"

// Represents a ServiceAccount associated with Projects in all zones.
// +genclient
type ProjectServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectServiceAccountSpec   `json:"spec,omitempty"`
	Status ProjectServiceAccountStatus `json:"status,omitempty"`
}

var _ v1alpha1.MuxResourceInterface = &ProjectServiceAccount{}

// +kubebuilder:object:root=true

// Represents a collection of ProjectServiceAccounts.
type ProjectServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ProjectServiceAccount `json:"items"`
}

// Provides the overall status of a ProjectServiceAccount.
type ProjectServiceAccountStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	// + optional
	Zones []ProjectServiceAccountZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a ProjectServiceAccount rolling out to a particular zone.
type ProjectServiceAccountZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	// + optional
	ReplicaStatus ProjectServiceAccountReplicaStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ProjectServiceAccount{},
		&ProjectServiceAccountList{},
	)
}
