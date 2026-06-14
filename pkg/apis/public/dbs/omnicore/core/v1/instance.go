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
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// InstanceHARolePrimary represents the primary role of an instance in HA setup.
	InstanceHARolePrimary = "Primary"

	// InstanceHARoleStandby represents the standby role of an instance in HA setup.
	InstanceHARoleStandby = "Standby"
)

// Service is a service provided by the operator.
type Service string

// InstanceMode describes how an instance will be managed by the operator.
type InstanceMode string

type ImageOSType string

const (
	// DebianOSType represents the standard Debian-based image.
	DebianOSType ImageOSType = "Debian"

	// UBI9OSType represents the Universal Base Image (UBI) version 9.
	UBI9OSType ImageOSType = "UBI9"
)

var DefaultEnabledServices map[Service]bool = map[Service]bool{
	Monitoring:       true,
	BackupAndRestore: true,
	Security:         true,
	Logging:          true,
	Patching:         true,
}

const (
	// Monitoring service provides the ability to collect
	// monitoring data from the database and the cluster.
	Monitoring Service = "Monitoring"

	// BackupAndRestore service provides database backups and restore functionalities.
	BackupAndRestore Service = "Backup"

	// Security service
	Security Service = "Security"

	// Logging service
	Logging Service = "Logging"

	// Patching service provides software and database patching.
	Patching Service = "Patching"

	// ManuallySetUpStandby means that operator will skip DB creation during
	// provisioning, instance will be ready for users to manually set up standby.
	ManuallySetUpStandby InstanceMode = "ManuallySetUpStandby"

	// Pause Mode means the instance will stop processing incoming API calls and
	// terminate any pending LRO operation after a grace period
	Pause InstanceMode = "Pause"

	Recovery InstanceMode = "Recovery"

	Maintenance InstanceMode = "Maintenance"

	// AdminPasswordTimeoutMinute is the timeout period for admin password
	AdminPasswordTimeoutMinute time.Duration = 60 * time.Minute

	// PVCs created by StatefulSet following the following pattern:
	// volumeClaimTemplate.Name + stsName + ordinal
	PVCNamePattern = "%s-%s-%d"

	// StatefulSet replica count for Instance in different state.
	StoppedReplicaCnt = 0
	NormalReplicaCnt  = 1

	// TaskTypeDatabase is used in pod labels to identify the pod is running a database engine.
	TaskTypeDatabase = "database"
	// TaskTypeMonitoring is used in pod labels to identify the pod is running monitoring.
	TaskTypeMonitoring = "monitoring"

	// TaskTypeMonitoringSvc is used in pod labels for identifying the pod running the monitoring svc (which is accessible by the customer)
	TaskTypeMonitoringSvc = "monitoringsvc"

	// TaskTypeOmniSvcNodeMgr is used in pod labels to identify the pod is running the omnisvc node manager.
	TaskTypeOmniSvcNodeMgr = "node-manager"

	// TaskTypeImageCatalog is used in resource labels to identify the config maps containing image digests.
	TaskTypeImageCatalog = "image-catalog"

	// Image names in InstanceSpec.Images
	// TODO(b/290761374): renaming the image keys
	ImageMonitor    = "monitoring"
	ImageLogging    = "logging"
	ImageLogrotate  = "logrotate"
	ImageDatabase   = "database"
	ImageInit       = "dbinit"
	ImageBackupInit = "backupinit"

	DefaultHealthcheckPeriodSeconds = 30

	DBContainer     = "database"
	DBInitContainer = "dbinit"

	DBSMetricProxySidecarContainerName = "otel-collector"
)

type InstanceComponentName string

const (
	Dataplane InstanceComponentName = "Dataplane"

	ControlPlaneAgents InstanceComponentName = "ControlPlaneAgents"
)

//+kubebuilder:object:generate=true

type InstanceComponentSpec struct {
	// Name of a component
	Name InstanceComponentName `json:"name"`

	// Version of a component
	//+optional
	Version string `json:"version"`

	// The list of container images in the component
	Images map[string]string `json:"images"`

	// OSType explicitly specifies the base operating system type (e.g., Debian, UBI9)
	// for images used by this component.
	// +optional
	// +kubebuilder:validation:Enum=Debian;UBI9
	// nullon(dbs-fleet,dbs-local)
	OSType ImageOSType `json:"osType,omitempty"`

	// Start time of the upgrade
	//+optional
	UpgradeScheduledAt *metav1.Time `json:"upgradeScheduledAt,omitempty"`
}

//+kubebuilder:object:generate=true

// InstanceSpec represents the database engine agnostic
// part of the spec describing the desired state of an Instance.
type InstanceSpec struct {
	// Version of a database.
	// +required
	// nullon(samwise-fleet,samwise-local)
	Version string `json:"version,omitempty"`

	// Deprecated: Replacement images for the database instance.
	//
	// nullon(samwise-fleet,samwise-local)
	// +optional
	Images map[string]string `json:"images,omitempty"`

	// Port is the port number that the database is listening on.
	// nullon(dbs-fleet,dbs-local)
	// +optional
	// +kubebuilder:validation:Minimum=1024
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:default=5432
	Port int `json:"port,omitempty"`

	// The list of instance components
	// An instance is composed of dataplane and controlPlaneAgent components
	// +optional
	Components map[InstanceComponentName]InstanceComponentSpec `json:"component,omitempty"`

	// DBNetworkServiceOptions allows to override some details of kubernetes
	// Service created to expose a connection to database.
	// +optional
	DBLoadBalancerOptions *DBLoadBalancerOptions `json:"dbLoadBalancerOptions,omitempty"`

	// Source IP CIDR ranges allowed for a client.
	// +optional
	SourceCidrRanges []string `json:"sourceCidrRanges,omitempty"`

	// Parameters allows to set database parameters for the database cluster. This field is
	// optional.
	//
	// Parameters will take a key/value pair corresponding to the parameter name/value as defined
	// by the database engine.
	//
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`

	// Services list the optional semi-managed services that
	// the customers can choose from.
	// +optional
	Services map[Service]bool `json:"services,omitempty"`

	// Resource specification for the database container.
	//
	// When any of the fields inside the resource changes, the operator restarts the
	// database instance with the new resource specification.
	//
	// +kubebuilder:validation:Required
	Resources *Resource `json:"resources"`

	// Mode specifies how this instance will be managed by the operator.
	// +optional
	// +kubebuilder:validation:Enum=ManuallySetUpStandby;Pause;Recovery;Maintenance
	Mode InstanceMode `json:"mode,omitempty"`

	// The maximum duration allowed for patching the StatefulSet during container image updates.
	// This threshold ensures that the database pods are updated within a specific time frame.
	// If this field is unset, it will fall back to using the default maximum duration.
	// +optional
	DatabasePatchingTimeout *metav1.Duration `json:"databasePatchingTimeout,omitempty"`

	// AdminUser represents the admin user specification. This field is required.
	//
	// This is the initial database user that the control plane creates. Additional
	// database users are managed by the end-user directly. This field can also be used to
	// reset the password of the initial user.
	AdminUser *AdminUserSpec `json:"adminUser,omitempty"`

	// SystemUserPasswordRefs is a list of system users and the corresponding secrets containing the password for those accounts
	// nullon(dbs-fleet,dbs-local)
	// +optional
	SystemUserPasswordRefs map[string]string `json:"systemUserPasswordRefs,omitempty"`

	// IsStopped stops the instance when set to true. This field is optional and default to false.
	//
	// When stopped, the compute resources (CPU, memory) of the instance are released. However, the instance
	// still keeps the storage resource and network endpoints so that restarting is transparent to the
	// downstream services.
	// See the status field for success or failures, if any.
	// +optional
	IsStopped *bool `json:"isStopped,omitempty"`

	// AvailabilityOptions contains adjustable settings for HA features
	// +optional
	AvailabilityOptions *AvailabilityOptions `json:"availabilityOptions,omitempty"`

	// +kubebuilder:default=false
	// +optional
	AllowExternalIncomingTrafficToInstance bool `json:"allowExternalIncomingTrafficToInstance,omitempty"`

	// AuditLogTarget configures the sink for the database audit logs
	// +optional
	AuditLogTarget *AuditLogTargetSpec `json:"auditLogTarget,omitempty"`

	// Replication configures replication connections to other db instances
	//
	// nullon(samwise-fleet)
	// +optional
	Replication *ReplicationSpec `json:"replication,omitempty"`

	// TLS is the desired server certificate configuration for the instance. This field is optional.
	// When this field is changed, the instance pods will restart to load the specified certificate
	//
	// +optional
	TLS *TLSSpec `json:"tls,omitempty"`

	// SchedulingConfig specifies how the instance should be scheduled on Kubernetes nodes.
	//
	// When any field inside the scheduling config changes, it can lead to rescheduling of the
	// k8s pod onto a different node based on the config.
	//
	// +optional
	SchedulingConfig *SchedulingConfig `json:"schedulingconfig,omitempty"`
}

// +kubebuilder:object:generate=true
type SchedulingConfig struct {
	// NodeAffinity describes node affinity scheduling rules for the instance.
	// +optional
	NodeAffinity *corev1.NodeAffinity `json:"nodeaffinity,omitempty"`
	// PodAffinity describes pod affinity scheduling rules for the instance.
	// +optional
	PodAffinity *corev1.PodAffinity `json:"podAffinity,omitempty"`
	// PodAntiAffinity describes pod anti-affinity scheduling rules for the instance.
	// +optional
	PodAntiAffinity *corev1.PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	// Tolerations to enable the management of whether to allow or disallow scheduling an
	// instance on a Kubernetes node that has a specific taint applied.
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
	// TopologySpreadConstraints describes how to spread pods across topology domains.
	// nullon(dbs-fleet)
	// +optional
	TopologySpreadConstraints []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
}

//+kubebuilder:object:generate=true

// WalArchiveSpec specifies wal archive settings.
type WalArchiveSpec struct {
	// ArchiveMode specifies archive_mode, see https://www.postgresql.org/docs/current/runtime-config-wal.html#GUC-ARCHIVE-MODE for details.
	// +kubebuilder:validation:Enum:=on;always
	// +kubebuilder:default:=on
	// +optional
	ArchiveMode string `json:"archiveMode,omitempty"`

	// Location is the location where archived wal logs are stored.
	Location string `json:"location,omitempty"`
}

// WalArchiveStatus represents the current wal archive settings.
type WalArchiveStatus struct {
	Location string `json:"location,omitempty"`
}

//+kubebuilder:object:generate=true

// TLSSpec contains secrets used for encrypted communication with the database instance.
type TLSSpec struct {
	// CertSecret contains the name of a certificate secret within the same namespace.
	// The secret must contain entries ca.crt (CA certificate), tls.key (server private key),
	// and tls.crt (server leaf certificate). This secret is used to set the TLS config
	// for the database instance.
	CertSecret *corev1.LocalObjectReference `json:"certSecret,omitempty"`

	// DataPlaneCertIssuer refers to the issuer representing the CA that will sign
	// data plane certificates, like for the database server. This field is
	// mutually exclusive with the `tls.certSecret` field.
	// nullon(dbs-fleet,dbs-local)
	DataPlaneCertIssuer *IssuerReference `json:"dataPlaneCertIssuer,omitempty"`

	// ControlPlaneAgentsCertIssuer refers to the issuer representing the CA that will
	// sign certificates for control plane agent components.
	// nullon(dbs-fleet,dbs-local)
	ControlPlaneAgentsCertIssuer *IssuerReference `json:"controlPlaneAgentsCertIssuer,omitempty"`

	// DataPlaneCertRequest contains attributes for data plane certificates.
	// This includes certificates for the database server, pgBackRest server, and
	// pgBouncer auth query client. Updating these fields after initial creation
	// will result in certificate rotation.
	// nullon(dbs-fleet,dbs-local)
	DataPlaneCertRequest *DataPlaneCertificateRequest `json:"dataPlaneCertRequest,omitempty"`

	// ControlPlaneAgentsCertRequest contains attributes for control plane agent
	// certificates. This includes certificates for internal components that
	// communicate securely with the Operator. Updating these fields after initial
	// creation will result in certificate rotation.
	// nullon(dbs-fleet,dbs-local)
	ControlPlaneAgentsCertRequest *CertificateRequest `json:"controlPlaneAgentsCertRequest,omitempty"`
}

//+kubebuilder:object:generate=true

// TLSStatus contains the active resources used for encrypted communication
// with the database instance.
type TLSStatus struct {
	// CertSecret contains the name of the active database server certificate
	// secret within the same namespace.
	//
	// If `spec.TLS` is provided, this value should refer to the same secret after the
	// database has been configured to use the provided server certificate.
	CertSecret *corev1.LocalObjectReference `json:"certSecret,omitempty"`

	// ControlPlaneAgentsCertIssuer refers to the issuer that was used to sign
	// certificates for control plane agent components.
	// nullon(dbs-fleet,dbs-local)
	ControlPlaneAgentsCertIssuer *IssuerReference `json:"controlPlaneAgentsCertIssuer,omitempty"`

	// ControlPlaneAgentsCertRequest refers to the certificate request fields that
	// were used to provision certificates for control plane agent components.
	// nullon(dbs-fleet,dbs-local)
	ControlPlaneAgentsCertRequest *CertificateRequest `json:"controlPlaneAgentsCertRequest,omitempty"`
}

// +kubebuilder:object:generate=true
type AdminUserSpec struct {
	// PasswordRef is the name of the secret containing the admin user's password. This value will be used during initial
	// provisioning or password reset to set the admin user to that password. The secret must be under the same project
	// as the Database cluster.
	// The name of the secret must follow this pattern `db-pw-<dbc name>`. Additionally, the key of the password (inside the secret) must
	// be the same as the database cluster name.
	// +optional
	PasswordRef *corev1.LocalObjectReference `json:"passwordRef,omitempty"`
}

// PasswordRef contains the secret information and the key to get the admin user password
type PasswordRef struct {
	// SecretRef is a reference to the secret that contains admin user password
	// +optional
	SecretRef corev1.SecretReference `json:"secretRef,omitempty"`
	// PasswordKey is the key used to search the secret for the password
	// +optional
	PasswordKey string `json:"passwordKey,omitempty"`
}

// AvailabilityOptions contains customization options for instance-level HA features
type AvailabilityOptions struct {
	// LivenessProbe enables or disables the liveness probe which is used to trigger a container restart.
	// When set to `Enabled`, the liveness probe runs periodic health checks on the database. It restarts the container if it fails three consecutive health checks. LivenessProbe is automatically disabled for HA instances.
	// When set to `Disabled`, the liveness probe is not running health checks on the database. The default value is Enabled.
	// +kubebuilder:validation:Enum=Enabled;Disabled;OpDisabled
	// +kubebuilder:default:=Enabled
	// +optional
	LivenessProbe string `json:"livenessProbe,omitempty"`

	// HealthcheckPeriodSeconds is the number of seconds the healthcheck prober will wait before checking the health of the primary and standby instances again and updating the status accordingly.
	// This field is propagated down from the DBCluster's spec
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=86400
	// +kubebuilder:default=30
	// +optional
	HealthcheckPeriodSeconds int `json:"healthcheckPeriodSeconds"`
}

// DBLoadBalancerOptions contains customization options for the Kubernetes
// LoadBalancer exposing database connections.
// +kubebuilder:object:generate=true
type DBLoadBalancerOptions struct {
	// GCP contains Google Cloud specific attributes for the Kubernetes LoadBalancer.
	// +optional
	GCP *DBLoadBalancerOptionsGCP `json:"gcp,omitempty"`
	// OnPrem contains On-Prem specific attributes for the Kubernetes LoadBalancer.
	// +optional
	OnPrem *DBLoadBalancerOptionsOnPrem `json:"onprem,omitempty"`
	// Annotation provided by the customer will be added to the service object of type loadbalancer.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// DBLoadBalancerOptionsGCP contains GCP specific options for the Kubernetes
// LoadBalancer created for database connections.
type DBLoadBalancerOptionsGCP struct {
	// A LoadBalancer can be internal or external.
	// See https://kubernetes.io/docs/concepts/services-networking/service/#internal-load-balancer
	// +kubebuilder:validation:Enum="";Internal;External
	// +optional
	LoadBalancerType string `json:"loadBalancerType,omitempty"`

	// LoadBalancerIP is a static IP address, see
	// https://cloud.google.com/compute/docs/ip-addresses/reserve-static-external-ip-address
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty"`

	// LoadBalancerInterface is the network interface to be used by the LoadBalancer.
	// +optional
	LoadBalancerInterface string `json:"loadBalancerInterface,omitempty"`
}

// DBLoadBalancerOptionsOnPrem contains OnPrem specific options for the
// LoadBalancer created for database connections.
type DBLoadBalancerOptionsOnPrem struct {
	// A LoadBalancer can be internal or external.
	// +kubebuilder:validation:Enum="";Internal;External
	// +optional
	LoadBalancerType string `json:"loadBalancerType,omitempty"`

	// LoadBalancerIP is a virtual IP address, see
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty"`

	// LoadBalancerInterface is the network interface to be used by the LoadBalancer.
	// +optional
	LoadBalancerInterface string `json:"loadBalancerInterface,omitempty"`
}

//+kubebuilder:object:generate=true

// RestoredFrom is the status showing the most recent restore source for current Instance.
type RestoredFrom struct {
	// Source Instance this Instance restores from.
	SourceInstance string `json:"sourceInstance,omitempty"`

	// Time point of the source Instance this Instance restores from.
	RestoredTime *metav1.Time `json:"restoredTime,omitempty"`
}

//+kubebuilder:object:generate=true

// AuditLogTargetSpec defines the available sinks for database audit logs
type AuditLogTargetSpec struct {
	Syslog *AuditLogSyslogTargetSpec `json:"syslog,omitempty"`
}

//+kubebuilder:object:generate=true

type AuditLogSyslogTargetSpec struct {
	// Host is the syslog server FQDN or IP address
	Host string `json:"host"`

	// CertsSecretRef contains the certificates to be used for the TLS connection to syslog server
	CertsSecretRef *corev1.SecretReference `json:"certsSecretRef"`
}

//+kubebuilder:object:generate=true

type InstanceComponentStatus struct {
	// Name of a component
	Name InstanceComponentName `json:"name"`

	// Version of a component
	Version string `json:"version"`

	// The list of container images in the components
	Images map[string]string `json:"images"`

	// OSType is the base operating system type (e.g., Debian, UBI)
	// for images used by this component.
	// +optional
	// +kubebuilder:validation:Enum=Debian;UBI9
	// nullon(dbs-fleet,dbs-local)
	OSType ImageOSType `json:"osType,omitempty"`

	// Start time of the upgrade
	// +optional
	UpgradeScheduledAt *metav1.Time `json:"upgradeScheduledAt,omitempty"`
}

//+kubebuilder:object:generate=true

// InstanceStatus defines the observed state of Instance
type InstanceStatus struct {
	// Phase is a summary of current state of the Instance.
	// +optional
	Phase InstancePhase `json:"phase,omitempty"`

	// Endpoint is presently expressed in the format of <instanceName>-svc.<ns>.
	Endpoint string `json:"endpoint,omitempty"`

	// URL represents an IP and a port number info needed in order to
	// establish a database connection from outside a cluster.
	URL string `json:"url,omitempty"`

	// Description is for a human consumption.
	// E.g. when an Instance is restored from a backup
	// this field is populated with the human readable
	// restore details.
	Description string `json:"description,omitempty"`

	// InstanceObservedGeneration is the latest generation observed by the controller.
	// +optional
	InstanceObservedGeneration int64 `json:"instanceObservedGeneration,omitempty"`

	// IsChangeApplied indicates whether instance changes have been applied
	// +optional
	IsChangeApplied metav1.ConditionStatus `json:"isChangeApplied,omitempty"`

	// EntityStatus represents the status of an Entity.
	EntityStatus `json:",inline"`

	// RestoredFrom shows the most recent restore source for current Instance.
	// +optional
	RestoredFrom *RestoredFrom `json:"restoredFrom,omitempty"`

	// CurrentParameters stores the last successfully set database parameters.
	CurrentParameters map[string]string `json:"currentParameters,omitempty"`

	// LastFailedParameterUpdate is used to avoid getting into the failed
	// parameter update loop.
	LastFailedParameterUpdate map[string]string `json:"lastFailedParameterUpdate,omitempty"`

	// ExternalConnectivity represents the external connectivity details instance.
	ExternalConnectivity *Connectivity `json:"externalConnectivity,omitempty"`

	// InternalIP represents the internal connectivity details of the instance.
	InternalConnectivity *Connectivity `json:"internalConnectivity,omitempty"`

	// ActiveImages stores the stable images used by the active containers.
	ActiveImages map[string]string `json:"ActiveImages,omitempty"`

	// LastFailedImages stores the images which failed the last patching workflow.
	LastFailedImages map[string]string `json:"LastFailedImages,omitempty"`

	// ActiveComponents stores the information of current components in the database instance
	// +optional
	ActiveComponents map[InstanceComponentName]InstanceComponentStatus `json:"ActiveComponents,omitempty"`

	// ReplicationStatus represents the current state of replication connections.
	ReplicationStatus *ReplicationStatus `json:"ReplicationStatus,omitempty"`

	// AllocatedResources represents the current configuration of memory/CPU/disks
	AllocatedResources *Resource `json:"allocatedResources,omitempty"`

	// WalArchiveSetting represents the current wal archive settings.
	// +optional
	// nullon(dbs-fleet,dbs-local)
	WalArchiveSetting *WalArchiveStatus `json:"walArchiveSetting,omitempty"`

	// AdminUser represents the status of database admin user.
	// +optional
	AdminUser *AdminUserStatus `json:"adminUser,omitempty"`

	// AlloyDbAdminUser represents the status of database admin user.
	// +optional
	// nullon(dbs-fleet,dbs-local)
	AlloyDbAdminUser *AdminUserStatus `json:"alloyDbAdminUser,omitempty"`

	// TLS is the current server certificate configuration for the instance.
	// +optional
	TLS *TLSStatus `json:"tls,omitempty"`

	// NodeManager holds status information about the instance's Node Manager.
	NodeManager *NodeManagerStatus `json:"nodeManager,omitempty"`
}

type AdminUserStatus struct {
	// PasswordResourceVersion is the Password Secret's resourceVersion when the
	// password was last updated on the database.
	// +optional
	PasswordResourceVersion string `json:"passwordResourceVersion,omitempty"`
}

//+kubebuilder:object:generate=true

type Resource struct {
	// The amount of memory allocated to the database container.
	Memory resource.Quantity `json:"memory,omitempty"`

	// The amount of CPU allocated to the database container.
	Cpu resource.Quantity `json:"cpu,omitempty"`

	// Priority of the database process. Range is -20 to 19.
	// Lower value indicates higher priority.
	// nullon(dbs-fleet,dbs-local,samwise-fleet,samwise-local)
	// +optional
	Priority *int32 `json:"priority,omitempty"`

	// The machine type of the VM that database container runs on.
	// nullon(samwise-fleet,samwise-local)
	// +optional
	MachineType string `json:"machineType,omitempty"`

	// The specifications of the disks allocated to the database container. This field is required.
	Disks []DiskSpec `json:"disks,omitempty"`

	// Limits describes the maximum compute resources allowed. Risk of OOMKill if memory usage exceeds the
	// request when the Kubernetes node hosting the database container has memory pressure.
	// This field is optional.
	// Setting cpu and/or memory limit to 0 leaves cpu and/or memory limit unset on the database container.
	// Learn more about not setting memory limit at https://kubernetes.io/docs/tasks/configure-pod-container/assign-memory-resource/#if-you-do-not-specify-a-memory-limit.
	// +optional
	// nullon(dbs-fleet,dbs-local)
	Limits *ResourceList `json:"limits,omitempty"`
}

// ToK8sResourceRequirements converts Resource to k8s corev1.ResourceRequirements.
func (r *Resource) ToK8sResourceRequirements() corev1.ResourceRequirements {
	req := corev1.ResourceRequirements{
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    r.Cpu,
			corev1.ResourceMemory: r.Memory,
		},
	}

	if r.Limits != nil {
		req.Limits = corev1.ResourceList{}
		if !r.Limits.Cpu.IsZero() {
			req.Limits[corev1.ResourceCPU] = r.Limits.Cpu
		}
		if !r.Limits.Memory.IsZero() {
			req.Limits[corev1.ResourceMemory] = r.Limits.Memory
		}
	} else {
		// If limits are not specified, default them to requests.
		req.Limits = corev1.ResourceList{
			corev1.ResourceCPU:    r.Cpu,
			corev1.ResourceMemory: r.Memory,
		}
	}
	return req
}

// ResourceList defines the compute resources for a component.
// +kubebuilder:object:generate=true
type ResourceList struct {
	// The amount of memory allocated to the database container. This field is optional.
	Memory resource.Quantity `json:"memory,omitempty"`

	// The amount of CPU allocated to the database container. This field is optional.
	Cpu resource.Quantity `json:"cpu,omitempty"`
}

// +kubebuilder:object:generate=true
// Connectivity stores the connectivity details of the Instance
type Connectivity struct {
	URL string `json:"url,omitempty"`
	IP  string `json:"IP,omitempty"`
}

type ServiceAccountRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// +kubebuilder:object:generate=true
type NodeManagerStatus struct {
	// Addr holds the Node Manager endpoint for the instance.
	// If can be either in the format "<host>:<port>" or just "<host>".
	// If port is not specified then the default NodeManager port will be used.
	Addr string `json:"addr,omitempty"`

	// SystemUsers contains the status of system users used by the node manager.
	SystemUsers map[string]NodeManagerSystemUserStatus `json:"systemUsers,omitempty"`
}

// +kubebuilder:object:generate=true
type NodeManagerSystemUserStatus struct {
	// PasswordUpdatedAt is the time when the password was last updated on the node manager.
	PasswordUpdatedAt metav1.Time `json:"passwordUpdatedAt,omitempty"`

	// PasswordResourceVersion is the resource version of the secret containing
	// the password the last time it was used to update the password on the node
	// manager.
	PasswordResourceVersion string `json:"passwordResourceVersion,omitempty"`
}

// Instance represents the contract for the Anthos DB Operator compliant
// database Operator providers to abide by.
type Instance interface {
	Entity
	InstanceSpec() *InstanceSpec
	InstanceStatus() *InstanceStatus
	GetPrefixedName() string
}

// InstanceList represents a db engine agnostic list of instances
type InstanceList interface {
	client.ObjectList
	Instances() []client.Object
}

type InstanceDbPlugin interface {
	NewInstance() Instance
	NewInstanceList() client.ObjectList

	NewInstanceHealth() InstanceHealth

	NewInstanceBackupPlan() InstanceBackupPlan
	NewInstanceBackupPlanList() InstanceBackupPlanList
	NewInstanceBackup() InstanceBackup

	NewInstanceRestore() InstanceRestore

	GetInstance(ctx context.Context, c client.Client, name, namespace string) (Instance, error)

	// This API is used by clone to change the PVC names of the origin dbcluster to match PVC names of the target dbcluster.
	// The returned map should be the map[PVCFullName]=PVCShortName
	GetPVCNames(instName string) map[string]string
	GetStsName(instName string) string
	GetInstanceGroupKind() metav1.GroupKind

	GetDBClusterLabelForInstance() string
	GetDBClusterNSLabelForInstance() string
	DBEngine() string
	DBEngineShortName() string
	GetCRLockFinalizerName() string

	OperatorNS() string
	CertResourceName() string
	DBPort(instance Instance) int
	DefaultReplicationUsername(version string) string
	ServiceAccountName(dbcName string, cpaVersion string) string
}

func InstanceToClientObject(i Instance) (client.Object, error) {
	obj, ok := i.(client.Object)
	if !ok {
		return nil, fmt.Errorf("couldn't convert to client Object; %+q", obj)
	}
	return obj, nil
}

// NewInstance creates a zero-value Instance object from the same GroupVersion
// as the object given in the parameters.
func NewInstance(scheme *runtime.Scheme, gvObj runtime.Object) (Instance, error) {
	return CreateZeroObjFromSameGroupVersion[Instance](scheme, gvObj, "Instance")
}

// NewInstanceList creates a zero-value InstanceList object from the same
// GroupVersion as the object given in the parameters.
func NewInstanceList(scheme *runtime.Scheme, gvObj runtime.Object) (InstanceList, error) {
	return CreateZeroObjFromSameGroupVersion[InstanceList](scheme, gvObj, "InstanceList")
}

func InitializeAvailabilityOptionsStruct(availabilityOptions *AvailabilityOptions) *AvailabilityOptions {
	if availabilityOptions.HealthcheckPeriodSeconds == 0 {
		availabilityOptions.HealthcheckPeriodSeconds = DefaultHealthcheckPeriodSeconds
	}
	return availabilityOptions
}
