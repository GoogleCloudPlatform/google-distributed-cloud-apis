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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LogAccessLevel string

const (
	// Application Operator
	AO LogAccessLevel = "ao"
	// Platform Admin
	PA LogAccessLevel = "pa"
	// Infrastructure Operator
	IO LogAccessLevel = "io"
	// DefaultMonitoringRetentionTime is the default amount of time Cortex
	// is configured to store metrics.
	DefaultMonitoringRetentionTime = "2160h"
	// A DynamicAdditionalSink name used for exporting IO logs to an external splunk source
	IoSplunkExternalName = "ioSplunkExternalSink"
)

var (
	// splunkExternal is a dynamic sink for logs to export logs to an external Splunk instance.
	IoSplunkExternalSink = DynamicAdditionalSink{
		Name:           IoSplunkExternalName,
		LogAccessLevel: "io",
	}
)

// Defines the specification or expected state of the `ObservabilityPipeline` object.
type ObservabilityPipelineSpec struct {
	// TODO(nathanharrell): 'Enabled' kept for compatibility while transitioning
	// to API 2.0, to be removed. b/235597166

	// Specifies if the Observability pipeline stack is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// The monitoring pipeline configuration.
	Monitoring ObservabilityMonitoring `json:"monitoring,omitempty"`

	// The Alertmanager pipeline configuration.
	Alerting ObservabilityAlerting `json:"alerting,omitempty"`

	// The logging pipeline configuration.
	Logging ObservabilityLogging `json:"logging,omitempty"`

	// The audit logging pipeline configuration.
	AuditLogging ObservabilityAuditLogging `json:"auditLogging,omitempty"`

	// The security pipeline configuration.
	SecurityLogging ObservabilitySecurityLogging `json:"securityLogging,omitempty"`
}

// Defines the configuration for the monitoring dashboards.
type ObservabilityMonitoring struct {
	// Retention time for metrics in hours.
	// +kubebuilder:validation:Pattern="^[0-9]+(ns|us|µs|ms|s|m|h|d|w|y)$"
	// https://cortexmetrics.io/docs/configuration/configuration-file/#generic-placeholders
	RetentionTime *string `json:"retentionTime,omitempty"`

	// The storage size for metrics data within an organization.
	// +kubebuilder:default="20Gi"
	// +kubebuilder:validation:Pattern="^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$"
	// https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/
	// +comment:used in //vendors/k8s.io/apimachinery/pkg/api/resource to parse Quantity
	LocalStorageSize string `json:"localStorageSize,omitempty"`

	// The configuration information for the Grafana instance to create.
	Grafana Grafana `json:"grafana,omitempty"`

	// The sink for all metrics. For example, using the project namespace of
	// an Application Operator as a value means that the metrics become visible
	// to the Application Operator of the project. Route by time series using
	// the `timeseries` label.
	Sink string `json:"sink,omitempty"`
}

// Defines the rules to create alerts based on metrics data.
type ObservabilityAlerting struct {
	// The storage size for alerting data within an organization.
	// +kubebuilder:default="1Gi"
	LocalStorageSize string `json:"localStorageSize,omitempty"`

	// The name of the `ConfigMap` object that contains the Alertmanager
	// configuration file.
	// +kubebuilder:validation:Pattern="^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$"
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#dns-subdomain-names
	AlertManagerConfig string `json:"alertmanagerConfig,omitempty"`

	Volumes      []corev1.Volume      `json:"volumes,omitempty"`
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
}

// Defines the expected state of the logging stack of the Observability
// pipeline.
type ObservabilityLogging struct {
	// The retention time for operational logs in hours.
	// +kubebuilder:default="720h"
	// +kubebuilder:validation:Pattern="^[0-9]+(ns|us|µs|ms|s|m|h|d|w|y)$"
	// https://cortexmetrics.io/docs/configuration/configuration-file/#generic-placeholders
	RetentionTime string `json:"retentionTime,omitempty"`

	// The storage size for logging data within an organization.
	// +kubebuilder:default="20Gi"
	// +kubebuilder:validation:Pattern="^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$"
	// https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/
	// +comment:used in //vendors/k8s.io/apimachinery/pkg/api/resource to parse Quantity
	LocalStorageSize string `json:"localStorageSize,omitempty"`

	// The sink for all logs. For example, using the project namespace of an
	// Application Operator as a value means that the logs become visible to
	// the Application Operator of the project.
	Sink string `json:"sink,omitempty"`

	// The configuration of an optional additional sink for all operational logs.
	AdditionalSink AdditionalLogSink `json:"additionalSink,omitempty"`

	// The additional sinks for operational logs apart from any predefined sink.
	// The difference with the `additionalSink` field is that `dynamicAdditionalSinks`
	// determines most of the sink variables, such as the sink's hostname, on the
	// user's behalf. The supported outputs include `ioSplunkExternalSink`.
	// +optional
	DynamicAdditionalSinks []string `json:"dynamicAdditionalSinks,omitempty"`
}

// Defines the expected state of the audit logging stack of the Observability pipeline.
type ObservabilityAuditLogging struct {
	// The retention time for audit logs in hours.
	// +kubebuilder:default="9600h"
	// +kubebuilder:validation:Pattern="^[0-9]+(ns|us|µs|ms|s|m|h|d|w|y)$"
	// https://cortexmetrics.io/docs/configuration/configuration-file/#generic-placeholders
	RetentionTime string `json:"retentionTime,omitempty"`

	// The storage size for audit logging data within an organization.
	// +kubebuilder:default="20Gi"
	// +kubebuilder:validation:Pattern="^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$"
	// https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/
	// +comment:used in //vendors/k8s.io/apimachinery/pkg/api/resource to parse Quantity
	LocalStorageSize string `json:"localStorageSize,omitempty"`

	// The configuration of an optional additional sink for all audit logs.
	AdditionalSink AdditionalLogSink `json:"additionalSink,omitempty"`

	// The additional sinks for audit logs apart from any predefined sink. The
	// difference with the `additionalSink` field is that `dynamicAdditionalSinks`
	// determines most of the sink variables, such as the sink's hostname, on the
	// user's behalf. The supported outputs include `ioSplunkExternalSink`.
	// +optional
	DynamicAdditionalSinks []string `json:"dynamicAdditionalSinks,omitempty"`
}

// Defines the expected state of the security logging stack of the Observability pipeline.
type ObservabilitySecurityLogging struct {
	// The retention time for security logs in hours.
	// +kubebuilder:default="720h"
	// +kubebuilder:validation:Pattern="^[0-9]+(ns|us|µs|ms|s|m|h|d|w|y)$"
	// https://cortexmetrics.io/docs/configuration/configuration-file/#generic-placeholders
	RetentionTime string `json:"retentionTime,omitempty"`

	// The storage size for security logging data within an organization.
	// +kubebuilder:default="20Gi"
	// +kubebuilder:validation:Pattern="^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$"
	// https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/
	// +comment:used in //vendors/k8s.io/apimachinery/pkg/api/resource to parse Quantity
	LocalStorageSize string `json:"localStorageSize,omitempty"`

	// The configuration of an optional additional sink for all security logs.
	AdditionalSink AdditionalLogSink `json:"additionalSink,omitempty"`

	// The additional sinks for security logs apart from any predefined sink. The
	// difference with the `additionalSink` field is that `dynamicAdditionalSinks`
	// determines most of the sink variables, such as the sink's hostname, on the
	// user's behalf. The supported outputs include `ioSplunkExternalSink`.
	// +optional
	DynamicAdditionalSinks []string `json:"dynamicAdditionalSinks,omitempty"`
}

// Defines the expected state of the Grafana instance provisioned.
type Grafana struct {
	// The storage size for the dashboards within an organization.
	// +kubebuilder:default="1Gi"
	// +kubebuilder:validation:Pattern="^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$"
	// https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/
	// +comment:used in //vendors/k8s.io/apimachinery/pkg/api/resource to parse Quantity
	StorageSize string `json:"storageSize,omitempty"`
}

// Configures the additional sinks to route logs. For more information,
// see [https://cloud.google.com/anthos/private-mode/docs/1.9/how-to/export-logs](https://cloud.google.com/anthos/private-mode/docs/1.9/how-to/export-logs).
type AdditionalLogSink struct {
	// TODO(nathanharrell): ClusterSelector kept for compatibility while
	// transitioning to API 2.0, to be removed. b/235597166

	ClusterSelector     ClusterSelector      `json:"clusterSelector,omitempty"`
	FluentBitConfigMaps []string             `json:"fluentbitConfigMaps,omitempty"`
	Volumes             []corev1.Volume      `json:"volumes,omitempty"`
	VolumeMounts        []corev1.VolumeMount `json:"volumeMounts,omitempty"`
}

type DynamicAdditionalSink struct {
	// Name of the sync.
	Name string `json:"name,omitempty"`
	// Choose access level for log entries. Default is 'IO'.
	// +kubebuilder:validation:Enum=pa;io
	LogAccessLevel LogAccessLevel `json:"logAccessLevel,omitempty"`
}

// TODO(nathanharrell): ClusterSelector kept for compatibility while
// transitioning to API 2.0, to be removed. b/235597166

// Selects target clusters.
type ClusterSelector struct {
	ExcludedClusterList []string `json:"exclude,omitempty"`
}

// Defines the observed state of the `ObservabilityPipeline` object.
type ObservabilityPipelineStatus struct {
	// The installed version of the Observability pipeline stack.
	Version string `json:"version,omitempty"`

	// A list of conditions observed in the Observability pipeline stack.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=mon
// Defines the Schema for the Observability Pipeline API.
// +genclient
type ObservabilityPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObservabilityPipelineSpec   `json:"spec,omitempty"`
	Status ObservabilityPipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `ObservabilityPipeline` objects.
type ObservabilityPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObservabilityPipeline `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObservabilityPipeline{}, &ObservabilityPipelineList{})
}

func (m ObservabilityMonitoring) GetRetentionTime() string {
	if m.RetentionTime == nil {
		return DefaultMonitoringRetentionTime
	}
	return *m.RetentionTime
}
