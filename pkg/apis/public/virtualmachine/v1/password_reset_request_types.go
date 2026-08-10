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
// +kubebuilder:resource:shortName={vmpasswordresetrequest, vmpasswordresetrequests}
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="VM",type="string",JSONPath=".spec.vmName"
// +kubebuilder:printcolumn:name="User",type="string",JSONPath=".spec.user"
// Represents a password reset request for a given VM.
// +gdcloud:manifest:relevant=false,oc=vmm
// +genclient
type VirtualMachinePasswordResetRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachinePasswordResetRequestSpec   `json:"spec"`
	Status            VirtualMachinePasswordResetRequestStatus `json:"status,omitempty"`
}

// Defines the `VirtualMachinePasswordResetRequest` specification.
type VirtualMachinePasswordResetRequestSpec struct {
	// The name of the VM to request a password reset.
	// This field is immutable and cannot be updated after creating a password reset request.
	VMName string `json:"vmName"`
	// The name of the user to perform the password reset for a specified VM.
	// This field is immutable and cannot be updated after creating a password reset request.
	User string `json:"user"`
	// The public key to encrypt the new password for a VM.
	// This field is immutable and cannot be updated after creating a password reset request.
	PublicKey string `json:"publicKey"`
}

// Defines the state of a `VirtualMachinePasswordResetRequest` resource.
type VirtualMachinePasswordResetRequestState string

const (
	// A field that specifies the request is in the configured state.
	VMPasswordResetRequestStateConfigured VirtualMachinePasswordResetRequestState = "configured"
	// A field that specifies the request is in the pending state.
	VMPasswordResetRequestStatePending VirtualMachinePasswordResetRequestState = "pending"
	// A field that specifies the request is in the failed state.
	VMPasswordResetRequestStateFailed VirtualMachinePasswordResetRequestState = "failed"
	// ConditionTypeVMReady means whether the VMPasswordResetRequest is ready or not.
	// This is being added to be used in SLO calculations, where we determine good or bad states based on this.
	ConditionTypeVMPasswordResetRequestReady = "Ready"
)

// Describes the status of the `VirtualMachinePasswordResetRequest` resource.
type VirtualMachinePasswordResetRequestStatus struct {
	// The state of the `VirtualMachinePasswordResetRequest` resource.
	State VirtualMachinePasswordResetRequestState `json:"state,omitempty"`
	// The reason for the current status of the resource.
	Reason string `json:"reason,omitempty"`
	// A field that provides additional information for the current status.
	Message string `json:"message,omitempty"`
	// Details of observed state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// A field that specifies the time when the request processed.
	ProcessedAt metav1.Time `json:"processedAt,omitempty"`
	// The new password encrypted using the public
	// key provided in the request and encoded using base64.
	// To decrypt the password, use base64 to decode the string
	// and decrypt the result using RSA decryption.
	EncryptedPassword string `json:"encryptedPassword,omitempty"`
	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []VMMError `json:"errors,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of `VirtualMachinePasswordResetRequest` resources.
type VirtualMachinePasswordResetRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachinePasswordResetRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachinePasswordResetRequest{}, &VirtualMachinePasswordResetRequestList{})
}
