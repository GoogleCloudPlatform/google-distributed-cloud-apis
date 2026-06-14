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

// gemini_package_types.go is the type definition for each gemini application (genai router)

// GeminiApplicationName describes which application is deployed by L1.
// currently supported apps:
// - GENAI-Router: The Generative AI Router
// - l2opr: The l2 operator
// - l2crd: The l2 crd components
// - vertex-1p-serving-stack: The Vertex 1P serving stack components
// +kubebuilder:validation:Enum=genai-router;l2opr;l2crd;vertex-1p-serving-stack
type GeminiApplicationName string

// GeminiApplicationPath maps to an GeminiApplication's `Path` field, which was
// originally used to fetch the target application's Helm chart by the L1
// Operator as part of enablement. This functionality is no longer active and
// has been superceded by OCLCM.
type GeminiApplicationPath string

// GeminiCluster describes which cluster the target application is deployed in.
// Currently supported values: system, admin.
// +kubebuilder:validation:string=geminiSystem;geminiAdmin
type GeminiCluster string

const (
	// points to system cluster
	GeminiSystem GeminiCluster = "geminiSystem"
	// points to admin cluster
	GeminiAdmin GeminiCluster = "geminiAdmin"

	// GeminiNodeLabel is the label key for nodes that can run Gemini pods.
	GeminiNodeLabel = "gemini.gdc.goog/large-gemini"
)

// GeminiApplication defines the type of Gemini Application
type GeminiApplication struct {
	Name GeminiApplicationName `json:"name"`
	// path of Application helm in harbor
	Path string `json:"path"`
}

// GeminiApplicationStatus keeps track of each application thats is deployed
type GeminiApplicationStatus struct {
	Application GeminiApplication `json:"application"`
	// Indicates if application is deployed successfully
	Deployed *bool `json:"deployed"`
	// Hold the messages if any
	// +optional
	Messages []string `json:"messages,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type GeminiPackage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GeminiPackageSpec   `json:"spec,omitempty"`
	Status            GeminiPackageStatus `json:"status,omitempty"`
}

// GeminiPackageSpec defines the desired state of Gemini Package
type GeminiPackageSpec struct {
	// applications list ex. ocr
	// +listType=map
	// +listMapKey=name
	Applications []GeminiApplication `json:"applications,omitempty"`
}

// PackageStatus defines the observed state of Package
type GeminiPackageStatus struct {
	// The generation observed by the application controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// The status of deployed applications
	ApplicationStatuses map[GeminiApplicationName]GeminiApplicationStatus `json:"applicationStatuses,omitempty"`
}

// +kubebuilder:object:root=true
// GeminiPackageList contains a list of Gemini Packages
type GeminiPackageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GeminiPackage `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&GeminiPackage{},
		&GeminiPackageList{},
	)
}
