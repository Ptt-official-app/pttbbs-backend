package oidcop

import (
	"context"
	"crypto/sha256"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/zitadel/oidc/v3/pkg/op"
	"golang.org/x/text/language"
)

func InitGinRouter(router *gin.Engine, extraOptions ...op.Option) (err error) {
	provider, err := NewProvider(extraOptions...)
	if err != nil {
		return err
	}

	router.GET("/.well-known/openid-configuration", gin.WrapH(provider))
	router.GET("/api/oidc/*path", gin.WrapH(provider))
	router.POST("/api/oidc/*path", gin.WrapH(provider))

	return nil
}

func NewProvider(extraOptions ...op.Option) (provider *op.Provider, err error) {
	config := &op.Config{
		CryptoKey:   sha256.Sum256([]byte(types.OIDC_OP_KEY)), // used for code encrypt.
		CryptoKeyId: types.OIDC_OP_KEY_ID,

		// will be used if the end_session endpoint is called without a post_logout_redirect_uri
		DefaultLogoutRedirectURI: types.OIDC_OP_POST_LOGOUT_URL,

		// enables code_challenge_method S256 for PKCE (and therefore PKCE in general)
		CodeMethodS256: true,

		// enables additional client_id/client_secret authentication by form post (not only HTTP Basic Auth)
		AuthMethodPost: true,

		// enables additional authentication by using private_key_jwt
		AuthMethodPrivateKeyJWT: true,

		// enables refresh_token grant use
		GrantTypeRefreshToken: true,

		// enables use of the `request` Object parameter
		RequestObjectSupported: true,

		// this example has only static texts (in English), so we'll set the here accordingly
		SupportedUILocales: []language.Tag{language.English},

		SupportedScopes: []string{
			"openid",
			// "profile",
			// "email",
			"offline_access",
		},

		SupportedClaims: []string{
			"sub",
			"aud",
			"exp",
			"iat",
			"iss",
			"auth_time",
			"nonce",
			"acr",
			"amr",
			"scopes",
			"client_id",
			"azp",
			// "name",
			// "family_name",
			// "given_name",
			// "locale",
			// "email",
		},

		DeviceAuthorization: op.DeviceAuthorizationConfig{
			Lifetime:     5 * time.Minute,
			PollInterval: 5 * time.Second,
			UserFormPath: "/device",
			UserCode:     op.UserCodeBase20,
		},
	}

	options := append([]op.Option{
		op.WithCustomEndpoints(
			op.NewEndpoint("/api/oidc/auth"),
			op.NewEndpoint("/api/oidc/token"),
			op.NewEndpoint("/api/oidc/user"),
			op.NewEndpoint("/api/oidc/revoke"),
			op.NewEndpoint("/api/oidc/end_session"),
			op.NewEndpoint("/api/oidc/keys"),
		),
		op.WithAccessTokenVerifierOpts(func(accessTokenVerifier *op.AccessTokenVerifier) {
			accessTokenVerifier.SupportedSignAlgs = []string{string(jose.EdDSA)}
		}),
	}, extraOptions...)

	if types.OIDC_OP_IS_ALLOW_HTTP {
		options = append(options, op.WithAllowInsecure())
	}

	storage, err := NewStorage()
	if err != nil {
		return nil, err
	}

	provider, err = op.NewProvider(config, storage, op.StaticIssuer(types.OIDC_OP_ISSUER),
		options...,
	)
	if err != nil {
		return nil, err
	}

	setupProvider(provider)

	op.NewIssuerInterceptor(provider.IssuerFromRequest)

	return provider, nil
}

func setupProvider(provider *op.Provider) {
	PROVIDER = provider
	AUTH_CALLBACK_URL = func(ctx context.Context, oidcID string) string {
		return types.URL_PREFIX + op.AuthCallbackURL(PROVIDER)(ctx, oidcID)
	}
}
