package api

import (
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthRateLimitByIP limits the requests per client ip to the login / register apis.
//
// The per-account limit of check2FAToken does not stop one client from
// guessing codes across many accounts, or from sending mails to any address
// through attemptregister.
func AuthRateLimitByIP(c *gin.Context) {
	clientIP := c.ClientIP()
	count, err := schema.IncrAuthRateLimit(clientIP, types.AUTH_RATE_LIMIT_WINDOW_TS_DURATION)
	if err != nil {
		// the apis behind this limit need redis as well, so they fail on their own.
		logrus.Errorf("AuthRateLimitByIP: unable to IncrAuthRateLimit: clientIP: %v e: %v", clientIP, err)
		return
	}

	if count > types.AUTH_RATE_LIMIT_PER_IP {
		processResult(c, nil, 429, ErrTooManyAttempts, "")
		c.Abort()
	}
}
