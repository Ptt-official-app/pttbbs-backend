package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		params     any
		path       any
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

// setupArticleWithComment creates an article and a comment owned by SYSOP, and registers another user chhsiao123.
func setupArticleWithComment(t *testing.T) (boardID bbs.BBoardID, articleID bbs.ArticleID, deletePath *DeleteCommentsPath, commentID types.CommentID) {
	t.Helper()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummarySYSOP_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}

	createArticleParams := &CreateArticleParams{
		PostType: "測試",
		Title:    "this is a test",
		Content: [][]*types.Rune{
			{{Utf8: "測試0", Color0: types.DefaultColor, Color1: types.DefaultColor}},
		},
	}
	gotCreateArticleResult, _, err := CreateArticle(testIP, user, createArticleParams, &CreateArticlePath{FBoardID: "SYSOP"}, nil)
	require.NoError(t, err)
	createArticle := gotCreateArticleResult.(CreateArticleResult)

	createCommentParams := &CreateCommentParams{
		CommentType: ptttype.COMMENT_TYPE_RECOMMEND,
		Content:     "test123",
	}
	createCommentPath := &CreateCommentPath{FBoardID: createArticle.FBoardID, FArticleID: createArticle.ArticleID}
	gotCreateCommentResult, _, err := CreateComment(testIP, user, createCommentParams, createCommentPath, nil)
	require.NoError(t, err)
	createComment := gotCreateCommentResult.(CreateCommentResult)

	boardID, err = toBoardID(createArticle.FBoardID, testIP, "SYSOP", nil)
	require.NoError(t, err)
	articleID = createArticle.ArticleID.ToArticleID()
	deletePath = &DeleteCommentsPath{FBoardID: createArticle.FBoardID, FArticleID: createArticle.ArticleID}

	return boardID, articleID, deletePath, createComment.CommentID
}

func TestDeleteComments_notOwner(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardID, articleID, path, commentID := setupArticleWithComment(t)

	params := &DeleteCommentsParams{TheList: []*apitypes.DeleteCommentParams{{CommentID: commentID, Reason: "x"}}}

	// a user who is not the article owner must be rejected, and the cached comment must stay untouched.
	other := &UserInfo{UserID: "chhsiao123", IsOver18: true}
	_, statusCode, err := DeleteComments(testIP, other, params, path, nil)
	assert.Error(t, err)
	assert.Equal(t, 403, statusCode)

	commentMap, err := schema.GetCommentMapByCommentIDs(boardID, articleID, []types.CommentID{commentID})
	require.NoError(t, err)
	require.Contains(t, commentMap, commentID)
	assert.Equal(t, ptttype.COMMENT_TYPE_RECOMMEND, commentMap[commentID].TheType)
}

func TestDeleteComments_unknownCommentID(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _, path, _ := setupArticleWithComment(t)

	params := &DeleteCommentsParams{TheList: []*apitypes.DeleteCommentParams{{CommentID: "no-such-cid", Reason: "x"}}}

	owner := &UserInfo{UserID: "SYSOP", IsOver18: true}
	assert.NotPanics(t, func() {
		_, statusCode, err := DeleteComments(testIP, owner, params, path, nil)
		assert.Error(t, err)
		assert.Equal(t, 400, statusCode)
	})
}
