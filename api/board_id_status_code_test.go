package api

import (
	"errors"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_boardIDErrToStatusCode(t *testing.T) {
	tests := []struct {
		err  error
		want int
	}{
		{err: mongo.ErrNoDocuments, want: 404},
		{err: ErrBoardNotPopular, want: 403},
		{err: schema.ErrInvalidPTTBID, want: 400},
		{err: errors.New("db down"), want: 500},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, boardIDErrToStatusCode(tt.err), "err: %v", tt.err)
	}
}

func TestBoardIDStatusCode_noBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	user := &UserInfo{UserID: "SYSOP"}
	fboardID := apitypes.FBoardID("NoSuchBoard")

	_, statusCode, err := LoadBottomArticles("localhost", user, nil, &LoadBottomArticlesPath{FBoardID: fboardID}, &gin.Context{})
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
	assert.Equal(t, 404, statusCode)

	_, statusCode, err = LoadGeneralArticles("localhost", user, NewLoadGeneralArticlesParams(), &LoadGeneralArticlesPath{FBoardID: fboardID}, &gin.Context{})
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
	assert.Equal(t, 404, statusCode)

	_, statusCode, err = LoadClassBoards("localhost", user, NewLoadClassBoardsParams(), &LoadClassBoardsPath{ClsID: 9999}, &gin.Context{})
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
	assert.Equal(t, 404, statusCode)
}

func TestBoardIDStatusCode_notPopular(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}, 123456890000000000)

	origIsAllGuest := types.IS_ALL_GUEST
	types.IS_ALL_GUEST = true
	defer func() { types.IS_ALL_GUEST = origIsAllGuest }()

	user := &UserInfo{UserID: "SYSOP"}
	_, statusCode, err := LoadBottomArticles("localhost", user, nil, &LoadBottomArticlesPath{FBoardID: "WhoAmI"}, &gin.Context{})
	assert.ErrorIs(t, err, ErrBoardNotPopular)
	assert.Equal(t, 403, statusCode)
}
