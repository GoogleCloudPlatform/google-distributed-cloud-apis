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
)

// +kubebuilder:object:generate=true

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=si

// SaasInstance represents an single instance of a Saas service.
type SaasInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SaasInstanceSpec   `json:"spec,omitempty"`
	Status SaasInstanceStatus `json:"status,omitempty"`
}

type SaasInstanceSpec struct {
	// SaasType identifies the type of service for the current SaasInstance. This field is immutable.
	//
	// The SaasType is defined by the service producer and it must be unique in the environment all
	// Saas services are deployed. For example "prometheus".
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	SaasType string `json:"saasType,omitempty"`

	// ConsumerResource linked to the current SaasInstance.
	//
	// Saas services providers would require consumers to create a resource of certain kind in order
	// to create an instance of the service. For example, the Prometheus Operator would require the
	// consumers to create a `Prometheus` resource in order to deploy the Prometheus service. In
	// this example there would be a SaasInstance resource that is linked this Prometheus resource.
	//
	// This field is immutable.
	//
	// +kubebuilder:validation:Required
	ConsumerResource ConsumerResource `json:"consumerResource,omitempty"`
}

type SaasInstanceStatus struct{}

// +kubebuilder:object:root=true

type SaasInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SaasInstance `json:"items"`
}

type ConsumerResource struct {
	ApiGroup  string `json:"apiGroup,omitempty"`
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func (cr *ConsumerResource) IsEmpty() bool {
	// This works, because all fields are comparable (string primitives).
	// If that changes (like an array field is added) the implementation of this function would have to change as well.
	return cr == nil || (*cr == ConsumerResource{})
}

func init() {
	SchemeBuilder.Register(&SaasInstance{}, &SaasInstanceList{})
}
