package api

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Ptt-official-app/pttbbs-backend/oidcop"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLogin(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, err := oidcop.NewProvider()
	if err != nil {
		logrus.Errorf("TestLogin: unable to oidcop.NewProvider: e: %v", err)
	}
	defer func() {
		oidcop.PROVIDER = nil
	}()

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

	token_db, err := schema.Get2FA("SYSOP")
	logrus.Infof("TestLogin: after schema.Get2FA: token_db: %v e: %v", token_db, err)

	params0 := &LoginParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Input:        "SYSOP",
		VerifyCode:   token_db.Token,
	}

	expected0 := &LoginResult{TokenType: "bearer", Username: "SYSOP"}
	expectedDB0 := []*schema.AccessToken{{UserID: "SYSOP"}}

	client := schema.NewClient(types.WEB_CLIENT_ID, types.CLIENT_TYPE_APP, nil, "localhost")
	_ = schema.UpdateClient(client)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

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
			args:       args{remoteAddr: "localhost", params: params0, c: ctx},
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
			result, ok := got.(*LoginResult)
			if !ok {
				t.Errorf("Login() not LoginResult: result: %v", result)
				return
			}
			tt.expected.AccessToken = result.AccessToken
			tt.expected.RefreshToken = result.RefreshToken

			tt.expected.AccessExpireTS = result.AccessExpireTS
			tt.expected.RefreshExpireTS = result.RefreshExpireTS

			logrus.Infof("TestLogin: to DeepEqual: result: %v expected: %v", result, tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Login() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLoginWrapper(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, err := oidcop.NewProvider()
	if err != nil {
		logrus.Errorf("TestLoginWrapper: unable to oidcop.NewProvider: e: %v", err)
	}
	defer func() {
		oidcop.PROVIDER = nil
	}()

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

	client := schema.NewClient(types.WEB_CLIENT_ID, types.CLIENT_TYPE_APP, nil, "localhost")

	_ = schema.UpdateClient(client)

	params0 := &LoginParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		Input:        "SYSOP",
		VerifyCode:   token_db.Token,
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
