package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthRateLimitByIP(t *testing.T) {
	setupTest()
	defer teardownTest()

	origLimit := types.AUTH_RATE_LIMIT_PER_IP
	types.AUTH_RATE_LIMIT_PER_IP = 3
	defer func() { types.AUTH_RATE_LIMIT_PER_IP = origLimit }()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/login", AuthRateLimitByIP, func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	post := func(remoteAddr string) int {
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		req.RemoteAddr = remoteAddr
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		return w.Code
	}

	for i := 0; i < 3; i++ {
		assert.Equal(t, http.StatusOK, post("203.0.113.7:1234"))
	}
	assert.Equal(t, http.StatusTooManyRequests, post("203.0.113.7:1234"))

	// other clients are not affected.
	assert.Equal(t, http.StatusOK, post("203.0.113.8:1234"))
}
