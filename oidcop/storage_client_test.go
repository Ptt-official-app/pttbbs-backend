package oidcop

import (
	"context"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_UnknownClient(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	known := schema.NewClient("known_client_id", types.CLIENT_TYPE_APP, nil, "localhost")
	require.NoError(t, schema.UpdateClient(known))

	t.Run("GetClientByClientID unknown", func(t *testing.T) {
		client, err := s.GetClientByClientID(ctx, "unknown_client_id")
		assert.ErrorIs(t, err, schema.ErrNotFound)
		assert.Nil(t, client)
	})

	t.Run("GetClientByClientID known", func(t *testing.T) {
		client, err := s.GetClientByClientID(ctx, "known_client_id")
		require.NoError(t, err)
		assert.Equal(t, "known_client_id", client.GetID())
	})

	t.Run("ClientCredentials unknown", func(t *testing.T) {
		client, err := s.ClientCredentials(ctx, "unknown_client_id", "secret")
		assert.ErrorIs(t, err, schema.ErrNotFound)
		assert.Nil(t, client)
	})

	t.Run("ClientCredentialsTokenRequest unknown", func(t *testing.T) {
		assert.NotPanics(t, func() {
			req, err := s.ClientCredentialsTokenRequest(ctx, "unknown_client_id", nil)
			assert.ErrorIs(t, err, schema.ErrNotFound)
			assert.Nil(t, req)
		})
	})

	t.Run("AuthorizeClientIDSecret unknown", func(t *testing.T) {
		err := s.AuthorizeClientIDSecret(ctx, "unknown_client_id", "secret")
		assert.ErrorIs(t, err, schema.ErrNotFound)
	})
}
