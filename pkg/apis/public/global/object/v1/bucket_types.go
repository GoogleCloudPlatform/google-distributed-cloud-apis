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

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
	v1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/object/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".metadata.name",name="Bucket Name", type=string
// +kubebuilder:printcolumn:JSONPath=".spec.description",name="Description", type=string
// +kubebuilder:printcolumn:JSONPath=".spec.storageClass",name="Storage Class", type=string
// +kubebuilder:printcolumn:JSONPath=".spec.location",name="Location", type=string
// +kubebuilder:printcolumn:JSONPath=".status.fullyQualifiedName",name="Fully Qualified Bucket Name", type=string
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".status.globalEndpoint",description="Global endpoint for the bucket"
// +kubebuilder:printcolumn:name="BucketReady",type="string",JSONPath=".status.conditions[?(@.type=='BucketReady')].status",description="Readiness of the Bucket"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='BucketReady')].reason",description="Reason of the Bucket Readiness"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='BucketReady')].message",description="Message from the controller"
// +kubebuilder:storageversion
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

// Defines the desired state of the Bucket Resource.
type BucketSpec struct {
	// The description of bucket contents.
	// +kubebuilder:validation:MaxLength:=200
	Description string `json:"description"`

	// The storage class of the bucket. Defines how frequently data needs to be accessed. The available options include `Standard` and `Nearline`. `Standard` is appropriate for hot data that is accessed frequently, such as websites, streaming videos, and mobile apps. It is used for data that can be stored for at least 30 days. `Nearline` is appropriate for data that can be stored for at least 60 days, including data backup and long-tail multimedia content.
	StorageClass v1.ObjectStorageClass `json:"storageClass"`

	// +optional
	// The policies of the bucket resource. If unspecified, default policies are applied.
	BucketPolicy *v1.GlobalBucketPolicy `json:"bucketPolicy,omitempty"`

	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location is immutable"
	// The physical place where object data in the bucket resides. This must map to the name of an existing `ObjectStorageLocation` resource.
	Location string `json:"location,omitempty"`
}

// Defines the observed state of the Bucket.
type BucketStatus struct {
	// The name of the provisioned bucket. This name is used to refer to the bucket when using external tools and libraries.
	FullyQualifiedName string `json:"fullyQualifiedName,omitempty"`

	// The conditions of the bucket. Supported conditions include `BucketReady`. If BucketReady is `True`, it indicates the bucket has been provisioned and is ready for use.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Zonal endpoints at which the bucket is reachable. Use these endpoints if a customized failover is required.
	ZonalEndpoints []string `json:"zonalEndpoints,omitempty"`

	// Global endpoint which will dynamically route traffic to any zone which contains data for this bucket. Use this endpoint if automatic failover is required.
	GlobalEndpoint string `json:"globalEndpoint,omitempty"`

	// The region where the bucket is stored.
	Region string `json:"region,omitempty"`

	// The status of the encryption on the bucket.
	Encryption v1.EncryptionStatus `json:"encryption,omitempty"`

	// ErrorStatus holds the most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func (b *Bucket) FullyQualifiedName() string {
	return b.Status.FullyQualifiedName
}

func (b *Bucket) SetFullyQualifiedName(fqn string) {
	b.Status.FullyQualifiedName = fqn
}

func (b *Bucket) StorageClass() v1.ObjectStorageClass {
	return b.Spec.StorageClass
}

func (b *Bucket) Conditions() *[]metav1.Condition {
	return &b.Status.Conditions
}

func (b *Bucket) EncryptionStatus() *v1.EncryptionStatus {
	return &b.Status.Encryption
}

func (b *Bucket) SetEncryptionStatus(status v1.EncryptionStatus) {
	b.Status.Encryption = status
}

func (b *Bucket) ErrorStatus() **corev1alpha1.ErrorStatus {
	return &b.Status.ErrorStatus
}

func (b *Bucket) LockingPolicy() *v1.LockingPolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.LockingPolicy
}

func (b *Bucket) LifecyclePolicy() *v1.LifecyclePolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.LifecyclePolicy
}

func (b *Bucket) CorsPolicy() *v1.CorsPolicy {
	if b.Spec.BucketPolicy == nil {
		return nil
	}
	return b.Spec.BucketPolicy.CorsPolicy
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
