package api

import (
	"bytes"
	"errors"
	"maps"
	"strconv"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/fav"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

func tryGetUserFavorites(
	userID bbs.UUserID,
	levelIdx schema.LevelIdx,
	startIdxStr string,
	ascending bool,
	limit int,
	c *gin.Context) (
	userFavorites []*schema.UserFavorites,
	nextIdxStr string,
	statusCode int,
	err error,
) {
	startIdx := 0
	if len(startIdxStr) > 0 {
		startIdx, err = strconv.Atoi(startIdxStr)
		if err != nil {
			return nil, "", 400, ErrInvalidParams
		}
	}

	doubleBufferIdx, mtime, statusCode, err := tryGetUserFavoritesCore(userID, c)
	if err != nil {
		return nil, "", statusCode, err
	}

	// get db
	userFavorites, nextIdx, err := getUserFavoritesFromDB(userID, doubleBufferIdx, mtime, levelIdx, startIdx, ascending, limit)
	if err != nil {
		return nil, "", 500, err
	}

	nextIdxStr = ""
	if nextIdx >= 0 {
		nextIdxStr = strconv.Itoa(nextIdx)
	}

	return userFavorites, nextIdxStr, 200, nil
}

func tryGetUserFavoritesCore(userID bbs.UUserID, c *gin.Context) (newDoubleBufferIdx int, newMTime types.NanoTS, statusCode int, err error) {
	// user-favorites-meta
	userFavoritesMeta, err := schema.GetUserFavoritesMeta(userID)
	if err != nil {
		return -1, 0, 500, err
	}

	dbMTime := types.NanoTS(0)
	newDoubleBufferIdx = 0
	if userFavoritesMeta != nil {
		dbMTime = userFavoritesMeta.MTime
		newDoubleBufferIdx = (userFavoritesMeta.DoubleBufferIdx + 1) % schema.N_USER_FAVORITES_DOUBLE_BUFFER
	}

	// backend get user-favorites
	theParams_b := &pttbbsapi.GetFavoritesParams{
		RetrieveTS: dbMTime.ToTime4(),
	}
	var result_b *pttbbsapi.GetFavoritesResult

	urlMap := map[string]string{
		"uid": string(userID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.GET_FAV_R)

	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return -1, 0, statusCode, err
	}

	// check mtime
	backendMTime := types.Time4ToNanoTS(result_b.MTime)
	if backendMTime <= dbMTime {
		return userFavoritesMeta.DoubleBufferIdx, dbMTime, 200, nil
	}

	updateNanoTS := types.NowNanoTS()
	err = deserializeUserFavoritesAndUpdateDB(userID, backendMTime, result_b.Content, updateNanoTS, newDoubleBufferIdx)
	if err != nil {
		return -1, 0, 500, err
	}

	return newDoubleBufferIdx, backendMTime, 200, nil
}

func tryGetAllUserFavorites(userID bbs.UUserID, c *gin.Context) (userFavoritesMeta *schema.UserFavoritesMeta, userFavorites []*schema.UserFavorites, err error) {
	_, _, _, err = tryGetUserFavoritesCore(userID, c)
	if err != nil {
		return nil, nil, err
	}

	return getAllUserFavoritesFromDB(userID)
}

func getAllUserFavoritesFromDB(userID bbs.UUserID) (userFavoritesMeta *schema.UserFavoritesMeta, userFavorites []*schema.UserFavorites, err error) {
	// user-favorites-meta
	userFavoritesMeta, err = schema.GetUserFavoritesMeta(userID)
	if err != nil {
		return nil, nil, err
	}
	if userFavoritesMeta == nil {
		userFavoritesMeta = &schema.UserFavoritesMeta{
			UserID: userID,
		}
		return userFavoritesMeta, []*schema.UserFavorites{}, nil
	}

	// get db
	userFavorites, err = schema.GetAllUserFavorites(userID, userFavoritesMeta.DoubleBufferIdx, userFavoritesMeta.MTime)
	if err != nil {
		return nil, nil, err
	}

	return userFavoritesMeta, userFavorites, nil
}

func deserializeUserFavoritesAndUpdateDB(userID bbs.UUserID, mTime types.NanoTS, contentWithVersion []byte, updateNanoTS types.NanoTS, doubleBufferIdx int) (err error) {
	if len(contentWithVersion) < 2 {
		return ErrInvalidFav
	}
	content := contentWithVersion[2:]

	file := bytes.NewReader(content)

	f, err := fav.ReadFavrec(file, nil, nil, 0)
	if err != nil {
		return err
	}

	meta, userFavorites := schema.FavToUserFavorites(f, userID, doubleBufferIdx, updateNanoTS, mTime)

	err = schema.UpdateUserFavorites(userID, doubleBufferIdx, userFavorites, mTime, updateNanoTS)
	if err != nil {
		return err
	}

	// ok if data is out-dated.
	err = schema.UpdateUserFavoritesMeta(meta)
	if errors.Is(err, schema.ErrNoMatch) {
		err = nil
	}
	if err != nil {
		return err
	}

	return nil
}

func getUserFavoritesFromDB(userID bbs.UUserID, doubleBufferIdx int, mtime types.NanoTS, levelIdx schema.LevelIdx, startIdx int, ascending bool, limit int) (userFavorites []*schema.UserFavorites, nextIdx int, err error) {
	userFavorites, err = schema.GetUserFavorites(userID, doubleBufferIdx, levelIdx, startIdx, ascending, limit+1, mtime)
	if err != nil {
		return nil, 0, err
	}
	if len(userFavorites) <= limit {
		nextIdx = -1
	} else {
		nextIdx = userFavorites[limit].Idx
		userFavorites = userFavorites[:limit]
	}

	return userFavorites, nextIdx, nil
}

func tryGetBoardSummaryMapFromUserFavorites(userID bbs.UUserID, userFavorites_db []*schema.UserFavorites, c *gin.Context) (boardSummaryMap_db map[ptttype.Bid]*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, statusCode int, err error) {
	bids := bidsInUserFavorites(userFavorites_db)

	return getBoardSummaryMapFromBids(userID, bids, c)
}

func bidsInUserFavorites(userFavorites []*schema.UserFavorites) (bids []ptttype.Bid) {
	bids = make([]ptttype.Bid, 0, len(userFavorites))
	for _, each := range userFavorites {
		if each.TheType != pttbbsfav.FAVT_BOARD {
			continue
		}

		eachBid := ptttype.Bid(each.TheID)

		bids = append(bids, eachBid)
	}

	return bids
}

func tryWriteFav(theFav *fav.Fav, remoteAddr string, userID bbs.UUserID, c *gin.Context) (statusCode int, err error) {
	theBytes := make([]byte, 0, MAX_USER_FAVORITES_BUF_SIZE)

	buf := bytes.NewBuffer(theBytes)

	err = theFav.WriteFavrec(buf)
	if err != nil {
		return 500, err
	}

	content := buf.Bytes()

	// backend get user-favorites
	theParams_b := &pttbbsapi.WriteFavoritesParams{
		Content: content,
	}
	var result_b *pttbbsapi.WriteFavoritesResult

	urlMap := map[string]string{
		"uid": string(userID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.WRITE_FAV_R)

	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil {
		return statusCode, err
	}

	return 200, nil
}

func checkUserFavBoard(userID bbs.UUserID, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, boardSummaries []*schema.BoardSummary, c *gin.Context) (newUserBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, err error) {
	doubleBufferIdx, mtime, _, err := tryGetUserFavoritesCore(userID, c)
	if err != nil {
		return nil, err
	}

	pttbids := make([]ptttype.Bid, len(boardSummaries))
	pttbidMap := make(map[ptttype.Bid]bbs.BBoardID)
	for idx, each := range boardSummaries {
		pttbids[idx] = each.Bid
		pttbidMap[each.Bid] = each.BBoardID
	}

	userFavoriteIDs, err := schema.GetUserFavoriteIDsByPttbids(userID, doubleBufferIdx, pttbids, mtime)
	if err != nil {
		return nil, err
	}
	for _, each := range userFavoriteIDs {
		bid, ok := pttbidMap[ptttype.Bid(each.TheID)]
		if !ok {
			continue
		}
		userBoardInfoMap[bid].Fav = true
	}

	return userBoardInfoMap, nil
}

func postAddFavorite(userID bbs.UUserID, rootFav *fav.Fav, remoteAddr string, levelIdx schema.LevelIdx, startIdx int, c *gin.Context, isBoard bool) (ret AddFavoriteResult, statusCode int, err error) {
	statusCode, err = tryWriteFav(rootFav, remoteAddr, userID, c)
	if err != nil {
		return nil, statusCode, err
	}

	startIdxStr := strconv.Itoa(startIdx)

	newUserFavorites, _, statusCode, err := tryGetUserFavorites(userID, levelIdx, startIdxStr, true, 1, c)
	if err != nil {
		return nil, statusCode, err
	}

	if len(newUserFavorites) != 1 {
		return nil, 500, ErrInvalidFav
	}

	newUserFavorite := newUserFavorites[0]
	if !isBoard {
		summary := apitypes.NewBoardSummaryFromUserFavorites(userID, newUserFavorite, nil, nil)
		return AddFavoriteResult(summary), 200, nil
	}

	boardSummaryMap_db, userBoardInfoMap, statusCode, err := tryGetBoardSummaryMapFromUserFavorites(userID, newUserFavorites, c)
	if err != nil {
		return nil, statusCode, err
	}

	boardSummaries_db := make([]*schema.BoardSummary, 0, len(boardSummaryMap_db))
	for each := range maps.Values(boardSummaryMap_db) {
		boardSummaries_db = append(boardSummaries_db, each)
	}

	userBoardInfoMap, err = checkUserReadBoard(userID, userBoardInfoMap, boardSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	boardSummary_db, ok := boardSummaryMap_db[ptttype.Bid(newUserFavorite.TheID)]
	if !ok {
		return nil, 500, ErrInvalidFav
	}
	userBoardInfo, ok := userBoardInfoMap[boardSummary_db.BBoardID]
	if !ok {
		return nil, 500, ErrInvalidFav
	}

	summary := apitypes.NewBoardSummaryFromUserFavorites(userID, newUserFavorite, boardSummary_db, userBoardInfo)

	summary.TokenUser = userID

	return AddFavoriteResult(summary), 200, nil
}
