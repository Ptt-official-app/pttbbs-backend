package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const REFRESH_R = "/account/refresh"

type RefreshParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	RefreshToken string `json:"refresh_token"`
}

type RefreshResult LoginResult

func RefreshWrapper(c *gin.Context) {
	params := &RefreshParams{}
	FormJSON(Refresh, params, c)
}

func Refresh(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*RefreshParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	isValidClient, client := checkClient(theParams.ClientID, theParams.ClientSecret)

	if !isValidClient {
		return nil, 400, ErrInvalidParams
	}

	clientInfo := getClientInfo(client)

	params_b := &pttbbsapi.RefreshParams{
		ClientInfo: clientInfo,
		Refresh:    theParams.RefreshToken,
	}

	result_bi, err := pttbbsapi.Refresh(remoteAddr, params_b, c)
	if err != nil {
		return nil, 500, err
	}

	result_b, ok := result_bi.(*pttbbsapi.RefreshResult)
	if !ok {
		return nil, 500, ErrInvalidBackendStatusCode
	}

	setTokenToCookie(c, result_b.Jwt)

	result = &RefreshResult{
		Username:        string(result_b.UserID),
		AccessToken:     result_b.Jwt,
		TokenType:       result_b.TokenType,
		RefreshToken:    result_b.Refresh,
		AccessExpireTS:  types.Time8(result_b.AccessExpire),
		RefreshExpireTS: types.Time8(result_b.RefreshExpire),
	}

	return result, 200, nil
}
