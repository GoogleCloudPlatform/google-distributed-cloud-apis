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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// --- Condition Types ---
const (
	// ConditionTypeVMIISourceDiskReady tracks if the source disk/file is verified, cloned, and ready.
	ConditionTypeVMIISourceDiskReady = "SourceDiskReady"

	// ConditionTypeVMIISucceeded tracks the end-to-end import, translation, and caching phases.
	ConditionTypeVMIISucceeded = "Succeeded"

	// ConditionTypeVMIITimeout tracks if any phase of the import timed out.
	ConditionTypeVMIITimeout = "Timeout"

	// ConditionTypeVMIIReady is the top-level summary condition.
	//
	// Deprecated: Ready is legacy. Use the new condition types.
	ConditionTypeVMIIReady = "Ready"
)

// --- Condition Reasons ---
const (
	// Applies to the SourceDiskReady condition.
	ReasonVMIISourceDiskAvailable                   = "SourceDiskAvailable"
	ReasonVMIIDiskPVCLookupFailed                   = "DiskPVCLookupFailed"
	ReasonVMIIVolumeSnapshotCloneVerificationFailed = "VolumeSnapshotCloneVerificationFailed"
	ReasonVMIIVolumeSnapshotMetadataFetchFailed     = "VolumeSnapshotMetadataFetchFailed"
	ReasonVMIIVolumeSnapshotCreationFailed          = "VolumeSnapshotCreationFailed"
	ReasonVMIIClonePVCCreationFailed                = "ClonePVCCreationFailed"
	ReasonVMIISourceFileImportFailed                = "SourceFileImportFailed"
	ReasonVMIISourceDiskLookupFailed                = "SourceDiskLookupFailed"
	ReasonVMIISourceFileImportInProgress            = "SourceFileImportInProgress"
	ReasonVMIISourceDiskWaiting                     = "SourceDiskWaiting"

	// Applies to the Succeeded condition.
	ReasonVMIIAllPhasesComplete         = "AllPhasesComplete"
	ReasonVMIITranslationFailed         = "TranslationFailed"
	ReasonVMIITranslationJobFailed      = "TranslationJobFailed"
	ReasonVMIIImportFailed              = "ImportFailed"
	ReasonVMIIImportJobNotCreated       = "ImportJobNotCreated"
	ReasonVMIIShadowProjectNotExists    = "ShadowProjectNotExists"
	ReasonVMIIShadowProjectSetupFailed  = "ShadowProjectSetupFailed"
	ReasonVMIIBucketNotCreated          = "BucketNotCreated"
	ReasonVMIIBucketSecretNotExists     = "BucketSecretNotExists"
	ReasonVMIIServiceAccountSetupFailed = "ServiceAccountSetupFailed"
	ReasonVMIIImageCacheFailed          = "ImageCacheFailed"
	ReasonVMIIImportPendingWait         = "ImportPendingWait"
	ReasonVMIISysprepSecretPending      = "SysprepSecretPending"
	ReasonVMIIBlankDiskPending          = "BlankDiskPending"
	ReasonVMIITranslationInProgress     = "TranslationInProgress"
	ReasonVMIITranslationJobStarted     = "TranslationJobStarted"
	ReasonVMIITranslationCleanUpPending = "TranslationCleanUpPending"
	ReasonVMIISSHCredsReady             = "SSHCredsReady"
	ReasonVMIITranslationVMPending      = "TranslationVMPending"
	ReasonVMIITranslationVMIPPending    = "TranslationVMIPPending"
	ReasonVMIITranslationJobPendingWait = "TranslationJobPendingWait"
	ReasonVMIITranslationJobPending     = "TranslationJobPending"
	ReasonVMIITranslationJobInProgress  = "TranslationJobInProgress"
	ReasonVMIITranslationJobComplete    = "TranslationJobComplete"
	ReasonVMIIImportJobPending          = "ImportJobPending"
	ReasonVMIIImportJobInProgress       = "ImportJobInProgress"
	ReasonVMIIImageCacheInProgress      = "ImageCacheInProgress"

	// Applies to timeouts.
	ReasonVMIICreationTimeout    = "CreationTimeout"
	ReasonVMIITranslationTimeout = "TranslationTimeout"
	ReasonVMIIImageUploadTimeout = "ImageUploadTimeout"
	ReasonVMIIDeletionTimeout    = "DeletionTimeout"
)

// VirtualMachineImageImportTransitionKey represents a transition key for VirtualMachineImageImport.
type VirtualMachineImageImportTransitionKey string

// VirtualMachineImageImportTransitionTime represents a transition time with a key.
type VirtualMachineImageImportTransitionTime struct {
	// Transition is the transition key.
	// +kubebuilder:validation:Required
	Transition VirtualMachineImageImportTransitionKey `json:"transition"`

	// Time is the last transition time.
	// +kubebuilder:validation:Required
	Time metav1.Time `json:"time"`
}

// Represents the operation to import and convert
// VM resources that contain data; for example, to import and convert
// 'VirtualMachineDisk` into a `VirtualMachineImage`.
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:shortName={vmimageimport, vmimageimports}
// +gdcloud:manifest:relevant=true,oc=vmm,component=compute,multigroup=true
// +gdcloud:manifest:entities="images",verbs=import,skipcodegen=true
// +gdcloud:manifest:rbac="import:vmm-admin,project-vm-image-admin"
type VirtualMachineImageImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineImageImportSpec   `json:"spec,omitempty"`
	Status VirtualMachineImageImportStatus `json:"status,omitempty"`
}

// Specifies the source and metadata for the image
// that you want to create.
type VirtualMachineImageImportSpec struct {
	// Refers to the resource from which contents are imported.
	// The source must be from the same namespace.
	Source ImageSourceReference `json:"source"`

	// Refers to the resource to which contents are imported.
	// Only supported in GDC Connected deployments.
	// This field is required if the specified `ImageSourceReference` is `GCS`.
	Destination ImageDestinationReference `json:"destination,omitempty"`

	// Specifies the properties of the `VirtualMachineImage` you want to create.
	ImageMetadata ImageMetadataInput `json:"imageMetadata"`

	// Specifies whether to prepare this image for a GDC air-gapped deployment
	// with a value to, for example, install the required packages.
	// If this is not specified, preparation occurs only if the image is from object storage.
	PrepareImage *bool `json:"prepareImage,omitempty"`

	// These are the options for image preparation. This is only valid when `prepareImage` is `true`.
	PrepareOptions ImagePrepareOptions `json:"prepareOptions,omitempty"`
}

// The specification for the `VirtualMachineImage`.
type ImageMetadataInput struct {
	// The image name, such as `ubuntu-20.04-server-cloudimg`.
	Name string `json:"name"`

	// The name of the OS to which this image belongs, such as `ubuntu-2004`.
	OperatingSystem OSName `json:"operatingSystem"`

	// The minimum size of the disk to which the image can be applied.
	// This specifies only the recommended size for future disks that are
	// created from this image. It does not represent the size of the image itself.
	// This field is required if the image is being imported from object storage.
	MinimumDiskSize *resource.Quantity `json:"minimumDiskSize,omitempty"`

	// Refers to the specific family that a particular image belongs to.
	// +kubebuilder:validation:Optional.
	// +kubebuilder:validation:Pattern="^[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?$"
	Family string `json:"family,omitempty"`
}

// Provides the image import status.
type VirtualMachineImageImportStatus struct {
	// Provide the `Ready` status of the import progress.
	Conditions []metav1.Condition `json:"conditions"`

	// Refers to the `VirtualMachineImage` once successfully created.
	// The image is in the same namespace as the image import.
	ImageName string `json:"imageName,omitempty"`

	// Last transition time of each transition.
	// +optional
	// +listType=map
	// +listMapKey=transition
	TransitionTime []VirtualMachineImageImportTransitionTime `json:"transitionTime,omitempty"`

	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []VMMError `json:"errors,omitempty"`
}

// Points to the resource from which the image contents are populated.
// Only one source can be specified.
type ImageSourceReference struct {
	// Specified when the image is populated from an existing `VirtualMachineDisk`.
	DiskRef *corev1.LocalObjectReference `json:"diskRef,omitempty"`

	// The bucket details for an image that is populated from object storage.
	ObjectStorage *ImageObjectStorageSourceReference `json:"objectStorage,omitempty"`

	// The bucket details for an image that is populated from gcs.
	GCS *ImageGCSSourceReference `json:"gcs,omitempty"`
}

// Points to the destination to which the image contents are uploaded.
// Only one destination can be specified.
// Only supported in GDC Connected deployments.
// This field is required if the specified `ImageSourceReference` is `GCS`.
type ImageDestinationReference struct {
	// The bucket details for an image that is populated to gcs.
	GCS *ImageGCSDestinationReference `json:"gcs,omitempty"`
}

// Represents the options for image preparation.
type ImagePrepareOptions struct {
	// Specifies whether to install the GDC air-gapped guest environment.
	// Defaults to `true`.
	InstallGuestEnvironment *bool `json:"installGuestEnvironment,omitempty"`
}

// Represents the object storage source from which to import an image.
type ImageObjectStorageSourceReference struct {
	// The name of the `Bucket` custom resource that holds this image.
	// The `Bucket` custom resource must be in the same namespace as this object.
	BucketRef corev1.LocalObjectReference `json:"bucketRef"`

	// The name of the image within the bucket.
	ObjectName string `json:"objectName"`
}

// Represents the GCS source from which to import an image.
type ImageGCSSourceReference struct {
	// The GCS Bucket name that holds this image.
	BucketName string `json:"bucketName"`

	// The name of the image within the bucket.
	ObjectName string `json:"objectName"`
}

// Represents the GCS destination to which to import an image.
type ImageGCSDestinationReference struct {
	// The GCS Bucket name that holds this image.
	BucketName string `json:"bucketName"`
}

// A list of `VirtualMachineImageImport` objects.
// +kubebuilder:object:root=true
type VirtualMachineImageImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VirtualMachineImageImport `json:"items"`
}

// GetTransitionTime returns the transition time for the specified key if present.
func (s *VirtualMachineImageImportStatus) GetTransitionTime(key VirtualMachineImageImportTransitionKey) (metav1.Time, bool) {
	for _, t := range s.TransitionTime {
		if t.Transition == key {
			return t.Time, true
		}
	}
	return metav1.Time{}, false
}

// SetTransitionTime sets or updates the transition time for the specified key.
func (s *VirtualMachineImageImportStatus) SetTransitionTime(key VirtualMachineImageImportTransitionKey, t metav1.Time) {
	for i, tt := range s.TransitionTime {
		if tt.Transition == key {
			s.TransitionTime[i].Time = t
			return
		}
	}
	s.TransitionTime = append(s.TransitionTime, VirtualMachineImageImportTransitionTime{
		Transition: key,
		Time:       t,
	})
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachineImageImport{},
		&VirtualMachineImageImportList{},
	)
}
