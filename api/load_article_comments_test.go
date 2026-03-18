package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func TestLoadArticleComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = deserializeUserDetailAndUpdateDB(testUserSYSOP_b, 123456890000000000)
	_, _ = deserializeUserDetailAndUpdateDB(testUserChhsiao123_b, 123456891000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	articleSummaries_b := []*bbs.ArticleSummary{testArticleSummary1_b, testArticleSummary2_b}
	_, _, _ = deserializeArticlesAndUpdateDB("SYSOP", "10_WhoAmI", articleSummaries_b, 123456892000000000, false)

	articleParams := &GetArticleDetailParams{}
	articlePath := &GetArticleDetailPath{
		FBoardID:   apitypes.FBoardID("WhoAmI"),
		FArticleID: apitypes.FArticleID("M.1607937174.A.081"),
	}
	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	_, _, _ = GetArticleDetail(testIP, user, articleParams, articlePath, nil)

	comments := []*apitypes.Comment{
		{
			FBoardID:   "WhoAmI",
			FArticleID: "M.1607937174.A.081",
			CommentID:  "FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
			TheType:    2,
			CreateTime: 1608388624,
			SortTime:   1608388624,
			Owner:      "chhsiao123",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "噓～",
						Big5:   []byte("\xbcN\xa1\xe3                                                "),
						Color0: types.Color{Foreground: 37, Background: 40},
						Color1: types.Color{Foreground: 37, Background: 40},
						DBCS:   []byte("\xbcN\xa1\xe3                                                "),
					},
				},
			},
			Idx: "1608388624000000000@FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
		},
		{
			FBoardID:   "WhoAmI",
			FArticleID: "M.1607937174.A.081",
			CommentID:  "FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg",
			TheType:    1,
			CreateTime: 1608388560,
			SortTime:   1608388560,
			Owner:      "chhsiao123",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                  "),
						Color0: types.Color{Foreground: 37, Background: 40},
						Color1: types.Color{Foreground: 37, Background: 40},
						DBCS:   []byte("\xb1\xc0                                                  "),
					},
				},
			},
			Idx: "1608388560000000000@FlIk36uaIAA:FQaNH8WkdAbEGD7yp2Zkvg",
		},
		{
			FBoardID:   "WhoAmI",
			FArticleID: "M.1607937174.A.081",
			CommentID:  "FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ",
			TheType:    3,
			CreateTime: 1608388500,
			SortTime:   1608388500,
			Owner:      "SYSOP",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "推",
						Big5:   []byte("\xb1\xc0                                                       "),
						Color0: types.Color{Foreground: 37, Background: 40},
						Color1: types.Color{Foreground: 37, Background: 40},
						DBCS:   []byte("\xb1\xc0                                                       "),
					},
				},
			},
			Idx: "1608388500000000000@FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ",
		},
	}

	params0 := NewLoadArticleCommentsParams()
	path0 := &LoadArticleCommentsPath{
		FBoardID:   "WhoAmI",
		FArticleID: "M.1607937174.A.081",
	}
	expected0 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
			comments[1],
			comments[2],
		},

		TokenUser: "SYSOP",
	}

	params1 := &LoadArticleCommentsParams{
		Descending: true,
		Max:        2,
	}
	expected1 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
			comments[1],
		},
		NextIdx: "1608388500000000000@FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ",

		TokenUser: "SYSOP",
	}

	params2 := &LoadArticleCommentsParams{
		StartIdx:   "1608388500000000000@FlIk0bNSyAA:3dK46zmOe5zmna12AC1gnQ",
		Descending: true,
		Max:        2,
	}
	expected2 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[2],
		},

		TokenUser: "SYSOP",
	}

	params3 := &LoadArticleCommentsParams{
		Descending: false,
		Max:        200,
	}
	expected3 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[2],
			comments[1],
			comments[0],
		},

		TokenUser: "SYSOP",
	}

	params4 := &LoadArticleCommentsParams{
		Descending: false,
		Max:        2,
	}
	expected4 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[2],
			comments[1],
		},
		NextIdx: "1608388624000000000@FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",

		TokenUser: "SYSOP",
	}

	params5 := &LoadArticleCommentsParams{
		StartIdx:   "1608388624000000000@FlIk7pJMoAA:cLGi8fC4fapuiBkTXHU2OA",
		Descending: false,
		Max:        2,
	}
	expected5 := &LoadArticleCommentsResult{
		List: []*apitypes.Comment{
			comments[0],
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
		expectedResult     *LoadArticleCommentsResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			expectedResult:     expected0,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params1, path: path0},
			expectedResult:     expected1,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params2, path: path0},
			expectedResult:     expected2,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params3, path: path0},
			expectedResult:     expected3,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params4, path: path0},
			expectedResult:     expected4,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params5, path: path0},
			expectedResult:     expected5,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := LoadArticleComments(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadArticleComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)

			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadArticleComments() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
