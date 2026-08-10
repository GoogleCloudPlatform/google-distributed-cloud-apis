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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmaccessrequest, vmaccessrequests}
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="VM",type="string",JSONPath=".spec.vm"
// +kubebuilder:printcolumn:name="User",type="string",JSONPath=".spec.user"
// +kubebuilder:printcolumn:name="TTL",type="string",JSONPath=".spec.ssh.ttl"
// Represents an access request to a VM.
// +gdcloud:manifest:relevant=false,oc=vmm
// +genclient
type VirtualMachineAccessRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineAccessRequestSpec   `json:"spec"`
	Status            VirtualMachineAccessRequestStatus `json:"status,omitempty"`
}

// Defines the `VirtualMachineAccessRequest` specification.
type VirtualMachineAccessRequestSpec struct {
	// Specifies the name of the VM to access.
	VM string `json:"vm"`
	// Specifies the username for accessing the VM.
	User string `json:"user"`
	// Holds the ssh credentials used to access the VM.
	SSH SSHSpec `json:"ssh"`
}

// Stores the ssh credentials used to establish the connection.
type SSHSpec struct {
	// Specifies the public key to program for SSH access.
	Key string `json:"key"`
	// Specifies the length of time for which this key is valid, expressed in
	// hours, minutes and seconds. The default value is `24h0m0s`.
	// +kubebuilder:default="24h0m0s"
	TTL metav1.Duration `json:"ttl,omitempty"`
}

// State of `VirtualMachineAccessRequest`.
type VirtualMachineAccessRequestState string

const (
	// Specifies that the request is configured.
	VMAccessRequestStateConfigured VirtualMachineAccessRequestState = "configured"
	// Specifies that the request is pending.
	VMAccessRequestStatePending VirtualMachineAccessRequestState = "pending"
	// Specifies that the request failed.
	VMAccessRequestStateFailed VirtualMachineAccessRequestState = "failed"

	// Deprecated: ConditionTypeVMARReady is legacy and used for SLO calculations.
	// Use ConditionTypeVMARSucceeded instead.
	ConditionTypeVMARReady = "Ready"

	// ConditionTypeVMARSucceeded specifies that the VMAccessRequest has been processed.
	ConditionTypeVMARSucceeded = "Succeeded"

	ReasonVMARConfigured                    = "RequestConfigured"
	ReasonVMARFailed                        = "RequestFailed"
	ReasonVMARPending                       = "RequestPending"
	ReasonVMARVMNotFound                    = "VirtualMachineNotFound"
	ReasonVMARGuestEnvironmentNotEnabled    = "GuestEnvironmentNotEnabled"
	ReasonVMARGuestEnvironmentDisabled      = "GuestEnvironmentDisabled"
	ReasonVMARGuestEnvironmentPending       = "GuestEnvironmentPending"
	ReasonVMARGuestEnvironmentBeingDisabled = "GuestEnvironmentBeingDisabled"
	ReasonVMARAccessManagementDisabled      = "AccessManagementDisabled"
	ReasonVMARInvalidKeyFormat              = "InvalidKeyFormat"
	ReasonVMARAccessManagementNotReady      = "AccessManagementStateNotReady"
	ReasonVMARRpcFailure                    = "RpcFailure"
	ReasonVMARSyncFailed                    = "SyncFailed"
	ReasonVMARAgentFailure                  = "AgentFailure"
)

// Describes the status of the VirtualMachineAccessRequest.
type VirtualMachineAccessRequestStatus struct {
	// Specifies the state of `VirtualMachineAccessRequest`.
	State VirtualMachineAccessRequestState `json:"state,omitempty"`
	// Indicates the reason for the current status.
	Reason string `json:"reason,omitempty"`
	// Provides additional context for the current status.
	Message string `json:"message,omitempty"`
	// Specifies the time when the request was processed.
	ProcessedAt metav1.Time `json:"processedAt,omitempty"`
	// Details of observed state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []VMMError `json:"errors,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of VirtualMachineAccessRequest objects.
type VirtualMachineAccessRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineAccessRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineAccessRequest{}, &VirtualMachineAccessRequestList{})
}
