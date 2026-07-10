package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const CHANGE_EMAIL_R = "/user/:user_id/changeemail"

type ChangeEmailParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	Jwt          string `json:"token" form:"token" url:"token"`
}

type ChangeEmailPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type ChangeEmailResult struct {
	Email string `json:"email"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func ChangeEmailWrapper(c *gin.Context) {
	params := &ChangeEmailParams{}
	path := &ChangeEmailPath{}

	LoginRequiredPathJSON(ChangeEmail, params, path, c)
}

func ChangeEmail(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*ChangeEmailParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*ChangeEmailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	// get email token info
	queryUserID, email, queryClientInfo, statusCode, err := getEmailTokenInfo(theParams.Jwt, pttbbsapi.CONTEXT_CHANGE_EMAIL, c)
	if err != nil {
		return nil, statusCode, err
	}

	if clientInfo != queryClientInfo {
		return nil, 403, ErrInvalidClient
	}

	if thePath.UserID != queryUserID {
		return nil, 403, ErrInvalidUser
	}

	// create db-record first to avoid race-condition
	updateNanoTS := types.NowNanoTS()
	err = schema.CreateUserEmail(queryUserID, email, false, updateNanoTS)
	if err != nil {
		return nil, 403, err
	}

	// get backend data
	theParams_b := &pttbbsapi.ChangeEmailParams{
		Jwt: theParams.Jwt,
	}

	var result_b *pttbbsapi.ChangeEmailResult

	urlMap := map[string]string{
		"uid": string(thePath.UserID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.CHANGE_EMAIL_R)

	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	// update db-record to complete the record.
	updateNanoTS = types.NowNanoTS()
	err = schema.UpdateUserEmailIsDefault(queryUserID, result_b.Email, true, updateNanoTS)
	if err != nil {
		return nil, statusCode, err
	}

	// update user-info

	return &ChangeEmailResult{Email: result_b.Email, TokenUser: user.UserID}, 200, nil
}
