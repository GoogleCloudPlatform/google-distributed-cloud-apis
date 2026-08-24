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

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/common/global/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Name",JSONPath=".metadata.name",type=string
// +kubebuilder:printcolumn:name="Active",JSONPath=".spec.active",type=boolean
// +kubebuilder:printcolumn:name="Location Type",JSONPath=".spec.locationType",type=string
// +kubebuilder:printcolumn:name="ILMReady",JSONPath=".status.conditions[?(@.type=='ILMReady')].status",type=string
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type=string
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// Defines the schema for the BucketLocationConfig API.
// +genclient
// +genclient:nonNamespaced
type BucketLocationConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BucketLocationSpec         `json:"spec,omitempty"`
	Status BucketLocationConfigStatus `json:"status,omitempty"`
}

// Defines the observed state of the BucketLocationConfig.
type BucketLocationConfigStatus struct {
	// The observations of the overall state of the resource.
	// Known condition types: Ready.
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The list of cluster statuses where the resource is rolled out to.
	// +listType=map
	// +listMapKey=name
	Clusters []BucketLocationClusterStatus `json:"clusters,omitempty"`
}

// Provides the status of a BucketLocation rolling out to a particular cluster.
// Follows pkg/apis/common/global/v1alpha1 ZoneStatus
type BucketLocationClusterStatus struct {
	// The name of the cluster where the replica this status represents is in.
	Name string `json:"name"`

	// The status of rolling out the replica to the cluster.
	RolloutStatus v1alpha1.ZoneRolloutStatus `json:"rolloutStatus,omitempty"`

	// The reconciliation status of the replica collected from the cluster.
	// Any condition within the field that has an .observedGeneration less than
	// .rolloutStatus.replicaGeneration is out of date.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// Contains a list of BucketLocationConfigs.
// +kubebuilder:object:root=true
type BucketLocationConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLocationConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BucketLocationConfig{}, &BucketLocationConfigList{})
}
