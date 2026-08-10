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
	// InstalledStatusType is the condition type that indicates Observability
	// stack has been installed or upgraded.
	InstalledStatusType = "Installed"

	// ReadyStatusType is the condition type that indicates Observability stack
	// is healthy.
	ReadyStatusType = "Ready"
)

const (
	// InstalledReasonSuccess is the condition reason that indicates the
	// Installed status being true because of a successful installation.
	InstalledReasonSuccess = "InstallationSuccess"

	// InstalledReasonFailure is the condition reason that indicates the
	// Installed status being false because of an installation failure.
	InstalledReasonFailure = "InstallationFailure"

	// InstalledReasonUpgrading is the condition reason that indicates the
	// Installed status being false because of an upgrade is in process.
	InstalledReasonUpgrading = "Upgrading"

	// InstalledReasonLegacy is the condition reason that indicates the
	// Installed status being true because of legacy OSS stack is present.
	InstalledReasonLegacy = "Legacy"

	// ReadyReasonHealthy is the condition reason that indicates the Ready
	// status being true because of an successful health check.
	ReadyReasonHealthy = "Healthy"

	// ReadyReasonNonupgradable is the condition reason that indicates the Ready
	// status being false because upgrade from current version is not supported.
	ReadyReasonNonupgradable = "UpgradeNotSupported"

	// ReadyReasonUpgradeFailure is the condition reason that indicates the
	// Ready status being false because of an unsuccessful upgrade.
	ReadyReasonUpgradeFailure = "UpgradeFailure"

	// ReadyReasonHealthCheckFailure is the condition reason that indicates the
	// Ready status being false/unknown because of a failure in health check.
	ReadyReasonHealthCheckFailure = "HealthCheckFailure"
)
