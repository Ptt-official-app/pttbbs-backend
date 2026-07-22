package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

const LOAD_MAN_ARTICLES_R = "/board/:bid/manuals"

type LoadManArticlesParams struct {
	LevelIdx apitypes.FArticleID `json:"level_idx,omitempty" form:"level_idx,omitempty" url:"level_idx,omitempty"`
}

type LoadManArticlesPath struct {
	FBoardID apitypes.FBoardID `uri:"bid"`
}

type LoadManArticlesResult struct {
	List []*apitypes.ManArticleSummary `json:"list"`

	TokenUser bbs.UUserID `json:"tokenuser,omitempty"`
}

func LoadManArticlesWrapper(c *gin.Context) {
	params := &LoadManArticlesParams{}
	path := &LoadManArticlesPath{}
	LoginRequiredPathQuery(LoadManArticles, params, path, c)
}

func LoadManArticles(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoadManArticlesParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*LoadManArticlesPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	// check board permission
	_, err = CheckUserBoardPermReadable(user, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	level := theParams.LevelIdx.ToManArticleID()
	articleSummaries_db, err := schema.GetManArticleSummaries(boardID, level)
	if err != nil {
		return nil, 500, err
	}

	r := NewLoadManArticlesResult(articleSummaries_db, userID)

	return r, 200, nil
}

func NewLoadManArticlesResult(articleSummaries_db []*schema.ManArticleSummary, userID bbs.UUserID) (r *LoadManArticlesResult) {
	articleSummaries := make([]*apitypes.ManArticleSummary, 0, len(articleSummaries_db))

	for _, each_db := range articleSummaries_db {
		each := apitypes.NewManArticleSummary(each_db, "")
		articleSummaries = append(articleSummaries, each)
	}
	return &LoadManArticlesResult{
		List:      articleSummaries,
		TokenUser: userID,
	}
}
