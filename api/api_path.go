package api

import (
	"strings"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func PathQuery(theFunc PathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
	err := c.ShouldBindQuery(params)
	if err != nil {
		processResult(c, nil, 400, err, "")
		return
	}

	pathProcess(theFunc, params, path, c)
}

func pathProcess(theFunc PathAPIFunc, params interface{}, path interface{}, c *gin.Context) {
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

	isOver18 := verifyIsOver18(c)

	user := &UserInfo{IsOver18: isOver18, UserID: bbs.UUserID(pttbbsapi.GUEST)}

	result, statusCode, err := theFunc(remoteAddr, user, params, path, c)
	processResult(c, result, statusCode, err, "")
}
