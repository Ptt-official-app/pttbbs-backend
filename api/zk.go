package api

import (
	"context"
	"net/http/httputil"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const ZK_CREATE_CHALLENGE_R = "/challenge"

const ZK_GET_CHALLENGE_R = "/challenge/:challenge"

const ZK_LINK_VERIFY_R = "/link-verify"

const ZK_SMT_ROOT_STATUS_R = "/smt-root/status"

const ZK_ISSUER_CERT_STATUS_R = "/issuer-cert/status"

func ZKProxyWrapper(zkProxy *httputil.ReverseProxy) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1. loginRequiredCore
		_, user, err := loginRequiredCore(c)
		if err != nil {
			logrus.Errorf("ZKProxyWrapper: unable to pass loginRequired: e: %v", err)
			processResult(c, nil, 403, ErrInvalidRemoteAddr, "")
			return
		}

		if user.UserID == bbs.UUserID(pttbbsapi.GUEST) {
			processResult(c, nil, 401, ErrInvalidUser, "")
			return
		}

		// 2. set context.
		//nolint:staticcheck // SA1029: ZK_USER_ID_KEY is one-time-use for userID in LinkVerifyProxy.
		ctx := context.WithValue(c.Request.Context(), types.ZK_USER_ID_KEY, user.UserID)
		c.Request = c.Request.WithContext(ctx)

		// 3. serveHTTP
		zkProxy.ServeHTTP(c.Writer, c.Request)
	}
}
