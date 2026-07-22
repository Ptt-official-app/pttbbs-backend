package schema

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client_c *db.Collection

type Client struct {
	// 可信任的 app-client

	ClientID string `bson:"client_id"`

	ClientName string `bson:"client_name"`

	TheApplicationType        op.ApplicationType  `bson:"apptype"`
	TheAccessTokenType        op.AccessTokenType  `bson:"token_type"`
	TheRedirectURIs           []string            `bson:"redirect_uris"`
	ThePostLogoutRedirectURIs []string            `bson:"logout_redirect_uris"`
	TheResponseTypes          []oidc.ResponseType `bson:"resps"`
	TheGrantTypes             []oidc.GrantType    `bson:"grants"`

	ClientType   types.ClientType `bson:"client_type"`
	RemoteAddr   string           `bson:"ip"`
	UpdateNanoTS types.NanoTS     `bson:"update_nano_ts"`
}

// AccessTokenType implements [op.Client].
func (c *Client) AccessTokenType() op.AccessTokenType {
	return c.TheAccessTokenType
}

// ApplicationType implements [op.Client].
func (c *Client) ApplicationType() op.ApplicationType {
	return c.TheApplicationType
}

// AuthMethod implements [op.Client].
// AuthMethod must return the authentication method (client_secret_basic, client_secret_post, none, private_key_jwt)
func (c *Client) AuthMethod() oidc.AuthMethod {
	return oidc.AuthMethodNone
}

// ClockSkew implements [op.Client].
func (c *Client) ClockSkew() time.Duration {
	return 0
}

// DevMode implements [op.Client].
func (c *Client) DevMode() bool {
	return false
}

// GetID implements [op.Client].
func (c *Client) GetID() string {
	return c.ClientID
}

// GrantTypes implements [op.Client].
func (c *Client) GrantTypes() []oidc.GrantType {
	return c.TheGrantTypes
}

// IDTokenLifetime implements [op.Client].
func (c *Client) IDTokenLifetime() time.Duration {
	return types.ACCESS_TOKEN_EXPIRE_TS_DURATION
}

// IDTokenUserinfoClaimsAssertion implements [op.Client].
func (c *Client) IDTokenUserinfoClaimsAssertion() bool {
	return true
}

// IsScopeAllowed implements [op.Client].
func (c *Client) IsScopeAllowed(scope string) bool {
	return true
}

// LoginURL implements [op.Client].
// LoginURL will be called to redirect the user (agent) to the login UI
func (c *Client) LoginURL(authRequestID string) string {
	return types.FRONTEND_LOGIN_URL + "?authRequestID=" + authRequestID
}

// PostLogoutRedirectURIs implements [op.Client].
func (c *Client) PostLogoutRedirectURIs() []string {
	return c.ThePostLogoutRedirectURIs
}

// RedirectURIs implements [op.Client].
func (c *Client) RedirectURIs() []string {
	return c.TheRedirectURIs
}

// ResponseTypes implements [op.Client].
func (c *Client) ResponseTypes() []oidc.ResponseType {
	return c.TheResponseTypes
}

// RestrictAdditionalAccessTokenScopes implements [op.Client].
func (c *Client) RestrictAdditionalAccessTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

// RestrictAdditionalIdTokenScopes implements [op.Client].
func (c *Client) RestrictAdditionalIdTokenScopes() func(scopes []string) []string {
	return func(scopes []string) []string {
		return scopes
	}
}

var EMPTY_CLIENT = &Client{}

var (
	CLIENT_CLIENT_ID_b      = getBSONName(EMPTY_CLIENT, "ClientID")
	CLIENT_CLIENT_SECRET_b  = getBSONName(EMPTY_CLIENT, "ClientSecret")
	CLIENT_REMOTE_ADDR_b    = getBSONName(EMPTY_CLIENT, "RemoteAddr")
	CLIENT_UPDATE_NANO_TS_b = getBSONName(EMPTY_CLIENT, "UpdateNanoTS")
)

func NewClient(clientID string, clientType types.ClientType, redirectURIs []string, remoteAddr string) *Client {
	nowNanoTS := types.NowNanoTS()

	applicationType := op.ApplicationTypeNative
	if clientType == types.CLIENT_TYPE_WEB {
		applicationType = op.ApplicationTypeUserAgent
	}

	return &Client{
		ClientID: clientID,

		TheApplicationType: applicationType,
		TheAccessTokenType: op.AccessTokenTypeJWT,
		TheRedirectURIs:    redirectURIs,

		TheResponseTypes: []oidc.ResponseType{oidc.ResponseTypeCode, oidc.ResponseTypeIDToken, oidc.ResponseTypeIDTokenOnly},

		TheGrantTypes: []oidc.GrantType{oidc.GrantTypeBearer, oidc.GrantTypeCode, oidc.GrantTypeRefreshToken},

		ClientType:   clientType,
		RemoteAddr:   remoteAddr,
		UpdateNanoTS: nowNanoTS,
	}
}

func UpdateClient(c *Client) (err error) {
	query := bson.M{
		CLIENT_CLIENT_ID_b: c.ClientID,
	}

	r, err := Client_c.CreateOnly(query, c)
	if err != nil {
		return err
	}
	if r.UpsertedCount > 0 {
		return nil
	}

	query[CLIENT_UPDATE_NANO_TS_b] = bson.M{
		"$lt": c.UpdateNanoTS,
	}
	r, err = Client_c.UpdateOneOnly(query, c)
	if err != nil {
		return err
	}
	if r.MatchedCount == 0 {
		return ErrNoMatch
	}

	return nil
}

func GetClient(clientID string) (ret *Client, err error) {
	query := bson.M{
		CLIENT_CLIENT_ID_b: clientID,
	}

	ret = &Client{}
	err = Client_c.FindOne(query, ret, nil)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return ret, nil
}
