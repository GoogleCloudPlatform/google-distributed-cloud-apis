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
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/common"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

const BackupNameLabel = "dbadmin.goog/backupname"

// +kubebuilder:object:generate=true

// ClonedDBClusterConfig defines the desired config of a cloned DBCluster.
type ClonedDBClusterConfig struct {
	// The name of cloned DBCluster.
	// +kubebuilder:validation:Required
	DBClusterName string `json:"dbclusterName,omitempty"`

	// To be supported: cross-namespace clone
	// DBClusterNamespace string
}

// +kubebuilder:object:generate=true

// RestoreSpec defines the desired state of Restore
type RestoreSpec struct {
	// The name of the source DBCluster to restore from.
	// This field is required.
	//
	// +kubebuilder:validation:required
	SourceDBCluster common.DBClusterRef `json:"sourceDBCluster"`

	// Previous point in time to restore to.
	// This field is optional.
	// Default to restore the latest available time point if not specified.
	//
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	// +kubebuilder:validation:optional
	PointInTime *metav1.Time `json:"pointInTime,omitempty"`

	// Settings for the cloned DBCluster. This lets you specify the name for the cloned DBCluster.
	// This field is optional.
	// Default to restore the source DBCluster if not specified.
	//
	// +kubebuilder:validation:optional
	ClonedDBClusterConfig *ClonedDBClusterConfig `json:"clonedDBClusterConfig,omitempty"`
}

// +kubebuilder:object:generate=true

// RestoreStatus defines the observed state of restore.
type RestoreStatus struct {
	occoreapi.EntityStatus `json:",inline"`
	Phase                  occoreapi.RestorePhase `json:"phase,omitempty"`
	// Creation time of the Restore
	// +optional
	CreateTime *metav1.Time `json:"createTime,omitempty"`
	// Completion time of the Restore
	// +optional
	CompleteTime *metav1.Time `json:"completeTime,omitempty"`
	// Actual point-in-time this restore brings the target DBCluster into.
	// Might be different from value specified in spec.PointInTime.
	// nullon(samwise-fleet)
	// +optional
	RestoredPointInTime *metav1.Time `json:"restoredPointInTime,omitempty"`
}

type Restore interface {
	occoreapi.Entity
	RestoreSpec() *RestoreSpec
	RestoreStatus() *RestoreStatus
	DBEngineName() string
}

type RestoreList interface {
	ctrlclient.ObjectList
	RestoreListItem() []Restore
}
