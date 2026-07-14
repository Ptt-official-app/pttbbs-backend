package api

import (
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const REGISTER_USER_R = "/account/register"

type RegisterUserParams struct {
	Token string `json:"token" form:"token" url:"token"`
}

func NewRegisterUserParams() *RegisterUserParams {
	return &RegisterUserParams{}
}

func RegisterUserWrapper(c *gin.Context) {
	params := NewRegisterUserParams()
	RedirectQuery(RegisterUser, params, c)
}

func RegisterUser(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result string, statusCode int, err error) {
	theParams, ok := params.(*RegisterUserParams)
	if !ok {
		return types.FRONTEND_ERR_URL, 303, ErrInvalidParams
	}

	email, err := getEmailFromEmailVerificationToken(theParams.Token)
	if err != nil {
		return types.FRONTEND_ERR_URL, 303, err
	}

	// create db-record first to avoid race-condition
	userID, err := genUserID(email)
	if err != nil {
		return types.FRONTEND_ERR_URL, 303, err
	}

	username, err := genUsername(userID)
	if err != nil {
		return types.FRONTEND_ERR_URL, 303, err
	}

	accessToken, refreshToken, _, err := genAccessAndRefreshTokens(c, username, types.WEB_CLIENT_ID, "")
	if err != nil {
		return types.FRONTEND_ERR_URL, 303, err
	}

	// result
	setTokenToCookie(c, accessToken, refreshToken)

	return types.FRONTEND_INIT_URL, 303, nil
}
