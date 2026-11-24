package schema

import (
	"errors"
	"testing"
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOIDCTokenNotFound(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, err := GetOIDCAccessToken("not-exist")
	assert.True(t, errors.Is(err, ErrNotFound), "access token: %v", err)

	_, err = GetOIDCRefreshToken("not-exist")
	assert.True(t, errors.Is(err, ErrNotFound), "refresh token: %v", err)

	assert.NoError(t, DeleteOIDCAccessToken("not-exist"))
}

func TestGetOIDCAccessTokensBySubjectAndClientID_expired(t *testing.T) {
	setupTest()
	defer teardownTest()

	exp := time.Now().Add(time.Hour)
	expired := &OIDCAccessToken{ID: "expired", ApplicationID: "app", Subject: "sub", Expiration: exp}
	alive := &OIDCAccessToken{ID: "alive", ApplicationID: "app", Subject: "sub", Expiration: exp}
	require.NoError(t, SetOIDCAccessToken(expired))
	require.NoError(t, SetOIDCAccessToken(alive))

	// the token expires in redis while its id stays in the subject/client set.
	require.NoError(t, db.RDBDel(rdb, RDB_PREFIX_OIDC_OP_ACCESS_TOKEN+expired.ID))

	got, err := GetOIDCAccessTokensBySubjectAndClientID("sub", "app")
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Equal(t, "alive", got[0].ID)

	assert.NoError(t, DeleteOIDCAccessToken(expired.ID))
}
