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
	apiext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// ServiceDescription represents a description of a Marketplace service.
type ServiceDescription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServiceDescriptionSpec `json:"spec,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.version`
// Represents a single version of a Marketplace service.
type ServiceVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServiceVersionSpec `json:"spec"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Desired Version",type=string,JSONPath=`.spec.serviceVersionRef.name`
// +kubebuilder:printcolumn:name="Cluster",type=string,JSONPath=`.spec.clusterRef.name`
// +kubebuilder:printcolumn:name="Installed Version",type=string,JSONPath=`.status.installedVersion`
// Represents an installed instance of a Marketplace service.
type ServiceInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceInstanceSpec   `json:"spec"`
	Status ServiceInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceDescription` custom resources.
type ServiceDescriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceDescription `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceVersion` custom resources.
type ServiceVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceVersion `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceInstance` custom resources.
type ServiceInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceInstance `json:"items"`
}

// ServiceDescriptionSpec defines a service description of a Marketplace service.
type ServiceDescriptionSpec struct {
	// Selector value corresponding to "marketplace.gdc.goog/service-name" label.
	ServiceName string `json:"serviceName"`
	// Display name for this service.
	DisplayName string `json:"displayName"`
	// Vendor for this service.
	Vendor string `json:"vendor"`
	// Short description for this service.
	Description string `json:"description"`
	// Type of this service for Overview page.
	ServiceType string `json:"serviceType"`
	// Type of icon asset.
	IconType IconType `json:"iconType"`
	// Source data for the icon asset, e.g. base64-encoded image.
	IconSource string `json:"iconSource"`
	// Service overview HTML-formatted section.
	Overview string `json:"overview"`
	// Service pricing HTML-formatted section.
	Pricing ServiceDescriptionDetails `json:"pricing"`
	// Service support HTML-formatted section.
	Support string `json:"support"`
	// Service terms HTML-formatted section.
	Terms ServiceDescriptionDetails `json:"terms"`
	// Service contact info HTML-formatted section.
	ContactInfo string `json:"contactInfo"`
	// Service documentation details.
	Documentation ServiceDescriptionDetails `json:"documentation"`
	// Additional details for "Categories".
	Categories []string `json:"categories"`
}

// ServiceDescriptionDetails defines the fields available for a given section.
type ServiceDescriptionDetails struct {
	// Description details HTML-formatted section.
	Details string `json:"details"`
	// Description URL, if available.
	// +optional
	Url string `json:"url,omitempty"`
}

// IconType represents the type of icon asset.
// +kubebuilder:validation:Enum=IMAGE;SVG_ICON
type IconType string

const (
	// Icon asset types
	IconTypeImage IconType = "IMAGE"
	IconTypeSVG   IconType = "SVG_ICON"

	// Service types
	HelmBased string = "helm"
	LinkBased string = "link"
)

// Defines an available version of a Marketplace service.
type ServiceVersionSpec struct {
	// The version of this service.
	Version string `json:"version"`

	// The name of the project where service Helm is stored.
	HarborProject string `json:"harborProject,omitempty"`

	// The name of the Helm chart whose lifecycle will be managed.
	Entrypoint string `json:"entrypoint"`

	// The default configuration of the Helm chart values.
	DefaultConfiguration string `json:"defaultConfiguration,omitempty"`
}

// Defines an installed instance of a Marketplace service.
type ServiceInstanceSpec struct {
	// ServiceVersionRef references the ServiceVersion for this instance.
	ServiceVersionRef corev1alpha1.NamespacedName `json:"serviceVersionRef"`

	// The name of the cluster to install the service instance.
	ClusterRef corev1alpha1.NamespacedName `json:"clusterRef"`

	// TargetNamespace specifies an optional namespace where the service instance should be installed.
	// +optional
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// The parameters to configure for the service instance.
	// This can contain arbitrary JSON data.
	Parameters *apiext.JSON `json:"parameters,omitempty"`
}

// Defines the observed state of a Marketplace service instance.
type ServiceInstanceStatus struct {
	// The currently running version.
	// An empty string indicates no version is installed.
	InstalledVersion string `json:"installedVersion"`

	// The indicator for whether the installation is successful, or if any errors
	// occurred.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ErrorStatus holds the most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

var (
	// The name of the service name label for a service.
	MarketplaceServiceVersionNameLabel = GroupVersion.Group + "/service-name"

	// The name of the feature flag annotation for a service.
	MarketplaceFeatureFlagAnnotation = GroupVersion.Group + "/service-flag"

	// Annotation to indicate the minimum required nodes.
	RequireMinimumNodes = GroupVersion.Group + "/require-minimum-nodes"

	// Annotation to indicate the minimum required CPU cores.
	RequireMinimumCpuCores = GroupVersion.Group + "/require-minimum-cpu-cores"

	// Annotation to indicate the minimum required memory.
	RequireMinimumMemory = GroupVersion.Group + "/require-minimum-memory"
)

func init() {
	SchemeBuilder.Register(
		&ServiceVersion{},
		&ServiceInstance{},
		&ServiceDescription{},
		&ServiceVersionList{},
		&ServiceInstanceList{},
		&ServiceDescriptionList{},
	)
}
