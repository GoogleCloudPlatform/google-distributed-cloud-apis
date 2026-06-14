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
	// AdminPasswordReady
	// True when the primary instance has successfully set the user's
	// admin user password to the requested value. False when this
	// operation has not been completed or had an error.
	//
	AdminPasswordReady ConditionType = "AdminPasswordReady"

	// DBCUpgradeInProgress indicates the cluster is currently upgrading.
	DBCUpgradeInProgress ConditionType = "DBCUpgradeInProgress"

	// LDTM indicates the cluster is currently performing some LDTM operation.
	LDTMScalingInProgress ConditionType = "LDTMScalingInProgress"

	// LDTM indicates the cluster is currently performing some LDTM operation.
	LDTMMinorVersionUpgradeInProgress ConditionType = "LDTMMinorVersionUpgradeInProgress"
	// HAReady
	// True when a DBCluster has HA enabled, and it has successfully set up a standby instance. It can transition to false in the following cases: HA is disabled, after a failover while a standby is being recreated, or as part of auto-healing if we need to recreate or fix a bad standby. This is owned by the dbcluster reconciler (under the HA workflow component).
	HAReady ConditionType = "HAReady"

	// MaintenanceWindowInProgress indicates the cluster is currently in a maintenance window.
	MaintenanceWindowInProgress ConditionType = "MaintenanceWindowInProgress"

	// Provisioned
	// This condition indicates that the DBCluster has finished its initial provisioning process. It turns True when the DatabaseReady condition on the DBCluster turns True; It will not turn to False afterward - except when the status is deleted such as during backup/restore.
	Provisioned ConditionType = "Provisioned"

	// K8ServiceReady
	// True when the DBC has internal and external (if required) load balancers created and they have IPs assigned to them. False otherwise.
	K8ServiceReady ConditionType = "K8ServiceReady"

	// DNSReady
	// True when the DBC has a DNS FQDN. False Otherwise. Applicable only to Lancer Evo.
	DNSReady ConditionType = "DNSReady"

	// LoadBalancerReady
	// True when the DBC has a Load Balancer setup and healthy on all nodes
	LoadBalancerReady ConditionType = "LoadBalancerReady"

	// VIPManagerReady
	// True when the DBC has a VIP manager is setup and vip is assigned correctly
	VIPManagerReady ConditionType = "VIPManagerReady"

	// EndpointsReady
	// True when the DBC has a Read-Write and Read-Only (if enabled) endpoints assigned
	EndpointsReady ConditionType = "EndpointsReady"

	// SIEDeleted
	// True when the ServiceIsolationEnvironment (SIE) is deleted. False otherwise.
	SIEDeleted ConditionType = "SIEDeleted"

	// L1ShadowNamespaceDeleted
	// True when the shadow namespace in L1 is deleted. False otherwise.
	L1ShadowNamespaceDeleted ConditionType = "L1ShadowNamespaceDeleted"

	// L2ShadowNamespaceDeleted
	// True when the shadow namespace in L2 is deleted. False otherwise.
	L2ShadowNamespaceDeleted ConditionType = "L2ShadowNamespaceDeleted"
)

// Condition reason for ConditionType=Provisioned on database cluster status condition
const (
	DBClusterCreated  ConditionReason = "DBClusterCreated"
	SASetupInProgress ConditionReason = "ServiceAccountSetupInProgress"
	SASetupDone       ConditionReason = "ServiceAccountSetupDone"
)

// Condition reason for ConditionType=SIEDeleted on database cluster status condition
const (
	SIEDeletionInProgress ConditionReason = "SIEDeletionInProgress"
	SIEDeletionComplete   ConditionReason = "SIEDeletionComplete"
)

// Condition reason for ConditionType=L1ShadowNamespaceDeleted on database cluster status condition
const (
	L1ShadowNamespaceDeletionInProgress ConditionReason = "L1ShadowNamespaceDeletionInProgress"
	L1ShadowNamespaceDeletionComplete   ConditionReason = "L1ShadowNamespaceDeletionComplete"
)

// Condition reason for ConditionType=L2ShadowNamespaceDeleted on database cluster status condition
const (
	L2ShadowNamespaceDeletionInProgress ConditionReason = "L2ShadowNamespaceDeletionInProgress"
	L2ShadowNamespaceDeletionComplete   ConditionReason = "L2ShadowNamespaceDeletionComplete"
)

// Condition reason for ConditionType=DBCUpgradeInProgress on database cluster status condition
const (
	LDTMStart     ConditionReason = "LDTMStarted"
	LDTMFailed    ConditionReason = "LDTMFailed"
	LDTMSucceeded ConditionReason = "LDTMSucceeded"

	DBCUpgradeFlowBegins            ConditionReason = "DBCUpgradeFlowBegins"
	UpgradeComplete                 ConditionReason = "UpgradeComplete"
	UpgradePreflightCheckInProgress ConditionReason = "PreflightCheckInProgress"
	UpgradePreflightCheckFailed     ConditionReason = "PreflightCheckFailed"
	UpgradeFailed                   ConditionReason = "UpgradeFailed"
	RollbackComplete                ConditionReason = "RollbackComplete"
	RollbackFailed                  ConditionReason = "RollbackFailed"

	K8ServiceStarting ConditionReason = "K8ServiceStarting"
	K8ServiceStarted  ConditionReason = "K8ServiceStarted"
	K8ServiceStopping ConditionReason = "K8ServiceStopping"
	K8ServiceStopped  ConditionReason = "K8ServiceStopped"

	DNSPending   ConditionReason = "DNSPending"
	DNSConfirmed ConditionReason = "DNSConfirmed"
)

// Condition reason for ConditionType=LoadBalancerReady on database cluster status condition
const (
	LoadBalancerHealthy    ConditionReason = "LoadBalancerHealthy"
	LoadBalancerDegraded   ConditionReason = "LoadBalancerDegraded"
	LoadBalancerUnhealthy  ConditionReason = "LoadBalancerUnhealthy"
	LoadBalancerInProgress ConditionReason = "LoadBalancerInProgress"
)

// Condition reason for ConditionType=VIPManagerReady on database cluster status condition
const (
	VIPManagerHealthy    ConditionReason = "VIPManagerHealthy"
	VIPManagerDegraded   ConditionReason = "VIPManagerDegraded"
	VIPManagerUnhealthy  ConditionReason = "VIPManagerUnhealthy"
	VIPManagerInProgress ConditionReason = "VIPManagerInProgress"
	VIPManagerSplitBrain ConditionReason = "VIPManagerSplitBrain"
	VIPManagerError      ConditionReason = "VIPManagerError"
)

// Condition reason for ConditionType=EndpointsReady on database cluster status condition
const (
	EndpointsAssigned            ConditionReason = "EndpointsAssigned"
	EndpointsAssigningInProgress ConditionReason = "EndpointsAssigningInProgress"
	EndpointsError               ConditionReason = "EndpointsError"
)
