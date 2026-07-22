package oidcop

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	jose "github.com/go-jose/go-jose/v4"
	"github.com/google/uuid"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

type Storage struct {
	// authRequests  map[string]*AuthRequest
	// codes map[string]string
	// tokens        map[string]*Token
	// clients       map[string]*Client
	// userStore     UserStore
	// services      map[string]Service
	// refreshTokens map[string]*RefreshToken
	signingKey *SigningKey // used for access-token jwt signing.
	// deviceCodes  map[string]deviceAuthorizationEntry
	// userCodes map[string]string
	// serviceUsers map[string]*Client
}

func NewStorage() (*Storage, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Storage{
		signingKey: NewSigningKey("key_id", privKey, pubKey),
	}, nil
}

// ClientCredentials implements [op.ClientCredentialsStorage].
func (s *Storage) ClientCredentials(ctx context.Context, clientID string, clientSecret string) (ret op.Client, err error) {
	client, err := schema.GetClient(clientID)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// ClientCredentialsTokenRequest implements [op.ClientCredentialsStorage].
func (s *Storage) ClientCredentialsTokenRequest(ctx context.Context, clientID string, scopes []string) (op.TokenRequest, error) {
	client, err := schema.GetClient(clientID)
	if err != nil {
		return nil, err
	}

	return &oidc.JWTTokenRequest{
		Subject:  client.ClientID,
		Audience: []string{clientID},
		Scopes:   scopes,
	}, nil
}

// AuthRequestByCode implements [op.Storage].
func (s *Storage) AuthRequestByCode(ctx context.Context, code string) (op.AuthRequest, error) {
	requestID, err := schema.GetRequestIDByOIDCCode(code)
	if err != nil {
		return nil, err
	}
	return s.AuthRequestByID(ctx, requestID)
}

// AuthRequestByID implements [op.Storage].
func (s *Storage) AuthRequestByID(ctx context.Context, requestID string) (op.AuthRequest, error) {
	request, err := schema.GetOIDCAuthRequestByRequestID(requestID)
	if err != nil {
		return nil, err
	}
	return request, nil
}

// AuthorizeClientIDSecret implements [op.Storage].
func (s *Storage) AuthorizeClientIDSecret(ctx context.Context, clientID string, clientSecret string) (err error) {
	_, err = schema.GetClient(clientID)
	if err != nil {
		return err
	}

	return nil
}

// CreateAccessAndRefreshTokens implements [op.Storage].
func (s *Storage) CreateAccessAndRefreshTokens(ctx context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshToken string, expiration time.Time, err error) {
	if teReq, ok := request.(op.TokenExchangeRequest); ok {
		return s.exchangeRefreshToken(ctx, teReq)
	}

	applicationID, authTime, amr := getInfoFromRequest(request)

	if currentRefreshToken == "" {
		refreshTokenID := uuid.NewString()
		accessToken, err := s.accessToken(applicationID, refreshTokenID, request.GetSubject(), request.GetAudience(), request.GetScopes())
		if err != nil {
			return "", "", time.Time{}, err
		}
		refreshToken, err := s.createRefreshToken(accessToken, amr, authTime)
		if err != nil {
			return "", "", time.Time{}, err
		}
		return accessToken.ID, refreshToken, accessToken.Expiration, nil
	}

	newRefreshToken = uuid.NewString()

	accessToken, err := s.accessToken(applicationID, newRefreshToken, request.GetSubject(), request.GetAudience(), request.GetScopes())
	if err != nil {
		return "", "", time.Time{}, err
	}

	if err := s.renewRefreshToken(currentRefreshToken, newRefreshToken, accessToken.ID); err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken.ID, newRefreshToken, accessToken.Expiration, nil
}

//nolint:unused
func (s *Storage) exchangeRefreshToken(ctx context.Context, request op.TokenExchangeRequest) (accessTokenID string, newRefreshToken string, expiration time.Time, err error) {
	applicationID := request.GetClientID()
	authTime := request.GetAuthTime()

	refreshTokenID := uuid.NewString()
	accessToken, err := s.accessToken(applicationID, refreshTokenID, request.GetSubject(), request.GetAudience(), request.GetScopes())
	if err != nil {
		return "", "", time.Time{}, err
	}

	refreshToken, err := s.createRefreshToken(accessToken, nil, authTime)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken.ID, refreshToken, accessToken.Expiration, nil
}

// createRefreshToken will store a refresh_token in-memory based on the provided information
func (s *Storage) createRefreshToken(accessToken *schema.OIDCAccessToken, amr []string, authTime time.Time) (token string, err error) {
	refreshToken := &schema.OIDCRefreshToken{
		ID:            accessToken.RefreshTokenID,
		Token:         accessToken.RefreshTokenID,
		AuthTime:      authTime,
		AMR:           amr,
		ApplicationID: accessToken.ApplicationID,
		Subject:       accessToken.Subject,
		Audience:      accessToken.Audience,
		Expiration:    time.Now().Add(5 * time.Hour),
		Scopes:        accessToken.Scopes,
		AccessToken:   accessToken.ID,
	}

	err = schema.SetOIDCRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	return refreshToken.Token, nil
}

func (s *Storage) accessToken(applicationID, refreshTokenID, subject string, audience, scopes []string) (accessToken *schema.OIDCAccessToken, err error) {
	accessToken = &schema.OIDCAccessToken{
		ID:             uuid.NewString(),
		ApplicationID:  applicationID,
		RefreshTokenID: refreshTokenID,
		Subject:        subject,
		Audience:       audience,
		Expiration:     time.Now().Add(types.ACCESS_TOKEN_EXPIRE_TS_DURATION),
		Scopes:         scopes,
	}

	err = schema.SetOIDCAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *Storage) renewRefreshToken(currentRefreshToken, newRefreshToken, newAccessToken string) (err error) {
	refreshToken, err := schema.GetOIDCRefreshToken(currentRefreshToken)
	if err != nil {
		return err
	}

	// deletes the refresh token
	err = schema.DeleteOIDCRefreshToken(currentRefreshToken)
	if err != nil {
		return err
	}

	// delete the access token which was issued based on this refresh token
	err = schema.DeleteOIDCAccessToken(refreshToken.AccessToken)
	if err != nil {
		return err
	}

	if refreshToken.Expiration.Before(time.Now()) {
		return ErrInvalidToken
	}

	// creates a new refresh token based on the current one
	refreshToken.Token = newRefreshToken
	refreshToken.ID = newRefreshToken
	refreshToken.Expiration = time.Now().Add(types.ACCESS_TOKEN_EXPIRE_TS_DURATION)
	refreshToken.AccessToken = newAccessToken

	err = schema.SetOIDCRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	return nil
}

// CreateAccessToken implements [op.Storage].
func (s *Storage) CreateAccessToken(ctx context.Context, request op.TokenRequest) (accessTokenID string, expiration time.Time, err error) {
	var clientID string
	switch req := request.(type) {
	case *schema.OIDCAuthRequest:
		// if authenticated for an app (auth code / implicit flow) we must save the client_id to the token
		clientID = req.ClientID
	case op.TokenExchangeRequest:
		clientID = req.GetClientID()
	}

	token, err := s.accessToken(clientID, "", request.GetSubject(), request.GetAudience(), request.GetScopes())
	if err != nil {
		return "", time.Time{}, err
	}
	return token.ID, token.Expiration, nil
}

// CreateAuthRequest implements [op.Storage].
func (s *Storage) CreateAuthRequest(ctx context.Context, authReq *oidc.AuthRequest, username string) (op.AuthRequest, error) {
	if len(authReq.Prompt) == 1 && authReq.Prompt[0] == "none" {
		// With prompt=none, there is no way for the user to log in
		// so return error right away.
		return nil, oidc.ErrLoginRequired()
	}

	request := schema.NewAuthRequest(authReq, username)
	request.ID = uuid.NewString()

	err := schema.SaveAuthRequest(request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

// DeleteAuthRequest implements [op.Storage].
func (s *Storage) DeleteAuthRequest(ctx context.Context, requestID string) (err error) {
	authReq, err := schema.GetOIDCAuthRequestByRequestID(requestID)
	if err != nil {
		return err
	}
	err = schema.DeleteOIDCAuthRequest(requestID)
	if err != nil {
		return err
	}

	return schema.DeleteOIDCCodeByCode(authReq.CodeChallenge.Challenge)
}

// GetClientByClientID implements [op.Storage].
func (s *Storage) GetClientByClientID(ctx context.Context, clientID string) (op.Client, error) {
	client, err := schema.GetClient(clientID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetKeyByIDAndClientID implements [op.Storage].
func (s *Storage) GetKeyByIDAndClientID(ctx context.Context, keyID string, clientID string) (*jose.JSONWebKey, error) {
	key_db, err := schema.GetOIDCKey(clientID, keyID)
	if err != nil {
		return nil, err
	}

	return &jose.JSONWebKey{
		KeyID: keyID,
		Use:   "sig",
		Key:   key_db.Key,
	}, nil
}

// GetPrivateClaimsFromScopes implements [op.Storage].
func (s *Storage) GetPrivateClaimsFromScopes(ctx context.Context, userID string, clientID string, scopes []string) (claims map[string]any, err error) {
	return nil, nil
}

// GetRefreshTokenInfo implements [op.Storage].
func (s *Storage) GetRefreshTokenInfo(ctx context.Context, clientID string, token string) (userID string, tokenID string, err error) {
	refreshToken, err := schema.GetOIDCRefreshToken(token)
	if err != nil {
		return "", "", err
	}
	if refreshToken.ApplicationID != clientID {
		return "", "", op.ErrInvalidRefreshToken
	}

	return refreshToken.Subject, refreshToken.ID, nil
}

// Health implements [op.Storage].
func (s *Storage) Health(context.Context) error {
	return nil
}

// KeySet implements [op.Storage].
func (s *Storage) KeySet(context.Context) ([]op.Key, error) {
	// as mentioned above, this example only has a single signing key without key rotation,
	// so it will directly use its public key
	//
	// when using key rotation you typically would store the public keys alongside the private keys in your database
	// and give both of them an expiration date, with the public key having a longer lifetime
	return []op.Key{&PublicKey{s.signingKey}}, nil
}

// RevokeToken implements [op.Storage].
func (s *Storage) RevokeToken(ctx context.Context, tokenOrTokenID string, userID string, clientID string) *oidc.Error {
	accessToken, err := schema.GetOIDCAccessToken(tokenOrTokenID)
	if err == nil { // revoke accessToken
		if accessToken.ApplicationID != clientID {
			return oidc.ErrInvalidClient().WithDescription("token was not issued for this client")
		}
		// if it is an access token, just remove it
		// you could also remove the corresponding refresh token if really necessary

		err = schema.DeleteOIDCAccessToken(tokenOrTokenID)
		if err != nil {
			return oidc.ErrServerError()
		}
		return nil
	}

	refreshToken, err := schema.GetOIDCRefreshToken(tokenOrTokenID)
	if err != nil {
		// if the token is neither an access nor a refresh token, just ignore it, the expected behavior of
		// being not valid (anymore) is achieved
		return nil
	}

	if refreshToken.ApplicationID != clientID {
		return oidc.ErrInvalidClient().WithDescription("token was not issued for this client")
	}

	err = schema.DeleteOIDCRefreshToken(tokenOrTokenID)
	if err != nil {
		return oidc.ErrServerError()
	}

	// if it is a refresh token, we will have to remove the access token as well
	err = schema.DeleteOIDCAccessToken(refreshToken.AccessToken)
	if err != nil {
		return oidc.ErrServerError()
	}

	return nil
}

// SaveAuthCode implements [op.Storage].
func (s *Storage) SaveAuthCode(ctx context.Context, id string, code string) error {
	return schema.SaveOIDCCode(id, code)
}

// SetIntrospectionFromToken implements [op.Storage].
func (s *Storage) SetIntrospectionFromToken(ctx context.Context, introspection *oidc.IntrospectionResponse, tokenID, subject, clientID string) (err error) {
	accessToken, err := schema.GetOIDCAccessToken(tokenID)
	if err != nil {
		return err
	}

	introspection.Expiration = oidc.FromTime(accessToken.Expiration)
	if accessToken.Expiration.Before(time.Now()) {
		return ErrInvalidToken
	}

	// check if the client is part of the requested audience
	for _, aud := range accessToken.Audience {
		if aud == clientID {
			// The introspection response only has to return a boolean (active) if the token is active.
			// This will automatically be done by the library if you don't return an error.
			// You can also return further information about the user / associated token.
			// e.g. the userinfo (equivalent to userinfo endpoint)

			userInfo := &oidc.UserInfo{}
			err := s.setUserinfo(ctx, userInfo, subject, clientID, accessToken.Scopes)
			if err != nil {
				return err
			}
			introspection.SetUserInfo(userInfo)
			//...and also the requested scopes...
			introspection.Scope = accessToken.Scopes
			//...and the client the token was issued to
			introspection.ClientID = accessToken.ApplicationID
			return nil
		}
	}
	return fmt.Errorf("token is not valid for this client")
}

func (s *Storage) setUserinfo(ctx context.Context, userInfo *oidc.UserInfo, userID, clientID string, scopes []string) (err error) {
	userEmail, err := schema.GetUserEmailByUserID(bbs.UUserID(userID))
	if err != nil {
		return err
	}
	if userEmail == nil {
		return schema.ErrNotFound
	}
	userProfile, err := schema.GetUserProfile(bbs.UUserID(userID))

	for _, scope := range scopes {
		switch scope {
		case oidc.ScopeOpenID:
			userInfo.Subject = string(userEmail.UserID)
		case oidc.ScopeEmail:
			userInfo.Email = userEmail.Email
			userInfo.EmailVerified = oidc.Bool(true)
		case oidc.ScopeProfile:
			if err != nil {
				return err
			}
			userInfo.PreferredUsername = userProfile.Username
			userInfo.Name = userProfile.GivenName + " " + userProfile.FamilyName
			userInfo.FamilyName = userProfile.FamilyName
			userInfo.GivenName = userProfile.GivenName
			userInfo.Locale = oidc.NewLocale(userProfile.Locale)
		}
	}
	return nil
}

// SetUserinfoFromScopes implements [op.Storage].
func (s *Storage) SetUserinfoFromScopes(ctx context.Context, userinfo *oidc.UserInfo, userID string, clientID string, scopes []string) error {
	return nil
}

// SetUserinfoFromToken implements [op.Storage].
func (s *Storage) SetUserinfoFromToken(ctx context.Context, userinfo *oidc.UserInfo, tokenID string, subject string, origin string) (err error) {
	token, err := schema.GetOIDCAccessToken(tokenID)
	if err != nil {
		return err
	}

	// the userinfo endpoint should support CORS. If it's not possible to specify a specific origin in the CORS handler,
	// and you have to specify a wildcard (*) origin, then you could also check here if the origin which called the userinfo endpoint here directly
	// note that the origin can be empty (if called by a web client)
	//
	// if origin != "" {
	//	client, ok := s.clients[token.ApplicationID]
	//	if !ok {
	//		return fmt.Errorf("client not found")
	//	}
	//	if err := checkAllowedOrigins(client.allowedOrigins, origin); err != nil {
	//		return err
	//	}
	//}
	if token.Expiration.Before(time.Now()) {
		return ErrInvalidToken
	}
	return s.setUserinfo(ctx, userinfo, token.Subject, token.ApplicationID, token.Scopes)
}

// SignatureAlgorithms implements [op.Storage].
func (s *Storage) SignatureAlgorithms(context.Context) ([]jose.SignatureAlgorithm, error) {
	return []jose.SignatureAlgorithm{s.signingKey.algorithm}, nil
}

// SigningKey implements [op.Storage].
func (s *Storage) SigningKey(context.Context) (op.SigningKey, error) {
	return s.signingKey, nil
}

// TerminateSession implements [op.Storage].
func (s *Storage) TerminateSession(ctx context.Context, username string, clientID string) error {
	tokens, err := schema.GetOIDCAccessTokensBySubjectAndClientID(username, clientID)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		err1 := schema.DeleteOIDCAccessToken(token.ID)
		if err1 != nil {
			err = err1
		}
		err2 := schema.DeleteOIDCRefreshToken(token.RefreshTokenID)
		if err2 != nil {
			err = err2
		}
	}

	return err
}

// TokenRequestByRefreshToken implements [op.Storage].
func (s *Storage) TokenRequestByRefreshToken(ctx context.Context, refreshTokenID string) (op.RefreshTokenRequest, error) {
	refreshToken, err := schema.GetOIDCRefreshToken(refreshTokenID)
	if err != nil {
		return nil, err
	}

	return schema.NewOIDCRefreshTokenRequest(refreshToken), nil
}

// ValidateJWTProfileScopes implements [op.Storage].
func (s *Storage) ValidateJWTProfileScopes(ctx context.Context, userID string, scopes []string) ([]string, error) {
	allowedScopes := make([]string, 0)
	for _, scope := range scopes {
		if scope == oidc.ScopeOpenID {
			allowedScopes = append(allowedScopes, scope)
		}
	}
	return allowedScopes, nil
}
