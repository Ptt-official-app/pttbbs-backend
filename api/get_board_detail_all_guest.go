package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func GetBoardDetailAllGuestWrapper(c *gin.Context) {
	params := &GetBoardDetailParams{}
	path := &GetBoardDetailPath{}
	PathQuery(GetBoardDetailAllGuest, params, path, c)
}

func GetBoardDetailAllGuest(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	user.UserID = bbs.UUserID(pttbbsapi.GUEST)

	return GetBoardDetail(remoteAddr, user, params, path, c)
}
