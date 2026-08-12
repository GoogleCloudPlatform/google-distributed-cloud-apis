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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=backendservicepolicies
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=bsp

// Represents policies to be applied to one or more load balancers.
//
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="SessionAffinity",type="string",JSONPath=".spec.sessionAffinity"
// +gdcloud:manifest:relevant=true,oc=unet,component=compute,entities="backend-service-policies"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:external-load-balancer-admin,internal-load-balancer-admin,load-balancer-developer,load-balancer-admin"
// +gdcloud:manifest:rbac="describe,list:external-load-balancer-viewer,internal-load-balancer-viewer"
// +gdcloud:manifest:skipcodegen=true
type BackendServicePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   BackendServicePolicySpec   `json:"spec"`
	Status BackendServicePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of BackendServicePolicy.
type BackendServicePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendServicePolicy `json:"items"`
}

// Describes the attributes that a user expects from this backend policy.
// +kubebuilder:validation:XValidation:rule="has(oldSelf.sessionAffinity) == has(self.sessionAffinity)", message="SessionAffinity is immutable"
// +kubebuilder:validation:XValidation:rule="has(oldSelf.selectors) == has(self.selectors)", message="selectors is immutable"
type BackendServicePolicySpec struct {
	// The value of the priority is used to compare against multiple BackendServicePolicies when matched to the same BackendService
	// The lower the value, the higher the priority.
	// If multiple policies have the same priority, alphanumerical comparison of BackendServicePolicy name will be used as a tiebreaker.
	//
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=100
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Priority is immutable"
	Priority int32 `json:"priority"`

	// The Session Affinity mode applied to the Backend Policy.
	// This field is optional. This field is immutable.
	//
	// Allowed values:
	// - NONE
	//   requests will be routed to any backend. This is the default value.
	// - CLIENT_IP_DST_PORT_PROTO
	//   requests from the same 4-tuple (source IP, destination IP,
	//   destination port, protocol) will be routed to the same destination
	//   backend.
	//
	// If multiple policies match the same BackendProject, the policies are
	// ORed. A Backend Service will have Session Affinity enabled if any of
	// the policies affecting it has Session Affinity enabled.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=NONE
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="SessionAffinity is immutable"
	SessionAffinity SessionAffinity `json:"sessionAffinity,omitempty"`

	// A selector defining which BackendService(s) this policy is applied to.
	// This field is required. This field is immutable.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="has(self.matchLabels)", message="MatchLabels is required"
	// +kubebuilder:validation:XValidation:rule="!has(self.matchLabels) || size(self.matchLabels) > 0", message="MatchLabels must have at least 1 label"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Selectors are immutable"
	Selectors metav1.LabelSelector `json:"selectors,omitempty"`

	// A configuration that is utilized to determine connection draining.
	// +optional
	// +kubebuilder:validation:Optional
	ConnectionDraining *ConnectionDraining `json:"connectionDraining,omitempty"`
}

// Defines the connection draining configuration.
type ConnectionDraining struct {
	// The value for how long connections can last on backends that are
	// terminating or being removed. A value of 0 disables connection draining.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	DrainingTimeoutSec int32 `json:"drainingTimeoutSec,omitempty"`
}

// The type of Session Affinity that will be utilized for services.
// +kubebuilder:validation:Enum=CLIENT_IP_DST_PORT_PROTO;NONE
type SessionAffinity string

const (
	// SessionAffinityClientIpDstPortProto indicates that requests from the same
	// 4-tuple hash will be routed to the same destination backend.
	SessionAffinityClientIpDstPortProto SessionAffinity = "CLIENT_IP_DST_PORT_PROTO"

	// SessionAffinityNone indicates that Session Affinity is disabled.
	SessionAffinityNone SessionAffinity = "NONE"
)

// Represents the status of the Backend Service Policy.
type BackendServicePolicyStatus struct {
	// A list of conditions describing the current state of the Backend Service Policy.
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
	SchemeBuilder.Register(&BackendServicePolicy{}, &BackendServicePolicyList{})
}
