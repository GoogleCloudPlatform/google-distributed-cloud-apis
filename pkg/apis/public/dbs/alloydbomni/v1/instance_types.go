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

	occoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

// InstanceSpec defines the desired state of AlloyDBOmniInstance
type InstanceSpec struct {
	// Instance specs that are common across all database engines.
	occoreapi.InstanceSpec `json:",inline"`
	// +optional
	// Specify AlloyDB Omni specific features
	Features Features `json:"features,omitempty"`
}

// +kubebuilder:object:generate=true
// Feature Spec
type Features struct {
	// GoogleMLExtension allow users to call Vertex AI models from the database.
	// This extension is optional.
	//
	// For information on how to configure the GoogleMLExtension, follow the user guide in
	// https://cloud.google.com/alloydb/docs/ai/configure-vertex-ai#gcloud.
	//
	// Note that enabling this extension for an existing database will cause a database restart.
	//
	// +optional
	GoogleMLExtension *GoogleMLExtensionSpec `json:"googleMLExtension,omitempty"`
}

// +kubebuilder:object:generate=true
// Google ML Extension Spec
type GoogleMLExtensionSpec struct {
	// The configuration for the GoogleMLExtension. This feature is required if this extension is
	// enabled.
	//
	// +required
	Config GoogleMLExtensionConfig `json:"config,omitempty"`
}

// +kubebuilder:object:generate=true
// Google ML Extension Config
type GoogleMLExtensionConfig struct {
	// VertexAIKeyRef is the name of the secrect where the Vertex AI private key is stored. This
	// field is required.
	//
	// +required
	VertexAIKeyRef string `json:"vertexAIKeyRef,omitempty"`
}

// InstanceStatus defines the observed state of AlloyDBOmniInstance
type InstanceStatus struct {
	// Instance status that is common across all database engines.
	occoreapi.InstanceStatus `json:",inline"`

	// PrimaryPodIP indicates the IP of AlloyDBOmni primary pod.
	PrimaryPodIP string `json:"primaryPodIP,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.endpoint",name="Endpoint",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.url",name="URL",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.labels['dbs\\.internal\\.dbadmin\\.goog/ha-role']",name="Role",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.phase`,name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=`.status.description`,name="Message",type="string"
// Instance is the Schema for the instances API
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceSpec   `json:"spec,omitempty"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of AlloyDBOmni Instance
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
