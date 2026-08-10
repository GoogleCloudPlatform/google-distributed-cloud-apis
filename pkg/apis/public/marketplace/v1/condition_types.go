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

type ServiceInstallationCondition string

const (
	// The condition type that indicates the deployment of a `ServiceInstance` is
	// in progress.
	DeployingConditionType ServiceInstallationCondition = "Deploying"
)

type ServiceInstallationReason string

const (
	// The condition reason for a successful uninstallation, indicated by a
	// `Deploying` status of `false`.
	DeployingStatusReasonUninstallSuccess ServiceInstallationReason = "UninstallSuccess"

	// The condition reason for a successful deployment, indicated by a
	// `Deploying` status of `false`.
	DeployingStatusReasonSuccess ServiceInstallationReason = "DeploymentSuccess"

	// The condition reason for the deployment being retried due to a failure,
	// indicated by a `Deploying` status of `false`.
	DeployingStatusReasonFailure ServiceInstallationReason = "DeploymentFailure"

	// The condition reason for an in-progress deployment, indicated by a
	// `Deploying` status of `true`.
	DeployingStatusReasonPending ServiceInstallationReason = "DeploymentPending"
)
