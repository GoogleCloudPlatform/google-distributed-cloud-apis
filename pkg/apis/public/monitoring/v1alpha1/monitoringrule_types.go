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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Defines the specification or expected state of the `MonitoringRule` object.
type MonitoringRuleSpec struct {
	// The rule evaluation interval.
	Interval string `json:"interval,omitempty"`

	// The limit number of alerts. A value of `0` means no limit.
	// +kubebuilder:default=0
	Limit int `json:"limit,omitempty"`

	// The list of record rules.
	RecordRules []RecordRule `json:"recordRules,omitempty"`

	// The list of alert rules.
	AlertRules []AlertRule `json:"alertRules,omitempty"`
}

// Defines the record rules configuration.
type RecordRule struct {
	// The time series in which to write the record rule. It must be a valid
	// metric name.
	Record string `json:"record,omitempty"`

	// The PromQL or LogQL expression to evaluate the record rule.
	Expr string `json:"expr,omitempty"`

	// The labels to add or overwrite.
	Labels map[string]string `json:"labels,omitempty"`
}

// Defines the alert rules configuration.
type AlertRule struct {
	// The alert name. Its value must be a valid label value.
	Alert string `json:"alert,omitempty"`

	// The PromQL or LogQL expression to evaluate the alert rule.
	Expr string `json:"expr,omitempty"`

	// The duration in seconds over which the specified condition must be met
	// to move the alert from the pending state to the open state.
	// +kubebuilder:default="0s"
	For string `json:"for,omitempty"`

	// The labels to add or overwrite. The required labels in this field are
	// `severity: [error, critical, warning, info, test]`,
	// `code: <short code for the error>`, and
	// `resource: <component, service, or hardware related to the alert>`.
	// Any additional labels are optional.
	Labels map[string]string `json:"labels,omitempty"`

	// The annotations to add.
	Annotations map[string]string `json:"annotations,omitempty"`

	// GroupbyLabels defines the list of logical/metadata labels used to aggregate alerts into incidents.
	GroupbyLabels []string `json:"groupby_labels,omitempty"`
}

// Defines the observed state of the `MonitoringRule` object.
type MonitoringRuleStatus struct {
	// Defines the observed state of the `MonitoringRule` object.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Defines the Schema for the Monitoring Rules API.
// +gdcloud:manifest:relevant=false,oc=mon
// +genclient
type MonitoringRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringRuleSpec   `json:"spec,omitempty"`
	Status MonitoringRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of monitoring rules.
type MonitoringRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringRule{}, &MonitoringRuleList{})
}
