package api

import (
	"context"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/mand"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLoadManArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	// load articles
	ctx := context.Background()
	req := &mand.ListRequest{
		BoardName: "WhoAmI",
		Path:      "",
	}
	resp, _ := mand.Cli.List(ctx, req)

	posts := resp.Entries

	logrus.Infof("TestGetManArticles: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBManArticlesAndUpdateDB("10_WhoAmI", "", posts, updateNanoTS)

	// params
	params0 := &LoadManArticlesParams{LevelIdx: ""}
	path0 := &LoadManArticlesPath{FBoardID: "WhoAmI"}
	expected0 := &LoadManArticlesResult{
		List: []*apitypes.ManArticleSummary{
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1584665022.A.ED0",
				CreateTime: 1584665022,
				MTime:      1584665022,
				Title:      "[問卦] 為何打麻將叫賭博但買股票叫投資？",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1608386280.A.BC9",
				CreateTime: 1608386280,
				MTime:      1608386280,
				Title:      "[心得] 測試一下特殊字～",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1608388506.A.85D",
				CreateTime: 1608388506,
				MTime:      1608388506,
				Title:      "[閒聊] 所以特殊字真的是有綠色的∼",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1607202240.A.30D",
				CreateTime: 1607202240,
				MTime:      1607202240,
				Title:      "[新聞] TVBS六都民調 侯奪冠、盧升第四、柯墊底",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1607937174.A.081",
				CreateTime: 1607937174,
				MTime:      1607937174,
				Title:      "[閒聊] 新書的情報",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1234567890.A.123",
				CreateTime: 1234567890,
				MTime:      1234567890,
				Title:      "[問題]然後呢？～",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1234560000.A.081",
				CreateTime: 1234560000,
				MTime:      1234560000,
				Title:      "[問題]這是 SYSOP",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1607937176.A.081",
				CreateTime: 1607937176,
				MTime:      1607937176,
				Title:      "[心得] 測試一下特殊字～",
			},
			{
				FBoardID:   "WhoAmI",
				ArticleID:  "M.1234567892.A.123",
				CreateTime: 1234567892,
				MTime:      1234567892,
				Title:      "[問題]然後呢？～",
			},
		},

		TokenUser: "SYSOP",
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     *LoadManArticlesParams
		path       *LoadManArticlesPath
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     *LoadManArticlesResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: "testIP", userID: "SYSOP", params: params0, path: path0},
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
			gotResult, gotStatusCode, err := LoadManArticles(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadManArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadManArticles() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
