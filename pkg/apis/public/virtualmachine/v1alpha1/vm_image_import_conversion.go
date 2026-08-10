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
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/virtualmachine/v1"
	vmviewv1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/virtualmachineview/v1alpha1"
)

// ConvertTo converts this VirtualMachineImageImport to the Hub version (v1).
func (v1a1 *VirtualMachineImageImport) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.VirtualMachineImageImport)

	dst.ObjectMeta = v1a1.ObjectMeta

	if v1a1.Spec.Source.Disk != nil {
		dst.Spec.Source.DiskRef = &corev1.LocalObjectReference{
			Name: v1a1.Spec.Source.Disk.Name,
		}
	}
	if v1a1.Spec.Source.ObjectStorage != nil {
		dst.Spec.Source.ObjectStorage = &v1.ImageObjectStorageSourceReference{
			BucketRef: corev1.LocalObjectReference{
				Name: v1a1.Spec.Source.ObjectStorage.Bucket,
			},
			ObjectName: v1a1.Spec.Source.ObjectStorage.ObjectName,
		}
	}

	dst.Spec.ImageMetadata.Name = fmt.Sprintf("%s-%s", v1a1.Spec.ImageMetadata.Name, v1a1.Spec.ImageMetadata.Version)
	dst.Spec.ImageMetadata.OperatingSystem = v1.OSName(v1a1.Spec.ImageMetadata.OperatingSystem.Name)
	dst.Spec.ImageMetadata.MinimumDiskSize = v1a1.Spec.ImageMetadata.MinimumDiskSize

	dst.Spec.PrepareImage = v1a1.Spec.PrepareImage
	dst.Spec.PrepareOptions.InstallGuestEnvironment = v1a1.Spec.PrepareOptions.InstallGuestEnvironment

	dst.Status.Conditions = v1a1.Status.Conditions

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (v1a1 *VirtualMachineImageImport) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.VirtualMachineImageImport)

	v1a1.ObjectMeta = src.ObjectMeta

	if src.Spec.Source.DiskRef != nil {
		v1a1.Spec.Source.Disk = &ImageDiskSourceReference{
			Name: src.Spec.Source.DiskRef.Name,
		}
	}
	if src.Spec.Source.ObjectStorage != nil {
		v1a1.Spec.Source.ObjectStorage = &ImageObjectStorageSourceReference{
			Bucket:     src.Spec.Source.ObjectStorage.BucketRef.Name,
			ObjectName: src.Spec.Source.ObjectStorage.ObjectName,
		}
	}

	name, ver := splitVersionAndName(src.Spec.ImageMetadata.Name)
	v1a1.Spec.ImageMetadata.Name = name
	v1a1.Spec.ImageMetadata.Version = ver
	v1a1.Spec.ImageMetadata.OperatingSystem.Name = vmviewv1alpha1.OSName(src.Spec.ImageMetadata.OperatingSystem)
	v1a1.Spec.ImageMetadata.MinimumDiskSize = src.Spec.ImageMetadata.MinimumDiskSize

	v1a1.Spec.PrepareImage = src.Spec.PrepareImage
	v1a1.Spec.PrepareOptions.InstallGuestEnvironment = src.Spec.PrepareOptions.InstallGuestEnvironment

	v1a1.Status.Conditions = src.Status.Conditions

	return nil
}

// Because version doesn't exist as a separate field in VMII v1, we append it to the image name.
// For converting back, we need to extract this version from the last hyphen. There are two failure modes:
// 1. If the version contains a hyphen, this only return the last part of the version (the rest will be in the name)
// This is okay because the imported image name will still be the same
// 2. If this is called on a v1 name without a hyphen (as they aren't required in v1), the version in v1alpha1 will be set to 1.0.
// This is because versions are necessary in v1alpha1, but it does mean that the resultant image name will be different.
// This case should be rare enough that it isn't an issue. Plausibly, v1alpha1 objects will never be converted from v1.
func splitVersionAndName(name string) (string, string) {
	loc := strings.LastIndex(name, "-")

	if loc == -1 || loc == len(name)-1 {
		return name, "1.0"
	}

	return name[:loc], name[loc+1:]
}
