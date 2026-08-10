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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +gdcloud:manifest:relevant=false,oc=m4gdc
// +genclient
// VMMigration CR
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type VMMigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VMMigrationSpec   `json:"spec"`
	Status            VMMigrationStatus `json:"status,omitempty"`
}

type VMMigrationSpec struct {
	DiscoveredVMRef corev1.LocalObjectReference `json:"discoveredVMRef"`          // Reference to the original DiscoveredVM
	ReadyToMigrate  bool                        `json:"readyToMigrate,omitempty"` // Flag to trigger migraiton
}

type VMMigrationStatus struct {
	Phase VMMigrationPhase `json:"phase,omitempty"` // Overall migration phase
	// TODO(unik): Uncomment after design is final, and replication conditions are committed
	// Conditions []MigrationCondition `json:"conditions,omitempty"` // Detailed conditions
	// Replication VMReplicationStatus  `json:"replication,omitempty"` // Nested replication status
}

type VMMigrationPhase string

const (
	VMMigrationPhasePending     VMMigrationPhase = "Pending"
	VMMigrationPhaseValidating  VMMigrationPhase = "Validating"
	VMMigrationPhaseReplicating VMMigrationPhase = "Replicating"
	VMMigrationPhaseAdapting    VMMigrationPhase = "Adapting"
	VMMigrationPhaseTurningUp   VMMigrationPhase = "TurningUp"
	VMMigrationPhaseCompleted   VMMigrationPhase = "Completed"
	VMMigrationPhaseFailed      VMMigrationPhase = "Failed"
)

// MigrationCondition
// type MigrationCondition struct {
//  Type               MigrationConditionType `json:"type"`
// 	Status             corev1.ConditionStatus `json:"status"`
// 	LastUpdateTime     metav1.Time            `json:"lastUpdateTime,omitempty"`
// 	LastTransitionTime metav1.Time            `json:"lastTransitionTime,omitempty"`
// 	Reason             string                 `json:"reason,omitempty"`
// 	Message            string                 `json:"message,omitempty"`
// }

// +kubebuilder:object:root=true
type VMMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VMMigration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMMigration{}, &VMMigrationList{})
}
