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
*/
package v1

import (
	"gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/common"
	pkgcommon "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/common/utils"
	eedbcapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/dbcluster/v1"
	fleetapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/fleet/fleet/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
	consts "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/oracle/consts"
	"gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/oracle/metadata"
	oracleapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/oracle/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	DBClusterLabel            = "oracle.dbadmin.goog/dbcluster"
	DBClusterRestoreFromLabel = "oracle.dbadmin.goog/restorefrom"
	ResourcePrefix            = metadata.DBEngineShortName + "-"
)

// +kubebuilder:object:generate=true

type OracleInstanceSpec oracleapi.InstanceSpec

func (i *OracleInstanceSpec) CommonSpec() *occoreapi.InstanceSpec {
	return &i.InstanceSpec
}

//+kubebuilder:object:generate=true

// DBClusterSpec defines the desired state of DBCluster
type DBClusterSpec struct {
	eedbcapi.DBClusterSpec `json:",inline"`
	PrimarySpec            OracleInstanceSpec `json:"primarySpec,omitempty"`
}

//+kubebuilder:object:generate=true

// DBClusterStatus defines the observed state of DBCluster
type DBClusterStatus struct {
	eedbcapi.DBClusterStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:storageversion
//+kubebuilder:printcolumn:JSONPath=".status.primary.endpoint",name="PrimaryEndpoint",type="string"
//+kubebuilder:printcolumn:JSONPath=`.status.primary.phase`,name="PrimaryPhase",type="string"
//+kubebuilder:printcolumn:JSONPath=`.status.phase`,name="DBClusterPhase",type="string"

// DBCluster is the Schema for the dbclusters API
type DBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DBClusterSpec   `json:"spec,omitempty"`
	Status DBClusterStatus `json:"status,omitempty"`
}

func (d *DBCluster) DBPorts() []corev1.ServicePort {
	port := d.Spec.PrimarySpec.Port
	if port == 0 {
		port = consts.SSLListenerPort
	}
	return []corev1.ServicePort{
		{
			Name:       common.DBPortName,
			Protocol:   corev1.ProtocolTCP,
			Port:       int32(port),
			TargetPort: intstr.FromInt(port),
		},
	}
}

func (d *DBCluster) GetPrefixedName() string {
	return ResourcePrefix + d.GetName()
}

func (d *DBCluster) GetInternalName() string {
	return pkgcommon.GetDBClusterInternalName(d.GetName())
}

//+kubebuilder:object:root=true

// DBClusterList contains a list of DBCluster
type DBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBCluster{}, &DBClusterList{})
}

func (d *DBCluster) DBClusterSpec() *eedbcapi.DBClusterSpec {
	return &d.Spec.DBClusterSpec
}

func (d *DBCluster) PrimarySpec() eedbcapi.InstanceSpec {
	return &d.Spec.PrimarySpec
}

func (d *DBCluster) EntityStatus() *occoreapi.EntityStatus {
	return &d.Status.DBClusterStatus.EntityStatus
}

func (d *DBCluster) DBClusterStatus() *eedbcapi.DBClusterStatus {
	return &d.Status.DBClusterStatus
}

func (d *DBCluster) DBEngineShortName() string {
	return metadata.DBEngineShortName
}

func (d *DBCluster) DBEngineName() string {
	return string(fleetapi.Oracle)
}

// DBClusters returns a list of generic eedbcapi.DBCluster
func (dl *DBClusterList) DBClusters() (clusters []eedbcapi.DBCluster) {
	for _, cluster := range dl.Items {
		clusters = append(clusters, &cluster)
	}
	return clusters
}

var (
	_ eedbcapi.DBCluster = &DBCluster{}
	_ occoreapi.Entity   = &DBCluster{}
)
