package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/boardd"
	"github.com/Ptt-official-app/pttbbs-backend/mockhttp"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestUpdateBoardSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralBoards(nil)

	updateNanoTS := types.NowNanoTS()

	boardSummaries0 := make([]*BoardSummary, len(ret.Boards))
	for idx, each_b := range ret.Boards {
		boardSummaries0[idx] = NewBoardSummary(each_b, updateNanoTS)
	}

	query0 := &BoardQuery{BBoardID: "1_test1"}
	boardSummary0 := &BoardSummary{
		BBoardID:  "1_test1",
		Brdname:   "test1",
		Title:     "測試1",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,
		NUser:     100,

		LastPostTime: types.NanoTS(1234567890000000000),

		UpdateNanoTS: updateNanoTS,

		IsOver18UpdateNanoTS: updateNanoTS,

		Gid: 3,
		Bid: 1,

		IdxByName:  "test1",
		IdxByClass: "tPq41Q@test1",
	}

	query1 := &BoardQuery{BBoardID: "2_test2"}
	boardSummary1 := &BoardSummary{
		BBoardID:  "2_test2",
		Brdname:   "test2",
		Title:     "測試2",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool2", "teemo2"},
		Total:     124,
		NUser:     101,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS,

		IsOver18UpdateNanoTS: updateNanoTS,

		Gid: 3,
		Bid: 2,

		IdxByName:  "test2",
		IdxByClass: "tPq41Q@test2",
	}

	boardSummary2 := &BoardSummary{
		BBoardID: bbs.BBoardID("3_test3"),

		Brdname:   "test3",
		Title:     "測試3",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "CPBL",
		BMs:       []bbs.UUserID{"okcool3", "teemo3"},
		Total:     125,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS,

		IsOver18UpdateNanoTS: updateNanoTS,

		Gid: 3,
		Bid: 3,

		IdxByName:  "test3",
		IdxByClass: "tPq41Q@test3",
	}

	query2 := &BoardQuery{BBoardID: "3_test3"}
	boardSummaries1 := []*BoardSummary{boardSummary1, boardSummary2}

	boardSummary3 := &BoardSummary{
		BBoardID: bbs.BBoardID("2_test2"),

		Brdname:   "test2",
		Title:     "測試4",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "CPBL",
		BMs:       []bbs.UUserID{"okcool4", "teemo4"},
		Total:     125,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS,

		IsOver18UpdateNanoTS: updateNanoTS,

		Gid: 3,
		Bid: 2,

		IdxByName:  "test2",
		IdxByClass: "tPq41Q@test2",
	}
	boardSummaries2 := []*BoardSummary{boardSummary2, boardSummary3}

	updateNanoTS1 := types.NowNanoTS()
	boardSummary4 := &BoardSummary{
		BBoardID: bbs.BBoardID("2_test2"),

		Brdname:   "test2",
		Title:     "測試4",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "CPBL",
		BMs:       []bbs.UUserID{"okcool4", "teemo4"},
		Total:     125,

		LastPostTime: types.NanoTS(1300000000000000000),

		UpdateNanoTS: updateNanoTS1,

		IsOver18UpdateNanoTS: updateNanoTS,

		Gid: 3,
		Bid: 2,

		IdxByName:  "test2",
		IdxByClass: "tPq41Q@test2",
	}
	boardSummaries3 := []*BoardSummary{boardSummary2, boardSummary4}

	type args struct {
		boardSummaries []*BoardSummary
		updateNanoTS   types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		query    *BoardQuery
		expected *BoardSummary
	}{
		// TODO: Add test cases.
		{
			args:     args{boardSummaries: boardSummaries0, updateNanoTS: updateNanoTS},
			query:    query0,
			expected: boardSummary0,
		},
		{
			args:     args{boardSummaries: boardSummaries0, updateNanoTS: updateNanoTS},
			query:    query1,
			expected: boardSummary1,
		},
		{
			args:     args{boardSummaries: boardSummaries1, updateNanoTS: updateNanoTS},
			query:    query2,
			expected: boardSummary2,
		},
		{
			args:     args{boardSummaries: boardSummaries2, updateNanoTS: updateNanoTS},
			query:    query1,
			expected: boardSummary1,
		},
		{
			args:     args{boardSummaries: boardSummaries3, updateNanoTS: updateNanoTS1},
			query:    query1,
			expected: boardSummary4,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			if err := UpdateBoardSummaries(tt.args.boardSummaries, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBoardSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := GetBoardSummary(tt.query.BBoardID)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}

func TestGetBoardSummariesByClsID(t *testing.T) {
	setupTest()
	defer teardownTest()

	ret := mockhttp.LoadGeneralBoards(nil)

	updateNanoTS := types.NowNanoTS()

	boardSummaries0 := make([]*BoardSummary, len(ret.Boards))
	for idx, each_b := range ret.Boards {
		boardSummaries0[idx] = NewBoardSummary(each_b, updateNanoTS)
	}

	_ = UpdateBoardSummaries(boardSummaries0, updateNanoTS)

	boardSummary0 := &BoardSummary{
		BBoardID:  "1_test1",
		Brdname:   "test1",
		Title:     "測試1",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,
		NUser:     100,

		LastPostTime: types.NanoTS(1234567890000000000),

		IsOver18UpdateNanoTS: updateNanoTS,

		UpdateNanoTS: updateNanoTS,
		Gid:          3,
		Bid:          1,

		IdxByName:  "test1",
		IdxByClass: "tPq41Q@test1",
	}

	boardSummary1 := &BoardSummary{
		BBoardID:  "2_test2",
		Brdname:   "test2",
		Title:     "測試2",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		BMs:       []bbs.UUserID{"okcool2", "teemo2"},
		Total:     124,
		NUser:     101,

		LastPostTime: types.NanoTS(1300000000000000000),

		IsOver18UpdateNanoTS: updateNanoTS,

		UpdateNanoTS: updateNanoTS,
		Gid:          3,
		Bid:          2,

		IdxByName:  "test2",
		IdxByClass: "tPq41Q@test2",
	}

	expected0 := []*BoardSummary{boardSummary0, boardSummary1}
	expected1 := []*BoardSummary{boardSummary0}

	type args struct {
		clsID    ptttype.Bid
		startIdx string
		isAsc    bool
		limit    int
		sortBy   ptttype.BSortBy
	}
	tests := []struct {
		name                   string
		args                   args
		expectedBoardSummaries []*BoardSummary
		wantErr                bool
	}{
		// TODO: Add test cases.
		{
			args:                   args{clsID: 3, limit: 100, sortBy: ptttype.BSORT_BY_CLASS, isAsc: true},
			expectedBoardSummaries: expected0,
		},
		{
			args:                   args{clsID: 3, limit: 100, startIdx: "test1", sortBy: ptttype.BSORT_BY_NAME, isAsc: true},
			expectedBoardSummaries: expected0,
		},
		{
			args:                   args{clsID: 3, limit: 100, startIdx: "test1", sortBy: ptttype.BSORT_BY_NAME, isAsc: false},
			expectedBoardSummaries: expected1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotBoardSummaries, err := GetBoardSummariesByClsID(tt.args.clsID, tt.args.startIdx, tt.args.isAsc, tt.args.limit, tt.args.sortBy)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardSummariesByClsID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotBoardSummaries, tt.expectedBoardSummaries)
		})
		wg.Wait()
	}
}

func TestNewBoardSummaryFromPBBoard(t *testing.T) {
	board0 := &boardd.Board{
		RawModerators: "okcool/teemo",
		Name:          "test",
		Bclass:        "測試",
		Bid:           7,
		Title:         "這裡是測試板",
		NumUsers:      12345,
		NumPosts:      234,
		Parent:        1,
	}
	expected0 := &BoardSummary{
		BBoardID: "7_test",
		Brdname:  "test",
		Title:    "這裡是測試板",
		Category: "測試",
		BMs:      []bbs.UUserID{"okcool", "teemo"},
		Total:    234,
		NUser:    12345,

		ParentID: "1_test1",

		Gid:          1,
		Bid:          7,
		IdxByName:    "test",
		IdxByClass:   "tPq41Q@test",
		BoardType:    "◎",
		UpdateNanoTS: 1234567890000000000,

		IsOver18UpdateNanoTS: 1234567890000000000,
	}

	type args struct {
		b_b          *boardd.Board
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		expected *BoardSummary
	}{
		// TODO: Add test cases.
		{
			args:     args{b_b: board0, updateNanoTS: 1234567890000000000},
			expected: expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			got := NewBoardSummaryFromPBBoard(tt.args.b_b, tt.args.updateNanoTS)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
		wg.Wait()
	}
}
