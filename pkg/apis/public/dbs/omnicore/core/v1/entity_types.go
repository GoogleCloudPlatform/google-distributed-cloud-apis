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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type (
	ConditionType   string
	ConditionReason string
)

//+kubebuilder:object:generate=true

type EntityStatus struct {
	// Internal: The generation observed by the controller.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Internal: Whether the resource was reconciled by the controller.
	Reconciled bool `json:"reconciled,omitempty"`

	// CriticalIncidents is a flat list of all active Critical Incidents.
	CriticalIncidents []CriticalIncident `json:"criticalIncidents,omitempty"`

	// Conditions represents the latest available observations of the
	// Entity's current state.
	// +listType=map
	// +listMapKey=type
	// +patchmergekey=type
	// +patchstrategy=merge
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,2,rep,name=conditions"`
}

type Entity interface {
	client.Object
	EntityStatus() *EntityStatus
}
