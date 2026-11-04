package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/gin-gonic/gin"
)

func TestLoadFavoriteBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := NewLoadFavoriteBoardsParams()
	path0 := &LoadFavoriteBoardsPath{
		UserID: "SYSOP",
	}

	testBoardSummary9_c := &apitypes.BoardSummary{
		FBoardID:  "test9",
		Brdname:   "test9",
		Title:     "測試9",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     100,
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,

		LastPostTime: 1234567890,
		StatAttr:     ptttype.NBRD_BOARD,
		Idx:          "3",
		Gid:          3,
		Bid:          9,
		Fav:          true,

		URL: "/board/test9/articles",
	}

	testBoardSummary8_c := &apitypes.BoardSummary{
		FBoardID:  "test8",
		Brdname:   "test8",
		Title:     "測試8",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     101,
		BMs:       []bbs.UUserID{"okcool2", "teemo2"},
		Total:     124,

		LastPostTime: 1300000000,
		StatAttr:     ptttype.NBRD_BOARD,
		Idx:          "5",
		Gid:          3,
		Bid:          8,
		Fav:          true,

		URL: "/board/test8/articles",
	}

	testFavoriteBoards0_c := []*apitypes.BoardSummary{
		{StatAttr: ptttype.NBRD_LINE, Idx: "0", BMs: []bbs.UUserID{}},
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":1", Idx: "1", URL: "/user/SYSOP/favorites?level_idx=:1", BMs: []bbs.UUserID{}},
		{StatAttr: ptttype.NBRD_LINE, Idx: "2", BMs: []bbs.UUserID{}},
		testBoardSummary9_c,
		{Title: testTitle0, StatAttr: ptttype.NBRD_FOLDER, LevelIdx: ":4", Idx: "4", URL: "/user/SYSOP/favorites?level_idx=:4", BMs: []bbs.UUserID{}},
		testBoardSummary8_c,
	}
	result0 := &LoadFavoriteBoardsResult{
		List:    testFavoriteBoards0_c,
		NextIdx: "",

		TokenUser: "SYSOP",
	}

	params1 := &LoadFavoriteBoardsParams{
		StartIdx:  "2",
		Max:       2,
		Ascending: true,
	}

	result1 := &LoadFavoriteBoardsResult{
		List:    testFavoriteBoards0_c[2:4],
		NextIdx: "4",

		TokenUser: "SYSOP",
	}

	params2 := &LoadFavoriteBoardsParams{
		StartIdx:  "4",
		Max:       2,
		Ascending: false,
	}

	list2 := []*apitypes.BoardSummary{
		testFavoriteBoards0_c[4],
		testFavoriteBoards0_c[3],
	}
	result2 := &LoadFavoriteBoardsResult{
		List:    list2,
		NextIdx: "2",

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
		expectedResult     any
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			expectedResult:     result0,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params1, path: path0},
			expectedResult:     result1,
			expectedStatusCode: 200,
		},
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params2, path: path0},
			expectedResult:     result2,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := LoadFavoriteBoards(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFavoriteBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "result", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadFavoriteBoards() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
