package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func TestLoadUserComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = deserializeUserDetailAndUpdateDB(testUserSYSOP_b, 123456890000000000)
	_, _ = deserializeUserDetailAndUpdateDB(testUserChhsiao123_b, 123456891000000000)

	// board
	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b, testBoardSummarySYSOP_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	articleSummaries_b := []*bbs.ArticleSummary{testArticleSummary1_b, testArticleSummary2_b}
	_, _, _ = deserializeArticlesAndUpdateDB("SYSOP", "10_WhoAmI", articleSummaries_b, 123456892000000000, false)

	// articles
	update0 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "19bWBI4Z", BoardArticleID: "10_WhoAmI:19bWBI4Z", ReadUpdateNanoTS: types.Time8(1608388624).ToNanoTS()}
	update1 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "1VrooM21", BoardArticleID: "10_WhoAmI:1VrooM21", ReadUpdateNanoTS: types.Time8(1608388624).ToNanoTS()}

	_, _ = schema.UserArticle_c.Update(update0, update0)
	_, _ = schema.UserArticle_c.Update(update1, update1)

	paramsLoadGeneralArticles := NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles := &LoadGeneralArticlesPath{FBoardID: "WhoAmI"}

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	LoadGeneralArticles("localhost", user, paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	paramsLoadGeneralArticles = NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles = &LoadGeneralArticlesPath{FBoardID: "SYSOP"}
	LoadGeneralArticles("localhost", user, paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	articleParams := &GetArticleDetailParams{}
	articlePath := &GetArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1607937174.A.081"),
	}
	_, _, _ = GetArticleDetail(testIP, user, articleParams, articlePath, nil)

	// tests
	params0 := NewLoadUserCommentsParams()
	path0 := &LoadUserCommentsPath{
		UserID: "chhsiao123",
	}

	expected0 := &LoadUserCommentsResult{
		List: []*apitypes.ArticleComment{
			{
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1607937174.A.081",
				CreateTime:        1607937174,
				MTime:             1607937100,
				Recommend:         3,
				NComments:         3,
				Owner:             "teemo",
				Title:             "再來呢？～",
				Money:             12,
				Class:             "問題",
				URL:               "http://localhost:3457/bbs/board/WhoAmI/article/M.1607937174.A.081",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1608388624000000000@FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
				CommentType:       ptttype.COMMENT_TYPE_BOO,
				CommentCreateTime: 1608388624,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:   "噓～",
							Color0: types.DefaultColor,
							Color1: types.DefaultColor,
							Big5:   []byte("\xbcN\xa1\xe3                                                "),
							DBCS:   []byte("\xbcN\xa1\xe3                                                "),
						},
					},
				},
			},
			{
				FBoardID:          "WhoAmI",
				FArticleID:        "M.1607937174.A.081",
				CreateTime:        1607937174,
				MTime:             1607937100,
				Recommend:         3,
				NComments:         3,
				Owner:             "teemo",
				Title:             "再來呢？～",
				Money:             12,
				Class:             "問題",
				URL:               "http://localhost:3457/bbs/board/WhoAmI/article/M.1607937174.A.081",
				Read:              types.READ_STATUS_MTIME,
				Idx:               "c#1608388560000000000@FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg",
				TheType:           apitypes.ARTICLE_COMMENT_TYPE_COMMENT,
				CommentID:         "FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg",
				CommentType:       ptttype.COMMENT_TYPE_RECOMMEND,
				CommentCreateTime: 1608388560,
				Comment: [][]*types.Rune{
					{
						{
							Utf8:   "推",
							Color0: types.DefaultColor,
							Color1: types.DefaultColor,
							Big5:   []byte("\xb1\xc0                                                  "),
							DBCS:   []byte("\xb1\xc0                                                  "),
						},
					},
				},
			},
		},

		TokenUser: "SYSOP",
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
		path       interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     interface{}
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: params0, path: path0},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := LoadUserComments(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadUserComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadUserComments() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
