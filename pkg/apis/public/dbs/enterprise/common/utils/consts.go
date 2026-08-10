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

package common

const (
	FleetOperatorNsName = "dbs-fleet-system"
	// CertificatesSecretName is the name of the Secret that contains the CA
	// certificate information for all DBClusters in the same namespace.
	CertificatesSecretName = "dbs-certificates"
	AlloyDBOmniGroupName   = "alloydbomni.dbadmin.gdc.goog"
	SamwiseGroupName       = "alloydbomni.dbadmin.goog"
	// DbClusterNamePrefixLengthInInternalName is how many chars we use from the
	// dbcluster name as the prefix for creating the internal name
	DbClusterNamePrefixLengthInInternalName = 24

	// ManagedByLabel is a k8s recommended label used to mark when a resource is managed by a controller.
	ManagedByLabel    = "app.kubernetes.io/managed-by"
	ManagedByOperator = "operator"
)
