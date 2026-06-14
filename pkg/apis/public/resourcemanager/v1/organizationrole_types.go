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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=or
// +gdcloud:manifest:relevant=false,oc=iam
// Provides a system namespace resource that propagates the `ClusterRole`
// configuration to all user clusters within the organization.
// +genclient
type OrganizationRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrganizationRoleSpec   `json:"spec,omitempty"`
	Status OrganizationRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// Contains a list of `OrganizationRole` resources.
type OrganizationRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationRole `json:"items"`
}

// Defines the desired state of the `OrganizationRole` resource.
type OrganizationRoleSpec struct {
	// The rules of the `ClusterRole` resource to create in all clusters.
	Rules []rbacv1.PolicyRule `json:"rules,omitempty"`
	// An optional field that describes the same aggregation logic as in the
	// Kubernetes `ClusterRole` object.
	AggregationRule *rbacv1.AggregationRule `json:"aggregationRule,omitempty"`
}

// Defines the observed state of the `OrganizationRole` object.
type OrganizationRoleStatus struct {
	// If the `Ready` condition is `True`, then all `ClusterRole` resources are
	// successfully propagated to all user clusters. If the `Ready` condition is
	// `False`, then some or all `ClusterRole` resources have failed to propagate.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The name of the propagated `ClusterRole` resource in all user clusters
	// within the organization.
	PropagatedName string `json:"propagatedName,omitempty"`

	// The list of propagation statuses on the clusters.
	Clusters []ClusterStatus `json:"clusters,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&OrganizationRole{}, &OrganizationRoleList{})
}
