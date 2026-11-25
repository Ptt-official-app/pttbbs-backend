package api

import (
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const LOGOUT_R = "/account/logout"

func LogoutWrapper(c *gin.Context) {
	FormLogout(Logout, nil, c)
}

// Logout revokes the access token (and its refresh token) in the token store,
// and removes the tokens from the cookies.
//
// The access token is a JWT. Without revoking it in the token store
// it would still be valid until it expires.
func Logout(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	token := utils.GetAccessToken(c)

	if token != "" {
		claim, err := parseAccessToken(token)
		if err == nil {
			err = revokeAccessToken(claim.JWTID)
			if err != nil {
				logrus.Warnf("api.Logout: unable to revokeAccessToken: tokenID: %v e: %v", claim.JWTID, err)
			}

			userUsernameOver18, err := schema.GetUserUsernameOver18ByUsername(claim.Subject)
			if err == nil && userUsernameOver18 != nil {
				userVisit := &schema.UserVisit{
					UserID:       userUsernameOver18.UserID,
					Action:       c.Request.Method + ":" + c.Request.URL.Path,
					UpdateNanoTS: types.NowNanoTS(),
				}
				_ = schema.UpdateUserVisit(userVisit)
			}
		}
	}

	refreshToken := utils.GetCookie(c, types.REFRESH_TOKEN_NAME)
	if refreshToken != "" {
		err = revokeRefreshToken(refreshToken)
		if err != nil {
			logrus.Warnf("api.Logout: unable to revokeRefreshToken: e: %v", err)
		}
	}

	removeTokenFromCookie(c)

	return nil, 200, nil
}
