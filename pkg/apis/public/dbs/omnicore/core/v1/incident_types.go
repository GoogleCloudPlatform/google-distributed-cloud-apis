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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// CriticalIncident contains all information about an ongoing critical incident.
// +kubebuilder:object:generate=true
type CriticalIncident struct {
	// Code is the error code of this particular error.
	// Error codes are DBSE+numeric strings, like "DBSE1012".
	Code string `json:"code"`
	// Message describes the incident/error that occurred.
	Message string `json:"message,omitempty"`
	// Resource contains information about the Database Service component that reported the incident
	// as well as about the K8s resource.
	Resource CriticalIncidentResource `json:"resource"`
	// StackTrace contains an unstructured list of messages from the stack trace.
	// +optional
	StackTrace []CriticalIncidentStackTraceMessage `json:"stackTrace,omitempty"`
	// MessageTemplateParams contains key-value pairs necessary for generating
	// a user-friendly data-driven version of Message in the UI.
	// +optional
	MessageTemplateParams map[string]string `json:"messageTemplateParams,omitempty"`
	// CreateTime is the timestamp when this Incident was created at the origin.
	CreateTime *metav1.Time `json:"createTime,"`
	// TransientUntil if present indicates that the issue should be considered transient until the specified time.
	// +optional
	TransientUntil *metav1.Time `json:"transientUntil,omitempty"`
}

// Location allows to find the originating k8s resource.
// Empty for internal incidents.
type Location struct {
	// Cluster is the name of the cluster of the affected K8S resource.
	// +optional
	Cluster string `json:"cluster,omitempty"`
	// Group is the Group name of the k8s resource.
	// +optional
	Group string `json:"group,omitempty"`
	// Version is the Version of the k8s resource.
	// +optional
	Version string `json:"version,omitempty"`
	// Kind is the Kind of the k8s resource.
	// +optional
	Kind string `json:"kind,omitempty"`
	// Namespace is the namespace of the affected K8S resource.
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// Name is the name of the affected K8S resource.
	// +optional
	Name string `json:"name,omitempty"`
}

// CriticalIncidentResource helps to locate the source of the critical incident.
type CriticalIncidentResource struct {
	// Component is an internal identifier of the Database Service subsystem that reported the incident.
	Component string `json:"component,"`
	// Location
	// +optional
	Location Location `json:"location,omitempty"`
}

// CriticalIncidentStackTraceMessage contains stack trace information
// available for the incident.
type CriticalIncidentStackTraceMessage struct {
	// Component is the name of a Database Service component that logged the message.
	// +optional
	Component string `json:"component,omitempty"`
	// Logged message.
	// +optional
	Message string `json:"message,omitempty"`
}
