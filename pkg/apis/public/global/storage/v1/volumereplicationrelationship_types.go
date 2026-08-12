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
	globalv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Leaving this in serves as a check to ensure VolumeReplicationRelationship properly implements a MuxResource
var _ globalv1alpha1.MuxResourceInterface = &VolumeReplicationRelationship{}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName="vrr"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Source Zone",type="string",JSONPath=".spec.source.zoneRef",description="The Zone identifier containing the source volume"
// +kubebuilder:printcolumn:name="Source PVC",type="string",JSONPath=".spec.source.pvc.pvcRef",description="The PVC identifier bound to the source volume"
// +kubebuilder:printcolumn:name="Source PVC Cluster",type="string",JSONPath=".spec.source.pvc.clusterRef",description="The Cluster identifier containing the source PVC"
// +kubebuilder:printcolumn:name="Source VM Disk",type="string",JSONPath=".spec.source.virtualMachineDisk.virtualMachineDiskRef",description="The Virtual Machine Disk identifier bound to the source volume"
// +kubebuilder:printcolumn:name="Dest. Zone",type="string",JSONPath=".spec.destination.zoneRef",description="The Zone identifier to replicate the volume to"
// +kubebuilder:printcolumn:name="Dest. PVC Cluster",type="string",JSONPath=".spec.destination.pvc.clusterRef",description="The Cluster identifier to contain the destination volume"
// +kubebuilder:printcolumn:name="Dest. Volume Override",type="string",JSONPath=".spec.destination.volumeOverrideName",description="The name to use for the destination volume (default is to use the source volume name)"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="The state of the volume replication relationship"

// Defines the Schema for the 'VolumeReplicationRelationship' API.
// +genclient
type VolumeReplicationRelationship struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeReplicationRelationshipSpec   `json:"spec,omitempty"`
	Status VolumeReplicationRelationshipStatus `json:"status,omitempty"`
}

func (r *VolumeReplicationRelationship) GetSpec() any {
	return globalv1alpha1.NewTypedFieldAccessor(&r.Spec).GetField()
}

func (r *VolumeReplicationRelationship) SetSpec(spec any) error {
	return globalv1alpha1.NewTypedFieldAccessor(&r.Spec).SetField(spec)
}

func (r *VolumeReplicationRelationship) GetStatus() globalv1alpha1.MuxStatusInterface {
	return globalv1alpha1.NewMuxStatusAccessor(&r.Status).GetStatus()
}

func (r *VolumeReplicationRelationship) SetStatus(status globalv1alpha1.MuxStatusInterface) error {
	return globalv1alpha1.NewMuxStatusAccessor(&r.Status).SetStatus(status)
}

// Defines the desired state of VolumeReplicationRelationship.
type VolumeReplicationRelationshipSpec struct {
	Source      VolumeReplicationSource      `json:"source"`
	Destination VolumeReplicationDestination `json:"destination"`
}

// Defines the origin of a volume replication operation.
type VolumeReplicationSource struct {
	// A reference to the 'Zone' custom resource.
	ZoneRef string `json:"zoneRef"`

	// A reference to a `PersistentVolumeClaim` resource.
	PVC *PVCVolumeSource `json:"pvc,omitempty"`

	// A reference to a `VirtualMachineDisk` resource.
	VirtualMachineDisk *VirtualMachineDiskVolumeSource `json:"virtualMachineDisk,omitempty"`

	// A reference to a `FileShare` resource.
	FileShare *FileShareSource `json:"fileShare,omitempty"`
}

type PVCVolumeSource struct {
	// A reference to the `Cluster` custom resource where the PVC to be replicated exists.
	ClusterRef string `json:"clusterRef"`

	// A name reference to the `PersistentVolumeClaim` resource to be replicated.
	PVCRef string `json:"pvcRef"`
}

type VirtualMachineDiskVolumeSource struct {
	// A name reference to the `VirtualMachineDisk` resource to be replicated.
	VirtualMachineDiskRef string `json:"virtualMachineDiskRef"`
}

type FileShareSource struct {
	// A name reference to the `FileShare` resource to be replicated.
	FileShareRef string `json:"fileShareRef"`
}

// Defines the target location for a volume replication operation.
type VolumeReplicationDestination struct {
	// A reference to the `Zone` custom resource.
	ZoneRef string `json:"zoneRef"`

	// A reference to a `PersistentVolumeClaim` resource.
	PVC *PVCVolumeDestination `json:"pvc,omitempty"`

	// The desired name for the destination volume.
	// The default is to use the same name as the source volume.
	VolumeOverrideName string `json:"volumeOverrideName,omitempty"`
}

type PVCVolumeDestination struct {
	// A reference to the `Cluster` custom resource where the destination PVC is created.
	ClusterRef string `json:"clusterRef"`
}

// Defines the state of the volume replication used in the `VolumeReplicationRelationshipStatus`.
type VolumeReplicationState string

// Supported states for the volume replication relationship.
const (
	VolumeReplicationStatePending      VolumeReplicationState = "Pending"
	VolumeReplicationStateError        VolumeReplicationState = "Error"
	VolumeReplicationStateBrokenOff    VolumeReplicationState = "Broken Off"
	VolumeReplicationStateIdle         VolumeReplicationState = "Idle"
	VolumeReplicationStateTransferring VolumeReplicationState = "Transferring"
	VolumeReplicationStateInSync       VolumeReplicationState = "InSync"
	VolumeReplicationStateOutOfSync    VolumeReplicationState = "OutOfSync"
	VolumeReplicationStateQuiescing    VolumeReplicationState = "Quiescing"
	VolumeReplicationStateEstablished  VolumeReplicationState = "Established"
)

// Defines the observed state of the 'VolumeReplicationRelationship'.
type VolumeReplicationRelationshipStatus struct {
	globalv1alpha1.MuxStatus `json:",inline"`

	// The list of zone status' where this resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Zones []VolumeReplicationRelationshipZoneStatus `json:"zones,omitempty"`
}

// Defines the observed state of the `VolumeReplicationRelationshipReplica` of a particular zone.
type VolumeReplicationRelationshipZoneStatus struct {
	globalv1alpha1.ZoneStatus `json:",inline"`

	// The reconciliation status of the replica collected from the zone.
	ReplicaStatus VolumeReplicationRelationshipReplicaStatus `json:"replicaStatus,omitempty"`
}

func (s *VolumeReplicationRelationshipStatus) GetZones() globalv1alpha1.ListMap[string, globalv1alpha1.ZoneStatusInterface] {
	return globalv1alpha1.NewZoneStatusListAccessor(&s.Zones).GetZones()
}

func (s *VolumeReplicationRelationshipStatus) SetZones(zones globalv1alpha1.ListMap[string, globalv1alpha1.ZoneStatusInterface]) error {
	return globalv1alpha1.NewZoneStatusListAccessor(&s.Zones).SetZones(zones)
}

func (s *VolumeReplicationRelationshipZoneStatus) GetReplicaStatus() any {
	return globalv1alpha1.NewTypedFieldAccessor(&s.ReplicaStatus).GetField()
}

func (s *VolumeReplicationRelationshipZoneStatus) SetReplicaStatus(replicaStatus any) error {
	return globalv1alpha1.NewTypedFieldAccessor(&s.ReplicaStatus).SetField(replicaStatus)
}

//+kubebuilder:object:root=true

// Represents a list of 'VolumeReplicationRelationship'
type VolumeReplicationRelationshipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeReplicationRelationship `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VolumeReplicationRelationship{}, &VolumeReplicationRelationshipList{})
}
