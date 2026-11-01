package lemmyutils

import (
	"net/http"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

type GetTokenParams struct {
	Username string `json:"username"`
}

type GetTokenResult = LoginResult

func TryLemmyGetToken(username string, c *gin.Context) (token string, err error) {
	err = TryLemmyLogin(c)
	if err != nil {
		return "", err
	}

	token, err = LemmyGetToken(username, c)
	if err == nil {
		return token, nil
	}

	if err != ErrInvalidBackendStatusCode {
		return "", err
	}

	err = LemmyLogin(c)
	if err != nil {
		return "", err
	}

	return LemmyGetToken(username, c)
}

func LemmyGetToken(username string, c *gin.Context) (token string, err error) {
	url := types.PTTBBS_LEMMY_PREFIX + "/pttbbs-lemmy/token"

	params := &GetTokenParams{
		Username: username,
	}

	var ret *GetTokenResult
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + LEMMY_ADMIN_TOKEN,
	}

	statusCode, err := utils.BackendPost(c, url, params, headers, &ret)
	if err != nil {
		return "", err
	}
	if statusCode != http.StatusOK {
		return "", ErrInvalidBackendStatusCode
	}

	return ret.JWT, nil
}
