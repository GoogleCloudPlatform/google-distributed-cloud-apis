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
	LabelMigrationDBCluster      = "migration.dbadmin.goog/dbcluster"
	LabelMigrationExternalServer = "migration.dbadmin.goog/externalserver"
	LabelDBClusterMigration      = "dbcluster.dbadmin.goog/migration"
)

type MigrationControl string

const (
	MigrationControlSetup   MigrationControl = "setup"
	MigrationControlStart   MigrationControl = "start"
	MigrationControlStop    MigrationControl = "stop"
	MigrationControlPromote MigrationControl = "promote"

	ConditionTypeMigrationComplete = "MigrationComplete"

	ConditionReasonMigrationUnknown  ConditionReason = "Unknown"
	ConditionReasonMigrationError    ConditionReason = "Error"
	ConditionReasonMigrationUnsynced ConditionReason = "Unsynced"
	ConditionReasonMigrationSyncing  ConditionReason = "Syncing"
	ConditionReasonMigrationRunning  ConditionReason = "Running"
	ConditionReasonMigrationStopped  ConditionReason = "Stopped"
	ConditionReasonMigrationPromoted ConditionReason = "Promoted"

	ConditionTypeMigrationSetupLROComplete     = "MigrationSetupLROComplete"
	ConditionReasonMigrationSetupLROUnknown    = "Unknown"
	ConditionReasonMigrationSetupLROInProgress = "InProgress"
	ConditionReasonMigrationSetupLROComplete   = "Complete"
	ConditionReasonMigrationSetupLROFailed     = "Failed"
)
