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
	"github.com/sirupsen/logrus"
)

func TestDeleteComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummarySYSOP_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	createArticlePath0 := &CreateArticlePath{
		FBoardID: "SYSOP",
	}

	createArticleRune0 := &types.Rune{
		Utf8:   "測試0",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	createArticleRune1 := &types.Rune{
		Utf8:   "測試1",
		Color0: types.DefaultColor,
		Color1: types.DefaultColor,
	}

	createArticleParams0 := &CreateArticleParams{
		PostType: "測試",
		Title:    "this is a test",
		Content: [][]*types.Rune{
			{createArticleRune0, createArticleRune1},
		},
	}

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}

	gotCreateArticleResult0, _, _ := CreateArticle(testIP, user, createArticleParams0, createArticlePath0, nil)
	gotCreateArticle0, _ := gotCreateArticleResult0.(CreateArticleResult)

	logrus.Infof("TestDeleteComments: gotCreateArticle0: %v", gotCreateArticle0)

	createCommentParams0 := &CreateCommentParams{
		CommentType: ptttype.COMMENT_TYPE_RECOMMEND,
		Content:     "test123",
	}
	createCommentPath0 := &CreateCommentPath{
		FBoardID:   gotCreateArticle0.FBoardID,
		FArticleID: gotCreateArticle0.ArticleID,
	}

	gotCreateCommentResult0, _, err := CreateComment(testIP, user, createCommentParams0, createCommentPath0, nil)
	createCommentResult0, _ := gotCreateCommentResult0.(CreateCommentResult)
	logrus.Infof("createCommentResult: %v gotCreateCommentResult0: %v e: %v", createCommentResult0, gotCreateCommentResult0, err)

	params0 := &DeleteCommentsParams{
		TheList: []*apitypes.DeleteCommentParams{
			{
				CommentID: createCommentResult0.CommentID,
				Reason:    "誤植",
			},
		},
	}
	path0 := &DeleteCommentsPath{
		FBoardID:   gotCreateArticle0.FBoardID,
		FArticleID: gotCreateArticle0.ArticleID,
	}

	expected0 := &DeleteCommentsResult{
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
		name               string
		args               args
		expectedResult     *DeleteCommentsResult
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
			gotResult, gotStatusCode, err := DeleteComments(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			logrus.Infof("TestDeleteComments: got: %v statusCode: %v e: %v", gotResult, gotStatusCode, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("DeleteComments() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
