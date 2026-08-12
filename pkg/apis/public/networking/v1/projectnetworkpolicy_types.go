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
	k8snetworkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
	rmv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pnp
// +kubebuilder:printcolumn:name="TYPE",type="string",JSONPath=".spec.policyType"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// Contains the Schema for the `ProjectNetworkPolicy` API.
// +genclient
type ProjectNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired configuration for `ProjectNetworkPolicy` resource.
	Spec ProjectNetworkPolicySpec `json:"spec,omitempty"`

	// The observed state for `ProjectNetworkPolicy` resource.
	Status ProjectNetworkPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// Defines a list of `ProjectNetworkPolicy` resources.
type ProjectNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectNetworkPolicy `json:"items"`
}

// Defines the desired state of `ProjectNetworkPolicy` resource.
// The array of ingress or egress rules for this policy applies to the specified
// subject or target.
// When multiple rules are present, or when multiple policies are present,
// the rules for each are combined additively. In other words, traffic is
// allowed if it matches at least one rule.
// +kubebuilder:validation:XValidation:rule="!(self.policyType == 'Ingress' && has(self.egress))",message="egress rules must not be specified when policyType is Ingress"
// +kubebuilder:validation:XValidation:rule="!(self.policyType == 'Egress' && has(self.ingress))",message="ingress rules must not be specified when policyType is Egress"
// +kubebuilder:validation:XValidation:rule="!(self.subject.subjectType == 'ManagedService' && has(self.ingress) && self.ingress.exists(r, has(r.protocols) && size(r.protocols) > 0))",message="protocols must not be specified for ManagedService subject"
// +kubebuilder:validation:XValidation:rule="!(self.subject.subjectType == 'ManagedService' && has(self.egress) && self.egress.exists(r, has(r.protocols) && size(r.protocols) > 0))",message="protocols must not be specified for ManagedService subject"
type ProjectNetworkPolicySpec struct {
	// The target of the project network policies. If unspecified,
	// all pods excluding the managed services in the project are
	// selected.
	// +kubebuilder:default:={subjectType:UserWorkload, userWorkloadSelector:{labelSelector:{workloads:{}}}}
	Subject ProjectNetworkPolicySubject `json:"subject,omitempty"`
	// The direction of traffic on which the policy rules are
	// applied. This must be set to one of `ingress` and `egress`.
	// If not set, then it defaults to `ingress`.
	// +kubebuilder:default:=Ingress
	PolicyType PolicyType `json:"policyType,omitempty"`
	// A list of ingress rules for this policy.
	// If this field is empty, the `ProjectNetworkPolicy` resource does not allow any
	// traffic and serves solely to ensure that the subjects it selects are isolated
	// by default.
	// +kubebuilder:validation:MaxItems:=1
	Ingress []ProjectNetworkPolicyIngressRule `json:"ingress,omitempty"`
	// A list of egress rules for this policy.
	// If this field is empty, the `ProjectNetworkPolicy` resource does not allow any
	// traffic and serves solely to ensure that subjects it selects are isolated
	// by default.
	// +kubebuilder:validation:MaxItems:=1
	Egress []ProjectNetworkPolicyEgressRule `json:"egress,omitempty"`
}

// Defines the observed state of a `ProjectNetworkPolicy` resource.
type ProjectNetworkPolicyStatus struct {
	// If `ready` is `true`, it means that all network policies are successfully
	// propagated to all user clusters. if `ready` is `false`, it means that some,
	// or all, network policies have failed to propagate.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The name of the propagated network policy realized in
	// all user clusters within the project.
	// This field is expected to be set when the property of `SubjectType="UserWorkload"`.
	PropagatedName string `json:"propagatedName,omitempty"`

	// The list of propagation status on the clusters.
	// This field is expected to be set when the property of `SubjectType="UserWorkload"`.
	Clusters []rmv1alpha1.ClusterStatus `json:"clusters,omitempty"`

	// The list of managed service namespaces
	// that the policy is propagated to. This field is expected to be set when the property of `SubjectType="ManagedService"`.
	PropagatedManagedServiceNamespaces []string `json:"propagatedManagedServiceNamespaces,omitempty"`

	// ErrorStatus holds most recent errors with last seen time.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// Defines the direction of traffic on which the policy rules are
// applied.
// +kubebuilder:validation:Enum=Ingress;Egress
type PolicyType string

const (
	// PolicyTypeIngress is for Ingress policies.
	PolicyTypeIngress PolicyType = "Ingress"
	// PolicyTypeEgress is for Ingress policies.
	PolicyTypeEgress PolicyType = "Egress"
)

// Defines the target for project network policies.
type ProjectNetworkPolicySubject struct {
	// The type of entities the policy rules apply to.
	// This must be set to one of `userWorkload` or `managedService`.
	// If not set, then it defaults to `userWorkload`.
	// If set to `userWorkload`, then all pods excluding the managed services in
	// the project are selected.
	// If set to `managedService`, then specified managed services are selected.
	// +kubebuilder:default:=UserWorkload
	SubjectType PolicySubjectType `json:"subjectType,omitempty"`

	// The managed services that the policy rules apply to.
	// Must be specified only with `SubjectType="ManagedService"`.
	ManagedServices *PolicyManagedServiceSubject `json:"managedServices,omitempty"`

	// WorkloadSelector selects the workloads in the project to which the policy rules apply.
	// If this field is nil or empty, this rule applies to all workloads in the project.
	// Deprecated: Use `userWorkloadSelector` instead.
	WorkloadSelector *metav1.LabelSelector `json:"workloadSelector,omitempty"`

	// UserWorkloadSelector selects the workloads (Pods or VMs) to which this
	// policy applies, based on a combination of cluster, namespace, and
	// workload labels.
	//
	// If this field is omitted or empty, the policy applies to all workloads
	// within the project, including those in Standard clusters.
	//
	// To select workloads from a specific Standard cluster, you must use the
	// `clusters` selector with a `matchLabels` entry for `kubernetes.io/metadata.name`.
	// +kubebuilder:validation:XValidation:rule="!(!has(self.labelSelector.clusters) && has(self.labelSelector.namespaces))",message="`labelSelector.namespaces` must not be specified without `labelSelector.clusters`"
	UserWorkloadSelector *WorkloadSelector `json:"userWorkloadSelector,omitempty"`
}

// Defines the target type of the network policies.
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

// Defines a managed service target.
// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1
type PolicyManagedServiceSubject struct {
	// The managed service types that the policy applies to.
	// Exactly one item must be specified.
	// +kubebuilder:validation:MinItems:=1
	// +kubebuilder:validation:MaxItems:=1
	MatchTypes []string `json:"matchTypes,omitempty"`
}

// ActionType defines the action to be taken on traffic matching a rule.
// +kubebuilder:validation:Enum=Allow;Deny
type ActionType string

const (
	// ActionTypeAllow allows traffic matching the rule.
	ActionTypeAllow ActionType = "Allow"
	// ActionTypeDeny denies traffic matching the rule.
	ActionTypeDeny ActionType = "Deny"
)

// Defines a single ingress rule for a `ProjectNetworkPolicy` resource.
// +kubebuilder:validation:XValidation:rule="self.action != 'Deny' || size(self.from) > 0",message="from must be specified for ingress Deny rules"
// +kubebuilder:validation:XValidation:rule="!((has(self.ports) && size(self.ports) > 0) && (has(self.protocols) && size(self.protocols) > 0))",message="ports and protocols must not be used simultaneously"
type ProjectNetworkPolicyIngressRule struct {
	// Action specifies whether the traffic matching this rule is allowed or denied.
	// If not set, it defaults to Allow.
	// +kubebuilder:default:=Allow
	Action ActionType `json:"action,omitempty"`

	// A list of ports for incoming traffic.
	// Each item in this list is combined using a logical `OR` operation. If this field is
	// empty or missing, this rule matches all ports, traffic is not restricted by
	// port.
	// If this field is present and contains at least one item, then this rule
	// allows traffic only if the traffic matches at least one port in the
	// list.
	// Deprecated: Use Protocols instead.
	Ports []ProjectNetworkPolicyPort `json:"ports,omitempty"`

	// Protocols allow for fine-grain matching of traffic on protocol-specific attributes.
	// If this field is empty or missing, this rule matches all protocols and all ports.
	Protocols []ProjectNetworkPolicyProtocol `json:"protocols,omitempty"`

	// A list of sources which are able to access the subject of the policy.
	// Items in this list are combined using a logical `OR` operation.
	// If this field is empty or missing, this rule matches all sources
	// , traffic is not restricted by source.
	// If this field contains at least one item, this rule allows traffic only
	// if the traffic matches at least one item in the `from` list.
	// A maximum of one item must be specified.
	// +kubebuilder:validation:MaxItems:=1
	From []ProjectNetworkPolicyPeer `json:"from,omitempty"`
}

// Defines a single egress rule for a `ProjectNetworkPolicy` resource.
// +kubebuilder:validation:XValidation:rule="self.action != 'Deny' || size(self.to) > 0",message="to must be specified for egress Deny rules"
// +kubebuilder:validation:XValidation:rule="!((has(self.ports) && size(self.ports) > 0) && (has(self.protocols) && size(self.protocols) > 0))",message="ports and protocols must not be used simultaneously"
type ProjectNetworkPolicyEgressRule struct {
	// Action specifies whether the traffic matching this rule is allowed or denied.
	// If not set, it defaults to Allow.
	// +kubebuilder:default:=Allow
	Action ActionType `json:"action,omitempty"`

	// A list of the destination ports showing outgoing traffic.
	// Each item in this list is combined using a logical `OR` operation. If this field is
	// empty or missing, this rule matches all ports, traffic is not restricted by
	// port.
	// If this field is present and contains at least one item, then this rule
	// allows traffic only if the traffic matches at least one port in the
	// list.
	// Deprecated: Use Protocols instead.
	Ports []ProjectNetworkPolicyPort `json:"ports,omitempty"`

	// Protocols allow for fine-grain matching of traffic on protocol-specific attributes.
	// If this field is empty or missing, this rule matches all protocols and all ports.
	Protocols []ProjectNetworkPolicyProtocol `json:"protocols,omitempty"`

	// A list of destinations for outgoing traffic of the subject for this rule.
	// Items in this list are combined using a logical `OR` operation.
	// If this field is empty or missing, this rule matches all destinations
	// , traffic is not restricted by destination.
	// If this field contains at least one item, this rule allows traffic only
	// if the traffic matches at least one item in the `to` list.
	// A maximum of one item must be specified.
	// +kubebuilder:validation:MaxItems:=1
	To []ProjectNetworkPolicyPeer `json:"to,omitempty"`
}

// Represents a peer to allow traffic from.
// Exactly one of the subfields must be specified.
// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1
type ProjectNetworkPolicyPeer struct {
	// The projects to apply the network policy to.
	// Deprecated: Use `projectSelector` instead.
	Projects *PolicyProjects `json:"projects,omitempty"`

	// ProjectSelector selects projects and workloads within those projects as a source of traffic.
	// If specified, allows traffic from workloads within the selected projects that match the workload selector.
	ProjectSelector *ProjectSelector `json:"projectSelector,omitempty"`

	// A policy on a particular `iPBlock`.
	// If empty, then all external IPs, excluding Kubernetes nodes and workloads in the
	// organization, are selected.
	// Deprecated: Use `ipBlocks` instead.
	IPBlock *k8snetworkingv1.IPBlock `json:"ipBlock,omitempty"`

	// A policy on particular `iPBlocks`.
	// If empty, then all external IPs, excluding Kubernetes nodes and workloads in the
	// organization, are selected.
	IPBlocks []k8snetworkingv1.IPBlock `json:"ipBlocks,omitempty"`
}

// Represents a collection of projects that is used to match a set of projects.
// +kubebuilder:validation:MaxProperties:=1
type PolicyProjects struct {
	// The selected projects which are chosen within
	// the organization based on their name.
	// The project namespace is derived from the project network policy's namespace.
	// If this field is empty or missing, this rule matches all projects.
	// A maximum of one item must be specified.
	// +kubebuilder:validation:MaxItems:=1
	// +kubebuilder:default:={}
	MatchNames []string `json:"matchNames,omitempty"`
}

// ProjectSelector selects projects and workloads.
// +kubebuilder:validation:MaxProperties:=2
// +kubebuilder:validation:XValidation:rule="!(has(self.workloads) && has(self.workloadSelector))",message="`workloadSelector` is deprecated and cannot be used with `userWorkloadSelector`"
// +kubebuilder:validation:XValidation:rule="!(size(self.projects.matchNames) == 0 && has(self.workloadSelector) && has(self.workloadSelector.labelSelector.clusters) && size(self.workloadSelector.labelSelector.clusters.matchLabels) == 1)",message="A project name must be specified in `projects.matchNames` when selecting a specific cluster"
type ProjectSelector struct {
	// The projects to apply the network policy to.
	// If empty, this rule matches all projects.
	// +kubebuilder:default:={}
	Projects *PolicyProjects `json:"projects,omitempty"`

	// The workloads to apply the network policy to.
	// If empty, all workloads in the selected projects are included.
	// Deprecated: Use `workloadSelector` instead.
	Workloads *metav1.LabelSelector `json:"workloads,omitempty"`

	// WorkloadSelector further refines the peer selection to specific workloads
	// (Pods or VMs) within the projects selected by the `projects` field.
	//
	// It allows filtering by a combination of cluster, namespace, and workload
	// labels. All specified selectors are combined using a logical AND.
	//
	// If this field is omitted, all workloads within the selected projects are
	// matched, including those in Standard clusters.
	//
	// If the `labelSelector.clusters` field is omitted, all clusters are matched.
	// To select workloads from a specific Standard cluster, you must use the `clusters`
	// selector with a `matchLabels` entry for `kubernetes.io/metadata.name`.
	// A project name must also be specified in `projects.matchNames` when selecting a specific cluster.
	WorkloadSelector *WorkloadSelector `json:"workloadSelector,omitempty"`
}

// Represents a port to allow traffic on.
// If all subfields are empty, all TCP traffic is selected.
type ProjectNetworkPolicyPort struct {
	// The protocol which traffic must match. The options are TCP, UDP, or SCTP. If unspecified, this
	// field defaults to TCP.
	Protocol *corev1.Protocol `json:"protocol,omitempty"`

	// The port on the given protocol. This can either be a numerical or named
	// port on a pod. If this field is not provided, this matches all port names and
	// numbers.
	// If present, only traffic on the specified protocol and port is matched.
	Port *intstr.IntOrString `json:"port,omitempty"`
}

// ProjectNetworkPolicyProtocol describes additional protocol-specific match rules.
// Exactly one field must be set.
//
// +kubebuilder:validation:MaxProperties:=1
// +kubebuilder:validation:MinProperties:=1
type ProjectNetworkPolicyProtocol struct {
	// TCP specific protocol matches.
	//
	// +optional
	TCP *ProjectNetworkPolicyProtocolTCP `json:"tcp,omitempty"`

	// UDP specific protocol matches.
	//
	// +optional
	UDP *ProjectNetworkPolicyProtocolUDP `json:"udp,omitempty"`

	// SCTP specific protocol matches.
	//
	// +optional
	SCTP *ProjectNetworkPolicyProtocolSCTP `json:"sctp,omitempty"`

	// ICMP specific protocol matches.
	//
	// +optional
	ICMP *ProjectNetworkPolicyProtocolICMP `json:"icmp,omitempty"`

	// DestinationNamedPort selects a destination port on a pod based on the
	// ContainerPort name.
	//
	// +optional
	DestinationNamedPort string `json:"destinationNamedPort,omitempty"`
}

// ProjectNetworkPolicyProtocolTCP defines the TCP protocol matches.
type ProjectNetworkPolicyProtocolTCP struct {
	// DestinationPort for the match.
	Port int32 `json:"port,omitempty"`
}

// ProjectNetworkPolicyProtocolUDP defines the UDP protocol matches.
type ProjectNetworkPolicyProtocolUDP struct {
	// DestinationPort for the match.
	Port int32 `json:"port,omitempty"`
}

// ProjectNetworkPolicyProtocolSCTP defines the SCTP protocol matches.
type ProjectNetworkPolicyProtocolSCTP struct {
	// DestinationPort for the match.
	Port int32 `json:"port,omitempty"`
}

// ProjectNetworkPolicyProtocolICMP defines the ICMP protocol matches.
type ProjectNetworkPolicyProtocolICMP struct {
	// If set, all ICMP types are allowed.
}

func init() {
	SchemeBuilder.Register(&ProjectNetworkPolicy{}, &ProjectNetworkPolicyList{})
}
