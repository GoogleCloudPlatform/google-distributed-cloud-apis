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

	eeexportapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/export/v1"
	eehaapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/ha/v1"
	eeimportapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/import/v1"
	cecoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

type DBClusterPhase string

const (
	DBClusterReconciling                   DBClusterPhase = "DBClusterReconciling"
	DBClusterCertificateProvisioningFailed DBClusterPhase = "DBClusterCertificateProvisioningFailed"
	DBClusterWaitingForCertificate         DBClusterPhase = "DBClusterWaitingForCertificate"
	DBClusterReady                         DBClusterPhase = "DBClusterReady"
	DBClusterPermanentlyDeleting           DBClusterPhase = "DBClusterPermanentlyDeleting"
	DBClusterStopped                       DBClusterPhase = "DBClusterStopped"
	DBClusterDeleted                       DBClusterPhase = "DBClusterDeleted"
	DBClusterDeletingBackupFailed          DBClusterPhase = "DBClusterDeletingBackupFailed"
	DBClusterDeletingBackupSuccess         DBClusterPhase = "DBClusterDeletingBackupSuccess"
	DBClusterDeletingInstanceFailed        DBClusterPhase = "DBClusterDeletingInstanceFailed"
	DBClusterDeleting                      DBClusterPhase = "DBClusterDeleting"
	DBClusterRestoring                     DBClusterPhase = "DBClusterRestoring"
	DBClusterRestoringFailed               DBClusterPhase = "DBClusterRestoringFailed"
	DBClusterPendingDisasterRecovery       DBClusterPhase = "DBClusterPendingDisasterRecovery"
	DBClusterDisasterRecovering            DBClusterPhase = "DBClusterDisasterRecovering"
	DBClusterDisasterRecoveryFailed        DBClusterPhase = "DBClusterDisasterRecoveryFailed"
	DBClusterCloning                       DBClusterPhase = "DBClusterCloning"
	DBClusterCloningFailed                 DBClusterPhase = "DBClusterCloningFailed"
	DBClusterWaitingForAuditLogTarget      DBClusterPhase = "DBClusterWaitingForAuditLogTarget"
	DBClusterUpgradeFailed                 DBClusterPhase = "DBClusterUpgradeFailed"
	DBClusterFailoverInProgress            DBClusterPhase = "DBClusterFailoverInProgress"
	DBClusterSwitchoverInProgress          DBClusterPhase = "DBClusterSwitchoverInProgress"
	DBClusterInMaintenance                 DBClusterPhase = "DBClusterInMaintenance"
	DBClusterSwitchoverRollbackFailed      DBClusterPhase = "DBClusterSwitchoverRollbackFailed"
)

const (
	DbPort string = "Port"

	// Different modes for DBCluster.
	// DBCluster in this mode goes through normal provision steps.
	NormalMode string = ""
	// DBCluster in this mode indicates a disaster recovery will be performed thus skip normal provision steps.
	DisasterRecoveryMode string = "disasterRecovery"

	// MaintenanceMode turns off DBCluster probes (liveness, startup).
	MaintenanceMode string = "maintenance"

	DBClusterKind string = "DBCluster"

	AnnotationForceReconcile string = "forceReconcile"
)

//+kubebuilder:object:generate=true

type PrimaryStatus struct {
	// Phase is a summary of current state of the Instance.
	// +optional
	Phase cecoreapi.InstancePhase `json:"phase,omitempty"`

	// Conditions represents the latest available observations
	// of the Instance's current state.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Endpoint is the address that can be used to establish database connections.
	// Deprecated: use the Endpoints field instead.
	Endpoint string `json:"endpoint,omitempty"`

	// URL consists of the address and port number that can be used to
	// establish a client connection to the database.
	//
	// This value is expressed in the following format: <address>:<port>.
	// Deprecated: use the Endpoints field instead.
	URL string `json:"url,omitempty"`

	// Endpoints contains all the endpoint through which the users can access this instance.
	// +listType=map
	// +listMapKey=name
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +optional
	Endpoints []cecoreapi.Endpoint `json:"endpoints,omitempty" patchStrategy:"merge" patchMergeKey:"name"`

	// LatestImport represents the latest import for the database instance
	// +optional
	LatestImport *PrimaryImportStatus `json:"latestImport,omitempty"`

	// LatestExport represents the latest export for the database instance
	// +optional
	LatestExport *PrimaryExportStatus `json:"latestExport,omitempty"`

	// CurrentParameters indicates the current values of the parameters.
	//
	// CurrentParameters allows to verify that the `spec.primarySpec.parameters` field has been
	// applied to the database. Only the parameters names in `spec.primarySpec.parameters` will be
	// included in this field.
	CurrentParameters map[string]string `json:"currentParameters,omitempty"`

	// LastFailedParameterUpdate is used to avoid getting into the failed
	// parameter update loop.
	//
	// nullon(samwise-fleet)
	LastFailedParameterUpdate map[string]string `json:"lastFailedParameterUpdate,omitempty"`

	// ActiveImages stores the stable images used by the active containers.
	// nullon(samwise-fleet)
	ActiveImages map[string]string `json:"ActiveImages,omitempty"`

	// LastFailedImages stores the images which failed the last patching workflow.
	// nullon(samwise-fleet)
	LastFailedImages map[string]string `json:"LastFailedImages,omitempty"`

	// AllocatedResources represents the current configuration of memory/CPU/disks
	AllocatedResources *cecoreapi.Resource `json:"allocatedResources,omitempty"`

	// CurrentDatabaseVersion is the current database version that the primary instance is running.
	//
	// This value should match the value of `spec.databaseVersion` after the primary instance is
	// provisioned or the upgrade or downgrade has concluded successfully.
	//
	// nullon(dbs-fleet)
	CurrentDatabaseVersion string `json:"currentDatabaseVersion,omitempty"`

	// CurrentControlPlaneAgentsVersion is the control plane agents version that the primary
	// instance is running.
	//
	// This value should match the value of `spec.controlPlaneAgentsVersion` after the primary
	// instance is provisioned or the upgrade or downgrade has concluded successfully.
	//
	// nullon(dbs-fleet)
	CurrentControlPlaneAgentsVersion string `json:"currentControlPlaneAgentsVersion,omitempty"`

	// CurrentDatabaseImage is the customized database image that the primary instance is using.
	//
	// This value should match the value of `spec.databaseImage` after the primary instance is
	// provisioned or the upgrade or downgrade has concluded successfully.
	//
	// nullon(dbs-fleet)
	// +optional
	CurrentDatabaseImage string `json:"currentDatabaseImage,omitempty"`

	// WalArchiveSetting represents the current wal archive settings.
	// nullon(dbs-fleet)
	// +optional
	WalArchiveSetting *cecoreapi.WalArchiveStatus `json:"walArchiveSetting,omitempty"`
}

//+kubebuilder:object:generate=true

type DBClusterStatus struct {
	cecoreapi.EntityStatus `json:",inline"`
	Phase                  DBClusterPhase `json:"phase,omitempty"`
	// Primary contains the status of the primary Instance.
	Primary PrimaryStatus `json:"primary,omitempty"`

	// +optional
	RestoredFrom *RestoredFrom `json:"restoredFrom,omitempty"`

	// UpgradeScheduledAt is a timestamp that indicates when the next upgrade is
	// scheduled to start. If it is nil, it means there is no upcoming upgrade scheduled.
	// nullon(samwise-fleet)
	// +optional
	UpgradeScheduledAt *metav1.Time `json:"upgradeScheduledAt,omitempty"`

	// ServiceAccounts contains the service accounts created by the control plane to
	// be used by different operations. By granting permissions to these service accounts,
	// the database can interact with other services within the kubernetes ecosystem.
	// For further information, including what permissions is required, refer to the
	// documentation of each operation.
	// +optional
	ServiceAccounts map[OpType]cecoreapi.ServiceAccountRef `json:"serviceAccounts,omitempty"`

	// CertificateReference refers to a secret and a key of the server CA certificate
	// that can be used to connect to the database.
	//
	// If `spec.TLS` is provided, this value should refer to the same secret after the
	// database has been configured to use the provided server certificate.
	CertificateReference *cecoreapi.CertificateRef `json:"certificateReference,omitempty"`

	// TLS is the current server certificate configuration for the database cluster.
	// +optional
	TLS *TLSStatus `json:"tls,omitempty"`

	// LatestFailoverStatus is the status of the most recently updated failover for the database cluster
	// This status is a copy of the status of the current or most recently updated failover operation for the database cluster.
	// This can be used to conveniently monitor the status of a currently running failover operation.
	// +optional
	LatestFailoverStatus *FailoverStatus `json:"latestFailoverStatus,omitempty"`

	// MigrationStatus represents the status of migration for the database cluster.
	// nullon(samwise-fleet)
	// +optional
	MigrationStatus *MigrationStatus `json:"migrationStatus,omitempty"`

	// AvailabilityZones represents the zone status for a multi-zone HA database cluster.
	// nullon(samwise-fleet)
	// +optional
	AvailabilityZones *AvailabilityZones `json:"availabilityZones,omitempty"`

	// PreviousAvailabilityZones represents the zone status for a multi-zone HA database cluster before a HA operation.
	// nullon(samwise-fleet)
	// +optional
	PreviousAvailabilityZones *AvailabilityZones `json:"previousAvailabilityZones,omitempty"`
}

//+kubebuilder:object:generate=true

// RestoredFrom is the status showing the most recent restore source for current DBCluster.
type RestoredFrom struct {
	// Source DBCluster this DBCluster restores from.
	SourceDBCluster string `json:"sourceDBCluster,omitempty"`

	// Time point of the source DBCluster this DBCluster restores from.
	RestoredTime *metav1.Time `json:"restoredTime,omitempty"`
}

type InstanceSpec interface {
	CommonSpec() *cecoreapi.InstanceSpec
}

type AvailabilityType string

const (
	DualZone AvailabilityType = "dualzone"
)

//+kubebuilder:object:generate=true

// Availability contains customization options for high availability (HA) features for the DBCluster
type Availability struct {
	// EnableHighAvailability sets this DBCluster to be a High Availability cluster.
	// +kubebuilder:default=false
	// nullon(samwise-fleet)
	// +optional
	EnableHighAvailability bool `json:"enableHighAvailability,omitempty"`

	// EnableAutoFailover means this DBCluster will trigger a failover if it detects the primary instance is unhealthy and standby instance is healthy.
	// If set to `true`, then automatic failover is enabled. If set to `false`, then autofailover will not be triggered even if the system detects that the primary instance is unhealthy. The default value is `true`.
	// When it is enabled, if the system detects that the primary instance is unhealthy for the given threshold, it will trigger a failover.
	// This feature is only applicable if this is a HA DBCluster and if the standby is healthy.
	// +kubebuilder:default=true
	// +optional
	EnableAutoFailover bool `json:"enableAutoFailover"`

	// HealthcheckPeriodSeconds is the number of seconds the healthcheck prober will wait before checking the health of the primary and standby instances again and updating the status accordingly.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=86400
	// +kubebuilder:default=30
	// +optional
	HealthcheckPeriodSeconds int `json:"healthcheckPeriodSeconds"`

	// AutoFailoverTriggerThreshold is the number of consecutive healthcheck failures on the primary instance that will trigger an automatic failover.
	// If set to 0, then it will use the system default value. Use the EnableAutoFailover flag to completely disable automatic failover.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=3
	// +optional
	AutoFailoverTriggerThreshold int `json:"autoFailoverTriggerThreshold"`

	// EnableAutoHeal means this DBCluster will trigger an autoheal if it detects the standby instance is unhealthy.
	// If set to `true`, then autoheal is enabled. If set to `false`, then autoheal will not be triggered even if the system detects that the standby instance is unhealthy. The default value is `true`.
	// When it is enabled, if the system detects that the standby instance is unhealthy for the given threshold, it will trigger an autoheal.
	// This feature is only applicable if this is a HA DBCluster.
	// +kubebuilder:default=true
	// +optional
	EnableAutoHeal bool `json:"enableAutoHeal"`

	// AutoHealTriggerThreshold is the number of consecutive healthcheck failures on the standby instance that will trigger automatic healing.
	// Use the EnableAutoHeal flag to completely disable automatic healing.
	// +kubebuilder:validation:Minimum=2
	// +kubebuilder:default=5
	// +optional
	AutoHealTriggerThreshold int `json:"autoHealTriggerThreshold"`

	// NumberOfStandbys is the number of standbys that should be created for this DBCluster.
	// If set to any value greater than `0`, then HA is enabled on the cluster and the system will create the indicated number of standby instances. The maximum allowed standby instances is 5.
	// To check the current status of HA on this DBCluster, look at the HAReady condition under the DBCluster status. If HAReady is `true`, then setup has been complete and ready.
	// If set to `0`, then HA is disabled on the cluster, and deletes any existing standby instances.
	// Any number between `0` and `5` inclusive is supported. The default value is `0`.
	//
	// Additional Documentation: https://cloud.google.com/alloydb/docs/omni/kubernetes-ha
	// +kubebuilder:validation:Maximum=5
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=0
	// nullon(dbs-fleet)
	// +optional
	NumberOfStandbys int `json:"numberOfStandbys,omitempty"`

	// EnableStandbyAsReadReplica determine whether the standbys can accept user
	// queries or not. If set to true, a new endpoint will be created to enable read-only access to the standby(s).
	//
	// nullon(dbs-fleet)
	EnableStandbyAsReadReplica bool `json:"enableStandbyAsReadReplica"`

	// ReplayReplicationSlotsOnStandbys if set to true will allow replication
	// slots that are written in WAL files to be replayed on the HA standbys.
	//
	// When using this option it is recommended to also enable the
	// LogReplicationSlot field on the upstream Replication resources so that
	// the corresponding replication slots are logged in the WAL files.
	//
	// This will ensure that in the event of an HA failover or switchover on the
	// primary DBCluster, the new HA primary instance retains WAL files which
	// have not yet been consumed by these replication slots.
	//
	// Note, modifying this field will cause all HA standbys to restart.
	//
	// nullon(dbs-fleet)
	ReplayReplicationSlotsOnStandbys *bool `json:"replayReplicationSlotsOnStandbys,omitempty"`

	// Type represents the availability type of this dbcluster.
	// nullon(samwise-fleet)
	// +kubebuilder:validation:Enum="";"dualzone"
	// +optional
	Type AvailabilityType `json:"type,omitempty"`

	// AvailabilityZones specifies the zones for a multi-zone HA database cluster.
	// If set, it will be used for initial deployment. The actual zones can vary
	// after initial deployment, e.g., post failover or disaster recovery operations.
	// nullon(samwise-fleet)
	// +optional
	AvailabilityZones *AvailabilityZones `json:"availabilityZones,omitempty"`
}

//+kubebuilder:object:generate=true

// TLSSpec defines certificate secrets for encrypted communication for instances that are
// part of the database cluster.
type TLSSpec struct {
	// CertSecret references the certificate secret within the same namespace.
	// The secret must contain entries ca.crt (CA certificate), tls.key (server private key),
	// and tls.crt (server leaf certificate). This secret is used to set the TLS config
	// for the database instances that a part of the database cluster.
	CertSecret *corev1.LocalObjectReference `json:"certSecret,omitempty"`
}

//+kubebuilder:object:generate=true

// TLSStatus contains the active resources used for encrypted communication
// with the database instance.
type TLSStatus struct {
	// DataPlaneCertIssuer refers to the issuer that was used to provision the
	// data plane certificates, like for the database server.
	DataPlaneCertIssuer *cecoreapi.IssuerReference `json:"dataPlaneCertIssuer,omitempty"`

	// ControlPlaneAgentsCertIssuer refers to the issuer that was used to sign
	// certificates for control plane agent components.
	// nullon(dbs-fleet,dbs-local)
	ControlPlaneAgentsCertIssuer *cecoreapi.IssuerReference `json:"controlPlaneAgentsCertIssuer,omitempty"`

	// DataPlaneCertRequest refers to the certificate request fields that were
	// used to provision the data plane certificates, like for the database
	// server.
	// nullon(dbs-fleet,dbs-local)
	DataPlaneCertRequest *cecoreapi.DataPlaneCertificateRequest `json:"dataPlaneCertRequest,omitempty"`

	// ControlPlaneAgentsCertRequest refers to the certificate request fields that
	// were used to provision certificates for control plane agent components.
	// nullon(dbs-fleet,dbs-local)
	ControlPlaneAgentsCertRequest *cecoreapi.CertificateRequest `json:"controlPlaneAgentsCertRequest,omitempty"`
}

//+kubebuilder:object:generate=true

type DBClusterSpec struct {
	// nullon(samwise-fleet)
	PrimaryCluster string `json:"primaryCluster,omitempty"`

	// IsDeleted indicates a request to delete the DBCluster. This field is
	// optional, and by default it is false.
	//
	// This fields applies to all instances of the database cluster. When set this
	// to true, the control plane will attempt to release the related resources,
	// including network endpoints.
	// See the status fields for indication of success or failures, if any.
	// +kubebuilder:default=false
	// +optional
	IsDeleted bool `json:"isDeleted,omitempty"`

	// Indicate the mode of this DBCluster.
	// +kubebuilder:validation:Enum="";"disasterRecovery";"maintenance"
	// +optional
	Mode string `json:"mode"`

	// Availability contains adjustable settings for DBCluster HA features
	// +optional
	Availability *Availability `json:"availability,omitempty"`

	// AllowExternalIncomingTraffic is used to toggle external load balancer creation.
	// +kubebuilder:default=false
	// +optional
	AllowExternalIncomingTraffic bool `json:"allowExternalIncomingTraffic,omitempty"`

	// TLS is the desired server certificate configuration for the cluster. This field is optional.
	// By default, this field is empty and a new self-signed CA and leaf certificate
	// are generated for the cluster. When this field is changed, the database cluster
	// pods will restart to load the specified certificate. The field `status.certificateReference`
	// indicates the current CA certificate secret and key.
	//
	// +optional
	TLS *TLSSpec `json:"tls,omitempty"`

	// DatabaseVersion is the desired database version for the cluster for example, "15.4.5". This
	// field is required.
	//
	// This version is applied to all instances of the database cluster. In the case of a new
	// database cluster, the instance is created using the specified version. In the case of an
	// existing database cluster, the operator attempts to upgrade or downgrade to the
	// specified `databaseVersion`. The field `status.currentDatabaseVersion` indicates the current
	// database version.
	//
	// See the list of available versions in {https://docs.cloud.google.com/alloydb/omni/kubernetes/current/docs/choose-compatible-versions}.
	//
	// nullon(dbs-fleet)
	DatabaseVersion string `json:"databaseVersion,omitempty"`

	// databaseImageOSType is an optional field to explicitly select the base operating system
	// of the AlloyDBOmni database container image (e.g., Debian or UBI9).
	// +optional
	// +kubebuilder:validation:Enum=Debian;UBI9
	// nullon(dbs-fleet)
	DatabaseImageOSType cecoreapi.ImageOSType `json:"databaseImageOSType,omitempty"`

	// ControlPlaneAgentsVersion is the desired control plane agents version for the cluster for
	// example, "0.5.2". This field is required.
	//
	// The `controlPlaneAgentsVersion` must be compatible with the chosen `databaseVersion`.
	// To know what versions are compatible, check the list of available versions in {link}.
	//
	// This version is applied to all instances of the database cluster. In the case of a new
	// database cluster, the instance is be created using the specified version. In the case of an
	// existing database cluster, the operator will aptempt to upgrade or downgrade to the
	// specified `controlPlaneAgentsVersion`. The field `status.currentControlPlaneAgentsVersion`
	// indicates the current version for control plane agents.
	//
	// TODO(b/320311538): replace link with the list of available versions.
	//
	// nullon(dbs-fleet)
	ControlPlaneAgentsVersion string `json:"controlPlaneAgentsVersion,omitempty"`

	// DatabaseImage is the URI of a customized database image within the container registry,
	// for example, "gcr.io/foo/bar/alloydbomni:15-7-2-customized". This field is optional.
	//
	// If `databaseImage` is specified, then the operator uses this container image for the
	// database instead of the  default database container image of the specified `databaseVersion`.
	// We recommend that the `databaseImage` container will be based on the default database
	// image used of the chosen `databaseVersion`.
	//
	// For more information about using a customized database image visit {link}.
	//
	// TODO(b/320311538): replace link with the guide for customizing database image.
	//
	// nullon(dbs-fleet)
	// +optional
	DatabaseImage string `json:"databaseImage,omitempty"`
}

type DBCluster interface {
	cecoreapi.Entity
	PrimarySpec() InstanceSpec
	DBClusterSpec() *DBClusterSpec

	DBClusterStatus() *DBClusterStatus

	DBEngineShortName() string
	DBEngineName() string
	DBPorts() []corev1.ServicePort
	GetPrefixedName() string
	GetInternalName() string
}

// DBClusterList is a common interface for all the DBClusterList implentations
type DBClusterList interface {
	// DBClusters is a list of DBCluster objects
	DBClusters() []DBCluster
}

//+kubebuilder:object:generate=true

type PrimaryImportStatus struct {
	// ImportName is the Name of the latest import
	// +required
	ImportName string `json:"importName,omitempty"`
	// CreationTimeStamp represents the creation time of the import for the database instance
	// +optional
	CreationTimeStamp metav1.Time `json:"creationTimeStamp,omitempty"`
	// Spec represents the spec of the import for the database instance
	// +optional
	Spec eeimportapi.ImportSpec `json:"spec,omitempty"`
	// Status represents the of the latest import for the database instance
	// +optional
	Status eeimportapi.ImportStatus `json:"status,omitempty"`
}

//+kubebuilder:object:generate=true

type PrimaryExportStatus struct {
	// ExportName is the Name of the latest export
	// +required
	ExportName string `json:"exportName,omitempty"`
	// CreationTimeStamp represents the creation time of the export for the database instance
	// +optional
	CreationTimeStamp metav1.Time `json:"creationTimeStamp,omitempty"`
	// Spec represents the spec of the export for the database instance
	// +optional
	Spec eeexportapi.ExportSpec `json:"spec,omitempty"`
	// Status represents the of the latest import for the database instance
	// +optional
	Status eeexportapi.ExportStatus `json:"status,omitempty"`
}

//+kubebuilder:object:generate=true

type FailoverStatus struct {
	// FailoverName is the Name of the latest failover
	// +required
	FailoverName string `json:"failoverName,omitempty"`

	// Status represents status of the latest failover for the database cluster
	// +optional
	Status eehaapi.FailoverStatus `json:"status,omitempty"`
}

//+kubebuilder:object:generate=true

type MigrationStatus struct {
	// Status represents status of the migration for the database cluster
	// +optional
	*cecoreapi.ReplicationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:generate=true

type AvailabilityZones struct {
	// Primary represents the zone of the primary instance.
	// +optional
	Primary Zone `json:"primary,omitempty"`
	// Standbys represents the zone of the standby instances.
	// +optional
	Standbys []Zone `json:"standbys,omitempty"`
}

// Zone contains the zone information.
type Zone struct {
	Name string `json:"name,omitempty"`
}

type OpType string

const (
	OpImport OpType = "import"
	OpExport OpType = "export"
)
