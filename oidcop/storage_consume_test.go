package oidcop

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

// small enough to stay within the redis timeout of the test environment.
const testConcurrency = 5

func TestStorage_AuthRequestByCode_SingleUse(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	authReq := schema.NewAuthRequest(&oidc.AuthRequest{ClientID: "test_client_id", CodeChallenge: "test_challenge", CodeChallengeMethod: oidc.CodeChallengeMethodS256}, "test_username")
	authReq.ID = "test_request_id"
	require.NoError(t, schema.SaveAuthRequest(authReq))
	require.NoError(t, s.SaveAuthCode(ctx, authReq.ID, "test_code"))

	var success atomic.Int32
	var wg sync.WaitGroup
	for range testConcurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := s.AuthRequestByCode(ctx, "test_code"); err == nil {
				success.Add(1)
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(1), success.Load())

	// the code stays unusable after the request is deleted, and DeleteAuthRequest does not need the code.
	require.NoError(t, s.DeleteAuthRequest(ctx, authReq.ID))
	_, err := s.AuthRequestByCode(ctx, "test_code")
	assert.Error(t, err)
}

func TestStorage_renewRefreshToken_SingleUse(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}

	refreshToken := &schema.OIDCRefreshToken{
		ID:            "test_refresh_token",
		Token:         "test_refresh_token",
		Subject:       "test_username",
		ApplicationID: "test_client_id",
		Expiration:    time.Now().Add(time.Hour),
		Scopes:        []string{oidc.ScopeOpenID},
	}
	require.NoError(t, schema.SetOIDCRefreshToken(refreshToken))

	var success atomic.Int32
	var wg sync.WaitGroup
	for i := range testConcurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newToken := "new_refresh_token_" + string(rune('a'+i))
			if err := s.renewRefreshToken("test_refresh_token", newToken, "new_access_token"); err == nil {
				success.Add(1)
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, int32(1), success.Load())

	_, err := schema.GetOIDCRefreshToken("test_refresh_token")
	assert.ErrorIs(t, err, schema.ErrNotFound)
}
