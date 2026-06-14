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
	"k8s.io/apimachinery/pkg/util/sets"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// The desired VM running state.
// +kubebuilder:validation:Enum:=Running;Stopped
type VirtualMachineRunningState string

const (
	VMGroup = "vm." + Group

	// Indicates the VM running state.
	// A `VirtualMachine` instance gets respawned if the previous instance failed in an error state, or was shutdown within the guest.
	VirtualMachineRunningStateRunning VirtualMachineRunningState = "Running"
	// Indicates the VM stopped state.
	// No `VirtualMachine` instance will be present. If a guest is already running, it gets stopped.
	VirtualMachineRunningStateStopped VirtualMachineRunningState = "Stopped"
)

type VirtualMachineState string

// VirtualMachineTransitionKey represents a transition key, which is either a VirtualMachineState or a VirtualMachineStateReason.
type VirtualMachineTransitionKey string

// Valid states of a virtual machine.
const (
	// VirtualMachine is accepted by the system and is waiting for allocation.
	VirtualMachineStatePending VirtualMachineState = "Pending"
	// Indicates resources associated with the VirtualMachine are being provisioned and prepared.
	VirtualMachineStateProvisioning VirtualMachineState = "Provisioning"
	// VirtualMachine is being prepared for running.
	VirtualMachineStateStarting VirtualMachineState = "Starting"
	// VirtualMachine has started and is actively running.
	VirtualMachineStateRunning VirtualMachineState = "Running"
	// VirtualMachine is currently stopped and isn't expected to start until requested.
	VirtualMachineStateStopped VirtualMachineState = "Stopped"
	// VirtualMachine is in the process of being stopped.
	VirtualMachineStateStopping VirtualMachineState = "Stopping"
	// VirtualMachine is paused. The resources for the VirtualMachine is still
	// allocated but the guest OS is paused.
	VirtualMachineStatePaused VirtualMachineState = "Paused"
	// VirtualMachine is in the process of deletion, as well as its associated resources.
	VirtualMachineStateTerminating VirtualMachineState = "Terminating"
	// VirtualMachine has some configuration error.
	// The error could be temporal if it is waiting for dependent resources.
	VirtualMachineStateErrorConfiguration VirtualMachineState = "ErrorConfiguration"
	// VirtualMachine is waiting for an IP to be assigned to it.
	VirtualMachineStatePendingIPAllocation VirtualMachineState = "PendingIPAllocation"
	// VirtualMachine's status could not be obtained.
	VirtualMachineStateUnknown VirtualMachineState = "Unknown"
	// Indicates that an error has occurred while scheduling the virtual machine,
	// e.g. due to unsatisfiable resource requests or unsatisfiable scheduling constraints.
	VirtualMachineStateUnschedulable VirtualMachineState = "ErrorUnschedulable"
	// Indicates that an error with one or more of the attached disks of the VirtualMachine.
	VirtualMachineStateDiskError VirtualMachineState = "DiskError"
	// Indicates that the attached Disk is being readied.
	VirtualMachineStateDiskPending VirtualMachineState = "WaitingForDisk"
	// Indicates that the virtual machine is currently in a crash loop waiting to be retried.
	VirtualMachineStateCrashLoopBackoff VirtualMachineState = "CrashLoopBackoff"
	// Indicates that the virtual machine is in the process of being migrated to another host.
	VirtualMachineStateMigrating VirtualMachineState = "Migrating"
)

type VirtualMachineStateReason string

// LINT.IfChange
const (
	// Network the VM connects to is not found or is being deleted.
	NetworkNotFound VirtualMachineStateReason = "NetworkNotFound"
	// VirtualMachineType the VM refers to is not found or is being deleted.
	MachineTypeNotFound VirtualMachineStateReason = "MachineTypeNotFound"
	// VirtualMachineDisk the VM attaches is not found or is being deleted.
	MachineDiskNotFound VirtualMachineStateReason = "VirtualMachineDiskNotFound"
	// VirtualMachineDisk the VM attaches is being provisioned or is not ready.
	MachineDiskNotReady VirtualMachineStateReason = "VirtualMachineDiskNotReady"
	// VirtualMachineDisk the virtual machine disk is configured incorrectly.
	MachineDiskMisconfig VirtualMachineStateReason = "VirtualMachineDiskMisconfig"
	// VirtualMachineDisk the virtual machine disk image metadata not found.
	MachineDiskImageMetadataNotFound VirtualMachineStateReason = "VirtualMachineDiskImageMetadataNotFound"
	// VirtualMachine's network interface failed creation.
	InterfaceCreationFailed VirtualMachineStateReason = "InterfaceCreationFailed"
	// VirtualMachine's network interface failed update.
	InterfaceUpdateFailed VirtualMachineStateReason = "InterfaceUpdateFailed"
	// VirtualMachineProvisioningFailed indicates the VM provisioning failed.
	VirtualMachineProvisioningFailed VirtualMachineStateReason = "VirtualMachineProvisioningFailed"
	// Secret the VirtualMachine refers to is not found or does not have the expected key(s) in its data.
	ReferencedSecretDataNotFound VirtualMachineStateReason = "ReferencedSecretDataNotFound"
)

// LINT.ThenChange(/oc/vmm/subs/common-admin/internal/controllers/vm_health_monitor.go)

const (
	// Indicates whether the guest environment is enabled.
	ConditionTypeGuestEnvironmentAccessEnabled = "GuestEnvironmentAccessEnabled"
	// ConditionTypeGuestEnvironmentSynced means whether guest environment is synced.
	ConditionTypeGuestEnvironmentSynced = "GuestEnvironmentSynced"
	// ConditionTypeEditable means whether VM is editable.
	// Editable condition is true iff KubeVirt VM is not instantiated.
	ConditionTypeEditable = "Editable"
	// ConditionTypeVMReady means whether the VM is ready or not.
	// This is being added to be used in SLO calculations, where we determine good or bad states based on this.
	ConditionTypeVMReady = "Ready"

	// Reasons for 'GuestEnvironmentAccessEnabled' condition.
	// Indicates that the guest environment access was enabled by the user.
	ReasonUserConfiguration = "UserConfiguration"

	// Reasons for 'GuestEnvironmentSynced' condition.
	// Indicates that the GED status is not synced.
	ReasonGuestEnvironmentNotConnected = "GuestEnvironmentNotConnected"
	// Indicates that the GED is not found.
	ReasonAccessManagementUnavailable = "AccessManagementUnavailable"
	// Indicates that the GED generation is outdated.
	ReasonAccessManagementProgressing = "AccessManagementProgressing"
	// Indicates that the GED is synced.
	ReasonAccessManagementReady = "AccessManagementReady"
	// Indicates that the guest agent is fully ready and communicating.
	ReasonGuestAgentReady = "GuestAgentReady"

	// Reason for 'Editable' condition.
	ReasonVirtualMachineNotInstantiated = "VirtualMachineNotInstantiated"
	ReasonVirtualMachineInstantiated    = "VirtualMachineInstantiated"
	ReasonVirtualMachineStopped         = "VirtualMachineStopped"

	// ConditionTypeLiveMigratableOnEviction indicates whether the VM is eligible for
	// live migration during a node eviction.
	ConditionTypeLiveMigratableOnEviction = "LiveMigratableOnEviction"

	// Reasons for 'ConditionTypeLiveMigratableOnEviction' condition being False.
	ReasonNotMigratableSkipAnnotation   = "SkipLiveMigrationAnnotation"
	ReasonNotMigratableMaxAttempts      = "MaxMigrationAttemptsExceeded"
	ReasonNotMigratableWebhookRejection = "WebhookRejection"
)

const (
	// Specifies the firmware interface type BIOS.
	FirmwareTypeBIOS = "bios"
	// Specifies the firmware interface type UEFI.
	FirmwareTypeUEFI = "uefi"
)

// Defines the type of Confidential Computing technology of the VM.
type ConfidentialInstanceType string

// Specifies the Confidential Computing technology type TDX.
const TDXType ConfidentialInstanceType = "TDX"

var SupportedConfidentialInstanceType = sets.NewString(string(TDXType))

// Tracks the VirtualMachine provision time.
type VirtualMachineProvisionTime struct {
	// Time taken for the first VM provision. i.e. Time taken from the object being created
	// till the VM is in running status.
	// +optional
	InitProvisionTime *metav1.Duration `json:"initProvisionTime,omitempty"`

	// Time taken for the most recent VM provision. It can be equal to InitProvisionTime if the
	// VM is only being provisioned once.
	// +optional
	LastProvisionTime *metav1.Duration `json:"lastProvisionTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={gvm,gvms,vm,vms}
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +genclient
// +gdcloud:manifest:relevant=true,oc=vmm,component=compute,multigroup=true
// +gdcloud:manifest:entities="instances",verbs=create;delete;list;describe;start;stop;reset;add-metadata;remove-metadata,skipcodegen=true
// +gdcloud:manifest:rbac="list,describe:project-vm-viewer"
// +gdcloud:manifest:rbac="create,delete,list,describe,start,stop,reset,add-metadata,remove-metadata:project-vm-admin"
// +gdcloud:manifest:entities="instances",verbs=add-access-config;attach-disk;delete-access-config;update;update-access-config,skipcodegen=true
// +gdcloud:manifest:rbac="add-access-config,attach-disk,delete-access-config,update,update-access-config:project-vm-admin"
// Represents the Virtual Machine's configuration and state.
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec,omitempty"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// Defines the specification of the Virtual Machine.
type VirtualMachineSpec struct {
	// Indicates the requested running state of the VirtualMachine.
	// Defaults to Running.
	// +optional
	RunningState *VirtualMachineRunningState `json:"runningState,omitempty"`

	// Specifies the list of disks attached to this vm. There must be exactly
	// one boot disk. Changes to disk attachments require a reboot to take effect.
	Disks []DiskAttachment `json:"disks"`

	// Specifies the CPU and Memory of the VM.
	// CPU and Memory can be defined directly or through the VirtualMachineType.
	// Changes to Compute require a reboot to take effect.
	// Compute is immutable when the VM is in `Unknown` state.
	// +optional
	Compute Compute `json:"compute,omitempty"`

	// Specifies the configurations of the confidential VM.
	// +optional
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty"`

	// Specifies the list of startup scripts for the VM.
	// Linux VMs must have `cloud-init` installed for `StartupScripts` to take effect. The scripts
	// are executed in alphabetical order, based on the name of each startup script.
	// In Windows VMs the type of script impacts the order of execution. Scripts are executed in the
	// order ps1, bat, cmd. If two scripts are of the same type they are executed in alphabetical
	// order, based on the name.
	// +optional
	StartupScripts []StartupScript `json:"startupScripts,omitempty"`

	// Specifies the VM's guest environment configuration.
	// If the field is nil the `enable` field in `AccessManagement` is `true` by default.
	// Otherwise, the non-nil configuration for each sub-feature inside the structure
	// overrides the default configuration of the sub-feature.
	// +optional
	GuestEnvironment *GuestEnvironment `json:"guestEnvironment,omitempty"`

	// Specifies the VM initialization options at boot time.
	// +optional
	Firmware *Firmware `json:"firmware,omitempty"`

	// Specifies the VM's security-related configurations.
	// +optional
	ShieldConfig *ShieldConfig `json:"shieldConfig,omitempty"`

	// Specifies the network configuration.
	Network *NetworkSpec `json:"network,omitempty"`

	// TemplateRef specifies the `VirtualMachineTemplate` to use as a base
	// for this VM and refers to a `VirtualMachineTemplate` in the same namespace.
	// The properties defined in the template will be used to create the VM.
	// If other fields in the `VirtualMachineSpec` are set directly when `TemplateRef` is set,
	// the values specified in the `VirtualMachineSpec` will take precedence over the values sourced from the
	// `TemplateRef`.
	// +optional
	TemplateRef *corev1.LocalObjectReference `json:"templateRef,omitempty"`

	// GracefulShutdown configures graceful shutdown for the VM.
	// This section is mutable even after the VM is running.
	// If not specified, default connectionDrainTimeoutSeconds of 60 seconds will be applied.
	// +optional
	GracefulShutdown *GracefulShutdown `json:"gracefulShutdown,omitempty"`
}

// GracefulShutdown configures graceful shutdown for a VM.
type GracefulShutdown struct {
	// ConnectionDrainTimeoutSeconds specifies the grace period after a VM stop, eviction, or deletion is initiated and before the guest is shutdown.
	// +optional
	// +default=60
	// +kubebuilder:validation:Maximum=3600
	// +kubebuilder:validation:Minimum=1
	ConnectionDrainTimeoutSeconds *int32 `json:"connectionDrainTimeoutSeconds,omitempty"`
}

// Specifies the network configuration.
type NetworkSpec struct {
	// The network interfaces attached to the VM.
	// If no unicast interfaces are specified, a `default` interface is
	// automatically added.
	// Users in a multicast-enabled organization can add the `multicast` interface.
	// The first interface specified will be treated as the default
	// interface when setting up the default route inside the VM.
	Interfaces []NetworkInterfaceSpec `json:"interfaces,omitempty"`

	// Specifies the configuration tier for networking performance on all interfaces in the VM in GDC airgapped.
	// Unused in GDC connected.
	// +kubebuilder:default="Default"
	// +optional
	NetworkPerformanceTier NetworkPerformanceTier `json:"networkPerformanceTier,omitempty"`
}

const (
	NetworkDefault   = "default"
	NetworkMulticast = "multicast"
)

// +kubebuilder:validation:Enum=Default;Tier_1
type NetworkPerformanceTier string

const (
	// Networking configurations for default performance.
	NetworkPerformanceTierDefault NetworkPerformanceTier = "Default"
	// Advanced networking configurations for enhanced throughput performance.
	NetworkPerformanceTier1 NetworkPerformanceTier = "Tier_1"
)

// Specifies the network interface configuration.
// In GDC connected, only Network and IPAddresses can be specified.
// In GDC airgapped, one of Network, Subnet, or IPAddresses must be specified. Subnet and IPAddresses can optionally be specified together.
type NetworkInterfaceSpec struct {
	// The network that the interface is connected to.
	// In GDC airgapped, valid values are: `default`, `multicast`.
	// +optional
	Network string `json:"network,omitempty"`

	// The subnet that the interface is connected to in GDC airgapped. Unused in GDC connected.
	// If unspecified, defaults to the default subnet of the specified network.
	// +optional
	Subnet string `json:"subnet,omitempty"`

	// The namespace that the Subnet the interface is connected to in GDC airgapped. Unused in GDC connected.
	// If unspecified, defaults to the same namespace as the VirtualMachine.
	// +optional
	SubnetNamespace string `json:"subnetNamespace,omitempty"`

	// The IP address to be assigned to the interface. Only the first IP address is assigned to the interface.
	// In GDC airgapped, an IP address will be dynamically allocated if unspecified.
	// In GDC connected, if the network is configured to use an external DHCP server, this field can optionally be used to specify a static address. If the network is not configured to use an external DHCP server, this field is required.
	// +optional
	IPAddresses []IPAddress `json:"ipAddresses,omitempty"`
}

type IPAddress struct {
	// The IP address.
	// In GDC connected, the address may contain a subnet mask. If subnet mask is not included, /32 is taken as default. For example, 1.2.3.4 will be taken as 1.2.3.4/32. Alternatively, the input can be 1.2.3.4/24.
	// In GDC airgapped, the address may not contain a subnet mask. Subnet information is retrieved from the Subnet object.
	Address string `json:"address"`

	// Whether the lifecycle of the Subnet associated with this IP address should be managed by the system in GDC airgapped. Unused in GDC connected.
	// If true (default), a Subnet is created automatically for this IP and deleted during VM deletion.
	// If false, the user must have created a Subnet and the Subnet must be in ready status.
	// +optional
	Managed *bool `json:"managed,omitempty"`
}

// Defines a startup script for a VM.
// Supports the specification of a startup script either as a plain text string
// or a Kubernetes secret. If the field `script` is specified, then the field
// `scriptSecretRef` should not be provided, and vice versa.
type StartupScript struct {
	// Specifies the name of a script.
	// Must match the regex `[\w][\w\-.]*` and be at most 255 characters.
	// If specifying a script for a Windows VM, the name must include a '-'
	// followed by the script extension as a suffix. For example, use
	// the name `hello-world-ps1` for a Powershell script named `hello-world`.
	Name string `json:"name"`
	// Specifies a plain text string that contains the script.
	// The script content size must be lesser than 2048 bytes.
	// +kubebuilder:validation:MaxLength=2048
	Script string `json:"script,omitempty"`
	// `ScriptSecretRef` references a secret resource that contains the startup script.
	// When this StartupScript is defined within a VirtualMachine, `ScriptSecretRef`
	// must refer to a zonal Kubernetes `Secret` in the same namespace.
	// When defined within a VirtualMachineTemplate, `ScriptSecretRef` must refer to a
	// global Kubernetes `Secret` in the same namespace.
	// The name specified in `ScriptSecretRef` must exactly match the name of the
	// referenced `Secret` resource.
	ScriptSecretRef *corev1.LocalObjectReference `json:"scriptSecretRef,omitempty"`
}

// Specifies the guest environment configuration.
type GuestEnvironment struct {
	// Specifies the access management configuration.
	// +optional
	AccessManagement *AccessManagementConfig `json:"accessManagement,omitempty"`

	// Specifies the Guest Health configuration.
	// +optional
	GuestHealthCheck *GuestHealthCheck `json:"guestHealthCheck,omitempty"`
}

// Specifies the `AccessManagement` feature configuration in the guest environment.
type AccessManagementConfig struct {
	// Specifies whether to `enable` the `AccessManagement` feature in the
	// VM's guest environment. See the `GuestEnvironment` field description
	// for information about the default value of the field.
	Enable bool `json:"enable"`
}

// Specifies the `GuestHealth` feature configuration in the guest environment.
type GuestHealthCheck struct {
	// Specifies whether to `enable` the `GuestHealth` feature in the
	// VM's guest environment.
	Enable bool `json:"enable"`
}

// Specifies the configurations of the confidential VM.
type ConfidentialInstanceConfig struct {
	// Specifies the type of Confidential Computing technology.
	// +optional
	ConfidentialInstanceType ConfidentialInstanceType `json:"confidentialInstanceType,omitempty"`
}

// Specifies the VM initialization options at boot time.
type Firmware struct {
	// Specifies whether to boot via UEFI or BIOS.
	// Defaults to `bios`.
	// +kubebuilder:validation:Enum:=uefi;bios
	// +optional
	// Deprecated: Use ShieldConfig.BootType instead.
	Type string `json:"type,omitempty"`

	// Enables or disables boot loader certificate verification.
	// This is to assist in blocking modified or malicious code from loading.
	// The default value is `true` if `type` is `uefi`.
	// If `type` is set to `bios`, the default value is `false` and cannot be modified
	// since boot loader certificate verification is not available for BIOS.
	// +optional
	// Deprecated: Use ShieldConfig.EnableSecureBoot instead.
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`
}

// Specifies the VM's security-related configurations.
type ShieldConfig struct {
	// Specifies whether to boot via UEFI or BIOS.
	// Defaults to `bios`.
	// +kubebuilder:validation:Enum:=uefi;bios
	// +kubebuilder:default=bios
	// +optional
	BootType string `json:"bootType,omitempty"`

	// Enables or disables boot loader certificate verification.
	// This is to assist in blocking modified or malicious code from loading.
	// The default value is `true` if `bootType` is `uefi`.
	// If `bootType` is set to `bios`, the default value is `false` and cannot be modified
	// since boot loader certificate verification is not available for BIOS.
	// +optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Whether to emulate a VTPM device.
	// Defaults to `false`.
	// +kubebuilder:default=false
	// +optional
	EnableVtpm *bool `json:"enableVtpm,omitempty"`
}

// NetworkStatus is the status for the network of the Virtual Machine.
type NetworkStatus struct {
	// +optional
	Interfaces []NetworkInterfaceStatus `json:"interfaces,omitempty"`
}

// NetworkInterfaceStatus is the status for the NetworkInterface resource.
type NetworkInterfaceStatus struct {
	// Name denotes the name of the network interface exposed inside the VM, e.g. "eth0", "eth1".
	Name string `json:"name"`

	// IpAddresses are the IP addresses assigned to the NetworkInterface.
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty"`

	// MacAddress is the MAC address assigned to the NetworkInterface.
	// +optional
	MacAddress string `json:"macAddress,omitempty"`

	// Subnet is the Subnet assigned to the NetworkInterface.
	// +optional
	Subnet SubnetReference `json:"subnet,omitempty"`
}

type SubnetReference struct {
	// Name is the name of the Subnet.
	Name string `json:"name"`

	// Namespace is the namespace of the Subnet.
	Namespace string `json:"namespace"`
}

// Wrapper for all VMM errors, including error codes.
type VMMError struct {
	corev1alpha1.BaseError `json:",inline"`
}

// Contains the observed state of the Virtual Machine.
type VirtualMachineStatus struct {
	// Observed state of the VM.
	// +optional
	State VirtualMachineState `json:"state,omitempty"`

	// Reason why the VM is in the observed state.
	// Populated if applicable for the observed state.
	// +optional
	Reason VirtualMachineStateReason `json:"reason,omitempty"`

	// Additional details about the state of the VM.
	// +optional
	Message string `json:"message,omitempty"`

	// Status of the VM networks.
	// +optional
	Network NetworkStatus `json:"network,omitempty"`

	// Disks provides information about disks attached to the VirtualMachine.
	// Provided only for hotpluggable disks for now, but in the future may
	// potentially also be provided for non-hotpluggable disks.
	// +optional
	Disks *VirtualMachineAttachedDisksInfo `json:"disks,omitempty"`

	// Details of observed state.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Last transition time of each state.
	// +optional
	// Deprecated: Use TransitionTime instead.
	StateTransitionTime map[VirtualMachineState]metav1.Time `json:"stateTransitionTime,omitempty"`

	// Last transition time of each transition.
	// +optional
	// +listType=map
	// +listMapKey=transition
	TransitionTime []VirtualMachineTransitionTime `json:"transitionTime,omitempty"`

	// Time taken to provision the VM.
	// +optional
	ProvisionTime *VirtualMachineProvisionTime `json:"provisionTime,omitempty"`

	// FailureHistory tracks recent VM failures.
	// +optional
	FailureHistory *FailureHistory `json:"failureHistory,omitempty"`

	// A list of any errors that occurred during the reconciliation of this resource.
	// +optional
	Errors []VMMError `json:"errors,omitempty"`
}

// VirtualMachineTransitionTime represents a transition time with a key.
type VirtualMachineTransitionTime struct {
	// Transition is the transition key.
	// +kubebuilder:validation:Required
	Transition VirtualMachineTransitionKey `json:"transition"`

	// Time is the last transition time.
	// +kubebuilder:validation:Required
	Time metav1.Time `json:"time"`
}

type VirtualMachineAttachedDisksInfo struct {
	// Attachments provides status of disks attached to the VirtualMachine.
	// Provided only for hotpluggable disks for now, but in the future may
	// potentially also be provided for non-hotpluggable disks.
	// +optional
	Attachments []VirtualMachineAttachedDiskStatus `json:"attachments,omitempty"`
}

// VirtualMachineAttachedDiskStatus holds the observed status of a disk
// attached to a VirtualMachine.
type VirtualMachineAttachedDiskStatus struct {
	// Name refers to the name of the disk as defined in the
	// VirtualMachineSpec.Disks[*].VirtualMachineDiskRef.Name.
	Name string `json:"name"`

	// Hotpluggable reflects the disk attachment's underlying capability type
	// as part of the status.
	// A 'true' value indicates this is a hotplug attachment, which supports
	// the disk being added to and removed from the VM while it is running.
	// A 'false' or absent value indicates this is a coldplug attachment, which
	// means the disk was attached at boot and requires the VM to be stopped
	// for disk removal.
	// +optional
	Hotpluggable *bool `json:"hotpluggable,omitempty"`

	// Phase represents the current phase of the process of attachment
	// of the disk to the VM.
	// Provided only for hotpluggable disks for now, but in the future may
	// potentially also be provided for non-hotpluggable disks.
	// +optional
	Phase VirtualMachineAttachedDiskPhase `json:"phase,omitempty"`

	// Reason provides a brief description for the current disk attachment phase.
	// Provided only for hotpluggable disks for now, but in the future may
	// potentially also be provided for non-hotpluggable disks.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message provides a more detailed message about the current disk attachment
	// phase.
	// Provided only for hotpluggable disks for now, but in the future may
	// potentially also be provided for non-hotpluggable disks.
	// +optional
	Message string `json:"message,omitempty"`
}

// VirtualMachineAttachedDiskPhase represents the current phase of the
// process of attachment of the disk to the VM.
// Provided only for hotpluggable disks for now, but in the future may
// potentially also be provided for non-hotpluggable disks.
// +kubebuilder:validation:Enum:=Attaching;Ready;Detaching;Other;Unknown
type VirtualMachineAttachedDiskPhase string

// Constants for VirtualMachineAttachedDiskPhase.
const (
	// AttachedDiskPhaseAttaching indicates the disk attachment process is
	// in progress.
	AttachedDiskPhaseAttaching VirtualMachineAttachedDiskPhase = "Attaching"

	// AttachedDiskPhaseReady indicates the disk is ready to be used.
	AttachedDiskPhaseReady VirtualMachineAttachedDiskPhase = "Ready"

	// AttachedDiskPhaseDetaching indicates the disk detachment process is
	// in progress.
	AttachedDiskPhaseDetaching VirtualMachineAttachedDiskPhase = "Detaching"

	// AttachedDiskPhaseOther indicates some other phase of disk attachment,
	// other than those explicitly specified, and the details
	// of this state are not relevant to be shared at this time.
	AttachedDiskPhaseOther VirtualMachineAttachedDiskPhase = "Other"

	// AttachedDiskPhaseUnknown indicates that the disk attachment state
	// could not be definitively determined at this time.
	AttachedDiskPhaseUnknown VirtualMachineAttachedDiskPhase = "Unknown"
)

type FailureHistory struct {
	// Number of times VM has failed.
	FailureCounter int `json:"failureCounter,omitempty"`
	//+listType=atomic
	FailureInfos []FailureInfo `json:"failureInfos,omitempty"`
}

type FailureInfo struct {
	// Failure reason
	Reason string `json:"reason,omitempty"`
	// Time when failure was observed.
	ObservedTime metav1.Time `json:"observedTime,omitempty"`
}

// Contains a list of `VirtualMachine` objects.
// +kubebuilder:object:root=true
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

// GetTransitionTime returns the transition time for the specified key if present.
func (s *VirtualMachineStatus) GetTransitionTime(key VirtualMachineTransitionKey) (metav1.Time, bool) {
	for _, t := range s.TransitionTime {
		if t.Transition == key {
			return t.Time, true
		}
	}
	return metav1.Time{}, false
}

// SetTransitionTime sets or updates the transition time for the specified key.
func (s *VirtualMachineStatus) SetTransitionTime(key VirtualMachineTransitionKey, t metav1.Time) {
	for i, tt := range s.TransitionTime {
		if tt.Transition == key {
			s.TransitionTime[i].Time = t
			return
		}
	}
	s.TransitionTime = append(s.TransitionTime, VirtualMachineTransitionTime{
		Transition: key,
		Time:       t,
	})
}

func init() {
	SchemeBuilder.Register(
		&VirtualMachine{},
		&VirtualMachineList{},
	)
}
