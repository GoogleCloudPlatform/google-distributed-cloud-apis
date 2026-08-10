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
	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BackupRepositoryManagerSpec defines the desired state of BackupRepositoryManager
type BackupRepositoryManagerSpec struct {
	// BackupRepositorySpec contains all the details needed to create a backup repository
	// inside a user cluster. Note that the ImportPolicy field of the BackupRepositorySpec
	// will not be honored here. The created BackupRepository will be "ReadWrite" if the
	// cluster it is being created in is the ReadWriteCluster. It will be "ReadOnly" if
	// the cluster it is being created in exists in the ReadOnlyClusters list. This field
	// is immutable.
	// +kubebuilder:validation:Required
	BackupRepositorySpec BackupRepositorySpec `json:"backupRepositorySpec"`

	// ReadWriteCluster specifies the single cluster (baremetal.cluster.gke.io/Cluster) which has permission
	// to create a ReadWrite backup repository using the data in the BackupRepositorySpec field. A ReadWrite
	// repository can be used to schedule/create Backups, BackupPlans, and Restores, and is effectively owned
	// by the cluster it is created in.
	// NOTE: A BackupRepository can only be used as ReadWrite by at most one k8s cluster. This field is mutable.
	// +optional
	ReadWriteCluster *NamespacedName `json:"readWriteCluster,omitempty"`

	// ReadOnlyClusters specifies the list of clusters (baremetal.cluster.gke.io/Cluster) which have permission
	// to create a ReadOnly backup repository using the data in the BackupRepositorySpec field.
	// A ReadOnly repository can only be used to import and view backups. No new backups/resources can be created in
	// this repository, but restores can use and reference read-only backups for restoration. Intended to import
	// backups from another cluster for a cross-cluster restore. There is no restriction on how often a
	// BackupRepository can be used as ReadOnly. This field is mutable.
	ReadOnlyClusters []NamespacedName `json:"readOnlyClusters,omitempty"`
}

// BackupRepositoryManagerStatus defines the observed state of BackupRepositoryManager
type BackupRepositoryManagerStatus struct {

	// BackupCount is the total number of backups that have been created inside the
	// storage bucket that the BackupRepositorySpec points to.
	BackupCount int `json:"backupCount,omitempty"`

	// BackupPlanCount is the total number of backup plans that have been created inside the
	// storage bucket that the BackupRepositorySpec points to.
	BackupPlanCount int `json:"backupPlanCount,omitempty"`

	// RestoreCount is the total number of restores that have been created inside the
	// storage bucket that the BackupRepositorySpec points to.
	RestoreCount int `json:"restoreCount,omitempty"`

	// RestorePlanCount is the total number of restore plans that have been created inside the
	// storage bucket that the BackupRepositorySpec points to.
	RestorePlanCount int `json:"restorePlanCount,omitempty"`

	// ClusterStatuses contains the current status in regard to creating the desired backup repository in each cluster.
	// The list will contain an entry for all clusters specified in the ReadOnlyClusters list as well as the ReadWriteCluster.
	// The backup repository being successfully created in a cluster is indicated by its entry in the ClusterStatuses array having
	// the "Ready" condition as "True", otherwise it will be "False".
	ClusterStatuses []ClusterStatus `json:"clusterStatuses,omitempty"`

	// Conditions indicates whether or not the Backup Repository Manager and its created backup repositories are in a healthy state.
	Conditions *[]metav1.Condition `json:"conditions,omitempty"`

	// The most recent errors from reconciliation with the observed times included.
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

// ClusterStatus contains backup repository creation status for a given cluster.
type ClusterStatus struct {
	Cluster    NamespacedName     `json:"cluster"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +genclient
// +genclient:nonNamespaced
// +gdcloud:manifest:relevant=false,oc=back
// BackupRepositoryManager is the Schema for the backuprepositorymanagers API
type BackupRepositoryManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupRepositoryManagerSpec   `json:"spec,omitempty"`
	Status BackupRepositoryManagerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// BackupRepositoryManagerList contains a list of BackupRepositoryManager
type BackupRepositoryManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupRepositoryManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupRepositoryManager{}, &BackupRepositoryManagerList{})
}
