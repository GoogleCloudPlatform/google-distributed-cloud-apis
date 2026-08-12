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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ConditionTypeReady is the condition type that indicates the overall readiness of
	// a resource. If ConditionTypeReady status is set as True, it indicates that this
	// resource has been reconciled correctly and the corresponding managed binding has
	// been created.
	ConditionTypeReady = "Ready"

	// ConditionTypeTier is the condition type that indicates the resource tier of the
	// ManagedSIEM instance.
	ConditionTypeTier = "Tier"

	// ConditionUnknown is the condition status that indicates the expected condition type is not present in the conditions list.
	ConditionUnknown = "Unknown"
	// ConditionReady is the condition status that indicates the ready condition type status is true.
	ConditionReady = "Ready"
)

func ReadyStatus(conditions []metav1.Condition) string {
	for _, condition := range conditions {
		if condition.Type == ConditionTypeReady {
			if condition.Status == metav1.ConditionTrue {
				return ConditionReady
			}
			return condition.Reason
		}
	}
	return ConditionUnknown
}

func TierStatus(conditions []metav1.Condition) string {
	for _, condition := range conditions {
		if condition.Type == ConditionTypeTier {
			return condition.Reason
		}
	}
	return ConditionUnknown
}
