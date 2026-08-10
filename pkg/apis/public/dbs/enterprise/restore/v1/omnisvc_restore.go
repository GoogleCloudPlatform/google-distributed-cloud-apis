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

//+kubebuilder:object:generate=true

// OmniSvcRestoreSpec defines the desired state of Restore that are specific to OmniSvc based product.
type OmniSvcRestoreSpec struct {
	// The Backup to restore from.
	// This field is optional.
	// You must specify either Backup or PointInTime.
	// If you specify Backup, then you must leave the ClonedDBClusterConfig field unspecified.
	// If you specify PointInTime, then you must provide a new DBCluster name in the ClonedDBClusterConfig field.
	// Otherwise, the Restore request will be rejected.
	//
	// +kubebuilder:validation:optional
	Backup string `json:"backup,omitempty"`
}

type OmniSvcRestore interface {
	Restore
	OmniSvcRestoreSpec() *OmniSvcRestoreSpec
}

type OmniSvcRestoreList interface {
	RestoreList
	OmniSvcRestoreListItem() []OmniSvcRestore
}
