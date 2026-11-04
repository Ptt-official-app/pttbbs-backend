package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

func LoadPopularBoardsAllGuestWrapper(c *gin.Context) {
	Query(LoadPopularBoardsAllGuest, nil, c)
}

func LoadPopularBoardsAllGuest(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	boardSummaries_db, err := schema.GetPopularBoardSummaries()
	if err != nil {
		return nil, 500, err
	}

	userBoardInfoMap := make(map[bbs.BBoardID]*apitypes.UserBoardInfo)
	for _, board := range boardSummaries_db {
		userBoardInfoMap[board.BBoardID] = &apitypes.UserBoardInfo{
			Stat: ptttype.NBRD_BOARD,
		}
	}

	user.UserID = bbs.UUserID(pttbbsapi.GUEST)

	result = NewLoadPopularBoardsResult(boardSummaries_db, userBoardInfoMap, user.UserID)

	return result, 200, nil
}
