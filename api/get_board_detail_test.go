package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/apitypes"
	"github.com/gin-gonic/gin"
)

func TestGetBoardDetail(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	LoadAutoCompleteBoards("", user, NewLoadAutoCompleteBoardsParams(), nil)

	params0 := &GetBoardDetailParams{}
	path0 := &GetBoardDetailPath{FBoardID: "test1"}

	result0 := GetBoardDetailResult(&apitypes.BoardDetail{
		FBoardID:  "test1",
		Brdname:   "test1",
		Title:     "測試1",
		BrdAttr:   0,
		BoardType: "◎",
		Category:  "測試",
		NUser:     100,
		BMs:       []bbs.UUserID{"okcool", "teemo"},
		Total:     123,

		LastPostTime: 1234567890,

		TokenUser: "SYSOP",
	})

	params1 := &GetBoardDetailParams{Fields: "bms,title"}
	path1 := &GetBoardDetailPath{FBoardID: "test1"}

	result1 := GetBoardDetailResult(&apitypes.BoardDetail{
		FBoardID: "test1",
		Title:    "測試1",
		BMs:      []bbs.UUserID{"okcool", "teemo"},

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
		wantResult     interface{}
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{remoteAddr: testIP, userID: "SYSOP", params: params0, path: path0},
			wantResult:     result0,
			wantStatusCode: 200,
		},
		{
			args:           args{remoteAddr: testIP, userID: "SYSOP", params: params1, path: path1},
			wantResult:     result1,
			wantStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := GetBoardDetail(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBoardDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetBoardDetail() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("GetBoardDetail() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
		wg.Wait()
	}
}
