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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// HelmPausedAnnotation is an annotation that can be applied to HarborInstance
	// object.
	// HaaS HarborInstance controllers will not do any helm reconciliation
	// on those objects with this annotations.
	HelmPausedAnnotation = "artifactregistry.private.gdc.goog/helmPaused"
)

// HasPaused returns true if the object has the `paused` annotation.
func HasHelmPaused(o metav1.ObjectMeta) bool {
	if metav1.HasAnnotation(o, HelmPausedAnnotation) {
		return o.Annotations[HelmPausedAnnotation] == "true"
	}
	return false
}
