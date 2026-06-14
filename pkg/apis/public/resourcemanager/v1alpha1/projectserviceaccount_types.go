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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectServiceAccountKeyAlgorithm string

const (
	ES256 ProjectServiceAccountKeyAlgorithm = "ES256"
)

// Defines the desired state of the `ProjectServiceAccount` resource.
type ProjectServiceAccountSpec struct {
	// The public keys used to verify the signature of the JWTs for the
	// `ProjectServiceAccount` resource.
	Keys []ProjectServiceAccountKey `json:"keys,omitempty"`
}

// Defines the observed state of the `ProjectServiceAccount` resource.
type ProjectServiceAccountStatus struct {
	// If the `Ready` condition is `True`, all `ServiceAccount` resources are
	// successfully propagated to all clusters of its project. If the `Ready`
	// condition is `False`, some `ServiceAccount` resources have failed to
	// propagate. The `Ready` condition can transition from `True` to `Unknown` if
	// the corresponding `ServiceAccount` resource in a user cluster is modified,
	// which triggers another propagation.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The name of the propagated `ServiceAccount` resource.
	PropagatedName string `json:"propagatedName,omitempty"`

	// The list of all selected cluster names and the conditions of the propagated
	// resources in the clusters.
	Clusters []ClusterStatus `json:"clusters,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=psa
// +gdcloud:manifest:relevant=false,oc=rm
// +genclient
// Defines a project resource that propagates the service account to all user
// clusters the project spans across. The namespace of the
// `ProjectServiceAccount` resource corresponds to the project.
type ProjectServiceAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectServiceAccountSpec   `json:"spec,omitempty"`
	Status ProjectServiceAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// Contains a list of `ProjectServiceAccount` resources.
type ProjectServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectServiceAccount `json:"items"`
}

// Contains the key component used to verify the JWT signed by the private key
// for the `ProjectServiceAccount` resource. The JWT is used as part of the
// authentication flow. Currently, the `ProjectServiceAccountKey` resource only
// supports user-managed keys. Users can create and delete user-managed key
// pairs.
// Users are responsible for rotating these keys periodically to ensure the
// security of their service accounts. Users retain the private key of these key
// pairs, and the `ProjectServiceAccountKey` resource retains only the public
// key.
type ProjectServiceAccountKey struct {
	// The algorithm of the key. Currently only ES256 keys are supported.
	Algorithm ProjectServiceAccountKeyAlgorithm `json:"algorithm"`

	// The ID of the key. This is used to determine which key to verify against.
	ID string `json:"id"`

	// The base64 encoded public key to verify against.
	Key string `json:"key"`

	// The expiration date for the key.
	ValidBefore metav1.Time `json:"validBefore"`

	// The start date when the key becomes valid.
	ValidAfter metav1.Time `json:"validAfter"`
}

func init() {
	SchemeBuilder.Register(&ProjectServiceAccount{}, &ProjectServiceAccountList{})
}
