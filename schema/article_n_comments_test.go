package schema

import (
	"cmp"
	"slices"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestGetArticleNCommentsByArticleIDs(t *testing.T) {
	setupTest()
	defer teardownTest()

	updateNanoTS := types.NowNanoTS()

	articleSummary1 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("1_test1"),
		ArticleID:  bbs.ArticleID("1VrooM21teemo"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    3,
		Owner:        bbs.UUserID("teemo"),
		Title:        "再來呢？～",
		Money:        12,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
	}

	articleSummary2 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("1_test1"),
		ArticleID:  bbs.ArticleID("1VrooM21teem2"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    4,
		Owner:        bbs.UUserID("teem2"),
		Title:        "再來呢2？～",
		Money:        15,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
	}

	articleSummaries1 := []*ArticleSummary{articleSummary1, articleSummary2}

	_ = UpdateArticleSummaries(articleSummaries1, updateNanoTS)

	articleNComments0 := []*ArticleNComments{
		{
			BBoardID:  "1_test1",
			ArticleID: "1VrooM21teem2",
		},
		{
			BBoardID:  "1_test1",
			ArticleID: "1VrooM21teemo",
		},
	}

	type args struct {
		bboardID   bbs.BBoardID
		articleIDs []bbs.ArticleID
	}
	tests := []struct {
		name                     string
		args                     args
		expectedArticleNComments []*ArticleNComments
		wantErr                  bool
	}{
		// TODO: Add test cases.
		{
			args:                     args{bboardID: "1_test1", articleIDs: []bbs.ArticleID{"1VrooM21teemo", "1VrooM21teem2"}},
			expectedArticleNComments: articleNComments0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArticleNComments, err := GetArticleNCommentsByArticleIDs(tt.args.bboardID, tt.args.articleIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleNCommentsByArticleIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			slices.SortStableFunc(gotArticleNComments, func(a, b *ArticleNComments) int {
				return cmp.Compare(a.ArticleID, b.ArticleID)
			})

			testutil.TDeepEqual(t, "got", gotArticleNComments, tt.expectedArticleNComments)
		})
	}
}
