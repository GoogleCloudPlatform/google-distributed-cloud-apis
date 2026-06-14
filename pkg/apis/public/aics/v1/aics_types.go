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

// aics_package_types.go is the type definition for each aics application (inferenceGateway)

// AICSApplicationName describes which application is deployed by L1.
// currently supported apps:
// - Inference-Gateway: Inference Gateway for Gemini Pro
// - l2opr: The l2 operator
// - l2crd: The l2 crd components
// - eps: Endpoint Picker Service
// +kubebuilder:validation:Enum=inference-gateway;l2opr;l2crd;eps
type AICSApplicationName string

// AICSApplicationPath maps to an AICSApplication's `Path` field, which was
// originally used to fetch the target application's Helm chart by the L1
// Operator as part of enablement. This functionality is no longer active and
// has been superceded by OCLCM.
type AICSApplicationPath string

// AICSCluster describes which cluster the target application is deployed in.
// Currently supported values: system, admin.
// +kubebuilder:validation:string=aicsSystem;aicsAdmin
type AICSCluster string

const (
	// points to system cluster
	AICSSystem AICSCluster = "aicsSystem"

	// points to admin cluster
	AICSAdmin AICSCluster = "aicsAdmin"
)

// AICSApplication defines the type of AICS Application
type AICSApplication struct {
	Name AICSApplicationName `json:"name"`

	// path of Application helm in harbor
	Path string `json:"path"`
}

// AICSApplicationStatus keeps track of each application thats is deployed
type AICSApplicationStatus struct {
	Application AICSApplication `json:"application"`

	// Indicates if application is deployed successfully
	Deployed *bool `json:"deployed"`

	// Hold the messages if any
	// +optional
	Messages []string `json:"messages,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// kubebuilder:object:generate=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// AICSPackage is the Schema for the packages API
type AICSPackage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AICSPackageSpec   `json:"spec,omitempty"`
	Status AICSPackageStatus `json:"status,omitempty"`
}

// AICSPackageSpec defines the desired state of AICS Package
type AICSPackageSpec struct {
	// applications list ex.inference-gateway
	// +listType=map
	// +listMapKey=name
	Applications []AICSApplication `json:"applications,omitempty"`
}

// PackageStatus defines the observed state of Package
type AICSPackageStatus struct {
	// The generation observed by the application controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// The status of deployed applications
	ApplicationStatuses map[AICSApplicationName]AICSApplicationStatus `json:"applicationStatuses,omitempty"`
}

// +kubebuilder:object:root=true
// AICSPackageList contains a list of AICS Packages
type AICSPackageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AICSPackage `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&AICSPackage{},
		&AICSPackageList{},
	)
}
