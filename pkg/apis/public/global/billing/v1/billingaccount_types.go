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
	common "gke-internal.googlesource.com/private-cloud/pkg/apis/common/global/v1alpha1"
	billing "gke-internal.googlesource.com/private-cloud/pkg/apis/public/billing/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="DisplayName",type="string",JSONPath=".spec.displayName"

// +genclient
// Represents a Billing Account.
// The identifier of the `BillingAccount` will be the "namespaced name",
// <namespace/name>. It will appear on the Invoices for this Billing Account,
// and be on the selectors for querying billing data in Dashboards.
// Limited access: This field might not be available as it may not be
// accredited for use in your deployment. You can access it when it's approved.
type BillingAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   billing.BillingAccountSpec `json:"spec,omitempty"`
	Status BillingAccountStatus       `json:"status,omitempty"`
}

var _ common.MuxResourceInterface = &BillingAccount{}

// Provides the overall status of a BillingAccount.
type BillingAccountStatus struct {
	common.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []BillingAccountZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a BillingAccount rolling out to a particular zone.
type BillingAccountZoneStatus struct {
	common.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus billing.BillingAccountStatus `json:"replicaStatus,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of BillingAccounts
type BillingAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []BillingAccount `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// +genclient
// Represents the link between a Billing Account and a Project or the
// Organization.
// Limited access: This field might not be available as it may not be
// accredited for use in your deployment. You can access it when it's approved.
type BillingAccountBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   billing.BillingAccountBindingSpec `json:"spec,omitempty"`
	Status BillingAccountBindingStatus       `json:"status,omitempty"`
}

var _ common.MuxResourceInterface = &BillingAccountBinding{}

// Provides the overall status of a BillingAccountBinding.
type BillingAccountBindingStatus struct {
	common.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []BillingAccountBindingZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a BillingAccountBinding rolling out to a particular
// zone.
type BillingAccountBindingZoneStatus struct {
	common.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus billing.BillingAccountBindingStatus `json:"replicaStatus,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of BillingAccountBinding
type BillingAccountBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []BillingAccountBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&BillingAccount{}, &BillingAccountList{},
		&BillingAccountBinding{}, &BillingAccountBindingList{},
	)
}
