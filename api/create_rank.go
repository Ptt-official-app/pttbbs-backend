package api

import (
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

func CreateRank(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*CreateRankParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}
	thePath, ok := path.(*CreateRankPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
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

	isValid, statusCode, err := isBoardValidUser(boardID, userID)
	if err != nil {
		return nil, statusCode, err
	}
	if !isValid {
		return nil, 403, ErrInvalidUser
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
