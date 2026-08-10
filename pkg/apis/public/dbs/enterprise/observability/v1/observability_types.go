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

// +kubebuilder:object:generate=true
// ParametersSpec defines the observability related configuration.
type ParametersSpec struct {
	// CustomMetrics specifies the custom metrics configuration.
	// +kubebuilder:validation:Required
	CustomMetrics *CustomMetricsSpec `json:"customMetrics"`
}

// +kubebuilder:object:generate=true
// CustomMetricsSpec defines the specifications for custom metrics.
type CustomMetricsSpec struct {
	// ResourceLimits specifies the resource limits for running the custom queries.
	// +optional
	ResourceLimits *CustomQueryResourceLimits `json:"resourceLimits,omitempty"`

	// Definitions is a list of custom metric definitions.
	// +listType=map
	// +listMapKey=metricGroup
	// +kubebuilder:validation:Required
	Definitions []CustomMetricDefinition `json:"definitions"`
}

// +kubebuilder:object:generate=true
// CustomQueryResourceLimits defines the resource constraints for custom metric queries.
type CustomQueryResourceLimits struct {
	// WorkMemory specifies the maximum amount of memory to be used by a query operation.
	// Example: "4MB"
	// +optional
	WorkMemory string `json:"workMemory,omitempty"`

	// MaxParallelWorkers specifies the maximum number of parallel workers for a query.
	// Setting to 0 disables parallel query execution.
	// +optional
	MaxParallelWorkers int `json:"maxParallelWorkers,omitempty"`

	// StatementTimeout specifies the maximum time allowed for any statement to run.
	// Example: "2s"
	// +optional
	StatementTimeout string `json:"statementTimeout,omitempty"`
}

// +kubebuilder:object:generate=true
// CustomMetricDefinition defines a single custom metric.
type CustomMetricDefinition struct {
	// MetricGroup is a name to group related metrics.
	// Used in the final metric name: alloydb_omni_custom_<metricGroup>_<metric_name>
	// +kubebuilder:validation:Required
	MetricGroup string `json:"metricGroup"`

	// Database is the name of the database to connect to for this query.
	// +kubebuilder:validation:Required
	Database string `json:"database"`

	// Query is the SQL SELECT statement to execute.
	// +kubebuilder:validation:Required
	Query string `json:"query"`

	// Metrics is a list of metrics to extract from the query result.
	// +kubebuilder:validation:Required
	// +listType=map
	// +listMapKey=name
	Metrics []MetricColumn `json:"metrics"`
}

// +kubebuilder:object:generate=true
// MetricColumn defines how to interpret a column from the query result.
type MetricColumn struct {
	// Name is the name of the column in the SQL query result.
	// This will be used as part of the Prometheus metric name or as a label name.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Desc is a human-readable description of the metric or label.
	// +kubebuilder:validation:Required
	Desc string `json:"desc"`

	// Usage specifies how this column will be used.
	// Valid values are "label", "gauge", "counter".
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=label;gauge;counter
	Usage string `json:"usage"`
}
