package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

func TestAttemptSetIDEmail(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserIDEmail_c.Drop()

	params0 := &AttemptSetIDEmailParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",

		Password: "123123",
		Email:    "test@ptt.test",
	}

	path0 := &AttemptSetIDEmailPath{
		UserID: "SYSOP",
	}

	result0 := &AttemptSetIDEmailResult{
		UserID: "SYSOP",
		Email:  "test@ptt.test",

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
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := AttemptSetIDEmail(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AttemptSetIDEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("AttemptSetIDEmail() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("AttemptSetIDEmail() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
