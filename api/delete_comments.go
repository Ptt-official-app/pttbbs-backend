package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const DELETE_COMMENTS_R = "/board/:bid/article/:aid/deletecomments"

type DeleteCommentsParams struct {
	TheList []*apitypes.DeleteCommentParams `json:"list" form:"list" url:"list"`
}

type DeleteCommentsPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type DeleteCommentsResult struct {
	Success   bool        `json:"success"`
	TokenUser bbs.UUserID `json:"tokenuser"`
}

func DeleteCommentsWrapper(c *gin.Context) {
	params := &DeleteCommentsParams{}
	path := &DeleteCommentsPath{}
	LoginRequiredPathJSON(DeleteComments, params, path, c)
}

func DeleteComments(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*DeleteCommentsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*DeleteCommentsPath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	userID := user.UserID
	boardID, err := toBoardID(thePath.FBoardID, remoteAddr, userID, c)
	if err != nil {
		return nil, 500, err
	}
	articleID := thePath.FArticleID.ToArticleID()

	oldContent, oldContentPrefix, oldSignatureDBCS, articleDetailSummary_db, oldSZ, oldsum, statusCode, err := editArticleGetArticleContentInfo(userID, boardID, articleID, c, true)
	if err != nil {
		return nil, statusCode, err
	}

	err = deleteCommentsUpdateComments(userID, remoteAddr, boardID, articleID, theParams.TheList)
	if err != nil {
		return nil, 500, err
	}

	statusCode, err = postUpdateComments(userID, remoteAddr, boardID, articleID, oldContent, oldContentPrefix, oldSignatureDBCS, articleDetailSummary_db, oldSZ, oldsum, c)
	if err != nil {
		return nil, statusCode, err
	}

	result = &DeleteCommentsResult{
		Success:   true,
		TokenUser: userID,
	}
	return result, 200, nil
}

func deleteCommentsUpdateComments(userID bbs.UUserID, remoteAddr string, boardID bbs.BBoardID, articleID bbs.ArticleID, deletes []*apitypes.DeleteCommentParams) (err error) {
	commentIDs := make([]types.CommentID, len(deletes))
	for idx, each := range deletes {
		commentIDs[idx] = each.CommentID
	}

	updateNanoTS := types.NowNanoTS()
	commentMap, err := schema.GetCommentMapByCommentIDs(boardID, articleID, commentIDs)
	if err != nil {
		return nil
	}

	comments := make([]*schema.Comment, 0, len(deletes))
	for _, each := range deletes {
		comment := commentMap[each.CommentID]
		newComment := each.ToComment(comment, userID, remoteAddr, updateNanoTS)
		comments = append(comments, newComment)
	}

	return schema.UpdateComments(comments, updateNanoTS)
}
