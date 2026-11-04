package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const CREATE_COMMENT_R = "/board/:bid/article/:aid/comment"

type CreateCommentParams struct {
	CommentType ptttype.CommentType `json:"type" form:"type" url:"type"`
	Content     string              `json:"content" form:"content" url:"content"`
}

type CreateCommentPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid" binding:"required"`
	FArticleID apitypes.FArticleID `uri:"aid" binding:"required"`
}

type CreateCommentResult *apitypes.Comment

func CreateCommentWrapper(c *gin.Context) {
	params := &CreateCommentParams{}
	path := &CreateCommentPath{}
	LoginRequiredPathJSON(CreateComment, params, path, c)
}

func CreateComment(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*CreateCommentParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*CreateCommentPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}
	articleID := thePath.FArticleID.ToArticleID()

	err = CheckUserBoardPermPostable(user, boardID, c)
	if err != nil {
		return nil, 403, err
	}

	// content-dbcs
	contentDBCS := types.Utf8ToBig5(theParams.Content)

	// backend
	theParams_b := &pttbbsapi.CreateCommentParams{
		CommentType: theParams.CommentType,
		Content:     contentDBCS,
	}
	var result_b *pttbbsapi.CreateCommentResult

	urlMap := map[string]string{
		"bid": string(boardID),
		"aid": string(articleID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.CREATE_COMMENT_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	dbComments := dbcs.ParseComments(userID, result_b.Content, result_b.Content)
	if len(dbComments) == 0 {
		return nil, 500, ErrInvalidParams
	}

	dbComment := dbComments[0]
	dbComment.CreateTime = types.Time4ToNanoTS(result_b.MTime)
	dbComment.BBoardID = boardID
	dbComment.ArticleID = articleID
	updateNanoTS := types.NowNanoTS()
	dbComment.SetSortTime(updateNanoTS)
	err = tryUpdateComments(dbComments, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	apiComment := apitypes.NewComment(dbComment, userID)

	return CreateCommentResult(apiComment), 200, nil
}
