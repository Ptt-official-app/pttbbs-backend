package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_initGinCORS(t *testing.T) {
	setupTest()
	defer teardownTest()

	origAllowOrigins := types.ALLOW_ORIGINS
	defer func() { types.ALLOW_ORIGINS = origAllowOrigins }()

	tests := []struct {
		name         string
		allowOrigins []string
		origin       string
		wantAllowed  bool
	}{
		{name: "wildcard", allowOrigins: []string{"*"}, origin: "http://a.example", wantAllowed: true},
		{name: "empty list allows all", allowOrigins: []string{}, origin: "http://a.example", wantAllowed: true},
		{name: "listed origin", allowOrigins: []string{"http://a.example"}, origin: "http://a.example", wantAllowed: true},
		{name: "unlisted origin", allowOrigins: []string{"http://a.example"}, origin: "http://b.example", wantAllowed: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			types.ALLOW_ORIGINS = tt.allowOrigins

			router := gin.New()
			assert.NotPanics(t, func() { initGinCORS(router) })
			router.GET("/", func(c *gin.Context) {})

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Origin", tt.origin)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantAllowed, w.Header().Get("Access-Control-Allow-Origin") != "")
		})
	}
}
