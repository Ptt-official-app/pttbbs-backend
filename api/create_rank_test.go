package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRank(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	for _, userID := range []bbs.UUserID{"SYSOP1", "SYSOP2", "SYSOP3"} {
		user_b := *testUserChhsiao123_b
		user_b.UUserID = userID
		user_b.Username = string(userID)
		_, _ = DeserializeUserDetailAndUpdateDBForTest(&user_b, 123456891000000000)
	}

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	schema.UpdateArticleSummaries([]*schema.ArticleSummary{{BBoardID: "10_WhoAmI", ArticleID: "1VrooM21", Owner: "test"}}, 1234567890000000000)

	path0 := &CreateRankPath{
		FBoardID:   "WhoAmI",
		FArticleID: "M.1607937174.A.081",
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     *CreateRankParams
		path       *CreateRankPath
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     any
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name:               "invalid rank",
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: &CreateRankParams{Rank: 1000000000}, path: path0},
			wantErr:            true,
			expectedStatusCode: 400,
		},
		{
			name:               "guest cannot do rank",
			args:               args{remoteAddr: testIP, userID: "guest", params: &CreateRankParams{Rank: 1}, path: path0},
			wantErr:            true,
			expectedStatusCode: 403,
		},
		{
			name:               "owner cannot do rank",
			args:               args{remoteAddr: testIP, userID: "test", params: &CreateRankParams{Rank: 1}, path: path0},
			wantErr:            true,
			expectedStatusCode: 403,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 1, TokenUser: "SYSOP"},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: &CreateRankParams{Rank: -1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: -1, TokenUser: "SYSOP"},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP1", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 0, TokenUser: "SYSOP1"},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP2", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 1, TokenUser: "SYSOP2"},
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP3", params: &CreateRankParams{Rank: 1}, path: path0},
			expectedResult:     &CreateRankResult{Rank: 2, TokenUser: "SYSOP3"},
			expectedStatusCode: 200,
		},
	}

	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := CreateRank(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("CreateRank() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CreateRank() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}

			if tt.expectedResult == nil {
				return
			}

			path0 := &GetArticleDetailPath{FBoardID: tt.args.path.FBoardID, FArticleID: tt.args.path.FArticleID}
			r0, _, _ := GetArticleDetail(tt.args.remoteAddr, user, nil, path0, nil)
			if r0 != nil {
				ret0 := r0.(*GetArticleDetailResult)
				expectedResult := tt.expectedResult.(*CreateRankResult)
				assert.Equal(t, expectedResult.Rank, ret0.Rank)
			}
		})
		wg.Wait()
	}
}

func TestCreateRank_boardNotReadable(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)

	board_b := *testBoardSummaryWhoAmI_b
	board_b.BrdAttr = ptttype.BRD_OVER18
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", []*bbs.BoardSummary{&board_b}, 123456890000000000)

	_ = schema.UpdateArticleSummaries([]*schema.ArticleSummary{{BBoardID: "10_WhoAmI", ArticleID: "1VrooM21", Owner: "test"}}, 1234567890000000000)

	path := &CreateRankPath{FBoardID: "WhoAmI", FArticleID: "M.1607937174.A.081"}

	// no over-18 cookie.
	user := &UserInfo{UserID: "chhsiao123", IsOver18: false}
	gotResult, statusCode, err := CreateRank(testIP, user, &CreateRankParams{Rank: 1}, path, nil)
	assert.Error(t, err)
	assert.Nil(t, gotResult)
	assert.Equal(t, 403, statusCode)

	// the rank must not be recorded.
	total, err := schema.GetArticleSummary("10_WhoAmI", "1VrooM21")
	assert.NoError(t, err)
	assert.Zero(t, total.Rank)
}
