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

	globalv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
	networkingv1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/networking/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=bes

// Represents a load balancer configuration.
//
// +kubebuilder:printcolumn:name="HealthCheck",type="string",JSONPath=".spec.healthCheckName"
type BackendService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   BackendServiceSpec   `json:"spec"`
	Status BackendServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of BackendService.
type BackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendService `json:"items"`
}

// Represents the status of BackendService.
type BackendServiceStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"` // embed the duck type for global status

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []BackendServiceZoneStatus `json:"zones,omitempty"`

	// ErrorStatus holds the most recent errors with last seen time.
	//
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

type BackendServiceZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"` // embed the duck type for zone status

	// The reconciliation status of the replica collected from the zone.
	ReplicaStatus networkingv1.BackendServiceStatus `json:"replicaStatus,omitempty"`
}

// Holds information about the forwarding rule.
type ForwardingRuleRef struct {
	// A name of the referenced forwarding rule object.
	// This field is required. This field is immutable.
	Name string `json:"name"`
}

// BackendServiceReadyConditionReason defines the set of reasons that explain
// why a particular BackendService Ready condition status is set.
type BackendServiceReadyConditionReason string

const (
	// ZoneNotFound indicates that the referenced Zone objects cannot
	// be found in the API.
	ZoneNotFound BackendServiceReadyConditionReason = "ZoneNotFound"
)

func init() {
	SchemeBuilder.Register(&BackendService{}, &BackendServiceList{})
}
