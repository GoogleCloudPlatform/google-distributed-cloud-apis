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
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=onp
// OrganizationNetworkPolicy is the Schema for the organizationnetworkpolicies API.
// +genclient
type OrganizationNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec contains the desired configuration for OrganizationNetworkPolicy.
	Spec OrganizationNetworkPolicySpec `json:"spec,omitempty"`

	// Status contains the observed state for OrganizationNetworkPolicy.
	Status OrganizationNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// OrganizationNetworkPolicyList contains a list of OrganizationNetworkPolicy.
type OrganizationNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationNetworkPolicy `json:"items"`
}

// OrganizationNetworkPolicySpec defines the desired state of OrganizationNetworkPolicy.
// The array of ingress rules for this policy applies to the specified subject/target.
// When multiple rules are present, or when multiple policies are present,
// the rules for each are combined additively. In other words, traffic is
// allowed if it matches at least one rule.
type OrganizationNetworkPolicySpec struct {
	// Subject specifies the managed services of the organization network policies.
	// +kubebuilder:default:={subjectType:ManagedService}
	Subject OrganizationNetworkPolicySubject `json:"subject,omitempty"`

	// Ingress define the ingress rule for the traffic.
	// If Ingress is empty or missing, it does not allow any traffic.
	// If this field contains at least one item, this rule allows
	// traffic only if the traffic matches at least one item in the from.
	// +kubebuilder:validation:Optional
	Ingress []OrganizationNetworkPolicyIngressRule `json:"ingress,omitempty"`
}

// OrganizationNetworkPolicyStatus defines the observed state of OrganizationNetworkPolicy.
type OrganizationNetworkPolicyStatus struct {
	// If Ready is True, it means that the OrganizationNetworkPolicy is successfully
	// propagated to the org admin cluster; if Ready is False, it means that
	// OrganizationNetworkPolicy have failed to propagate.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// GeneratedAuthorizationPolicies is the list of authorization policies.
	// that are generated from OrganizationNetworkPolicy.
	GeneratedAuthorizationPolicies []AuthorizationPolicyRef `json:"generatedauthorizationpolicies,omitempty"`
}

// AuthorizationPolicyRef is a reference to the generated authorization policy.
type AuthorizationPolicyRef struct {
	// Name of the referent.
	Name string `json:"name,omitempty"`

	// Namespace of the referent.
	Namespace string `json:"namespace,omitempty"`
}

// OrganizationNetworkPolicySubject represents the organization service targets.
// Must choose exactly one of the properties for the target.
type OrganizationNetworkPolicySubject struct {
	// SubjectType specifies the type of entites the policy rules apply to.
	// If not set, then it defaults to `ManagedService`.
	// +kubebuilder:default:=ManagedService
	SubjectType OrganizationNetworkPolicySubjectType `json:"subjectType,omitempty"`

	// Support multiple different services including:
	// Org Multi-Tenant Service:
	// - UIConsole
	// - APIServer
	// Services represents the service you want to select
	// Support for MT managed services like ODS in the future.
	Services *ManagedServiceSubject `json:"services,omitempty"`
}

// OrganizationNetworkPolicySubjectType defines the target type of the policies.
// +kubebuilder:validation:Enum=ManagedService
type OrganizationNetworkPolicySubjectType string

const (
	// SubjectTypeManagedService is the subject type for managed services.
	SubjectTypeManagedService OrganizationNetworkPolicySubjectType = "ManagedService"
)

// ManagedServiceSubject defines a managed service target.
type ManagedServiceSubject struct {
	// MatchTypes specifies the org managed service types that the policy applies to.
	// +kubebuilder:validation:MinItems:=1
	MatchTypes []string `json:"matchTypes,omitempty"`
}

// OrganizationNetworkPolicyIngressRule defines a single ingress rule for a OrganizationNetworkPolicy.
type OrganizationNetworkPolicyIngressRule struct {
	// List of sources which should be able to access the subject of the policy.
	// Items in this list are combined using a logical OR operation.
	// If this field is empty or missing, this rule matches all sources (traffic not restricted by source).
	// If this field contains at least one item, this rule allows traffic only
	// if the traffic matches at least one item in the "from" list.
	// At max one item can be specified.
	// +optional
	From []OrganizationNetworkPolicyPeer `json:"from,omitempty"`
}

// OrganizationNetworkPolicyPeer describes a peer to allow traffic from.
type OrganizationNetworkPolicyPeer struct {
	// IPBlock defines policy on a particular IPBlock.
	// If empty, then allows all traffic (0.0.0.0/0).
	// +optional
	IPBlock *networkingv1.IPBlock `json:"ipBlock,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&OrganizationNetworkPolicy{},
		&OrganizationNetworkPolicyList{},
	)
}
