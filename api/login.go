package api

import (
	"fmt"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const LOGIN_R = "/account/login"

type LoginParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	// Input can be username or email
	Input string `json:"input" form:"input"`

	VerifyCode string `json:"verify_code" form:"verify_code"`
}

func NewLoginParams() *LoginParams {
	return &LoginParams{}
}

type LoginResult struct {
	Username        string      `json:"username"`
	AccessToken     string      `json:"access_token"`
	TokenType       string      `json:"token_type"`
	RefreshToken    string      `json:"refresh_token"`
	AccessExpireTS  types.Time8 `json:"access_expire"`
	RefreshExpireTS types.Time8 `json:"refresh_expire"`
}

// LoginLog record user login info, no matter success or not
type LoginLog struct {
	ClientInfo
	LoginID   string
	LoginTime types.NanoTS
	LoginIP   string
	IsSuccess bool
}

func (l *LoginLog) String() string {
	var success string
	if l.IsSuccess {
		success = "\033[97;42mSuccess\033[0m"
	} else {
		success = "\033[97;41mFail\033[0m"
	}
	return fmt.Sprintf("ID: %s login %s from %s at %v Client: %v \n", l.LoginID, success, l.LoginIP, l.LoginTime.ToTime(), l.ClientInfo)
}

func LoginWrapper(c *gin.Context) {
	params := NewLoginParams()
	FormJSON(Login, params, c)
}

func Login(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoginParams)
	// record user login
	loginLog := &LoginLog{
		ClientInfo: ClientInfo{
			ClientID: theParams.ClientID,
		},
		LoginID:   theParams.Input,
		LoginIP:   remoteAddr,
		LoginTime: types.NowNanoTS(),
		IsSuccess: false, // default is false
	}
	defer func() {
		logrus.Infof("%v", loginLog)
	}()

	if !ok {
		return nil, 400, ErrInvalidParams
	}

	// XXX skip client-info for now.
	/*
		isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)
		if !isValidClient {
			return nil, 401, ErrInvalidParams
		}

		clientInfo := getClientInfo(client)
	*/
	clientInfo := ""

	userID, username, _, err := loginInputToUsernameEmail(theParams.Input)
	if err != nil {
		return nil, 401, err
	}

	err = check2FAToken(userID, theParams.VerifyCode)
	if err != nil {
		return nil, 401, err
	}

	// gen tokens
	accessToken, accessExpireTS, err := genAccessToken(userID, clientInfo)
	if err != nil {
		return nil, 500, err
	}

	refreshToken, refresExpireTS, err := genRefreshToken(userID, clientInfo)
	if err != nil {
		return nil, 500, err
	}

	// update: loginLog success login
	loginLog.IsSuccess = true

	// result
	result = NewLoginResult(username, accessToken, accessExpireTS, refreshToken, refresExpireTS)

	setTokenToCookie(c, accessToken)

	return result, 200, nil
}

func NewLoginResult(username string, accessToken string, accessExpireTS types.Time8, refreshToken string, refreshExpireTS types.Time8) *LoginResult {
	return &LoginResult{
		Username:        username,
		TokenType:       "bearer",
		AccessToken:     accessToken,
		AccessExpireTS:  accessExpireTS,
		RefreshToken:    refreshToken,
		RefreshExpireTS: refreshExpireTS,
	}
}
