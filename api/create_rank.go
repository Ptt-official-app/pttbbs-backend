package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const CREATE_RANK_R = "/board/:bid/article/:aid/rank"

type CreateRankParams struct {
	Rank int `json:"rank" form:"rank" url:"rank"`
}

type CreateRankPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid" binding:"required"`
	FArticleID apitypes.FArticleID `uri:"aid" binding:"required"`
}

type CreateRankResult struct {
	Rank int `json:"rank"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func CreateRankWrapper(c *gin.Context) {
	params := &CreateRankParams{}
	path := &CreateRankPath{}
	LoginRequiredPathJSON(CreateRank, params, path, c)
}

func CreateRank(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*CreateRankParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}
	thePath, ok := path.(*CreateRankPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	// rank is -1 (dislike), 0 (cancel) or 1 (like).
	if theParams.Rank < -1 || theParams.Rank > 1 {
		return nil, 400, ErrInvalidParams
	}

	userID := user.UserID
	// all the guests share the same guest-record, so ranking requires login.
	if userID == bbs.UUserID(pttbbsapi.GUEST) {
		return nil, 403, ErrInvalidUser
	}

	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, boardIDErrToStatusCode(err), err
	}
	articleID := thePath.FArticleID.ToArticleID()

	// check permission
	articleSummary, err := schema.GetArticleSummary(boardID, articleID)
	if err != nil {
		return nil, 500, err
	}
	if articleSummary == nil {
		return nil, 400, ErrNoArticle
	}

	if articleSummary.Owner == userID {
		return nil, 403, ErrInvalidUser
	}

	_, err = CheckUserBoardPermReadable(user, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	// update rank
	updateNanoTS := types.NowNanoTS()
	origRank, err := schema.UpdateRank(boardID, articleID, userID, theParams.Rank, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	newRank, err := schema.UpdateArticleRank(boardID, articleID, theParams.Rank-origRank, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	return &CreateRankResult{Rank: newRank, TokenUser: userID}, 200, nil
}
