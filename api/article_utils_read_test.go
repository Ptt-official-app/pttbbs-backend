package api

import (
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_setUserReadArticle(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserArticle_c.Drop()

	boardID := bbs.BBoardID("10_WhoAmI")
	articleID := bbs.ArticleID("1VtW-QXT")
	updateNanoTS := types.NowNanoTS()
	content := [][]*types.Rune{}

	guestID := bbs.UUserID(pttbbsapi.GUEST)
	setUserReadArticle(content, guestID, boardID, articleID, updateNanoTS)
	setUserReadArticle(content, "testuser0", boardID, articleID, updateNanoTS)

	// guest: nothing is persisted, neither with the real nor with the empty board-id.
	for _, bid := range []bbs.BBoardID{boardID, ""} {
		got, err := schema.FindUserReadArticles(guestID, bid, []bbs.ArticleID{articleID})
		require.NoError(t, err)
		assert.Empty(t, got, "guest, bid: %v", bid)
	}

	// user: stored with the board-id so that the read-state lookups can find it.
	got, err := schema.FindUserReadArticles("testuser0", boardID, []bbs.ArticleID{articleID})
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Equal(t, boardID, got[0].BoardID)
	assert.Equal(t, types.ToBoardArticleID(boardID, articleID), got[0].BoardArticleID)

	got, err = schema.FindUserReadArticlesByBoardArticleIDs("testuser0", []types.BoardArticleID{types.ToBoardArticleID(boardID, articleID)})
	require.NoError(t, err)
	assert.Len(t, got, 1)
}

func Test_deserializeArticlesAndUpdateDB_userReadArticle(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserArticle_c.Drop()
	defer schema.Article_c.Drop()

	boardID := bbs.BBoardID("10_WhoAmI")
	articleID := bbs.ArticleID("1VtW-QXT")
	newList := func() []*bbs.ArticleSummary {
		return []*bbs.ArticleSummary{{
			BBoardID:   boardID,
			ArticleID:  articleID,
			CreateTime: 1608388506,
			MTime:      1608388508,
			Owner:      "SYSOP",
			FullTitle:  []byte("[測試] test title"),
			RealTitle:  []byte("test title"),
			Class:      []byte("測試"),
			Idx:        "1608388506@1VtW-QXT",
			Read:       true,
		}}
	}

	guestID := bbs.UUserID(pttbbsapi.GUEST)
	_, readMap, err := deserializeArticlesAndUpdateDB(guestID, boardID, newList(), types.NowNanoTS(), false)
	require.NoError(t, err)
	assert.True(t, readMap[articleID])
	got, err := schema.FindUserReadArticles(guestID, boardID, []bbs.ArticleID{articleID})
	require.NoError(t, err)
	assert.Empty(t, got)
	got, err = schema.FindUserReadArticles(guestID, "", []bbs.ArticleID{articleID})
	require.NoError(t, err)
	assert.Empty(t, got)

	_, _, err = deserializeArticlesAndUpdateDB("testuser0", boardID, newList(), types.NowNanoTS(), false)
	require.NoError(t, err)
	got, err = schema.FindUserReadArticles("testuser0", boardID, []bbs.ArticleID{articleID})
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Equal(t, boardID, got[0].BoardID)
}
