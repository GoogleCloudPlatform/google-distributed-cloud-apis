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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=fri

// Represents a frontend API to create internal forwarding rule.
//
// +kubebuilder:printcolumn:name="BackendService",type="string",JSONPath=".spec.backendServiceRef.name"
// +kubebuilder:printcolumn:name="CIDR",type="string",JSONPath=".status.cidr"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +gdcloud:manifest:relevant=true,oc=unet,component=compute,entities="forwarding-rules"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:internal-load-balancer-admin,load-balancer-admin"
// +gdcloud:manifest:rbac="describe,list:internal-load-balancer-viewer"
// +gdcloud:manifest:skipcodegen=true
type ForwardingRuleInternal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   ForwardingRuleInternalSpec   `json:"spec"`
	Status ForwardingRuleInternalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of ForwardingRuleInternal.
type ForwardingRuleInternalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ForwardingRuleInternal `json:"items"`
}

// Describes the attributes that a user expects from a forwarding rule.
type ForwardingRuleInternalSpec struct {
	ForwardingRuleSpecCommon `json:",inline"`
}

// Represents the status of forwarding rule.
type ForwardingRuleInternalStatus struct {
	ForwardingRuleStatusCommon `json:",inline"`
}

// Describes common attributes that a user expects from a forwarding rule.
//
// +kubebuilder:validation:XValidation:rule="has(oldSelf.cidrRef) == has(self.cidrRef)", message="CIDRRef is immutable"
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.backendServiceRef) || has(self.backendServiceRef)", message="BackendServiceRef is required once set"
type ForwardingRuleSpecCommon struct {

	// A reference to object holding the CIDR to use for this forwarding rule.
	// It has to reference object in the same namespace as this forwarding rule.
	// If not specified, an IPv4 /32 CIDR will be auto-reserved from the
	// global or zonal IP pool. This field is optional.
	// This field is immutable.
	//
	// +optional
	CIDRRef *CIDRRef `json:"cidrRef,omitempty"`

	// A list of L4 ports for which packets will be forwarded to the backends
	// configured with this forwarding rule. At least one port has to be
	// specified. The provided port-protocol pair has to be unique in the list.
	// For internal forwarding rules within the same VPC network or all external
	// forwarding rules, two or more forwarding rules cannot use the same
	// [CIDR, Protocol] pair if they share at least one port number.
	// This field is required. This field is immutable.
	//
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:MaxItems:=100
	// +kubebuilder:validation:XValidation:rule="self.all(a, self.exists_one(b, a.port == b.port && a.protocol == b.protocol))",message="Ports has duplicate port-protocol pair"
	Ports []Port `json:"ports"`

	// A reference to BackendService used for this forwarding rule. It has to
	// reference BackendService in the same namespace as this forwarding rule.
	// This field is immutable once set.
	//
	// +optional
	BackendServiceRef *BackendServiceRef `json:"backendServiceRef,omitempty"`
}

// Holds information about the CIDR.
type CIDRRef struct {
	// A name of the referenced cidr object.
	// This field is required. This field is immutable.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="CIDRRef name is immutable"
	// +kubebuilder:validation:MaxLength=250
	Name string `json:"name"`
}

// Contains information on L4 port on which service needs to be served.
type Port struct {
	// Specifies Layer-4 protocol which traffic must match. Only TCP and UDP
	// are supported. This field is required. This field is immutable.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=TCP;UDP
	Protocol *corev1.Protocol `json:"protocol"`

	// A number of the port that will be exposed by this service.
	// This field is required. This field is immutable.
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Required
	Port int32 `json:"port"`
}

func (p *Port) Equal(port *Port) bool {
	if p.Port == port.Port && *p.Protocol == *port.Protocol {
		return true
	}
	return false
}

// Holds information about the backend.
type BackendServiceRef struct {
	// A name of the referenced backend service object.
	// This field is required. This field is immutable.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="BackendServiceRef name is immutable"
	// +kubebuilder:validation:MaxLength=250
	Name string `json:"name"`
}

// Represents common status of ForwardingRule
type ForwardingRuleStatusCommon struct {
	// The resulting cidr value used for this forwarding rule.
	//
	// +optional
	CIDR string `json:"cidr,omitempty"`

	// A list of conditions describing the current state of the forwarding rule.
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

	// Subnets holds the list of subnets that are used for this forwarding rule.
	//
	// +optional
	Subnets []SubnetReference `json:"subnets,omitempty"`
}

// SubnetReference holds the reference to a subnet.
type SubnetReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// ForwardingRuleReadyConditionReason defines the set of reasons that explain
// why a particular ForwardingRule Ready condition status is set.
type ForwardingRuleReadyConditionReason string

const (
	// BackendServiceNotFound indicates that the referenced BackendService objects cannot be found in the API.
	BackendServiceNotFound ForwardingRuleReadyConditionReason = "BackendServiceNotFound"
	// BackendServiceNotReady indicates that the referenced BackendService object is not Ready.
	BackendServiceNotReady ForwardingRuleReadyConditionReason = "BackendServiceNotReady"
	// CIDRError indicates that there was an issue when handling IPAM for the ForwardingRule.
	CIDRError ForwardingRuleReadyConditionReason = "CIDRError"
	// CIDRNotAssigned indicates that the referenced Subnet object does not have an IP assigned.
	CIDRNotAssigned ForwardingRuleReadyConditionReason = "CIDRNotAssigned"
	// SubnetNotReady indicates that the referenced Subnet object is not Ready.
	SubnetNotReady ForwardingRuleReadyConditionReason = "SubnetNotReady"
	// SubnetNotFound indicates that the referenced Subnet object cannot be found in the API.
	SubnetNotFound ForwardingRuleReadyConditionReason = "SubnetNotFound"
	// PortOverlap indicates that the used port overlap with other ForwardingRule.
	PortOverlap ForwardingRuleReadyConditionReason = "PortOverlap"
	// BackendServicePoliciesError indicates we weren't able to retrieve the BackendServicePolicies
	BackendServicePoliciesError ForwardingRuleReadyConditionReason = "BackendServicePoliciesError"
	// ProjectClustersError indicates we weren't able to retrieve the clusters to manage services
	ProjectClustersError ForwardingRuleReadyConditionReason = "ProjectClustersError"
	// ServiceError indicates we weren't able to create our service
	ServiceError ForwardingRuleReadyConditionReason = "ServiceError"
)

func init() {
	SchemeBuilder.Register(&ForwardingRuleInternal{}, &ForwardingRuleInternalList{})
}
