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
	"gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/enterprise/common"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	cecoreapi "gke-internal.googlesource.com/private-cloud/pkg/apis/public/dbs/omnicore/core/v1"
)

//+kubebuilder:object:generate=true

type ImportSpec struct {
	// DBClusterRef is the dbcluster name within the same namespace to import into.
	// +required
	DBClusterRef               common.DBClusterRef `json:"dbclusterRef,omitempty"`
	cecoreapi.CommonImportSpec `json:",inline"`
}

//+kubebuilder:object:generate=true

type ImportStatus struct {
	cecoreapi.ImportStatus `json:",inline"`
}

// Import is a L1 import interface
type Import interface {
	cecoreapi.Entity
	ImportSpec() ImportSpec
	ImportStatus() *ImportStatus
	DBEngineName() string
}

type ImportList interface {
	ctrlclient.ObjectList
	ImportListItem() []Import
}
