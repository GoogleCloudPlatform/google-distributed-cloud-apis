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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/googlecloudplatform/google-distributed-cloud-apis/pkg/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Represents a configuration for an identity provider that supports OIDC or SAML.
// +gdcloud:manifest:relevant=true,oc=iam,component=iam,entities="identity-provider-configs"
// +gdcloud:manifest:verbs=create;delete;describe;list;
// +gdcloud:manifest:rbac="create,delete,describe,list:organization-iam-admin,idp-federation-admin"
// +gdcloud:manifest:rbac="describe,list:idp-federation-viewer"
// +gdcloud:manifest:skipcodegen=true
// +genclient
type IdentityProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityProviderConfigSpec   `json:"spec,omitempty"`
	Status IdentityProviderConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Contains a list of `IdentityProviderConfig` resources.
type IdentityProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderConfig `json:"items"`
}

// Provides the specification, or desired state, of an `IdentityProviderConfig` resource.
// Either OIDCConfig or SAMLConfig has to be provided but not both.
type IdentityProviderConfigSpec struct {
	// OIDC specific configuration.
	// +optional
	OIDCConfig *OIDCProviderConfig `json:"oidc,omitempty"`

	// SAML specific configuration.
	// +optional
	SAMLConfig *SAMLProviderConfig `json:"saml,omitempty"`
}

// Provides the status of an `IdentityProviderConfig` resource.
type IdentityProviderConfigStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// The most recent errors with the observed times included.
	ErrorStatus *corev1alpha1.ErrorStatus `json:"errorStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&IdentityProviderConfig{}, &IdentityProviderConfigList{})
}

// OIDCProviderConfig contains parameters needed for OIDC Authentication flow.
type OIDCProviderConfig struct {
	// ClientID is an ID for OIDC client application.
	ClientID string `json:"clientID"`

	// ClientSecret is the shared secret between OIDC client application
	// and OIDC provider.
	// +optional
	ClientSecret string `json:"clientSecret,omitempty"`

	// CertificateAuthorityData contains a standard Base64 encoded, PEM formatted certificate authority
	// certificate for OIDC provider.
	// +optional
	CertificateAuthorityData string `json:"certificateAuthorityData,omitempty"`

	// URI for the OIDC provider. This URI is the prefix of "/.well-known/openid-configuration". For
	// example, if the IDP's well-known configuration endpoint is
	// "https://{oauth-provider-hostname}/.well-known/openid-configuration", then the Issuer URI is
	// "https://{oauth-provider-hostname}".
	// +kubebuilder:validation:Format=uri
	IssuerURI string `json:"issuerURI"`

	// KubectlRedirectURI is the URI to redirect users authenticating to an
	// OIDC provider with the kubectl plugin.
	// +kubebuilder:validation:Format=uri
	// +optional
	KubectlRedirectURI string `json:"kubectlRedirectURI,omitempty"`

	// CloudConsoleRedirectURI is the URI to redirect users going through the
	// OAuth flow using cloud console.
	// +kubebuilder:validation:Format=uri
	// +optional
	CloudConsoleRedirectURI string `json:"cloudConsoleRedirectURI,omitempty"`

	// Comma-separated list of identifiers used to specify what access privileges
	// are being requested in addition to "openid" scope.
	// +optional
	Scopes string `json:"scopes,omitempty"`

	// Comma-separated list of key-value pairs that will be query-encoded and sent
	// with the authentication endpoint request.
	// +optional
	ExtraParams string `json:"extraParams,omitempty"`

	// Flag that denotes if the access-token should be included in the request as
	// part of the bearer token by `gcloud anthos auth login` and `kubectl oidc
	// login`. Defaults to false.
	// +optional
	EnableAccessToken bool `json:"enableAccessToken,omitempty"`

	// Name of the claim in the OIDC ID Token that holds the username. If this is
	// missing from the ID Token, authentication will fail.
	UserClaim string `json:"userClaim"`

	// Prefix to prepend to the user name.
	// +optional
	UserPrefix string `json:"userPrefix,omitempty"`

	// Name of the claim in the OIDC ID Token that holds the user's groups information.
	// +optional
	GroupsClaim string `json:"groupsClaim,omitempty"`

	// Prefix to prepend to the group name.
	// +optional
	GroupPrefix string `json:"groupPrefix,omitempty"`

	// Flag to denote if HTTP reverse proxy is used for connecting to
	// authentication provider. This should be set to true when authentication provider
	// is not reachable by Google Cloud Console.
	// +optional
	DeployCloudConsoleProxy bool `json:"deployCloudConsoleProxy,omitempty"`

	// Optional Common Expression Language (CEL) for mapping user attributes from
	// the identity provider to the web application.
	// +optional
	AttributeMapping map[string]string `json:"attributeMapping,omitempty"`

	// Optional ClientAuth params define the way AIS authenticates itself while sending the Token API.
	// Currently used during authorization code flow.
	// +optional
	ClientAuthParams *ClientAuthConfig `json:"clientAuthParams,omitempty"`

	// Optional configuration for the OIDC encrypted token feature.
	// +optional
	EncryptedTokens *EncryptedTokensConfig `json:"encryptedTokens,omitempty"`
}

// ClientAuthConfig determines the way AIS authenticates itself to the Identity Provider.
type ClientAuthConfig struct {
	// Type of the ClientAuthentication method. Default method is client_secret_post that sends
	// client_secret in the Token API.
	// Other method supported is private_key_jwt that sends a signed jwt in the Token API.
	// private_key_jwt method requires SigningKey configuration to be valid.
	// +optional
	Method string `json:"method,omitempty"`

	// Signature Algorithm used to compute the JWT.
	// Only supported algorithm at this point is RS256.
	// +optional
	SignatureAlg string `json:"signatureAlg,omitempty"`

	// SigningKey denotes the parameters required to retrieve the key used to sign the JWT.
	// Key is stored as a kubernetes secret.
	// +optional
	SigningKey ClientAuthSigningKeySecret `json:"signingKey,omitempty"`
}

// ClientAuthSigningKeySecret contains the details of the signing key.
type ClientAuthSigningKeySecret struct {
	// Name of the secret which stores the SigningKey.
	Name string `json:"name"`

	// Namespace of the secret which stores the SigningKey.
	Namespace string `json:"namespace"`

	// kid is the unique identifier that represents the signing key.
	// Identity Provider looks up public key corresponding to the kid in order to verify payload.
	// +optional
	Kid string `json:"kid"`
}

// EncryptedTokensConfig provides configuration for the OIDC token encryption feature.
type EncryptedTokensConfig struct {
	// Specifies if OIDC token(s) must be decrypted before being parsed. The token decryption feature
	// is only enabled when this field is set to true.
	Enabled bool `json:"enabled"`

	// Kubernetes secret where the token decryption key(s) is/are stored.
	DecryptionKeys []*DecryptionKeysSecret `json:"decryptionKeys"`
}

// DecryptionKeysSecret specifies the namespace, and name of a kubernetes secret that holds the
// decryption keys for decrypting the encrypted tokens.
// kid is the unique key identifier for a decryption key that's used by the Provider to let the
// token recipient know which decryption key needs to be used for decrypting the token(s).
type DecryptionKeysSecret struct {
	// Kubernetes Name of the secret.
	Name string `json:"name"`

	// Kubernetes NameSpace of the secret.
	NameSpace string `json:"namespace"`

	// kid is the unique key identifier that represents a decryption key.
	Kid string `json:"kid"`
}

// SAMLProviderConfig contains parameters needed for SAML Authentication flow.
type SAMLProviderConfig struct {
	// The SAML entity ID for the SAML provider, specified in a URI format.
	// For example: https://www.idp.com/saml.
	// +kubebuilder:validation:Format=uri
	IDPEntityID string `json:"idpEntityID"`

	// The URI to the SAML provider's SSO endpoint. For example: https://www.idp.com/saml/sso.
	// +kubebuilder:validation:Format=uri
	IDPSingleSignOnURI string `json:"idpSingleSignOnURI"`

	// The IDP certificates that will be used to verify the SAML response. These certificates should
	// be standard Base64 encoded, and PEM formatted. Only a maximum of 2 certificates are
	// supported to facilitate IDP certificate rotation.
	// +kubebuilder:validation:MaxItems=2
	IDPCertificateDataList []string `json:"idpCertificateDataList"`

	// Name of the attribute in the SAML response that holds the username. If this attribute is
	// missing from the SAML response, authentication will fail.
	// +optional
	UserAttribute string `json:"userAttribute,omitempty"`

	// Optional prefix to prepend to the user name.
	// +optional
	UserPrefix string `json:"userPrefix,omitempty"`

	// Name of the attribute in the SAML response that holds the user's groups.
	// +optional
	GroupsAttribute string `json:"groupsAttribute,omitempty"`

	// Optional prefix to prepend to each group name.
	// +optional
	GroupPrefix string `json:"groupPrefix,omitempty"`

	// Optional Common Expression Language (CEL) for mapping user attributes from
	// the identity provider to the web application.
	// +optional
	AttributeMapping map[string]string `json:"attributeMapping,omitempty"`

	// Optional configuration for the SAML encrypted assertion feature.
	// +optional
	EncryptedAssertions *EncryptedAssertionConfig `json:"encryptedAssertions,omitempty"`

	// Optional configuration for the SAML authentication request signing feature.
	// +optional
	SignedRequests *SignedRequestConfig `json:"signedRequests,omitempty"`
}

// EncryptedAssertionConfig provides configuration for the SAML assertion encryption feature.
type EncryptedAssertionConfig struct {
	// Specifies if SAML assertions must be decrypted before being parsed. The assertion decryption
	// feature is only enabled when this is set to true.
	Enabled bool `json:"enabled"`

	// Kubernetes secret where the Assertion encryption certificate, and the assertion decryption key
	// are stored. This secret needs to be of type TLS.
	// (see: https://kubernetes.io/docs/concepts/configuration/secret/#tls-secrets).
	DecryptionKeys []*KubernetesSecretConfig `json:"decryptionKeys"`
}

// SignedRequestConfig provides configuration for the SAML authentication request signing feature.
type SignedRequestConfig struct {
	// Specifies if SAML authentication request must be signed. The signed request feature is only
	// enabled when this is set to true.
	Enabled bool `json:"enabled"`

	// Kubernetes secret where the SAML request signing certificate key pair is stored. This secret
	// needs to be of type TLS. (see: https://kubernetes.io/docs/concepts/configuration/secret/#tls-secrets).
	SigningKey *KubernetesSecretConfig `json:"signingKey"`
}

// KubernetesSecretConfig specifies the namespace, and name of a kubernetes secret.
type KubernetesSecretConfig struct {
	// Kubernetes Name of the secret.
	Name string `json:"name"`

	// Kubernetes NameSpace of the secret.
	NameSpace string `json:"namespace"`
}
