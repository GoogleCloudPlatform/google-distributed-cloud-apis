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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// `NotebookSpec` defines the desired state of the Notebook.
type NotebookSpec struct {
	// A specification for a Notebook instance that describes the desired state of the Notebook's workload in the cluster.
	Template NotebookTemplateSpec `json:"template,omitempty"`
}

// `NotebookTemplateSpec` defines the desired state of the cluster.
type NotebookTemplateSpec struct {
	// Defines the containers, environment variables for the container,
	// and other properties such as the scheduler name and security context.
	// For more information, refer to https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#podspec-v1-core.
	Spec corev1.PodSpec `json:"spec,omitempty"`
	// <code>UserCluster</code> is the name of the user cluster where the Notebook workload is created. The cluster
	// must have sufficient resources to run the workload specified in <code>PodSpec</code>.
	UserCluster string `json:"userCluster,omitempty"`
}

// `NotebookError` represents an error that occurred while updating Notebook resources.
type NotebookError struct {
	// A string that contains a detailed message that describes the <code>Pod</code> state.
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`
	// A string that contains the error code of the <code>Pod</code> in its current state, such as <code>VTXWB-E0000</code>.
	// +optional
	Code string `json:"code,omitempty"`
}

// `NotebookStatus` defines the observed state of the Notebook.
type NotebookStatus struct {
	// Deprecated, use <code>Notebook</code> CR <code>Events</code>. An array of <code>NotebookCondition</code> objects.
	// TODO(b/430501986): Replace NotebookCondition with metav1.Condition.
	// +optional
	Conditions []NotebookCondition `json:"conditions,omitempty"`
	// The number of <code>Pods</code> created by the <code>StatefulSet</code> controller with their <code>Condition</code> set to <code>Ready</code>.
	ReadyReplicas int32 `json:"readyReplicas"`
	// Deprecated, use <code>NotebookStatus.Phase</code>. The state of the underlying container. For more information,
	// refer to https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#containerstate-v1-core.
	// +optional
	ContainerState corev1.ContainerState `json:"containerState"`
	// A string that contains the current <code>Pod</code> phase.
	// +optional
	Phase string `json:"phase,omitempty"`
	// A string that contains the reason the <code>Pod</code> is in its current state.
	// +optional
	Reason string `json:"reason,omitempty"`
	// If the Notebook is not currently running or the controller failed to query pod status, the time when the Notebook pod was first observed to not be running.
	// +optional
	PodUnavailableTime metav1.Time `json:"podUnavailableTime,omitempty"`
	// The last error to occur while updating Notebook resources. This field is empty on success.
	// +optional
	Error NotebookError `json:"error,omitempty"`
}

type NotebookCondition struct {
	// The type of the condition. There are three possible values, <code>Running</code>, <code>Waiting</code>, and <code>Terminated</code>.
	Type string `json:"type"`
	// The time when the Notebook was last probed.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// A string that contains the reason the container is in its current state.
	// +optional
	Reason string `json:"reason,omitempty"`
	// A string that contains a detailed message the container is in its current state.
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +gdcloud:manifest:relevant=false,oc=vtxwb
// +genclient
// `Notebook` is the Schema for the Notebooks API.
type Notebook struct {
	metav1.TypeMeta `json:",inline"`
	// Metadata that all persisted resources must have. Refer to
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta
	// in the Kubernetes API documentation for fields of `metadata`.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotebookSpec   `json:"spec,omitempty"`
	Status NotebookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// `NotebookList` contains a list of `Notebook` instances.
type NotebookList struct {
	// Refer to Kubernetes API documentation to learn about <code>metadata</code> fields,
	metav1.TypeMeta `json:",inline"`
	// Describes metadata that synthetic resources must have, including lists and various status objects. Refer to
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#listmeta-v1-meta
	// in the Kubernetes API documentation for fields of `metadata`.
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notebook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Notebook{}, &NotebookList{})
}
