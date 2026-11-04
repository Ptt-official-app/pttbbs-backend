package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/mand"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestGetManArticleDetail(t *testing.T) {
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

	logrus.Infof("TestGetManArticleDetail: posts: %v", len(posts))

	updateNanoTS := types.NowNanoTS()
	_, _ = DeserializePBManArticlesAndUpdateDB("10_WhoAmI", "", posts, updateNanoTS)

	// params
	path0 := &GetManArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1608386280.A.BC9"),
	}

	testContentSignature0 := make([][]*types.Rune, 0, len(testUtf8Content3Utf8)+len(testUtf8Signature3Utf8))
	testContentSignature0 = append(testContentSignature0, testUtf8Content3Utf8...)
	testContentSignature0 = append(testContentSignature0, testUtf8Signature3Utf8...)
	expectedResult0 := &GetManArticleDetailResult{
		BBoardID:   apitypes.FBoardID("WhoAmI"),
		ArticleID:  apitypes.FArticleID("M.1608386280.A.BC9"),
		CreateTime: types.Time8(1608386280),
		MTime:      types.Time8(1608386280),

		Title:   "[心得] 測試一下特殊字～",
		Content: testContentSignature0,

		TokenUser: "SYSOP",
	}

	expectedArticleDetailSummary0 := &schema.ManArticleDetailSummary{
		BBoardID:   bbs.BBoardID("10_WhoAmI"),
		ArticleID:  types.ManArticleID("M.1608386280.A.BC9"),
		ContentMD5: "1slmiNd1Hf0Pb5oaFzKYHw",
		// ContentMTime:          types.NanoTS(1608386280000000000),

		CreateTime: types.NanoTS(1608386280000000000),
		MTime:      types.NanoTS(1608386280000000000),
		Title:      "[心得] 測試一下特殊字～",

		Idx: 1,

		ContentMTime: 1608386280000000000,
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		path       *GetManArticleDetailPath
		c          *gin.Context
	}
	tests := []struct {
		name                         string
		args                         args
		expectedResult               *GetManArticleDetailResult
		expectedArticleDetailSummary *schema.ManArticleDetailSummary
		expectedStatusCode           int
		wantErr                      bool
		boardID                      bbs.BBoardID
		articleID                    types.ManArticleID
	}{
		// TODO: Add test cases.
		{
			args:                         args{path: path0, userID: "SYSOP", remoteAddr: "testIP"},
			boardID:                      "10_WhoAmI",
			articleID:                    "M.1608386280.A.BC9",
			expectedResult:               expectedResult0,
			expectedArticleDetailSummary: expectedArticleDetailSummary0,
			expectedStatusCode:           200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := GetManArticleDetail(tt.args.remoteAddr, user, nil, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetManArticleDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("GetManArticleDetail() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}

			gotArticleDetailSummary, err := schema.GetManArticleDetailSummary(tt.boardID, tt.articleID)
			if err != nil {
				t.Errorf("GetManArticleDetail(): unable to get summary: e: %v", err)
			}

			gotArticleDetailSummary.ContentID = ""
			gotArticleDetailSummary.ContentUpdateNanoTS = 0
			gotArticleDetailSummary.UpdateNanoTS = 0
			testutil.TDeepEqual(t, "summary", gotArticleDetailSummary, tt.expectedArticleDetailSummary)
		})
		wg.Wait()
	}
}
