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
	v1 "k8s.io/api/core/v1"
)

type StorageType string

const (
	StorageTypeGCS StorageType = "GCS"
	StorageTypeS3  StorageType = "S3"
)

//+kubebuilder:object:generate=true

// S3Options contains data for configuring access to S3 compatible Object Store
type S3Options struct {
	// Bucket is a required field, (ex: dbs-dump-bucket)
	// A user is to ensure proper write access to the storage bucket from within the
	// Operator.
	// +required
	Bucket *string `json:"bucket"`

	// Object key for the dump files. (ex: ods-dump/scottschema.dmp).
	// +required
	Key *string `json:"key"`

	// Region is S3 region the bucket resides in.
	// +required
	Region string `json:"region,omitempty"`

	// Endpoint is S3 end point.
	// +required
	Endpoint string `json:"endpoint,omitempty"`

	// SecretRef is a reference to the secret that stores bucket access information.
	// +optional
	SecretRef v1.SecretReference `json:"secretRef,omitempty"`

	// CertRef is a reference to the secret that stores CA certs used to assess the S3 endpoint.
	// The value of key 'ca.crt' inside this secret will be used.
	// Default to skip SSL verification if not specified.
	// nullon(dbs-fleet,dbs-local)
	// +optional
	CertRef v1.SecretReference `json:"certRef,omitempty"`

	// CABundle is a pool of PEM encoded CA certs which will be used to validate the storageGrid's server
	// certificate.
	// nullon(samwise-fleet)
	// +optional
	CABundle []string `json:"caBundle,omitempty"`
}

//+kubebuilder:object:generate=true

// GCSOptions contains data for configuring access to GCS
type GCSOptions struct {
	// Bucket is a required field, (ex: dbs-dump-bucket)
	// A user is to ensure proper write access to the storage bucket from within the
	// Operator.
	// +required
	Bucket *string `json:"bucket"`

	// Object key for the dump files. (ex: ods-dump/scottschema.dmp).
	// +required
	Key *string `json:"key"`

	// SecretRef is a reference to the secret that stores GCS access information.
	// +optional
	SecretRef v1.SecretReference `json:"secretRef,omitempty"`
}

//+kubebuilder:object:generate=true

type StorageSpec struct {
	// Type of Repository (ex: S3, GCS), which tells the agent which storage system/API to use.
	// +required
	// +kubebuilder:validation:Enum=GCS;S3
	Type StorageType `json:"type"`

	// S3Options is a reference to S3 dependent options (Ex: S3 Access Secret, End Point, Region).
	// +optional
	S3Options *S3Options `json:"s3Options,omitempty"`

	// GCSOptions is a reference to GCS dependent options.
	// +optional
	GCSOptions *GCSOptions `json:"gcsOptions,omitempty"`
}
