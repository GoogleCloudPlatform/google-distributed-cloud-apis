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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

// The list of valid OSes for use.
type OSName string

// These are valid OS names for VM images in GDC. This is a similar concept with
// the OS name in GCE https://cloud.google.com/compute/docs/images/os-details.
// Refer the naming in GCE in
// https://cloud.google.com/sdk/gcloud/reference/compute/images/import#--os

const (
	Ubuntu2004  OSName = "ubuntu-2004"
	Ubuntu2204  OSName = "ubuntu-2204"
	Ubuntu2404  OSName = "ubuntu-2404"
	Windows2019 OSName = "windows-2019"
	Windows2022 OSName = "windows-2022"
	Windows2025 OSName = "windows-2025"
	Windows10   OSName = "windows-10"
	Windows11   OSName = "windows-11"
	RHEL8       OSName = "rhel-8"
	RockyLinux8 OSName = "rocky-linux-8"
	RockyLinux9 OSName = "rocky-linux-9"
	GardenLinux OSName = "garden-linux"
	SUSECHost15 OSName = "suse-chost-15"
	ContainerOS OSName = "container-os"
	Unknown     OSName = "unknown"
)

var SupportedOSNames = sets.NewString(
	string(Ubuntu2004),
	string(Ubuntu2204),
	string(Ubuntu2404),
	string(Windows2019),
	string(Windows2022),
	string(Windows2025),
	string(Windows10),
	string(Windows11),
	string(RHEL8),
	string(RockyLinux8),
	string(GardenLinux),
	string(SUSECHost15),
	string(ContainerOS),
	string(Unknown),
)

// TODO(b/535433346): Refactor OS names classification to use attributes and filters.

// WindowsServerOSNames are the valid OS names for Windows Server images.
var WindowsServerOSNames = sets.NewString(
	string(Windows2019),
	string(Windows2022),
	string(Windows2025),
)

// WindowsClientOSNames are the valid OS names for Windows Client images.
var WindowsClientOSNames = sets.NewString(
	string(Windows10),
	string(Windows11),
)

// ImportableWindowsOSNames are the valid Windows OS names for BYO image import.
var ImportableWindowsOSNames = WindowsServerOSNames.Union(WindowsClientOSNames)

// ImportableOSNames are the valid OS names for BYO image import.
var ImportableOSNames = sets.NewString(
	string(Ubuntu2004),
	string(Ubuntu2204),
	string(Ubuntu2404),
	string(RHEL8),
	string(GardenLinux),
	string(SUSECHost15),
).Union(ImportableWindowsOSNames)

const (
	// ConditionTypeImageCacheReady indicates whether the image has been cached.
	ConditionTypeImageCacheReady = "ImageCacheReady"
)

// Represents the disk image that can be used on virtual machine.
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={vmimage, vmimages}
// +gdcloud:manifest:relevant=true,oc=vmm,component=compute,multigroup=true
// +gdcloud:manifest:entities="images",verbs=delete;update
// +gdcloud:manifest:rbac="describe,list,update,delete:vmm-admin"
// +gdcloud:manifest:rbac="describe,list:vmm-viewer"
// +gdcloud:manifest:skipcodegen=true
type VirtualMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineImageSpec   `json:"spec,omitempty"`
	Status            VirtualMachineImageStatus `json:"status,omitempty"`
}

// Defines the specification of the virtual machine image.
type VirtualMachineImageSpec struct {
	// The details of the OS.
	// +kubebuilder:validation:Required.
	OperatingSystem OperatingSystemSpec `json:"operatingSystem"`

	// The minimum size of the disk the image can be applied to.
	// This specifies only the recommended size for the future disks that are
	// created from this image. It does not represent the size of the Image itself.
	// +kubebuilder:validation:Optional.
	MinimumDiskSize *resource.Quantity `json:"minimumDiskSize,omitempty"`

	// Refers to the GCS resource from which image are stored.
	// +kubebuilder:validation:Optional.
	GCS GCSReference `json:"gcs,omitempty"`

	// Refers to the specific family that a particular image belongs to.
	// +kubebuilder:validation:Optional.
	// +kubebuilder:validation:Pattern="^[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?$"
	Family string `json:"family,omitempty"`

	// The timestamp when the virtual machine image build process was initiated.
	BuildTime *metav1.Time `json:"buildTime,omitempty"`
}

// Contains the observed state of the `VirtualMachineImage` object.
type VirtualMachineImageStatus struct {
	// The conditions of the virtual machine image.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []VMMError `json:"errors,omitempty"`

	// The zone where this virtual machine image is stored.
	// +optional
	StorageLocation string `json:"storageLocation,omitempty"`
}

// Contains the operating system information of an image.
type OperatingSystemSpec struct {
	// The name of the OS to which this image belongs, e.g. "windows-2016".
	Name OSName `json:"name"`
}

// Represents the GCS source from which to store the image.
type GCSReference struct {
	// The URL of the GCS object.
	URL string `json:"url"`
}

// Contains a list of `VirtualMachineImage` objects.
// +kubebuilder:object:root=true
type VirtualMachineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VirtualMachineImage `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineImage{},
		&VirtualMachineImageList{},
	)
}
