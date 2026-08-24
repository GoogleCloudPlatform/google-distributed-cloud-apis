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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

type ObjectStorageClass string

const (
	Standard ObjectStorageClass = "Standard"
	Nearline ObjectStorageClass = "Nearline"
)

type EncryptionType string

const (
	CMEK EncryptionType = "CMEK"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".metadata.name",name="Bucket Name", type=string
// +kubebuilder:printcolumn:JSONPath=".spec.description",name="Description", type=string
// +kubebuilder:printcolumn:JSONPath=".spec.storageClass",name="Storage Class", type=string
// +kubebuilder:printcolumn:JSONPath=".spec.location",name="Location", type=string
// +kubebuilder:printcolumn:JSONPath=".status.fullyQualifiedName",name="Fully-Qualified-Bucket-Name", type=string
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".status.endpoint",description="Endpoint for the bucket"
// +kubebuilder:printcolumn:name="Region",type="string",JSONPath=".status.region",description="Region for the bucket"
// +kubebuilder:printcolumn:name="BucketReady",type="string",JSONPath=".status.conditions[?(@.type=='BucketReady')].status",description="Readiness of the Bucket"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='BucketReady')].reason",description="Reason of the Bucket Readiness"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='BucketReady')].message",description="Message from the controller"
// +kubebuilder:storageversion
// +gdcloud:manifest:relevant=true,oc=obj,component=storage,entities="buckets"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,describe,list:audit-logs-platform-restore-bucket-creator"
// +gdcloud:manifest:rbac="create,update,delete,describe,list:project-vm-image-admin,project-bucket-admin,bucket-admin,dr-system-admin"
// +gdcloud:manifest:rbac="describe,list:project-vm-image-viewer,project-bucket-object-viewer,project-bucket-object-admin,dr-system-viewer,audit-logs-platform-bucket-viewer,bucket-object-admin,bucket-object-viewer"
// +gdcloud:manifest:skipcodegen=true
// Defines the schema for the Buckets API.
// +genclient
type Bucket struct {
	// Please do not forget to update the bucket webhook accordingly if you update the Bucket API

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketSpec   `json:"spec,omitempty"`
	Status BucketStatus `json:"status,omitempty"`
}

// Contains a list of Buckets.
// +kubebuilder:object:root=true
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// BucketSpec defines the desired state of the Bucket Resource.
type BucketSpec struct {
	// Description of bucket contents.
	// +kubebuilder:validation:MaxLength:=200
	Description string `json:"description"`

	// Defines how frequently data needs to be accessed. The available options include `Standard` and `Nearline`. `Standard` is appropriate for hot data that is accessed frequently, such as websites, streaming videos, and mobile apps. It is used for data that can be stored for at least 30 days. `Nearline` is appropriate for data that can be stored for at least 60 days, including data backup and long-tail multimedia content.
	StorageClass ObjectStorageClass `json:"storageClass"`

	// +optional
	// Defines policies of the bucket resource. If unspecified, default policies are applied.
	BucketPolicy *BucketPolicy `json:"bucketPolicy,omitempty"`

	// +optional
	// Defines the physical place where object data in the bucket resides. If unspecified, defaults to the location in which the bucket is being created.
	Location string `json:"location,omitempty"`
}

// Defines the observed state of the Bucket.
type BucketStatus struct {
	// The name of the provisioned bucket. This name is used to refer to the bucket when using external tools and libraries.
	FullyQualifiedName string `json:"fullyQualifiedName,omitempty"`

	// Specifies the status of the bucket. Supported conditions include `BucketReady`. If BucketReady is `True`, it indicates the bucket has been provisioned and is ready for use.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The DNS endpoint at which the bucket is reachable.
	Endpoint string `json:"endpoint,omitempty"`

	// The region where the bucket is stored.
	Region string `json:"region,omitempty"`

	// The status of the encryption on the bucket.
	Encryption EncryptionStatus `json:"encryption,omitempty"`

	// ErrorStatus holds the most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// Defines policies of the Bucket.
type BucketPolicy struct {
	// Policy for object locking. When set, object versioning is enabled and all objects stored in the bucket will be subject to this policy. A locked object cannot be deleted until the lock expires.
	// Can only be enabled when creating the bucket and cannot be disabled afterwards.
	// When not enabled, object locking and versioning are disabled and cannot be enabled.
	// +optional
	LockingPolicy *LockingPolicy `json:"lockingPolicy,omitempty"`
	// Policy for custom CORS policy user set on the bucket.
	// CorsPolicy has to be enabled when additional CORS policy is needed on the buckets. Otherwise, the bucket will only have the default UI console CORS policy.
	// CorsPolicy can always be added or modified later after the bucket is created.
	// +optional
	CorsPolicy *CorsPolicy `json:"corsPolicy,omitempty"`
	// Policy for custom lifecycle policy user set on the bucket.
	// LifecyclePolicy can always be added, modified, removed later after the bucket is created.
	LifecyclePolicy *LifecyclePolicy `json:"lifecyclePolicy,omitempty"`
}

// EncryptionStatus defines the status of the encryption on the bucket.
type EncryptionStatus struct {
	// Defines the type of encryption to be used for the bucket.
	// Available options are:
	// - CMEK - Customer Managed Encryption Key which creates a KMS backed key rooted in the HSM which the customer is billed for.
	//			The customer can access these keys and manage them through KMS.
	Type EncryptionType `json:"type,omitempty"`

	// KeyRef references the key which is used as the default key to encrypt objects in the bucket.
	KeyRef *corev1.ObjectReference `json:"keyRef,omitempty"`
}

func (b *Bucket) FullyQualifiedName() string {
	return b.Status.FullyQualifiedName
}

func (b *Bucket) SetFullyQualifiedName(fqn string) {
	b.Status.FullyQualifiedName = fqn
}

func (b *Bucket) StorageClass() ObjectStorageClass {
	return b.Spec.StorageClass
}

func (b *Bucket) Conditions() *[]metav1.Condition {
	return &b.Status.Conditions
}

func (b *Bucket) EncryptionStatus() *EncryptionStatus {
	return &b.Status.Encryption
}

func (b *Bucket) SetEncryptionStatus(status EncryptionStatus) {
	b.Status.Encryption = status
}

func (b *Bucket) ErrorStatus() **corev1alpha1.ErrorStatus {
	return &b.Status.ErrorStatus
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}

// EncryptionKeyRef returns the EncryptionKeyRef from the bucket's status.
// This method conforms to the S3 Proxy's Bucket interface.
func (b *Bucket) EncryptionKeyRef() *corev1.ObjectReference {
	return b.Status.Encryption.KeyRef
}

// IsVersioned returns whether the bucket policy implies that the bucket is versioned.
func (b *Bucket) IsVersioned() bool {
	return b.Spec.BucketPolicy != nil && b.Spec.BucketPolicy.LockingPolicy != nil
}

func (b *Bucket) IsMZBucket() bool {
	return false
}

func (b *Bucket) GetLocation() string {
	return b.Spec.Location
}

func (b *Bucket) LockingPolicy() *LockingPolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.LockingPolicy
}

func (b *Bucket) LifecyclePolicy() *LifecyclePolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.LifecyclePolicy
}

func (b *Bucket) CorsPolicy() *CorsPolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.CorsPolicy
}

func (b *Bucket) ZonalEndpoints() []string {
	if b.Status.Endpoint == "" {
		return []string{}
	}
	return []string{b.Status.Endpoint}
}
