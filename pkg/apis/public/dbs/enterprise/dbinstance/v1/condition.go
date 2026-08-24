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
	occoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

var (
	// ConditionDBNodesReady indicates whether all DBNodes required for this
	// DBInstance have been provisioned and are ready.
	ConditionDBNodesReady occoreapi.ConditionType = "DBNodesReady"
	// ReadonDBNodesNotReady indicates that the minimum number of DBNodes are not ready.
	// This is used with a False status for the condition.
	ReasonDBNodesNotReady occoreapi.ConditionReason = "DBNodesNotReady"
	// ReasonMinDBNodesReady indicates that at least the minimum number of
	// DBNodes (as specified in the spec) are ready. This is used with a True
	// status for the condition.
	ReasonMinDBNodesReady occoreapi.ConditionReason = "MinDBNodesReady"
	// ReasonAllDBNodesReady indicates that the DBInstance has the required
	// number of ready DBNodes (as specified in the spec). This is used with a
	// True status for the condition.
	ReasonAllDBNodesReady occoreapi.ConditionReason = "AllDBNodesReady"

	// ConditionStreaming indicates whether the DBNodes are actively replicating
	// from the primary.
	ConditionStreaming occoreapi.ConditionType = "Streaming"
	// ReasonNotStreaming indicates that the minimum number of DBNodes are not streaming.
	// This is used with a False status for the condition.
	ReasonNotStreaming occoreapi.ConditionReason = "DBNodesNotStreaming"
	// ReasonMinDBNodesStreaming indicates that at least the minimum number of
	// DBNodes (as specified in the spec) are replicating. It's used with a True
	// status for the condition.
	ReasonMinDBNodesStreaming occoreapi.ConditionReason = "MinDBNodesStreaming"
	// ReasonAllDBNodesStreaming indicates that the DBInstance has the required
	// number of replicating DBNodes. This is used with a True status for the
	// condition.
	ReasonAllDBNodesSreaming occoreapi.ConditionReason = "AllDBNodesStreaming"

	// ConditionNetworkReady shows whether a LoadBalancer has been created and
	// is ready to serve the DBNodes.
	ConditionNetworkReady occoreapi.ConditionType   = "NetworkReady"
	ReasonNetworkNotReady occoreapi.ConditionReason = "NetworkNotReady"
	ReasonNetworkReady    occoreapi.ConditionReason = "NetworkReady"

	// ConditionAvailable indicates whether the required number of DBNodes have
	// been successfully provisioned and are accessible via the LoadBalancer.
	//
	// This condition is True when both DBNodesReady and NetworkReady conditions
	// are True.
	//
	// Note, a True status does not necessarily mean that the DBNodes are
	// actively replicating, but rather the DNodes can serve (potentially stale)
	// data.
	ConditionAvailable     occoreapi.ConditionType   = "Available"
	ReasonDBNodesAvailable occoreapi.ConditionReason = "DBNodesAvailable"

	// ReasonInvalidSpec may be used on all Conditions of the DBIntance to indicate
	// that something in the Spec is invalid.
	ReasonInvalidSpec occoreapi.ConditionReason = "InvalidSpec"
)
