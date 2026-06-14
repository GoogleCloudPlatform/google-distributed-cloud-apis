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
)

// Enables us to choose the QoS policy for a given NFS share.
// +kubebuilder:validation:Enum=Standard;Premium
type FileShareTier string

const (
	Standard FileShareTier = "Standard"
	Premium  FileShareTier = "Premium"
)

// Enables us to specify the protocol on which a share will be exposed.
// +kubebuilder:validation:Enum=NFSv4
type FileShareProtocol string

const (
	NFSv4 FileShareProtocol = "NFSv4"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=file
// +genclient
// Represents a single indivisible unit of storage. It is a zonal artifact, and the
// existence of a FileShare in a zone indicates that there is, or is expected to be, an
// actual volume existing in that zone holding data for that share. Modifications to the
// FileShare may be restricted based on the nature of the backing volume. For instance, if the
// FileShare is read only, we might not be able to accept all changes to the specification.
type FileShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired state of a file share.
	// +optional
	Spec FileShareSpec `json:"spec,omitempty"`

	// The observed state of a file share.
	// Read-only.
	// +optional
	Status FileShareStatus `json:"status,omitempty"`
}

// Defines the desired state of a file share.
type FileShareSpec struct {
	// The amount of client-writable space in the share. Changes in capacity
	// are considered best-effort. Units are bytes.
	Capacity *resource.Quantity `json:"capacity"`

	// The performance tier of the file share. Changes in tier are considered
	// best-effort. If not specified, the default is Standard.
	// +optional
	Tier FileShareTier `json:"tier,omitempty"`

	// The protocol on which the file share should be exposed.
	// If not specified, the default is NFSv4.
	// +optional
	Protocol FileShareProtocol `json:"protocol,omitempty"`
}

// Represents the observed state of a file share.
type FileShareStatus struct {
	// A list of observed conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The current capacity of the share. Units are bytes.
	// +optional
	CurrentCapacity *resource.Quantity `json:"currentCapacity,omitempty"`

	// The current performance tier of the file share.
	// +optional
	CurrentTier FileShareTier `json:"currentTier,omitempty"`

	// The information necessary to connect to the file share in the zone. It
	// is expected that once this information is set, it is fixed to these values for the entire
	// life of the volume, regardless if the volume's primary zone changes.
	// +optional
	ConnectionInfo *FileShareConnectionInformation `json:"connectionInfo,omitempty"`

	// The name of the storage volume found to be associated with this file share.
	// +optional
	StorageVolumeName string `json:"storageVolumeName,omitempty"`
}

// Contains the information necessary to connect to the file share in the zone.
type FileShareConnectionInformation struct {
	// A string representation of the network location of the file server. This
	// may either be a raw address or domain name, either of which are assumed to be resolvable
	// or routable on the network a pod in this namespace would be assigned to.
	Server string `json:"server"`

	// The port on which the file server is serving the file share.
	Port int32 `json:"port"`

	// The path at which the file share exists on the file server. The exact semantic
	// nature is protocol specific.
	Path string `json:"path"`
}

// +kubebuilder:object:root=true

// Represents a collection of file shares.
type FileShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []FileShare `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&FileShare{},
		&FileShareList{},
	)
}
