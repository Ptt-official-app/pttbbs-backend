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

func TestLoadBottomArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	update0 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "19bWBI4Z", ReadUpdateNanoTS: types.Time8(1534567891).ToNanoTS()}
	update1 := &schema.UserArticle{UserID: "SYSOP", BoardID: "10_WhoAmI", ArticleID: "1VrooM21", ReadUpdateNanoTS: types.Time8(1234567800).ToNanoTS()}

	_, _ = schema.UserArticle_c.Update(update0, update0)
	_, _ = schema.UserArticle_c.Update(update1, update1)

	// load articles
	ctx := context.Background()
	brdname := &boardd.BoardRef_Name{Name: "WhoAmI"}
	req := &boardd.ListRequest{
		Ref:            &boardd.BoardRef{Ref: brdname},
		Offset:         0,
		Length:         100 + 1,
		IncludeBottoms: true,
	}
	resp, _ := boardd.Cli.List(ctx, req)

	posts := resp.Posts

	logrus.Infof("TestGetArticleDetail: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBArticlesAndUpdateDB("10_WhoAmI", posts, updateNanoTS, true)

	// params
	path := &LoadBottomArticlesPath{FBoardID: "WhoAmI"}
	expectedResult := &LoadGeneralArticlesResult{
		List: []*apitypes.ArticleSummary{
			{
				FBoardID:   apitypes.FBoardID("WhoAmI"),
				ArticleID:  apitypes.FArticleID("M.1234560000.A.081"),
				IsDeleted:  false,
				CreateTime: types.Time8(1234560000),
				MTime:      types.Time8(1234560000),
				Recommend:  13,
				Owner:      "SYSOP",
				Title:      "這是 SYSOP",
				Class:      "問題",
				Money:      0,
				Filemode:   0,
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1234560000.A.081",
				Read:       false,

				Idx: "1234560000@19bUG021",

				Editable:  true,
				Deletable: true,
			},
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
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1234567890.A.123",
				Read:       true,

				Idx:       "1234567890@19bWBI4Z",
				Deletable: true,
			},
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
				URL:        "http://localhost:3457/bbs/board/WhoAmI/article/M.1607937174.A.081",
				Read:       false,

				Idx:       "1607937174@1VrooM21",
				Deletable: true,
			},
		},
		NextIdx: "",

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
				remoteAddr: testIP,
				userID:     "SYSOP",
				path:       path,
				c:          &gin.Context{},
			},
			expectedResult:     expectedResult,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := LoadBottomArticles(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBottomArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadBottomArticles() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
