package schema

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func TestGetUserNickname(t *testing.T) {
	setupTest()
	defer teardownTest()

	testUserDetail0.UpdateNanoTS = testUserInfoSummary0.UpdateNanoTS
	UpdateUserDetail(testUserDetail0)

	type args struct {
		userID bbs.UUserID
	}
	tests := []struct {
		name             string
		args             args
		expectedNickname string
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			args:             args{userID: "test_userid"},
			expectedNickname: "test_nickname",
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			gotUserNickname, err := GetUserNameInfo(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserNickname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotNickname := ""
			if gotUserNickname != nil {
				gotNickname = gotUserNickname.Nickname
			}
			if gotNickname != tt.expectedNickname {
				t.Errorf("GetUserNickname() = %v, want %v", gotNickname, tt.expectedNickname)
			}
		})
	}
	wg.Wait()
}
