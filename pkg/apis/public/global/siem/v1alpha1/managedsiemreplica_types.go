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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
	siemv1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/siem/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// Represents a organization scoped managed SIEM instance replica.
// +genclient
type ManagedSIEMReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ManagedSIEMSpec          `json:"spec"`
	Status ManagedSIEMReplicaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a list of organization scoped managed SIEM instances.
type ManagedSIEMReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedSIEMReplica `json:"items"`
}

// Provides the specification of an organization scoped managed SIEM instance.
type ManagedSIEMSpec struct {
	// The resource provisioning of the SIEM instance.
	Resource siemv1alpha1.SIEMResourceProvisioning `json:"resource,omitempty"`
}

// Provides the status of a organization scoped managed SIEM replica instance.
type ManagedSIEMReplicaStatus struct {
	// The observations of this project's overall state.
	Conditions []metav1.Condition `json:"conditions"`
	// A list of current errors and the timestamp this field gets updated.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
	// The version of Elastic in this replica. Used to configure cross-cluster search.
	Version string `json:"version,omitempty"`
	// The ingest fqdn of this replica. Used to configure cross-cluster search.
	IngestFqdn string `json:"ingestFqdn,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ManagedSIEMReplica{},
		&ManagedSIEMReplicaList{},
	)
}
