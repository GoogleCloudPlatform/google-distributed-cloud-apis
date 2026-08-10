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

const (
	// GlobalNamespaceReady indicates if the Project's associated Namespace is reconciled.
	GlobalNamespaceReady = "Namespace"
	// GlobalNamespaceCleanup is the condition type that indicates that the Project is being
	// deleted, and the associated namespace is deleted.
	GlobalNamespaceCleanup = "NamespaceCleanup"

	// ReconciledReason indicates the resource is reconciled successfully.
	ReconciledReason = "Reconciled"

	// TerminatingReason indicates the resource is being deleted.
	TerminatingReason = "Terminating"

	// KubernetesClientErrorReason indicates an error is returned when the k8s client
	// performs read or write action.
	KubernetesClientErrorReason = "KubernetesClientError"
)
