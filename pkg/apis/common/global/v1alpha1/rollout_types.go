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

// +kubebuilder:validation:Enum=Immediate;Custom

// Enumerates the rollout strategy for a global MUX resource.
type RolloutStrategyType string

const (
	// Rolls out a global MUX resource to all zones immediately at the same time.
	ImmediateRolloutStrategyType RolloutStrategyType = "Immediate"

	// Rolls out a global MUX resource based on a custom rollout controller.
	CustomRolloutStrategyType RolloutStrategyType = "Custom"
)

// Describes how to roll out a global MUX resource to each zone.
type RolloutStrategy struct {
	// The type of the rollout strategy.
	Type RolloutStrategyType `json:"type"`
}
