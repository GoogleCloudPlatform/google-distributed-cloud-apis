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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Defines the selectors to identify logs to collect. The different selectors (MatchClusters, MatchPodNamePrefixes, and MatchContainerNamePrefixes) are combined with a logical `AND` relationship. If a selector is not specified, its default behavior is to match all resources (e.g., all clusters, all pods, or all containers).
type SIEMProjectForwarderSelectors struct {
	// The resource names of the clusters to collect logs from. These refer to
	// `Cluster` resources in the `clusters.gdc.goog` API group. The default configuration
	// is to collect logs from all clusters. The relationship between different
	// clusters is an `OR` relationship.
	// +optional
	// +listType=set
	MatchClusters []string `json:"matchClusters,omitempty"`

	// The pod name prefixes to collect logs from. The Observability platform
	// scrapes all pods with names that start with the specified prefixes. The values
	// must contain `[a-z0-9-]` characters only. The relationship between different
	// list elements is an `OR` relationship. The default configuration is to collect
	// logs from all pods if not specified.
	// +optional
	// +listType=set
	MatchPodNamePrefixes []string `json:"matchPodNamePrefixes,omitempty"`

	// The container name prefixes to collect logs from. The Observability platform
	// scrapes all containers with names that start with the specified prefixes. The values
	// must contain `[a-z0-9-]` characters only. The relationship between different
	// list elements is an `OR` relationship. The default configuration is to collect
	// logs from all containers if not specified.
	// +optional
	// +listType=set
	MatchContainerNamePrefixes []string `json:"matchContainerNamePrefixes,omitempty"`
}

// Defines the specification or expected state of the SIEMProjectForwarder object.
type SIEMProjectForwarderSpec struct {
	// The matching pattern that identifies pods or containers to collect logs from.
	// +optional
	Selector *SIEMProjectForwarderSelectors `json:"selector,omitempty"`

	// The predefined parser for log entries.
	// +kubebuilder:validation:Enum=klog_text;klog_json;klogr;gdch_json;json
	// +optional
	Parser *OperationalLogParser `json:"parser,omitempty"`

	// A service name to apply as a label. For user workloads,
	// you can consider this field for a workload name.
	// +optional
	ServiceName *string `json:"serviceName,omitempty"`

	// The additional static fields to apply to log entries.
	// This field is a mapping of key-value pairs, where the field name is the key and the field value is the value.
	// +optional
	AdditionalFields map[string]string `json:"additionalFields,omitempty"`

	// Destinations specifies the SIEM output configurations.
	Destinations SIEMProjectDestinations `json:"destinations"`
}

// Defines the possible SIEM destination configurations for projects.
type SIEMProjectDestinations struct {
	// Defines a list of Elastic destination configurations.
	ElasticOutputs []SIEMProjectElasticOutput `json:"elasticOutputs"`
}

// Defines the configuration for an Elastic SIEM output.
type SIEMProjectElasticOutput struct {
	// The host name of the target Elastic service.
	Host string `json:"host"`

	// TCP port of the Elastic instance.
	// +kubebuilder:default=9200
	// +optional
	Port *int `json:"port,omitempty"`

	// The reference to a basic-auth Secret containing the HTTP credentials (username and password)
	// for the Elastic instance.
	HTTPCredentialSecretRef corev1.LocalObjectReference `json:"httpCredentialSecretRef"`

	// +kubebuilder:default=true
	// +optional
	TLS *bool `json:"tls,omitempty"`

	// +kubebuilder:default=10
	// +optional
	NetConnectTimeoutSeconds *int `json:"netConnectTimeoutSeconds,omitempty"`

	// The index for the Elastic instance.
	// +kubebuilder:default="operational"
	// +optional
	Index *string `json:"index,omitempty"`

	// The version of the Elasticsearch instance. Only versions greater than or equal to 7 are supported.
	// The version can be specified as a major version (e.g., "7", "8") or a full semantic version (e.g., "7.17.3", "8.5.1").
	// +required
	ElasticsearchVersion string `json:"elasticsearchVersion"`
}

// Defines the observed state of the SIEMProjectForwarder object.
type SIEMProjectForwarderStatus struct {
	// A list of conditions observed in the logging stack.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=log
// Defines the Schema for the siemprojectforwarders API.
// +genclient
// +gdcloud:manifest:relevant=true,oc=log,component=logging,entities="siem-project-forwarders"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create:siemprojectforwarder-creator"
// +gdcloud:manifest:rbac="delete,update:siemprojectforwarder-editor"
// +gdcloud:manifest:rbac="describe,list:siemprojectforwarder-creator,siemprojectforwarder-editor,siemprojectforwarder-viewer"
type SIEMProjectForwarder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SIEMProjectForwarderSpec   `json:"spec,omitempty"`
	Status            SIEMProjectForwarderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// Contains a list of SIEMProjectForwarder objects.
type SIEMProjectForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SIEMProjectForwarder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SIEMProjectForwarder{}, &SIEMProjectForwarderList{})
}
