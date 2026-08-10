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

	"gke-internal.googlesource.com/private-cloud/pkg/apis/common/global/v1alpha1"
)

type BucketLocationType string

const (
	AsyncDualZoneLocationType BucketLocationType = "AsyncDualZone"
	SyncDualZoneLocationType  BucketLocationType = "SyncDualZone"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Name",JSONPath=".metadata.name",type=string
// +kubebuilder:printcolumn:name="Active",JSONPath=".spec.active",type=boolean
// +kubebuilder:printcolumn:name="Location Type",JSONPath=".spec.locationType",type=string
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type=string
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// Defines the schema for the BucketLocation API.
// +genclient
// +genclient:nonNamespaced
type BucketLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketLocationSpec   `json:"spec,omitempty"`
	Status BucketLocationStatus `json:"status,omitempty"`
}

// BucketLocationSpec defines the desired state of the BucketLocation Resource.
type BucketLocationSpec struct {
	LocationInfo `json:",inline"`

	// Supported types include: [ AsyncDualZone, SyncDualZone ].
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="LocationType is immutable"
	LocationType BucketLocationType `json:"locationType"`

	// +kubebuilder:default=true
	// If active = false, new buckets cannot be created with this BucketLocation.
	Active *bool `json:"active,omitempty"`
}

// Only one of its members should be specified.
type LocationInfo struct {
	// +optional
	// BucketLocation configuration specific to AsyncDualZone buckets.
	// Must be present if LocationType = AsyncDualZone
	AsyncDualZone *AsyncDualZoneSpec `json:"asyncDualZone,omitempty"`

	// +optional
	// BucketLocation configuration specific to SyncDualZone buckets.
	// Must be present if LocationType = SyncDualZone
	SyncDualZone *SyncDualZoneSpec `json:"syncDualZone,omitempty"`
}

type AsyncDualZoneSpec struct {
	// Slice of 2 location.mz.global.private.gdc.goog/v1alpha1 Zone resources.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ZoneNames is immutable"
	ZoneNames []string `json:"zoneNames"`

	// 2-character prefix used in the BucketPrefix resource
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="CombinationPrefix is immutable"
	CombinationPrefix string `json:"combinationPrefix"`
}

type SyncDualZoneSpec struct {
	// Slice of 2 location.mz.global.private.gdc.goog/v1alpha1 Zone resources.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ZoneNames is immutable"
	ZoneNames []string `json:"zoneNames"`

	// 2-character prefix used in the BucketPrefix resource
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="CombinationPrefix is immutable"
	CombinationPrefix string `json:"combinationPrefix"`
}

// Defines the observed state of the BucketLocation.
type BucketLocationStatus struct {
	v1alpha1.MuxStatus `json:",inline"`

	// The list of zone statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []BucketLocationZoneStatus `json:"zones,omitempty"`
}

// Provides the status of a BucketLocation rolling out to a particular zone.
type BucketLocationZoneStatus struct {
	v1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	ReplicaStatus BucketLocationReplicaStatus `json:"replicaStatus,omitempty"`
}

// Contains a list of BucketLocations.
// +kubebuilder:object:root=true
type BucketLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLocation `json:"items"`
}

func (spec *BucketLocationSpec) CombinationPrefix() string {
	if spec == nil {
		return ""
	}
	switch spec.LocationType {
	case AsyncDualZoneLocationType:
		if spec.AsyncDualZone == nil {
			return ""
		}
		return spec.AsyncDualZone.CombinationPrefix
	case SyncDualZoneLocationType:
		if spec.SyncDualZone == nil {
			return ""
		}
		return spec.SyncDualZone.CombinationPrefix
	default:
		return ""
	}
}

func (spec *BucketLocationSpec) ZoneNames() []string {
	if spec == nil {
		return []string{}
	}
	switch spec.LocationType {
	case AsyncDualZoneLocationType:
		if spec.AsyncDualZone == nil {
			return []string{}
		}
		return spec.AsyncDualZone.ZoneNames
	case SyncDualZoneLocationType:
		if spec.SyncDualZone == nil {
			return []string{}
		}
		return spec.SyncDualZone.ZoneNames
	}
	return []string{}
}

func init() {
	SchemeBuilder.Register(&BucketLocation{}, &BucketLocationList{})
}
