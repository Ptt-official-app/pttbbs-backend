package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestGetBoardIsPopular(t *testing.T) {
	setupTest()
	defer teardownTest()

	updateNanoTS := types.NowNanoTS()

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

		UpdateNanoTS: updateNanoTS,
		Gid:          3,
		Bid:          2,

		IdxByName:  "test2",
		IdxByClass: "tPq41Q@test2",
		IsPopular:  true,
	}

	UpdateBoardSummaries([]*BoardSummary{boardSummary0, boardSummary1}, updateNanoTS)

	boardID0 := bbs.BBoardID("1_test1")
	want0 := &BoardIsPopular{
		BBoardID:  "1_test1",
		IsPopular: false,
	}

	boardID1 := bbs.BBoardID("2_test2")
	want1 := &BoardIsPopular{
		BBoardID:  "2_test2",
		IsPopular: true,
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		boardID bbs.BBoardID
		want    *BoardIsPopular
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			boardID: boardID0,
			want:    want0,
		},
		{
			boardID: boardID1,
			want:    want1,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			got, gotErr := GetBoardIsPopular(tt.boardID)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetBoardIsPopular() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("GetBoardIsPopular() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			testutil.TDeepEqual(t, "got", got, tt.want)
		})
	}
}
