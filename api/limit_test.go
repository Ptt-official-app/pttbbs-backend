package api

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_isValidLimit(t *testing.T) {
	tests := []struct {
		limit int
		want  bool
	}{
		{limit: -1 << 62, want: false},
		{limit: -1, want: false},
		{limit: 0, want: false},
		{limit: 1, want: true},
		{limit: DEFAULT_MAX_LIST, want: true},
		{limit: MAX_LIMIT, want: true},
		{limit: MAX_LIMIT + 1, want: false},
		{limit: 1 << 62, want: false},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, isValidLimit(tt.limit), "limit: %v", tt.limit)
	}
}

// Every list handler must reject an out-of-range limit with 400 before touching
// the database or the backend.
func TestListHandlersInvalidLimit(t *testing.T) {
	setupTest()
	defer teardownTest()

	user := &UserInfo{UserID: "SYSOP"}

	type handler func(limit int) (any, int, error)
	handlers := map[string]handler{
		"GetArticleBlocks": func(limit int) (any, int, error) {
			return GetArticleBlocks("localhost", user, &GetArticleBlocksParams{Max: limit}, &GetArticleBlocksPath{}, &gin.Context{})
		},
		"GetManArticleBlocks": func(limit int) (any, int, error) {
			return GetManArticleBlocks("localhost", user, &GetManArticleBlocksParams{Max: limit}, &GetManArticleBlocksPath{}, &gin.Context{})
		},
		"LoadArticleComments": func(limit int) (any, int, error) {
			return LoadArticleComments("localhost", user, &LoadArticleCommentsParams{Max: limit}, &LoadArticleCommentsPath{}, &gin.Context{})
		},
		"LoadAutoCompleteBoards": func(limit int) (any, int, error) {
			return LoadAutoCompleteBoards("localhost", user, &LoadAutoCompleteBoardsParams{Max: limit}, &gin.Context{})
		},
		"LoadClassBoards": func(limit int) (any, int, error) {
			return LoadClassBoards("localhost", user, &LoadClassBoardsParams{Max: limit}, &LoadClassBoardsPath{}, &gin.Context{})
		},
		"LoadFavoriteBoards": func(limit int) (any, int, error) {
			return LoadFavoriteBoards("localhost", user, &LoadFavoriteBoardsParams{Max: limit}, &LoadFavoriteBoardsPath{UserID: "SYSOP"}, &gin.Context{})
		},
		"LoadGeneralArticles": func(limit int) (any, int, error) {
			return LoadGeneralArticles("localhost", user, &LoadGeneralArticlesParams{Max: limit}, &LoadGeneralArticlesPath{}, &gin.Context{})
		},
		"LoadGeneralArticlesByKeyword": func(limit int) (any, int, error) {
			return LoadGeneralArticles("localhost", user, &LoadGeneralArticlesParams{Max: limit, Keyword: "a"}, &LoadGeneralArticlesPath{}, &gin.Context{})
		},
		"LoadGeneralBoards": func(limit int) (any, int, error) {
			return LoadGeneralBoards("localhost", user, &LoadGeneralBoardsParams{Max: limit}, &gin.Context{})
		},
		"LoadUserArticles": func(limit int) (any, int, error) {
			return LoadUserArticles("localhost", user, &LoadUserArticlesParams{Max: limit}, &LoadUserArticlesPath{UserID: "SYSOP"}, &gin.Context{})
		},
		"LoadUserComments": func(limit int) (any, int, error) {
			return LoadUserComments("localhost", user, &LoadUserCommentsParams{Max: limit}, &LoadUserCommentsPath{UserID: "SYSOP"}, &gin.Context{})
		},
	}

	for name, h := range handlers {
		for _, limit := range []int{-1, 0, MAX_LIMIT + 1, 1 << 62} {
			result, statusCode, err := h(limit)
			assert.Nil(t, result, "%v limit: %v", name, limit)
			assert.Equal(t, 400, statusCode, "%v limit: %v", name, limit)
			assert.ErrorIs(t, err, ErrInvalidLimit, "%v limit: %v", name, limit)
		}
	}
}
