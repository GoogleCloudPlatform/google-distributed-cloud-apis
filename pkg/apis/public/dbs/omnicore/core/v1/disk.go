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
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:generate=true

// DiskSpec defines the desired state of a disk.
type DiskSpec struct {
	// Name of the disk. This field is required.
	//
	// The allowed values are: "DataDisk", "LogDisk", "BackupDisk", "BackupRepoDisk", and "ObsDisk".
	//
	// "DataDisk" is mounted in the database pod to store the database files.
	// "LogDisk" is mounted in the database pod to store the archived logs and is only used
	// for vendor specific backup.
	// "BackupDisk" is mounted in the database pod to store the backup configs, and is also
	// used for other purposes related to backup features.
	// "BackupRepoDisk" is mounted in the backup repo pod to store the backups and archived WALs
	// for the backup with local storage.
	// "ObsDisk" is mounted in the database pod to store the observability data.
	//
	// +required
	// +kubebuilder:validation:Enum=DataDisk;LogDisk;BackupDisk;ObsDisk;BackupRepoDisk;
	Name string `json:"name"`

	// Disk size in bytes for example, "10Gi" for 10 Gibibytes. This field is required.
	//
	// The allowed size unit prefixes are: "Ki", "Mi", "Gi", "Ti, "Pi" and "Ei" for 2-base. Also
	// "K", "M", "G", "T, "P" and "E" for 10-base. See https://en.wikipedia.org/wiki/Unit_prefix.
	//
	// +required
	// +kubebuilder:validation:Pattern=^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
	Size string `json:"size"`

	// StorageClass points to a particular CSI storage class. This field is optional.
	//
	// If the field is not set, then the default CSI storage class for the Kubernetes cluster is
	// used. If there is no default for the Kubernetes cluster,  then the Persistence Volume Claim
	// will fail and the database cluster will fail to provision.
	//
	// You can read more about storage classes in
	// https://kubernetes.io/docs/concepts/storage/storage-classes.
	//
	// +optional
	StorageClass string `json:"storageClass,omitempty"`

	// Additional annotations added to the Persistent Volume Claim. This field is optional.
	//
	// This allows to integrate with other tools.
	//
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// VolumeName is the binding reference to the Persistent Volume tied to this disk. This field
	// is optional.
	//
	// This allows to reuse an existing volume.
	//
	// Note that if this field is specified, the value "storageClass" will not take effect. You can
	// learn more about this in
	// https://kubernetes.io/docs/concepts/storage/persistent-volumes/#binding.
	//
	// +optional
	VolumeName string `json:"volumeName,omitempty"`

	// AccessModes contains the desired access modes for the volume.
	//
	// Refer to https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes for
	// more information.
	//
	// +optional
	AccessModes []corev1.PersistentVolumeAccessMode `json:"accessModes,omitempty"`

	// A label query over volumes to consider for binding. This field is optional.
	//
	// If this field is set, then the volume with matching labels is used as the backing
	// volume for the disk.
	//
	// Refer to https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#PersistentVolumeClaimSpec
	// for more information.
	//
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// Path on the VM where the disk is mounted. Must be a directory path
	// +optional
	// nullon(samwise-fleet,samwise-local,dbs-fleet,dbs-local)
	Path string `json:"path,omitempty"`

	// DataSource field can be used to specify either:
	// * An existing VolumeSnapshot object (snapshot.storage.k8s.io/v1)
	// * An existing PVC (Normal)
	// Note: This feature must be enabled in the operator, and the annotation
	// "dbcluster.dbadmin.goog/allow-datasource" must be set to "true" on the
	// DBCluster for this field to work.
	// +optional
	DataSource *corev1.TypedLocalObjectReference `json:"dataSource,omitempty"`
}

// DiskType is a type that points to the disk type
type DiskType string

const (
	DataDisk       = "DataDisk"
	LogDisk        = "LogDisk"
	ObsDisk        = "ObsDisk"
	BackupDisk     = "BackupDisk"
	BackupRepoDisk = "BackupRepoDisk"
)

// PVCShortName returns the name used for the PVC related to this DiskSpec used
// inside a Statefulset. This is not the name of the actual created PVC, but it
// can be used to refer to the PVC inside the Statefulset spec (e.g., volume mount).
func (ds DiskSpec) PVCShortName() string {
	return strings.ToLower(ds.Name)
}

// PVCFullName returns the name of the actual PVC created for the DiskSpec.
func (ds DiskSpec) PVCFullName(stsName string, idx int) string {
	return fmt.Sprintf(PVCNamePattern, ds.PVCShortName(), stsName, idx)
}
