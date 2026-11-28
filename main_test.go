package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newGinEngine_trustedProxies(t *testing.T) {
	setupTest()
	defer teardownTest()

	origTrustedProxies := types.TRUSTED_PROXIES
	defer func() { types.TRUSTED_PROXIES = origTrustedProxies }()

	tests := []struct {
		name           string
		trustedProxies []string
		remoteAddr     string
		xff            string
		wantIP         string
	}{
		{
			name:           "untrusted peer cannot spoof",
			trustedProxies: []string{"127.0.0.1", "::1"},
			remoteAddr:     "203.0.113.7:1234",
			xff:            "127.0.0.1",
			wantIP:         "203.0.113.7",
		},
		{
			name:           "trusted proxy forwards client ip",
			trustedProxies: []string{"127.0.0.1", "::1"},
			remoteAddr:     "127.0.0.1:1234",
			xff:            "203.0.113.7",
			wantIP:         "203.0.113.7",
		},
		{
			name:           "empty list trusts nobody",
			trustedProxies: []string{""},
			remoteAddr:     "127.0.0.1:1234",
			xff:            "203.0.113.7",
			wantIP:         "127.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			types.TRUSTED_PROXIES = tt.trustedProxies

			router, err := newGinEngine()
			require.NoError(t, err)

			var got string
			router.GET("/ip", func(c *gin.Context) { got = c.ClientIP() })

			req := httptest.NewRequest(http.MethodGet, "/ip", nil)
			req.RemoteAddr = tt.remoteAddr
			req.Header.Set("X-Forwarded-For", tt.xff)
			router.ServeHTTP(httptest.NewRecorder(), req)

			assert.Equal(t, tt.wantIP, got)
		})
	}

	types.TRUSTED_PROXIES = []string{"not-an-ip"}
	_, err := newGinEngine()
	assert.Error(t, err)
}

func Test_newGinEngine_securityHeaders(t *testing.T) {
	setupTest()
	defer teardownTest()

	router, err := newGinEngine()
	require.NoError(t, err)
	router.GET("/test", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	router.NoRoute(func(c *gin.Context) { c.String(http.StatusOK, "index.html") })

	for _, path := range []string{"/test", "/no-such-route"} {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))

		assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"), path)
		assert.Equal(t, "SAMEORIGIN", w.Header().Get("X-Frame-Options"), path)
		assert.Equal(t, "frame-ancestors 'self'", w.Header().Get("Content-Security-Policy"), path)
	}
}
