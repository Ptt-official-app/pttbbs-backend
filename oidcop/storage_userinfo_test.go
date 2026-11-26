package oidcop

import (
	"context"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func TestStorage_UserinfoAndIntrospection(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	// tokens are issued with the username as subject.
	require.NoError(t, schema.CreateUsername("testUserID", "test_username", 123456890000000000))
	require.NoError(t, schema.CreateUserEmail("testUserID", "test@localhost.dev", false, 123456890000000000))

	const clientID = "test_client_id"
	authReq := schema.NewAuthRequestByUsername("test_username", clientID)
	assert.Equal(t, []string{clientID}, authReq.GetAudience())

	token, err := s.accessToken(clientID, "", authReq.GetSubject(), authReq.GetAudience(), []string{oidc.ScopeOpenID, oidc.ScopeEmail})
	require.NoError(t, err)

	t.Run("userinfo", func(t *testing.T) {
		userInfo := &oidc.UserInfo{}
		err := s.SetUserinfoFromToken(ctx, userInfo, token.ID, token.Subject, "")
		require.NoError(t, err)
		assert.Equal(t, "test_username", userInfo.Subject)
		assert.Equal(t, "test@localhost.dev", userInfo.Email)
	})

	t.Run("introspection", func(t *testing.T) {
		introspection := &oidc.IntrospectionResponse{}
		err := s.SetIntrospectionFromToken(ctx, introspection, token.ID, token.Subject, clientID)
		require.NoError(t, err)
		assert.Equal(t, "test_username", introspection.Subject)
		assert.Equal(t, clientID, introspection.ClientID)
		assert.Equal(t, []string{oidc.ScopeOpenID, oidc.ScopeEmail}, []string(introspection.Scope))
		assert.True(t, token.Expiration.After(time.Now()))
	})

	t.Run("introspection by other client", func(t *testing.T) {
		introspection := &oidc.IntrospectionResponse{}
		err := s.SetIntrospectionFromToken(ctx, introspection, token.ID, token.Subject, "other_client_id")
		assert.Error(t, err)
	})

	t.Run("unknown user", func(t *testing.T) {
		unknown, err := s.accessToken(clientID, "", "no_such_username", []string{clientID}, []string{oidc.ScopeOpenID})
		require.NoError(t, err)

		err = s.SetUserinfoFromToken(ctx, &oidc.UserInfo{}, unknown.ID, unknown.Subject, "")
		assert.ErrorIs(t, err, schema.ErrNotFound)
	})
}
