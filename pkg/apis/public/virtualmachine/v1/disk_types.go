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
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

const (
	// Indicates whether a `VirtualMachineDisk` is provisioned
	// and ready for consumption.
	ConditionTypeVirtualMachineDiskReady = "Ready"

	// Indicates whether a `VirtualMachineDisk` is being expanded.
	// Once disk expansion has completed, the condition will be removed.
	ConditionTypeVirtualMachineDiskExpansionInProgress = "ExpansionInProgress"

	// Indicates that a disk is currently provisioning.
	ConditionReasonDiskProvisioningPending = "DiskProvisioningPending"

	// Applies to the ExpansionInProgress condition to indicate that the expansion is pending until the VM is in Stopped state.
	ConditionReasonPendingVMStopped = "PendingVMStopped"

	// Applies to the ExpansionInProgress condition to indicate that the disk is being expanded.
	ConditionReasonExpandingDisk = "ExpandingDisk"

	// Indicates whether replication from / to a VirtualMachineDisk is ready.
	ConditionTypeVirtualMachineDiskReplicationReady = "Ready"

	// Indicates that the volume replication relationship associated with the
	// VirtualMachineDisk is pending.
	ConditionReasonVirtualMachineDiskReplicationRelationshipPending = "ReplicationRelationshipPending"

	// Indicates that the volume replication relationship associated with the
	// VirtualMachineDisk is ready.
	ConditionReasonVirtualMachineDiskReplicationRelationshipReady = "ReplicationRelationshipReady"
)

// Defines the phase of the VirtualMachineDisk.
type VirtualMachineDiskPhase string

const (
	// Disk waiting for provisioning.
	DiskPhasePending VirtualMachineDiskPhase = "Pending"
	// Disk provisioning is scheduled.
	DiskPhaseImportScheduled VirtualMachineDiskPhase = "Scheduled"
	// Disk is being provisioned.
	DiskPhaseImportInProgress VirtualMachineDiskPhase = "InProgress"
	// Disk is provisioning is complete.
	DiskPhaseSucceeded VirtualMachineDiskPhase = "Succeeded"
	// Disk is provisioning failed.
	DiskPhaseFailed VirtualMachineDiskPhase = "Failed"
	// Disk's phase is unknown.
	DiskPhaseUnknown VirtualMachineDiskPhase = "Unknown"
	// Disk provisioning is paused.
	DiskPhasePaused VirtualMachineDiskPhase = "Paused"
	// Disk waiting for replication to be initiated from the primary.
	DiskPhaseWaitingForReplicationInitiation = "WaitingForReplicationInitiation"
)

// Defines the progress of the `VirtualMachineDisk`.
type VirtualMachineDiskProgress string

// Represents the attachment relationship between the `VirtualMachine` and the
// `VirtualMachineDisk`.
type DiskAttachment struct {
	// Specifies whether this disk is the boot device for the `VirtualMachine`.
	// There must be exactly one disk marked as `boot`.
	Boot *bool `json:"boot,omitempty"`
	// Specifies whether the disk should be deleted when the `VirtualMachine` is deleted.
	// `AutoDelete` only applies while a disk is attached to a `VirtualMachine`. A
	// `VirtualMachineDisk` lifecycle is decoupled from the `VirtualMachine` once
	// it is no longer referenced in `.spec.disks`.
	AutoDelete *bool `json:"autoDelete,omitempty"`
	// Refers to a `VirtualMachineDisk` in the same `namespace`.
	VirtualMachineDiskRef corev1.LocalObjectReference `json:"virtualMachineDiskRef"`
}

// Reference to the Image Source.
type ImageDiskSource struct {
	// The name of the `VirtualMachineImage` API object.
	Name string `json:"name"`
	// The namespace of the `VirtualMachineImage` API object.
	// For golden images use `vm-system`. For images in the current project, set
	// this as empty or use the current `namespace` value.
	Namespace string `json:"namespace,omitempty"`
}

type ImageFamilySource struct {
	// The name of the image family to be used.
	// +kubebuilder:validation:Pattern="^[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?$"
	Name string `json:"name"`

	// The namespace of the image family.
	// For image families of golden images use `vm-system`.
	// For image families in the current project,
	// set this as empty or use the current `namespace` value.
	Namespace string `json:"namespace,omitempty"`
}

// Defines the source for the disk. Specify exactly one of the supported sources to use to populate a disk.
type DiskSource struct {
	// Denotes that the disk is created from a disk Image.
	Image *ImageDiskSource `json:"image,omitempty"`

	// Denotes the name of the image family this image belongs to.
	// User must provide either an image or an image family, but not both.
	ImageFamily *ImageFamilySource `json:"imageFamily,omitempty"`
}

type DiskType string

// Supported disk types.
const (
	// Standard persistent disk.
	DiskTypeStandard DiskType = "Standard"

	// Local SSD disk.
	// Supported in GDC air-gapped, unsupported in GDC connected.
	DiskTypeLocal DiskType = "Local"

	// Performance persistent disk.
	DiskTypePerformance DiskType = "Performance"
)

type ReplicationRole string

// Replication roles.
const (
	// The source disk that is being replicated.
	ReplicationRolePrimary ReplicationRole = "Primary"

	// The target disk that is being replicated to.
	ReplicationRoleSecondary ReplicationRole = "Secondary"
)

// LINT.IfChange(disk_spec)

// Defines the desired state of `VirtualMachineDisk`.
type VirtualMachineDiskSpec struct {
	// Specifies the source from which the disk contents are populated. If
	// this field is omitted a blank disk gets provisioned.
	Source *DiskSource `json:"source,omitempty"`
	// Specifies the size of the disk: 5GiB, 600MiB, and so on.
	// Size must be specified for a blank disk. For disks from
	// other sources, the size depends on the source.
	//
	// For image source, the size is optional and is inferred as being equivalent to `Image.spec.minimumDiskSize`.
	// If the size is specified it has to be greater than the `Image.spec.minimumDiskSize`.
	Size         resource.Quantity `json:"size,omitempty"`
	originalSize string            `json:"-"`
	// Specifies the type of the disk. Defaults to `Standard`.
	// In GDC air-gapped, supported types are: `Standard`, `Local`, `Performance`.
	// In GDC connected, supported types are: `Standard`.
	Type DiskType `json:"type,omitempty"`
}

// LINT.ThenChange(./disk_types_equality.go)

// Defines the observed state of VirtualMachineDisk.
type VirtualMachineDiskStatus struct {
	// The current phase of the Disk.
	Phase VirtualMachineDiskPhase `json:"phase,omitempty"`
	// Progress is the current progress of the Disk provision.
	// Value is between 0 and 100 inclusive, N/A if not available.
	Progress VirtualMachineDiskProgress `json:"progress,omitempty"`
	// The current size of the disk.
	Size resource.Quantity `json:"size,omitempty"`
	// The latest observations of the `VirtualMachineDisk` state.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// The list of `VirtualMachine` instances to which the `VirtualMachineDisk` is attached.
	VirtualMachineAttachments []VirtualMachineAttachment `json:"virtualMachineAttachments,omitempty"`
	// The status of disk replication.
	ReplicationStatus *DiskReplicationStatus `json:"replicationStatus,omitempty"`
	// The time taken to provision the `VirtualMachineDisk` and to reach a `Ready` state.
	// The time taken to provision the `VirtualMachineDisk` and to reach a `Ready` state.
	// For example, this is the time spent downloading an image, and so on.
	ProvisionTime *metav1.Duration `json:"provisionTime,omitempty"`
	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []VMMError `json:"errors,omitempty"`
}

// Contains information about which `VirtualMachine` this `VirtualMachineDisk` is
// attached to, and how it is attached.
type VirtualMachineAttachment struct {
	// The name of an attached `VirtualMachine`.
	NameRef corev1.LocalObjectReference `json:"nameRef"`
	// The UID of the attached `VirtualMachine`.
	UID types.UID `json:"uid"`
	// Reflective of how the `VirtualMachine` attaches this disk.
	// The disk is deleted when all VMs attached have `autoDelete` set to `true` and all VMs are deleted.
	// If a minimum of one VM has `autoDelete` set to `false`, the disk will not be deleted.
	AutoDelete bool `json:"autoDelete"`
}

type DiskReplicationStatus struct {
	// The role that this disk has in the replication relationship.
	// Valid values are "primary" and "secondary".
	Role ReplicationRole `json:"role,omitempty"`
	// The primary disk that this disk is being replicated from, if
	// this disk is a secondary disk. Will be nil for primary disks.
	PrimaryDisk *ReplicatedDiskInfo `json:"primaryDisk,omitempty"`
	// The secondary disk that this disk is being replicated to, if
	// this disk is a primary disk. Will be nil for secondary disks.
	SecondaryDisk *ReplicatedDiskInfo `json:"secondaryDisk,omitempty"`
	// The name of the volume replication relationship object that is managing
	// the replication.
	VolumeReplicationRelationship string `json:"volumeReplicationRelationship,omitempty"`
	// The latest observations of the replications state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// A list of any errors that occurred during replication.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

type ReplicatedDiskInfo struct {
	// The name of the `VirtualMachineDisk` instance.
	Name string `json:"name,omitempty"`
	// The zone of the `VirtualMachineDisk` instance.
	Zone string `json:"zone,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={gdisk,vmdisk,gdisks,vmdisks}
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Progress",type="string",JSONPath=".status.progress"
// Schema for the virtualmachinedisks API.
// +genclient
// +gdcloud:manifest:relevant=true,oc=vmm,component=compute,multigroup=true
// +gdcloud:manifest:entities="disks",verbs=create;update,skipcodegen=true
// +gdcloud:manifest:rbac="create,update:project-vm-admin"
type VirtualMachineDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineDiskSpec   `json:"spec,omitempty"`
	Status VirtualMachineDiskStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Contains a list of VirtualMachineDisk.
type VirtualMachineDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineDisk `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachineDisk{}, &VirtualMachineDiskList{})
}

// IsBlank returns whether the disk was created as a blank disk, i.e. without importing data from any sources.
// It doesn't reflect the current disk usage.
func (d *VirtualMachineDisk) IsBlank() bool {
	return d.Spec.Source == nil
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (spec *VirtualMachineDiskSpec) UnmarshalJSON(data []byte) error {
	// We need to unmarshal into specAlias with a different type SpecAlias
	// because if we call json.Unmarshal with a VirtualMachineDiskSpec object,
	// it will recursively call this function and cause stack overflow.
	type SpecAlias VirtualMachineDiskSpec
	var specAlias SpecAlias

	if err := json.Unmarshal(data, &specAlias); err != nil {
		return err
	}

	*spec = VirtualMachineDiskSpec(specAlias)

	if err := spec.setOriginalSize(data); err != nil {
		return err
	}

	return nil
}

func (spec *VirtualMachineDiskSpec) setOriginalSize(jsonData []byte) error {
	type SpecWithSizeOnly struct {
		Size string `json:"size,omitempty"`
	}

	specWithSizeOnly := SpecWithSizeOnly{}
	if err := json.Unmarshal(jsonData, &specWithSizeOnly); err != nil {
		return err
	}

	spec.originalSize = specWithSizeOnly.Size
	return nil
}

// MarshalJSON implements the json.Marshaller interface.
func (spec VirtualMachineDiskSpec) MarshalJSON() ([]byte, error) {
	// We need to marshal into specAlias with a different type SpecAlias
	// because if we call json.Marshal with a VirtualMachineDiskSpec object,
	// it will recursively call this function and cause stack overflow.
	type SpecAlias VirtualMachineDiskSpec
	specAlias := SpecAlias(spec)

	// Why do we marshal into bytes first and later unmarshal in setSizeToOriginalSizeInJSON to set the "size" field?
	// An appearingly more elegant solution is to define a new struct like below and marshal it:
	// specCopy := struct {
	// 	   Source *DiskSource `json:"source,omitempty"`
	// 	   Size   string      `json:"size,omitempty"`
	// 	   Type   DiskType    `json:"type,omitempty"`
	// }{
	// 	   Source: spec.Source,
	// 	   Size: spec.originalSize,
	// 	   Type: spec.Type,
	// }
	// Essentially a struct the same as VirtualMachineDiskSpec, only the "Size" field is of type "string".
	// However, this means each time we change VirtualMachineDiskSpec, we need to remember to change this struct.
	// If we add a new field to VirtualMachineDiskSpec, we need to add the same field to this struct, otherwise that field
	// will not be included in the JSON. This is error prone, so we went with the current solution.

	jsonData, err := json.Marshal(specAlias)
	if err != nil {
		return nil, err
	}

	if jsonData, err = spec.preserveOriginalSizeUnitIfPossible(jsonData); err != nil {
		return nil, fmt.Errorf("failed to preserve original size unit (if possible): %w", err)
	}
	return jsonData, nil
}

func (spec *VirtualMachineDiskSpec) preserveOriginalSizeUnitIfPossible(jsonData []byte) ([]byte, error) {
	// The VirtualMachineDisk object was newly created, not from unmarshalling
	if spec.originalSize == "" {
		return jsonData, nil
	}

	originalSize, err := resource.ParseQuantity(spec.originalSize)
	if err != nil { // This should never happen
		return nil, fmt.Errorf("failed to parse original size: %w", err)
	}

	// Disk size has changed since the unmarshalling (e.g. a client reads the disk, changes the size, and writes the disk to the API server)
	// The spec.originalSize field wouldn't be sent as part of the payload to the apiserver and is only kept locally in-memory.
	if !spec.Size.Equal(originalSize) {
		return jsonData, nil
	}

	jsonData, err = spec.setSizeToOriginalSizeInJSON(jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to set size to the original size in JSON: %w", err)
	}

	return jsonData, nil
}

func (spec *VirtualMachineDiskSpec) setSizeToOriginalSizeInJSON(jsonData []byte) ([]byte, error) {
	jsonMap := map[string]interface{}{}
	if err := json.Unmarshal(jsonData, &jsonMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal disk JSON: %w, %q", err, string(jsonData))
	}

	jsonMap["size"] = spec.originalSize

	return json.Marshal(jsonMap)
}
