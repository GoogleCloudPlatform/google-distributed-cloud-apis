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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="ProjectCreatedStatus",type="string",JSONPath=".status.conditions[?(@.type=='ProjectCreated')].reason"
// +kubebuilder:printcolumn:name="ProjectSubjectAssignedStatus",type="string",JSONPath=".status.conditions[?(@.type=='ProjectSubjectAssigned')].reason"
// +kubebuilder:printcolumn:name="ProjectDeletedStatus",type="string",JSONPath=".status.conditions[?(@.type=='ProjectDeleted')].reason"
// +gdcloud:manifest:relevant=true,oc=haas,component=harbor
// +gdcloud:manifest:entities="harbor-projects",verbs=create;describe
// +gdcloud:manifest:rbac="create,describe:harbor-instance-admin;harbor-project-creator"
// +gdcloud:manifest:skipcodegen=true
// Represents a harbor project in a harbor instance.
// A custom resource establishes the expectation that a project must exist.
// Namespace is the GDCH project name that the harbor instance and its harbor projects belong to.
type HarborInstanceProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   HarborInstanceProjectSpec   `json:"spec"`
	Status HarborInstanceProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `HarborInstanceProject` resources.
type HarborInstanceProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []HarborInstanceProject `json:"items"`
}

// Defines the specification or expected state of the `HarborProject` object.
type HarborInstanceProjectSpec struct {
	// The name of the harbor project. Has to match harbor's naming rules.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^[a-z0-9]+(?:[._-][a-z0-9]+)*$"
	// +kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	ProjectName string `json:"projectName"`
	// The harbor instance that the harbor project belongs to.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	HarborInstanceRef corev1.LocalObjectReference `json:"harborInstanceRef"`
	// The user or group that creates the harbor project,
	// and the subject will be granted as the first harbor Project Admin
	// to manage the harbor project, and grant more users access, at harbor UI.
	// It only represents individual users for now.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	CreatorSubject *rbacv1.Subject `json:"creatorSubject,omitempty"`
	// Whether a project will scan images automatically on push.
	// Defaults to false if not specified.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// +kubebuilder:default=false
	AutoScan *bool `json:"autoScan,omitempty"`
}

const (
	// The condition type that indicates whether a harbor project has been created.
	HarborInstanceProjectCreatedStatusType = "ProjectCreated"
	// The condition type that indicates whether a harbor project subject has been assigned.
	HarborInstanceProjectSubjectAssignedStatusType = "ProjectSubjectAssigned"
	// The condition type that indicates whether a harbor project has been deleted.
	HarborInstanceProjectDeletedStatusType = "ProjectDeleted"
	// The condition type that indicates whether a project has been exempted from signature verification.
	HarborInstanceProjectSignatureVerificationExemptedStatusType = "ProjectSignatureVerificationExempted"

	// The condition reason that indicates that a project has been successfully created.
	HarborInstanceProjectCreatedReason = "Created"
	// The condition reason that indicates that an error occurred at creation.
	HarborInstanceProjectCreateErrorReason = "CreateError"
	// The condition reason that indicates that the subject is assigned.
	HarborInstanceProjectSubjectAssignedReason = "SubjectAssigned"
	// The condition reason that indicates that an error occurred during the assignment of the subject.
	HarborInstanceProjectSubjectAssignErrorReason = "SubjectAssignError"
	// The condition reason that indicates that an error occurred at deletion.
	HarborInstanceProjectDeleteErrorReason = "DeleteError"
	// The condition reason that indicates that an error occurred during enabling signature verification exemption for a project.
	HarborInstanceProjectSignatureVerificationExemptionErrorReason = "SignatureVerificationExemptionError"
	// The condition reason that indicates that signature verification exemption is enabled for a project.
	HarborInstanceProjectSignatureVerificationExemptionReason = "SignatureVerificationExemptionEnabled"
)

// Defines the observed state of the `HarborInstanceProject` object.
type HarborInstanceProjectStatus struct {
	// Conditions include `ProjectCreated`, `ProjectDeleted` and `ProjectSubjectAssigned`
	// `ProjectCreated` means that the harbor project is created in the harbor instance with true or false status.
	// `ProjectDeleted` means that the harbor project is deleted in the harbor instance with true or false status.
	// `ProjectSubjectAssigned` means that the user or group who initiates the custom resource creation
	// is granted as first Harbor Admin role in harbor, with true or false status.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(
		&HarborInstanceProject{},
		&HarborInstanceProjectList{},
	)
}
