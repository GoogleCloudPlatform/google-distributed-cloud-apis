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
)

// WorkloadSelector specifies the selectors to select workloads (Pods or/and VMs) inside a Project.
// +kubebuilder:validation:MinProperties:=1
// +kubebuilder:validation:MaxProperties:=1
// +kubebuilder:default:={labelSelector:{workloads:{}}}
type WorkloadSelector struct {
	// LabelSelector selects the workloads based on Kubernetes labels.
	LabelSelector *WorkloadLabelSelector `json:"labelSelector,omitempty"`
}

// WorkloadLabelSelector selects the workloads in a Project based on Kubernetes labels.
// +kubebuilder:default:={workloads:{}}
type WorkloadLabelSelector struct {
	// Clusters are selected based on Kubernetes labels applied to cluster objects.
	// This field should only be specified when selecting a Standard or a Shared user cluster.
	// This selector currently only supports selection by cluster name.
	// To select by name, use `matchLabels` with the key `kubernetes.io/metadata.name`.
	// +kubebuilder:validation:XValidation:rule="size(self.matchLabels)==1",message="`matchLabels` must have exactly one key specified."
	// +kubebuilder:validation:XValidation:rule="'kubernetes.io/metadata.name' in self.matchLabels",message="`matchLabels` must have the key `kubernetes.io/metadata.name`."
	// +kubebuilder:validation:XValidation:rule="size(self.matchLabels['kubernetes.io/metadata.name']) > 0",message="Value of the key `kubernetes.io/metadata.name` must not be empty."
	// +kubebuilder:validation:XValidation:rule="!has(self.matchExpressions)",message="`matchExpressions` must not be specified."
	Clusters *metav1.LabelSelector `json:"clusters,omitempty"`
	// Namespaces are selected based on Kubernetes labels applied to namespace objects.
	// The `clusters` field is required if this field is specified.
	// Namespaces are selected only from the clusters selected by `clusters` field.
	Namespaces *metav1.LabelSelector `json:"namespaces,omitempty"`
	// Workloads are selected based on their Kuberentes labels.
	// If `clusters` and `namespaces` fields are specified, then workloads are selected
	// only from the selected clusters and namespaces.
	// If `clusters` and `namespaces` fields are omitted, workloads are selected
	// from the project based on the specified workload labels.
	// If empty, all workloads are selected.
	// +kubebuilder:default:={}
	Workloads *metav1.LabelSelector `json:"workloads,omitempty"`
}
