package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestAddFavoriteBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	boardSummaries_b := []*bbs.BoardSummary{testBoardSummaryWhoAmI_b}
	_, _, err := deserializeBoardsAndUpdateDB("SYSOP", boardSummaries_b, 123456890000000000)
	if err != nil {
		logrus.Errorf("TestAddFavoriteBoard: unable to deserializeBoardsAndUpdateDB: e: %v", err)
	}

	paramsLoad0 := &LoadGeneralBoardsParams{
		StartIdx: "vFSt-Q@WhoAmI",
	}

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	result_i, _, err := LoadGeneralBoardsByClass("localhost", user, paramsLoad0, nil)
	result, _ := result_i.(*LoadGeneralBoardsResult)
	logrus.Infof("TestAddFavoriteBoard: after LoadGeneralBoardsByClass: result: %v", result.List[0])

	params0 := &AddFavoriteBoardParams{
		FBoardID: "WhoAmI",
	}

	path0 := &AddFavoriteBoardPath{
		UserID: "SYSOP",
	}

	ret0 := AddFavoriteResult(testBoardSummary10)

	params1 := &AddFavoriteBoardParams{
		FBoardID: "WhoAmI",
		LevelIdx: ":1",
	}

	ret1 := AddFavoriteResult(&apitypes.BoardSummary{
		FBoardID:  "WhoAmI",
		Brdname:   "WhoAmI",
		Title:     "呵呵，猜猜我是誰！",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "嘰哩",
		NUser:     0,
		BMs:       []bbs.UUserID{},
		Total:     0,

		LastPostTime: 0,
		StatAttr:     ptttype.NBRD_FAV,
		Idx:          "1",
		Gid:          5,
		Bid:          10,

		URL: "/board/WhoAmI/articles",

		TokenUser: "SYSOP",
	})

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     interface{}
		path       interface{}
		c          *gin.Context
	}
	tests := []struct {
		name           string
		args           args
		wantResult     AddFavoriteResult
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{params: params0, path: path0, userID: "SYSOP", remoteAddr: "localhost"},
			wantResult:     ret0,
			wantStatusCode: 200,
		},
		{
			args:           args{params: params0, path: path0, userID: "SYSOP", remoteAddr: "localhost"},
			wantResult:     ret0,
			wantStatusCode: 200,
		},
		{
			args:           args{params: params1, path: path0, userID: "SYSOP", remoteAddr: "localhost"},
			wantResult:     ret1,
			wantStatusCode: 200,
		},
		{
			args:           args{params: params1, path: path0, userID: "SYSOP", remoteAddr: "localhost"},
			wantResult:     ret1,
			wantStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := AddFavoriteBoard(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddFavoriteBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, _ := gotResult.(AddFavoriteResult)
			testutil.TDeepEqual(t, "got", got, tt.wantResult)
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("AddFavoriteBoard() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
		wg.Wait()
	}
}
