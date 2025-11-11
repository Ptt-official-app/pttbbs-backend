package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
)

func toBoardID(fboardID apitypes.FBoardID, remoteAddr string, userID bbs.UUserID, c *gin.Context) (boardID bbs.BBoardID, err error) {
	boardID, err = fboardID.ToBBoardID()
	if err != nil {
		return "", err
	}

	// if this site is not for all guest:
	//    we can return boardID.
	if !types.IS_ALL_GUEST {
		return boardID, nil
	}

	// XXX ensure that board is popular.
	boardIsPopular, err := schema.GetBoardIsPopular(boardID)
	if err != nil {
		return "", err
	}
	if !boardIsPopular.IsPopular {
		return "", ErrBoardNotPopular
	}

	return boardID, nil
}

func bidToBoardID(bid ptttype.Bid) (boardID bbs.BBoardID, err error) {
	return schema.GetBoardIDByPttbid(bid)
}

// DeserializeBoards
//
// each_b.Reason happens only with invalid permission.
func DeserializeBoardsAndUpdateDB(boardSummaries_b []*bbs.BoardSummary, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, err error) {
	boardSummaries = make([]*schema.BoardSummary, 0, len(boardSummaries_b))
	for _, each_b := range boardSummaries_b {
		if each_b.Reason != 0 {
			continue
		}
		each := schema.NewBoardSummary(each_b, updateNanoTS)

		boardSummaries = append(boardSummaries, each)
	}
	if len(boardSummaries) == 0 {
		return nil, nil
	}

	err = schema.UpdateBoardSummaries(boardSummaries, updateNanoTS)
	if err != nil {
		return nil, err
	}

	return boardSummaries, nil
}

func deserializeBoardsAndUpdateDB(userID bbs.UUserID, boardSummaries_b []*bbs.BoardSummary, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, err error) {
	if len(boardSummaries_b) == 0 {
		return nil, nil, nil
	}

	boardSummaries, err = DeserializeBoardsAndUpdateDB(boardSummaries_b, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	userReadBoards := make([]*schema.UserReadBoard, 0, len(boardSummaries_b))
	userBoardInfoMap = make(map[bbs.BBoardID]*apitypes.UserBoardInfo)
	for _, each_b := range boardSummaries_b {
		if each_b.Reason != 0 {
			continue
		}

		userBoardInfoMap[each_b.BBoardID] = &apitypes.UserBoardInfo{
			Read: each_b.Read,
			Stat: each_b.StatAttr,
		}

		if each_b.Read {
			each_db := &schema.UserReadBoard{
				UserID:           userID,
				BBoardID:         each_b.BBoardID,
				ReadUpdateNanoTS: updateNanoTS,
			}
			userReadBoards = append(userReadBoards, each_db)
		}
	}

	err = schema.UpdateUserReadBoards(userReadBoards, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	return boardSummaries, userBoardInfoMap, err
}

// DeserializePBBoardsAndUpdateDB
func DeserializePBBoardsAndUpdateDB(boardSummaries_b []*boardd.Board, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, err error) {
	boardSummaries = make([]*schema.BoardSummary, 0, len(boardSummaries_b))
	for _, each_b := range boardSummaries_b {
		each := schema.NewBoardSummaryFromPBBoard(each_b, updateNanoTS)

		boardSummaries = append(boardSummaries, each)
	}
	if len(boardSummaries) == 0 {
		return nil, nil
	}

	err = schema.UpdateBoardSummaries(boardSummaries, updateNanoTS)
	if err != nil {
		return nil, err
	}

	return boardSummaries, nil
}

func deserializePBBoardsAndUpdateDB(boardSummaries_b []*boardd.Board, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, err error) {
	if len(boardSummaries_b) == 0 {
		return nil, nil, nil
	}

	boardSummaries, err = DeserializePBBoardsAndUpdateDB(boardSummaries_b, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	// fav or folder
	userBoardInfoMap = make(map[bbs.BBoardID]*apitypes.UserBoardInfo)
	for _, each := range boardSummaries {
		brdStat := ptttype.NBRD_BOARD
		if each.BrdAttr.HasPerm(ptttype.BRD_GROUPBOARD) {
			brdStat = ptttype.NBRD_FOLDER
		}
		userBoardInfoMap[each.BBoardID] = &apitypes.UserBoardInfo{
			Stat: brdStat,
		}
	}

	return boardSummaries, userBoardInfoMap, err
}

// DeserializeBoards
//
// each_b.Reason happens only with invalid permission.
func DeserializeBoardDetailsAndUpdateDB(boardDetails_b []*bbs.BoardDetail, updateNanoTS types.NanoTS) (boardDetails []*schema.BoardDetail, err error) {
	boardDetails = make([]*schema.BoardDetail, 0, len(boardDetails_b))
	for _, each_b := range boardDetails_b {
		if each_b.Reason != 0 {
			continue
		}
		each := schema.NewBoardDetail(each_b, updateNanoTS)

		boardDetails = append(boardDetails, each)
	}
	if len(boardDetails) == 0 {
		return nil, nil
	}

	err = schema.UpdateBoardDetails(boardDetails, updateNanoTS)
	if err != nil {
		return nil, err
	}

	return boardDetails, nil
}

func isBoardValidUser(boardID bbs.BBoardID, c *gin.Context) (isValid bool, statusCode int, err error) {
	if types.IS_ALL_GUEST {
		return true, 200, nil
	}

	var result_b *pttbbsapi.IsBoardValidUserResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.IS_BOARD_VALID_USER_R)
	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return false, statusCode, err
	}
	if !result_b.IsValid {
		return false, 403, ErrInvalidUser
	}

	return true, 200, nil
}

//nolint:unused
func isBoardSummariesValidUser(boardSummaries []*schema.BoardSummary, c *gin.Context) (validBoardSummaries []*schema.BoardSummary, err error) {
	boardIDs := make([]bbs.BBoardID, len(boardSummaries))
	for idx, each := range boardSummaries {
		boardIDs[idx] = each.BBoardID
	}

	var result_b *pttbbsapi.IsBoardsValidUserResult

	params := &pttbbsapi.IsBoardsValidUserParams{
		BoardIDs: boardIDs,
	}

	url := pttbbsapi.IS_BOARDS_VALID_USER_R
	statusCode, err := utils.BackendGet(c, url, params, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, err
	}

	validBoardSummaries = make([]*schema.BoardSummary, 0, len(boardSummaries))
	for _, each := range boardSummaries {
		isValid, ok := result_b.IsValid[each.BBoardID]
		if !ok || !isValid {
			continue
		}
		validBoardSummaries = append(validBoardSummaries, each)
	}

	return validBoardSummaries, nil
}

func getBoardSummaryMapFromBids(userID bbs.UUserID, bids []ptttype.Bid, c *gin.Context) (boardSummaryMap_db map[ptttype.Bid]*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, statusCode int, err error) {
	// backend get boards by bids
	theParams_b := &pttbbsapi.LoadBoardsByBidsParams{
		Bids: bids,
	}
	var result_b *pttbbsapi.LoadBoardsByBidsResult

	url := pttbbsapi.LOAD_BOARDS_BY_BIDS_R
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, nil, statusCode, err
	}

	// update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializeBoardsAndUpdateDB(userID, result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, nil, 500, err
	}

	boardSummaryMap_db = map[ptttype.Bid]*schema.BoardSummary{}
	for _, each := range boardSummaries_db {
		boardSummaryMap_db[each.Bid] = each
	}

	return boardSummaryMap_db, userBoardInfoMap, 200, nil
}

// https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L953
func checkUserReadBoard(userID bbs.UUserID, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, theList []*schema.BoardSummary) (newUserBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, err error) {
	checkBBoardIDMap := make(map[bbs.BBoardID]int)
	queryBBoardIDs := make([]bbs.BBoardID, 0, len(theList))
	for idx, each := range theList {
		if each == nil {
			continue
		}

		eachBoardInfo, ok := userBoardInfoMap[each.BBoardID]
		if (eachBoardInfo.Stat&ptttype.NBRD_LINE != 0) || (eachBoardInfo.Stat&ptttype.NBRD_FOLDER != 0) {
			continue
		}

		if ok && eachBoardInfo.Read {
			continue
		}

		if each.BrdAttr&(ptttype.BRD_GROUPBOARD|ptttype.BRD_SYMBOLIC) != 0 {
			continue
		}

		if each.Total == 0 {
			continue
		}

		// check with read-time
		checkBBoardIDMap[each.BBoardID] = idx
		queryBBoardIDs = append(queryBBoardIDs, each.BBoardID)
	}

	dbResults, err := schema.FindUserReadBoards(userID, queryBBoardIDs)
	if err != nil {
		return nil, err
	}

	// setup read in the list
	// no need to update db, because we don't read the newest yet.
	// the Read flag is set based on the existing db.UpdateNanoTS
	for _, each := range dbResults {
		eachBoardID := each.BBoardID
		eachReadNanoTS := each.ReadUpdateNanoTS

		eachBoardInfo, ok := userBoardInfoMap[eachBoardID]
		if !ok {
			continue
		}

		listIdx, ok := checkBBoardIDMap[eachBoardID]
		if !ok {
			continue
		}

		eachInTheList := theList[listIdx]
		eachLastPostNanoTS := eachInTheList.LastPostTime

		isRead := eachReadNanoTS > eachLastPostNanoTS
		eachBoardInfo.Read = isRead
	}

	return userBoardInfoMap, nil
}
