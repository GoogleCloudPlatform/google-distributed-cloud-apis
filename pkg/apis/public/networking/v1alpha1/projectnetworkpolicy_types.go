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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	rmv1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/resourcemanager/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pnp
// ProjectNetworkPolicy is the Schema for the projectnetworkpolicies API.
// +genclient
type ProjectNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec contains the desired configuration for ProjectNetworkPolicy.
	Spec ProjectNetworkPolicySpec `json:"spec,omitempty"`

	// Status contains the observed state for ProjectNetworkPolicy.
	Status ProjectNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// ProjectNetworkPolicyList contains a list of ProjectNetworkPolicy.
type ProjectNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectNetworkPolicy `json:"items"`
}

// ProjectNetworkPolicySpec defines the desired state of ProjectNetworkPolicy.
// The array of ingress or egress rules for this policy applies to the specified
// subject/ target.
// When multiple rules are present, or when multiple policies are present,
// the rules for each are combined additively. In other words, traffic is
// allowed if it matches at least one rule.
type ProjectNetworkPolicySpec struct {
	// Subject specifies the target of the project network policies. If not
	// specified, all pods excluding the managed services in the project are
	// selected.
	// +kubebuilder:default:={subjectType:UserWorkload}
	Subject ProjectNetworkPolicySubject `json:"subject,omitempty"`
	// PolicyType specifies the direction of traffic on which the policy rules are
	// applied. This must be set to one of `Ingress` and `Egress`.
	// If not set, then it defaults to `Ingress`.
	// +kubebuilder:default:=Ingress
	PolicyType PolicyType `json:"policyType,omitempty"`
	// Ingress defines the list of ingress rules for this policy.
	// If this field is empty or nil, the ProjectNetworkPolicy does not allow any
	// traffic (and serves solely to ensure that subjects it selects are isolated
	// by default).
	// +optional
	// +kubebuilder:validation:MaxItems:=1
	Ingress []ProjectNetworkPolicyIngressRule `json:"ingress,omitempty"`
	// Egress defines the list of egress rules for this policy.
	// If this field is empty or nil, the ProjectNetworkPolicy does not allow any
	// traffic (and serves solely to ensure that subjects it selects are isolated
	// by default).
	// +optional
	// +kubebuilder:validation:MaxItems:=1
	Egress []ProjectNetworkPolicyEgressRule `json:"egress,omitempty"`
}

// ProjectNetworkPolicyStatus defines the observed state of ProjectNetworkPolicy.
type ProjectNetworkPolicyStatus struct {
	// If Ready is True, it means that all NetworkPolicies are successfully
	// propagated to all user clusters; if Ready is False, it means that some
	// (or all) NetworkPolicies have failed to propagate.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// PropagatedName is the name of the propagated NetworkPolicy realized in
	// all user clusters within the project.
	// Expected to be set when SubjectType="UserWorkload".
	PropagatedName string `json:"propagatedName,omitempty"`

	// Clusters is the list of propagation status on the clusters.
	// Expected to be set when SubjectType="UserWorkload".
	Clusters []rmv1alpha1.ClusterStatus `json:"clusters,omitempty"`

	// PropagatedManagedServiceNamespaces is the list of managed service namespaces
	// that the policy is propagated to. Expected to be set when SubjectType="ManagedService".
	PropagatedManagedServiceNamespaces []string `json:"propagatedManagedServiceNamespaces,omitempty"`
}

// PolicyType specifies the direction of traffic on which the policy rules are
// applied.
// +kubebuilder:validation:Enum=Ingress;Egress
type PolicyType string

const (
	// PolicyTypeIngress is for Ingress policies.
	PolicyTypeIngress PolicyType = "Ingress"
	// PolicyTypeEgress is for Ingress policies.
	PolicyTypeEgress PolicyType = "Egress"
)

// ProjectNetworkPolicySubject defines the target for project network policies.
type ProjectNetworkPolicySubject struct {
	// SubjectType specifies the type of entities the policy rules apply to.
	// This must be set to one of `UserWorkload` and `ManagedService`.
	// If not set, then it defaults to `UserWorkload`.
	// If set to UserWorkload, then all pods excluding the managed services in
	// the project are selected.
	// If set to ManagedService, then specified managed services are selected.
	// +kubebuilder:default:=UserWorkload
	SubjectType PolicySubjectType `json:"subjectType,omitempty"`

	// ManagedServices selects the managed services that the policy rules apply to.
	// Must be specified only with SubjectType=ManagedService.
	ManagedServices *PolicyManagedServiceSubject `json:"managedServices,omitempty"`
}

// PolicySubjectType defines the target type of the network policies.
// +kubebuilder:validation:Enum=UserWorkload;ManagedService
type PolicySubjectType string

const (
	// PolicySubjectTypeUserWorkload is the subject type for user pods/ workloads
	// in the project.
	PolicySubjectTypeUserWorkload PolicySubjectType = "UserWorkload"
	// PolicySubjectTypeManagedService is the subject type for managed services
	// in the project.
	PolicySubjectTypeManagedService PolicySubjectType = "ManagedService"
)

// PolicyManagedServiceSubject defines a managed service target.
// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1
type PolicyManagedServiceSubject struct {
	// MatchTypes specifies the managed service types that the policy applies to.
	// Exactly one item can be specified.
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:MaxItems:=1
	MatchTypes []string `json:"matchTypes,omitempty"`
}

// ProjectNetworkPolicyIngressRule defines a single ingress rule for a ProjectNetworkPolicy.
type ProjectNetworkPolicyIngressRule struct {
	// List of ports for incoming traffic.
	// Each item in this list is combined using a logical OR. If this field is
	// empty or missing, this rule matches all ports (traffic not restricted by
	// port).
	// If this field is present and contains at least one item, then this rule
	// allows traffic only if the traffic matches at least one port in the
	// list.
	// +optional
	Ports []ProjectNetworkPolicyPort `json:"ports,omitempty"`

	// List of sources which should be able to access the subject of the policy.
	// Items in this list are combined using a logical OR operation.
	// If this field is empty or missing, this rule matches all sources
	// (traffic not restricted by source).
	// If this field contains at least one item, this rule allows traffic only
	// if the traffic matches at least one item in the "from" list.
	// At max one item can be specified.
	// +optional
	// +kubebuilder:validation:MaxItems:=1
	From []ProjectNetworkPolicyPeer `json:"from,omitempty"`
}

// ProjectNetworkPolicyEgressRule defines a single egress rule for a ProjectNetworkPolicy.
type ProjectNetworkPolicyEgressRule struct {
	// List of destination ports outgoing traffic.
	// Each item in this list is combined using a logical OR. If this field is
	// empty or missing, this rule matches all ports (traffic not restricted by
	// port).
	// If this field is present and contains at least one item, then this rule
	// allows traffic only if the traffic matches at least one port in the
	// list.
	// +optional
	Ports []ProjectNetworkPolicyPort `json:"ports,omitempty"`

	// List of destinations for outgoing traffic of the subject for this rule.
	// Items in this list are combined using a logical OR operation.
	// If this field is empty or missing, this rule matches all destinations
	// (traffic not restricted by destination).
	// If this field contains at least one item, this rule allows traffic only
	// if the traffic matches at least one item in the "to" list.
	// If this
	// At max one item can be specified.
	// +optional
	// +kubebuilder:validation:MaxItems:=1
	To []ProjectNetworkPolicyPeer `json:"to,omitempty"`
}

// ProjectNetworkPolicyPeer describes a peer to allow traffic from.
// Exactly one of the subfields must be specified.
// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1
type ProjectNetworkPolicyPeer struct {
	// Projects defines the projects to apply the network policy to.
	// +optional
	Projects *PolicyProjects `json:"projects,omitempty"`

	// IPBlock defines policy on a particular IPBlock.
	// If empty, then all external IPs (excludes k8s nodes, workloads in the
	// organization) are selected.
	// +optional
	IPBlock *networkingv1.IPBlock `json:"ipBlock,omitempty"`
}

// PolicyProjects is used to match a set of projects.
// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1
type PolicyProjects struct {
	// MatchNames selects the projects within the organization on their name.
	// The project namespace is derived from the project network policy's namespace.
	// If this field is empty or missing, this rule matches all projects.
	// At max one item can be specified.
	// +kubebuilder:validation:MaxItems:=1
	MatchNames []string `json:"matchNames,omitempty"`
}

// ProjectNetworkPolicyPort describes a port to allow traffic on.
// If all subfields are empty, all TCP traffic is selected.
type ProjectNetworkPolicyPort struct {
	// The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this
	// field defaults to TCP.
	// +optional
	Protocol *v1.Protocol `json:"protocol,omitempty"`

	// The port on the given protocol. This can either be a numerical or named
	// port on a pod. If this field is not provided, this matches all port names and
	// numbers.
	// If present, only traffic on the specified protocol AND port will be matched.
	// +optional
	Port *intstr.IntOrString `json:"port,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ProjectNetworkPolicy{}, &ProjectNetworkPolicyList{})
}
