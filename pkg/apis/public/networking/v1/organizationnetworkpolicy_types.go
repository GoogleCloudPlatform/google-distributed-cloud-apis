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
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=onp
// Defines the Schema for the `OrganizationNetworkPolicy` API.
// +genclient
// +gdcloud:manifest:relevant=true,oc=unet,component=networking,entities="organization-network-policies"
// +gdcloud:manifest:verbs=create;delete;describe;update
// +gdcloud:manifest:rbac="create,delete,describe,update:org-network-policy-admin"
// +gdcloud:manifest:skipcodegen=true
type OrganizationNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired configuration for `OrganizationNetworkPolicy`.
	Spec OrganizationNetworkPolicySpec `json:"spec,omitempty"`

	// The observed state for `OrganizationNetworkPolicy`.
	Status OrganizationNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// Defines a list of `OrganizationNetworkPolicy` resources.
type OrganizationNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationNetworkPolicy `json:"items"`
}

// Defines the desired state of a `OrganizationNetworkPolicy` resource.
// The array of ingress rules for this policy applies to the specified target.
// When multiple rules are present, or when multiple policies are present,
// the rules for each are combined additively. Traffic is
// allowed if it matches at least one rule.
type OrganizationNetworkPolicySpec struct {
	// The managed services of the organization network policies.
	Subject OrganizationNetworkPolicySubject `json:"subject"`

	// The ingress rule for the traffic.
	// If `ingress` is empty or missing, it does not allow any traffic.
	// If this field contains at least one item, this rule allows
	// traffic only if the traffic matches at least one item in the `from` field.
	// +kubebuilder:validation:Optional
	Ingress []OrganizationNetworkPolicyIngressRule `json:"ingress,omitempty"`
}

// Defines the observed state of `OrganizationNetworkPolicy` resource.
type OrganizationNetworkPolicyStatus struct {
	// If `ready` is `true`, it means that the `OrganizationNetworkPolicy` resource is successfully
	// propagated to the org admin cluster. If `ready` is `false`, it means that the
	// `OrganizationNetworkPolicy` has failed to propagate.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The list of authorization policies
	// that are generated from the `OrganizationNetworkPolicy` resource.
	GeneratedAuthorizationPolicies []AuthorizationPolicyRef `json:"generatedauthorizationpolicies,omitempty"`

	// ErrorStatus holds most recent errors with last seen time.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// Represents a reference to the generated authorization policy.
type AuthorizationPolicyRef struct {
	// The name of the referent.
	Name string `json:"name,omitempty"`

	// The namespace of the referent.
	Namespace string `json:"namespace,omitempty"`
}

// Represents the organization service targets.
// Choose exactly one of the properties for the target.
// +kubebuilder:validation:XValidation:rule="has(self.services)",message="`services` must be specified"
type OrganizationNetworkPolicySubject struct {
	// The type of entities the policy rules apply to.
	// If not set, then it defaults to `ManagedService`.
	// +kubebuilder:default:=ManagedService
	SubjectType OrganizationNetworkPolicySubjectType `json:"subjectType,omitempty"`

	// The service to select.
	// Supports the organization multi-tenant service, including `UIConsole` and `APIServer`.
	// +kubebuilder:validation:XValidation:rule="has(self.matchTypes) && size(self.matchTypes) > 0",message="`matchTypes` must contain at least one service type"
	Services *ManagedServiceSubject `json:"services,omitempty"`
}

// Defines the target type of the policies.
// +kubebuilder:validation:Enum=ManagedService
type OrganizationNetworkPolicySubjectType string

const (
	// SubjectTypeManagedService is the subject type for managed services.
	SubjectTypeManagedService OrganizationNetworkPolicySubjectType = "ManagedService"
)

// Defines a managed service target.
type ManagedServiceSubject struct {
	// The organization managed service types that the policy applies to.
	MatchTypes []string `json:"matchTypes,omitempty"`
}

// Defines a single ingress rule for a `OrganizationNetworkPolicy` resource.
type OrganizationNetworkPolicyIngressRule struct {
	// A list of sources which are able to access the subject of the policy.
	// Items in this list are combined using a logical `OR` operation.
	// If this field is empty or missing, this rule matches all sources, the traffic is not restricted by source.
	// If this field contains at least one item, this rule allows traffic only
	// if the traffic matches at least one item in the `from` list.
	// A maximum of one item must be specified.
	From []OrganizationNetworkPolicyPeer `json:"from,omitempty"`
}

// Defines a peer to allow traffic from.
type OrganizationNetworkPolicyPeer struct {
	// A policy on a particular `iPBlock`.
	// If empty, it allows all traffic (0.0.0.0/0).
	IPBlock *networkingv1.IPBlock `json:"ipBlock,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&OrganizationNetworkPolicy{},
		&OrganizationNetworkPolicyList{},
	)
}
