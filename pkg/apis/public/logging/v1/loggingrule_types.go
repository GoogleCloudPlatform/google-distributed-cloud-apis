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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Source string

const (
	Operational Source = "operational"
	Audit       Source = "audit"
)

// Defines the specification or expected state of the `LoggingRule` object.
type LoggingRuleSpec struct {
	// The log source on which to base alerts. Accepted values are `operational` and `audit`.
	// +kubebuilder:default="operational"
	// +kubebuilder:validation:Enum=operational;audit
	Source Source `json:"source,omitempty"`

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
	// The time series in which to write the record rule. It must
	// be a valid metric name.
	Record string `json:"record,omitempty"`

	// The PromQL or LogQL expression to evaluate the record rule.
	Expr string `json:"expr,omitempty"`

	// The labels to add or overwrite.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// Defines the alert rules configuration.
type AlertRule struct {
	// The alert name. Its value must be a valid label value.
	Alert string `json:"alert,omitempty"`

	// The PromQL or LogQL expression to evaluate the alert rule.
	Expr string `json:"expr,omitempty"`

	// The duration in seconds over which the specified condition must be
	// met to move the alert from the pending state to the open state.
	// +kubebuilder:default="0s"
	For string `json:"for,omitempty"`

	// The labels to add or overwrite. The required labels in this field are
	// `severity: [error, critical, warning, info, high, moderate, low]`,
	// `code: <short code for the error>`, and `resource: <component, service,
	// `errorcode: <short code for the error. code will be deprecated after migration to errorcode>`
	// `_gdch_project: <Appropriate project. For IO it will be infra-obs and for PA it will be platform-obs>`
	Labels map[string]string `json:"labels,omitempty"`

	// The annotations to add.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Defines the observed state of the `LoggingRule` object.
type LoggingRuleStatus struct {
	// A list of conditions observed in the logging alerting stack.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The name of the Loki host instance where the `LoggingRule` object is
	// currently installed.
	LokiInstance string `json:"lokiInstance,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=log,component=logging,entities="rules"
// +gdcloud:manifest:verbs=create;list;describe;delete;update
// +gdcloud:manifest:rbac="create,describe,list:loggingrule-pa-creator,loggingrule-creator,project-cortex-alertmanager-editor"
// +gdcloud:manifest:rbac="delete,describe,list,update:loggingrule-pa-editor,loggingrule-editor,project-cortex-alertmanager-editor"
// +gdcloud:manifest:rbac="describe,list:loggingrule-pa-viewer,loggingrule-viewer,project-cortex-alertmanager-viewer"
// +gdcloud:manifest:skipcodegen=true
// Defines the Schema for the Logging Rules API.
// +genclient
type LoggingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoggingRuleSpec   `json:"spec,omitempty"`
	Status LoggingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// Contains a list of logging rules.
type LoggingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingRule{}, &LoggingRuleList{})
}
