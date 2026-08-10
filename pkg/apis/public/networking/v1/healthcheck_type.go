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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=hc

// Specifies the backend service health checks.
//
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="CheckIntervalSec",type="integer",JSONPath=".spec.checkIntervalSec"
// +kubebuilder:printcolumn:name="TimeoutSec",type="integer",JSONPath=".spec.timeoutSec"
// +kubebuilder:printcolumn:name="HealthyThreshold",type="integer",JSONPath=".spec.healthyThreshold"
// +kubebuilder:printcolumn:name="UnhealthyThreshold",type="integer",JSONPath=".spec.unhealthyThreshold"
type HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   HealthCheckSpec   `json:"spec"`
	Status HealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of HealthCheck.
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthCheck `json:"items"`
}

// Describes the attributes that a user expects from a health check.
type HealthCheckSpec struct {
	// Defines the parameters for health check probes.
	//
	// +kubebuilder:validation:XValidation:rule="[has(self.tcpHealthCheck), has(self.httpHealthCheck), has(self.httpsHealthCheck)].exists_one(probeExists, probeExists)",message="Exactly one probe type must be specified."
	// +kubebuilder:validation:XValidation:rule="(has(self.tcpHealthCheck) == has(oldSelf.tcpHealthCheck)) && (has(self.httpHealthCheck) == has(oldSelf.httpHealthCheck)) && (has(self.httpsHealthCheck) == has(oldSelf.httpsHealthCheck))",message="Probe type is immutable"
	ProbeHandler `json:",inline"`

	// The amount of time in seconds from the start of one probe to the start of
	// the next one. Defaults to 5 when omitted or set to 0.
	// It is discouraged to set TimeoutSec longer than CheckIntervalSec.
	// This field is immutable.
	//
	// +kubebuilder:default=5
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="CheckIntervalSec is immutable"
	// +optional
	CheckIntervalSec int32 `json:"checkIntervalSec,omitempty"`

	// A time (in seconds) to wait before claiming failure.
	// Defaults to 5 when omitted or set to 0.
	// It is discouraged to set TimeoutSec longer than CheckIntervalSec.
	// This field is immutable.
	//
	// +kubebuilder:default=5
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="TimeoutSec is immutable"
	// +optional
	TimeoutSec int32 `json:"timeoutSec,omitempty"`

	// A number of sequential probes that must succeed for the endpoint to be
	// considered healthy. Defaults to 2 when omitted or set to 0.
	// This field is immutable.
	//
	// +kubebuilder:default=2
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="HealthyThreshold is immutable"
	// +optional
	HealthyThreshold int32 `json:"healthyThreshold,omitempty"`

	// A number of sequential probes that must fail for the endpoint to be
	// considered unhealthy. Defaults to 2 when omitted or set to 0.
	// This field is immutable.
	//
	// +kubebuilder:default=2
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UnhealthyThreshold is immutable"
	// +optional
	UnhealthyThreshold int32 `json:"unhealthyThreshold,omitempty"`
}

// Defines the available probes for the health check.
// One and only one of the fields must be specified.
type ProbeHandler struct {
	// Defines probes using TCP port.
	//
	// +optional
	TCPHealthCheck *TCPHealthCheck `json:"tcpHealthCheck,omitempty"`

	// Defines probes using HTTP port.
	//
	// +optional
	HTTPHealthCheck *HTTPHealthCheck `json:"httpHealthCheck,omitempty"`

	// Defines probes using HTTPS port.
	//
	// +optional
	HTTPSHealthCheck *HTTPSHealthCheck `json:"httpsHealthCheck,omitempty"`
}

// Specifies parameters for TCP health check probes.
type TCPHealthCheck struct {
	// A number of the port on which the health check will be performed.
	// Defaults to 80. This field is immutable.
	//
	// +kubebuilder:default=80
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Port is immutable"
	// +optional
	Port int32 `json:"port,omitempty"`
}

// Specifies parameters for HTTP health check probes.
type HTTPHealthCheck struct {
	// A number of the port on which the health check will be performed.
	// Defaults to 80. This field is immutable.
	//
	// +kubebuilder:default=80
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Port is immutable"
	// +optional
	Port int32 `json:"port,omitempty"`

	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the host header is set to the destination
	// IP address to which health check packets are sent. This field is immutable.
	//
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Host is immutable"
	// +optional
	Host string `json:"host,omitempty"`

	// The request path of the HTTP health check request. The default value
	// is `/`. This field does not support query parameters. Must comply with
	// RFC3986. This field is immutable.
	//
	// +kubebuilder:default="/"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="RequestPath is immutable"
	// +optional
	RequestPath string `json:"requestPath,omitempty"`

	// Creates a content-based HTTP health check. In addition to the
	// required 200 status code, user can configure the health
	// check to pass only when the backend sends this specific ASCII
	// response string within the first 1024 bytes of the HTTP response body.
	// This field is immutable.
	//
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Response is immutable"
	// +optional
	Response string `json:"response,omitempty"`
}

// Specifies parameters for HTTPS health check probes.
type HTTPSHealthCheck struct {
	// A number of the port on which the health check will be performed.
	// Defaults to 443. This field is immutable.
	//
	// +kubebuilder:default=443
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Port is immutable"
	// +optional
	Port int32 `json:"port,omitempty"`

	// The value of the host header in the HTTPS health check request. If
	// left empty (default value), the host header is set to the destination
	// IP address to which health check packets are sent. This field is immutable.
	//
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Host is immutable"
	// +optional
	Host string `json:"host,omitempty"`

	// The request path of the HTTPS health check request. The default value
	// is `/`. This field does not support query parameters. Must comply with
	// RFC3986. This field is immutable.
	//
	// +kubebuilder:default="/"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="RequestPath is immutable"
	// +optional
	RequestPath string `json:"requestPath,omitempty"`

	// Creates a content-based HTTPS health check. In addition to the
	// required 200 status code, user can configure the health
	// check to pass only when the backend sends this specific ASCII
	// response string within the first 1024 bytes of the HTTP response body.
	// This field is immutable.
	//
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Response is immutable"
	// +optional
	Response string `json:"response,omitempty"`
}

// Represents the status of a health check.
type HealthCheckStatus struct {
	// A list of conditions describing the current state of the health check.
	// Known condition types are:
	// * "Ready"
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// ErrorStatus holds the most recent errors with last seen time.
	//
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&HealthCheck{}, &HealthCheckList{})
}
