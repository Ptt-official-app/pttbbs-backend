package api

import (
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

const GET_MY_USER_INFO_R = "/user/me"

func GetMyUserInfoWrapper(c *gin.Context) {
	params := &GetUserInfoParams{}
	LoginRequiredQuery(GetMyUserInfo, params, c)
}

// GetMyUserInfo returns the info of the user of the access token.
//
// The user is looked up by user-id, because the username is not the user-id
// (genUsername assigns a random username on registration).
func GetMyUserInfo(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	getUserDetail := func() (*schema.UserDetail, error) {
		return schema.GetUserDetailByUserID(user.UserID)
	}

	return getUserInfo(user, getUserDetail, c)
}
