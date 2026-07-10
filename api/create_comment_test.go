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

func TestCreateComment(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	params0 := &CreateCommentParams{
		CommentType: ptttype.COMMENT_TYPE_RECOMMEND,
	}
	path0 := &CreateCommentPath{
		FBoardID:   "WhoAmI",
		FArticleID: "M.1607937174.A.082",
	}
	expected0 := &apitypes.Comment{
		FBoardID:   "WhoAmI",
		FArticleID: "M.1607937174.A.082",
		TheType:    ptttype.COMMENT_TYPE_RECOMMEND,
		Owner:      "SYSOP",
		CommentID:  "FoL-3rSuO0A:PlQBjivX5w0hPmg3SIsGjw",
		CreateTime: 1607937324,
		SortTime:   1622139048,
		Content: [][]*types.Rune{
			{
				{
					Utf8:   "test123",
					Color0: types.DefaultColor,
					Color1: types.DefaultColor,
					Big5:   []byte("test123                                                  "),
					DBCS:   []byte("test123                                                  "),
				},
			},
		},
		Idx: "1622139048149857000@FofoE7wzVug:PlQBjivX5w0hPmg3SIsGjw",

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
		expectedResult     CreateCommentResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				remoteAddr: testIP,
				userID:     "SYSOP",
				params:     params0,
				path:       path0,
				c:          nil,
			},
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
			gotResult, gotStatusCode, err := CreateComment(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, _ := gotResult.(CreateCommentResult)
			got.SortTime = tt.expectedResult.SortTime
			got.CommentID = tt.expectedResult.CommentID
			got.Idx = tt.expectedResult.Idx
			testutil.TDeepEqual(t, "got", got, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CreateComment() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
