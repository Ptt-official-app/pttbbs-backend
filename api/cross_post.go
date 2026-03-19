package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const CROSS_POST_R = "/board/:bid/article/:aid/crosspost"

type CrossPostParams struct {
	XBoardID apitypes.FBoardID `json:"xbid" form:"xbid" url:"xbid"`
}

type CrossPostPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid" binding:"required"`
	FArticleID apitypes.FArticleID `uri:"aid" binding:"required"`
}

type CrossPostResult struct {
	Article *apitypes.ArticleSummary `json:"article"`
	Comment *apitypes.Comment        `json:"comment"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func CrossPostWrapper(c *gin.Context) {
	params := &CrossPostParams{}
	path := &CrossPostPath{}
	LoginRequiredPathJSON(CrossPost, params, path, c)
}

func CrossPost(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*CrossPostParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*CrossPostPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	xBoardID, err := toBoardID(theParams.XBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}

	articleID := thePath.FArticleID.ToArticleID()

	// backend
	theParams_b := &pttbbsapi.CrossPostParams{
		XBoardID: xBoardID,
	}
	var result_b *pttbbsapi.CrossPostResult

	urlMap := map[string]string{
		"bid": string(boardID),
		"aid": string(articleID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.CROSS_POST_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	// set article
	theList_b := []*bbs.ArticleSummary{result_b.ArticleSummary}
	updateNanoTS := types.NowNanoTS()
	articleSummaries_db, _, err := deserializeArticlesAndUpdateDB(userID, xBoardID, theList_b, updateNanoTS, false)
	if err != nil {
		return nil, 500, err
	}

	articleSummary_db := articleSummaries_db[0]
	dummyContent := [][]*types.Rune{}

	setUserReadArticle(dummyContent, userID, articleSummary_db.ArticleID, updateNanoTS)

	articleSummary := apitypes.NewArticleSummaryFromWithRegex(articleSummary_db, "")

	// comment
	dbComments := dbcs.ParseComments(userID, result_b.Comment, result_b.Comment)
	if len(dbComments) == 0 {
		return nil, 500, ErrInvalidParams
	}

	dbComment := dbComments[0]
	dbComment.CreateTime = types.Time4ToNanoTS(result_b.CommentMTime)
	dbComment.BBoardID = boardID
	dbComment.ArticleID = articleID
	dbComment.SetSortTime(updateNanoTS)
	err = tryUpdateComments(dbComments, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	apiComment := apitypes.NewComment(dbComment, "")

	return &CrossPostResult{Article: articleSummary, Comment: apiComment, TokenUser: userID}, 200, nil
}
