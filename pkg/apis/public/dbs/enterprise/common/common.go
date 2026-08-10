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

package common

const (
	AnnotationLastKnownPrimaryInstanceName = "dbs.internal.dbadmin.goog/lastPrimaryInstanceName"
	AnnotationLastKnownPrimaryZone         = "dbs.internal.dbadmin.goog/lastKnownPrimaryZone"

	AnnotationForceRaaSMigration = "dbs.internal.dbadmin.goog/migrateToRaaS"

	AnnotationDisableHealthcheck = "dbs.internal.dbadmin.goog/disableHealthcheck"
	AnnotationHealthcheckBlocker = "dbs.internal.dbadmin.goog/healthcheckBlocker"

	AnnotationInitialHASetup = "dbs.internal.dbadmin.goog/initialHASetup"

	// AnnotationOrphanedPrimaryInstances keeps track of the orphaned primary instances that
	// cannot be deleted during failover, e.g., due to the primary zone down. It's supposed to
	// be updated by the failover controller (when old primary cannot be deleted within timeout)
	// and the DBCluster controller (when the orphaned primary instances become reachable for deletion).
	AnnotationOrphanedPrimaryInstances = "dbs.internal.dbadmin.goog/orphanedPrimaryInstances"

	AnnotationInitialHASetupAfterFailover   = "failover"
	AnnotationInitialHASetupAfterSwitchover = "switchover"
	AnnotationInitialHASetupStandbyCreation = "create"

	// AnnotationPreSuccessStartTime indicates the timestamp when the switchover's internal
	// phase switch to PreSuccess phase
	// TODO(b/430108661) refactor the annotation to be a field in switchover CR.
	AnnotationPreSuccessStartTime = "switchover.dbadmin.goog/preSuccessStartTime"

	DBPortName = "db"

	// Deprecated: Use InstanceHealth object instead
	// TODO(b/503779273): these constants should be deleted after migrating new health check controls are complete
	AnnotationHealthcheckFailures    = "dbs.internal.dbadmin.goog/consecutiveHealthcheckFailures"
	AnnotationLastHealthcheckRunTime = "dbs.internal.dbadmin.goog/lastHealthcheckRunTime"
)

type DBClusterRef string

type BackupPlanRef string

type BackupRef string
