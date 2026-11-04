package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const LOGOUT_R = "/account/logout"

func LogoutWrapper(c *gin.Context) {
	FormLogout(Logout, nil, c)
}

func Logout(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	jwt := pttbbsapi.GetJwt(c) // get jwt from access-token

	if jwt == "" {
		jwt = utils.GetCookie(c, types.ACCESS_TOKEN_NAME)
	}

	userID, _, _, err := pttbbsapi.VerifyJwt(jwt, true)
	if err == nil {
		userVisit := &schema.UserVisit{
			UserID:       userID,
			Action:       c.Request.Method + ":" + c.Request.URL.Path,
			UpdateNanoTS: types.NowNanoTS(),
		}
		_ = schema.UpdateUserVisit(userVisit)
	}

	removeTokenFromCookie(c)

	return nil, 200, nil
}
