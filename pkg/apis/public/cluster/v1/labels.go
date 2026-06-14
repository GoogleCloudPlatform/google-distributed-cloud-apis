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

const (
	// User project of vanilla cluster.
	UserProjectLabel = "cluster.gdc.goog/user-project"

	// EnableNodeExternalEgressLabel is the label on cluster CR to indicate
	// whether node egress traffic to destinations outside the organization is
	// allowed for both control plane and worker nodes of a Vanilla cluster.
	//
	// If the label is not specified or set to false on the cluster CR, by
	// default node egress traffic to outside the organization is forbidden.
	//
	// If the label is set to true, egress NAT and egress network policies are
	// configured to allow node egress traffic to destinations outside the
	// organization.
	EnableNodeExternalEgressLabel = "cluster.gdc.goog/enable-node-egress-to-outside-the-org"

	// DedicatedNodeLabel is the label key for dedicated nodes.
	DedicatedNodeLabel = "cluster.gdc.goog/dedicated-node"

	// AutoScalingEnabledLabel is the label key to indicate if autoscaling is enabled.
	AutoScalingEnabledLabel = "cluster.gdc.goog/autoscaling-enabled"
)
