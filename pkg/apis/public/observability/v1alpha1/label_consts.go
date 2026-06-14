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

const (
	// AuditLogDestination identifies intended audience for K8s API Server logs
	// where requestURI field contains a namespace with such label.
	// Valid values for this label are:
	// apis/public/logging/v1alpha1/LogAccessLevel.IO"
	// apis/public/logging/v1alpha1/LogAccessLevel.PA"
	// If label is not specified or contains invalid value it defaults to "io"
	// Users should not change this label's value.
	AuditLogDestination = Group + "/auditlog-destination"
)
