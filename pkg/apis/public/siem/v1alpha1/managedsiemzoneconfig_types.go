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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=siem,component=siem,entities="zonal-configs"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:siem-instance-admin"
// +gdcloud:manifest:rbac="describe,list:siem-instance-viewer"

// Represents the resource provisioning for a SIEMaaS instance in a specific zone.
// +genclient
type ManagedSIEMResourceZoneConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedSIEMResourceZoneConfigSpec   `json:"spec"`
	Status            ManagedSIEMResourceZoneConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type ManagedSIEMResourceZoneConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ManagedSIEMResourceZoneConfig `json:"items"`
}

// Represents the resource provisioning of a SIEM instance.
type SIEMResourceProvisioning struct {
	// The size tier of the SIEM instance.
	// +kubebuilder:validation:Enum:=NONE;SMALL;MEDIUM;LARGE;XLARGE
	Tier PredefinedResourceTier `json:"tier"`
}

type PredefinedResourceTier string

const (
	// PredefinedResourceTierNone: no resource (disable)
	PredefinedResourceTierNone PredefinedResourceTier = "NONE"
	// PredefinedResourceTierSmall: 6TB (6Ti)
	PredefinedResourceTierSmall PredefinedResourceTier = "SMALL"
	// PredefinedResourceTierMedium: 10TB (10Ti)
	PredefinedResourceTierMedium PredefinedResourceTier = "MEDIUM"
	// PredefinedResourceTierLarge: 20TB (20Ti)
	PredefinedResourceTierLarge PredefinedResourceTier = "LARGE"
	// PredefinedResourceTierXLarge: 30TB (30Ti)
	PredefinedResourceTierXLarge PredefinedResourceTier = "XLARGE"
)

type ManagedSIEMResourceZoneConfigSpec struct {
	ManagedSIEMRef corev1.LocalObjectReference `json:"managedSIEMRef"`
	Zone           string                      `json:"zone"`
	Resource       SIEMResourceProvisioning    `json:"resource"`
}

type ManagedSIEMResourceZoneConfigStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&ManagedSIEMResourceZoneConfig{},
		&ManagedSIEMResourceZoneConfigList{},
	)
}
