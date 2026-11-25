package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginRequiredRedirectPathQuery(theFunc LoginRequiredRedirectPathAPIFunc, params any, path any, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredRedirectPathProcess(theFunc, params, path, c)
}

func loginRequiredRedirectPathProcess(theFunc LoginRequiredRedirectPathAPIFunc, params any, path any, c *gin.Context) {
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

	userID, isOver18, statusCode, err := getLoginUser(c)
	if err != nil {
		processResult(c, nil, statusCode, err, "")
		return
	}

	user := &UserInfo{IsOver18: isOver18, UserID: userID}

	redirectPath, statusCode, err := theFunc(remoteAddr, user, params, path, c)
	processRedirectResult(c, redirectPath, statusCode, err, userID)
}
