package oidcop

import (
	"context"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

const (
	testClientID      = "test_client_id"
	testOtherClientID = "other_client_id"
	testUsername      = "test_username"
)

// testIssueTokens issues tokens through the code flow (scope offline_access).
func testIssueTokens(t *testing.T, s *Storage, clientID string) (accessTokenID string, refreshTokenID string) {
	t.Helper()

	authReq := schema.NewAuthRequestByUsername(testUsername, clientID)
	accessTokenID, refreshTokenID, _, err := s.CreateAccessAndRefreshTokens(context.Background(), authReq, "")
	require.NoError(t, err)
	require.NotEmpty(t, accessTokenID)
	require.NotEmpty(t, refreshTokenID)

	return accessTokenID, refreshTokenID
}

func testAccessTokenExists(t *testing.T, accessTokenID string) bool {
	t.Helper()

	_, err := schema.GetOIDCAccessToken(accessTokenID)
	if err != nil {
		require.ErrorIs(t, err, schema.ErrNotFound)
		return false
	}
	return true
}

func testRefreshTokenExists(t *testing.T, refreshTokenID string) bool {
	t.Helper()

	_, err := schema.GetOIDCRefreshToken(refreshTokenID)
	if err != nil {
		require.ErrorIs(t, err, schema.ErrNotFound)
		return false
	}
	return true
}

func TestStorage_CreateAccessAndRefreshTokens(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	accessTokenID, refreshTokenID := testIssueTokens(t, s, testClientID)

	accessToken, err := schema.GetOIDCAccessToken(accessTokenID)
	require.NoError(t, err)
	assert.Equal(t, testClientID, accessToken.ApplicationID)
	assert.Equal(t, testUsername, accessToken.Subject)
	assert.Equal(t, refreshTokenID, accessToken.RefreshTokenID)

	refreshToken, err := schema.GetOIDCRefreshToken(refreshTokenID)
	require.NoError(t, err)
	assert.Equal(t, testClientID, refreshToken.ApplicationID)
	assert.Equal(t, testUsername, refreshToken.Subject)
	assert.Equal(t, accessTokenID, refreshToken.AccessToken)

	// refresh: the refresh token and its access token are rotated.
	req, err := s.TokenRequestByRefreshToken(ctx, refreshTokenID)
	require.NoError(t, err)
	newAccessTokenID, newRefreshTokenID, _, err := s.CreateAccessAndRefreshTokens(ctx, req, refreshTokenID)
	require.NoError(t, err)
	assert.NotEqual(t, refreshTokenID, newRefreshTokenID)

	assert.False(t, testAccessTokenExists(t, accessTokenID))
	assert.False(t, testRefreshTokenExists(t, refreshTokenID))

	newAccessToken, err := schema.GetOIDCAccessToken(newAccessTokenID)
	require.NoError(t, err)
	assert.Equal(t, testUsername, newAccessToken.Subject)
	assert.Equal(t, newRefreshTokenID, newAccessToken.RefreshTokenID)

	newRefreshToken, err := schema.GetOIDCRefreshToken(newRefreshTokenID)
	require.NoError(t, err)
	assert.Equal(t, newAccessTokenID, newRefreshToken.AccessToken)
	assert.Equal(t, testClientID, newRefreshToken.ApplicationID)

	// the old refresh token can not be used again, and no access token is left behind.
	_, _, _, err = s.CreateAccessAndRefreshTokens(ctx, req, refreshTokenID)
	require.Error(t, err)

	tokens, err := schema.GetOIDCAccessTokensBySubjectAndClientID(testUsername, testClientID)
	require.NoError(t, err)
	require.Len(t, tokens, 1)
	assert.Equal(t, newAccessTokenID, tokens[0].ID)
}

func TestStorage_CreateAccessAndRefreshTokens_expiredRefreshToken(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	accessTokenID, refreshTokenID := testIssueTokens(t, s, testClientID)

	refreshToken, err := schema.GetOIDCRefreshToken(refreshTokenID)
	require.NoError(t, err)
	refreshToken.Expiration = time.Now().Add(-time.Minute)
	require.NoError(t, schema.SetOIDCRefreshToken(refreshToken))

	_, _, _, err = s.CreateAccessAndRefreshTokens(ctx, schema.NewOIDCRefreshTokenRequest(refreshToken), refreshTokenID)
	assert.ErrorIs(t, err, ErrInvalidToken)

	assert.False(t, testAccessTokenExists(t, accessTokenID))
	assert.False(t, testRefreshTokenExists(t, refreshTokenID))

	tokens, err := schema.GetOIDCAccessTokensBySubjectAndClientID(testUsername, testClientID)
	require.NoError(t, err)
	assert.Empty(t, tokens)
}

func TestStorage_CreateAccessToken(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}

	// implicit flow: the token belongs to the client of the auth request.
	authReq := schema.NewAuthRequestByUsername(testUsername, testClientID)
	accessTokenID, _, err := s.CreateAccessToken(context.Background(), authReq)
	require.NoError(t, err)

	accessToken, err := schema.GetOIDCAccessToken(accessTokenID)
	require.NoError(t, err)
	assert.Equal(t, testClientID, accessToken.ApplicationID)
	assert.Equal(t, testUsername, accessToken.Subject)
	assert.Empty(t, accessToken.RefreshTokenID)
}

func TestStorage_GetRefreshTokenInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	_, refreshTokenID := testIssueTokens(t, s, testClientID)

	userID, tokenID, err := s.GetRefreshTokenInfo(ctx, testClientID, refreshTokenID)
	require.NoError(t, err)
	assert.Equal(t, testUsername, userID)
	assert.Equal(t, refreshTokenID, tokenID)

	_, _, err = s.GetRefreshTokenInfo(ctx, testOtherClientID, refreshTokenID)
	assert.ErrorIs(t, err, op.ErrInvalidRefreshToken)

	_, _, err = s.GetRefreshTokenInfo(ctx, testClientID, "no_such_token")
	assert.ErrorIs(t, err, schema.ErrNotFound)
}

func TestStorage_RevokeToken(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	t.Run("access token by other client", func(t *testing.T) {
		accessTokenID, _ := testIssueTokens(t, s, testClientID)

		oidcErr := s.RevokeToken(ctx, accessTokenID, testUsername, testOtherClientID)
		require.NotNil(t, oidcErr)
		assert.Equal(t, oidc.InvalidClient, oidcErr.ErrorType)
		assert.True(t, testAccessTokenExists(t, accessTokenID))
	})

	t.Run("access token", func(t *testing.T) {
		accessTokenID, refreshTokenID := testIssueTokens(t, s, testClientID)

		assert.Nil(t, s.RevokeToken(ctx, accessTokenID, testUsername, testClientID))
		assert.False(t, testAccessTokenExists(t, accessTokenID))
		assert.True(t, testRefreshTokenExists(t, refreshTokenID))
	})

	t.Run("refresh token by other client", func(t *testing.T) {
		accessTokenID, refreshTokenID := testIssueTokens(t, s, testClientID)

		oidcErr := s.RevokeToken(ctx, refreshTokenID, testUsername, testOtherClientID)
		require.NotNil(t, oidcErr)
		assert.Equal(t, oidc.InvalidClient, oidcErr.ErrorType)
		assert.True(t, testRefreshTokenExists(t, refreshTokenID))
		assert.True(t, testAccessTokenExists(t, accessTokenID))
	})

	t.Run("refresh token also revokes its access token", func(t *testing.T) {
		accessTokenID, refreshTokenID := testIssueTokens(t, s, testClientID)

		assert.Nil(t, s.RevokeToken(ctx, refreshTokenID, testUsername, testClientID))
		assert.False(t, testRefreshTokenExists(t, refreshTokenID))
		assert.False(t, testAccessTokenExists(t, accessTokenID))
	})

	t.Run("unknown token", func(t *testing.T) {
		assert.Nil(t, s.RevokeToken(ctx, "no_such_token", testUsername, testClientID))
	})
}

func TestStorage_TerminateSession(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}

	accessTokenID0, refreshTokenID0 := testIssueTokens(t, s, testClientID)
	accessTokenID1, refreshTokenID1 := testIssueTokens(t, s, testClientID)
	otherAccessTokenID, otherRefreshTokenID := testIssueTokens(t, s, testOtherClientID)

	require.NoError(t, s.TerminateSession(context.Background(), testUsername, testClientID))

	for _, id := range []string{accessTokenID0, accessTokenID1} {
		assert.False(t, testAccessTokenExists(t, id))
	}
	for _, id := range []string{refreshTokenID0, refreshTokenID1} {
		assert.False(t, testRefreshTokenExists(t, id))
	}

	// the session with the other client is kept.
	assert.True(t, testAccessTokenExists(t, otherAccessTokenID))
	assert.True(t, testRefreshTokenExists(t, otherRefreshTokenID))
}

func TestStorage_CreateAuthRequest(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	_, err := s.CreateAuthRequest(ctx, &oidc.AuthRequest{ClientID: testClientID, Prompt: oidc.SpaceDelimitedArray{oidc.PromptNone}}, "")
	assert.ErrorIs(t, err, oidc.ErrLoginRequired())

	req, err := s.CreateAuthRequest(ctx, &oidc.AuthRequest{ClientID: testClientID, State: "test_state", Scopes: oidc.SpaceDelimitedArray{oidc.ScopeOpenID}}, "")
	require.NoError(t, err)
	require.NotEmpty(t, req.GetID())

	got, err := s.AuthRequestByID(ctx, req.GetID())
	require.NoError(t, err)
	assert.Equal(t, testClientID, got.GetClientID())
	assert.Equal(t, "test_state", got.GetState())
	assert.Equal(t, []string{testClientID}, got.GetAudience())
	assert.False(t, got.Done())
}
