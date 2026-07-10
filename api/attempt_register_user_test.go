package api

import (
	"reflect"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAttemptRegisterUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := &AttemptRegisterUserParams{
		Email: "test2@ptt.test",
	}

	expect0 := &AttemptRegisterUserResult{}

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     interface{}
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
			user := &UserInfo{UserID: "temp"}
			gotResult, gotStatusCode, err := AttemptRegisterUser(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AttemptRegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("AttemptRegisterUser() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("AttemptRegisterUser() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}
