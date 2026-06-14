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
	eecoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/ha/v1"
	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	FailoverMutateWebhookPath         = "/mutate-fleet-dbadmin-gdc-goog-v1-failover"
	FailoverStatusValidateWebhookPath = "/validate-fleet-dbadmin-gdc-goog-v1-failover-status"
)

//+kubebuilder:webhook:path=/mutate-fleet-dbadmin-gdc-goog-v1-failover,mutating=true,failurePolicy=fail,sideEffects=None,groups=fleet.dbadmin.gdc.goog,resources=failovers,verbs=create;update,versions=v1,name=mfailover.fleet.dbadmin.gdc.goog,admissionReviewVersions={v1,v1beta1}
//+kubebuilder:webhook:path=/validate-fleet-dbadmin-gdc-goog-v1-failover-status,mutating=false,failurePolicy=fail,sideEffects=None,groups=fleet.dbadmin.gdc.goog,resources=failovers/status,verbs=create;update;delete,versions=v1,name=vfailover.fleet.dbadmin.gdc.goog,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
//+kubebuilder:printcolumn:JSONPath=`.status.state`,name="state",type="string"
//+kubebuilder:printcolumn:JSONPath=".status.internal.phase",name="phase",type="string"
//+kubebuilder:printcolumn:JSONPath=".spec.dbclusterRef",name="dbcluster",type="string"

// Failover represents the parameters and status of a single failover operation.
type Failover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   eecoreapi.FailoverSpec   `json:"spec,omitempty"`
	Status eecoreapi.FailoverStatus `json:"status,omitempty"`
}

func (f *Failover) FailoverSpec() *eecoreapi.FailoverSpec {
	return &f.Spec
}

func (f *Failover) FailoverStatus() *eecoreapi.FailoverStatus {
	return &f.Status
}

func (f *Failover) EntityStatus() *occoreapi.EntityStatus {
	return &f.Status.EntityStatus
}

// +kubebuilder:object:root=true

// FailoverList contains a list of Failovers
type FailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Failover `json:"items"`
}

func (f *FailoverList) FailoverListItem() []eecoreapi.Failover {
	var items []eecoreapi.Failover
	for i := range f.Items {
		items = append(items, &f.Items[i])
	}
	return items
}

func init() {
	SchemeBuilder.Register(&Failover{}, &FailoverList{})
}

var (
	_ eecoreapi.Failover     = &Failover{}
	_ eecoreapi.FailoverList = &FailoverList{}
)
