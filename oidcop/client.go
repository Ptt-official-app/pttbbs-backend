package oidcop

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

type Client struct {
	ID                     string
	Secret                 string
	redirectURIs           []string
	postLogoutRedirectURIs []string
	applicationType        op.ApplicationType
	responseTypes          []oidc.ResponseType
	grantTypes             []oidc.GrantType
	accessTokenType        op.AccessTokenType
}

// AccessTokenType implements [op.Client].
func (c *Client) AccessTokenType() op.AccessTokenType {
	return c.accessTokenType
}

// ApplicationType must return the type of the client (app, native, user agent)
func (c *Client) ApplicationType() op.ApplicationType {
	return c.applicationType
}

// AuthMethod must return the authentication method (client_secret_basic, client_secret_post, none, private_key_jwt)
func (c *Client) AuthMethod() oidc.AuthMethod {
	return oidc.AuthMethodNone
}

// ClockSkew implements [op.Client].
func (c *Client) ClockSkew() time.Duration {
	return time.Duration(0)
}

// DevMode enables the use of non-compliant configs such as redirect_uris (e.g. http schema for user agent client)
func (c *Client) DevMode() bool {
	return types.SERVICE_MODE == types.DEV
}

// GetID must return the client_id
func (c *Client) GetID() string {
	return c.ID
}

// GrantTypes must return all allowed grant types (authorization_code, refresh_token, urn:ietf:params:oauth:grant-type:jwt-bearer)
func (c *Client) GrantTypes() []oidc.GrantType {
	return c.grantTypes
}

// IDTokenLifetime must return the lifetime of the client's id_tokens
func (c *Client) IDTokenLifetime() time.Duration {
	return types.ACCESS_TOKEN_EXPIRE_TS_DURATION
}

// IDTokenUserinfoClaimsAssertion allows specifying if claims of scope profile, email, phone and address are asserted into the id_token
// even if an access token if issued which violates the OIDC Core spec
// (5.4. Requesting Claims using Scope Values: https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims)
// some clients though require that e.g. email is always in the id_token when requested even if an access_token is issued
func (c *Client) IDTokenUserinfoClaimsAssertion() bool {
	return false
}

// IsScopeAllowed enables Client specific custom scopes validation
func (c *Client) IsScopeAllowed(scope string) bool {
	return true
}

// LoginURL will be called to redirect the user (agent) to the login UI
func (c *Client) LoginURL(authRequestID string) string {
	return types.FRONTEND_PREFIX + "/login?authRequestID=" + authRequestID
}

// PostLogoutRedirectURIs must return the registered post_logout_redirect_uris for sign-outs
func (c *Client) PostLogoutRedirectURIs() []string {
	return c.postLogoutRedirectURIs
}

// RedirectURIs must return the registered redirect_uris for Code and Implicit Flow
func (c *Client) RedirectURIs() []string {
	return c.redirectURIs
}

// ResponseTypes must return all allowed response types (code, id_token token, id_token)
// these must match with the allowed grant types
func (c *Client) ResponseTypes() []oidc.ResponseType {
	return c.responseTypes
}

// RestrictAdditionalAccessTokenScopes allows specifying which custom scopes shall be asserted into the JWT access_token
func (c *Client) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

// RestrictAdditionalIdTokenScopes allows specifying which custom scopes shall be asserted into the id_token
func (c *Client) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}
