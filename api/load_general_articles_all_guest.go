package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func LoadGeneralArticlesAllGuestWrapper(c *gin.Context) {
	params := NewLoadGeneralArticlesParams()
	path := &LoadGeneralArticlesPath{}
	PathQuery(LoadGeneralArticlesAllGuest, params, path, c)
}

func LoadGeneralArticlesAllGuest(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	user.UserID = bbs.UUserID(pttbbsapi.GUEST)

	return LoadGeneralArticles(remoteAddr, user, params, path, c)
}
