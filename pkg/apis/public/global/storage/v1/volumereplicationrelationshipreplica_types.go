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
	"github.com/go-openapi/strfmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	globalv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
)

// Leaving this in serves as a check to ensure VolumeReplicationRelationshipReplica properly implements a MuxReplica
var _ globalv1alpha1.MuxReplicaInterface = &VolumeReplicationRelationshipReplica{}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName="vrrr"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Source Zone",type="string",JSONPath=".spec.source.zoneRef",description="The Zone identifier containing the source volume"
// +kubebuilder:printcolumn:name="Source Cluster",type="string",JSONPath=".spec.source.clusterRef",description="The Cluster identifier containing the source volume"
// +kubebuilder:printcolumn:name="Source PVC",type="string",JSONPath=".spec.source.pvcRef",description="The PVC identifier bound to the source volume"
// +kubebuilder:printcolumn:name="Dest. Zone",type="string",JSONPath=".spec.destination.zoneRef",description="The Zone identifier to replicate the volume to"
// +kubebuilder:printcolumn:name="Dest. Cluster",type="string",JSONPath=".spec.destination.clusterRef",description="The Cluster identifier to contain the destination volume"
// +kubebuilder:printcolumn:name="Dest. PVC",type="string",JSONPath=".spec.destination.pvcOverrideName",description="The name to use for the destination PVC (default is to use the source PVC name)"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="The state of the volume replication relationship"

// Defines the Schema for the `VolumeReplicationRelationshipReplica` API.
// +genclient
type VolumeReplicationRelationshipReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeReplicationRelationshipSpec          `json:"spec,omitempty"`
	Status VolumeReplicationRelationshipReplicaStatus `json:"status,omitempty"`
}

func (r *VolumeReplicationRelationshipReplica) GetSpec() any {
	return globalv1alpha1.NewTypedFieldAccessor(&r.Spec).GetField()
}

func (r *VolumeReplicationRelationshipReplica) SetSpec(spec any) error {
	return globalv1alpha1.NewTypedFieldAccessor(&r.Spec).SetField(spec)
}

func (r *VolumeReplicationRelationshipReplica) GetStatus() any {
	return globalv1alpha1.NewTypedFieldAccessor(&r.Status).GetField()
}

func (r *VolumeReplicationRelationshipReplica) SetStatus(status any) error {
	return globalv1alpha1.NewTypedFieldAccessor(&r.Status).SetField(status)
}

// Defines the observed state of the `VolumeReplicationRelationshipReplica`.
type VolumeReplicationRelationshipReplicaStatus struct {
	// The current status of the volume replication relationship.
	// +optional
	State *VolumeReplicationState `json:"state,omitempty"`

	// A human-readable message indicating details about why the volume replication relationship is in this state.
	// +optional
	Message *string `json:"message,omitempty"`

	// The unique UUID identifying a SnapMirror relationship
	// +optional
	ReplicationID *strfmt.UUID `json:"replicationID,omitempty"`

	// The cached information about the source volume. This information is only populated by the source zone.
	// +optional
	SourceVolume *LocalVolume `json:"sourceVolume,omitempty"`

	// The cached information about the destination volume. This information is only populated by the destination zone.
	// +optional
	DestinationVolume *LocalVolume `json:"destinationVolume,omitempty"`

	// The timestamp of when the last snapshot was exported
	// +optional
	LagTime *string `json:"lagTime,omitempty"`

	// The name of the last exported snapshot
	// +optional
	ExportedSnapshotName *string `json:"exportedSnapshotName,omitempty"`

	// The type of the last transfer operation, such as a snapshot.
	// +optional
	LastTransferType *string `json:"lastTransferType,omitempty"`

	// The timestamp of when the last transfer operation occurred
	// +optional
	LastTransferTime *string `json:"lastTransferTime,omitempty"`

	// The information on the overall state of the `VolumeReplicationRelationshipReplica`.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Represents a volume within the local storage system.
type LocalVolume struct {
	// The name of the storage class defined by the `PersistentVolumeClaim`.
	StorageClassName string `json:"storageClassName"`

	// The storage cluster on which the Volume was allocated, as used inside
	// of the peering between zones.
	StorageClusterName string `json:"storageClusterName"`

	// The name of the internal volume on the storage appliance.
	InternalVolumeName string `json:"internalVolumeName"`

	// The storage size of the volume requested by the `PersistentVolumeClaim`.
	RequestedStorageSize string `json:"requestedStorageSize"`

	// The mode of the internal volume on the storage appliance.
	VolumeMode *corev1.PersistentVolumeMode `json:"volumeMode,omitempty"`

	// Specifies whether PVC has been deleted
	IsDeleted bool `json:"isDeleted"`
}

// +kubebuilder:object:root=true

// Contains a list of `VolumeReplicationRelationshipReplica` resources.
type VolumeReplicationRelationshipReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeReplicationRelationshipReplica `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeReplicationRelationshipReplica{}, &VolumeReplicationRelationshipReplicaList{})
}
