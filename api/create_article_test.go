package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestCreateArticle(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, err := deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)
	if err != nil {
		logrus.Errorf("TestCreateArticle: unable to deserializeBoardsAndUpdateDB: e: %v", err)
	}

	path0 := &CreateArticlePath{
		FBoardID: "WhoAmI",
	}

	rune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	rune1 := &types.Rune{
		Utf8:   "測試1",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	params0 := &CreateArticleParams{
		PostType: "測試",
		Title:    "this is a test",
		Content: [][]*types.Rune{
			{rune0, rune1},
		},
	}

	expected0 := CreateArticleResult(&apitypes.ArticleSummary{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		ArticleID:  apitypes.FArticleID("M.1607937174.A.082"),
		IsDeleted:  false,
		CreateTime: types.Time8(1607937174),
		MTime:      types.Time8(1607937100),
		Recommend:  0,
		Owner:      "SYSOP",
		Title:      "this is a test",
		Class:      "測試",
		URL:        "http://localhost:5173/board/WhoAmI/article/M.1607937174.A.082",
		Read:       false,

		TokenUser: "SYSOP",
	})

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
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
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
			gotResult, gotStatusCode, err := CreateArticle(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)

			logrus.Infof("TestCreateArticle: got: %v statusCode: %v e: %v", gotResult, gotStatusCode, err)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CreateArticle() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
