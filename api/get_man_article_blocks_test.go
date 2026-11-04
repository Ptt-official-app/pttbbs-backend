package api

import (
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

func TestGetManArticleBlocks(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	// load articles
	ctx := t.Context()
	req := &mand.ListRequest{
		BoardName: "WhoAmI",
		Path:      "",
	}
	resp, _ := mand.Cli.List(ctx, req)

	posts := resp.Entries

	logrus.Infof("TestGetManArticleBlocks: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBManArticlesAndUpdateDB("10_WhoAmI", "", posts, updateNanoTS)

	//
	params0 := NewGetManArticleBlocksParams()
	path0 := &GetManArticleBlocksPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1608386280.A.BC9"),
	}

	contentSignature3 := make([][]*types.Rune, 0, len(testUtf8Content3Utf8)+len(testUtf8Signature3Utf8))
	contentSignature3 = append(contentSignature3, testUtf8Content3Utf8...)
	contentSignature3 = append(contentSignature3, testUtf8Signature3Utf8...)
	expected0 := &GetManArticleBlocksResult{
		Content:    contentSignature3,
		CreateTime: types.Time8(1608386280),
		MTime:      types.Time8(1608386280),

		Title: "[心得] 測試一下特殊字～",

		TokenUser: "SYSOP",
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     *GetManArticleBlocksParams
		path       *GetManArticleBlocksPath
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     *GetManArticleBlocksResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{userID: "SYSOP", params: params0, path: path0},
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
			gotResult, gotStatusCode, err := GetManArticleBlocks(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetManArticleBlocks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ret := gotResult.(*GetManArticleBlocksResult)
			ret.NextIdx = ""
			tt.expectedResult.NextIdx = ""
			testutil.TDeepEqual(t, "got", ret, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("GetManArticleBlocks() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
