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

// Defines the strategy for locking the objects of the Bucket.
type LockingPolicy struct {
	// Specifies the minimum number of days that each version of every object will be retained. An object cannot be deleted during the retention period. If a bucket contains any object, it cannot be deleted either.
	// When unspecified, no default object retention period is set.
	// Can be modified after creation, however the change will only take effect for new objects and versions. Existing objects and versions will still use the previous value.
	// +optional
	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:validation:Maximum:=36500
	DefaultObjectRetentionDays *int32 `json:"defaultObjectRetentionDays,omitempty"`
}
