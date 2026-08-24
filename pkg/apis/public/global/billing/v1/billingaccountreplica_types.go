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
	common "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
	billing "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/billing/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="DisplayName",type="string",JSONPath=".spec.displayName"

// +genclient
// Represents a replicated BillingAccount that will be synced to a particular
// zonal API server.
type BillingAccountReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   billing.BillingAccountSpec   `json:"spec,omitempty"`
	Status billing.BillingAccountStatus `json:"status,omitempty"`
}

var _ common.MuxReplicaInterface = &BillingAccountReplica{}

// +kubebuilder:object:root=true

// Contains a list of BillingAccountReplica
type BillingAccountReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []BillingAccountReplica `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// +genclient
// Represents a replicated BillingAccountBinding that will be synced to a
// particular zonal API server.
type BillingAccountBindingReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   billing.BillingAccountBindingSpec   `json:"spec,omitempty"`
	Status billing.BillingAccountBindingStatus `json:"status,omitempty"`
}

var _ common.MuxReplicaInterface = &BillingAccountBindingReplica{}

// +kubebuilder:object:root=true

// Contains a list of BillingAccountBindingReplica
type BillingAccountBindingReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []BillingAccountBindingReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&BillingAccountReplica{}, &BillingAccountReplicaList{},
		&BillingAccountBindingReplica{}, &BillingAccountBindingReplicaList{},
	)
}
