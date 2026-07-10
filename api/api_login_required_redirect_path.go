package api

import (
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoginRequiredRedirectPathQuery(theFunc LoginRequiredRedirectPathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredRedirectPathProcess(theFunc, params, path, c)
}

func loginRequiredRedirectPathProcess(theFunc LoginRequiredRedirectPathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindUri(path)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 403, ErrInvalidRemoteAddr, "")
		return
	}

	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin, "")
		return
	}

	userID, err := verifyJwt(c)
	if err != nil {
		userID = bbs.UUserID(pttbbsapi.GUEST)
	}

	isOver18 := verifyIsOver18(c)

	user := &UserInfo{IsOver18: isOver18, UserID: userID}

	redirectPath, statusCode, err := theFunc(remoteAddr, user, params, path, c)
	processRedirectResult(c, redirectPath, statusCode, err, userID)
}
