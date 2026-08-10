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
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// MaintenanceWindow specifies a recurring time window for applying for patch and minor version upgrades.
type MaintenanceWindow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MaintenanceWindowSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowList represents a collection of MaintenanceWindows.
type MaintenanceWindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MaintenanceWindow `json:"items"`
}

// MaintenanceWindowSpec provides the specification (i.e., desired state) of
// a MaintenanceWindow.
type MaintenanceWindowSpec struct {
	// UpgradeType indicates the type of the MaintenanceWindow, which is one of
	// "MinorUpgrade" and "PatchUpgrade".
	UpgradeType UpgradeType `json:"upgradeType"`

	// Recurrence encodes a RRULE string to indicate how the window recurs.
	// https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html
	// Ex. TimeWindow.StartTime = 2022-04-18T02:00:00Z
	//     TimeWindow.EndTime =  2022-04-18T06:00:00Z
	//     Recurrence = FREQ=WEEKLY;BYDAY=MO,WE
	// These parameters would create a schedule that starts on 04/18/2022 from 2 a.m to 6 a.m.
	// Every Monday and Wednesday after, the same 2.am-6.am pattern would recur.
	Recurrence string `json:"recurrence"`

	// TimeWindow contains the start and end times for the MaintenanceWindow.
	TimeWindow TimeWindow `json:"timeWindow"`

	// Exclusions is an array of TimeWindows that were skipped. Exclusions will only contain times specific up to the hour.
	Exclusions []TimeWindow `json:"exclusions,omitempty"`
}

// +kubebuilder:validation:Enum=MinorUpgrade;PatchUpgrade

// UpgradeType specifies the type of an upgrade.
// Only one of the following upgrade types may be specified.
type UpgradeType string

const (
	// UpgradeTypeMinor indicates that the UpgradeType is for a minor upgrade.
	UpgradeTypeMinor UpgradeType = "MinorUpgrade"

	// UpgradeTypePatch indicates that the UpgradeType is for a patch upgrade.
	UpgradeTypePatch UpgradeType = "PatchUpgrade"
)

func init() {
	SchemeBuilder.Register(&MaintenanceWindow{}, &MaintenanceWindowList{})
}
