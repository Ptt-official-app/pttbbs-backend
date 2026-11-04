package api

import (
	"maps"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

const LOAD_FAVORITE_BOARDS_R = "/user/:user_id/favorites"

type LoadFavoriteBoardsParams struct {
	LevelIdx  schema.LevelIdx `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
	StartIdx  string          `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Ascending bool            `json:"asc,omitzero"  form:"asc,omitempty" url:"asc,omitempty"`
	Max       int             `json:"limit,omitzero" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadFavoriteBoardsPath struct {
	UserID bbs.UUserID `uri:"user_id"`
}

type LoadFavoriteBoardsResult struct {
	List    []*apitypes.BoardSummary `json:"list"`
	NextIdx string                   `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func NewLoadFavoriteBoardsParams() *LoadFavoriteBoardsParams {
	return &LoadFavoriteBoardsParams{
		Ascending: DEFAULT_ASCENDING,
		Max:       DEFAULT_MAX_LIST,
	}
}

func LoadFavoriteBoardsWrapper(c *gin.Context) {
	params := NewLoadFavoriteBoardsParams()
	path := &LoadFavoriteBoardsPath{}
	LoginRequiredPathQuery(LoadFavoriteBoards, params, path, c)
}

func LoadFavoriteBoards(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*LoadFavoriteBoardsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadFavoriteBoardsPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	if userID != thePath.UserID {
		return nil, 403, ErrInvalidUser
	}

	userFavorites_db, nextIdx, statusCode, err := tryGetUserFavorites(thePath.UserID, theParams.LevelIdx, theParams.StartIdx, theParams.Ascending, theParams.Max, c)
	if err != nil {
		return nil, statusCode, err
	}

	boardSummaryMap_db, userBoardInfoMap, statusCode, err := tryGetBoardSummaryMapFromUserFavorites(thePath.UserID, userFavorites_db, c)
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

	for each := range maps.Values(userBoardInfoMap) {
		each.Fav = true
	}

	r := NewLoadFavoriteBoardsResult(userID, userFavorites_db, boardSummaryMap_db, userBoardInfoMap, nextIdx)

	return r, 200, nil
}

func NewLoadFavoriteBoardsResult(userID bbs.UUserID, userFavorites_db []*schema.UserFavorites, boardSummaryMap_db map[ptttype.Bid]*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, nextIdx string) (result *LoadFavoriteBoardsResult) {
	theList := make([]*apitypes.BoardSummary, len(userFavorites_db))
	for idx, each := range userFavorites_db {
		if each.TheType != pttbbsfav.FAVT_BOARD {
			theList[idx] = apitypes.NewBoardSummaryFromUserFavorites(userID, each, nil, nil)
			theList[idx].TokenUser = ""
			continue
		}

		boardSummary_db, ok := boardSummaryMap_db[ptttype.Bid(each.TheID)]
		if !ok {
			continue
		}

		userBoardInfo, ok := userBoardInfoMap[boardSummary_db.BBoardID]
		if !ok {
			continue
		}
		theList[idx] = apitypes.NewBoardSummaryFromUserFavorites(userID, each, boardSummary_db, userBoardInfo)
		theList[idx].TokenUser = ""
	}

	result = &LoadFavoriteBoardsResult{
		List:    theList,
		NextIdx: nextIdx,

		TokenUser: userID,
	}

	return result
}
