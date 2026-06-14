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
	// IntraProjectPolicyCreationTimestampKey is the key for the annotation which
	// holds the creation time of the generated 'allow-intra-project-traffic'
	// ProjectNetworkPolicy.
	IntraProjectPolicyCreationTimestampKey = "networking.gdc.goog/intra-project-policy-creation-time"

	// FirstPartyServicePortsKey specify the ports of the backend pods that 1P services expose.
	// This annotation is added on ShadowProject resources.
	// The annotation is consumed by ProjectNetworkPolicy controller to realize
	// the backend ports that the policy applies to.
	// The value of the annotation must be a valid JSON string in the format
	// specified by list of ProjectNetworkPolicyPorts.
	// Examples:
	// - `[{"protocol":"TCP","port":80}]`
	// - `[{"protocol":"TCP","port":80},{"protocol":"UDP","port":53}]`
	FirstPartyServicePortsKey = "networking.gdc.goog/1p-service-ports"

	// ProjectDisableEgressKey annotation is set by DBS to inform UNET to disable egress IPs belonging to a
	// ServiceIsolationEnvironment.
	// "networking.gdc.goog/disable-egress-ip" = "true" disables IPs
	// Else if the annotation is missing or any other value, the IPs are not disabled.
	ProjectDisableEgressKey = "networking.gdc.goog/disable-egress-ip"

	// LoadBalancerTypeAnnotationKey is the annotation key on service of load balancer type.
	LoadBalancerTypeAnnotationKey = "networking.gke.io/load-balancer-type"
	ILBAnnotationValue            = "internal"
)
