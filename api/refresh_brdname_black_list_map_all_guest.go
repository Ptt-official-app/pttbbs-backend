package api

import (
	"github.com/gin-gonic/gin"
)

const REFRESH_BRDNAME_BLACK_LIST_MAP_R = "/boards/refresh-black-list"

type RefreshBrdnameBlackListMapResult struct {
	Total int `json:"total"`
}

func RefreshBrdnameBlackListMapAllGuestWrapper(c *gin.Context) {
	Query(RefreshBrdnameBlackListMapAllGuest, nil, c)
}

func RefreshBrdnameBlackListMapAllGuest(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	nBoardID, err := RefreshBrdnameBlackListMap()
	if err != nil {
		return 0, 500, err
	}

	result = &RefreshBrdnameBlackListMapResult{
		Total: nBoardID,
	}

	return result, 200, nil
}
