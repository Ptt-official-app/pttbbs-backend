package api

import (
	"context"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/dbcs"
	"github.com/Ptt-official-app/pttbbs-backend/mand"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpdateManArticleContentInfo(
	boardID bbs.BBoardID,
	articleID types.ManArticleID,
	content [][]*types.Rune,
	contentMD5 string,
	updateNanoTS types.NanoTS,
) (err error) {
	if contentMD5 == "" {
		return nil
	}

	contentID, contentBlocks := dbcs.ParseContentBlocks(boardID, bbs.ArticleID(articleID), content, contentMD5, updateNanoTS)

	err = schema.UpdateManContentBlocks(contentBlocks, updateNanoTS)
	if err != nil {
		return err
	}

	contentInfo := &schema.ManArticleContentInfo{
		ContentMD5: contentMD5,

		ContentID: contentID,

		ContentUpdateNanoTS: updateNanoTS,
	}

	err = schema.UpdateManArticleContentInfo(boardID, articleID, contentInfo)

	return err
}

func DeserializePBManArticlesAndUpdateDB(boardID bbs.BBoardID, levelIdx types.ManArticleID, entries []*mand.Entry, updateNanoTS types.NanoTS) (articleSummaries []*schema.ManArticleSummary, err error) {
	articleSummaries = make([]*schema.ManArticleSummary, 0, len(entries))

	for idx, each := range entries {
		each := schema.NewManArticleSummaryFromPB(each, boardID, levelIdx, updateNanoTS, idx)
		if each == nil {
			continue
		}

		articleSummaries = append(articleSummaries, each)
	}

	err = schema.UpdateManArticleSummaries(articleSummaries, updateNanoTS)
	if err != nil {
		return nil, err
	}

	return articleSummaries, nil
}

func ManArticleLockKey(boardID bbs.BBoardID, articleID types.ManArticleID) (key string) {
	return "m:" + string(boardID) + ":" + string(articleID)
}

func TryGetManArticleContentInfo(userID bbs.UUserID, bboardID bbs.BBoardID, articleID types.ManArticleID, c *gin.Context, isSystem bool, isContent bool) (content [][]*types.Rune, contentMD5 string, articleDetailSummary *schema.ManArticleDetailSummary, err error) {
	updateNanoTS := types.NanoTS(0)

	articleDetailSummary, err = tryGetManArticleDetailSummary(userID, bboardID, articleID, c, isSystem)
	if err != nil {
		return nil, "", nil, err
	}

	// already update-to-date
	if articleDetailSummary.MTime <= articleDetailSummary.ContentMTime && articleDetailSummary.MTime < articleDetailSummary.ContentUpdateNanoTS {
		contentInfo, err := schema.GetManArticleContentInfo(bboardID, articleID, isContent)
		if err != nil {
			return nil, "", nil, err
		}
		return contentInfo.Content, contentInfo.ContentMD5, articleDetailSummary, nil
	}

	lockKey := ManArticleLockKey(bboardID, articleID)
	lockToken, err := schema.TryLock(lockKey, ARTICLE_LOCK_TS_DURATION)
	if err != nil {
		return nil, "", nil, err
	}
	defer func() { _ = schema.Unlock(lockKey, lockToken) }()

	// 5. get article from pttbbs
	ctx := context.Background()

	brdname := bboardID.ToBrdname()
	path := string(articleID)
	req := &mand.ArticleRequest{
		BoardName:  brdname,
		Path:       path,
		SelectType: mand.ArticleRequest_SELECT_FULL,
		MaxLength:  -1,
	}

	resp, err := mand.Cli.Article(ctx, req)
	if err != nil {
		logrus.Errorf("TryGetManArticleContentInfo: unable to get content: boardID: %v articleID: %v e: %v", bboardID, articleID, err)
		return nil, "", nil, err
	}

	if resp == nil || resp.Content == nil { // XXX possibly the article is deleted. Need to check error-code and mark the article as deleted.
		logrus.Errorf("TryGetManArticleContentInfo: no article: boardID: %v articleID: %v", bboardID, articleID)
		return nil, "", nil, ErrNoArticle
	}

	updateNanoTS = types.NowNanoTS()

	contentStr := string(resp.Content)

	content, _, contentMD5, _, _, _, _, _, _ = dbcs.ParseContentStr(contentStr, articleDetailSummary.ContentMD5, false)

	err = UpdateManArticleContentInfo(bboardID, articleID, content, contentMD5, updateNanoTS)
	if err != nil {
		return nil, "", nil, err
	}

	if contentMD5 == "" {
		contentInfo, err := schema.GetManArticleContentInfo(bboardID, articleID, isContent)
		if err != nil {
			return nil, "", nil, err
		}
		content = contentInfo.Content
		contentMD5 = contentInfo.ContentMD5
	}

	// everything is good, update content-mtime
	_ = schema.UpdateManArticleContentMTime(bboardID, articleID, articleDetailSummary.MTime)

	return content, contentMD5, articleDetailSummary, nil
}

func tryGetManArticleDetailSummary(userID bbs.UUserID, boardID bbs.BBoardID, articleID types.ManArticleID, c *gin.Context, isSystem bool) (articleDetailSummary *schema.ManArticleDetailSummary, err error) {
	articleDetailSummary, err = schema.GetManArticleDetailSummary(boardID, articleID)
	if err != nil { // something went wrong with db.
		return nil, err
	}
	if articleDetailSummary == nil {
		return nil, ErrNoArticle
	}

	return articleDetailSummary, nil
}
