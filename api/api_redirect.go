package api

import (
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RedirectQuery(theFunc RedirectAPIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	redirectProcess(theFunc, params, c)
}

func RedirectFormJSON(theFunc RedirectAPIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		err = c.ShouldBindWith(params, binding.Form)
		if err != nil {
			processResult(c, nil, 400, err, "")
			return
		}
	}

	redirectProcess(theFunc, params, c)
}

func redirectProcess(theFunc RedirectAPIFunc, params interface{}, c *gin.Context) {
	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 403, ErrInvalidRemoteAddr, "")
		return
	}

	if !isValidOriginReferer(c) {
		processResult(c, nil, 403, ErrInvalidOrigin, "")
		return
	}

	isOver18 := verifyIsOver18(c)

	userID := bbs.UUserID(pttbbsapi.GUEST)
	user := &UserInfo{IsOver18: isOver18, UserID: userID}

	redirectPath, statusCode, err := theFunc(remoteAddr, user, params, c)
	processRedirectResult(c, redirectPath, statusCode, err, userID)
}
