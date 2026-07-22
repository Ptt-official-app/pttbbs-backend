package api

import (
	"context"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLoadUserArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b, testBoardSummarySYSOP_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	update0 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "19bWBI4Z", BoardArticleID: "10_WhoAmI:19bWBI4Z", ReadUpdateNanoTS: types.Time8(1534567891).ToNanoTS()}
	update1 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "1VrooM21", BoardArticleID: "10_WhoAmI:1VrooM21", ReadUpdateNanoTS: types.Time8(1234567800).ToNanoTS()}

	_, _ = schema.UserArticle_c.Update(update0, update0)
	_, _ = schema.UserArticle_c.Update(update1, update1)

	// load articles
	ctx := context.Background()
	brdname := &boardd.BoardRef_Name{Name: "WhoAmI"}
	req := &boardd.ListRequest{
		Ref:          &boardd.BoardRef{Ref: brdname},
		IncludePosts: true,
		Offset:       0,
		Length:       100 + 1,
	}
	resp, _ := boardd.Cli.List(ctx, req)

	posts := resp.Posts

	logrus.Infof("TestLoadUserArticles: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBArticlesAndUpdateDB("10_WhoAmI", posts, updateNanoTS, false)

	ctx = context.Background()
	brdname = &boardd.BoardRef_Name{Name: "SYSOP"}
	req = &boardd.ListRequest{
		Ref:          &boardd.BoardRef{Ref: brdname},
		IncludePosts: true,
		Offset:       0,
		Length:       100 + 1,
	}
	resp, _ = boardd.Cli.List(ctx, req)

	posts = resp.Posts

	logrus.Infof("TestLoadUserArticles: posts: %v", len(posts))

	updateNanoTS = types.NowNanoTS()
	_, _ = DeserializePBArticlesAndUpdateDB("1_SYSOP", posts, updateNanoTS, false)

	// params
	paramsLoadGeneralArticles := NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles := &LoadGeneralArticlesPath{FBoardID: "WhoAmI"}

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	LoadGeneralArticles("localhost", user, paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	paramsLoadGeneralArticles = NewLoadGeneralArticlesParams()
	pathLoadGeneralArticles = &LoadGeneralArticlesPath{FBoardID: "SYSOP"}
	LoadGeneralArticles("localhost", user, paramsLoadGeneralArticles, pathLoadGeneralArticles, &gin.Context{})

	articleSummary, _ := schema.GetArticleSummary("10_WhoAmI", "19bWBI4Z")
	logrus.Infof("TestLoadUserATestLoadUserArticles: articleSummary: %v", articleSummary)

	articleSummary, _ = schema.GetArticleSummary("10_WhoAmI", "1VrooM21")
	logrus.Infof("TestLoadUserATestLoadUserArticles: articleSummary: %v", articleSummary)

	articleSummary, _ = schema.GetArticleSummary("1_SYSOP", "1VrooM21")
	logrus.Infof("TestLoadUserATestLoadUserArticles: articleSummary: %v", articleSummary)

	params0 := NewUserArticlesParams()
	path0 := &LoadUserArticlesPath{
		UserID: "teemo",
	}

	expectedResult0 := &LoadUserArticlesResult{
		List: []*apitypes.ArticleSummary{
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1607937174.A.081"),
				IsDeleted:  false,
				CreateTime: types.Time8(1607937174),
				MTime:      types.Time8(1607937100),
				Recommend:  3,
				Owner:      "teemo",
				Title:      "新書的情報",
				Class:      "閒聊",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:5173/board/WhoAmI/article/M.1607937174.A.081",
				Read:       false,
				Idx:        "1607937174@1VrooM21",
			},
			{
				FBoardID:   apitypes.FBoardID("SYSOP"),
				ArticleID:  apitypes.FArticleID("M.1234567892.A.123"),
				IsDeleted:  false,
				CreateTime: types.Time8(1234567892),
				MTime:      types.Time8(1234567889),
				Recommend:  24,
				Owner:      "teemo",
				Title:      "然後呢？～",
				Class:      "問題",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:5173/board/SYSOP/article/M.1234567892.A.123",
				Read:       false,
				Idx:        "1234567892@19bWBK4Z",
			},
		},

		TokenUser: "SYSOP",
	}

	params1 := NewUserArticlesParams()
	path1 := &LoadUserArticlesPath{
		UserID: "okcool",
	}

	expectedResult1 := &LoadUserArticlesResult{
		List: []*apitypes.ArticleSummary{
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1234567890.A.123"),
				IsDeleted:  false,
				CreateTime: types.Time8(1234567890),
				MTime:      types.Time8(1234567889),
				Recommend:  8,
				Owner:      "okcool",
				Title:      "然後呢？～",
				Class:      "問題",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:5173/board/WhoAmI/article/M.1234567890.A.123",
				Read:       true,
				Idx:        "1234567890@19bWBI4Z",
			},
		},

		TokenUser: "SYSOP",
	}

	params2 := NewUserArticlesParams()
	path2 := &LoadUserArticlesPath{
		UserID: "nonexists",
	}
	expectedResult2 := &LoadUserArticlesResult{
		List: []*apitypes.ArticleSummary{},

		TokenUser: "SYSOP",
	}

	params3 := &LoadUserArticlesParams{
		Descending: true,
		Max:        2,
	}
	path3 := &LoadUserArticlesPath{
		UserID: "teemo",
	}

	expectedResult3 := &LoadUserArticlesResult{
		List: []*apitypes.ArticleSummary{
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1607937174.A.081"),
				IsDeleted:  false,
				CreateTime: types.Time8(1607937174),
				MTime:      types.Time8(1607937100),
				Recommend:  3,
				Owner:      "teemo",
				Title:      "新書的情報",
				Class:      "閒聊",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:5173/board/WhoAmI/article/M.1607937174.A.081",
				Read:       false,
				Idx:        "1607937174@1VrooM21",
			},
			{
				FBoardID:   apitypes.FBoardID("SYSOP"),
				ArticleID:  apitypes.FArticleID("M.1234567892.A.123"),
				IsDeleted:  false,
				CreateTime: types.Time8(1234567892),
				MTime:      types.Time8(1234567889),
				Recommend:  24,
				Owner:      "teemo",
				Title:      "然後呢？～",
				Class:      "問題",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:5173/board/SYSOP/article/M.1234567892.A.123",
				Read:       false,
				Idx:        "1234567892@19bWBK4Z",
			},
		},

		TokenUser: "SYSOP",
	}

	params4 := &LoadUserArticlesParams{
		StartIdx:   "1234567892@19bWBK4Z",
		Descending: true,
		Max:        1,
	}
	path4 := &LoadUserArticlesPath{
		UserID: "teemo",
	}

	expectedResult4 := &LoadUserArticlesResult{
		List: []*apitypes.ArticleSummary{
			{
				FBoardID:   apitypes.FBoardID("SYSOP"),
				ArticleID:  apitypes.FArticleID("M.1234567892.A.123"),
				IsDeleted:  false,
				CreateTime: types.Time8(1234567892),
				MTime:      types.Time8(1234567889),
				Recommend:  24,
				Owner:      "teemo",
				Title:      "然後呢？～",
				Class:      "問題",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:5173/board/SYSOP/article/M.1234567892.A.123",
				Read:       false,
				Idx:        "1234567892@19bWBK4Z",
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
			args: args{
				remoteAddr: "127.0.0.1",
				userID:     "SYSOP",
				params:     params0,
				path:       path0,
			},
			expectedResult:     expectedResult0,
			expectedStatusCode: 200,
		},
		{
			args: args{
				remoteAddr: "127.0.0.1",
				userID:     "SYSOP",
				params:     params1,
				path:       path1,
			},
			expectedResult:     expectedResult1,
			expectedStatusCode: 200,
		},
		{
			args: args{
				remoteAddr: "127.0.0.1",
				userID:     "SYSOP",
				params:     params2,
				path:       path2,
			},
			expectedResult:     expectedResult2,
			expectedStatusCode: 200,
		},
		{
			name: "limit to 2",
			args: args{
				remoteAddr: "127.0.0.1",
				userID:     "SYSOP",
				params:     params3,
				path:       path3,
			},
			expectedResult:     expectedResult3,
			expectedStatusCode: 200,
		},
		{
			name: "with start-idx",
			args: args{
				remoteAddr: "127.0.0.1",
				userID:     "SYSOP",
				params:     params4,
				path:       path4,
			},
			expectedResult:     expectedResult4,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := LoadUserArticles(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadUserArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadUserArticles() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
