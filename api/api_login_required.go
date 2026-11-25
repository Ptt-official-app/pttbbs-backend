package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginRequiredJSON(theFunc LoginRequiredAPIFunc, params any, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func LoginRequiredQuery(theFunc LoginRequiredAPIFunc, params any, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	loginRequiredProcess(theFunc, params, c)
}

func loginRequiredProcess(theFunc LoginRequiredAPIFunc, params any, c *gin.Context) {
	remoteAddr, user, statusCode, err := loginRequiredCore(c)
	if err != nil {
		processResult(c, nil, statusCode, err, "")
		return
	}

	result, statusCode, err := theFunc(remoteAddr, user, params, c)
	processResult(c, result, statusCode, err, user.UserID)
}

func loginRequiredCore(c *gin.Context) (remoteAddr string, user *UserInfo, statusCode int, err error) {
	remoteAddr = strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		return "", nil, 403, ErrInvalidRemoteAddr
	}

	userID, isOver18, statusCode, err := getLoginUser(c)
	if err != nil {
		return "", nil, statusCode, err
	}

	user = &UserInfo{IsOver18: isOver18, UserID: userID}

	return remoteAddr, user, 200, nil
}
