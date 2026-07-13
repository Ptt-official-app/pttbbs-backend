package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

func Test_bidToBoardID(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, _ = deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 1234567890000000000)

	type args struct {
		bid ptttype.Bid
	}
	tests := []struct {
		name            string
		args            args
		expectedBoardID bbs.BBoardID
		wantErr         bool
	}{
		// TODO: Add test cases.
		{
			args:            args{bid: 10},
			expectedBoardID: "10_WhoAmI",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotBoardID, err := bidToBoardID(tt.args.bid)
			if (err != nil) != tt.wantErr {
				t.Errorf("bidToBoardID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoardID, tt.expectedBoardID) {
				t.Errorf("bidToBoardID() = %v, want %v", gotBoardID, tt.expectedBoardID)
			}
		})
		wg.Wait()
	}
}

func Test_isBoardSummariesValidUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	boardSummaries_db, _, _ := deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 1234567890000000000)

	type args struct {
		boardSummaries []*schema.BoardSummary
		c              *gin.Context
	}
	tests := []struct {
		name                        string
		args                        args
		expectedValidBoardSummaries []*schema.BoardSummary
		wantErr                     bool
	}{
		// TODO: Add test cases.
		{
			args:                        args{boardSummaries: boardSummaries_db},
			expectedValidBoardSummaries: boardSummaries_db,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotValidBoardSummaries, err := isBoardSummariesValidUser(tt.args.boardSummaries, "SYSOP")
			if (err != nil) != tt.wantErr {
				t.Errorf("isBoardSummariesValidUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotValidBoardSummaries, tt.expectedValidBoardSummaries)
		})
		wg.Wait()
	}
}
