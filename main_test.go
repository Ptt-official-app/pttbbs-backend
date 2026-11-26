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
