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

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status

// +genclient
// +genclient:nonNamespaced
// FeatureGate represents the maturity stage of features.
type FeatureGate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FeatureGateSpec   `json:"spec,omitempty"`
	Status FeatureGateStatus `json:"status,omitempty"`
}

// FeatureGateSpec represents the specification of featureGates.
type FeatureGateSpec struct {
	// Stages are the default feature stages defined in source code.
	Stages map[string]FeatureStage `json:"stages,omitempty"`

	// Features is a map of feature name to its detailed configuration (stage, type).
	// This is the new standard for defining features.
	// +optional
	Features map[string]FeatureDetail `json:"features,omitempty"`

	// Overrides are the feature stages overridden in runtime.
	Overrides map[string]FeatureStage `json:"overrides,omitempty"`
}

// FeatureGateStatus defines the observed state of FeatureGate
type FeatureGateStatus struct {
	// ClientPods lists the pods that are clients of this FeatureGate, with details about each pod's feature enablement status.
	ClientPods []ClientPod `json:"clientPods,omitempty"`
}

// ClientPod describes a single pod that is a client of the FeatureGate, including its namespace, name, and whether the feature is enabled for this pod.
type ClientPod struct {
	// Namespace is the namespace where the client pod is located.
	Namespace string `json:"namespace,omitempty"`

	// Name is the name of the client pod.
	Name string `json:"name,omitempty"`

	// Enabled indicates whether the feature is currently enabled on this pod.
	Enabled bool `json:"enabled,omitempty"`

	// Feature specifies the name of the feature this setting applies to.
	Feature string `json:"feature,omitempty"`
}

// FeatureStage represents a feature's maturity.
// +kubebuilder:validation:Enum=dev;test;preview;production;accredited
type FeatureStage string
type FeatureType string

const (
	Dev                FeatureStage = "dev"
	Test               FeatureStage = "test"
	Preview            FeatureStage = "preview"
	Production         FeatureStage = "production"
	Accredited         FeatureStage = "accredited"
	FeatureTypeSCR     FeatureType  = "scr"
	FeatureTypeRuntime FeatureType  = "runtime"
)

type FeatureDetail struct {
	Stage FeatureStage `json:"stage"`
	Type  FeatureType  `json:"type,omitempty"`
}

// +kubebuilder:object:root=true

// FeatureGateList is a list of FeatureGate resources.
type FeatureGateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []FeatureGate `json:"items"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="MinStage",type=string,JSONPath=`.spec.minStage`
// +kubebuilder:resource:scope=Cluster

// +genclient
// +genclient:nonNamespaced
// Stage represents the current instance's features stage.
type Stage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec StageSpec `json:"spec"`
}

// StageSpec is the specification of the Stage.
type StageSpec struct {
	// MinStage specifies the minimum maturity stage required for features to be enabled in the instance.
	// This field is immutable.
	// +kubebuilder:validation:Required
	MinStage FeatureStage `json:"minStage"`
}

// +kubebuilder:object:root=true

// StageList contains a list of Stage.
type StageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Stage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FeatureGate{}, &FeatureGateList{})
	SchemeBuilder.Register(&Stage{}, &StageList{})
}
