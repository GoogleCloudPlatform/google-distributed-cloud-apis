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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MaintenanceAllowedCondition     = "MaintenanceAllowed"
	MinMaintenancePolicyDuration    = 4 * time.Hour
	MaxMaintenanceExclusionDuration = 90 * 24 * time.Hour
	MaintenancePolicyResourceType   = "maintenancepolicies"
)

const (
	InsideMaintenanceWindowReason    = "InsideMaintenanceWindow"
	InsideMaintenanceExclusionReason = "InsideMaintenanceExclusion"
	OutsideMaintenanceWindowReason   = "OutsideMaintenanceWindow"
)

// WeeklyCycle represents the weekly-recurring time window for operations.
type WeeklyCycle struct {
	// Day of the week when the operations can be started.
	// One or more days may be specified.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=7
	// +kubebuilder:validation:=UniqueItems=true
	DaysOfWeek []DayOfWeek `json:"daysOfWeek"`
	// Time within the day to start the operations.
	// UTC timezone is assumed.
	StartTime TimeOfDay `json:"startTime"`
	// Duration of the time window.
	Duration metav1.Duration `json:"duration"`
}

// MaintenanceWindow represents a period of time when maintenance is allowed.
// One and only one of the fields must be present.
type MaintenanceWindow struct {
	// Right now, there is only one such field to choose from, but the sake of forward compatibility it is marked as optional.

	// Weekly cycle.
	// +optional
	WeeklyCycle *WeeklyCycle `json:"weeklyCycle,omitempty"`
}

// MaintenanceExclusion is a period of time when maintenance is forbidden even if otherwise allowed by maintenance window.
// Ranges in MaintenanceExclusion use inclusive start values and exclusive end values (half-closed intervals). In interval
// notation, this is [start, end).
type MaintenanceExclusion struct {
	// Period of time when maintenance is forbidden.
	DateTimeRange DateTimeRange `json:"dateTimeRange,omitempty"`
}

type MaintenancePolicySpec struct {
	// Maintenance window that is applied to services covered by this policy.
	MaintenanceWindow MaintenanceWindow `json:"maintenanceWindow"`

	// Named periods when maintenance is forbidden even if otherwise allowed by maintenanceWindow.
	// +optional
	MaintenanceExclusions map[string]MaintenanceExclusion `json:"maintenanceExclusions,omitempty"`
}

type MaintenancePolicyStatus struct {
	// Conditions field contain conditions for MaintenancePolicies.
	// The maintenance policy controller sets the "MaintenanceAllowed" condition.
	//
	// The following are known values for the reason field:
	// - InsideMaintenanceWindow: 	 We are currently inside the scheduled Maintenance Window and the maintenance is allowed.
	//								 The status of the condition will be True.
	//
	// - OutsideMaintenanceWindow:   We are currently inside the Maintenance Window, but the exclusion is in effect,
	//								 the maintenance is not allowed. The status of the condition will be False.
	//
	// - InsideMaintenanceExclusion: We are currently outside the Maintenance Window, the maintenance is not allowed.
	//								 The status of the condition will be False.
	//
	// +listType=map
	// +listMapKey=type
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,7,rep,name=conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=mp
// +gdcloud:manifest:relevant=false,oc=ez

type MaintenancePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MaintenancePolicySpec   `json:"spec,omitempty"`
	Status MaintenancePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type MaintenancePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MaintenancePolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MaintenancePolicy{}, &MaintenancePolicyList{})
}
