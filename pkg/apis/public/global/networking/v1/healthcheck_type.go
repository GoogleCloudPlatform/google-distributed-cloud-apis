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
// +kubebuilder:resource:shortName=hc

// Specifies the backend service health checks.
//
// +kubebuilder:printcolumn:name="CheckIntervalSec",type="integer",JSONPath=".spec.checkIntervalSec"
// +kubebuilder:printcolumn:name="TimeoutSec",type="integer",JSONPath=".spec.timeoutSec"
// +kubebuilder:printcolumn:name="HealthyThreshold",type="integer",JSONPath=".spec.healthyThreshold"
// +kubebuilder:printcolumn:name="UnhealthyThreshold",type="integer",JSONPath=".spec.unhealthyThreshold"
type HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   networkingv1.HealthCheckSpec `json:"spec"`
	Status HealthCheckStatus            `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of HealthCheck.
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthCheck `json:"items"`
}

// Represents the status of a health check.
type HealthCheckStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"` // embed the duck type for global status

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []HealthCheckZoneStatus `json:"zones,omitempty"`

	// ErrorStatus holds the most recent errors with last seen time.
	//
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// Represents the status of a given zone's health check.
type HealthCheckZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"` // embed the duck type for zone status

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus networkingv1.HealthCheckStatus `json:"replicaStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&HealthCheck{}, &HealthCheckList{})
}
