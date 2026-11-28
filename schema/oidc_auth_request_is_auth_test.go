package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

// The auth request is created by the authorize endpoint without a user;
// the user is bound to it only after login (SetAuthRequestIsAuth).
func TestOIDCAuthRequest_SetAuthRequestIsAuth(t *testing.T) {
	setupTest()
	defer teardownTest()

	authReq := NewAuthRequest(&oidc.AuthRequest{ClientID: "client"}, "")
	authReq.ID = "auth-request-is-auth"
	require.NoError(t, SaveAuthRequest(authReq))

	before, err := GetOIDCAuthRequestByRequestID(authReq.ID)
	require.NoError(t, err)
	assert.False(t, before.Done())
	assert.Empty(t, before.GetSubject())
	assert.True(t, before.GetAuthTime().IsZero())

	start := time.Now()
	require.NoError(t, SetAuthRequestIsAuth(authReq.ID, "user"))

	after, err := GetOIDCAuthRequestByRequestID(authReq.ID)
	require.NoError(t, err)
	assert.True(t, after.Done())
	assert.Equal(t, "user", after.GetSubject())
	assert.False(t, after.GetAuthTime().Before(start.Truncate(time.Second)))

	// another auth request is not affected.
	other := NewAuthRequest(&oidc.AuthRequest{ClientID: "client"}, "")
	other.ID = "auth-request-other"
	require.NoError(t, SaveAuthRequest(other))
	got, err := GetOIDCAuthRequestByRequestID(other.ID)
	require.NoError(t, err)
	assert.False(t, got.Done())
	assert.Empty(t, got.GetSubject())
}
