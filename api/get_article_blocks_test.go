package api

import (
	"context"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestGetArticleBlocks(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

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

	logrus.Infof("TestGetArticleDetail: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBArticlesAndUpdateDB("10_WhoAmI", posts, updateNanoTS, false)

	//
	params0 := NewGetArticleBlocksParams()
	path0 := &GetArticleBlocksPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1608386280.A.BC9"),
	}

	expected0 := &GetArticleBlocksResult{
		Content:    testContent3Utf8[4:],
		Owner:      bbs.UUserID("SYSOP"),
		CreateTime: types.Time8(1608386280),
		MTime:      types.Time8(1608386280),

		Title:     "測試一下特殊字～",
		Money:     0,
		Recommend: 9,
		Class:     "心得",
		IP:        "172.22.0.1",
		Host:      "",
		BBS:       "批踢踢 docker(pttdocker.test)",

		TokenUser: "SYSOP",
	}

	params1 := NewGetArticleBlocksParams()
	path1 := &GetArticleBlocksPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1607202240.A.30D"),
	}

	expected1 := &GetArticleBlocksResult{
		Content:    testContent11Utf8[4:54],
		Owner:      bbs.UUserID("cheinshin"),
		CreateTime: types.Time8(1607202240),
		MTime:      types.Time8(1607202240),

		Title:     "TVBS六都民調 侯奪冠、盧升第四、柯墊底",
		Money:     0,
		Recommend: 23,
		Class:     "新聞",
		IP:        "49.216.65.39",
		Host:      "臺灣",
		BBS:       "批踢踢實業坊(ptt.cc)",
		NextIdx:   "FsQhVG-oT3A:IKCj3KzpwP5pcJxOAPNDNQ^1",

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
		expectedResult     *GetArticleBlocksResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{userID: "SYSOP", params: params0, path: path0},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
		{
			args:               args{userID: "SYSOP", params: params1, path: path1},
			expectedResult:     expected1,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := GetArticleBlocks(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ret := gotResult.(*GetArticleBlocksResult)
			ret.NextIdx = ""
			tt.expectedResult.NextIdx = ""
			testutil.TDeepEqual(t, "got", ret, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("GetArticleBlocks() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
