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
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	LabelDBC          = "dbs.internal.dbadmin.goog/dbc"
	MdDbc             = "dbc"
	MdPrimaryInstance = "primaryInstance"
	MdStandbyInstance = "standbyInstance"
)

// +kubebuilder:object:generate=true
type WorkflowSpec struct {
	// Metadata is intended to allow different workflows to attach data needed to
	// execute this workflow, e.g. which DBC/instance/ other resource this is
	// attached to
	Metadata map[string]string `json:"metadata,omitempty"`

	// CurrentStep is the current step of the workflow
	// +optional
	CurrentStep string `json:"currentStep,omitempty"`

	// Attempt allows the workflow to smartly retry and choose to fail if too many
	// retries have occurred
	// +kubebuilder:default:=0
	Attempt int `json:"attempt"`

	// StartTime is when the workflow began
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// EndTime is when the workflow has reached a terminal state
	// +optional
	EndTime *metav1.Time `json:"endTime,omitempty"`

	// CurrentStepTime allows us to see when the current step was initiated which
	// allows us to time out at the step level.
	// +optional
	CurrentStepTime *metav1.Time `json:"currentStepTime,omitempty"`

	// Cleanup is used to mark this object as safe for deletion.
	// +kubebuilder:default:=false
	Cleanup bool `json:"cleanup"`

	// RequeueTime if set, then tells the reconciler to requeue this job to run
	// at the specified time
	// +optional
	RequeueTime *metav1.Time `json:"requeueTime,omitempty"`
}
type CommonJob interface {
	occoreapi.Entity
	WorkflowSpec() *WorkflowSpec
}

func (w *WorkflowSpec) InitMetadata(dbc, primaryInst, standbyInst string) {
	if w.Metadata == nil {
		w.Metadata = map[string]string{}
	}
	w.Metadata[MdDbc] = dbc
	w.Metadata[MdPrimaryInstance] = primaryInst
	w.Metadata[MdStandbyInstance] = standbyInst
}
