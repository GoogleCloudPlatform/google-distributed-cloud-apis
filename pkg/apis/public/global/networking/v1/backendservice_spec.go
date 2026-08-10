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
	networkingv1 "gke-internal.googlesource.com/private-cloud/pkg/apis/public/networking/v1"
)

// Note: Global and Zonal BackendService does not share spec.
// When modifying this API, please be mindful of corresponding changes to the zonal BackendService API.

// Describes the attributes that a user expects from this backend service.
//
// +kubebuilder:validation:XValidation:rule="has(oldSelf.healthCheckName) == has(self.healthCheckName)", message="HealthCheckName is immutable"
type BackendServiceSpec struct {
	// A list of backends for this backend service. Only 1 Backend
	// can be specified per Zone or per User Cluster. This field is optional.
	// This field is mutable.
	//
	// +optional
	BackendRefs []BackendRef `json:"backendRefs,omitempty"`

	// A list of target ports that this BackendService will translate.
	// The provided port-protocol pair has to be unique in the list.
	// This field is optional.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MaxItems:=100
	// +kubebuilder:validation:XValidation:rule="self.all(a, self.exists_one(b, a.port == b.port && a.protocol == b.protocol))",message="TargetPorts has duplicate port-protocol pair"
	TargetPorts []networkingv1.TargetPort `json:"targetPorts,omitempty"`

	// A name of the health check parameters object for this backend service.
	// HealthCheck is applicable only for VM backends. It has to reference
	// HealthCheck in the same namespace as this backend service.
	// This field is optional. This field is immutable.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="HealthCheckName is immutable"
	// +kubebuilder:validation:MaxLength=250
	HealthCheckName *string `json:"healthCheckName,omitempty"`
}

// Holds information about the backend.
type BackendRef struct {
	// A name of the referenced Backend object. The referenced Backend has to be
	// in the same namespace as this backend service.
	// This field is required. This field is immutable.
	Name string `json:"name"`

	// A name of zone in which the referenced backend is created in.
	// This field is required. This field is immutable.
	Zone string `json:"zone"`
}
