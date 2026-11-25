package api

import (
	"time"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const REFRESH_R = "/account/refresh"

type RefreshParams struct {
	ClientID     string `json:"client_id" form:"client_id"`
	ClientSecret string `json:"client_secret" form:"client_secret"`

	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

type RefreshResult LoginResult

func RefreshWrapper(c *gin.Context) {
	params := &RefreshParams{}
	FormJSON(Refresh, params, c)
}

// Refresh rotates the OIDC refresh token issued by Login:
// a new access token and a new refresh token are issued,
// and the old refresh token (and the access token issued with it) are revoked.
func Refresh(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*RefreshParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	if theParams.RefreshToken == "" {
		return nil, 400, ErrInvalidParams
	}

	refreshToken_db, err := schema.GetOIDCRefreshToken(theParams.RefreshToken)
	if err != nil || refreshToken_db == nil {
		return nil, 401, ErrInvalidToken
	}

	if theParams.ClientID != "" && theParams.ClientID != refreshToken_db.ApplicationID {
		return nil, 401, ErrInvalidToken
	}

	if refreshToken_db.Expiration.Before(time.Now()) {
		_ = schema.DeleteOIDCRefreshToken(theParams.RefreshToken)
		return nil, 401, ErrInvalidToken
	}

	username := refreshToken_db.Subject
	clientID := refreshToken_db.ApplicationID

	accessToken, newRefreshToken, expireTS, err := genAccessAndRefreshTokens(c, username, clientID, theParams.RefreshToken)
	if err != nil {
		logrus.Warnf("api.Refresh: unable to genAccessAndRefreshTokens: username: %v clientID: %v e: %v", username, clientID, err)
		return nil, 401, ErrInvalidToken
	}

	if clientID == types.WEB_CLIENT_ID {
		setTokenToCookie(c, accessToken, newRefreshToken)
	}

	loginResult := NewLoginResult(username, accessToken, newRefreshToken, expireTS, "")
	result = (*RefreshResult)(loginResult)

	return result, 200, nil
}
