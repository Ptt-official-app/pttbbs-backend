package api

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// markTestArticleDeleted marks the article as deleted in db.
func markTestArticleDeleted(t *testing.T, boardID bbs.BBoardID, articleID bbs.ArticleID) {
	t.Helper()

	_, err := schema.Article_c.UpdateOneOnly(
		&schema.ArticleQuery{BBoardID: boardID, ArticleID: articleID},
		bson.M{schema.ARTICLE_IS_DELETED_b: true},
	)
	assert.NoError(t, err)
}

func TestTryGetArticleContentInfo_deleted(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}, 123456890000000000)
	_, _, _ = deserializeArticlesAndUpdateDB("SYSOP", "10_WhoAmI", []*bbs.ArticleSummary{testArticleSummary1_b}, 123456892000000000, false)
	markTestArticleDeleted(t, "10_WhoAmI", "1VrooM21")

	_, _, _, _, _, _, _, _, _, fileSize, hash, statusCode, err := TryGetArticleContentInfo("SYSOP", "10_WhoAmI", "1VrooM21", &gin.Context{}, false, false, true)

	assert.ErrorIs(t, err, ErrAlreadyDeleted)
	assert.Equal(t, 404, statusCode)
	assert.Equal(t, 0, fileSize)
	assert.Equal(t, 0, int(hash))
}

func TestLoadArticleComments_noPerm(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}, 123456890000000000)
	_, _, _ = deserializeArticlesAndUpdateDB("SYSOP", "10_WhoAmI", []*bbs.ArticleSummary{testArticleSummary1_b}, 123456892000000000, false)
	markTestArticleDeleted(t, "10_WhoAmI", "1VrooM21")

	user := &UserInfo{UserID: "SYSOP"}
	path := &LoadArticleCommentsPath{FBoardID: "WhoAmI", FArticleID: "M.1607937174.A.081"}
	result, statusCode, err := LoadArticleComments("localhost", user, NewLoadArticleCommentsParams(), path, &gin.Context{})

	assert.ErrorIs(t, err, ErrNoArticle)
	assert.Nil(t, result)
	assert.Equal(t, 403, statusCode)
}

func TestLoadClassBoards_noPerm(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)

	updateNanoTS := types.NowNanoTS()
	hiddenClass := &schema.BoardSummary{
		BBoardID:   "3_hidden",
		Brdname:    "hidden",
		BrdAttr:    ptttype.BRD_HIDE,
		Bid:        3,
		IdxByName:  "hidden",
		IdxByClass: "hidden",
	}
	_ = schema.UpdateBoardSummaries([]*schema.BoardSummary{hiddenClass}, updateNanoTS)

	user := &UserInfo{UserID: "chhsiao123"}
	path := &LoadClassBoardsPath{ClsID: 3}
	result, statusCode, err := LoadClassBoards("localhost", user, NewLoadClassBoardsParams(), path, &gin.Context{})

	assert.ErrorIs(t, err, ErrPermBoardReadHidden)
	assert.Nil(t, result)
	assert.Equal(t, 403, statusCode)
}

func TestGetManArticleDetail_noArticle(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}, 123456890000000000)

	user := &UserInfo{UserID: "SYSOP"}
	path := &GetManArticleDetailPath{FBoardID: apitypes.FBoardID("WhoAmI"), FArticleID: "M.1608386280.A.BC9"}
	result, statusCode, err := GetManArticleDetail("localhost", user, nil, path, &gin.Context{})

	assert.ErrorIs(t, err, ErrNoArticle)
	assert.Nil(t, result)
	assert.Equal(t, 500, statusCode)
}
