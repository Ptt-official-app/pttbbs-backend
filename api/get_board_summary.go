package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const GET_BOARD_SUMMARY_R = "/board/:bid/summary"

type GetBoardSummaryParams struct{}

type GetBoardSummaryPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type GetBoardSummaryResult *apitypes.BoardSummary

func GetBoardSummaryWrapper(c *gin.Context) {
	params := &GetBoardSummaryParams{}
	path := &GetBoardSummaryPath{}
	LoginRequiredPathQuery(GetBoardSummary, params, path, c)
}

func GetBoardSummary(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	thePath, ok := path.(*GetBoardSummaryPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 400, err
	}

	_, err = CheckUserBoardPermReadable(user, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	// backend get-board-summary
	theParams_b := &pttbbsapi.LoadBoardSummaryParams{}

	var result_b pttbbsapi.LoadBoardSummaryResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.LOAD_BOARD_SUMMARY_R)
	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	// update to db
	theList_b := []*bbs.BoardSummary{(*bbs.BoardSummary)(result_b)}
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializeBoardsAndUpdateDB(userID, theList_b, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	// check isRead
	userBoardInfoMap, err = checkUserReadBoard(userID, userBoardInfoMap, boardSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	userBoardInfo, ok := userBoardInfoMap[boardID]
	if !ok {
		return nil, 500, ErrNoBoard
	}

	boardSummary_db := boardSummaries_db[0]
	boardSummary := apitypes.NewBoardSummary(boardSummary_db, "", userBoardInfo, userID)

	// result
	result = GetBoardSummaryResult(boardSummary)

	return result, 200, nil
}
