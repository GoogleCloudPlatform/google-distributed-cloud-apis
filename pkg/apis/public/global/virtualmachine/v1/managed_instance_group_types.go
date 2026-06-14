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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=managedinstancegroups,singular=managedinstancegroup,shortName={mig,migs}
// +kubebuilder:printcolumn:name="Target Size",type="integer",JSONPath=".spec.targetSize"
// +kubebuilder:printcolumn:name="Updated Replicas",type="integer",JSONPath=".status.updatedReplicas"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ManagedInstanceGroup represents the configuration and state of a Managed Instance Group.
type ManagedInstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedInstanceGroupSpec   `json:"spec,omitempty"`
	Status ManagedInstanceGroupStatus `json:"status,omitempty"`
}

// ManagedInstanceGroupSpec defines the desired state of ManagedInstanceGroup.
type ManagedInstanceGroupSpec struct {
	// BaseInstanceName specifies the base name (prefix) attached to all instances in this group.
	// Must comply with RFC 1035 format (lowercase alphanumeric characters or hyphens, starting with a letter and ending with a letter or digit)
	// and cannot exceed 58 characters to leave room for generated instance name suffixes.
	// +optional
	// +kubebuilder:validation:MaxLength=58
	// +kubebuilder:validation:Pattern=`^[a-z]([-a-z0-9]*[a-z0-9])?$`
	BaseInstanceName string `json:"baseInstanceName,omitempty"`

	// DistributionPolicy specifies the intended distribution of managed instances across target zones.
	// Currently, only single-zone managed instance groups are supported. If multiple zones are specified,
	// only the first zone in the list will be used by the controller, and remaining zones will be ignored.
	// +optional
	DistributionPolicy DistributionPolicy `json:"distributionPolicy,omitempty"`

	// InstanceTemplate references the instance template object used for creating new instances in the group.
	InstanceTemplate corev1.LocalObjectReference `json:"instanceTemplate"`

	// TargetSize is the target number of running instances managed by this group.
	// Must be a non-negative integer (greater than or equal to 0).
	// +kubebuilder:validation:Minimum=0
	TargetSize int32 `json:"targetSize"`
}

// DistributionPolicy specifies distribution details across target zones.
type DistributionPolicy struct {
	// Zones is the list of target zones where managed instances will be created.
	// Currently, only single-zone deployment is supported; only the first zone in this list will be targeted.
	// +optional
	Zones []TargetZone `json:"zones,omitempty"`
}

// TargetZone specifies a single target zone for instance placement.
type TargetZone struct {
	// Zone is the name of the target zone (e.g., "us-central1-a").
	Zone string `json:"zone"`
}

// ManagedInstanceGroupStatus defines the observed state of ManagedInstanceGroup.
type ManagedInstanceGroupStatus struct {
	// Conditions represents the observations of this Managed Instance Group's overall state.
	// +optional
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the most recent generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// UpdatedReplicas tracks the total number of instances currently running the latest InstanceTemplate.
	// +optional
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`

	// CurrentActions tracks the list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
	// +optional
	CurrentActions CurrentActions `json:"currentActions,omitempty"`
}

// CurrentActions tracks the list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
type CurrentActions struct {
	// None is the number of instances in the managed instance group that are running and have no scheduled actions.
	// +optional
	None int32 `json:"none,omitempty"`

	// Creating is the number of instances in the managed instance group that are scheduled to be created or are currently being created.
	// +optional
	Creating int32 `json:"creating,omitempty"`

	// Deleting is the number of instances in the managed instance group that are scheduled to be deleted or are currently being deleted.
	// +optional
	Deleting int32 `json:"deleting,omitempty"`

	// Recreating is the number of instances in the managed instance group that are scheduled to be recreated or are currently being recreated.
	// Recreating an instance deletes the existing root persistent disk and creates a new disk from the image that is defined in the instance template.
	// +optional
	Recreating int32 `json:"recreating,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedInstanceGroupList contains a list of ManagedInstanceGroup.
type ManagedInstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedInstanceGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&ManagedInstanceGroup{},
		&ManagedInstanceGroupList{},
	)
}
