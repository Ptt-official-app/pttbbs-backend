package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/gin-gonic/gin"
)

const LOAD_GENERAL_BOARDS_BY_CLASS_R = "/boards/byclass"

func LoadGeneralBoardsByClassWrapper(c *gin.Context) {
	params := NewLoadGeneralBoardsParams()
	LoginRequiredQuery(LoadGeneralBoardsByClass, params, c)
}

func LoadGeneralBoardsByClass(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	return loadGeneralBoardsCore(remoteAddr, user, params, c, pttbbsapi.LOAD_GENERAL_BOARDS_BY_CLASS_R)
}
