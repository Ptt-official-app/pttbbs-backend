package schema

import (
	"cmp"
	"slices"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateUserReadArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserArticle_c.Drop()

	userReadArticles := []*UserReadArticle{
		{
			UserID:           bbs.UUserID("testuser0"),
			BoardID:          bbs.BBoardID("10_WhoAmI"),
			ArticleID:        bbs.ArticleID("testarticle0"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:           bbs.UUserID("testuser0"),
			BoardID:          bbs.BBoardID("10_WhoAmI"),
			ArticleID:        bbs.ArticleID("testarticle1"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:           bbs.UUserID("testuser0"),
			BoardID:          bbs.BBoardID("10_WhoAmI"),
			ArticleID:        bbs.ArticleID("testarticle2"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	type args struct {
		userReadArticles []*UserReadArticle
		updateNanoTS     types.NanoTS
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		findUserID     bbs.UUserID
		findBoardID    bbs.BBoardID
		findArticleIDs []bbs.ArticleID
		want           []*UserReadArticle
	}{
		// TODO: Add test cases.
		{
			args:           args{userReadArticles: userReadArticles, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:     bbs.UUserID("testuser0"),
			findBoardID:    bbs.BBoardID("10_WhoAmI"),
			findArticleIDs: []bbs.ArticleID{"testarticle0", "testarticle1", "testarticle2"},
			want:           userReadArticles,
		},
		{
			args:           args{userReadArticles: userReadArticles, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:     bbs.UUserID("testuser0"),
			findBoardID:    bbs.BBoardID("10_WhoAmI"),
			findArticleIDs: []bbs.ArticleID{"testarticle0", "testarticle1", "testarticle2"},
			want:           userReadArticles,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadArticles(tt.args.userReadArticles, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadArticles() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := FindUserReadArticles(tt.findUserID, tt.findBoardID, tt.findArticleIDs)
			if err != nil {
				t.Errorf("UpdateUserReadArticles: e: %v", err)
			}

			slices.SortStableFunc(got, func(a, b *UserReadArticle) int {
				return cmp.Compare(a.ArticleID, b.ArticleID)
			})
			testutil.TDeepEqual(t, "got", got, tt.want)
		})
	}
}

func TestUpdateUserReadArticle(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserArticle_c.Drop()

	userReadArticle := &UserReadArticle{
		UserID:           bbs.UUserID("testuser0"),
		BoardID:          bbs.BBoardID("10_WhoAmI"),
		ArticleID:        bbs.ArticleID("testarticle0"),
		ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
	}
	type args struct {
		userReadArticle *UserReadArticle
	}
	tests := []struct {
		name     string
		args     args
		expected *UserReadArticle
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{userReadArticle: userReadArticle},
			expected: userReadArticle,
		},
		{
			args:     args{userReadArticle: userReadArticle},
			expected: userReadArticle,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadArticle(tt.args.userReadArticle); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadArticle() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := &UserReadArticle{}

			query := bson.M{
				USER_ARTICLE_USER_ID_b:    tt.args.userReadArticle.UserID,
				USER_ARTICLE_BOARD_ID_b:   tt.args.userReadArticle.BoardID,
				USER_ARTICLE_ARTICLE_ID_b: tt.args.userReadArticle.ArticleID,
			}
			_ = UserArticle_c.FindOne(query, got, nil)

			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}

func TestUpdateUserReadArticles_sameArticleIDInOtherBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserArticle_c.Drop()

	newRow := func(boardID bbs.BBoardID, nanoTS types.NanoTS) *UserReadArticle {
		return &UserReadArticle{
			UserID:           "testuser0",
			BoardID:          boardID,
			ArticleID:        "samearticle",
			ReadUpdateNanoTS: nanoTS,
		}
	}

	err := UpdateUserReadArticles([]*UserReadArticle{newRow("10_WhoAmI", 1000), newRow("11_Other", 1000)}, 1000)
	require.NoError(t, err)

	// only the row of 11_Other is newer.
	err = UpdateUserReadArticles([]*UserReadArticle{newRow("11_Other", 2000)}, 2000)
	require.NoError(t, err)

	got, err := FindUserReadArticles("testuser0", "10_WhoAmI", []bbs.ArticleID{"samearticle"})
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Equal(t, types.NanoTS(1000), got[0].ReadUpdateNanoTS)

	got, err = FindUserReadArticles("testuser0", "11_Other", []bbs.ArticleID{"samearticle"})
	require.NoError(t, err)
	require.Len(t, got, 1)
	assert.Equal(t, types.NanoTS(2000), got[0].ReadUpdateNanoTS)
}
