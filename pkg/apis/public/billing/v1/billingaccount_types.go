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

// Contains API Schema definitions for Billing Account related APIs.
// +kubebuilder:object:generate=true
// +groupName=billing.gdc.goog
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="DisplayName",type="string",JSONPath=".spec.displayName"
// +kubebuilder:subresource:status

// +genclient
// +gdcloud:manifest:relevant=false,oc=bil
// Represents a Billing Account.
// The identifier of the `BillingAccount` will be the "namespaced name",
// <namespace/name>. It will appear on the Invoices for this Billing Account,
// and be on the selectors for querying billing data in Dashboards.
// <br>
// Limited access: This field might not be available as it may not be
// accredited for use in your deployment. You can access it when it's approved.
type BillingAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   BillingAccountSpec   `json:"spec"`
	Status BillingAccountStatus `json:"status,omitempty"`
}

// Defines the metadata of a Billing Account.
type BillingAccountSpec struct {
	// A human readable name or description of this Billing Account.
	// This name will be a selector for querying billing data in dashboards.
	DisplayName string `json:"displayName"`

	// The link to an external payment account.
	PaymentSystemConfig PaymentSystemConfig `json:"paymentSystemConfig"`
}

// Stores the link to an external payment account.
// Fields are one-of.
type PaymentSystemConfig struct {
	// A link to a customized account defined by Infrastructure Operators (IO).
	// To configure the customized account using `CustomConfig`:
	// <ul>
	// <li>It should include a key `payment-config-type`, with value defined by IO.</li>
	// <li>It should list the names of identifiers of the account in keys, and values in values.</li>
	// </ul>
	// Example: <br>
	//   "payment-config-type": "Example" <br>
	//   "account-id": "test-account-id-1"
	CustomConfig map[string]string `json:"customConfig,omitempty"`

	// A link to a Cloud Billing account.
	CloudBillingConfig PaymentSystemCloudBillingConfig `json:"cloudBillingConfig,omitempty"`
}

// Stores the config for a Cloud Billing account.
type PaymentSystemCloudBillingConfig struct {
	// The Cloud Billing account ID.
	AccountID string `json:"accountID"`
}

// Provides the status of a BillingAccount.
type BillingAccountStatus struct {
	// Conditions contains the latest time and state when Billing Platform
	// processes the BillingAccount.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
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
// Represents the link between a Billing Account and
// a Project or the Organization.
// Limited access: This field might not be available as it may not be
// accredited for use in your deployment. You can access it when it's approved.
type BillingAccountBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   BillingAccountBindingSpec   `json:"spec"`
	Status BillingAccountBindingStatus `json:"status,omitempty"`
}

// Defines the spec of a BillingAccountBinding.
type BillingAccountBindingSpec struct {
	// The Billing Account to use.
	// Required.
	BillingAccountRef corev1alpha1.NamespacedName `json:"billingAccountRef"`
}

// Shows the status of the BillingAccountBinding.
type BillingAccountBindingStatus struct {
	// Conditions contains the latest time and state when Billing Platform
	// processed the binding.
	// If the `Effective` condition is `true`, the binding in the
	// `ObservedGeneration` of the object is effective, otherwise, the previous
	// binding, or no binding is effective.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// BindingHistory keeps an append-only ledger of all past active billing account bindings.
	BindingHistory []BindingHistoryRecord `json:"bindingHistory,omitempty"`
}

// BindingHistoryRecord represents a historic binding point of a billing account to a project or organization.
type BindingHistoryRecord struct {
	BillingAccountRef corev1alpha1.NamespacedName `json:"billingAccountRef"`
	EffectiveTime     metav1.Time                 `json:"effectiveTime"`
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
