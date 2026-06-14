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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Defines the specification or expected state of the `Dashboard` object.
type DashboardSpec struct {
	// The `ConfigMap` object that contains the JSON model of the Grafana
	// dashboard to include.
	ConfigMapRef DashboardConfigMap `json:"configMapRef,omitempty"`

	// The name of the folder where the system must store the dashboards.
	// +kubebuilder:validation:MinLength=1
	Foldername string `json:"foldername,omitempty"`
}

// Provides information about the `ConfigMap` object that contains the JSON
// model of a Grafana dashboard.
type DashboardConfigMap struct {
	// The name of the `ConfigMap` object containing the JSON model of
	// the dashboard.
	Name string `json:"name,omitempty"`

	// The namespace of the `ConfigMap` object containing the JSON model of
	// the dashboard.
	Namespace string `json:"namespace,omitempty"`

	// The key of the `ConfigMap` object containing the JSON model of the
	// dashboard.
	Key string `json:"key,omitempty"`
}

// Defines the observed state of the `Dashboard` object.
type DashboardStatus struct {
	// A list of conditions observed in the dashboard.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// A mapping of all the dashboards created by the custom resource using the
	// unique ID (`uid`) of the dashboard as the key.
	Dashboards map[string]DashboardInfo `json:"dashboards,omitempty"`
}

// Contains a list of information about the connection of a dashboard created in Grafana.
type DashboardInfo struct {
	// The unique ID of the dashboard.
	UID string `json:"uid,omitempty"`

	// The path to the dashboard.
	URL string `json:"url,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=false,oc=mon
// Defines the Schema for the Dashboards API.
// +genclient
type Dashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DashboardSpec   `json:"spec,omitempty"`
	Status DashboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of dashboards in Grafana.
type DashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dashboard{}, &DashboardList{})
}
