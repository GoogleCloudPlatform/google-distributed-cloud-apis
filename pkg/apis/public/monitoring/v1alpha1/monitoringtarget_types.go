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

// Defines the specification or expected state of the `MonitoringTarget` object.
type MonitoringTargetSpec struct {
	// The matching pattern that identifies pods for this job. To establish a
	// relationship between different selectors, use `AND`.
	Selector MonitoringTargetSelectors `json:"selector,omitempty"`

	// The endpoint exposed for this job. The endpoint uses the style of
	// Prometheus.
	PodMetricsEndpoints MonitoringTargetPodMetricsEndpoints `json:"podMetricsEndpoints,omitempty"`
}

// Provides selectors that determine which pods to monitor.
type MonitoringTargetSelectors struct {
	// The clusters to consider for this job. The default configuration is to
	// consider all the clusters applicable to the project. The relationship
	// between different clusters is an `OR` relationship. For example, the
	// value `["admin", "system"]` indicates to consider the admin cluster `OR`
	// the system cluster.
	MatchClusters []string `json:"matchClusters,omitempty"`

	// The pod labels to consider for this job. The default configuration is to
	// consider no filter based on labels. The relationship between different
	// pairs is an `AND` relationship, so all pairs are considered.
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// The annotations to consider for this job. The default configuration is
	// to consider no filter based on annotations. The relationship between
	// different pairs is an `AND` relationship, so all pairs are considered.
	MatchAnnotations map[string]string `json:"matchAnnotations,omitempty"`
}

// Configures the metric endpoints for scraped pods.
type MonitoringTargetPodMetricsEndpoints struct {
	// The port from which metrics are scraped.
	Port MonitoringTargetPodMetricsPort `json:"port,omitempty"`

	// The path from which metrics are scraped.
	Path MonitoringTargetPodMetricsPath `json:"path,omitempty"`

	// The scheme to use when scraping metrics.
	Scheme MonitoringTargetPodMetricsScheme `json:"scheme,omitempty"`

	// The query parameters to use when scraping metrics from the `path`.
	Params map[string][]string `json:"params,omitempty"`

	// The frequency for Prometheus to scrape the metric endpoints defined in
	// the `podMetricsEndpoints` field.
	// +kubebuilder:default="60s"
	ScrapeInterval string `json:"scrapeInterval,omitempty"`

	// The time for Prometheus to wait for the response from the metric
	// endpoints defined in the `podMetricsEndpoints` field.
	// +kubebuilder:default="10s"
	ScrapeTimeout string `json:"scrapeTimeout,omitempty"`

	// The filter for either including (`allowlist`) or excluding (`denylist`)
	// metrics based on labels.
	MetricsRelabelings []MonitoringTargetMetricsRelabeling `json:"metricsRelabelings,omitempty"`
}

// Determines the port to use for scraping metrics from pods.
type MonitoringTargetPodMetricsPort struct {
	// The port to collect metrics from. If annotations are provided, they take
	// priority over this field.
	// +kubebuilder:default=80
	Value int `json:"value,omitempty"`

	// The port to collect metrics from using annotations.
	Annotation string `json:"annotation,omitempty"`
}

// Determines the path to use for scraping metrics from pods.
type MonitoringTargetPodMetricsPath struct {
	// The path to collect metrics from. If annotations are provided, they take
	// priority over this field.
	// +kubebuilder:default="/metrics"
	Value string `json:"value,omitempty"`

	// The path to collect metrics from using annotations.
	Annotation string `json:"annotation,omitempty"`
}

// Determines the scheme to use for scraping metrics from pods.
type MonitoringTargetPodMetricsScheme struct {
	// The scheme to use when collecting metrics. If annotations are provided,
	// they take priority over this field.
	// +kubebuilder:default="http"
	// +kubebuilder:validation:Enum=http;https
	Value string `json:"value,omitempty"`

	// The scheme to use when collecting metrics using annotations.
	Annotation string `json:"annotation,omitempty"`
}

// Defines a filter for keeping or discarding metrics based on labels.
type MonitoringTargetMetricsRelabeling struct {
	// The selected values from existing labels. The content is concatenated
	// using the `separator` and matched against the `regex` expression for the
	// `replace`, `keep`, and `drop` actions.
	SourceLabels []string `json:"sourceLabels,omitempty"`

	// The separator value placed between concatenated source label values.
	Separator string `json:"separator,omitempty"`

	// The regular expression to match the extracted value.
	Regex string `json:"regex,omitempty"`

	// The action to perform when the `regex` expression matches the extracted
	// value.
	Action string `json:"action,omitempty"`

	// The label to which to write the resulting value in a `replace` action.
	// This field is mandatory for `replace` actions. Capture groups of regular
	// expressions are available.
	TargetLabel string `json:"targetLabel,omitempty"`

	// The replacement value to use if the regular expression matches the
	// extracted value in a `replace` action. Capture groups of regular
	// expressions are available.
	Replacement string `json:"replacement,omitempty"`
}

// Defines the observed state of the `MonitoringTarget` object.
type MonitoringTargetStatus struct {
	// A list of conditions observed in the monitoring stack.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:validation:XValidation:rule="self.metadata.name.size() <= 63",message="MonitoringTarget name must be 63 characters or less"
// +kubebuilder:validation:XValidation:rule="!self.metadata.name.contains('.')",message="MonitoringTarget name must not contain dots"

// Defines the Schema for the monitoring targets API.
// +gdcloud:manifest:relevant=false,oc=mon
// +genclient
type MonitoringTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringTargetSpec   `json:"spec,omitempty"`
	Status MonitoringTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of monitoring targets.
type MonitoringTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringTarget{}, &MonitoringTargetList{})
}
