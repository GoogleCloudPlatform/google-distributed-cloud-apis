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

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false

// ManagedDNSZoneReplica represents a managed DNS zone.
type ManagedDNSZoneReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ManagedDNSZoneSpec          `json:"spec"`
	Status ManagedDNSZoneReplicaStatus `json:"status,omitempty"`
}

var _ v1alpha1.MuxReplicaInterface = &ManagedDNSZoneReplica{}

// +kubebuilder:object:root=true

// ManagedDNSZoneReplicaList represents a list of managed DNS zones.
type ManagedDNSZoneReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ManagedDNSZoneReplica `json:"items"`
}

// ManagedDNSZoneSpec provides the specification of a managed DNS zone.
type ManagedDNSZoneSpec struct {
	// The domain name for the DNS zone.
	// +kubebuilder:validation:MinLength=1
	// We restrict MaxLength to 239 (instead of 253) because the reconciler creates downstream
	// stub zones prefixed with "stub-public-" (12 chars) or "stub-private-" (13 chars).
	// This ensures the generated stub zone resource name does not exceed the Kubernetes limit of 253 characters.
	// +kubebuilder:validation:MaxLength=239
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9]([-a-zA-Z0-9]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([-a-zA-Z0-9]{0,61}[a-zA-Z0-9])?)*$`
	DNSName string `json:"dnsName"`
	// Human readable DNS zone description. Optional field.
	// +optional
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:Pattern=`^[ -~]*$`
	Description *string `json:"description,omitempty"`
	// The visibility of the DNS zone, i.e., whether it is a public zone,
	// or a private zone visible only to clients in the default customer
	// VPC network.
	// +kubebuilder:validation:Enum=PUBLIC;PRIVATE
	Visibility Visibility `json:"visibility"`
}

type Visibility string

const (
	Public  Visibility = "PUBLIC"
	Private Visibility = "PRIVATE"
)

// ManagedDNSZoneReplicaStatus provides the status of a managed DNS zone.
type ManagedDNSZoneReplicaStatus struct {
	// Conditions that will be used:
	// - Ready: This condition indicates whether the DNS zone has been
	//          added into the appropriate zonefile.
	// - Deleting: The DNS zone was marked for deletion.
	Conditions []metav1.Condition `json:"conditions"`
	// Name servers hosting the DNS zone.
	NameServers []string `json:"nameServers"`
}

func init() {
	SchemeBuilder.Register(
		&ManagedDNSZoneReplica{},
		&ManagedDNSZoneReplicaList{},
	)
}
