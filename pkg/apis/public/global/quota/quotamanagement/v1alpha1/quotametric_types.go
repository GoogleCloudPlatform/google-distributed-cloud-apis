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

// Represents the quota metric that the quota defined against.
type Metric struct {
	// The name that identifies the quota metric. e.g. compute.googleapis.com/tpus
	Name string `json:"name"`
	// A concise name for the metric, which can be displayed in user interfaces.
	DisplayName string `json:"displayName"`
}
