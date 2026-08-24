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
	eehealthapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/health/v1"
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:generate=true

// DBClusterHealthStatus represents the current health state of a DBCluster's instances.
type DBClusterHealthStatus struct {
	eehealthapi.DBClusterHealthStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=dbs
type DBClusterHealth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status DBClusterHealthStatus `json:"status,omitempty"`
}

func (in *DBClusterHealth) DBClusterHealthStatus() *eehealthapi.DBClusterHealthStatus {
	return &in.Status.DBClusterHealthStatus
}

// EntityStatus returns the DBClusterHealth entity status.
func (in *DBClusterHealth) EntityStatus() *occoreapi.EntityStatus {
	return &in.Status.EntityStatus
}

// +kubebuilder:object:root=true
// DBClusterHealthList contains a list of DBClusterHealth
type DBClusterHealthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBClusterHealth `json:"items"`
}

func (j *DBClusterHealthList) DBClusterHealthListItem() []eehealthapi.DBClusterHealth {
	var items []eehealthapi.DBClusterHealth
	for i := range j.Items {
		items = append(items, &j.Items[i])
	}
	return items
}

func init() {
	SchemeBuilder.Register(&DBClusterHealth{}, &DBClusterHealthList{})
}

var (
	_ eehealthapi.DBClusterHealth     = &DBClusterHealth{}
	_ eehealthapi.DBClusterHealthList = &DBClusterHealthList{}
)
