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
	// Constants for labels.

	// IsDefaultCertificateIssuerLabel represents a CertificateIssuer label that
	// marks an issuer as the default CertificateIssuer.
	IsDefaultCertificateIssuerLabel      = Group + "/is-default-issuer"
	IsDefaultCertificateIssuerLabelValue = "true"

	// Constants for annotations
	// Certificate manual reissuance key and values
	CertManualReissuanceKey             = Group + "/manual-reissuance"
	CertManualReissuanceValueRequested  = "requested"
	CertManualReissuanceValueInProgress = "in-progress"
	CertManualReissuanceValueFinished   = "finished"

	// Constants for status
	ConditionReady = "Ready"

	// Constants for reasons
	ReasonReconciling      = "Reconciling"
	ReasonReconcileBackoff = "ReconcileBackoff"

	// Certificate condition reasons
	CertReasonIssuerNotExist                 = "IssuerNotExist"
	CertReasonIssuerNotReady                 = "IssuerNotReady"
	CertReasonInvalidDNSName                 = "InvalidDNSName"
	CertReasonPending                        = "Pending"
	CertReasonUsingFallbackCA                = "UsingFallbackCA"
	CertReasonUsingMatchedCert               = "UsingMatchedCert"
	CertReasonIssued                         = "Issued"
	CertificateRequestReasonTemplateConflict = "TemplateConflict"
	CertReasonDefaultIssuerRemoved           = "DefaultIssuerRemoved"

	// CertificateIssuer condition reasons
	IssuerReasonFallbackCANotReady    = "FallbackCANotReady"
	IssuerReasonFallbackCAReady       = "FallbackCAReady"
	IssuerReasonAllCSRsReady          = "AllCSRsReady"
	IssuerReasonNotAllCSRsReady       = "NotAllCSRsReady"
	IssuerReasonCAaaSReady            = "CAaaSReady"
	IssuerReasonCAaaSNotReady         = "CAaaSNotReady"
	IssuerReasonACMEClientReady       = "ACMEClientReady"
	IssuerReasonACMEClientNotReady    = "ACMEClientNotReady"
	IsserReasonACMEClientDeleteFailed = "ACMEClientDeleteFailed"

	// CSRStatus condition reasons
	CSRReasonSigned            = "Signed"
	CSRReasonWaitingForSigning = "WaitingForSigning"

	//SignedBYOCert condition reason
	SignedBYOCertReasonUnknown  = "Unknown"
	SignedBYOCertReasonAccepted = "Accepted"
	SignedBYOCertReasonRejected = "Rejected"

	// Label for CertificateRequest reference in RevokeCertificateRequest
	CertificateRequestRefLabel = Group + "/certificate-request-ref"
)

// Condition types and reasons for CertificateAuthority
const (
	// CertificateAuthorityConditionReasonReady is the reason why a certificate authority is ready for use.
	CertificateAuthorityConditionReasonReady = "Ready"
	// CertificateAuthorityConditionReasonInternalResourceNotReady explains why a
	// certificate authority is not ready for use.
	CertificateAuthorityConditionReasonInternalResourceNotReady = "InternalResourceNotReady"
	// CertificateAuthorityConditionReasonIssuing explains why a certificate authority is not ready for use.
	CertificateAuthorityConditionReasonIssuing = "IssuingCACertificate"
	// CertificateAuthorityConditionReasonExpired explains why a certificate authority is not ready for use.
	CertificateAuthorityConditionReasonExpired = "CertificateExpired"
	// CertificateAuthorityConditionReasonPathLengthConstraintViolated explains that the CA violates an RFC 5280 path length constraint.
	CertificateAuthorityConditionReasonPathLengthConstraintViolated = "PathLengthConstraintViolated"
	CertificateAuthorityConditionTypeCACertificateReady             = "CACertificateReady"

	// CertificateAuthorityConditionTypeIssuing indicates that issuing a certificate for the
	// certificate authority is not complete.
	CertificateAuthorityConditionTypeIssuing = "Issuing"
	// CertificateAuthorityConditionReasonGeneratingCSR indicates that the controller is
	// generating a CSR for the CA certificate issuance.
	CertificateAuthorityConditionReasonGeneratingCSR = "GeneratingCSR"
	// CertificateAuthorityConditionReasonWaitingForCSRSigning indicates that CA certificate
	// issuance is pending on CSR signing.
	CertificateAuthorityConditionReasonWaitingForCSRSigning = "WaitingForCSRSigning"
	// CertificateAuthorityConditionReasonCertificateIssued indicates that CA certificate is
	// issued and the certificate authority is ready for use.
	CertificateAuthorityConditionReasonCertificateIssued = "Issued"
	// CertificateAuthorityConditionReasonSIsoENotReady indicates that service isolation environment is not ready.
	CertificateAuthorityConditionReasonSIsoENotReady = "ServiceIsolationEnvironmentNotReady"
	// CertificateAuthorityConditionTypeDbsReady indicates that dbs resources are ready.
	CertificateAuthorityConditionTypeDbsReady = "DbsResourcesReady"
	// CertificateAuthorityConditionTypeSmallstepCAReady indicates that smallstep CA is ready.
	CertificateAuthorityConditionTypeSmallstepCAReady = "SmallstepCAReady"
	// CertificateAuthorityConditionReasonPending indicates that resources is still reconciling.
	CertificateAuthorityConditionReasonPending = "Pending"
	// CertificateAuthorityConditionTypeBackendReady indicates that smallstep backend is ready.
	CertificateAuthorityConditionTypeBackendReady = "CABackendReady"
	// CertificateAuthorityConditionTypeDNSRegistrationReady indicates that dns registration is ready.
	CertificateAuthorityConditionTypeDNSRegistrationReady = "DNSRegistrationReady"
)

const (
	PKISystemNamespace = "pki-system"
	PKISuffix          = "-pki"

	// These const are reserved suffix for cert-manager resources created corresponding to CertificateAuthority.
	SelfSignedIssuerSuffix = PKISuffix + "-selfsigned"
	CertificateSuffix      = PKISuffix + "-cert"
	CAIssuerSuffix         = PKISuffix + "-issuer"

	// These const are reserved suffix for cert-manager resources created corresponding to Certificate.
	CertManagerCertSuffix         = PKISuffix + "-ct"
	CertManagerFallbackCertSuffix = PKISuffix + "-fallback-ct"

	SIsoENameSuffix = "-pki-cas"

	// Constants for annotation

	// Define an annotation to signify to use cert-manager as backend for CA.
	// If the annotation is missing or set as false, smallstep as backend will be used to host CA.

	// We'll continue using cert-manager as the backend for infra web PKI. To ensure backward
	// compatibility, if the annotation is missing and the CA is deployed in the pki-system
	// namespace., we will stil use cert-manager as backend.
	CAUseCertManagerBackendAnnotationKey = Group + "/use-cert-manager-backend"
)
