# IdentityProviderConfig

The IdentityProviderConfig contains the parameters of an IdP, and it support both SamlConfig and OIDCConfig.

The IdentityProviderConfig CR will be synced to the ClientConfig, where a controller will reconcile and connect the IdP to the organization admin cluster.
