package api

import (
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const ATTEMPT_LOGIN_R = "/account/attemptlogin"

type AttemptLoginParams struct {
	Input string `json:"input" form:"input"`
}

type AttemptLoginResult struct{}

func AttemptLoginWrapper(c *gin.Context) {
	params := &AttemptLoginParams{}
	FormJSON(AttemptLogin, params, c)
}

func AttemptLogin(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*AttemptLoginParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	userID, username, email, err := loginInputToUsernameEmail(theParams.Input)
	if err != nil {
		return &AttemptLoginResult{}, 200, nil
	}

	err = gen2FATokenAndSendEmail(userID, username, email, types.ATTEMPT_LOGIN_TITLE, types.ATTEMPT_LOGIN_TEMPLATE_CONTENT, types.EXPIRE_ATTEMPT_LOGIN_EMAIL_TS_DURATION)
	if err != nil {
		return nil, 200, nil
	}

	return &AttemptLoginResult{}, 200, nil
}
