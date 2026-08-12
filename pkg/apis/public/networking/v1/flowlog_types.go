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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Defines a list of flow log filters used for finding relevant
// flows. Flow events matching any of the provided filter rules are logged.
//
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +gdcloud:manifest:relevant=true,oc=unet,component=networking,entities="flow-logs"
// +gdcloud:manifest:verbs=create;delete;describe;list;update
// +gdcloud:manifest:rbac="create,delete,describe,list,update:flowlog-admin"
// +gdcloud:manifest:rbac="describe,list:flowlog-viewer"
// +gdcloud:manifest:skipcodegen=true
type FlowLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired configuration for a flow log.
	Spec FlowLogSpec `json:"spec,omitempty"`

	// The observed state of a flow log.
	Status FlowLogStatus `json:"status,omitempty"`
}

// Defines a list of flow log resources.
//
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FlowLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// A list of flow log entries.
	Items []FlowLog `json:"items,omitempty"`
}

// Represents the flow log filters that are applied.
// When multiple filters are present, the flow is logged if at least one filter
// matches the flow event.
type FlowLogSpec struct {
	// Specifies if this flow log is enabled. When disabled, the backend
	// flow filters are disabled, and no corresponding logs are collected.
	// Defaults to `true` if not specified.
	//
	// +kubebuilder:default:=true
	// +optional
	Enable *bool `json:"enable,omitempty"`

	// A list of filters used for matching flow events.
	// Flow events matching any of the provided filter rules are logged.
	//
	// +kubebuilder:validation:MinItems:=1
	Filters []FlowLogFilter `json:"filters"`

	// The amount of time this flow log rules is applied for. After the time is reached,
	// the flow logging rule is disabled.
	// If empty, this flow logging rule is enabled indefinitely.
	//
	// +optional
	Lifetime *Lifetime `json:"lifetime,omitempty"`

	// The fields that are logged for matching flow events.
	// If empty, defaults to logging all fields.
	//
	// +kubebuilder:default:=all
	// +optional
	LogDetailLevel *LogDetailLevel `json:"logDetailLevel,omitempty"`
}

// Defines the lifetime of a flow log. A value for `expiration` or
// `duration` must be specified, but not both.
//
// +kubebuilder:validation:MinProperties=1
// +kubebuilder:validation:MaxProperties=1
type Lifetime struct {
	// The time when this filter rule expires
	// and becomes inactive. Expiration must be a time in the future.
	// It includes the time required to propagate resources down to child
	// clusters so the value should account for an additional buffer of around
	// one minute to ensure that all clusters can begin logging and capture the
	// necessary traffic.
	//
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	// +optional
	Expiration *metav1.Time `json:"expiration,omitempty"`

	// The amount of time the flow log will be active for, starting from
	// when it is reconciled.
	// It includes the time required to propagate resources down to child
	// clusters so the value should account for an additional buffer of around
	// one minute to ensure that all clusters can begin logging and capture the
	// necessary traffic.
	//
	// +kubebuilder:validation:XValidation:rule="self == '' || duration(self) >= duration('1m')",message="Duration must be greater than 1 minute"
	// +optional
	Duration *metav1.Duration `json:"duration,omitempty"`
}

// Defines a collection of filter criteria that is applied at the
// same time.
// Each flow log filter contains several optional matching fields.
// The matching logic for each filter follows these rules:
// <br>
//
//	First, when a matching field is optional and not specified, it implies no
//	filtering is applied on this field of a flow.
//	For example, if no sources are provided, it means all sources are matched.
//
//	Next, when multiple fields are specified in one filter, all fields must match
//	the target flow.
//	For example, if  a source value of `srcNS/pod1` and a destination value of `dstNS/pod2` are
//	specified at the same time, it matches the flow from pod `srcNS/pod1`
//	to destination `dstNS/pod2`.
//
//	Finally, when a field is a list, specifying it multiple times means matching
//	any of the values.
//
// +kubebuilder:validation:XValidation:rule="!(has(self.endpoint) && (has(self.source) || has(self.destination)))", message="Endpoint cannot be set along with source or destination"
type FlowLogFilter struct {
	// A filter that filters flow events by a list of source rules.
	//
	// +optional
	Source *NetworkEndpointFilter `json:"source,omitempty"`

	// A filter that filters flow events by a list of destination rules.
	//
	// +optional
	Destination *NetworkEndpointFilter `json:"destination,omitempty"`

	// The endpoint filters flow events if the event source or destination matches
	// any given endpoint in this list.
	// If `endpoint` is set, `source` and `destination` must not be specified.
	// When specified, each endpoint corresponds to two filters: one with
	// `source` set to this endpoint and all other filter fields kept the same;
	// Another with `destination` set to this endpoint and all other filter
	// fields kept the same.
	//
	// +optional
	Endpoint *NetworkEndpointFilter `json:"endpoint,omitempty"`

	// A filter that filters flow events by L4 protocols defined in [v1.Protocol].
	// Each protocol must be specified at most once.
	//
	// +optional
	L4Protocols []v1.Protocol `json:"l4Protocols,omitempty"`

	// A filter that filters flow events by verdict classification.
	//
	// +optional
	Verdicts []PolicyVerdict `json:"verdicts,omitempty"`

	// A list of clusters and nodes used to match flows.
	//
	// +optional
	ClusterNodeSelectors []ClusterNodeSelector `json:"clusterNodeSelectors,omitempty"`
}

// Defines the observed state of flow logs.
type FlowLogStatus struct {
	// The current status of flow logs.
	// Known condition types are:
	// `Reconciled`: The flow log is reconciled and provisioned successfully;
	// and `Logging`: the flow log is currently reconciled and active.
	//
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The propagation status of this flow log in each cluster where
	// the resource is propagated.
	// The `Propagated` condition is set to `true` in the `Conditions` if this
	// resource is synced to the cluster, and its `ObservedGeneration` is set
	// to the generation of the propagated resource in the target cluster.
	// If this resource is successfully pruned from a cluster,
	// the corresponding `PropagationStatus` must be removed from the list.
	//
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=cluster
	// +listMapKey=node
	// +listMapKey=namespace
	// +listMapKey=name
	// +optional
	Clusters []PropagationStatus `json:"clusters,omitempty"`

	// The time the flow log becomes active.
	// This field is set by reconciler when it first interacts with the object, or when the `FlowLogSpec` resource
	// is changed.
	//
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`

	// The time the flow log becomes inactive.
	// This field is set by reconciler as such:
	// <ol>
	// <li>When the flow has infinity life time (`FlowLogSpec.Lifetime` is
	//     unspecified), `EndTime` may be empty.</li>
	// <li>When `FlowLogSpec.Lifetime.Expiration`` is set, its value is copied
	//     to `EndTime` directly.</li>
	// <li>When `FlowLogSpec.Lifetime.Duration` is set, `EndTime` will be set to
	//     the value of `StartTime` added to the value of `FlowLogSpec.Lifetime.Duration`</li>
	// </ol>
	// +optional
	EndTime *metav1.Time `json:"endTime,omitempty"`
}

const (
	// FlowLogConditionReconciled indicates the flow log is reconciled and
	// provisioned successfully.
	FlowLogConditionReconciled = "Reconciled"
	// FlowLogConditionLogging indicates the flow log is currently reconciled and
	// active (within lifetime).
	FlowLogConditionLogging = "Logging"
)

const (
	// FlowLogConditionLoggingReasonPending indicates the flow log is not logging currently as its start time is not current.
	FlowLogConditionLoggingReasonPending = "Pending"
	// FlowLogConditionLoggingReasonActive indicates the flow log is logging now.
	FlowLogConditionLoggingReasonActive = "Active"
	// FlowLogConditionLoggingReasonDisabled indicates the flow log is not logging now as it's disabled.
	FlowLogConditionLoggingReasonDisabled = "Disabled"
	// FlowLogConditionLoggingReasonPruning indicates the flow log is not logging now as it's being pruned.
	FlowLogConditionLoggingReasonPruning = "Pruning"
	// FlowLogConditionLoggingReasonExpired indicates the flow log is not logging now as it has expired.
	FlowLogConditionLoggingReasonExpired = "Expired"
)

const (
	// FlowLogConditionReconciledReasonSucceeded indicates the flow log has reconciled successfully.
	FlowLogConditionReconciledReasonSucceeded = "Succeeded"
	// FlowLogConditionReconciledReasonFailed indicates the flow log failed to reconcile.
	FlowLogConditionReconciledReasonFailed = "Failed"
	// FlowLogConditionReconciledReasonFailed indicates the flow log is being pruned.
	FlowLogConditionReconciledReasonPruning = "Pruning"
)

// Represents the information used to locate a node or nodes inside of a specified cluster.
// Either cluster or node or both must be specified.
//
// +kubebuilder:validation:MinProperties=1
type ClusterNodeSelector struct {
	// The name of the cluster.
	// If a value is not provided, all clusters will be searched for the desired node or nodes.
	//
	// +optional
	Cluster *string `json:"cluster,omitempty"`

	// A wildcard pattern used to search by the node name.
	// For example, `k8s*` or `*.domain.com`.
	//
	// +optional
	Node *string `json:"node,omitempty"`
}

// Defines a list of verdict classifying flows.
//
// +enum
type PolicyVerdict string

const (
	// PolicyVerdictUnknown is used if there is no verdict for this flow event
	PolicyVerdictUnknown PolicyVerdict = "unknown"
	// PolicyVerdictForwarded is used for flow events where the trace point has forwarded
	// this packet or connection to the next processing entity.
	PolicyVerdictForwarded PolicyVerdict = "forwarded"
	// PolicyVerdictDropped is used for flow events where the connection or packet has
	// been dropped (e.g. due to a malformed packet, it being rejected by a
	// network policy etc). The exact drop reason may be found in drop_reason_desc.
	PolicyVerdictDropped PolicyVerdict = "dropped"
	// PolicyVerdictError is used for flow events where an error occurred during processing
	PolicyVerdictError PolicyVerdict = "error"
	// PolicyVerdictAudit is used on policy verdict events in policy audit mode, to
	// denominate flows that would have been dropped by policy if audit mode
	// was turned off
	PolicyVerdictAudit PolicyVerdict = "audit"
	// PolicyVerdictRedirected is used for flow events which have been redirected to the proxy
	PolicyVerdictRedirected PolicyVerdict = "redirected"
	// PolicyVerdictTraced is used for flow events which have been observed at a trace point,
	// but no particular verdict has been reached yet
	PolicyVerdictTraced PolicyVerdict = "traced"
	// PolicyVerdictTranslated is used for flow events where an address has been translated
	PolicyVerdictTranslated PolicyVerdict = "translated"
)

// A list of predefined combinations of fields that are logged
// when filtered flow events are captured.
//
// +enum
type LogDetailLevel string

const (
	// LogDetailLevelAll logs all fields available.
	LogDetailLevelAll LogDetailLevel = "all"

	// LogDetailLevelLegacy logs the following fields:
	//   - time
	//   - protocol
	//   - source/destination IP addresses
	//   - source/destination port
	LogDetailLevelLegacy LogDetailLevel = "legacy"

	// LogDetailLevelK8SShort logs the following fields:
	//   - time
	//   - source/destination namespace
	//   - source/destination pod
	LogDetailLevelK8SShort LogDetailLevel = "k8s_short"

	// LogDetailLevelK8SLong logs all fields selected by "k8s_short", plus the
	// following fields:
	//   - drop_reason
	//   - verdict
	//   - traffic_direction
	//   - is_reply
	LogDetailLevelK8SLong LogDetailLevel = "k8s_long"
)

// Defines the propagation status for a specific cluster.
type PropagationStatus struct {
	// The cluster name where this resource is propagated.
	Cluster string `json:"cluster"`
	// The node name where this resource is propagated.
	Node string `json:"node"`
	// The namespace where this resource is propagated.
	Namespace string `json:"namespace"`
	// The name of the propagated resource.
	Name string `json:"name"`

	// The current status of the programmed resources.
	//
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&FlowLog{}, &FlowLogList{})
}
