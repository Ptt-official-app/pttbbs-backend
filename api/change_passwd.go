package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const CHANGE_PASSWD_R = "/user/:user_id/updatepasswd" //nolint // passwd as route

type ChangePasswdParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	OrigPassword    string `json:"orig_password" form:"orig_password"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
}

type ChangePasswdPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type ChangePasswdResult LoginResult

func ChangePasswdWrapper(c *gin.Context) {
	params := &ChangePasswdParams{}
	path := &ChangePasswdPath{}
	LoginRequiredPathJSON(ChangePasswd, params, path, c)
}

func ChangePasswd(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*ChangePasswdParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*ChangePasswdPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	if theParams.Password != theParams.PasswordConfirm {
		return nil, 400, ErrInvalidParams
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	// get backend data
	theParams_b := &pttbbsapi.ChangePasswdParams{
		ClientInfo: clientInfo,
		OrigPasswd: theParams.OrigPassword,
		Passwd:     theParams.Password,
	}

	var result_b *pttbbsapi.ChangePasswdResult

	urlMap := make(map[string]string)
	urlMap["uid"] = string(thePath.UserID)
	url := utils.MergeURL(urlMap, pttbbsapi.CHANGE_PASSWD_R)

	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return nil, statusCode, err
	}

	// result
	ret := NewChangePasswdResult(result_b)

	setTokenToCookie(c, result_b.Jwt)

	return ret, 200, nil
}

func NewChangePasswdResult(result_b *pttbbsapi.ChangePasswdResult) *ChangePasswdResult {
	return &ChangePasswdResult{
		Username:        string(result_b.UserID),
		AccessToken:     result_b.Jwt,
		TokenType:       "bearer",
		RefreshToken:    result_b.Refresh,
		AccessExpireTS:  types.Time8(result_b.AccessExpire),
		RefreshExpireTS: types.Time8(result_b.RefreshExpire),
	}
}
