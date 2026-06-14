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

// UpgradeStrategyType defines the type of upgrade strategy.
// +kubebuilder:validation:Enum=NodePool
type UpgradeStrategyType string

const (
	// UpgradeStrategyTypeNodePool indicates the strategy applies to node pools.
	UpgradeStrategyTypeNodePool UpgradeStrategyType = "NodePool"
)

// UpgradeStrategy defines how upgrade concurrency is handled.
type UpgradeStrategy struct {
	// Type indicates the kind of upgrade strategy.
	// +kubebuilder:validation:Required
	Type UpgradeStrategyType `json:"type"`

	// NodePoolStrategy defines the concurrency strategy for node pools.
	// It is required when Type is "NodePool".
	// +optional
	NodePoolStrategy *NodePoolUpgradeStrategy `json:"nodePoolStrategy,omitempty"`
}

// NodePoolUpgradeStrategy orchestrates parallelization of entire nodepools (wave-based execution)
// and defines the concurrency strategy for the nodes within those pools.
type NodePoolUpgradeStrategy struct {
	// NodePoolUpgradePhases is a list of upgrade phases that define the wave-based execution strategy.
	NodePoolUpgradePhases []UpgradePhaseConfig `json:"nodePoolUpgradePhases,omitempty"`

	// NodeStrategy defines the node-level concurrency strategy.
	NodeStrategy NodeUpgradeStrategy `json:"nodeStrategy,omitempty"`
}

// NodeUpgradeStrategy defines how individual nodes are upgraded concurrently within a nodepool.
type NodeUpgradeStrategy struct {
	// +kubebuilder:validation:Minimum=1
	// ConcurrentNodes is the number of nodes to upgrade concurrently.
	ConcurrentNodes int `json:"concurrentNodes,omitempty"`

	// MinimumAvailableNodes is the minimum number of nodes that must be available during the upgrade.
	MinimumAvailableNodes int `json:"minimumAvailableNodes,omitempty"`
}

// UpgradePhaseConfig provides granular control over a specific phase of nodepool upgrades.
type UpgradePhaseConfig struct {
	// Maximum number of elements to upgrade concurrently. -1 to specify all remaining.
	MaxConcurrency int `json:"maxConcurrency,omitempty"`
	// Percent success for considering this group of node pools as completed.
	CompletionThreshold int `json:"completionThreshold,omitempty"`
}
