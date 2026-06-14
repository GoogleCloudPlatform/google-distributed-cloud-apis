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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:XValidation:rule="has(self.duration) ? has(self.startTime) : true",message="startTime must be set if duration is set"
// +kubebuilder:validation:XValidation:rule="has(self.endTime) ? (has(self.startTime) && has(self.duration)) : true",message="startTime and duration must both be set if endTime is set"

// Timer is a field to capture the wall-clock time of events.
type Timer struct {
	// StartTime is when the timer counting started.
	StartTime *metav1.MicroTime `json:"startTime,omitempty"`

	// EndTime is when the timer counting ended (if it has ended).
	// If EndTime is set, then StartTime and Duration must also be set.
	EndTime *metav1.MicroTime `json:"endTime,omitempty"`

	// Duration stores the current duration of the timer.
	// If Duration is set, then StartTime must also be set.
	Duration *metav1.Duration `json:"duration,omitempty"`

	// Estimate is an estimate for the current timer.
	Estimate *metav1.Duration `json:"estimate,omitempty"`
}

// Start starts the timer.
//
// *idempotent* If the timer is already started, does nothing.
func (t *Timer) Start() {
	if t != nil && t.StartTime == nil {
		now := metav1.NowMicro()
		t.StartTime = &now
	}
}

// End stops the timer and returns the final duration.
//
// *idempotent* If the timer has already ended, just returns the final duration.
func (t *Timer) End() time.Duration {
	if t != nil && t.EndTime == nil {
		now := metav1.NowMicro()
		t.EndTime = &now
	}
	return t.Tick()
}

// Tick updates the timer's duration and returns it.
func (t *Timer) Tick() time.Duration {
	if t == nil || t.StartTime == nil {
		return 0
	}
	startTime := t.StartTime.Time
	endTime := time.Now()
	if t.EndTime != nil {
		endTime = t.EndTime.Time
	}
	currentDuration := endTime.Sub(startTime)
	t.Duration = &metav1.Duration{Duration: currentDuration}
	return currentDuration
}

// Overdue returns if, compared to the timer's estimate, the event is overdue.
func (t *Timer) IsOverdue() bool {
	return t != nil && t.Estimate != nil && t.Tick() > t.Estimate.Duration
}
