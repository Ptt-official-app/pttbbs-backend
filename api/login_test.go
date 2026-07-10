package api

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLogin(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.AccessToken_c.Drop()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	_ = schema.CreateUserEmail("SYSOP", "root@localhost.dev", false, 123456890000000000)

	paramsAttemptLogin := &AttemptLoginParams{
		Input: "SYSOP",
	}

	userInfo := &UserInfo{
		UserID:   "SYSOP",
		IsOver18: true,
	}

	AttemptLogin("127.0.0.1", userInfo, paramsAttemptLogin, nil)

	token_db, _ := schema.Get2FA("SYSOP")

	params0 := &LoginParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Input:        "SYSOP",
		VerifyCode:   token_db,
	}

	expected0 := &LoginResult{TokenType: "bearer", Username: "SYSOP"}
	expectedDB0 := []*schema.AccessToken{{UserID: "SYSOP"}}

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name       string
		args       args
		expected   *LoginResult
		expectedDB []*schema.AccessToken
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			args:       args{remoteAddr: "localhost", params: params0},
			expected:   expected0,
			expectedDB: expectedDB0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := Login(tt.args.remoteAddr, userInfo, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			/*
								query := make(map[string]interface{})
								query[schema.ACCESS_TOKEN_USER_ID_b] = "testuserid1"

				s				var ret []*schema.AccessToken
								err = schema.AccessToken_c.Find(query, 0, &ret, nil, nil)
								logrus.Infof("api.TestLogin: after Find: query: %v ret: %v e: %v", query, ret, err)
								if err != nil {
									t.Errorf("Login(): unable to find: e: %v", err)
								}
								for _, each := range ret {
									each.UpdateNanoTS = 0
								}
								if len(ret) < 1 {
									t.Errorf("Login(): unable to find access-token")
									return
								}
								expected.AccessToken = ret[0].AccessToken
			*/
			result := got.(*LoginResult)
			tt.expected.AccessToken = result.AccessToken
			tt.expected.RefreshToken = result.RefreshToken

			tt.expected.AccessExpireTS = result.AccessExpireTS
			tt.expected.RefreshExpireTS = result.RefreshExpireTS

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Login() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLoginWrapper(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	_ = schema.CreateUserEmail("SYSOP", "root@localhost.dev", false, 123456890000000000)

	paramsAttemptLogin := &AttemptLoginParams{
		Input: "SYSOP",
	}

	userInfo := &UserInfo{
		UserID:   "SYSOP",
		IsOver18: true,
	}

	AttemptLogin("127.0.0.1", userInfo, paramsAttemptLogin, nil)

	token_db, _ := schema.Get2FA("SYSOP")

	params0 := &LoginParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Input:        "SYSOP",
		VerifyCode:   token_db,
	}
	type args struct {
		params *LoginParams
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{params: params0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, c, r := testSetRequest(
				LOGIN_R,
				LOGIN_R,
				tt.args.params,
				"",
				"",
				nil,
				"POST",
				LoginWrapper,
			)

			r.ServeHTTP(w, c.Request)

			if w.Code != http.StatusOK {
				t.Errorf("code: %v", w.Code)
			}

			setCookie := w.Header().Get("Set-Cookie")
			logrus.Infof("setCookie: %v", setCookie)
		})
	}
}
