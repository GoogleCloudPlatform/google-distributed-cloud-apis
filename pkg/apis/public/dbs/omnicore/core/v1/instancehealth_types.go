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
	"sigs.k8s.io/controller-runtime/pkg/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true

// InstanceHealthStatus represents the current state of an instance's health.
type InstanceHealthStatus struct {
	EntityStatus `json:",inline"`

	// consecutiveHealthcheckFailures is the number of consecutive health check failures previously .
	// +kubebuilder:default=0
	// +required
	ConsecutiveHealthcheckFailures int `json:"consecutiveHealthcheckFailures"`

	// LastHealthcheckRunTime is the timestamp of the last healthcheck run.
	// +optional
	LastHealthcheckRunTime *metav1.Time `json:"lastHealthcheckRunTime,omitempty"`
}

type InstanceHealth interface {
	Entity
	InstanceHealthStatus() *InstanceHealthStatus
}

type InstanceHealthList interface {
	client.ObjectList
	InstanceHealthListItem() []InstanceHealth
}
