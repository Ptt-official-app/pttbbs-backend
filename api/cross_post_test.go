package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func TestCrossPost(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b, testBoardSummarySYSOP_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	path0 := &CrossPostPath{
		FBoardID:   "SYSOP",
		FArticleID: "M.1608386280.A.BC9",
	}

	params0 := &CrossPostParams{
		XBoardID: "WhoAmI",
	}

	expected0 := &CrossPostResult{
		Article: &apitypes.ArticleSummary{
			FBoardID:    "WhoAmI",
			ArticleID:   "M.1607937174.A.081",
			CreateTime:  1607937174,
			MTime:       1607937100,
			Owner:       "SYSOP",
			Title:       "this is a test",
			Class:       "轉",
			URL:         "http://localhost:5173/board/WhoAmI/article/M.1607937174.A.081",
			SubjectType: ptttype.SUBJECT_FORWARD,
		},

		Comment: &apitypes.Comment{
			FBoardID:   "SYSOP",
			FArticleID: "M.1608386280.A.BC9",
			TheType:    ptttype.COMMENT_TYPE_FORWARD,
			CreateTime: 1607937174,
			Owner:      "SYSOP",
			Content: [][]*types.Rune{
				{
					{
						Utf8:   "WhoAmI",
						Big5:   []byte("WhoAmI"),
						DBCS:   []byte("WhoAmI"),
						Color0: types.DefaultColor,
						Color1: types.DefaultColor,
					},
				},
			},
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
		expectedResult     *CrossPostResult
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
			gotResult, gotStatusCode, err := CrossPost(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrossPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			result, _ := gotResult.(*CrossPostResult)
			result.Comment.CommentID = ""
			result.Comment.SortTime = 0
			result.Comment.Idx = ""
			testutil.TDeepEqual(t, "got", result, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CrossPost() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
