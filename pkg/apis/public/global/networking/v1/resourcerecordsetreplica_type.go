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

// ResourceRecordSetReplica represents a resource record set.
type ResourceRecordSetReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ResourceRecordSetSpec          `json:"spec"`
	Status ResourceRecordSetReplicaStatus `json:"status,omitempty"`
}

var _ v1alpha1.MuxReplicaInterface = &ResourceRecordSetReplica{}

// +kubebuilder:object:root=true

// ResourceRecordSetReplicaList represents a list of resource record sets.
type ResourceRecordSetReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ResourceRecordSetReplica `json:"items"`
}

// ResourceRecordSetSpec provides the specification of a resource record set.
type ResourceRecordSetSpec struct {
	// The fully qualified domain name (FQDN) of the RRset.
	Name string `json:"name"`
	// The time to live in seconds for the RRset. This value indicates how long
	// a resolver should serve cached versions of resource records in this
	// RRset before querying for an updated version. Default value is 300.
	// +kubebuilder:default=300
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2000000000
	TTLSeconds *uint32 `json:"ttlSeconds,omitempty"`
	// The type of the resource records in the RRset.
	// +kubebuilder:validation:Enum=A;CNAME;MX;TXT;PTR
	Type string `json:"type"`
	// The data for all resource records in the RRset. Each entry represents a
	// separate resource record.
	RRData []string `json:"rrData"`
	// The managed DNS zone that the RRset belongs to.
	// The name of the ManagedDNSZone CR in the same namespace as this
	// ResourceRecordSet CR.
	// The domain name of the managed DNS zone must be a suffix of the FQDN
	// specified in the `Name` field above.
	DNSZone string `json:"dnsZone"`
}

// ResourceRecordSetReplicaStatus provides the status of a resource record set.
type ResourceRecordSetReplicaStatus struct {
	// Conditions that will be used:
	// - Ready: This condition indicates whether the RRset has been added to
	//          or updated in the appropriate zonefile.
	// - Deleting: The RRset was marked for deletion.
	Conditions []metav1.Condition `json:"conditions"`
}

func init() {
	SchemeBuilder.Register(
		&ResourceRecordSetReplica{},
		&ResourceRecordSetReplicaList{},
	)
}
