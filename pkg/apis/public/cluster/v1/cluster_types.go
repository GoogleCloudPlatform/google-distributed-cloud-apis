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

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// Specifies a GDC user cluster in an air-gapped configuration.
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.clusterState",description="The current cluster state"
// +kubebuilder:printcolumn:name="K8s Version",type="string",JSONPath=".status.versionStatus.kubernetesVersion",description="The current cluster Kubernetes version"
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of clusters.
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Cluster `json:"items"`
}

type ClusterSpec struct {
	// The GDC air-gapped version information of the user cluster during cluster creation.
	// Optional. Default to use the latest applicable version. Immutable.
	InitialVersion *ClusterVersion `json:"initialVersion,omitempty"`
	// The release channel a cluster is subscribed to.
	// When a cluster is subscribed to a release channel, GDC maintains the
	// cluster versions for users.
	// Optional. Mutable.
	// +kubebuilder:default={channel:UNSPECIFIED}
	ReleaseChannel *ReleaseChannel `json:"releaseChannel,omitempty"`
	// The cluster network configuration.
	// If unset, the default configurations with pod and service CIDR sizes are used.
	// Optional. Mutable.
	// +kubebuilder:default={podCIDRSize:21,serviceCIDRSize:23}
	ClusterNetwork *ClusterNetwork `json:"clusterNetwork,omitempty"`
	// The load balancer configuration.
	// If unset, the default configuration with the ingress service IP address size is used.
	// Optional. Mutable.
	// +kubebuilder:default={ingressServiceIPSize:20}
	LoadBalancer *LoadBalancer `json:"loadBalancer,omitempty"`
	// The list of node pools for the cluster worker nodes.
	// Optional. Mutable.
	NodePools []*NodePool `json:"nodePools,omitempty"`
	// The cluster level configuration for all node pools.
	// Optional. Mutable.
	// +kubebuilder:validation:Optional
	NodePoolDefaults *NodePoolDefaults `json:"nodePoolDefaults,omitempty"`
}

type NodePoolDefaults struct {

	// The cluster level configuration for all nodes.
	// Optional. Mutable.
	// +kubebuilder:validation:Optional
	NodeConfigDefaults *NodeConfigDefaults `json:"nodeConfigDefaults,omitempty"`
}

type NodeConfigDefaults struct {

	// Specifies the containerd config for all node.
	// Optional. Mutable.
	// +kubebuilder:validation:Optional
	ContainerdConfig *ContainerdConfig `json:"containerdConfig,omitempty"`
}

type ContainerdConfig struct {

	// Specifies private registries used to pull images on the node.
	// Containerd is the only container runtime supported.
	// +kubebuilder:validation:Optional
	PrivateRegistries []*PrivateRegistry `json:"privateRegistries,omitempty"`
}

type PrivateRegistry struct {
	// Specifies the private registry host. This must consist of the host or host:port.
	// +kubebuilder:validation:Required
	Host string `json:"host"`
	// Specifies the secret that stores the private registry's CA bundle.
	// The secret must be either in the cluster namespace, or have the annotation
	// `baremetal.cluster.gke.io/mark-source` so it can be forwarded to the
	// cluster namespace.
	// +kubebuilder:validation:Optional
	CACertSecretRef *corev1.SecretReference `json:"caCertSecretRef,omitempty"`

	// Specifies the secret for the private registry access credential.
	// The secret must be either in the cluster namespace, or have the annotation
	// `baremetal.cluster.gke.io/mark-source` so it can be forwarded to the
	// cluster namespace.
	// +kubebuilder:validation:Optional
	PullCredentialSecretRef *corev1.SecretReference `json:"pullCredentialSecretRef,omitempty"`
}

// Specifies the version information of a GDC user cluster in an air-gapped
// configuration.
type ClusterVersion struct {
	// The Kubernetes version of the GDC user cluster.
	KubernetesVersion string `json:"kubernetesVersion"`
}

// Specifies the cluster network configuration.
type ClusterNetwork struct {
	// The size of network ranges from which pod virtual IP addresses are allocated.
	// If unset, a default value `21` is used.
	// +kubebuilder:default=21
	PodCIDRSize *int32 `json:"podCIDRSize,omitempty"`
	// The size of network ranges from which service virtual IP addresses are allocated.
	// If unset, a default value `23` is used.
	// +kubebuilder:default=23
	ServiceCIDRSize *int32 `json:"serviceCIDRSize,omitempty"`
}

// Specifies the load balancer configuration.
type LoadBalancer struct {
	// The size of non-overlapping IP pools used by the load balancer typed services.
	// If unset, a default value `20` is used.
	// +kubebuilder:default=20
	IngressServiceIPSize *int32 `json:"ingressServiceIPSize,omitempty"`
}

// Specifies the `NodePool` custom resource configuration.
type NodePool struct {
	// The name of the node pool.
	Name string `json:"name"`
	// The desired number of nodes in the provisioned node pool.
	NodeCount *int32 `json:"nodeCount,omitempty"`
	// The name of the machine types that are used to provision nodes.
	MachineTypeName string `json:"machineTypeName"`
	// The taints assigned to the nodes of this node pool.
	Taints []corev1.Taint `json:"taints,omitempty"`
	// The labels assigned to nodes of this node pool.
	// It contains a list of key/value pairs.
	Labels map[string]string `json:"labels,omitempty"`
	// AcceleratorOptions indicates the desired configuration of accelerators within the NodePool.
	// It's only valid if the chosen MachineType contains accelerators.
	AcceleratorOptions *AcceleratorConfig `json:"acceleratorOptions,omitempty"`
	// AutoScaling configuration for the node pool.
	// If this field is omitted or null, auto-scaling is disabled.
	// If present, auto-scaling is enabled according to the specified min/max nodes.
	// Optional. Mutable.
	AutoScaling *AutoScaling `json:"autoScaling,omitempty"`
}

// AutoScaling configuration for a node pool.
// The presence of this configuration block in the parent resource indicates that
// auto-scaling is enabled for the node pool.
type AutoScaling struct {
	// Maximum number of nodes in the NodePool.
	// Needs to be greater than 0.
	// +kubebuilder:validation:Minimum=1
	// Required.
	MaxNodes int `json:"maxNodes,omitempty"`

	// Minimum number of nodes in the NodePool.
	// Needs to be greater than 0.
	// +kubebuilder:validation:Minimum=1
	// Required.
	MinNodes int `json:"minNodes,omitempty"`
}

type AcceleratorConfig struct {
	// GPUPartitionScheme indicates the scheme that will be used to partition the GPUs into MIGs.
	// This scheme and GPU DeviceModel jointly decide the MIG profiles. For example,
	// mixed-1 on H100L 94GB together determines GPU to be partitioned into one 4g.47gb
	// and one 3g.47gb.
	GPUPartitionScheme string `json:"gpuPartitionScheme"`
}

// Indicates which release channel a cluster is subscribed to.
type ReleaseChannel struct {
	// If unset, defaults to be `UNSPECIFIED`.
	// +kubebuilder:default=UNSPECIFIED
	Channel *Channel `json:"channel,omitempty"`
}

// Indicates a specific type of release channel.
// +kubebuilder:validation:Enum=UNSPECIFIED
type Channel string

const (
	// Pauses the auto cluster version upgrades.
	Unspecified Channel = "UNSPECIFIED"
)

// Defines the observed state of the cluster.
type ClusterStatus struct {
	// The latest observations of the cluster state.
	// Conditions such as `Reconciling` and `Stalled` indicate whether the last cluster
	// reconciliation succeeded.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// The observed error status of the cluster.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
	// A list of observed statuses of the worker node pools.
	WorkerNodePoolStatuses []*NodePoolStatus `json:"workerNodePoolStatuses,omitempty"`
	// Whether the control plane is ready.
	ControlPlaneConditions []metav1.Condition `json:"controlPlaneConditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// Whether the user cluster components have successfully deployed.
	ComponentsConditions []metav1.Condition `json:"componentsConditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// The installed version information of the cluster.
	VersionStatus *VersionStatus `json:"versionStatus,omitempty"`
	// The state of the cluster. The following states are available:
	// <ul>
	// <li>`Running`: the cluster has been created and is usable.</li>
	// <li>`Reconciling`: some work is actively being done on the cluster.</li>
	// <li>`Deleting`: the cluster is being deleted.</li>
	// <li>`Error`: some errors occurred while reconciling/provisioning the
	// cluster.</li>
	// </ul>
	ClusterState *ClusterState `json:"clusterState,omitempty"`

	// Subnets holds the list of subnets that are used for this cluster.
	//
	// +kubebuilder:validation:Optional
	Subnets []SubnetReference `json:"subnets,omitempty"`
}

// SubnetReference holds the reference to a subnet.
type SubnetReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// Indicates the state of the cluster.
type ClusterState string

// Defines the installed version information of the cluster.
type VersionStatus struct {
	// The component version of the cluster.
	ComponentVersion *string `json:"componentVersion,omitempty"`
	// The Kubernetes version of the cluster.
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`
}

// Defines the observed state of a `NodePool` resource.
type NodePoolStatus struct {
	// The name of the node pool.
	Name string `json:"name"`
	// The latest observations of the node pool's state.
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
	// The number of nodes that are ready to serve.
	ReadyNodes *int32 `json:"readyNodes,omitempty"`
	// The number of nodes that are reconciling.
	ReconcilingNodes *int32 `json:"reconcilingNodes,omitempty"`
	// The number of nodes that are stalled.
	StalledNodes *int32 `json:"stalledNodes,omitempty"`
	// The number of nodes whose statuses are unknown.
	UnknownNodes *int32 `json:"unknownNodes,omitempty"`
	// The time a node pool is in `ready` status.
	// This value will never change once it's set.
	ReadyTimestamp *metav1.Time `json:"readyTimestamp,omitempty"`
	// The version of Kubernetes running on this node pool's nodes.
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
