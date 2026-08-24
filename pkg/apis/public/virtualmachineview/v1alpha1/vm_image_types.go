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
	"k8s.io/apimachinery/pkg/util/sets"
)

// OSName is the set of OS'es that can be used in OS.
type OSName string

// These are valid OS names for VM images in GDCH. This is a similar concept with the
// OS name in GCE https://cloud.google.com/compute/docs/images/os-details. Refer the
// naming in GCE in https://cloud.google.com/sdk/gcloud/reference/compute/images/import#--os
const (
	Ubuntu_2004   OSName = "ubuntu-2004"
	Ubuntu_2204   OSName = "ubuntu-2204"
	Ubuntu_2404   OSName = "ubuntu-2404"
	Windows_2019  OSName = "windows-2019"
	Windows_2022  OSName = "windows-2022"
	Windows_2025  OSName = "windows-2025"
	Windows_10    OSName = "windows-10"
	Windows_11    OSName = "windows-11"
	RHEL_8        OSName = "rhel-8"
	ROCKY_LINUX_8 OSName = "rocky-linux-8"
	ROCKY_LINUX_9 OSName = "rocky-linux-9"
)

var SupportedOSNames = sets.NewString(
	string(Ubuntu_2004),
	string(Ubuntu_2204),
	string(Ubuntu_2404),
	string(Windows_2019),
	string(Windows_2022),
	string(Windows_2025),
	string(Windows_10),
	string(Windows_11),
	string(RHEL_8),
	string(ROCKY_LINUX_8),
	string(ROCKY_LINUX_9),
)

// +kubebuilder:object:root=true
// +genclient
// +gdcloud:manifest:relevant=false,oc=vmm
// VirtualMachineImage represents disk image that can be used on virtual machine.
type VirtualMachineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ImageMetadata ImageMetadata `json:"imageMetadata"`
}

// ImageMetadata contains the information of the virtual machine image.
type ImageMetadata struct {
	// The name of the image, e.g. "ubuntu-20.04-server-cloudimg".
	Name string `json:"name"`

	// The version of the image, e.g. `gdch-2.0.4-xxx`.
	Version string `json:"version"`

	// URL is the url of the Docker registry source of the image.
	URL string `json:"url"`

	OperatingSystem OperatingSystemInfo `json:"operatingSystem"`

	// The minimum size of the disk the image can be applied to.
	MinimumDiskSize *resource.Quantity `json:"minimumDiskSize,omitempty"`
}

// OperatingSystemInfo contains contains the operating system information of an image.
type OperatingSystemInfo struct {
	// The name of the OS to which this image belongs, e.g. "windows-2016".
	Name OSName `json:"name"`
}

// +kubebuilder:object:root=true
// VirtualMachineImageList is a list of VirtualMachineImage objects.
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
