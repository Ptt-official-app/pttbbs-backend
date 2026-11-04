package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckExistsUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := &CheckExistsUserParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",

		Username: "SYSOP2",
	}

	expect0 := &CheckExistsUserResult{IsExists: false}

	type args struct {
		remoteAddr string
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
			args:               args{remoteAddr: testIP, params: params0},
			expectedResult:     expect0,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			user := &UserInfo{UserID: "guest", IsOver18: true}
			gotResult, gotStatusCode, err := CheckExistsUser(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckExistsUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("CheckExistsUser() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("CheckExistsUser() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
