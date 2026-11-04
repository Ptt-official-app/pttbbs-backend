package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const SET_ID_EMAIL_R = "/user/:user_id/setidemail"

type SetIDEmailParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	Jwt          string `json:"token" form:"token" url:"token"`
}

type SetIDEmailPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type SetIDEmailResult struct {
	Email string `json:"idemail"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func SetIDEmailWrapper(c *gin.Context) {
	params := &SetIDEmailParams{}
	path := &SetIDEmailPath{}

	LoginRequiredPathJSON(SetIDEmail, params, path, c)
}

func SetIDEmail(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*SetIDEmailParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*SetIDEmailPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	// get email token info
	queryUserID, idEmail, queryClientInfo, statusCode, err := getEmailTokenInfo(theParams.Jwt, pttbbsapi.CONTEXT_SET_ID_EMAIL, c)
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
	err = schema.CreateUserIDEmail(queryUserID, idEmail, updateNanoTS)
	if err != nil {
		return nil, 403, err
	}

	// get backend data
	theParams_b := &pttbbsapi.SetIDEmailParams{
		Jwt:   theParams.Jwt,
		IsSet: true,
	}

	var result_b *pttbbsapi.SetIDEmailResult

	urlMap := map[string]string{
		"uid": string(thePath.UserID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.SET_ID_EMAIL_R)

	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	// update db-record to complete the record.
	updateNanoTS = types.NowNanoTS()
	err = schema.UpdateUserIDEmailIsSet(queryUserID, result_b.Email, true, updateNanoTS)
	if err != nil {
		return nil, statusCode, err
	}

	return &SetIDEmailResult{Email: result_b.Email, TokenUser: user.UserID}, 200, nil
}
