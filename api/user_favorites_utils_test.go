package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Test_checkUserFavBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	userBoardInfoMap0 := map[bbs.BBoardID]*apitypes.UserBoardInfo{
		"9_test9":   {},
		"10_WhoAmI": {},
	}
	boardSummaries0 := []*schema.BoardSummary{
		{
			BBoardID:     "9_test9",
			Brdname:      "test9",
			Title:        "測試9",
			LastPostTime: 1234567890000000000,
			UpdateNanoTS: 1234567890000000000,
			Gid:          3,
			Bid:          9,
			IdxByName:    "test9",
		},
		{
			BBoardID:     "10_WhoAmI",
			Brdname:      "WhoAmI",
			Title:        "猜猜我是誰",
			LastPostTime: 1234567890000000000,
			UpdateNanoTS: 1234567890000000000,
			Gid:          5,
			Bid:          10,
			IdxByName:    "WhoAmI",
		},
	}
	wantUserBoardInfoMap0 := map[bbs.BBoardID]*apitypes.UserBoardInfo{
		"9_test9":   {Fav: true},
		"10_WhoAmI": {},
	}

	type args struct {
		userID           bbs.UUserID
		userBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo
		boardSummaries   []*schema.BoardSummary
		c                *gin.Context
	}
	tests := []struct {
		name                    string
		args                    args
		wantNewUserBoardInfoMap map[bbs.BBoardID]*apitypes.UserBoardInfo
		wantErr                 bool
	}{
		// TODO: Add test cases.
		{
			args:                    args{userID: "SYSOP", userBoardInfoMap: userBoardInfoMap0, boardSummaries: boardSummaries0},
			wantNewUserBoardInfoMap: wantUserBoardInfoMap0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotNewUserBoardInfoMap, err := checkUserFavBoard(tt.args.userID, tt.args.userBoardInfoMap, tt.args.boardSummaries, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkUserFavBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			logrus.Infof("got: %v want: %v", gotNewUserBoardInfoMap["10_WhoAmI"], tt.wantNewUserBoardInfoMap["10_WhoAmI"])
			testutil.TDeepEqual(t, "got", gotNewUserBoardInfoMap, tt.wantNewUserBoardInfoMap)
		})
		wg.Wait()
	}
}

func Test_tryGetUserFavoritesCore_noFavorites(t *testing.T) {
	setupTest()
	defer teardownTest()

	// user without favorites meta in db and backend mtime 0 must not panic.
	gotIdx, gotMTime, gotStatus, err := tryGetUserFavoritesCore("NoFavUser", nil)
	if err != nil {
		t.Fatalf("tryGetUserFavoritesCore() error = %v", err)
	}
	if gotIdx != 0 || gotMTime != 0 || gotStatus != 200 {
		t.Errorf("tryGetUserFavoritesCore() = (%v, %v, %v), want (0, 0, 200)", gotIdx, gotMTime, gotStatus)
	}
}
