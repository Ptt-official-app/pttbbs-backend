package schema

import (
	"cmp"
	"reflect"
	"slices"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestUpdateRank(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		boardID      bbs.BBoardID
		articleID    bbs.ArticleID
		ownerID      bbs.UUserID
		rank         int
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name          string
		args          args
		expectedTotal int
		expectedOrig  int
		wantErr       bool
	}{
		// TODO: Add test cases.
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner", rank: 1, updateNanoTS: 1234567890000000000},
			expectedTotal: 1,
			expectedOrig:  0,
		},
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner", rank: 1, updateNanoTS: 1234567890000000000},
			wantErr:       true,
			expectedTotal: 1,
			expectedOrig:  0,
		},
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner", rank: -1, updateNanoTS: 1234567890000000001},
			expectedTotal: -1,
			expectedOrig:  1,
		},
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner1", rank: 1, updateNanoTS: 1234567890000000000},
			expectedTotal: 0,
			expectedOrig:  0,
		},
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner2", rank: 1, updateNanoTS: 1234567890000000000},
			expectedTotal: 1,
			expectedOrig:  0,
		},
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner3", rank: 1, updateNanoTS: 1234567890000000000},
			expectedTotal: 2,
			expectedOrig:  0,
		},
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner4", rank: 1, updateNanoTS: 1234567890000000000},
			expectedTotal: 3,
			expectedOrig:  0,
		},
		{
			args:          args{boardID: "test_board", articleID: "test_article", ownerID: "test_owner4", rank: -1, updateNanoTS: 1234567890000000001},
			expectedTotal: 1,
			expectedOrig:  1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			origRank, err := UpdateRank(tt.args.boardID, tt.args.articleID, tt.args.ownerID, tt.args.rank, tt.args.updateNanoTS)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRank() error = %v, wantErr %v", err, tt.wantErr)
			}
			if origRank != tt.expectedOrig {
				t.Errorf("UpdateRank() origRank: %v expected: %v", origRank, tt.expectedOrig)
			}

			total, err := SumRank(tt.args.boardID, tt.args.articleID)
			if err != nil {
				t.Errorf("UpdateRank: unable to sum: e: %v", err)
			}
			if total != tt.expectedTotal {
				t.Errorf("UpdateRank: total: %v expected: %v", total, tt.expectedTotal)
			}
		})
		wg.Wait()
	}
}

func TestSumRankByBoardID(t *testing.T) {
	setupTest()
	defer teardownTest()

	UpdateRank("test_board", "test_article", "test_owner1", 1, 1234567890000000000)
	UpdateRank("test_board", "test_article", "test_owner2", 1, 1234567890000000000)
	UpdateRank("test_board", "test_article", "test_owner3", -1, 1234567890000000000)
	UpdateRank("test_board", "test_article", "test_owner4", 1, 1234567890000000000)

	UpdateRank("test_board", "test_article1", "test_owner1", -1, 1234567890000000000)
	UpdateRank("test_board", "test_article1", "test_owner2", -1, 1234567890000000000)
	UpdateRank("test_board", "test_article1", "test_owner3", -1, 1234567890000000000)
	UpdateRank("test_board", "test_article1", "test_owner4", 1, 1234567890000000000)

	UpdateRank("test_board", "test_article2", "test_owner1", 1, 1234567890000000000)
	UpdateRank("test_board", "test_article2", "test_owner2", 1, 1234567890000000000)
	UpdateRank("test_board", "test_article2", "test_owner3", 1, 1234567890000000000)
	UpdateRank("test_board", "test_article2", "test_owner4", 1, 1234567890000000000)

	type args struct {
		boardID    bbs.BBoardID
		articleIDs []bbs.ArticleID
	}
	tests := []struct {
		name        string
		args        args
		expectedRet []*RankAgged
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			args: args{boardID: "test_board"},
			expectedRet: []*RankAgged{
				{BBoardID: "test_board", ArticleID: "test_article", Rank: 2},
				{BBoardID: "test_board", ArticleID: "test_article1", Rank: -2},
				{BBoardID: "test_board", ArticleID: "test_article2", Rank: 4},
			},
		},
		{
			args: args{boardID: "test_board", articleIDs: []bbs.ArticleID{"test_article1", "test_article2"}},
			expectedRet: []*RankAgged{
				{BBoardID: "test_board", ArticleID: "test_article1", Rank: -2},
				{BBoardID: "test_board", ArticleID: "test_article2", Rank: 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := SumRankByBoardID(tt.args.boardID, tt.args.articleIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumRankByBoardID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			slices.SortStableFunc(gotRet, func(a, b *RankAgged) int {
				return cmp.Compare(a.ArticleID, b.ArticleID)
			})
			if !reflect.DeepEqual(gotRet, tt.expectedRet) {
				t.Errorf("SumRankByBoardID() = %v, want %v", gotRet, tt.expectedRet)
			}
		})
	}
}
