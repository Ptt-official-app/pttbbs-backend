package api

import (
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/gin-gonic/gin"
)

func TestChangePasswd(t *testing.T) {
	setupTest()
	defer teardownTest()

	params := &ChangePasswdParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",

		OrigPassword:    "123123",
		Password:        "123124",
		PasswordConfirm: "123124",
	}

	path := &ChangePasswdPath{
		UserID: "SYSOP",
	}

	expected := &ChangePasswdResult{
		Username:  "SYSOP",
		TokenType: "bearer",
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
		expectedResult     *ChangePasswdResult
		expectedStatusCode int
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: params, path: path, c: nil},
			expectedResult:     expected,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := ChangePasswd(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangePasswd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			result := gotResult.(*ChangePasswdResult)
			tt.expectedResult.AccessToken = result.AccessToken
			testutil.TDeepEqual(t, "result", result, tt.expectedResult)
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("ChangePasswd() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
