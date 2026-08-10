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

const (
	// EnableDefaultExternalEgressAllowKey is the key for enabling default egress
	// allow for a Project to destinations outside the organization.
	// If the label is not specified on a Project:
	// - the egress traffic from the project is not allowed to destinations
	//   outside the organization, and;
	// - egress ProjectNetworkPolicies in the project will show a validation error
	//   in the status field and not be enforced in the dataplane.
	// Otherwise:
	// - the egress traffic is allowed to all destinations by default, and;
	// - egress ProjectNetworkPolicies can be defined to restrict the egress traffic.
	// The label value is ignored so specifying the label key is sufficient to
	// enable default egress allow.
	EnableDefaultExternalEgressAllowKey = "networking.gdc.goog/enable-default-egress-allow-to-outside-the-org"

	// DelegateEnableDenyLoggingKey is the annotation key to delegate network
	// logging for allowed flows to a network policy. See:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/network-policy-logging#networklogging_spec
	DelegateEnableDenyLoggingKey = "policy.network.gke.io/enable-deny-logging"

	// DelegateEnableLoggingKey is the annotation key to delegate network
	// logging for denied flows to a namespace. See:
	// ttps://cloud.google.com/kubernetes-engine/docs/how-to/network-policy-logging#networklogging_spec
	DelegateEnableLoggingKey = "policy.network.gke.io/enable-logging"

	// IngressKey is used to uniquely identify istio ingress services on the same cluster.
	// This unique identification is used to map the istio ingress services to IP addresses.
	IngressKey = "networking.gdc.goog/ingress"
)
