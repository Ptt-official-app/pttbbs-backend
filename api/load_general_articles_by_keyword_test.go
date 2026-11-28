package api

import (
	"errors"
	"strings"
	"testing"
)

func TestLoadGeneralArticlesByKeywordTooManyKeywords(t *testing.T) {
	setupTest()
	defer teardownTest()

	user := &UserInfo{UserID: "SYSOP"}
	path := &LoadGeneralArticlesPath{FBoardID: "WhoAmI"}
	keywords := strings.TrimSpace(strings.Repeat("然後 ", MAX_ARTICLE_KEYWORDS+1))
	params := &LoadGeneralArticlesParams{Keyword: keywords, Max: 200, Descending: true}

	_, statusCode, err := LoadGeneralArticlesByKeyword("localhost", user, params, path, nil)
	if statusCode != 400 || !errors.Is(err, ErrTooManyKeywords) {
		t.Errorf("LoadGeneralArticlesByKeyword() statusCode: %v e: %v, want 400 ErrTooManyKeywords", statusCode, err)
	}
}
