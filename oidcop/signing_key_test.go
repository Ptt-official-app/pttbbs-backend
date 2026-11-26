package oidcop

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStorage_SigningKey(t *testing.T) {
	origFile := types.OIDC_OP_SIGNING_KEY_FILE
	defer func() { types.OIDC_OP_SIGNING_KEY_FILE = origFile }()

	_, privKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)
	der, err := x509.MarshalPKCS8PrivateKey(privKey)
	require.NoError(t, err)

	dir := t.TempDir()
	keyFile := filepath.Join(dir, "signing.pem")
	require.NoError(t, os.WriteFile(keyFile, pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}), 0o600))

	t.Run("from file: same key after restart", func(t *testing.T) {
		types.OIDC_OP_SIGNING_KEY_FILE = keyFile

		s1, err := NewStorage()
		require.NoError(t, err)
		s2, err := NewStorage()
		require.NoError(t, err)

		assert.Equal(t, privKey, s1.signingKey.key)
		assert.Equal(t, s1.signingKey.key, s2.signingKey.key)
		assert.Equal(t, s1.signingKey.pubkey, s2.signingKey.pubkey)
	})

	t.Run("no file: random key", func(t *testing.T) {
		types.OIDC_OP_SIGNING_KEY_FILE = ""

		s1, err := NewStorage()
		require.NoError(t, err)
		s2, err := NewStorage()
		require.NoError(t, err)

		assert.NotEqual(t, s1.signingKey.key, s2.signingKey.key)
	})

	t.Run("invalid file", func(t *testing.T) {
		badFile := filepath.Join(dir, "bad.pem")
		require.NoError(t, os.WriteFile(badFile, []byte("not a pem"), 0o600))

		types.OIDC_OP_SIGNING_KEY_FILE = badFile
		_, err := NewStorage()
		assert.ErrorIs(t, err, ErrInvalidSigningKey)

		types.OIDC_OP_SIGNING_KEY_FILE = filepath.Join(dir, "missing.pem")
		_, err = NewStorage()
		assert.Error(t, err)
	})
}
