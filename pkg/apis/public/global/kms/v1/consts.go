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

const (


	// GlobalRootKeyNotReadyReason indicates that the GlobalRootKey is not ready.
	// The MZAEADKey depends on the readiness of the GlobalRootKey for its PrimaryKey field.
	GlobalRootKeyNotReadyReason = "GlobalRootKeyNotReadyReason"

	// EnsureZoneListError indicates that the MZAEADKey reconciler failed to
	// populate the zone list to the MZAEADKey CR.
	EnsureZoneListError = "EnsureZoneListError"

	// GenerateAndWrapAEADKeyError indicates that the MZAEADKey reconciler failed
	// to create the zonal AEADKeys
	// and wrapped key material.
	GenerateAndWrapAEADKeyError = "GenerateAndWrapAEADKeyError"

	// EnsureZonalAEADKeyCRCreatedError indicates that the MZAEADKey reconciler
	// failed to create the zonal AEADKey CR in the current zone.
	EnsureZonalAEADKeyCRCreatedError = "EnsureZonalAEADKeyCRCreatedError"

	// EnsureZonalAEADKeyReadyError indicates that the MZAEADKey
	// reconciler failed to fetch the wrapped AEADKey material from the MZAEADKey
	// CR, unwrap and rewrap it, and save it to the zonal AEADKey CR.
	EnsureZonalAEADKeyReadyError = "EnsureZonalAEADKeyReadyError"

	// DeleteMZEAEADKeyError indicates that the MZAEADKey reconciler failed to
	// delete an MZAEADKEY.
	DeleteMZAEADKeyError = "DeleteMZAEADKeyError"

	// ZonesNotReady indicates that one or more zones are not ready.
	ZonesNotReady = "ZonesNotReady"
)
