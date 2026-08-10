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

// +kubebuilder:object:root=true
// +genclient
// Represents a view of a Realm.
type RealmView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RealmViewSpec `json:"spec,omitempty"`
}

type RealmViewSpec struct {
	// Name of the realm.
	Name string `json:"name"`

	// Client IDs of the configured OIDC clients for a realm.
	// +optional
	OIDCRealmClientIDs []string `json:"oidcRealmClientIDs,omitempty"`

	// Secret references for base64 PEM encoded x509 realm OIDC secrets.
	// +optional
	OIDCClientSecrets []string `json:"oidcClientSecrets,omitempty"`

	// Secret reference for the base64 PEM encoded x509 realm SAML secret.
	// +optional
	EnabledSAMLClientSecret string `json:"enabledSAMLClientSecret,omitempty"`
}

// +kubebuilder:object:root=true
// Represents a list of `RealmView` objects.
type RealmViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RealmView `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&RealmView{},
		&RealmViewList{},
	)
}
