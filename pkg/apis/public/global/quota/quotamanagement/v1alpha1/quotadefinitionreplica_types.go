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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// Represents the replica of the global QuotaDefinition resource in each zone.
type QuotaDefinitionReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuotaDefinitionSpec          `json:"spec"`
	Status QuotaDefinitionReplicaStatus `json:"status,omitempty"`
}

type QuotaDefinitionSpec struct {
	// The quota metrics that the quotas are defined against.
	Metrics []Metric `json:"metrics,omitempty"`

	// The quotas defined for the service.
	Quotas []Quota `json:"quotas,omitempty"`

	// The rules that bind API methods to quota metrics. Binding an API method to a metric causes that metric's configured quota behaviors to apply to the method call. Used for API rate quota only.
	// +optional
	APIRules []APIRule `json:"apiRules,omitempty"`

	// Whether to enable rate quota in advisory mode for the service. If enabled,
	// then no rate limit enforcement will happen even if configured rate limits
	// are exceeded. Telemetry will still be updated to show exceeded limits.
	// +optional
	EnableAdvisoryMode bool `json:"enableAdvisoryMode"`
}

type APIRule struct {
	// The selector selects the APIs to which this rule applies.
	Selector APISelector `json:"selector"`

	// A list of objects consist of Quota metrics to update when the selected methods are called, and the associated cost applied to each metric.
	// The key of the map is the quota metric name, and the values are the amount increased for the metric against which the quotas are defined
	// +listType=map
	// +listMapKey=name
	MetricCosts []MetricCost `json:"metricCosts"`
}

type APISelector struct {
	// A list of string patterns that match one or more API methods. The pattern can be either a full path of a method or a "*" indicating a wildcard.
	// For example, the field can be ["*"] that selects all the methods in the service; or ["google.api.foo.v1alpha.Foo.setGpus", "google.api.foo.v1alpha.Foo.deleteGpus"] that selects two methods of the service Foo.
	MatchMethods []string `json:"matchMethods"`
}

type MetricCost struct {
	// The name of the metric for cost charging.
	Name string `json:"name"`

	// The cost applied to the metric.
	Cost int64 `json:"cost"`
}

type QuotaDefinitionReplicaStatus struct {
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a list of QuotaDefinitionReplicas
type QuotaDefinitionReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuotaDefinitionReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(&QuotaDefinitionReplica{}, &QuotaDefinitionReplicaList{})
}
