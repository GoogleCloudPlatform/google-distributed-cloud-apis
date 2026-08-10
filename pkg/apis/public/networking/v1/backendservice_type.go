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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "gke-internal.googlesource.com/private-cloud/pkg/apis/core/v1alpha1"
)

// Note: Global and Zonal BackendService does not share spec.
// When modifying this API, please be mindful of corresponding changes to the global BackendService API.

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=bes

// Represents a load balancer configuration.
//
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="HealthCheck",type="string",JSONPath=".spec.healthCheckName"
type BackendService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   BackendServiceSpec   `json:"spec"`
	Status BackendServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Contains a list of BackendService.
type BackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendService `json:"items"`
}

// Describes the attributes that a user expects from this backend service.
//
// +kubebuilder:validation:XValidation:rule="has(oldSelf.healthCheckName) == has(self.healthCheckName)", message="HealthCheckName is immutable"
type BackendServiceSpec struct {

	// A list of backends for this backend service. Only 1 Backend
	// can be specified per Zone or per User Cluster. This field is optional.
	// This field is mutable.
	//
	// +optional
	// +kubebuilder:validation:MaxItems:=50
	BackendRefs []BackendRef `json:"backendRefs,omitempty"`

	// A list of target ports that this BackendService will translate.
	// The provided port-protocol pair has to be unique in the list.
	// This field is optional.
	//
	// +optional
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MaxItems:=100
	// +kubebuilder:validation:XValidation:rule="self.all(a, self.exists_one(b, a.port == b.port && a.protocol == b.protocol))",message="TargetPorts has duplicate port-protocol pair"
	TargetPorts []TargetPort `json:"targetPorts,omitempty"`

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
}

// Holds information about an L4 port that will be translated to
// specified targetPort.
type TargetPort struct {
	Port `json:",inline"`

	// A port to which the Port value will be translated to.
	// This field is required.
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Required
	TargetPort int32 `json:"targetPort"`
}

// Aggregates the health status of all Backends in this BackendService.
type BackendServiceHealth struct {
	// Overall health of the Backends in this BackendService.
	// Possible values: Healthy, Partial, Unhealthy, Draining.
	// Healthy: All endpoints across all backends are Ready.
	// Unhealthy: All endpoints across all backends are Unhealthy.
	// Partial: A mix of Ready, Unhealthy, or Draining endpoints.
	// Draining: Some endpoints are Draining, but none are Unhealthy.
	// +optional
	AggregateState string `json:"aggregateState,omitempty"`

	// Total number of "available" (Ready) endpoints
	// out of the total number of endpoints across all BackendRefs.
	// Draining endpoints are considered as not available.
	// Example: "8/9"
	// +optional
	AvailabilityRatio string `json:"availabilityRatio,omitempty"`

	// Total number of endpoints resolved for this Backend.
	// +optional
	TotalEndpoints int32 `json:"totalEndpoints,omitempty"`

	// Number of endpoints in a Ready state.
	// +optional
	HealthyEndpoints int32 `json:"healthyEndpoints,omitempty"`

	// Number of endpoints in a Not Ready state.
	// +optional
	UnhealthyEndpoints int32 `json:"unhealthyEndpoints,omitempty"`

	// Number of endpoints in a Draining state.
	// +optional
	DrainingEndpoints int32 `json:"drainingEndpoints,omitempty"`

	// Last time the health state was changed.
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
}

// Represents the status of BackendService.
type BackendServiceStatus struct {
	// A list of conditions describing the current state of the backend service.
	// Known condition types are:
	// * "Ready"
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Holds the most recent errors with last seen time.
	//
	// +optional
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`

	// Aggregated and detailed health status for backend endpoints.
	// +optional
	BackendServiceHealth *BackendServiceHealth `json:"backendServiceHealth,omitempty"`

	// A list of forwarding rules using this backend service.
	//
	// +optional
	ForwardingRuleRefs []ForwardingRuleRef `json:"forwardingRuleRefs,omitempty"`

	// A name of the active BackendServicePolicy that applies to this BackendService
	// This field is optional.
	//
	// +optional
	// +kubebuilder:validation:Optional
	BackendServicePolicyRef *string `json:"backendServicePolicyRef,omitempty"`
}

// Holds information about the forwarding rule.
type ForwardingRuleRef struct {
	// A name of the referenced forwarding rule object.
	// This field is required. This field is immutable.
	Name string `json:"name"`
}

// BackendServiceReadyConditionReason defines the set of reasons that explain
// why a particular BackendService Ready condition status is set.
type BackendServiceReadyConditionReason string

const (
	// HealthCheckNotFound indicates that the referenced HealthCheck object
	// cannot be found in the API.
	HealthCheckNotFound BackendServiceReadyConditionReason = "HealthCheckNotFound"
	// BackendNotFound indicates that the one or more of the referenced Backend
	// objects cannot be found in the API.
	BackendNotFound BackendServiceReadyConditionReason = "BackendNotFound"
	// HealthCheckNotReady indicates that the referenced HealthCheck object
	// is not Ready.
	HealthCheckNotReady BackendServiceReadyConditionReason = "HealthCheckNotReady"
	// BackendNotReady indicates that the referenced Backend object
	// is not Ready.
	BackendNotReady BackendServiceReadyConditionReason = "BackendNotReady"
	// BackendConflict indicates that the referenced Backend objects are
	// conflicting. You can specify only 1 Backend per Zone (Project) or
	// 1 per Cluster
	BackendConflict BackendServiceReadyConditionReason = "BackendConflict"
	// AggregateStateHealthy indicates that all endpoints across all Backends are Ready.
	AggregateStateHealthy string = "HEALTHY"
	// AggregateStateUnhealthy indicates that all endpoints across all Backends are Unhealthy.
	AggregateStateUnhealthy string = "UNHEALTHY"
	// AggregateStatePartial indicates a mix of Ready, Unhealthy, or Draining endpoints.
	AggregateStatePartial string = "PARTIAL"
	// AggregateStateDraining indicates that some endpoints are Draining, but none are Unhealthy.
	AggregateStateDraining string = "DRAINING"
)

func init() {
	SchemeBuilder.Register(&BackendService{}, &BackendServiceList{})
}
