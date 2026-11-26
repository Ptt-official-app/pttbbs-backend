package schema

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetEmailVerification(t *testing.T) {
	setupTest()
	defer teardownTest()

	expireTSDuration := time.Duration(1) * time.Second
	email := "test@email.dev"

	got, err := SetEmailVerification(email, expireTSDuration)
	require.NoError(t, err)
	assert.False(t, strings.HasSuffix(got, "="), "unexpected = in token: %v", got)

	// the token is 32 random bytes, no email hash is appended.
	tokenBytes, err := base64.RawURLEncoding.DecodeString(got)
	require.NoError(t, err)
	assert.Len(t, tokenBytes, EMAIL_VERIFICATION_TOKEN_LEN)

	emailVerification, err := GetEmailVerification(got)
	require.NoError(t, err)
	assert.Equal(t, email, emailVerification.Email)

	got2, err := SetEmailVerification(email, expireTSDuration)
	require.NoError(t, err)
	assert.NotEqual(t, got, got2)

	// Sleep past the TTL: sleeping exactly the TTL races with Redis expiring got2.
	time.Sleep(expireTSDuration + 500*time.Millisecond)

	_, err = GetEmailVerification(got)
	assert.Error(t, err)
	_, err = GetEmailVerification(got2)
	assert.Error(t, err)
}

// TestGetEmailVerification_legacyFormat checks that pending tokens issued before
// the salt and hash were removed can still be read until they expire.
func TestGetEmailVerification_legacyFormat(t *testing.T) {
	setupTest()
	defer teardownTest()

	// token = random bytes + argon2 hash, value = base64 of {"Email":..,"Salt":..}.
	legacyToken := "bGVnYWN5LXRva2VuLWFuZC1oYXNoLWxvbmctZW5vdWdo"
	legacyValue := base64.URLEncoding.EncodeToString([]byte(`{"Email":"legacy@email.dev","Salt":"c2FsdA=="}`))

	require.NoError(t, db.RDBSet(rdb, RDB_PREFIX_EMAIL+legacyToken, legacyValue, 10*time.Second))
	defer func() { _ = db.RDBDel(rdb, RDB_PREFIX_EMAIL+legacyToken) }()

	got, err := GetEmailVerification(legacyToken)
	require.NoError(t, err)
	assert.Equal(t, "legacy@email.dev", got.Email)
}
