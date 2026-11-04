package api

import (
	"reflect"
	"sync"
	"testing"

	pttbbsai "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
)

func TestChangeEmail(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.UserEmail_c.Drop()

	jwt0, _ := pttbbsai.CreateEmailToken("SYSOP", "{\"c\":\"default_client_id\",\"t\":\"app\"}", "test@ptt.test", pttbbsai.CONTEXT_CHANGE_EMAIL)

	params0 := &ChangeEmailParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Jwt:          jwt0,
	}

	path0 := &ChangeEmailPath{
		UserID: "SYSOP",
	}

	result0 := &ChangeEmailResult{
		Email: "test@ptt.test",

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
			gotResult, gotStatusCode, err := ChangeEmail(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("ChangeEmail() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("ChangeEmail() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
