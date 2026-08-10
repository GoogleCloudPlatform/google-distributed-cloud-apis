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
)

type LogType string

// Defines the token configuration for the SIEM export.
type Token struct {
	// The name of the token.
	Name string `json:"name"`
	// The field of the token.
	Field string `json:"field"`
}

// Defines the Splunk output configuration.
type SplunkOutput struct {
	// The host name of the target Splunk service.
	Host string `json:"host"`

	// The authentication token for the HTTP Event Collector interface.
	Token Token `json:"token"`

	// The Transport Layer Security (TLS) protocol.
	// For more information, see https://docs.fluentbit.io/manual/administration/transport-security.
	// +kubebuilder:default="On"
	// +kubebuilder:validation:Enum=On;Off
	// +optional
	TLS *string `json:"tls,omitempty"`

	// The maximum time in seconds to wait for a TCP connection to be established.
	// This value includes the TLS handshake time.
	// +kubebuilder:default=10
	// +optional
	NetConnectTimeout *int `json:"netConnectTimeout,omitempty"`
}

// Defines the HTTP credential for the Elastic SIEM.
type HTTPCredential struct {
	// HTTP user for the Elastic instance
	HTTPUser string `json:"http_user"`

	// HTTP password for the Elastic instance
	HTTPPasswd string `json:"http_passwd"`
}

// Defines the Elastic output configuration.
type ElasticOutput struct {
	// The host name of the target Elastic service.
	Host string `json:"host"`

	// The HTTP credential for the Elastic instance.
	HTTPCredential HTTPCredential `json:"http_credential"`

	// TCP port of the Elastic instance.
	// +kubebuilder:default=9200
	// +optional
	Port *int `json:"port,omitempty"`

	// +kubebuilder:default="On"
	// +kubebuilder:validation:Enum=On;Off
	// +optional
	TLS *string `json:"tls,omitempty"`

	// +kubebuilder:default=10
	// +optional
	NetConnectTimeout *int `json:"netConnectTimeout,omitempty"`

	// The index for the Elastic instance. If not specified, the log type (audit or operational) will be used.
	// +optional
	Index *string `json:"index,omitempty"`
}

// SIEMDestinations defines the possible SIEM destination configurations.
// Users must specify outputs under `splunk` or `elastic` fields.
type SIEMDestinations struct {
	// Splunk defines a list of Splunk destination configurations.
	// +optional
	SplunkOutputs []SplunkOutput `json:"splunkOutputs,omitempty"`

	// Elastic defines a list of Elastic destination configurations.
	// +optional
	ElasticOutputs []ElasticOutput `json:"elasticOutputs,omitempty"`
}

// Defines the specification or expected state of the `SIEMOrgForwarder` resource.
type SIEMOrgForwarderSpec struct {
	// The type of logs to export to a SIEM destination. Accepted values are `operational` and `audit`.
	// +kubebuilder:validation:Enum=operational;audit
	Source LogType `json:"source"`

	// Destinations specifies the SIEM output configurations.
	// Current available destinations are Splunk and Elastic
	// +optional
	Destinations *SIEMDestinations `json:"destinations,omitempty"`

	// TODO (b/427090537): This field is to support the backward compatibility
	// To be removed in 1.17
	// +optional
	SplunkOutputs []SplunkOutput `json:"splunkOutputs,omitempty"`
}

// Defines the observed state of the `SIEMOrgForwarder` resource.
type SIEMOrgForwarderStatus struct {
	// The observed state of the `SIEMOrgForwarder` resource.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=log,component=logging,entities="siem-org-forwarders"
// +gdcloud:manifest:verbs=create;list;describe;delete;update
// +gdcloud:manifest:rbac="create,describe,list:siemexport-org-creator"
// +gdcloud:manifest:rbac="delete,describe,list,update:siemexport-org-editor"
// +gdcloud:manifest:rbac="describe,list:siemexport-org-viewer"
// +gdcloud:manifest:skipcodegen=true
// Defines the Schema for the `SIEMOrgForwarder` API.
// +genclient
// This API defines the type of logs, which can be audit or operational, and the external SIEM destination to send the logs.
type SIEMOrgForwarder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SIEMOrgForwarderSpec   `json:"spec,omitempty"`
	Status            SIEMOrgForwarderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// Contains a list of `SIEMOrgForwarder` objects.
type SIEMOrgForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SIEMOrgForwarder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SIEMOrgForwarder{}, &SIEMOrgForwarderList{})
}
