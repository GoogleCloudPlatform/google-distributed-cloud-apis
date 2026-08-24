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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// Represents the configuration of a CLoud NAT Gateway
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=cloudnatgateways,singular=cloudnatgateway,shortName=cng
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +gdcloud:manifest:relevant=false,oc=unet
type CloudNATGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the configuration of a Cloud NAT Gateway
	Spec CloudNATGatewaySpec `json:"spec"`

	// Status holds the most recently observed status of the gateway
	// +kubebuilder:default={conditions: {{type: "Ready", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}}
	Status CloudNATGatewayStatus `json:"status,omitempty"`
}

// Contains the list of CloudNAT Gateways
//
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CloudNATGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudNATGateway `json:"items"`
}

// CloudNATGatewaySpec describes the attributes of the gateway.
type CloudNATGatewaySpec struct {
	// WorkloadSelector defines which endpoints (Pods or VMs) will route egress traffic through this
	// gateway. Only the endpoints that have all selector labels are matched (AND matching).
	//
	// This field is immutable
	// +required
	// +kubebuilder:validation:XValidation:rule="!has(self.labelSelector.namespaces)",message="`workloadSelector.labelSelector.namespaces` field is not supported."
	// +kubebuilder:validation:XValidation:rule="has(self.labelSelector.workloads)",message="`workloadSelector.labelSelector.workloads` field is required."
	// +kubebuilder:validation:XValidation:rule="!has(self.labelSelector.workloads.matchExpressions)",message="`workloadSelector.labelSelector.workloads.matchExpressions` field is not supported."
	// +kubebuilder:validation:XValidation:rule="has(self.labelSelector.workloads.matchLabels) && size(self.labelSelector.workloads.matchLabels) > 0",message="`workloadSelector.labelSelector.workloads.matchLabels` must contain at least one entry."
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="WorkloadSelector is immutable"
	// +kubebuilder:validation:XValidation:rule="self.labelSelector.workloads.matchLabels == oldSelf.labelSelector.workloads.matchLabels",message="`workloadSelector.labelSelector.workloads.matchLabels` is immutable"
	WorkloadSelector *WorkloadSelector `json:"workloadSelector"`

	// SubnetsRefs specifies the list of subnets containing the egress IPs that will be used by the
	// CloudNATGateway. The subnets must be of leaf type and be in the same project namespace as this
	// CloudNATGateway. The maximum total number of IPs across all subnets must be lower or equal than 100. If the
	// number of IPs exceed 100, the gateway will use the first 100 IPs from the specified subnets in
	// the specified order.
	//
	// This field is mutable
	// +required
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:items:MinLength=1
	SubnetRefs []string `json:"subnetRefs"`

	// ConnectionOptions specifies the timeout for the new connections
	//
	// This field is mutable. Modifying this field will only apply to new connections.
	// +optional
	ConnectionOptions *CloudNATGatewayConnectionOptions `json:"connectionOptions,omitempty"`
}

// CloudNATGatewayTimeouts specifies the timeout for the different connections types.
type CloudNATGatewayConnectionOptions struct {
	// Timeout for non-TCP connections made through this gateway.
	//
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=86400
	NonTCPTimeoutSeconds *uint32 `json:"nonTCPTimeoutSeconds,omitempty"`

	// Timeout for established TCP connections made through this gateway.
	//
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=86400
	TCPTimeoutSeconds *uint32 `json:"tcpTimeoutSeconds,omitempty"`

	// Teardown timeout for TCP connections made through this gateway.
	//
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=86400
	TCPTeardownTimeoutSeconds *uint32 `json:"tcpTeardownTimeoutSeconds,omitempty"`

	// Establishment timeout for TCP connections initiated through this gateway.
	//
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=86400
	TCPEstablishmentTimeoutSeconds *uint32 `json:"tcpEstablishmentTimeoutSeconds,omitempty"`
}

// CloudNATGatewayStatus represents the status of the gateway.
type CloudNATGatewayStatus struct {
	// Conditions describe the current condition of the gateway.
	//
	// Known conditions types are:
	// * "Ready"
	// * "SubnetsReady"
	// * "PerimeterConfigurationReady"
	// * "EgressRoutesReady"
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:default={{type: "Ready", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"},{type: "SubnetsReady", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"},{type: "PerimeterConfigurationReady", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"},{type: "EgressRoutesReady", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// ErrorStatus holds the most recent errors with last seen time.
	//
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// Subnets contains the name and the status of each subnet currently in use by this gateway.
	//
	// +optional
	Subnets []CloudNATSubnetStatus `json:"subnets,omitempty"`
}

// CloudNATSubnetStatus contains the name and status message of one subnet.
type CloudNATSubnetStatus struct {
	// Name of the subnet.
	//
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Status message of the subnet.
	//
	// +required
	// +kubebuilder:validation:MinLength=1
	Status string `json:"status"`
}

const (
	// Condition type to indicate that the subnets from SubnetRefs are ready.
	CloudNATGatewaySubnetsReadyConditionType = "SubnetsReady"

	// Condition type to indicate that the Perimeter Configuration is ready.
	CloudNATGatewayPerimeterConfigurationReadyConditionType = "PerimeterConfigurationReady"

	// Condition type to indicate that the egress routes are ready.
	CloudNATGatewayEgressRoutesReadyConditionType = "EgressRoutesReady"

	// CloudNATGatewayInternalServerErrorReason indicates that there is some internal server error.
	CloudNATGatewayInternalServerErrorReason = "InternalServerError"

	// CloudNATGatewayInvalidSelectorReason indicates that the WorkloadSelector field is misconfigured.
	CloudNATGatewayInvalidSelectorReason = "InvalidSelector"

	// CloudNATGatewayGatewayNotReadyReason indicates the gateway is not ready because one or more of
	// the other conditions are not ready
	CloudNATGatewayOtherConditionsNotReadyReason = "OtherConditionsNotReady"

	// CloudNATGatewayInvalidSubnetsReason indicates some of the subnets are invalid.
	CloudNATGatewayInvalidSubnetsReason = "InvalidSubnets"

	// CloudNATGatewayPerimeterConfigNotReadyReason indicates that some perimeter cluster
	// configuration is not ready.
	CloudNATGatewayPerimeterConfigurationErrorReason = "PerimeterConfigurationError"

	// CloudNATGatewayEgressRoutesNotReadyReason indicates that the egress routes are not ready.
	CloudNATGatewayEgressRoutesErrorReason = "EgressRoutesError"
)

func init() {
	SchemeBuilder.Register(&CloudNATGateway{}, &CloudNATGatewayList{})
}
