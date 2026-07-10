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

func TestDeleteArticles(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	createpath0 := &CreateArticlePath{
		FBoardID: "WhoAmI",
	}

	createRune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}
	createParams0 := &CreateArticleParams{
		PostType: "測試",
		Title:    "this is a test",
		Content: [][]*types.Rune{
			{createRune0},
		},
	}
	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	got0, _, _ := CreateArticle(testIP, user, createParams0, createpath0, nil)
	gotCreateArticle0, _ := got0.(CreateArticleResult)
	path0 := &DeleteArticlesPath{
		FBoardID: "WhoAmI",
	}
	params0 := &DeleteArticlesParams{
		ArticleIDs: []apitypes.FArticleID{gotCreateArticle0.ArticleID},
	}

	expected0 := &DeleteArticlesResult{
		Success:   true,
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
		name           string
		args           args
		wantResult     interface{}
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			name:           "test delete article success",
			args:           args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			wantResult:     expected0,
			wantStatusCode: 200,
			wantErr:        false,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}

			gotResult, gotStatusCode, err := DeleteArticles(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			logrus.Infof("TestDeleteArticles: got: %v statusCode: %v e: %v", gotResult, gotStatusCode, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "got", gotResult, tt.wantResult)

			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("DeleteArticles() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
