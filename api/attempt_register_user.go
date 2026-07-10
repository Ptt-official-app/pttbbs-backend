package api

import (
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const ATTEMPT_REGISTER_USER_R = "/account/attemptregister"

type AttemptRegisterUserParams struct {
	Email string `json:"email" form:"email"`
}

type AttemptRegisterUserResult struct{}

func AttemptRegisterUserWrapper(c *gin.Context) {
	params := &AttemptRegisterUserParams{}
	FormJSON(AttemptRegisterUser, params, c)
}

func AttemptRegisterUser(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*AttemptRegisterUserParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	err = genEmailVerificationTokenAndSendEmail(theParams.Email, types.ATTEMPT_REGISTER_USER_TITLE, types.REGISTER_USER_URL, types.ATTEMPT_REGISTER_USER_TEMPLATE_CONTENT)
	if err != nil {
		return nil, 500, err
	}

	return &AttemptRegisterUserResult{}, 200, nil
}
