package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const GET_USER_ID_R = "/userid"

type GetUserIDResult struct {
	UserID bbs.UUserID `json:"user_id"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func GetUserIDWrapper(c *gin.Context) {
	LoginRequiredQuery(GetUserID, nil, c)
}

func GetUserID(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	userID := user.UserID
	return &GetUserIDResult{
		UserID: userID,

		TokenUser: userID,
	}, 200, nil
}
