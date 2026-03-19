package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const CREATE_ARTICLE_R = "/board/:bid/article"

type CreateArticleParams struct {
	PostType string          `json:"class" form:"class" url:"class"`
	Title    string          `json:"title" form:"title" url:"title"`
	Content  [][]*types.Rune `json:"content" form:"content" url:"content"`
}

type CreateArticlePath struct {
	FBoardID apitypes.FBoardID `uri:"bid" binding:"required"`
}

type CreateArticleResult *apitypes.ArticleSummary

func CreateArticleWrapper(c *gin.Context) {
	params := &CreateArticleParams{}
	path := &CreateArticlePath{}
	LoginRequiredPathJSON(CreateArticle, params, path, c)
}

func CreateArticle(remoteAddr string, user *UserInfo, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*CreateArticleParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*CreateArticlePath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		logrus.Errorf("CreateArticle: unable to find boardID: boardID: %v e: %v", thePath.FBoardID, err)
		return nil, 500, err
	}

	err = CheckUserBoardPermPostable(user, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	theType := types.Utf8ToBig5(theParams.PostType)
	theTitle := types.Utf8ToBig5(theParams.Title)
	content := simplifyContent(theParams.Content)
	contentDBCS := dbcs.Utf8ToDBCS(content)

	// backend
	theParams_b := &pttbbsapi.CreateArticleParams{
		PostType: theType,
		Title:    theTitle,
		Content:  contentDBCS,
	}
	var result_b pttbbsapi.CreateArticleResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.CREATE_ARTICLE_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	// update to db
	theList_b := []*bbs.ArticleSummary{(*bbs.ArticleSummary)(result_b)}
	updateNanoTS := types.NowNanoTS()
	articleSummaries_db, _, err := deserializeArticlesAndUpdateDB(userID, boardID, theList_b, updateNanoTS, false)
	if err != nil {
		return nil, 500, err
	}

	articleSummary_db := articleSummaries_db[0]

	dummyContent := [][]*types.Rune{}

	setUserReadArticle(dummyContent, userID, articleSummary_db.ArticleID, updateNanoTS)

	articleSummary := apitypes.NewArticleSummaryFromWithRegex(articleSummary_db, userID)

	return CreateArticleResult(articleSummary), 200, nil
}
