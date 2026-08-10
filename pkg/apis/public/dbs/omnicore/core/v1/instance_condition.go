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

	// DatabaseParameterUpdated
	// true when the following conditions are met
	//    - Database Parameters in instance spec match the ones being used by the
	//        running database
	//    - Database Parameters in instance spec were written to database config without (validation) errors, but database could not restart with the applied parameters. In this case the user will have to manually change back parameters if they want the database to start.
	//    - Database Parameters in instance spec failed, had errors (e.g, validation errors) and all the parameters were successfully restored to the ones before the change.
	//  Is False if:
	//    - Attempting to restart the database after a parameter update
	//    - Attempting to restore the parameters after a failed update
	DatabaseParameterUpdated ConditionType = "DatabaseParameterUpdated"

	// DatabaseProvisioned
	// This condition indicates that the Instance has finished its initial provisioning process. It turns True when the DatabaseReady condition on the Instance turns True; It will not turn to False afterward - except when the status is deleted such as during backup/restore.
	DatabaseProvisioned ConditionType = "DatabaseProvisioned"

	// DatabaseReady indicates whether the database is accessible.
	//	True when the database process on this instance is expected to be running. An internal admin user
	//    should be able to connect to this database process and execute a valid read/write query. False if the process is not running, or is unable to execute an otherwise valid query (e.g. if it is unable to commit transactions due to running out of disk space). This is currently owned by the instanceReconciler, as well as the health check prober.
	//  Standby:
	//  True when the database process on this instance is expected to be running. If a standby is expected to be a ready standby, it should be readable. False if the database process is not running, or is unable to replicate data from the primary. This is currently owned by the instanceReconciler, as well as the health check prober.
	DatabaseReady ConditionType = "DatabaseReady"

	// DBDaemonReady indicates whether the dbdaamon is ready.
	// True when the instance controller is able to establish a connection to the DBDaemon of this instance. Note that DBDaemon connectivity does not mean the DBDaemon is able to necessarily connect to the database process. False when the instance controller is not able to connect to the DBDaemon.
	DBDaemonReady ConditionType = "DBDaemonReady"

	// NodeManagerReady indicates whether the omnisvc NodeManager is ready.
	NodeManagerReady ConditionType = "NodeManagerReady"

	// InstanceAdminPasswordReady
	// True when this instance has had a set user's admin password
	// request and it has set the user password correctly. False when
	// this operation has not been completed or had an error.
	// Standby:
	//   On standbys, this value is set to true when it is true on the
	//   primary as that value is replicated from the primary.
	//
	InstanceAdminPasswordReady ConditionType = "AdminPasswordReady"

	// InstanceHAReady indicates whether HA is running successfully.
	// 	True for a primary instance when it has been set up to receive connections from a standby. This condition is primarily used to capture the workflow state for an HA instance, and is meant to be used by DBS internally rather than by the user. It can transition to false if we need to re-run setup steps. This is owned by the instance reconciler (under the HA workflow component).
	//  Standby:
	//	True for a standby when it has been configured to be run as a standby and has successfully replicated initial data from the primary. This condition is meant to capture the workflow state for an HA instance and is meant to be used by DBS internally rather than by the user. It can transition to false if we need to re-run setup steps. This is owned by the instance reconciler (under the HA workflow component).
	InstanceHAReady ConditionType = "HAReady"

	// MonitoringReady
	// True when the monitoring pod is created and is in a ready state. False otherwise.
	MonitoringReady ConditionType = "MonitoringReady"

	RolledBack ConditionType = "RolledBack"

	// Upgraded indicates whether an instance is upgraded or not
	Upgraded ConditionType = "Upgraded"

	// MLExtensionReady
	// True when instance has ML Volume mount and ML is enabled. False otherwise.
	MLExtensionReady ConditionType = "MLExtensionReady"

	OtelContainerInjection ConditionType = "OtelContainerInjected"
)

const (
	PatchingImage                     ConditionReason = "PatchingImage"
	PatchingDatabase                  ConditionReason = "PatchingDatabase"
	DatabasePodsStarting              ConditionReason = "PodsStarting"
	DatabasePodsStarted               ConditionReason = "Ready"
	DatabasePodResizing               ConditionReason = "Resizing"
	DatabasePodsStopping              ConditionReason = "PodsStopping"
	DatabasePodsStopped               ConditionReason = "PodsStopped"
	DatabasePodsStoppingSkipped       ConditionReason = "PodsStoppingSkipped"
	DatabaseStopped                   ConditionReason = "DatabaseStopped"
	DatabaseNotRunning                ConditionReason = "DatabaseNotRunning"
	DatabaseError                     ConditionReason = "Error"
	DatabaseProvisioning              ConditionReason = "Provisioning"
	HealthCheckFailed                 ConditionReason = "HealthCheckFailed"
	RestartFailedAfterParameterUpdate ConditionReason = "RestartFailedAfterParameterUpdate"
	UpdatingParameters                ConditionReason = "UpdatingParameters"

	MonitoringWaitingForInternalConnectivity ConditionReason = "WaitingForInternalConnectivity"

	MonitoringPodStarting ConditionReason = "MonitoringPodStarting"
	MonitoringPodStarted  ConditionReason = "MonitoringPodStarted"
	MonitoringPodStopping ConditionReason = "MonitoringPodStopping"
	MonitoringPodStopped  ConditionReason = "MonitoringPodStopped"

	MonitorStarted ConditionReason = "MonitorStarted"
	MonitorStopped ConditionReason = "MonitorStopped"
	MonitorError   ConditionReason = "MonitorError"

	DatabaseRecoveryBootstrapInProgress ConditionReason = "RecoveryBootstrapInProgress"
	DatabasePendingRecovery             ConditionReason = "PendingRecovery"

	ServiceDisabled ConditionReason = "ServiceDisabled"

	L2CaCertCopyingInProgress             ConditionReason = "L2CaCertCopyingInProgress"
	DBDaemonCertificateCreationInProgress ConditionReason = "DBDaemonCertCreationInProgress"
	DBDaemonNotAccessible                 ConditionReason = "DBDaemonNotAccessible"
	DBDaemonConnectionReady               ConditionReason = "DBDaemonConnectionReady"

	NodeManagerWaitingForCertificate ConditionReason = "WaitingForCertificate"
	NodeManagerWaitingForService     ConditionReason = "WaitingForService"
	NodeManagerWaitingForDBDaemon    ConditionReason = "WaitingForDBDaemon"
	NodeManagerPodsStarting          ConditionReason = "PodsStarting"
	NodeManagerError                 ConditionReason = "Error"
	NodeManagerNotEnabled            ConditionReason = "NotEnabled"
	NodeManagerIsReady               ConditionReason = "Ready"

	AdminPasswordApplied     ConditionReason = "Ready"
	ReconcilingAdminPassword ConditionReason = "InstanceReconciling"

	HASetupInProgress      ConditionReason = "SetupInProgress"
	HAFailoverInProgress   ConditionReason = "FailoverInProgress"
	HASwitchoverInProgress ConditionReason = "SwitchoverInProgress"
	HAReplicationError     ConditionReason = "ReplicationError"
	HASetupDone            ConditionReason = "Ready"
	HAPrimaryUnhealthy     ConditionReason = "PrimaryUnhealthy"
	HAStandbyUnhealthy     ConditionReason = "StandbyUnhealthy"

	UpgradeStarted                  ConditionReason = "UpgradeStarted"
	RollbackStarted                 ConditionReason = "RollbackStarted"
	PreUpgradeBackupCreated         ConditionReason = "PreUpgradeBackupCreated"
	PreUpgradeBackupCreationSkipped ConditionReason = "PreUpgradeBackupCreationSkipped"
	PreUpgradeBackupCreationFailed  ConditionReason = "PreUpgradeBackupCreationFailed"
	StsStartedAndUpdated            ConditionReason = "StatefulSetStartedAndUpdated"

	StsStarted           ConditionReason = "StatefulSetStarted"
	DataRestored         ConditionReason = "DataRestored"
	DataRestoreSkipped   ConditionReason = "DataRestoreSkipped"
	DataRestoreFailed    ConditionReason = "DataRestoreFailed"
	DataRestoreTimeout   ConditionReason = "DataRestoreTimeout"
	ImagePatchingTimeout ConditionReason = "ImagePatchingTimeout"

	StsStoppingTimeout               ConditionReason = "StatefulSetStoppingTimeout"
	DataPatched                      ConditionReason = "DataPatched"
	DataPatchingTimeout              ConditionReason = "DataPatchingTimeout"
	DataPatchingFailed               ConditionReason = "DataPatchingFailed"
	PreUppgradeBackupCreationTimeout ConditionReason = "PreUppgradeBackupCreationTimeout"

	RestoringDatabaseParameters               ConditionReason = "RestoringDatabaseParameters"
	DatabaseParameterUpdateCompleted          ConditionReason = "DatabaseParameterUpdateCompleted"
	DatabaseCannotRestartAfterParameterUpdate ConditionReason = "DatabaseCannotRestartAfterParameterUpdate"
	RestartingDatabase                        ConditionReason = "RestartingDatabase"

	MLExtensionInProgress ConditionReason = "MLExtensionInProgress"
	MLExtensionEnabled    ConditionReason = "MLExtensionEnabled"

	OtelContainerIsInjected ConditionReason = "Injected"
)
