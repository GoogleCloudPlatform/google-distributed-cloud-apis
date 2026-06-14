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

	"gke-internal.googlesource.com/private-cloud/pkg/apis/common/global/v1alpha1"
	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
	rmv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/resourcemanager/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// Represents a replicated ProjectServiceAccount resource that will be synced to a
// particular zonal API server.
// A ProjectServiceAccount resource will have a replica for each zone. Upon an update of
// the ProjectServiceAccount resource, the replicas will be progressively updated based
// on the resource's rollout strategy.
// +genclient
type ProjectServiceAccountReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectServiceAccountSpec          `json:"spec,omitempty"`
	Status ProjectServiceAccountReplicaStatus `json:"status,omitempty"`
}

var _ v1alpha1.MuxReplicaInterface = &ProjectServiceAccountReplica{}

// +kubebuilder:object:root=true

// Represents a collection of project service account replicas.
type ProjectServiceAccountReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ProjectServiceAccountReplica `json:"items"`
}

// Provides the desired state of a project.
type ProjectServiceAccountSpec struct {
	// The public keys used to verify the signature of the JWTs for the
	// `ProjectServiceAccount` resource.
	// + optional
	Keys []ProjectServiceAccountKey `json:"keys,omitempty"`
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
	Algorithm rmv1.ProjectServiceAccountKeyAlgorithm `json:"algorithm"`

	// The ID of the key. This is used to determine which key to verify against.
	ID string `json:"id"`

	// The base64 encoded public key to verify against.
	Key string `json:"key"`

	// The expiration time for the key.
	ValidBefore metav1.Time `json:"validBefore"`

	// The start time when the key becomes valid.
	ValidAfter metav1.Time `json:"validAfter"`
}

// Provides the status of a project replica.
type ProjectServiceAccountReplicaStatus struct {
	// Conditions represents the observations of this project's overall
	// state.
	// +listType=map
	// +listMapKey=type
	// + optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ErrorStatus contain a list of current errors and the timestamp this field
	// gets updated.
	// + optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ProjectServiceAccountReplica{},
		&ProjectServiceAccountReplicaList{},
	)
}
