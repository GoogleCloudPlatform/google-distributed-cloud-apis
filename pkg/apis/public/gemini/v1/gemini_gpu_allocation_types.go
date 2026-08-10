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

// GeminiGPUAllocationSpec defines the desired state of GeminiGPUAllocation.
type GeminiGPUAllocationSpec struct {
	// GPUAllocationUnits specifies the requested GPU units.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=0
	GPUAllocationUnits int64 `json:"gpuAllocationUnits"`
}

// GeminiGPUAllocationStatus defines the observed state of GeminiGPUAllocation.
type GeminiGPUAllocationStatus struct {
	// AllocatedGPUUnits specifies the actual GPU units allocated after pro-rata degradation scaling.
	AllocatedGPUUnits int64 `json:"allocatedGpuUnits"`

	// HealthyGPURatio represents the normalized decimal ratio of currently healthy
	// physical GPU units to total configured GPU units across the cluster (e.g. "1.00", "0.75").
	HealthyGPURatio string `json:"healthyGpuRatio,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=geminigpuallocations,singular=geminigpuallocation,scope=Namespaced,shortName=gga

// GeminiGPUAllocation is the Schema for the geminigpuallocations API.
type GeminiGPUAllocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GeminiGPUAllocationSpec   `json:"spec,omitempty"`
	Status GeminiGPUAllocationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GeminiGPUAllocationList contains a list of GeminiGPUAllocation.
type GeminiGPUAllocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GeminiGPUAllocation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GeminiGPUAllocation{}, &GeminiGPUAllocationList{})
}
