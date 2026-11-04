package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const LOAD_GENERAL_BOARDS_R = "/boards"

type LoadGeneralBoardsParams struct {
	Title     string `json:"title,omitempty" form:"title,omitempty" url:"title,omitempty"`
	Keyword   string `json:"keyword,omitempty" form:"keyword,omitempty" url:"keyword,omitempty"`
	StartIdx  string `json:"start_idx,omitempty" form:"start_idx,omitempty" url:"start_idx,omitempty"`
	Ascending bool   `json:"asc,omitzero"  form:"asc,omitempty" url:"asc,omitempty"`
	Max       int    `json:"limit,omitzero" form:"limit,omitempty" url:"limit,omitempty"`
}

type LoadGeneralBoardsResult struct {
	List    []*apitypes.BoardSummary `json:"list"`
	NextIdx string                   `json:"next_idx"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func NewLoadGeneralBoardsParams() *LoadGeneralBoardsParams {
	return &LoadGeneralBoardsParams{
		Ascending: DEFAULT_ASCENDING,
		Max:       DEFAULT_MAX_LIST,
	}
}

func LoadGeneralBoardsWrapper(c *gin.Context) {
	params := NewLoadGeneralBoardsParams()
	LoginRequiredQuery(LoadGeneralBoards, params, c)
}

func LoadGeneralBoards(remoteAddr string, user *UserInfo, params any, c *gin.Context) (result any, statusCode int, err error) {
	return loadGeneralBoardsCore(remoteAddr, user, params, c, pttbbsapi.LOAD_GENERAL_BOARDS_R)
}

func loadGeneralBoardsCore(remoteAddr string, user *UserInfo, params any, c *gin.Context, url string) (result any, statusCode int, err error) {
	theParams, ok := params.(*LoadGeneralBoardsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	// backend load-general-baords
	theParams_b := &pttbbsapi.LoadGeneralBoardsParams{
		StartIdx: theParams.StartIdx,
		Title:    types.Utf8ToBig5(theParams.Title),
		Keyword:  types.Utf8ToBig5(theParams.Keyword),
		NBoards:  theParams.Max,
		Asc:      theParams.Ascending,
	}
	var result_b *pttbbsapi.LoadGeneralBoardsResult

	statusCode, err = utils.BackendGet(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	return postLoadBoards(user.UserID, result_b, url, c)
}

func postLoadBoards(userID bbs.UUserID, result_b *pttbbsapi.LoadGeneralBoardsResult, url string, c *gin.Context) (result *LoadGeneralBoardsResult, statusCode int, err error) {
	// update to db
	updateNanoTS := types.NowNanoTS()
	boardSummaries_db, userBoardInfoMap, err := deserializeBoardsAndUpdateDB(userID, result_b.Boards, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	// check isRead
	userBoardInfoMap, err = checkUserReadBoard(userID, userBoardInfoMap, boardSummaries_db)
	if err != nil {
		return nil, 500, err
	}

	userBoardInfoMap, err = checkUserFavBoard(userID, userBoardInfoMap, boardSummaries_db, c)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadGeneralBoardsResult(boardSummaries_db, userBoardInfoMap, result_b.NextIdx, url, userID)

	return r, 200, nil
}

func NewLoadGeneralBoardsResult(boardSummaries_db []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo, nextIdx string, url string, userID bbs.UUserID) *LoadGeneralBoardsResult {
	theList := make([]*apitypes.BoardSummary, len(boardSummaries_db))

	isByClass := url == pttbbsapi.LOAD_GENERAL_BOARDS_BY_CLASS_R
	for i, each_db := range boardSummaries_db {
		idx := each_db.IdxByName
		if isByClass {
			idx = each_db.IdxByClass
		}

		userBoardInfo, ok := userBoardInfoMap[each_db.BBoardID]
		if !ok {
			continue
		}

		each := apitypes.NewBoardSummary(each_db, idx, userBoardInfo, "")
		theList[i] = each
	}

	return &LoadGeneralBoardsResult{
		List:    theList,
		NextIdx: nextIdx,

		TokenUser: userID,
	}
}
