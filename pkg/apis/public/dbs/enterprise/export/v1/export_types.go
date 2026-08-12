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
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/enterprise/common"
	cecoreapi "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/public/dbs/omnicore/core/v1"
)

//+kubebuilder:object:generate=true

type ExportSpec struct {
	// DBClusterRef is the dbcluster name within the same namespace to export from.
	// +required
	DBClusterRef               common.DBClusterRef `json:"dbclusterRef,omitempty"`
	cecoreapi.CommonExportSpec `json:",inline"`
}

//+kubebuilder:object:generate=true

type ExportStatus struct {
	cecoreapi.ExportStatus `json:",inline"`
}

// Export is a L1 export interface
type Export interface {
	cecoreapi.Entity
	ExportSpec() ExportSpec
	ExportStatus() *ExportStatus
	DBEngineName() string
}

type ExportList interface {
	ctrlclient.ObjectList
	ExportListItem() []Export
}
