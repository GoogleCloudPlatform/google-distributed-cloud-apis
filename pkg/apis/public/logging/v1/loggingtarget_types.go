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
	logmonv1alpha1 "sigs.k8s.io/cluster-operators/logmon-operator/api/v1alpha1"
)

type LogAccessLevel string

const (
	// Application Operator
	AO LogAccessLevel = "ao"
	// Platform Admin
	PA LogAccessLevel = "pa"
	// Infrastructure Operator
	IO LogAccessLevel = "io"
)

// Defines the specification or expected state of the `LoggingTarget` object.
type LoggingTargetSpec struct {
	//  The matching pattern that identifies pods or containers to collect
	// logs from. The relationship between different selectors is an
	// `AND` relationship, so all selectors are considered.
	Selector LoggingTargetSelectors `json:"selector,omitempty"`
	// The access level for log entries. The default value is `AO` for
	// Application Operator.
	// +kubebuilder:validation:Enum=ao;pa;io
	// +optional
	LogAccessLevel LogAccessLevel `json:"logAccessLevel,omitempty"`
	// The predefined parser for log entries.
	// +kubebuilder:validation:Enum=klog_text;klog_json;klogr;gdch_json;json
	// +optional
	Parser logmonv1alpha1.OperationalLogParser `json:"parser,omitempty"`
	// A service name to apply as a label. For user workloads,
	// you can consider this field for a workload name.
	ServiceName string `json:"serviceName,omitempty"`
	// The additional static fields to apply to log entries. This field is
	// a mapping of key-value pairs, where the field name is the key and
	// the field value is the value.
	// +optional
	AdditionalFields map[string]string `json:"additionalFields,omitempty"`
}

// Provides selectors that determine which pods or containers to collect logs from.
type LoggingTargetSelectors struct {
	// The clusters to collect logs from. The default configuration is to
	// collect logs from all clusters. The relationship between different
	// clusters is an `OR` relationship. For example, the value
	// `["admin", "system"]` indicates to consider the admin cluster
	// `OR` the system cluster.
	// +optional
	MatchClusters []string `json:"matchClusters,omitempty"`
	// The pod name prefixes to collect logs from. The Observability platform
	// scrapes all pods with names that start with the specified prefixes. The values
	// must contain `[a-z0-9-]` characters only. The relationship between different
	// list elements is an `OR` relationship.
	// +optional
	MatchPodNames []string `json:"matchPodNames,omitempty"`
	// The container name prefixes to collect logs from. The Observability platform
	// scrapes all containers with names that start with the specified prefixes. The values
	// must contain `[a-z0-9-]` characters only. The relationship between different
	// list elements is an `OR` relationship.
	// +optional
	MatchContainerNames []string `json:"matchContainerNames,omitempty"`
}

// Defines the observed state of the `LoggingTarget` object.
type LoggingTargetStatus struct {
	// A list of conditions observed in the logging stack.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=log,component=logging,entities="targets"
// +gdcloud:manifest:verbs=create;list;describe;delete;update
// +gdcloud:manifest:rbac="create,describe,list:loggingtarget-pa-creator,loggingtarget-creator"
// +gdcloud:manifest:rbac="delete,describe,list,update:loggingtarget-pa-editor,loggingtarget-editor"
// +gdcloud:manifest:rbac="describe,list:loggingtarget-pa-viewer,loggingtarget-viewer"
// +gdcloud:manifest:skipcodegen=true
// Defines the Schema for the operational logging targets API.
// +genclient
type LoggingTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingTargetSpec   `json:"spec,omitempty"`
	Status            LoggingTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// Contains a list of logging targets.
type LoggingTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingTarget{}, &LoggingTargetList{})
}
