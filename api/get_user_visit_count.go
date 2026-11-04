package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const GET_USER_VISIT_COUNT_R = "/uservisitcount"

type GetUserVisitCountResult struct {
	Total int64 `json:"total"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func GetUserVisitCountWrapper(c *gin.Context) {
	Query(GetUserVisitCount, nil, c)
}

func GetUserVisitCount(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	// get backend(go-pttbbs || c-pttbbs) user count
	url := pttbbsapi.GET_USER_VISIT_COUNT_R
	var result_b *pttbbsapi.GetUserVisitCountResult
	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, 500, err
	}

	// get pttbbsweb user count
	currentUserVisitCount := schema.GetUserVisitCount()
	// total user
	result = &GetUserVisitCountResult{Total: int64(result_b.Total) + currentUserVisitCount}
	return result, 200, nil
}
