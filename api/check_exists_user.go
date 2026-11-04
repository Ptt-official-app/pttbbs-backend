package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const CHECK_EXISTS_USER_R = "/account/existsuser"

type CheckExistsUserParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	Username string `json:"username" form:"username"`
}

type CheckExistsUserResult struct {
	IsExists bool `json:"is_exists"`
}

func CheckExistsUserWrapper(c *gin.Context) {
	params := &CheckExistsUserParams{}
	FormJSON(CheckExistsUser, params, c)
}

func CheckExistsUser(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*CheckExistsUserParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	isValidClient, _ := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	// check existing user
	theParams_b := &pttbbsapi.CheckExistsUserParams{
		Username: theParams.Username,
	}
	var result_b *pttbbsapi.CheckExistsUserResult

	url := pttbbsapi.CHECK_EXISTS_USER_R
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	return &CheckExistsUserResult{IsExists: result_b.IsExists}, 200, nil
}
