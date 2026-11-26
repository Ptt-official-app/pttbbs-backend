package oidcop

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAUTH_CALLBACK_URL(t *testing.T) {
	setupTest()
	defer teardownTest()

	router := gin.New()
	require.NoError(t, InitGinRouter(router))
	defer func() {
		PROVIDER = nil
		AUTH_CALLBACK_URL = nil
	}()

	// same kind of context api.Login passes in: a gin context without any issuer.
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	got := AUTH_CALLBACK_URL(c, "req-1")
	assert.Equal(t, types.OIDC_OP_ISSUER+"/api/oidc/auth/callback?id=req-1", got)

	// the returned URL must be served by the router the OP is mounted on.
	u, err := url.Parse(got)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest(http.MethodGet, u.RequestURI(), nil))
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}
