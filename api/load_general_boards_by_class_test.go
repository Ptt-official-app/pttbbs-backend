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

func TestLoadGeneralBoardsByClass(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserBoard_c.Drop()

	update0 := &schema.UserBoard{UserID: "SYSOP", BBoardID: "1_test1", ReadUpdateNanoTS: types.Time8(1234567891).ToNanoTS()}
	update1 := &schema.UserBoard{UserID: "SYSOP", BBoardID: "2_test2", ReadUpdateNanoTS: types.Time8(1234567891).ToNanoTS()}

	_, _ = schema.UserBoard_c.Update(update0, update0)
	_, _ = schema.UserBoard_c.Update(update1, update1)

	params := &LoadGeneralBoardsParams{}
	expectedResult := &LoadGeneralBoardsResult{
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
				URL:          "/board/test1/articles",
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
				URL:          "/board/test2/articles",
			},
		},
		NextIdx: "vFSt-Q@test3",

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
			args:               args{remoteAddr: "localhost", userID: "SYSOP", params: params, c: nil},
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
			gotResult, gotStatusCode, err := LoadGeneralBoardsByClass(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGeneralBoardsByClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			testutil.TDeepEqual(t, "got", gotResult, tt.expectedResult)

			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("LoadGeneralBoardsByClass() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
