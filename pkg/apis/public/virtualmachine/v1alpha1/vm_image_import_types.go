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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	vmviewv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/virtualmachineview/v1alpha1"
)

// VirtualMachineImageImport represents the operation to import and convert
// VM resources that contain data e.g., VirtualMachineDisk into a VirtualMachineImage.
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=vmm
type VirtualMachineImageImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineImageImportSpec   `json:"spec,omitempty"`
	Status VirtualMachineImageImportStatus `json:"status,omitempty"`
}

// TODO(b/277386522) Remove "under development" comments once VMM_BYO_IMG feature gate is on in production.

// VirtualMachineImageImportSpec specifies the source and metadata for the image
// to be created.
type VirtualMachineImageImportSpec struct {
	// Source refers to the resource that is used to import the contents from.
	// Currently, the source has to be from the same namespace.
	// +kubebuilder:validation:Required.
	Source ImageSourceReference `json:"source"`
	// ImageMetadata specifies the properties of the to be created VirtualMachineImage.
	// +kubebuilder:validation:Required.
	ImageMetadata ImageMetadataInput `json:"imageMetadata"`
	// Specifies whether or not to prepare this image for GDCH (e.g. install required packages)
	// Defaults to true for the `objectStorage` source and false for the `disk` source.
	// Note: this functionality is under development and may not be available.
	// +kubebuilder:validation:Optional.
	PrepareImage *bool `json:"prepareImage,omitempty"`
	// The options for image preparation. Only valid when `prepareImage` is true.
	// Note: this functionality is under development and may not be available.
	// +kubebuilder:validation:Optional.
	PrepareOptions ImagePrepareOptions `json:"prepareOptions,omitempty"`
}

// ImageMetadataInput is the specification for the VirtualMachineImage.
type ImageMetadataInput struct {
	// Name is name of the image, e.g. "ubuntu-20.04-server-cloudimg".
	// +kubebuilder:validation:Required.
	Name string `json:"name"`

	// Version is version of the image, e.g. `gdch-2.0.4-xxx`.
	// This allows to have multiple versions of image with name.
	// +kubebuilder:validation:Required.
	Version string `json:"version"`

	// OperatingSystem is details of the OS.
	// +kubebuilder:validation:Required.
	OperatingSystem vmviewv1alpha1.OperatingSystemInfo `json:"operatingSystem"`

	// MinimumDiskSize The minimum size of the disk the image can be applied to.
	// This specifies only the recommended size for the future disks that are
	// created from this image. It does not represent the size of the Image itself.
	// This field is required if the image is being imported from object storage.
	// +kubebuilder:validation:Optional.
	MinimumDiskSize *resource.Quantity `json:"minimumDiskSize,omitempty"`
}

// VirtualMachineImageImportStatus provides status of the image import.
type VirtualMachineImageImportStatus struct {
	// Conditions provide the 'Ready' status of the import progress.
	Conditions []metav1.Condition `json:"conditions"`
	// ImageName refers to the VirtualMachineImage once successfully created.
	// The image will be in the same namespace as the import.
	ImageName string `json:"imageName,omitempty"`
}

// ImageSourceReference points to the resource with which the image contents
// should be populated from. Only one source must be specified.
type ImageSourceReference struct {
	// Disk is specified when the image is populated from an existing VirtualMachineDisk.
	Disk *ImageDiskSourceReference `json:"disk,omitempty"`

	// The bucket details for an image that is populated from object storage.
	// Note: this functionality is under development and may not be available.
	ObjectStorage *ImageObjectStorageSourceReference `json:"objectStorage,omitempty"`
}

// ImageDiskSourceReference references a VirtualMachineDisk.
type ImageDiskSourceReference struct {
	// Name of the disk.
	// +kubebuilder:validation:Required.
	Name string `json:"name"`
}

// Represents the options for image preparation.
type ImagePrepareOptions struct {
	// Specifies whether or not to install the GDCH guest environment.
	// Defaults to true.
	InstallGuestEnvironment *bool `json:"installGuestEnvironment,omitempty"`
}

// Represents the object storage source to import an image from.
type ImageObjectStorageSourceReference struct {
	// The name of the `Bucket` CR that holds this image.
	// The `Bucket` CR must be in the same namespace as this object.
	// +kubebuilder:validation:Required.
	Bucket string `json:"bucket"`
	// The name of the image within the bucket.
	// +kubebuilder:validation:Required.
	ObjectName string `json:"objectName"`
}

// VirtualMachineImageImportList is a list of VirtualMachineImageImport objects.
// +kubebuilder:object:root=true
type VirtualMachineImageImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VirtualMachineImageImport `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineImageImport{},
		&VirtualMachineImageImportList{},
	)
}
