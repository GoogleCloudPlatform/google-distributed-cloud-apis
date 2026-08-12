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
	"sigs.k8s.io/controller-runtime/pkg/client"

	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

type DBInstanceType string

const (
	// ReadPoolType defines a DBInstance type in which all nodes are read
	// replicas of the primary node of the DBCluster.
	ReadPoolType DBInstanceType = "ReadPool"
)

//+kubebuilder:object:generate=true

type DBInstanceSpec struct {
	// InstanceType indicates the type of the DBInstance.
	// Currently only supports "ReadPool" type.
	// +kubebuilder:validation:Enum="ReadPool"
	// +kubebuilder:default="ReadPool"
	InstanceType DBInstanceType `json:"instanceType,omitempty"`

	// NodeCount determines the number of DBNodes that should
	// be created for this DBInstance.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=20
	// +kubebuilder:default=1
	NodeCount int `json:"nodeCount"`

	// DBClusterParent is the DBCluster this DBInstance replicates from.
	DBClusterParent *corev1.LocalObjectReference `json:"dbcParent,omitempty"`

	// Resource specifies the resources used for the Instances (i.e, DBNodes)
	// belonging to this DBInstance.
	// If omitted, the instance will use the same resources as the DBCluster's Primary Instance.
	// +optional
	Resources *occoreapi.Resource `json:"resources,omitempty"`

	// IsStopped stops the DBNodes in this DBInstance when true. This field is
	// optional and defaults to false. Stopping the DBCluster's Primary DBNode
	// does not automatically stop the DBNodes of ReadPool DBInstances.
	//
	// When stopped, the compute resources (CPU, memory) of the instance are
	// released. However, the DBNode still keeps the storage resource.
	// +optional
	IsStopped *bool `json:"isStopped,omitempty"`

	// ProgressTimeout determines the number of seconds the controller will
	// attempt to provision a DBNode, or which a DBNode could be not-ready for,
	// before it considers the DBNode to have failed.
	//
	// A value of 0 means that no timeout will be used.
	// +kubebuilder:default="30m"
	ProgressTimeout metav1.Duration `json:"progressTimeout,omitempty"`

	// LogReplicationSlot determines whether replication slot should be logged on the primary's slot allowlist.
	// By default it is false. Enabling this on a running system requires database restart of the primary.
	// +optional
	LogReplicationSlot *bool `json:"logReplicationSlot,omitempty"`

	// SchedulingConfig specifies how the instance should be scheduled on Kubernetes nodes.
	// If omitted, the instance will use the same scheduling config as the DBCluster's Primary Instance.
	//
	// When any field inside the scheduling config changes, it can lead to rescheduling of the
	// k8s pod onto a different node based on the config.
	//
	// +optional
	SchedulingConfig *occoreapi.SchedulingConfig `json:"schedulingconfig,omitempty"`
}

//+kubebuilder:object:generate=true

type DBInstanceStatus struct {
	occoreapi.EntityStatus `json:",inline"`

	// Endpoints are the endpoints from which the DBNodes in the DBInstance
	// can be accessed.
	//
	// +kubebuilder:validation:Type=array
	// +kubebuilder:validation:Items=Type=object
	// +patchStrategy=merge
	// +patchMergeKey=name
	// +listType=map
	// +listMapKey=name
	Endpoints []occoreapi.Endpoint `json:"endpoints,omitempty"`
}

// DBInstance maintains a pool of Instances (i.e, DBNodes) which hold the same
// configuration.
type DBInstance interface {
	occoreapi.Entity
	DBInstanceSpec() *DBInstanceSpec
	DBInstanceStatus() *DBInstanceStatus
}

// DBInstanceList is a common interface for all the DBInstanceList implementations
type DBInstanceList interface {
	client.ObjectList
	DBInstances() []DBInstance
}
