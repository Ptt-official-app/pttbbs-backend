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
)

func TestLoadPopularBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserBoard_c.Drop()

	update0 := &schema.UserBoard{UserID: "SYSOP", BBoardID: "1_test1", ReadUpdateNanoTS: types.Time8(1234567891).ToNanoTS()}
	update1 := &schema.UserBoard{UserID: "SYSOP", BBoardID: "2_test2", ReadUpdateNanoTS: types.Time8(1234567891).ToNanoTS()}

	_, _ = schema.UserBoard_c.Update(update0, update0)
	_, _ = schema.UserBoard_c.Update(update1, update1)

	expectedResult := &LoadPopularBoardsResult{
		List: []*apitypes.BoardSummary{
			{
				FBoardID:  "WhoAmI",
				Brdname:   "WhoAmI",
				Title:     "呵呵，猜猜我是誰！",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "嘰哩",
				NUser:     0,
				BMs:       []bbs.UUserID{},
				Read:      false,
				Total:     0,

				LastPostTime: 0,
				StatAttr:     ptttype.NBRD_BOARD,

				Gid: 5,
				Bid: 10,
				URL: "/board/WhoAmI/articles",
			},
			{
				FBoardID:  "SYSOP",
				Brdname:   "SYSOP",
				Title:     "站長好!",
				BrdAttr:   32,
				BoardType: "◎",
				Category:  "嘰哩",
				NUser:     0,
				BMs:       []bbs.UUserID{},
				Read:      false,
				Total:     0,

				LastPostTime: 0,
				StatAttr:     ptttype.NBRD_BOARD,

				Gid: 2,
				Bid: 1,

				Fav: true,
				URL: "/board/SYSOP/articles",
			},
			{
				FBoardID:  "Note",
				Brdname:   "Note",
				Title:     "動態看板及歌曲投稿",
				BrdAttr:   0,
				BoardType: "◎",
				Category:  "嘰哩",
				NUser:     0,
				BMs:       []bbs.UUserID{},
				Read:      false,
				Total:     0,

				LastPostTime: 0,
				StatAttr:     ptttype.NBRD_BOARD,

				Gid: 5,
				Bid: 8,
				Fav: true,
				URL: "/board/Note/articles",
			},
		},

		TokenUser: "SYSOP",
	}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     any
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
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: nil, c: nil},
			expectedResult:     expectedResult,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := LoadPopularBoards(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPopularBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			testutil.TDeepEqual(t, "result", gotResult, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadPopularBoards() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
