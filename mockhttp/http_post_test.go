package mockhttp

import (
	"reflect"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func Test_parseResult(t *testing.T) {
	setupTest()
	defer teardownTest()

	result1_b := &api.LoginResult{Jwt: "test_token", TokenType: "bearer"}
	var result1 *api.LoginResult

	result2_b := &api.RegisterResult{Jwt: "test_token", TokenType: "bearer"}
	var result2 *api.RegisterResult

	type args struct {
		backendResult any
		httpResult    any
	}
	tests := []struct {
		name     string
		args     args
		expected any
	}{
		// TODO: Add test cases.
		{
			args:     args{backendResult: result1_b, httpResult: &result1},
			expected: &result1,
		},
		{
			args:     args{backendResult: result2_b, httpResult: &result2},
			expected: &result2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseResult(tt.args.backendResult, tt.args.httpResult)

			httpValue := reflect.ValueOf(tt.args.httpResult).Elem().Interface()
			expectedValue := reflect.ValueOf(tt.expected).Elem().Interface()

			testutil.TDeepEqual(t, "result", httpValue, expectedValue)
		})
	}
}
