package api

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestEditArticleDetail(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	createPath0 := &CreateArticlePath{
		FBoardID: "WhoAmI",
	}

	createRune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	createRune1 := &types.Rune{
		Utf8:   "測試1",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	createParams0 := &CreateArticleParams{
		PostType: "測試",
		Title:    "this is a test",
		Content: [][]*types.Rune{
			{createRune0, createRune1},
		},
	}

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	got0, _, _ := CreateArticle(testIP, user, createParams0, createPath0, nil)
	gotCreateArticle0, _ := got0.(CreateArticleResult)

	content0 := [][]*types.Rune{
		{
			{
				Utf8:   "測試2",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
		{
			{
				Utf8:   "測試3",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
			},
		},
	}

	params0 := &EditArticleParams{
		Content: content0,
	}
	path0 := &EditArticlePath{
		FBoardID:   "WhoAmI",
		FArticleID: gotCreateArticle0.ArticleID,
	}

	expectedContent0 := [][]*types.Rune{
		{
			{
				Utf8:   "測試2",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb4\xfa\xb8\xd52\r"),
			},
		},
		{
			{
				Utf8:   "測試3",
				Color0: types.DefaultColor,
				Color1: types.DefaultColor,
				DBCS:   []byte("\xb4\xfa\xb8\xd53\r"),
			},
		},
	}
	expected0 := &EditArticleResult{
		Content: expectedContent0,
		MTime:   1583511858,
		Title:   "this is a test",
		Class:   "測試",

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
		expectedResult     *EditArticleResult
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := EditArticleDetail(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			logrus.Infof("TestEditArticleDetail: got: %v statusCode: %v e: %v", gotResult, gotStatusCode, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditArticleDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("EditArticleDetail() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
