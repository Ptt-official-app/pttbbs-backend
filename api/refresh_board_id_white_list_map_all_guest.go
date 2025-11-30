package api

import (
	"github.com/gin-gonic/gin"
)

const REFRESH_BOARD_ID_WHITE_LIST_MAP_R = "/boards/refresh-white-list"

type RefreshBoardIDWhiteListMapResult struct {
	Total int `json:"total"`
}

func RefreshBoardIDWhiteListMapAllGuestWrapper(c *gin.Context) {
	Query(LoadPopularBoardsAllGuest, nil, c)
}

func RefreshBoardIDWhiteListMapAllGuest(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	nBoardID, err := RefreshBoardIDWhiteListMap()
	if err != nil {
		return 0, 500, err
	}

	result = &RefreshBoardIDWhiteListMapResult{
		Total: nBoardID,
	}

	return result, 200, nil
}
