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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	rmv1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/resourcemanager/v1alpha1"
)

// Note: Global and Zonal FlowTagger does not share spec.
// When modifying this API, please be mindful of corresponding changes to the global FlowTagger API.

// Defines the Schema for the `FlowTagger` API.
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName={ft,fts}
// +gdcloud:manifest:relevant=false,oc=unet
type FlowTagger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the desired configuration for flow tagger resource.
	Spec FlowTaggerSpec `json:"spec,omitempty"`

	// Status for the flow tagger.
	Status FlowTaggerStatus `json:"status,omitempty"`
}

// Defines the specification or expected state of the `FlowTagger` resource.
type FlowTaggerSpec struct {
	// Source for the flow tagger.
	// +kubebuilder:validation:Optional
	Source *FlowTaggerEntity `json:"source,omitempty"`

	// Destination for the flow tagger.
	// +kubebuilder:validation:Optional
	Destination *FlowTaggerEntity `json:"destination,omitempty"`

	// Protocol for the flow tagger.
	// If not specified, it defaults to ALL.
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=ALL
	// +kubebuilder:validation:Enum=TCP;UDP;ALL
	Protocol FlowTaggerProtocol `json:"protocol,omitempty"`

	// Trace ID.
	// If not specified, a value will be assigned by the backend.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Optional
	TraceID int32 `json:"traceID,omitempty"`
}

// FlowTaggerProtocol defines the protocol for the flow tagger.
type FlowTaggerProtocol string

const (
	// FlowTaggerProtocolTCP is for TCP protocol.
	FlowTaggerProtocolTCP FlowTaggerProtocol = "TCP"
	// FlowTaggerProtocolUDP is for UDP protocol.
	FlowTaggerProtocolUDP FlowTaggerProtocol = "UDP"
	// FlowTaggerProtocolALL is for all protocols.
	FlowTaggerProtocolALL FlowTaggerProtocol = "ALL"
)

// Defines the properties of the flow tagger entity either source or destination. Exactly one of `resourceInfo` or `ip` must be specified.
// +kubebuilder:validation:XValidation:rule="has(self.resourceInfo) != has(self.ip)", message="Exactly one of resourceInfo or ip must be specified"
type FlowTaggerEntity struct {
	// IP address assigned to flow tagger entity.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Format=ip
	IP string `json:"ip,omitempty"`

	// Port number.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Optional
	Port uint16 `json:"port,omitempty"`

	// ResourceInfo provides more information about flow tagger resource.
	// This is mutually exclusive with the IP field.
	// +kubebuilder:validation:Optional
	ResourceInfo *ResourceInfo `json:"resourceInfo,omitempty"`
}

// ResourceInfo contains information about a flow tagger resource.
// +kubebuilder:validation:XValidation:rule="size(self.name) > 0 && size(self.kind) > 0 && size(self.__namespace__) > 0 && size(self.clusterInfo.name) > 0 && size(self.clusterInfo.kind) > 0", message="name, kind, namespace, clusterInfo.name, and clusterInfo.kind are required and must not be empty."
// +kubebuilder:validation:XValidation:rule="self.kind in ['Pod', 'Service']", message="kind must be either Pod or Service"
type ResourceInfo struct {
	corev1.TypedObjectReference `json:",inline"`

	// ClusterRef provides a reference to the cluster where the resource is located.
	ClusterRef corev1.TypedObjectReference `json:"clusterInfo"`
}

// Defines the observed state of the `FlowTagger` resource.
type FlowTaggerStatus struct {
	// The observed state of the `FlowTagger` resource.
	// If `ready` is `true`, it means that flow tagger is propagated to all clusters successfully.
	// If `ready` is `false`, it means that it failed to propagate flow tagger to one or more clusters.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The name of the propagated flow tagger realized in all remote clusters within the project.
	PropagatedName string `json:"propagatedName,omitempty"`

	// The list of propagation status on the clusters.
	Clusters []rmv1alpha1.ClusterStatus `json:"clusters,omitempty"`

	// Trace ID assigned.
	TraceID int32 `json:"traceID,omitempty"`
}

// Contains a list of `FlowTagger` objects.
//
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
type FlowTaggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FlowTagger resources.
	Items []FlowTagger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlowTagger{}, &FlowTaggerList{})
}
