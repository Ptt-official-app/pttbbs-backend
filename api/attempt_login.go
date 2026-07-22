package api

import (
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const ATTEMPT_LOGIN_R = "/account/attemptlogin"

type AttemptLoginParams struct {
	Input string `json:"input" form:"input"`

	OIDCID string `json:"oidc_id"`
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
		logrus.Errorf("api.AttemptLogin: unable to get loginInputToUsernameEmail: input: %v e: %v", theParams.Input, err)
		return &AttemptLoginResult{}, 200, nil
	}

	err = gen2FATokenAndSendEmail(userID, username, email, theParams.OIDCID, types.ATTEMPT_LOGIN_TITLE, types.ATTEMPT_LOGIN_TEMPLATE_CONTENT, types.EXPIRE_ATTEMPT_LOGIN_EMAIL_TS_DURATION)
	if err != nil {
		logrus.Errorf("api.AttemptLogin: unable to gen2FATokenAndSendEmail: input: %v userID: %v email: %v e: %v", theParams.Input, userID, email, err)
		return &AttemptLoginResult{}, 200, nil
	}

	return &AttemptLoginResult{}, 200, nil
}
