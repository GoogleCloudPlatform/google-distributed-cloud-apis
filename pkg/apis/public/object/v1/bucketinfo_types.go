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
)

// +gdcloud:manifest:relevant=false,oc=obj
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=bucketinfos,singular=bucketinfo
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".metadata.name",name="Bucket FQDN",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.storageClass",name="Storage Class",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.location",name="Location",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.bucketName",name="Bucket Name",type=string
// +kubebuilder:printcolumn:JSONPath=".spec.region",name="Region",type="string",description="Region for the bucket"
// +kubebuilder:storageversion
// Defines the schema for the BucketInfo API.
// +genclient
type BucketInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketInfoSpec   `json:"spec,omitempty"`
	Status BucketInfoStatus `json:"status,omitempty"`
}

// BucketInfoSpec defines the desired state of the BucketInfo Resource.
type BucketInfoSpec struct {
	// +optional
	// The description of bucket contents.
	// +kubebuilder:validation:MaxLength:=200
	Description string `json:"description"`

	// Defines how frequently data needs to be accessed. The available options include `Standard` and `Nearline`. `Standard` is appropriate for hot data that is accessed frequently, such as websites, streaming videos, and mobile apps. It is used for data that can be stored for at least 30 days. `Nearline` is appropriate for data that can be stored for at least 60 days, including data backup and long-tail multimedia content.
	StorageClass ObjectStorageClass `json:"storageClass"`

	// +optional
	// Defines policies of the bucket resource. If unspecified, default policies are applied.
	BucketPolicy *GlobalBucketPolicy `json:"bucketPolicy,omitempty"`

	// +optional
	// Defines the physical place where object data in the bucket resides. If unspecified, defaults to the location in which the bucket is being created.
	Location string `json:"location,omitempty"`

	// The non-namespaced name of the provisioned bucket. This is used to refer to the bucket
	// when using local tools and libraries.
	BucketName string `json:"bucketName,omitempty"`

	// Zonal DNS endpoints at which the bucket is reachable. Use these endpoints
	// if customized failover is required.
	ZonalEndpoints []string `json:"zonalEndpoints,omitempty"`

	// Global endpoint which will dynamically route traffic to any zone which contains data for this bucket. Use this endpoint if automatic failover is required.
	GlobalEndpoint string `json:"globalEndpoint,omitempty"`

	// The region where the bucket is stored.
	Region string `json:"region,omitempty"`

	// The status of the encryption on the bucket.
	Encryption EncryptionStatus `json:"encryption,omitempty"`

	// +optional
	// Only used for synchronous buckets. Determines if S3 operations should
	// revert to asynchronous replication due to one of the replication zones
	// being unavailable. This prevents synchronous buckets from becoming read-only
	// in the event that one of the replication zones is down. If empty, defaults
	// to false.
	AllowDegradedWrites *bool `json:"allowDegradedWrites,omitempty"`
}

// Defines the observed state of the BucketInfo.
type BucketInfoStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Defines policies of the Bucket.
type GlobalBucketPolicy struct {
	// Policy for object locking. When set, object versioning is enabled and all objects stored in the bucket will be subject to this policy. A locked object cannot be deleted until the lock expires.
	// Can only be enabled when creating the bucket and cannot be disabled afterwards.
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

// +kubebuilder:object:root=true
// Contains a list of BucketInfos.
type BucketInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketInfo `json:"items"`
}

func (b *BucketInfo) EncryptionStatus() *EncryptionStatus {
	return &b.Spec.Encryption
}

// FQDN returns the fully qualified bucket domain name.
func (b *BucketInfo) FullyQualifiedName() string {
	return b.Name
}

// IsVersioned returns whether the bucket policy implies that the bucket is versioned.
func (b *BucketInfo) IsVersioned() bool {
	return true
}

// StorageClass returns the bucket's storage class.
func (b *BucketInfo) StorageClass() ObjectStorageClass {
	return b.Spec.StorageClass
}

func (b *BucketInfo) IsMZBucket() bool {
	return true
}

func (b *BucketInfo) GetLocation() string {
	return b.Spec.Location
}

func (b *BucketInfo) LockingPolicy() *LockingPolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.LockingPolicy
}

func (b *BucketInfo) LifecyclePolicy() *LifecyclePolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.LifecyclePolicy
}

func (b *BucketInfo) CorsPolicy() *CorsPolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.CorsPolicy
}

func (b *BucketInfo) Conditions() *[]metav1.Condition {
	return &b.Status.Conditions
}

func (b *BucketInfo) ZonalEndpoints() []string {
	if b.Spec.ZonalEndpoints == nil {
		return []string{}
	}
	return b.Spec.ZonalEndpoints
}

func init() {
	SchemeBuilder.Register(&BucketInfo{}, &BucketInfoList{})
}
