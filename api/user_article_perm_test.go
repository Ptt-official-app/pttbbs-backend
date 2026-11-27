package api

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/stretchr/testify/assert"
)

func TestCheckUserArticlePermDeletableCore(t *testing.T) {
	setupTest()
	defer teardownTest()

	articleID := bbs.ArticleID("1VrooM21SYSOP")
	tests := []struct {
		name      string
		owner     bbs.UUserID
		isDeleted bool
		perm      *UserBoardPermReadable
		wantErr   error
	}{
		{name: "owner", owner: testPermUserID, perm: &UserBoardPermReadable{}},
		{name: "deleted, even for the owner", owner: testPermUserID, isDeleted: true, perm: &UserBoardPermReadable{IsSYSOP: true}, wantErr: ErrNoArticle},
		{name: "BM", owner: "other", perm: &UserBoardPermReadable{IsBM: true}},
		{name: "sysop", owner: "other", perm: &UserBoardPermReadable{IsSYSOP: true}},
		{name: "police is not enough", owner: "other", perm: &UserBoardPermReadable{IsPolice: true}, wantErr: ErrInvalidUser},
		{name: "board friend is not enough", owner: "other", perm: &UserBoardPermReadable{IsBoardFriend: true}, wantErr: ErrInvalidUser},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articlePermInfo := &schema.ArticlePermInfo{BBoardID: testPermBoardID, ArticleID: articleID, Owner: tt.owner, IsDeleted: tt.isDeleted}

			err := checkUserArticlePermDeletableCore(testPermUserID, testPermBoardID, articleID, articlePermInfo, tt.perm)
			assert.ErrorIs(t, err, tt.wantErr)

			// the batch version must agree with the single-article version.
			permMap := checkUserArticlesPermDeletableCore(testPermUserID, []*schema.ArticlePermInfo{articlePermInfo}, tt.perm)
			assert.ErrorIs(t, permMap[articleID], tt.wantErr)
		})
	}
}

func TestCheckUserArticlesPermEditableCore(t *testing.T) {
	setupTest()
	defer teardownTest()

	articlesPermInfo := []*schema.ArticlePermInfo{
		{ArticleID: "owned", Owner: testPermUserID},
		{ArticleID: "others", Owner: "other"},
		{ArticleID: "deleted", Owner: testPermUserID, IsDeleted: true},
	}

	// only the owner can edit: being BM or SYSOP does not matter.
	got := checkUserArticlesPermEditableCore(testPermUserID, articlesPermInfo, &UserBoardPermReadable{IsSYSOP: true, IsBM: true})

	expected := map[bbs.ArticleID]error{
		"owned":   nil,
		"others":  ErrInvalidUser,
		"deleted": ErrNoArticle,
	}
	assert.Equal(t, expected, got)
}
