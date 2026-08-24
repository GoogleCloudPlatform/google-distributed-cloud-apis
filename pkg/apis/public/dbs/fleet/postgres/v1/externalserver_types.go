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
	v1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/externalserver/v1"
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:generate=true

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// ExternalServer is the Schema for the ExternalServer API.
type ExternalServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalServerSpec   `json:"spec,omitempty"`
	Status ExternalServerStatus `json:"status,omitempty"`
}

func (in *ExternalServer) EntityStatus() *occoreapi.EntityStatus {
	return &in.Status.EntityStatus
}

func (in *ExternalServer) ExternalServerSpec() occoreapi.ExternalServerSpec {
	return in.Spec.ExternalServerSpec
}

func (in *ExternalServer) ExternalServerStatus() *occoreapi.ExternalServerStatus {
	return &in.Status.ExternalServerStatus
}

// ExternalServerSpec defines metadata of an external database server used for migration.
type ExternalServerSpec struct {
	occoreapi.ExternalServerSpec `json:",inline"`
}

type ExternalServerStatus struct {
	occoreapi.ExternalServerStatus `json:",inline"`
}

//+kubebuilder:object:root=true

type ExternalServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalServer{}, &ExternalServerList{})
}

var _ v1.ExternalServer = &ExternalServer{}
