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
	corev1 "k8s.io/api/core/v1"
)

const (
	// RootKeyFinalizer is the finalizer added to the root key secret to prevent premature deletion.
	RootKeyFinalizer = "kms.gdc.goog/finalizer"
	// RootKeyDeleteRequestAnnotation is the annotation to request deletion of the root key secret.
	RootKeyDeleteRequestAnnotation = "kms.gdc.goog/delete-rootkey-request"
	// PausedAnnotation is the annotation to pause reconciling specific KMS resources.
	PausedAnnotation = "kms.gdc.goog/paused"

	// MZAnnotation is the annotation that indicates the resource should only be reconciled by global reconcilers.
	MZAnnotation = "kms.gdc.goog/mz"
	// MZGlobalNamespaceAnnotation is the annotation that indicates the related global namespace.
	MZGlobalNamespaceAnnotation = "kms.gdc.goog/mz-global-namespace"

	// InternalCryptoBackendTypeAnnotation is the annotation key used on a CryptoBackend
	// resource to instruct the reconciler to use a specific internal provider
	// logic path (e.g., "SoftHSM") instead of the default production type.
	// This is for internal testing and CI/CD ONLY and should not be used by end-users.
	InternalCryptoBackendTypeAnnotation = "kms.gdc.goog/internal-cryptobackend-type"

	// SoftHSMPartitionLabelAnnotation is the annotation key for specifying the token label
	// for the internal SoftHSM backend. The KMS service will use this to find the correct slot.
	SoftHSMPartitionLabelAnnotation = "softhsm.internal.kms.gdc.goog/partition-label"

	// SoftHSMEndpointAnnotation is the annotation key for specifying the direct connection
	// endpoint for the internal SoftHSM cryptobackend service.
	SoftHSMEndpointAnnotation = "softhsm.internal.kms.gdc.goog/endpoint"

	// SecretTypeKMSRootKeyLocal contains a list of versioned root key materials.
	// Required fields:
	// - Secret.Data["active-key-material"] - key material that will be used for new wrap operations.
	// - Secret.Data["active-version"] - ID of the active version.
	// Optional fields:
	// - Secret.Data["previous-key-material"] - key material that used to be the active key material at a point. It gets moved to previous-key-material during root key rotation.
	// - Secret.Data["previous-version"] - ID of the key version that used to be the active version at one point. It gets moved to previous-version during root key rotation.
	SecretTypeKMSRootKeyLocal corev1.SecretType = "kms.gdc.goog/local-root"

	// SecretTypeKMSRootKeyCTM contains a list of keyIDs to keys in the Thales Cipher Trust Manager.
	// Required fields:
	// - Secret.Data["addr"] - comma separated list of HSM addresses in an HSM cluster.
	// - Secret.Data["username"] - username to login to the HSM.
	// - Secret.Data["password"] - password to login to HSM.
	// - Secret.Data["key-id"] - unique identifier of the root key on the HSM, returned by the HSM.
	// - Secret.Data["active-version"] - version ID of the active version of the root key.
	// Optional fields:
	// - Secret.Data["domain"] - domain the key exists in. Optional, CTM will default to root domain.
	SecretTypeKMSRootKeyCTM corev1.SecretType = "kms.gdc.goog/ctm-root"

	// SecretTypeKMSRootKeyGlobal contains the public/private key pair used to import the global root key material,
	// as well as the global root key material itself after it gets imported.
	// Required fields:
	// - Secret.Data["publicKey"] - public key used for import.
	// - Secret.Data["wrappedPrivateKey"] - private key used for import, wrapped by zonal KMS root key.
	// - Secret.Data["rootKeyIDForPrivateKey"] - ID of the zonal root key that wrapped the private key.
	// Optional fields:
	// - Secret.Data["wrappedGlobalRootKey"] - global root key material, wrapped by zonal KMS root key.
	// - Secret.Data["rootKeyIDForGlobalRootKey"] - ID of the zonal root key that wrapped the global root key.
	SecretTypeKMSRootKeyGlobal corev1.SecretType = "kms.gdc.goog/global-root"

	// CTMHSMAddressesKey is the key of a comma separated list of HSM addresses in an HSM cluster.
	CTMHSMAddressesKey = "addr"
	// CTMUsernameKey is the key of the username to login to the HSM.
	CTMUsernameKey = "username"
	// CTMPasswordKey is the key of the password to login to the HSM.
	CTMPasswordKey = "password"
	// CTMDomainKey is the key of the domain the user belongs to. Optional, if not set, user is assumed to be part of the root domain.
	CTMDomainKey = "domain"
	// CTMRootCACertsKey is the key of the root CA certs to connect to the HSM.
	CTMRootCACertsKey = "rootCACerts"

	// CTMKeyIDKey is the key of `id` field of the root key on the CTM key resource.
	CTMKeyIDKey = "key-id"
	// CTMActiveVersionKey is the key of version ID of the active version of the root key.
	CTMActiveVersionKey = "active-version"

	// LocalActiveKeyMaterialKey is the key of the root key  material that is currently active.
	LocalActiveKeyMaterialKey = "active-key-material"
	// LocalActiveVersionKey is the version ID mapping to active key material.
	LocalActiveVersionKey = "active-version"
	// LocalPreviousKeyMaterialKey is the key of the root key  material that was active at one point. It gets moved to previous key material during root key rotation.
	LocalPreviousKeyMaterialKey = "previous-key-material"
	// LocalPreviousVersionKey is the version ID mapping to previous key material.
	LocalPreviousVersionKey = "previous-version"
	// GRKSecretPublicKey is the public key for importing the global root key material.
	GRKSecretPublicKey = "publicKey"
	// GRKSecretWrappedPrivateKey is the private key for importing the global root key material, wrapped by the zonal KMS root key.
	GRKSecretWrappedPrivateKey = "wrappedPrivateKey"
	// GRKSecretRootKeyIDForPrivateKey is the ID of the zonal KMS root key used to wrap the private key material.
	GRKSecretRootKeyIDForPrivateKey = "rootKeyIDForPrivateKey"
	// GRKSecretWrappedGlobalRootKey is the global root key material, wrapped by the zonal KMS root key.
	GRKSecretWrappedGlobalRootKey = "wrappedGlobalRootKey"
	// GRKSecretRootKeyIDForGlobalRootKey is the ID of the zonal KMS root key used to wrap the global root key material.
	GRKSecretRootKeyIDForGlobalRootKey = "rootKeyIDForGlobalRootKey"

	ReconcileSuccess = "ReconcileSuccess"

	KeyImportLabel     = "imported"
	KeyImportLabelTrue = "true"

	// Status field in the root key label.
	RootKeyStatusLabel = "status"
	// Ready condition.
	RootKeyStatusValueReady = "ready"
	// Deprecated condition.
	RootKeyStatusValueDeprecated = "deprecated"
	// Unusable condition.
	RootKeyStatusValueUnusable = "unusable"

	RootKeyUnavailableReason = "RootKeyUnavailable"
	// Passthrough key creation is not enabled in configuration.
	PassthroughKeysDisabledReason = "PassthroughKeysDisabled"
	// Passthrough key creation is enabled but not implemented by this version.
	PassthroughKeysNotImplementedReason = "PassthroughKeysNotImplemented"
	// Passthrough key creation failed on the HSM.
	PassthroughKeyCreationFailedReason = "PassthroughKeyCreationFailed"
	// Passthrough key creation succeeded on the HSM.
	PassthroughKeyCreatedReason = "PassthroughKeyCreated"
	// ROtation job failed and retried for over the retryLimitTime.
	RotationJobExceededRetryLimitTimeReason = "JobExceededRetryLimitTime"

	InvalidRootKeyReason = "InvalidRootKey"
	KeyDeprecatedReason  = "KeyDeprecated"

	// RotationJob is the finalizer added to the rotation job to prevent deletion until completion.
	RotationJobFinalizer = "rotationjob.kms.gdc.goog/finalizer"
	// RotationJob has finished rewrapping existing resources and is complete.
	RotationJobStatusCompleted        = "Completed"
	RotationJobRewrapInProgressReason = "RewrapKMSResourcesInProgress"
	RotationJobRewrapFailedReason     = "RewrapKMSResourcesFailed"
	RotationJobStatusFailed           = "Failed"

	RotationJobDeletionFailed  = "DeletionFailed"
	RotationJobIncompleteError = "JobIncomplete"

	KeyImportCompleted      = "KeyImportCompleted"
	KeyImportReconcileError = "ReconcileKeyToImportError"
	KeyImportAwaitingReason = "AwaitingKeyToImport"
	KeyImportSuccess        = "ImportSuccess"

	KeyExportCompleted      = "KeyExportCompleted"
	KeyExportReconcileError = "ReconcileKeyToExportError"
)
