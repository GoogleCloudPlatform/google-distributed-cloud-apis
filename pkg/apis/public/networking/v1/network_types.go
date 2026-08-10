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
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Represents a filter that selects a set of network endpoints based on
// the filter conditions specified.
//
// +kubebuilder:validation:MinProperties=1
type NetworkEndpointFilter struct {
	// A filter that selects flow events that match the IP address or IP address range. Each of the IP addresses can be
	// specified as an exact match, like `1.1.1.1` or
	// `1200:0000:AB00:1234:0000:2552:7777:1313`, or as a CIDR range like
	// `1.1.1.0/24` or `1200:0000:AB00:1234:0000:2552:7777:1313/120`.
	// If not specified, any IP address is matched.
	//
	// +optional
	IPBlocks []string `json:"ipBlocks,omitempty"`

	// A filter that selects flow events that match the label selector. Selectors
	// support the full Kubernetes label selector syntax.
	//
	// +optional
	Labels []metav1.LabelSelector `json:"labels,omitempty"`

	// A list of namespaces and pods used to match flows.
	//
	// +optional
	NamespacePodSelectors []NamespacePodSelector `json:"namespacePodSelectors,omitempty"`

	// A filter that selects flows by their L4 ports. If this field is not provided, this matches
	// all port numbers.
	// An example value for a single port is `80`.
	// If present, only traffic on the specified protocol and port is matched.
	//
	// +optional
	Ports []intstr.IntOrString `json:"ports,omitempty"`
}

// Represents the information used to locate pods inside of the specified namespace.
// Specify a value for `namespace`, `pod`, or `namespace` and `pod`.
//
// +kubebuilder:validation:MinProperties=1
type NamespacePodSelector struct {
	// The flow events that match the namespace name.
	// For example, `kube-system`.
	//
	// +optional
	Namespace *string `json:"namespace,omitempty"`

	// The flow events that match the given pod name prefix.
	// For example, `xwing`, `coredns-`.
	//
	// +optional
	Pod *string `json:"pod,omitempty"`
}
