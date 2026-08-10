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
)

// NOTE: the getter names in the interfaces defined in this file do not conform
// to go/go-style/decisions#getters, to avoid naming conflicts between getters
// and exported fields.

// Lets you work with the .status field of a global mux resource.
// Attempting to retrieve a non-existent field will be a no-op and return a zero
// value.
type MuxStatusInterface interface {
	GetConditions() []metav1.Condition
	SetConditions(conditions []metav1.Condition)
	GetRollout() RolloutStatus
	SetRollout(rolloutStatus RolloutStatus)
	GetZones() ListMap[string, ZoneStatusInterface]
	SetZones(zones ListMap[string, ZoneStatusInterface]) error
}

// Lets you work with any API field that represents a map as a list of
// subobjects whose keys are embedded as one or more scalar fields within the
// subobjects. Subobjects in the list must have unique keys.
// Attempting to retrieve a non-existent item will be a no-op and return a zero
// value.
// Attempting to set a value will result in any existing value that has an the
// same key to be replaced.
// See the following links for the corresponding API convention:
// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#lists-of-named-subobjects-preferred-over-maps
type ListMap[K comparable, V any] interface {
	Keys() []K
	Values() []V
	Get(k K) V
	Set(v V) error
	Clear()
}

// Lets you work with an item in the .status.zones field of a global mux
// resource.
// Attempting to retrieve a non-existent field will be a no-op and return a zero
// value.
type ZoneStatusInterface interface {
	GetName() string
	SetName(name string)
	GetRolloutStatus() ZoneRolloutStatus
	SetRolloutStatus(rolloutStatus ZoneRolloutStatus)
	GetReplicaStatus() any
	SetReplicaStatus(replicaStatus any) error
}

// +kubebuilder:object:generate=true

// Provides the status of rolling out a global mux resource to a zone.
type ZoneRolloutStatus struct {
	// The observations of the current rollout.
	// Known condition types: Replicated, Synced.
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The .metadata.generation of the replica of the global mux resource in the
	// zonal API server. If the .status.zones[*].replicaStatus contains any
	// condition with an .observedGeneration less than this field, the condition
	// is out of date.
	ReplicaGeneration int64 `json:"replicaGeneration,omitempty"`
}

// +kubebuilder:object:generate=true

// Provides a minimum shape to adhere to the zone status specification for
// global mux resources.
// This duck type is intended to allow implementations of zone statuses embedded
// in global mux resource statuses and verifications that resource statuses meet
// the expectations.
// This is not a real zone status.
type ZoneStatus struct {
	// The name of the zone where the replica this status represents is in.
	Name string `json:"name"`

	// The status of rolling out the replica to the zone.
	RolloutStatus ZoneRolloutStatus `json:"rolloutStatus,omitempty"`
}

func (s *ZoneStatus) GetName() string {
	return s.Name
}

func (s *ZoneStatus) SetName(name string) {
	s.Name = name
}

func (s *ZoneStatus) GetRolloutStatus() ZoneRolloutStatus {
	return s.RolloutStatus
}

func (s *ZoneStatus) SetRolloutStatus(rolloutStatus ZoneRolloutStatus) {
	s.RolloutStatus = rolloutStatus
}

// Provides the current rollout strategy being used to roll out a global mux
// resource.
type RolloutStatus struct {
	// The current rollout strategy.
	Strategy RolloutStrategy `json:"strategy,omitempty"`
}

// +kubebuilder:object:generate=true

// Provides a minimum shape to adhere to the global mux resource status
// specification.
// This duck type is intended to allow implementations of global mux resource
// statuses and verifications that resource statuses meet the expectations.
// This is not a real status.
type MuxStatus struct {
	// The observations of the overall state of the resource.
	// Known condition types: Ready.
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The current strategy used to roll out the resource to each zone.
	Rollout RolloutStatus `json:"rollout,omitempty"`
}

func (s *MuxStatus) GetConditions() []metav1.Condition {
	return s.Conditions
}

func (s *MuxStatus) SetConditions(conditions []metav1.Condition) {
	s.Conditions = conditions
}

func (s *MuxStatus) GetRollout() RolloutStatus {
	return s.Rollout
}

func (s *MuxStatus) SetRollout(rolloutStatus RolloutStatus) {
	s.Rollout = rolloutStatus
}
