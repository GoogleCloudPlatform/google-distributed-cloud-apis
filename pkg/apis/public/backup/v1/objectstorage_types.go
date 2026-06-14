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

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Represents an object storage endpoint as well
// as the credentials to access it.
type ObjectStorageEndpoint struct {
	// The endpoint used to access object storage. For example, `https://s3.svc`.
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint" reflect:"unexport"`

	// The bucket within the endpoint to scope reference.
	// +kubebuilder:validation:Required
	Bucket string `json:"bucket" reflect:"unexport"`

	// The reference to S3 access credentials used to access object storage.
	// +kubebuilder:validation:Required
	SecretReference corev1.SecretReference `json:"secretReference" reflect:"unexport"`
}

// Refer to go/gdch1.0-object-storage-backup for more context.

// +gdcloud:manifest:relevant=false,oc=back
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// Configures backups of object storage.
type ObjectStorageBackupConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The source object storage to be backed up.
	// +kubebuilder:validation:Required
	Source ObjectStorageEndpoint `json:"source" reflect:"unexport"`

	// The target to which the backup is written.
	// +kubebuilder:validation:Required
	Target ObjectStorageEndpoint `json:"target" reflect:"unexport"`

	// Specifies whether to store each transfer in a timestamped sub-path of the bucket.
	// If set to `True`, the structure is `YYYY-MM-DD-hh:mm:ss`.
	// This helps to store copies taken from multiple points in time of a backup.
	// Defaults to `False`.
	// +optional
	// +kubebuilder:default:=false
	Timestamp bool `json:"timestamp" reflect:"unexport"`
}

func init() {
	SchemeBuilder.Register(&ObjectStorageBackupConfig{})
}
