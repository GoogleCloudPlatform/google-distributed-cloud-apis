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

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// RCConditionHealthy is used to report whether the replication is currently
	// working and healthy.
	RCConditionHealthy ConditionType   = "Healthy"
	RCReasonHealthy    ConditionReason = "Healthy"
	RCReasonUnstable   ConditionReason = "Unstable"
	RCReasonUnhealthy  ConditionReason = "Unhealthy"

	// RCConditionReady indicates whether the database has been fully configured
	// for replication according to the parameters provided in the
	// ReplicationConfig spec. Note, this doesn't necessarily mean replication
	// is currently active.
	//
	// The Ready condition will be a summary of these more fine-grained
	// conditions depending on the type and role of the ReplicationConfig.
	RCConditionReady ConditionType = "Ready"
	// RCReasonReady is used with a True status when replication is configured.
	RCReasonReady ConditionReason = "Ready"
	// RCReasonPending is used with a False status when replication cannot be
	// configured due to a dependency not being satisfied, e.g., parent database
	// is not ready.
	RCReasonPending ConditionReason = "Pending"
	// RCReasonInProgress is used with a False status when configuration is in
	// progress.
	RCReasonInProgress ConditionReason = "InProgress"
	// RCReasonDeleting is used with a False status when the ReplicationConfig
	// is being deleted.
	RCReasonDeleting ConditionReason = "Deleting"
	// RCReasonError is a Condition Reason that may be used in any of the
	// ReplicationConfig Conditions to indicate an error has occurred preventing
	// the controller from proceeding. This may be accompanied with a
	// CriticalIncident on the ReplicationConfig status providing more
	// information on the error.
	RCReasonError ConditionReason = "Error"

	// RCConditionConfigurationReady indicates whether the database
	// configurations for replication to work have been written and
	// applied to the database.
	//
	// For Upstream ReplicationConfigs this means that the
	// 95upstreamreplication.conf file has been created with the necessary
	// configurations.
	//
	// For Downstream ReplicationConfigs this means:
	//  - The 96downstreamreplication.conf file has been created with necessary
	//    configurations including the `primary_conninfo` config.
	//  - The pgpass file has been created with the replication password.
	//  - The standby.signal file has been created.
	RCConditionConfigurationReady ConditionType = "ConfigurationReady"

	// RCConditionUserReady is used on Upstream ReplicationConfigs to indicate
	// whether the replication user has been created with the specified password
	// and is capable of accepting replication connections.
	RCConditionUserReady ConditionType = "UserReady"

	// RCConditionSetupComplete is used on Physical Downstream
	// ReplicationConfigs to indicate whether replication setup has succeeded.
	//
	// If setup has been successfully completed the condition's status will
	// be set to True and reason will be set to "Ready". Otherwise,
	// the status will be False and the reason will be:
	//   - InProgress: The controller is still attempting the setup strategies.
	//   - Failed: All of the initialization strategies specified on the spec have
	//     failed and the controller is unable to continue with the setup.
	RCConditionSetupComplete ConditionType = "SetupComplete"

	RCConditionSlotReady ConditionType = "SlotReady"

	// RConditionConfigurationValidated indicates the replication config is validated
	RCConditionConfigurationValidated ConditionType = "ConfigurationValidated"

	RCConditionConfigurationValid ConditionReason = "ConfigurationValid"

	SynchronousSettingTrue        = "true"
	SynchronousSettingFalse       = "false"
	SynchronousSettingAutoManaged = "auto_managed"
)

// +kubebuilder:object:generate=true
type ReplicationConfigSpec struct {
	// Parent is a reference to the database this ReplicationConfig belongs to.
	// +kubebuilder:validation:Required
	Parent ReplicationDatabase `json:"parent"`

	// ReplicationType determines the type of replication which will be used
	// (i.e., Physical, Logical).
	ReplicationType ReplicationType `json:"type"`

	// ReplicationRole determines the role of the ReplicationConfig's parent in
	// the replication. An Upstream role means the parent is the source of
	// replication and a Downstream role means the parent is the destination of
	// the replication.
	ReplicationRole ReplicationRole `json:"role"`

	PhysicalUpstream *ReplicationPhysicalUpstreamSpec `json:"physicalUpstream,omitempty"`
	LogicalUpstream  *ReplicationLogicalUpstreamSpec  `json:"logicalUpstream,omitempty"`

	PhysicalDownstream *ReplicationPhysicalDownstreamSpec `json:"physicalDownstream,omitempty"`
}

// +kubebuilder:object:generate=true
type ReplicationDatabase struct {
	// DBNode is a reference to the DBNode the ReplicationConfig belongs to.
	// It should be non-nil if the ReplicationConfig belongs to a DBS DBNode.
	// The DBNode should be in the same namespace as the ReplicationConfig.
	// +optional
	DBNode *corev1.LocalObjectReference `json:"dbnode,omitempty"`
}

// +kubebuilder:object:generate=true
type ReplicationConfigStatus struct {
	// EntityStatus represents the status of an Entity.
	EntityStatus `json:",inline"`

	PhysicalUpstream   *ReplicationPhysicalUpstreamStatus   `json:"physicalUpstream,omitempty"`
	PhysicalDownstream *ReplicationPhysicalDownstreamStatus `json:"physicalDownstream,omitempty"`

	LogicalUpstream *ReplicationLogicalUpstreamStatus `json:"logicalUpstream,omitempty"`
}

// +kubebuilder:object:generate=true
type ReplicationUpstreamStatus struct {
	// PasswordResourceVersion is the Password Secret's resourceVersion when the
	// password was last updated on the database.
	// +optional
	PasswordResourceVersion string `json:"passwordResourceVersion,omitempty"`

	// ClientAddr is the address of the downstream client connected to this replication slot.
	// +optional
	ClientAddr string `json:"clientAddr,omitempty"`
	// ClientHostname is the hostname of the downstream client connected to this replication slot.
	// +optional
	ClientHostname string `json:"clientHostname,omitempty"`
	// ClientPort is the source port of the downstream client connected to this replication slot.
	// +optional
	ClientPort uint16 `json:"clientPort,omitempty"`
	// StartedAt is the time at which the downstream client connected to the server.
	// +optional
	StartedAt metav1.Time `json:"startedAt,omitempty"`
	// State is the current state of replication. It can take one of the following values:
	// See document for the `state` column of the `pg_stat_replication` table for more info:
	// https://www.postgresql.org/docs/current/monitoring-stats.html#MONITORING-PG-STAT-REPLICATION-VIEW
	// +optional
	State string `json:"state,omitempty"`

	// SynchronousEnabled is whether this replication application name is marked as synchronous
	// from the database property synchronous_standby_names as read from the database itself.
	// If the database is not connectable, this value will remain unchanged.
	SynchronousEnabled bool `json:"synchronousEnabled,omitempty"`

	// SynchronousStatus returns the value from the column sync_state from the
	// database table pg_stat_replication. The currently supported values are
	// empty string, async, potential, sync, quorum. If there is no active connection
	// this value will be empty.
	// If the database is not connectable, this value will remain unchanged.
	SynchronousStatus string `json:"synchronousStatus,omitempty"`
}

// +kubebuilder:object:generate=true
type ReplicationUpstreamUserSpec struct {
	// User is the name of a database user that will be created on the Instance
	// for this ReplicationConfig. Multiple Upstream ReplicationConfigs can
	// share the same user.
	//
	// Note, the User specified here will be managed by the ReplicationConfig
	// controller and removed when there are no longer any ReplicationConfigs
	// with this username. Do not specify a user here if you wish it to have
	// a lifecycle outside of the ReplicationConfig's lifecycle.
	//
	// +optional
	User string `json:"username,omitempty"`

	// Password is a reference to a Secret holding the User's password.
	// Any update made to the Secret will be captured and reflected on the
	// database user.
	//
	// +optional
	Password *corev1.SecretReference `json:"password,omitempty"`

	// SkipUserCreateAndUpdate is true when replication setup should skip user
	// creation and updates.
	// The primary use case for this is when this database is upstream and
	// downstream - i.e. a cascading database. This skips the step to create/set
	// the replication user since it should be propagated from the upstream.
	// +kubebuilder:default=false
	SkipUserCreateAndUpdate bool `json:"skipUserCreateAndUpdate,omitempty"`
}

// +kubebuilder:object:generate=true
type ReplicationPhysicalUpstreamSpec struct {
	ReplicationUpstreamUserSpec `json:",inline"`

	// SlotName is the replication slot that will be configured on the database.
	// This must be unique among all PhysicalUpstream and LogicalUpstream specs
	// on the same instance.
	// +kubebuilder:validation:Required
	SlotName string `json:"slotName"`

	// Synchronous sets whether this replication should be synchronous or not
	// `true` means this replication should be set as synchronous.
	// `false` means that this replication should be asynchronous.
	// `auto_managed` means that the synchronous state of this replication will be
	//    managed by the DBS system. It will generally be set to synchronous after
	//    initial set up is complete, but might be temporarily disabled during
	//    certain operations.
	// `` is used for backward compatibility but should be treated as false
	// +kubebuilder:validation:Enum="true";"false";"auto_managed";""
	// +kubebuilder:default="false"
	Synchronous string `json:"synchronous"`

	// LogReplicationSlot if true will cause the replication slot used by the
	// ReplicationConfig to be logged in the WAL files. If the corresponding
	// downstream has ReplayReplicationSlots enabled then they will be able to
	// replicate this replication slot on the downstream database.
	LogReplicationSlot bool `json:"logReplicationSlot,omitempty"`

	// PreserveSlot if true will keep the replication slot on the database even
	// after the upstream ReplicationConfig has been deleted.
	PreserveSlot bool `json:"preserveSlot,omitempty"`
}

// +kubebuilder:object:generate=true
type ReplicationPhysicalUpstreamStatus struct {
	ReplicationUpstreamStatus `json:",inline"`
}

// +kubebuilder:object:generate=true
type ReplicationLogicalUpstreamSpec struct {
	ReplicationUpstreamUserSpec `json:",inline"`

	// SlotName is the replication slot that will be configured on the database.
	// This must be unique among all PhysicalUpstream and LogicalUpstream specs
	// on the same instance.
	// +kubebuilder:validation:Required
	SlotName string `json:"slotName"`

	// Synchronous sets whether this replication should be synchronous or not
	// `true` means this replication should be set as synchronous.
	// `false` means that this replication should be asynchronous.
	// +kubebuilder:validation:Enum="true";"false";""
	// +kubebuilder:default="false"
	// +optional
	Synchronous string `json:"synchronous,omitempty"`

	// PluginName is the logical decoding plugin that should be associated with the slot.
	// See https://www.postgresql.org/docs/current/logicaldecoding-explanation.html#LOGICALDECODING-EXPLANATION-OUTPUT-PLUGINS
	// for more information.
	// +kubebuilder:validation:Required
	PluginName string `json:"pluginName"`

	// ApplicationName is required when Synchronous is set to true. The ApplicationName will
	// be used as the standby name when configuring synchronous replication.
	// +optional
	ApplicationName string `json:"applicationName"`

	// DatabaseName is the database with which the logical replication slot will be associated.
	// +kubebuilder:validation:Required
	DatabaseName string `json:"databaseName"`
}

// +kubebuilder:object:generate=true
type ReplicationLogicalUpstreamStatus struct {
	ReplicationUpstreamStatus `json:",inline"`
}

// +kubebuilder:object:generate=true
type ReplicationPhysicalDownstreamSpec struct {
	// SlotName is the replication slot that the database will use on the upstream server.
	// +kubebuilder:validation:Required
	SlotName string `json:"slotName"`

	// User is the database user which will be used to establish the replication
	// connection.
	// +kubebuilder:validation:Required
	User string `json:"username"`

	// Password is a reference to a Secret holding the User's password.
	// +kubebuilder:validation:Required
	Password *corev1.SecretReference `json:"password"`

	// Host is the hostname or address of the upstream database server to connect to.
	// +kubebuilder:validation:Required
	Host string `json:"host"`

	// Port is the port number of the upstream database server to connect to.
	// kubebuilder:default:=5432
	Port uint16 `json:"port,omitempty"`

	// SetupStrategies determine how the initial setup will be done so that
	// the downstream database can start streaming from the upstream.
	//
	// Multiple strategies can be provided to provide fallbacks in case a
	// strategy fails. They will be attempted in the same order they are
	// provided in this list. If a strategy succeeds then the rest of the
	// strategies in the list will be ignored. If all strategies fail then the
	// ReplicationConfig will be in a permanently failed state and the user must
	// delete and recreate the ReplicationConfig to retry.
	//
	// Note, not all failures that occur in a setup strategy will make it
	// fallback to the next strategy. Depending on what type of error occurs, we
	// might retry the same strategy again or fallback to the next strategy.
	//
	// Typically, errors such as connection errors will result in the strategy
	// being retried since the actual strategy wouldn't have gotten the chance
	// to actually be attempted. On the other hand, if a strategy determines
	// that it would not be able to succeed no matter how many times it is
	// retried then it would fallback. Documentation on each strategy should
	// specify in what circumstances it would fallback to the next strategy.
	SetupStrategies []PhysicalReplicationSetupStrategy `json:"setupStrategies,omitempty"`

	// RerunSetupAt will rerun the replication setup if the time on this is more
	// recent than the time of the lastSetup. This will be based on the time of
	// the last setup being completed successfully.
	//
	// If this controller detects that setup needs to be rerun, then it will
	// first promote this instance to bring it back into a read/write state.
	// Then it will execute the setupStrategies.
	//
	// The main purpose of this is to rerun some setups, such as pg_rewind if
	// the downstream has gotten briefly out of sync with the upstream, and we
	// can avoid the re-creation of the standby from scratch.
	RerunSetupAt *metav1.Time `json:"rerunSetupAt,omitempty"`

	// InvalidateSyncStatusAt if specified causes the SynchronousEnabled and
	// SynchronousState fields in the status to be invalidated until their
	// values are propagated from the upstream with a timestamp newer than the
	// time specified in this field.
	//
	// Once this field is updated, in the next reconcile the controller will
	// reset the status values if their associated timestamp is older than the
	// time specified in this field. If the values remain after the reconcile it
	// means they have a timestamp more recent than the specified time. You may
	// compare the Generation and ObservedGeneration of the ReplicationConfig to
	// know whether the controller has performed a reconcile after updating this
	// field.
	InvalidateSyncStatusAt *metav1.Time `json:"invalidateSyncStatusAt,omitempty"`

	// ReplayReplicationSlots if set to true will cause replication slots to be
	// replayed on the database if they exist in the WAL files received from the
	// upstream, allowing replication slots to be replicated from the upstream.
	//
	// This feature is only supported on AlloyDBOmni databases which have the
	// enable_log_replication_slots parameter set to on.
	//
	// After being promoted, any replication slots that were previously replicated
	// from the upstream will remain in the database, except for the replication
	// slot that was being used by the downstream itself.
	// Any replication slots that do remain after promotion are considered
	// "unmanaged" by RaaS, until a corresponding Upstream ReplicationConfig is
	// created with the same slot name.
	ReplayReplicationSlots bool `json:"replayReplicationSlots,omitempty"`

	// ReplicationSlotsAllowList is a list of replication slot names which will
	// be allowed stay in the downstream. Any replication slots which exist on
	// the downstream but do not exist in this list will be dropped.
	//
	// Note, this won't prevent the replication slots from getting replayed on
	// the database if ReplayReplicationSlots is enabled and there are
	// replication slots in the WAL files. It will however periodically delete
	// those slots if they do not exist in the allow-list.
	//
	// If set to nil or empty then no slot names will be allowed.
	ReplicationSlotsAllowList []string `json:"replicationSlotsAllowList,omitempty"`
}

// +kubebuilder:object:generate=true
type PhysicalReplicationSetupStrategy struct {
	// PGBaseBackup is a replication setup strategy that uses pg_basebackup
	// to retrieve a backup of the upstream database.
	//
	// This strategy will never fallback to next strategy and will always be
	// retried on errors. It should typically be used as the last-resort
	// strategy that is expected to succeed as long as the upstream database is
	// available, but may be slow to complete.
	// +optional
	PGBaseBackup *PGBaseBackupStrategy `json:"pgBaseBackup,omitempty"`

	// PGRewind is a downstream replication setup strategy that uses
	// pg_rewind to put the downstream in-sync with the upstream. It is useful for
	// cases where the two databases where previously replicating from each other
	// but have since diverged.
	//
	// This strategy will check connectivity with the upstream before running
	// pg_rewind. If it fails due to the upstream being unreachable it will be
	// retried, however if the pg_rewind command is run and returns
	// unsuccessfully then it will fallback to the next strategy.
	// +optional
	PGRewind *PGRewindStrategy `json:"pgRewind,omitempty"`
}

func (s *PhysicalReplicationSetupStrategy) StrategyName() string {
	switch {
	case s.PGBaseBackup != nil:
		return "PGBaseBackup"
	case s.PGRewind != nil:
		return "PGRewind"
	default:
		return "Unknown"
	}
}

// +kubebuilder:object:generate=true
type PGBaseBackupStrategy struct {
	// Checkpoint controls how the PostgreSQL server performs a checkpoint over
	// before initiating the base backup.
	//
	// Accepted values are:
	//   - fast: This option tells PostgreSQL to perform a "fast" checkpoint.
	//     It is the quickest way to create a checkpoint, but it may cause some
	//     additional load on the server during the backup process.
	//   - spread: This option instructs PostgreSQL to spread the checkpoint
	//     over a longer period. It minimizes the impact on the server's
	//     performance during the backup but might take longer to complete the
	//     checkpoint.
	//
	// +kubebuilder:validation:Enum=fast;spread
	// +kubebuilder:default:=fast
	// +optional
	Checkpoint string `json:"checkpoint,omitempty"`

	// MaxRate sets the maximum transfer rate at which data is collected from
	// the source server.
	//
	// This can be useful to limit the impact of pg_basebackup on the server.
	// Values are in kilobytes per second. Use a suffix of M to indicate
	// megabytes per second. A suffix of k is also accepted, and has no effect.
	// Valid values are between 32 kilobytes per second and 1024 megabytes per
	// second.
	//
	// +kubebuilder:validation:Pattern=^[0-9]+[kKmM]?$
	// +optional
	MaxRate string `json:"maxRate,omitempty"`

	// WalMethod determines if and how WAL records should be collected during
	// backup. This will include all write-ahead logs generated during the
	// backup. Unless the method none is specified, it is possible to start a
	// postmaster in the target directory without the need to consult the WAL
	// archive, thus making the output a completely standalone backup.
	//
	// Accepted values are:
	//   - none: Don't include write-ahead logs in the backup.
	//   - fetch: The write-ahead log files are collected at the end of the backup.
	//   - stream: Stream write-ahead log data while the backup is being taken.
	//
	// +kubebuilder:validation:Enum=none;fetch;stream
	// +kubebuilder:default:=stream
	// +optional
	WalMethod string `json:"walMethod,omitempty"`
}

// +kubebuilder:object:generate=true
type PGRewindStrategy struct{}

// +kubebuilder:object:generate=true
type ReplicationPhysicalDownstreamStatus struct {
	// State is the state of replication as seen in the pg_stat_wal_receiver table
	// of the downstream database server.
	State string `json:"state,omitempty"`

	// SetupStrategies contains information on the execution of each attempted
	// setup strategy. They appear in this list in the same order as the
	// strategies were defined in the spec.
	SetupStrategies []PhysicalReplicationSetupStrategyStatus `json:"setupStrategies,omitempty"`

	// SynchronousEnabled indicates whether the application_name used by this
	// replication has been included in the synchronous_standby_names config on
	// the upstream database.
	// When using this field, keep in mind that there may be a delay between
	// when the change is made on the upstream and when the new value gets
	// reflected here. Also, this field will contain the last value pushed down
	// from the upstream. If upstream is unreachable it will retain its last
	// value until it can reach the upstream again.
	SynchronousEnabled bool `json:"synchronousEnabled,omitempty"`

	// SynchronousStatus contains the value from the sync_state column from the
	// database table pg_stat_replication on the upstream side of the
	// connection.
	// When using this field, keep in mind that there may be a delay between
	// when the change is made on the upstream and when the new value gets
	// reflected here. Also, this field will contain the last value pushed down
	// from the upstream. If the upstream is unreachable it will retain its last
	// value until it can reach the upstream again.
	SynchronousStatus string `json:"synchronousStatus,omitempty"`

	// PasswordResourceVersion is the Password Secret's resourceVersion when the
	// password was last updated on the database.
	// +optional
	PasswordResourceVersion string `json:"passwordResourceVersion,omitempty"`
}

type RCSetupStrategyState string

const (
	RCSetupInProgress RCSetupStrategyState = "InProgress"
	RCSetupSuccess    RCSetupStrategyState = "Success"
	RCSetupError      RCSetupStrategyState = "Error"
	RCSetupFallback   RCSetupStrategyState = "Fallback"
	RCSetupUnknown    RCSetupStrategyState = "Unknown"
)

// +kubebuilder:object:generate=true
type PhysicalReplicationSetupStrategyStatus struct {
	// Strategy is the name of the strategy type this status is for.
	Strategy string `json:"strategy"`
	// State is the current state of this setup strategy. It can take the
	// following values:
	//
	//  - InProgress: The strategy is currently executing.
	//  - Success: The strategy has successfully completed and no more setup strategies will be attempted.
	//  - Error: The strategy has failed but will be retried. The Retries field will show how many times this strategy has been retried.
	//  - Fallback: The strategy has failed and will not be reattempted. Instead we will fallback to the next available strategy if it exists.
	//
	// +kubebuilder:validation:Enum=Unknown;InProgress;Success;Error;Fallback
	State RCSetupStrategyState `json:"state"`
	// Message is a description of why the setup attempt is in the state it is.
	// +optional
	Message string `json:"message"`
	// StartedAt is the time at which the most recent attempt of this strategy
	// was started.
	// +optional
	StartedAt metav1.Time `json:"startedAt,omitempty"`
	// EndedAt is the time at which the most recent attempt of this strategy
	// ended..
	// +optional
	EndedAt metav1.Time `json:"endedAt,omitempty"`
	// Retries is the number of times this strategy has been retried.
	// +optional
	Retries int64 `json:"retries"`
}

type ReplicationConfig interface {
	Entity
	ReplicationConfigSpec() *ReplicationConfigSpec
	ReplicationConfigStatus() *ReplicationConfigStatus

	PhysicalUpstreamSpec() *ReplicationPhysicalUpstreamSpec
	PhysicalUpstreamStatus() *ReplicationPhysicalUpstreamStatus
	PhysicalDownstreamSpec() *ReplicationPhysicalDownstreamSpec
	PhysicalDownstreamStatus() *ReplicationPhysicalDownstreamStatus

	LogicalUpstreamSpec() *ReplicationLogicalUpstreamSpec
	LogicalUpstreamStatus() *ReplicationLogicalUpstreamStatus

	UpstreamUserSpec() *ReplicationUpstreamUserSpec
	UpstreamStatus() *ReplicationUpstreamStatus
}

type ReplicationConfigList interface {
	client.ObjectList
	ReplicationConfigs() []client.Object
}

// NewReplicationConfig creates a zero-value ReplicationConfig object from the
// same GroupVersion as the object given in the parameters.
func NewReplicationConfig(scheme *runtime.Scheme, gvObj runtime.Object) (ReplicationConfig, error) {
	return CreateZeroObjFromSameGroupVersion[ReplicationConfig](scheme, gvObj, "ReplicationConfig")
}

// NewReplicationConfigList creates a zero-value ReplicationConfigList object
// from the same GroupVersion as the object given in the parameters.
func NewReplicationConfigList(scheme *runtime.Scheme, gvObj runtime.Object) (ReplicationConfigList, error) {
	return CreateZeroObjFromSameGroupVersion[ReplicationConfigList](scheme, gvObj, "ReplicationConfigList")
}
