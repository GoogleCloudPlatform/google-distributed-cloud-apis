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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
	projectv1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/resourcemanager/v1alpha1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// Represents a description of a Marketplace service.
type ServiceDescriptionReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceDescriptionSpec          `json:"spec,omitempty"`
	Status ServiceDescriptionReplicaStatus `json:"status,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Version",type=string,JSONPath=`.spec.version`
// Represents a single version of a Marketplace service.
type ServiceVersionReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceVersionSpec          `json:"spec"`
	Status ServiceVersionReplicaStatus `json:"status,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// Represents a single group of Marketplace services.
type ServiceCatalogReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceCatalogSpec          `json:"spec,omitempty"`
	Status ServiceCatalogReplicaStatus `json:"status,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// Represents a ServiceCatalog to Project mapping relations.
type ServiceCatalogBindingReplica struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceCatalogBindingSpec          `json:"spec,omitempty"`
	Status ServiceCatalogBindingReplicaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceDescriptionReplica` custom resources.
type ServiceDescriptionReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceDescriptionReplica `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceVersionReplica` custom resources.
type ServiceVersionReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceVersionReplica `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceCatalogReplica` custom resources.
type ServiceCatalogReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceCatalogReplica `json:"items"`
}

// +kubebuilder:object:root=true

// Represents a collection of `ServiceCatalogBindingReplica` custom resources.
type ServiceCatalogBindingReplicaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ServiceCatalogBindingReplica `json:"items"`
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

type ServiceDescriptionReplicaStatus struct{}

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

type ServiceVersionReplicaStatus struct{}

// ServiceCatalogSpec defines catalog owner information and a list of accessible Service Descriptions.
type ServiceCatalogSpec struct {
	// Display name for this ServiceCatalog.
	DisplayName string `json:"displayName"`

	// Catalog Owner Information.
	Owner CatalogOwner `json:"owner"`

	// List of Service Descriptions Included in the ServiceCatalog.
	ServiceDescriptionRef []corev1alpha1.NamespacedName `json:"serviceDescriptionRef"`
}

// CatalogOwner represents Catalog Owner Information including the Catalog Admin Account and the Catalog Owning Team.
type CatalogOwner struct {
	// Catalog Admin Account.
	Admin string `json:"admin"`

	// Catalog Owning Team.
	Team string `json:"team"`
}

type ServiceCatalogReplicaStatus struct{}

// Provides the specification, or desired state, of a ServiceCatalogBinding resource.
type ServiceCatalogBindingSpec struct {
	// ServiceCatalogRef represents the ServiceCatalog in this ServiceCatalogBinding.
	ServiceCatalogRef corev1alpha1.NamespacedName `json:"serviceCatalogRef"`

	// Selector is used to specify a set of rules to match Projects.
	Selector ServiceCatalogBindingSelector `json:"selector,omitempty"`
}

// Provides a set of rules to match Projects.
// Must choose exactly 0 or 1 of the selectors.
// 0 selector matches all Projects.
// +kubebuilder:validation:MaxProperties=1
type ServiceCatalogBindingSelector struct {
	// NameSelector is used to specify the list of names of Projects.
	NameSelector *NameSelector `json:"nameSelector,omitempty"`

	// LabelSelector is used to specify the list of label and values of Projects.
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
}

// Provides a list of Project Name for Service Catalog Binding to match with.
type NameSelector struct {
	MatchNames []string `json:"matchNames,omitempty"`
}

// Matches checks if the given Project matches the ServiceCatalogBindingSelector.
func (s *ServiceCatalogBindingSelector) Matches(project *projectv1alpha1.Project) (bool, error) {
	switch {
	case s.NameSelector != nil:
		for _, name := range s.NameSelector.MatchNames {
			if name == project.Name {
				return true, nil
			}
		}
		return false, nil

	case s.LabelSelector != nil:
		labelSelector, err := metav1.LabelSelectorAsSelector(s.LabelSelector)
		if err != nil {
			return false, fmt.Errorf("failed to parse LabelSelector: %w", err)
		}

		return labelSelector.Matches(labels.Set(project.GetLabels())), nil

	default: // matches all Projects if no selector is set.
		return true, nil
	}
}

type ServiceCatalogBindingReplicaStatus struct{}

func init() {
	SchemeBuilder.Register(
		&ServiceDescriptionReplica{},
		&ServiceVersionReplica{},
		&ServiceCatalogReplica{},
		&ServiceCatalogBindingReplica{},
		&ServiceDescriptionReplicaList{},
		&ServiceVersionReplicaList{},
		&ServiceCatalogReplicaList{},
		&ServiceCatalogBindingReplicaList{},
	)
}
