package api

import (
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func JSON(theFunc APIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	process(theFunc, params, c)
}

func Form(theFunc APIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindWith(params, binding.Form)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	process(theFunc, params, c)
}

func FormJSON(theFunc APIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindJSON(params)
	if err != nil {
		err = c.ShouldBindWith(params, binding.Form)
		if err != nil {
			processResult(c, nil, 400, err, "")
			return
		}
	}

	process(theFunc, params, c)
}

func FormLogout(theFunc APIFunc, params interface{}, c *gin.Context) {
	processLogout(theFunc, params, c)
}

func Query(theFunc APIFunc, params interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	process(theFunc, params, c)
}

func process(theFunc APIFunc, params interface{}, c *gin.Context) {
	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 400, ErrInvalidRemoteAddr, "")
		return
	}

	isOver18 := verifyIsOver18(c)

	user := &UserInfo{IsOver18: isOver18, UserID: bbs.UUserID(pttbbsapi.GUEST)}

	result, statusCode, err := theFunc(remoteAddr, user, params, c)
	processResult(c, result, statusCode, err, "")
}

func processLogout(theFunc APIFunc, params interface{}, c *gin.Context) {
	remoteAddr := strings.TrimSpace(c.ClientIP())
	if !isValidRemoteAddr(remoteAddr) {
		processResult(c, nil, 400, ErrInvalidRemoteAddr, "")
		return
	}

	isOver18 := verifyIsOver18(c)

	user := &UserInfo{IsOver18: isOver18, UserID: bbs.UUserID(pttbbsapi.GUEST)}

	_, _, _ = theFunc(remoteAddr, user, params, c)

	c.Redirect(303, types.FRONTEND_PREFIX)
}
