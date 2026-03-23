package api

import (
	"github.com/gin-gonic/gin"
)

const REFRESH_BRDNAME_WHITE_LIST_MAP_R = "/boards/refresh-white-list"

type RefreshBrdnameWhiteListMapResult struct {
	Total int `json:"total"`
}

func RefreshBrdnameWhiteListMapAllGuestWrapper(c *gin.Context) {
	Query(RefreshBrdnameWhiteListMapAllGuest, nil, c)
}

func RefreshBrdnameWhiteListMapAllGuest(remoteAddr string, user *UserInfo, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	nBoardID, err := RefreshBrdnameWhiteListMap()
	if err != nil {
		return 0, 500, err
	}

	result = &RefreshBrdnameWhiteListMapResult{
		Total: nBoardID,
	}

	return result, 200, nil
}
