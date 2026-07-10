package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/mockhttp"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func TestLoadClassBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)

	ret := mockhttp.LoadGeneralBoards(nil)
	updateNanoTS := types.NowNanoTS()

	boardSummaries0 := make([]*schema.BoardSummary, len(ret.Boards))
	for idx, each_b := range ret.Boards {
		boardSummaries0[idx] = schema.NewBoardSummary(each_b, updateNanoTS)
	}

	update0 := &schema.UserBoard{UserID: "SYSOP", BBoardID: "1_test1", ReadUpdateNanoTS: types.Time8(1234567891).ToNanoTS()}
	update1 := &schema.UserBoard{UserID: "SYSOP", BBoardID: "2_test2", ReadUpdateNanoTS: types.Time8(1234567891).ToNanoTS()}

	_, _ = schema.UserBoard_c.Update(update0, update0)
	_, _ = schema.UserBoard_c.Update(update1, update1)

	_ = schema.UpdateBoardSummaries(boardSummaries0, updateNanoTS)

	classSummary0 := &schema.BoardSummary{
		BBoardID:  "3_3..........",
		Brdname:   "3..........",
		Title:     "市民廣場     報告  站長  ㄜ！",
		BoardType: "Σ",
		Category:  "....",

		Gid:        1,
		Bid:        3,
		IdxByName:  "3..........",
		IdxByClass: "....@..........",
	}

	_ = schema.UpdateBoardSummaries([]*schema.BoardSummary{classSummary0}, updateNanoTS)

	params0 := &LoadClassBoardsParams{
		SortBy:    ptttype.BSORT_BY_NAME,
		Ascending: true,
		Max:       10,
	}

	path0 := &LoadClassBoardsPath{
		ClsID: 3,
	}

	expected0 := &LoadClassBoardsResult{
		List: []*apitypes.BoardSummary{
			{
				FBoardID:  "test1",
				Brdname:   "test1",
				Title:     "測試1",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "測試",
				NUser:     100,
				BMs:       []bbs.UUserID{"okcool", "teemo"},
				Read:      true,
				Total:     123,

				LastPostTime: 1234567890,
				StatAttr:     ptttype.NBRD_BOARD,

				Gid: 3,
				Bid: 1,

				Idx: "test1",
				URL: "/board/test1/articles",
			},
			{
				FBoardID:  "test2",
				Brdname:   "test2",
				Title:     "測試2",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "測試",
				NUser:     101,
				BMs:       []bbs.UUserID{"okcool2", "teemo2"},
				Read:      false,
				Total:     124,

				LastPostTime: 1300000000,
				StatAttr:     ptttype.NBRD_BOARD,

				Gid: 3,
				Bid: 2,
				Idx: "test2",
				URL: "/board/test2/articles",
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
		expectedResult     interface{}
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
			gotResult, gotStatusCode, err := LoadClassBoards(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadClassBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadClassBoards() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}

func Test_loadClassBoards(t *testing.T) {
	setupTest()
	defer teardownTest()
	type args struct {
		userID   bbs.UUserID
		clsID    ptttype.Bid
		startIdx string
		asc      bool
		max      int
		sortBy   ptttype.BSortBy
		c        *gin.Context
	}
	tests := []struct {
		name                      string
		args                      args
		expectedBoardSummaries_db []*schema.BoardSummary
		expectedUserBoardInfoMap  map[bbs.BBoardID]*apitypes.UserBoardInfo
		expectedNextIdx           string
		wantErr                   bool
	}{
		// TODO: Add test cases.
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotBoardSummaries_db, gotUserBoardInfoMap, gotNextIdx, err := loadClassBoards(tt.args.userID, tt.args.clsID, tt.args.startIdx, tt.args.asc, tt.args.max, tt.args.sortBy, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadClassBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "boardSummaries", gotBoardSummaries_db, tt.expectedBoardSummaries_db)
			testutil.TDeepEqual(t, "userBoardInfo", gotUserBoardInfoMap, tt.expectedUserBoardInfoMap)
			if gotNextIdx != tt.expectedNextIdx {
				t.Errorf("loadClassBoards() gotNextIdx = %v, want %v", gotNextIdx, tt.expectedNextIdx)
			}
		})
		wg.Wait()
	}
}
