package api

import (
	"bytes"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/Ptt-official-app/pttbbs-backend/utils"
	"github.com/gin-gonic/gin"
)

const EDIT_ARTICLE_R = "/board/:bid/article/:aid/edit"

type EditArticleParams struct {
	PostType string          `json:"class" form:"class" url:"class"`
	Title    string          `json:"title" form:"title" url:"title"`
	Content  [][]*types.Rune `json:"content" form:"content" url:"content"`
}

type EditArticlePath struct {
	FBoardID   apitypes.FBoardID   `uri:"bid"`
	FArticleID apitypes.FArticleID `uri:"aid"`
}

type EditArticleResult struct {
	MTime         types.Time8     `json:"modified"` //
	Content       [][]*types.Rune `json:"content"`
	ContentPrefix [][]*types.Rune `json:"prefix"`
	Title         string          `json:"title"` //
	Class         string          `json:"class"` // can be: R: 轉, [class]
	TokenUser     bbs.UUserID     `json:"tokenuser"`
}

func EditArticleDetailWrapper(c *gin.Context) {
	params := &EditArticleParams{}
	path := &EditArticlePath{}
	LoginRequiredPathJSON(EditArticleDetail, params, path, c)
}

func EditArticleDetail(remoteAddr string, user *UserInfo, params any, path any, c *gin.Context) (result any, statusCode int, err error) {
	theParams, ok := params.(*EditArticleParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*EditArticlePath)
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
	err = CheckUserArticlePermEditable(user, boardID, articleID, true, c)
	if err != nil {
		return nil, 403, err
	}

	_, oldContentPrefix, oldSignatureDBCS, _, oldSZ, oldsum, statusCode, err := editArticleGetArticleContentInfo(userID, boardID, articleID, c, false)
	if err != nil {
		return nil, statusCode, err
	}

	allContentDBCS, err := editArticleCompileContent(boardID, articleID, theParams.Content, oldContentPrefix, oldSignatureDBCS)
	if err != nil {
		return nil, 500, err
	}

	// 4. do lock. if failed, return the data in db.
	lockKey := ArticleLockKey(boardID, articleID)
	err = schema.TryLock(lockKey, ARTICLE_LOCK_TS_DURATION)
	if err != nil {
		return nil, 500, err
	}
	defer func() { _ = schema.Unlock(lockKey) }()

	// edit article
	var theType []byte
	var theTitle []byte
	if theParams.Title != "" {
		theType = types.Utf8ToBig5(theParams.PostType)
		theTitle = types.Utf8ToBig5(theParams.Title)
	}

	theParams_b := &pttbbsapi.EditArticleParams{
		OldSZ:    oldSZ,
		OldSum:   oldsum,
		PostType: theType,
		Title:    theTitle,
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
		return nil, statusCode, err
	}

	updateNanoTS := types.NowNanoTS()
	content, contentPrefix, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, _ := dbcs.ParseContent(result_b.Content, "")

	// update article
	// we need update-article-content be the 1st to upload,
	// because it's possible that there is no first-comments.
	// only article-content is guaranteed.

	err = UpdateArticleContentInfo(boardID, articleID, content, contentPrefix, contentMD5, ip, host, bbs, signatureMD5, signatureDBCS, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	contentMTime := types.Time4ToNanoTS(result_b.MTime)
	_ = schema.UpdateArticleContentMTime(boardID, articleID, contentMTime)

	// update article-edit-summary
	articleEditSummary := schema.NewArticleEditSummary(contentMTime, result_b.RealTitle, result_b.FullTitle, result_b.Class, updateNanoTS)
	err = schema.UpdateArticleEditSummary(boardID, articleID, articleEditSummary, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	result = &EditArticleResult{
		MTime:     contentMTime.ToTime8(),
		Content:   theParams.Content,
		Title:     articleEditSummary.Title,
		Class:     articleEditSummary.Class,
		TokenUser: userID,
	}
	return result, 200, nil
}

func editArticleCompileContent(boardID bbs.BBoardID, articleID bbs.ArticleID, content [][]*types.Rune, prefix [][]*types.Rune, signatureDBCS []byte) (allContentDBCS [][]byte, err error) {
	commentDBCSs, err := schema.GetAllCommentDBCSs(boardID, articleID)
	if err != nil {
		return nil, err
	}

	commentsDBCS := commentDBCSsToCommentsDBCS(commentDBCSs)

	prefixDBCS := dbcs.Utf8ToDBCS(prefix)

	contentDBCS := dbcs.Utf8ToDBCS(content)

	signaturesDBCS := bytes.Split(signatureDBCS, []byte{'\n'})

	lenAllContentDBCS := len(prefixDBCS) + len(contentDBCS) + len(commentsDBCS) + len(signaturesDBCS)
	allContentDBCS = make([][]byte, 0, lenAllContentDBCS)
	allContentDBCS = append(allContentDBCS, prefixDBCS...)
	allContentDBCS = append(allContentDBCS, contentDBCS...)
	allContentDBCS = append(allContentDBCS, signaturesDBCS...)
	allContentDBCS = append(allContentDBCS, commentsDBCS...)

	return allContentDBCS, nil
}
