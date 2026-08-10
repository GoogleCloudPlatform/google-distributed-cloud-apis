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

	semver "github.com/Masterminds/semver/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TimeWindow defines a chunk of time.
type TimeWindow struct {
	// Start indicates the start of the window.
	Start metav1.Time `json:"start,omitempty"`
	// End indicates the end of the window.
	End metav1.Time `json:"end,omitempty"`
}

// IsAfterEnd returns true if now is on or after the window End time (when End is non-zero).
func (t TimeWindow) IsAfterEnd(now time.Time) bool {
	return !t.End.IsZero() && !now.Before(t.End.Time)
}

// +kubebuilder:validation:Type=string

// Wrapper/alias type for `semver.Version` (Semantic Version).
// Extended to support KRM APIs.
type SemanticVersion semver.Version

func (v SemanticVersion) String() string {
	return semver.Version(v).String()
}

func (v *SemanticVersion) UnmarshalJSON(b []byte) error {
	return (*semver.Version)(v).UnmarshalJSON(b)
}

func (v SemanticVersion) MarshalJSON() ([]byte, error) {
	return semver.Version(v).MarshalJSON()
}

func (v *SemanticVersion) Equal(o *SemanticVersion) bool {
	return (*semver.Version)(v).Equal((*semver.Version)(o))
}

// +kubebuilder:validation:Enum=Preflight;Postflight;PreRollback;

// UpgradeCheckType defines when an upgrade check runs.
//
// "Preflight" will run before upgrade.
// "Postflight" will run after upgrade.
// "PreRollback" will run before rollback.
type UpgradeCheckType string

// LINT.IfChange
const (
	// UpgradeCheckTypePreflight indicates the check job will run before upgrade.
	UpgradeCheckTypePreflight UpgradeCheckType = "Preflight"
	// UpgradeCheckTypePostflight indicates the check job will run after upgrade.
	UpgradeCheckTypePostflight UpgradeCheckType = "Postflight"
	// UpgradeCheckTypePreRollback indicates the check job will be run before rollback.
	UpgradeCheckTypePreRollback UpgradeCheckType = "PreRollback"
)

// LINT.ThenChange(//build/rules/upgrade_check/providers.bzl)

// NodeUpgradeDisruption defines the disruption level for node upgrades.
type NodeUpgradeDisruption string

const (
	// RequiresWorkloadDrain forces a workload drain.
	RequiresWorkloadDrain NodeUpgradeDisruption = "RequiresWorkloadDrain"
	// NoDisruption indicates no workload drain is required.
	NoDisruption NodeUpgradeDisruption = "NoDisruption"
	// RequiresNodeReboot indicates a node reboot is required.
	RequiresNodeReboot NodeUpgradeDisruption = "RequiresNodeReboot"
)
