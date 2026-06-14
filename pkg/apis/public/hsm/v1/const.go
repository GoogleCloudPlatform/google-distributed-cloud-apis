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

type CredentialID string

const (

	// HSMSecretResourceType represents whether a secret is created by hsmcluster or hsmtenant.
	HSMSecretResourceType = "hsm.security.private.gdc.goog/secret-resource-type"
	// HSMSecretResourceNamespace represents the hsmcluster or hsmtenant object's namespace.
	HSMSecretResourceNamespace = "hsm.security.private.gdc.goog/secret-resource-namespace"
	// HSMSecretResourceName represents the hsmcluster or hsmtenant object's name.
	HSMSecretResourceName = "hsm.security.private.gdc.goog/secret-resource-name"
	// OperableComponent represents the operable component / team that a k8s object is created by.
	OperableComponent = "system.private.gdc.goog/component"
	// CredentialId represents the ID of a k8s secret.
	CredentialId = "system.private.gdc.goog/credentialid"

	// These IDs will be used to reference Rotation instructions in Toil Tasks.
	// See go/gdch-io-docs.
	CredIDProm     CredentialID = "HSM-0001" // Rotate Prometheus Auth Token
	CredIDBackup   CredentialID = "HSM-0002" // Rotate Backup Key and Backup Password
	CredIDSSH      CredentialID = "HSM-0003" // Rotate SSH Credentials
	CredIDKsadmin  CredentialID = "HSM-0004" // Rotate KSAdmin Password
	CredIDTenantPA CredentialID = "HSM-0005" // Rotate Tenant Platform Admin Passwords
	CredIDAdmin    CredentialID = "HSM-0006" // Rotate Admin Password
	CredIDLuna     CredentialID = "HSM-0007" // Rotate Luna HSM Passwords
	// HSM T0008 is CTM internal CA rotation.
	// HSM T0009 is CTM master key rotation.
	CredIDServer      CredentialID = "HSM-0010" // Rotate Server Passwords
	CredIDONTAP       CredentialID = "HSM-0011" // Rotate ONTAP Credentials
	CredIDStorageGRID CredentialID = "HSM-0012" // Rotate StorageGRID Credentials
	// HSM T0013 is back up of HSM Credentials.
	CredIDKMSCrypto        CredentialID = "HSM-0014" // Rotate KMS Crypto Credentials
	CredIDKMSAdmin         CredentialID = "HSM-0015" // Rotate KMS Admin Credentials
	CredIDKMIPRegistration CredentialID = "HSM-0016" // Rotate KMIP Credentials
)
