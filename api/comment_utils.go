package api

import (
	"bytes"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/cmsys"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/queue"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

func tryUpdateFirstComments(
	firstComments []*schema.Comment,
	firstCommentsMD5 string,
	firstCommentsUpdateNanoTS types.NanoTS,
	articleDetailSummary *schema.ArticleDetailSummary) (
	err error,
) {
	if firstCommentsMD5 == articleDetailSummary.FirstCommentsMD5 {
		return nil
	}

	err = schema.UpdateComments(firstComments, firstCommentsUpdateNanoTS)
	if err != nil {
		return err
	}

	articleFirstComments := &schema.ArticleFirstComments{
		BBoardID:                  articleDetailSummary.BBoardID,
		ArticleID:                 articleDetailSummary.ArticleID,
		FirstCommentsMD5:          firstCommentsMD5,
		FirstCommentsUpdateNanoTS: firstCommentsUpdateNanoTS,
	}
	err = schema.UpdateArticleFirstComments(articleFirstComments)
	if err != nil {
		return err
	}

	// remove old first comments.
	return nil
}

func tryUpdateComments(comments []*schema.Comment, updateNanoTS types.NanoTS) (err error) {
	err = schema.UpdateComments(comments, updateNanoTS)
	if err != nil {
		return err
	}

	// remove old first comments.
	return nil
}

func commentDBCSsToCommentsDBCS(commentDBCSs []*schema.CommentDBCS) (commentsDBCS [][]byte) {
	splitCommentDBCSs := make([][][]byte, len(commentDBCSs))
	lenCommentsDBCS := 0
	for idx, each := range commentDBCSs {
		splitCommentDBCS := bytes.Split(each.DBCS, []byte{'\n'})
		lenCommentsDBCS += len(splitCommentDBCS)
		splitCommentDBCSs[idx] = splitCommentDBCS
	}

	commentsDBCS = make([][]byte, 0, lenCommentsDBCS)
	for _, each := range splitCommentDBCSs {
		commentsDBCS = append(commentsDBCS, each...)
	}

	return commentsDBCS
}

func postUpdateComments(userID bbs.UUserID, remoteAddr string, boardID bbs.BBoardID, articleID bbs.ArticleID, oldContent [][]*types.Rune, oldContentPrefix [][]*types.Rune, oldSignatureDBCS []byte, articleDetailSummary *schema.ArticleDetailSummary, oldSZ int, oldsum cmsys.Fnv64_t, c *gin.Context) (statusCode int, err error) {
	allContentDBCS, err := editArticleCompileContent(boardID, articleID, oldContent, oldContentPrefix, oldSignatureDBCS)
	if err != nil {
		return 500, err
	}

	// 4. do lock. if failed, return the data in db.
	lockKey := ArticleLockKey(boardID, articleID)
	lockToken, err := schema.TryLock(lockKey, ARTICLE_LOCK_TS_DURATION)
	if err != nil {
		return 500, err
	}
	defer func() { _ = schema.Unlock(lockKey, lockToken) }()

	// edit article
	theParams_b := &pttbbsapi.EditArticleParams{
		OldSZ:    oldSZ,
		OldSum:   oldsum,
		PostType: nil,
		Title:    nil,
		Content:  allContentDBCS,
	}
	var result_b *pttbbsapi.EditArticleResult

	urlMap := map[string]string{
		"bid": string(boardID),
		"aid": string(articleID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.EDIT_ARTICLE_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return statusCode, err
	}

	updateNanoTS := types.NowNanoTS()
	content, contentPrefix, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, commentsDBCS := dbcs.ParseContent(result_b.Content, "")

	err = UpdateArticleContentInfo(boardID, articleID, content, contentPrefix, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, updateNanoTS)
	if err != nil {
		return 500, err
	}

	contentMTime := types.Time4ToNanoTS(result_b.MTime)
	_ = schema.UpdateArticleContentMTime(boardID, articleID, contentMTime)

	commentQueue := &queue.CommentQueue{
		BBoardID:          boardID,
		ArticleID:         articleID,
		OwnerID:           articleDetailSummary.Owner,
		CommentDBCS:       commentsDBCS,
		ArticleCreateTime: articleDetailSummary.CreateTime,
		ArticleMTime:      contentMTime,
		UpdateNanoTS:      updateNanoTS,
	}

	_ = queue.ProcessCommentQueue(commentQueue)

	return 200, nil
}
