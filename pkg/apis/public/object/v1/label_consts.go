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

const (
	// StorageClassLabel is set by controllers to indicate the storage class
	// that a resource corresponds to, e.g. on Secrets containing object storage
	// API access keys.
	StorageClassLabel = "object.gdc.goog/storage-class"

	// SubjectTypeLabel is set by controllers on Secrets to identify the subject
	// type that a resource corresponds to, e.g. "ServiceAccount", "User", etc.
	SubjectTypeLabel = "object.gdc.goog/subject-type"

	// TenantCategoryLabel is populated by OPA mutation webhook while creating
	// the project. The namespace created by the project also comes with the label.
	// TenantCategoryLabel indicates the object storage tenant category for the
	// project/namespace, e.g. "system" or "user".
	TenantCategoryLabel = "object.gdc.goog/tenant-category"

	// SecretTypeLabelis set by controllers on secrets to identify what type of encryption
	// key it belongs to, e.g. "BEK"
	SecretTypeLabel = "object.gdc.goog/secret-type"

	// EncryptionVersionLabel is added by the defaulter to buckets which need encryption
	// enabled, and the version is always needed before proxy initiates the encryption, e.g. "v1".
	// The label cannot be midified by the user.
	EncryptionVersionLabel = "object.gdc.goog/encryption-version"

	// ConvertedFromLabel is added by the convertor to bucket to indicate the intended
	// version of the bucket to create. If this label is not present on the bucket CR, it
	// means the bucket was created in current supported version, e.g "v1", "v1alpha1".
	ConvertedFromLabel = "object.gdc.goog/converted-from"

	// BucketNameLabel is set by controllers on KEKRef secrets to identify the bucket that the
	// secret it belongs to.
	BucketNameLabel = "object.gdc.goog/bucket-name"

	// BucketProjectLabel is set by controllers on KEKRef secrets to identify the project of the
	// bucket that the secret it belongs to.
	BucketProjectLabel = "object.gdc.goog/bucket-project"

	// BucketTypeLabel is set by the bucket controller on buckets to identify the type of project
	// that the bucket belongs to, e.g. "normal", "shadow".
	BucketTypeLabel = "object.gdc.goog/bucket-type"
)

// ObjectStorageTenantType indicates the tenant category for the project/namespace,
// e.g. "system" or "user".
type ObjectStorageTenantType string

const (
	System ObjectStorageTenantType = "system"
	User   ObjectStorageTenantType = "user"
)

// BucketType indicates the project type the bucket belongs to, e.g. "normal" or "shadow".
type BucketType string

const (
	Normal BucketType = "normal"
	Shadow BucketType = "shadow"
)
