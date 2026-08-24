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

const (
	// Ready condition is true iff VMEA is enabled and ingress/egress are configured.
	ConditionTypeReady = "Ready"

	// Reason for 'Ready' condition.
	// Indicates VMEA getting configured to latest changes.
	ReasonPending = "Pending"
	// Indicates error encountered while reconciling VMEA.
	ReasonError = "Error"
	// Indicates VMEA is updated to latest configuration.
	ReasonUpdated = "Updated"

	// IngressReady Condition

	// ConditionTypeVMEAIngressReady indicates whether the external LoadBalancer resources
	// are provisioned and the IP address is allocated.
	ConditionTypeVMEAIngressReady    = "IngressReady"
	ReasonIngressConfigured          = "IngressConfigured"
	ReasonIngressFailed              = "IngressFailed"
	ReasonIngressConfigurationFailed = "IngressConfigurationFailed"
	ReasonIngressDisabled            = "IngressDisabled"
	ReasonIngressCleanupFailed       = "IngressCleanupFailed"
	ReasonIngressIPPending           = "IngressIPPending"
	ReasonIngressResourcesPending    = "IngressResourcesPending"

	// EgressReady Condition

	// ConditionTypeVMEAEgressReady indicates whether the Egress NAT is established
	// and the project-level IP is confirmed.
	ConditionTypeVMEAEgressReady    = "EgressReady"
	ReasonEgressConfigured          = "EgressConfigured"
	ReasonEgressFailed              = "EgressFailed"
	ReasonEgressConfigurationFailed = "EgressConfigurationFailed"
	ReasonEgressDisabled            = "EgressDisabled"
	ReasonEgressCleanupFailed       = "EgressCleanupFailed"
	ReasonEgressIPPending           = "EgressIPPending"

	// Message for 'Ready' condition.
	// Indicates whether the VMEA is waiting for VM to start.
	MessageWaitingForVMStart = "waiting for VM to start"
	// Indicates error encountered while enabling ingress.
	MessageFailedEnablingIngress = "failed to enable ingress"
	// Indicates error encountered while enabling egress.
	MessageFailedEnablingEgress = "failed to enable egress"
	// Indicates error encountered while updating egress ip.
	MessageFailedGettingEgressIP = "failed to get egress ip"
	// Indicates error encountered while checking if ingress is enabled.
	MessageFailedCheckingIngressEnable = "failed to check ingress enablement"
	// Indicates VM ingress service is pending with ip.
	MessageWaitingForIngressIP = "waiting for ingress ip"
	// Indicates whether the VMEA is disabled.
	MessageDisabledByUser = "disabled by user"
	// Indicates error encountered while disabling ingress.
	MessageFailedDisablingIngress = "failed to disable ingress"
	// Indicates error encountered while disabling egress.
	MessageFailedDisablingEgress = "failed to disable egress"

	// Event Reasons

	EventReasonVirtualMachineNotFound = "VirtualMachineNotFound"
	EventReasonEnablingIngressFailed  = "EnablingIngressFailed"
	EventReasonEnablingEgressFailed   = "EnablingEgressFailed"
	EventReasonEgressIPFetchFailed    = "EgressIPFetchFailed"
	EventReasonUpdatedEgressIP        = "UpdatedEgressIP"
	EventReasonUpdatedIngressIP       = "UpdatedIngressIP"
	EventReasonDisabledExternalAccess = "DisabledExternalAccess"
)

// Represents the request of accessing the external
// VRF for a VirtualMachine.
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmexternalaccess, vmexternalaccesses}
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Enabled",type=boolean,JSONPath=".spec.enabled"
// +kubebuilder:printcolumn:name="Ingress-IP",type=string,JSONPath=".status.ingressIP"
// +kubebuilder:printcolumn:name="Egress-IP",type=string,JSONPath=".status.egressIP"
// +gdcloud:manifest:relevant=false,oc=vmm
type VirtualMachineExternalAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineExternalAccessSpec   `json:"spec,omitempty"`
	Status VirtualMachineExternalAccessStatus `json:"status,omitempty"`
}

// Defines the specification of the `VirtualMachineExternalAccess` object.
type VirtualMachineExternalAccessSpec struct {
	// Specifies whether the external vrf is accessible for the VirtualMachine.
	Enabled bool `json:"enabled"`

	// The list of ports that are exposed by the VirtualMachine ingress service.
	// +patchMergeKey=port
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=port
	// +listMapKey=protocol
	Ports []ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`
}

// ServicePort contains information on service's port.
type ServicePort struct {
	// The name of this port within the service.
	Name string `json:"name,omitempty"`

	// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP".
	Protocol corev1.Protocol `json:"protocol,omitempty"`

	// The port that will be exposed by this service.
	// +kubebuilder:validation:Required.
	Port int32 `json:"port"`
}

// Defines the observed state of the `VirtualMachineExternalAccess` object.
type VirtualMachineExternalAccessStatus struct {
	// IngressIP specifies the IP address on the VirtualMachine ingress service.
	IngressIP string `json:"ingressIP,omitempty"`

	// EgressIP specifies the IP address on the egress NAT which is used by the VirtualMachine.
	EgressIP string `json:"egressIP,omitempty"`

	// Details of observed state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []VMMError `json:"errors,omitempty"`
}

// Contains a list of VirtualMachineExternalAccess.
// +kubebuilder:object:root=true

type VirtualMachineExternalAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VirtualMachineExternalAccess `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineExternalAccess{},
		&VirtualMachineExternalAccessList{},
	)
}
