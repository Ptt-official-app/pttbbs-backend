package oidcop

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/db"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zitadel/oidc/v3/pkg/oidc"
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

	got := AUTH_CALLBACK_URL(c, "req-1", "secret-1")
	assert.Equal(t, types.OIDC_OP_ISSUER+"/api/oidc/auth/callback?id=req-1&s=secret-1", got)

	// the returned URL must be served by the router the OP is mounted on.
	u, err := url.Parse(got)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, httptest.NewRequest(http.MethodGet, u.RequestURI(), nil))
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}

func TestRequireAuthCallbackSecret(t *testing.T) {
	setupTest()
	defer teardownTest()

	router := gin.New()
	require.NoError(t, InitGinRouter(router))
	defer func() {
		PROVIDER = nil
		AUTH_CALLBACK_URL = nil
	}()

	const clientID = "test-callback-secret-client"
	const redirectURI = "https://rp.example/callback"
	require.NoError(t, schema.UpdateClient(schema.NewClient(clientID, types.CLIENT_TYPE_APP, []string{redirectURI}, "127.0.0.1")))

	authReq := schema.NewAuthRequest(&oidc.AuthRequest{ClientID: clientID, RedirectURI: redirectURI, ResponseType: oidc.ResponseTypeCode, Scopes: oidc.SpaceDelimitedArray{oidc.ScopeOpenID}}, "")
	authReq.ID = "test-callback-secret"
	require.NoError(t, schema.SaveAuthRequest(authReq))
	defer func() { _ = schema.DeleteOIDCAuthRequest(authReq.ID) }()
	authReqID := authReq.ID

	secret, err := schema.SetAuthRequestIsAuth(authReqID, "SYSOP")
	require.NoError(t, err)

	serveRecorder := func(req *http.Request) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		return w
	}
	serve := func(req *http.Request) int {
		return serveRecorder(req).Code
	}
	get := func(rawQuery string) int {
		return serve(httptest.NewRequest(http.MethodGet, AUTH_CALLBACK_PATH+"?"+rawQuery, nil))
	}

	// whoever only knows the auth request id cannot reach the callback.
	assert.Equal(t, http.StatusForbidden, get("id="+authReqID))
	assert.Equal(t, http.StatusForbidden, get("id="+authReqID+"&s=wrong"))

	// a valid secret must not be usable for another auth request id given in the POST body,
	// which op.AuthorizeCallback reads before the query.
	otherReqID := "test-callback-secret-other"
	otherSecret, err := schema.SetAuthRequestIsAuth(otherReqID, "SYSOP")
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, AUTH_CALLBACK_PATH+"?id="+otherReqID+"&s="+url.QueryEscape(otherSecret), strings.NewReader("id="+authReqID))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Equal(t, http.StatusForbidden, serve(req))

	// the secret returned by login passes the check once, and the OP redirects with the code.
	w := serveRecorder(httptest.NewRequest(http.MethodGet, AUTH_CALLBACK_PATH+"?id="+authReqID+"&s="+url.QueryEscape(secret), nil))
	require.Equal(t, http.StatusFound, w.Code)
	location, err := url.Parse(w.Header().Get("Location"))
	require.NoError(t, err)
	assert.Equal(t, redirectURI, location.Scheme+"://"+location.Host+location.Path)
	assert.NotEmpty(t, location.Query().Get("code"))
	assert.Equal(t, http.StatusForbidden, get("id="+authReqID+"&s="+url.QueryEscape(secret)))

	// other OP endpoints are not affected.
	assert.NotEqual(t, http.StatusForbidden, serve(httptest.NewRequest(http.MethodGet, "/api/oidc/keys", nil)))

	// an auth request that is already logged in cannot be re-bound to another user.
	_, err = schema.SetAuthRequestIsAuth(authReqID, "attacker")
	assert.ErrorIs(t, err, db.ErrRDBAlreadyExists)
}
