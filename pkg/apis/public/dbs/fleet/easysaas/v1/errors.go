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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Error struct {
	// The code for the error. It includes a prefix with letters followed by
	// a four-digit numeric code.
	// +kubebuilder:validation:Pattern:=`^[A-Z]{2,4}[0-9]{4}$`
	Code string `json:"code"`

	// The human-readable error message.
	Message string `json:"message"`
}

type ErrorStatus struct {
	Errors []Error `json:"errors,omitempty"`

	LastUpdate metav1.Time `json:"lastUpdate,omitempty"`
}

type ErrorCode string

const (
	ModuleInstanceReconcileError ErrorCode = "EZRE0001"
	LocalRolloutReconcileError   ErrorCode = "EZRE0002"
)

func NewError(code ErrorCode, message string) Error {
	return Error{
		Code:    string(code),
		Message: message,
	}
}

func (err Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}
