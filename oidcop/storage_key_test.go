package oidcop

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestStorage_GetKeyByIDAndClientID(t *testing.T) {
	setupTest()
	defer teardownTest()

	s := &Storage{}
	ctx := context.Background()

	pubKey, _, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)

	key := &schema.OIDCKey{ClientID: "client_id", KeyID: "key_id", Key: &pubKey}
	query := bson.M{schema.OIDC_KEY_CLIENT_ID_b: key.ClientID, schema.OIDC_KEY_KEY_ID_b: key.KeyID}
	_, err = schema.OIDCKey_c.CreateOnly(query, key)
	require.NoError(t, err)

	t.Run("unknown key", func(t *testing.T) {
		assert.NotPanics(t, func() {
			got, err := s.GetKeyByIDAndClientID(ctx, "unknown_key_id", "client_id")
			assert.ErrorIs(t, err, schema.ErrNotFound)
			assert.Nil(t, got)
		})
	})

	t.Run("known key", func(t *testing.T) {
		got, err := s.GetKeyByIDAndClientID(ctx, "key_id", "client_id")
		require.NoError(t, err)
		assert.Equal(t, "key_id", got.KeyID)
		assert.Equal(t, pubKey, got.Key)
		assert.True(t, got.Valid())
	})
}
