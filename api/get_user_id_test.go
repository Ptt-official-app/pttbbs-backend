package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

func TestGetUserID(t *testing.T) {
	expected0 := &GetUserIDResult{UserID: "SYSOP", TokenUser: "SYSOP"}
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
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: nil},
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
			gotResult, gotStatusCode, err := GetUserID(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("GetUserID() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("GetUserID() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
