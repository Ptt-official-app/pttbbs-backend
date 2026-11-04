package api

import (
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

const REPLY_COMMENTS_R = "/board/:bid/article/:aid/replycomments"

type ReplyCommentsParams struct {
	TheList []*apitypes.ReplyCommentParams `json:"list" form:"list" url:"list"`
}

type ReplyCommentsPath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type ReplyCommentsResult struct {
	Success bool `json:"success"`

	TokenUser bbs.UUserID `json:"tokenuser"`
}

func ReplyCommentsWrapper(c *gin.Context) {
	params := &ReplyCommentsParams{}
	path := &ReplyCommentsPath{}
	LoginRequiredPathJSON(ReplyComments, params, path, c)
}

func ReplyComments(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*ReplyCommentsParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*ReplyCommentsPath)
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
	err = CheckUserArticlePermReadable(user, boardID, articleID, true, c)
	if err != nil {
		return nil, 403, err
	}

	oldContent, oldContentPrefix, oldSignatureDBCS, articleDetailSummary_db, oldSZ, oldsum, statusCode, err := editArticleGetArticleContentInfo(userID, boardID, articleID, c, true)
	if err != nil {
		return nil, statusCode, err
	}

	err = replyCommentsUpdateReplies(userID, remoteAddr, boardID, articleID, theParams.TheList)
	if err != nil {
		return nil, 500, err
	}

	statusCode, err = postUpdateComments(userID, remoteAddr, boardID, articleID, oldContent, oldContentPrefix, oldSignatureDBCS, articleDetailSummary_db, oldSZ, oldsum, c)
	if err != nil {
		return nil, statusCode, err
	}

	result = &ReplyCommentsResult{
		Success: true,

		TokenUser: userID,
	}
	return result, 200, nil
}

func replyCommentsUpdateReplies(userID bbs.UUserID, remoteAddr string, boardID bbs.BBoardID, articleID bbs.ArticleID, replies []*apitypes.ReplyCommentParams) (err error) {
	updateNanoTS := types.NowNanoTS()
	replyComments := make([]*schema.Comment, len(replies))

	commentIDs := make([]types.CommentID, len(replies))
	for idx, each := range replies {
		commentIDs[idx] = each.CommentID
	}
	sortTimeMap, err := schema.GetCommentSortTimeMapByCommentIDs(boardID, articleID, commentIDs)
	if err != nil {
		return err
	}

	for idx, each := range replies {
		sortTime := sortTimeMap[each.CommentID]
		replyComments[idx] = each.ToComment(userID, remoteAddr, boardID, articleID, sortTime, updateNanoTS)
	}

	return schema.UpdateComments(replyComments, updateNanoTS)
}
