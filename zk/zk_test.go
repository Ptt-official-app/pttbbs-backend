package zk

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
)

func TestNewZKProxyDoesNotForwardSession(t *testing.T) {
	var upstreamHeader http.Header
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upstreamHeader = r.Header.Clone()
		http.SetCookie(w, &http.Cookie{Name: "upstream", Value: "1"})
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"verified":false}`))
	}))
	defer upstream.Close()

	origZKPrefix := types.ZK_PREFIX
	types.ZK_PREFIX = upstream.URL
	defer func() { types.ZK_PREFIX = origZKPrefix }()

	for name, proxy := range map[string]http.Handler{"zk": NewZKProxy(), "link-verify": NewZKLinkVerifyProxy()} {
		t.Run(name, func(t *testing.T) {
			upstreamHeader = nil

			req := httptest.NewRequest(http.MethodGet, "/challenge/test", nil)
			req.Header.Set("Authorization", "bearer access-token")
			req.Header.Set("Cookie", "token=access-token; refresh_token=refresh-token")
			req.Header.Set("X-CSRFToken", "csrf-token")
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			proxy.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Empty(t, upstreamHeader.Get("Authorization"))
			assert.Empty(t, upstreamHeader.Get("Cookie"))
			assert.Empty(t, upstreamHeader.Get("X-CSRFToken"))
			assert.Equal(t, "application/json", upstreamHeader.Get("Content-Type"))
			assert.Equal(t, req.Host, upstreamHeader.Get("X-Forwarded-Host"))
			assert.NotEmpty(t, upstreamHeader.Get("X-Forwarded-For"))

			assert.Empty(t, w.Header().Values("Set-Cookie"))
			assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
		})
	}
}
