package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func TestSaveAuthRequest_withoutCodeChallenge(t *testing.T) {
	setupTest()
	defer teardownTest()

	// authorization request without PKCE: no code_challenge.
	authReq := NewAuthRequest(&oidc.AuthRequest{ClientID: "client"}, "user")
	authReq.ID = "auth-request-id"
	require.Nil(t, authReq.CodeChallenge)

	require.NotPanics(t, func() {
		require.NoError(t, SaveAuthRequest(authReq))
	})

	got, err := GetOIDCAuthRequestByRequestID(authReq.ID)
	require.NoError(t, err)
	assert.Equal(t, "client", got.ClientID)

	require.NotPanics(t, func() {
		require.NoError(t, DeleteOIDCAuthRequest(authReq.ID))
	})
}

func TestSaveAuthRequest_withCodeChallenge(t *testing.T) {
	setupTest()
	defer teardownTest()

	authReq := NewAuthRequest(&oidc.AuthRequest{ClientID: "client", CodeChallenge: "challenge", CodeChallengeMethod: oidc.CodeChallengeMethodS256}, "user")
	authReq.ID = "auth-request-id2"

	require.NoError(t, SaveAuthRequest(authReq))
	require.NoError(t, DeleteOIDCAuthRequest(authReq.ID))
}
