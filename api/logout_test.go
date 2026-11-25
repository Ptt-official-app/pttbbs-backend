package api

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/oidcop"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestLogout(t *testing.T) {
	setupTest()
	defer teardownTest()

	theContext, _ := gin.CreateTestContext(httptest.NewRecorder())

	jsonBytes := []byte{}

	req, _ := http.NewRequest("POST", LOGOUT_R, bytes.NewBuffer(jsonBytes))

	theContext.Request = req

	logrus.Infof("theContext: %v", theContext)

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
			args: args{
				remoteAddr: "localhost",
				c:          theContext,
			},

			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST), IsOver18: true}
			gotResult, gotStatusCode, err := Logout(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("Logout() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("Logout() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
		wg.Wait()
	}
}

func TestLogoutRevokesToken(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, err := oidcop.NewProvider()
	if err != nil {
		t.Fatalf("unable to oidcop.NewProvider: e: %v", err)
	}
	defer func() {
		oidcop.PROVIDER = nil
	}()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_ = schema.CreateUserEmail("SYSOP", "root@localhost.dev", false, 123456890000000000)

	client := schema.NewClient(types.WEB_CLIENT_ID, types.CLIENT_TYPE_APP, nil, "localhost")
	_ = schema.UpdateClient(client)

	accessToken, refreshToken := testGenTokensWithRefreshToken(t)

	userID, _, _, _, _, err := verifyAccessToken(accessToken)
	if err != nil || userID != "SYSOP" {
		t.Fatalf("verifyAccessToken() before logout: userID: %v e: %v", userID, err)
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req, _ := http.NewRequest("POST", LOGOUT_R, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", "Bearer "+accessToken)
	c.Request = req

	user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST)}
	_, statusCode, err := Logout("localhost", user, nil, c)
	if err != nil || statusCode != 200 {
		t.Fatalf("Logout(): statusCode: %v e: %v", statusCode, err)
	}

	if userID, _, _, _, _, err = verifyAccessToken(accessToken); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("verifyAccessToken() after logout: userID: %v e: %v, want ErrInvalidToken", userID, err)
	}
	if _, err = schema.GetOIDCRefreshToken(refreshToken); err == nil {
		t.Errorf("GetOIDCRefreshToken() after logout succeeded unexpectedly")
	}
}

func TestLogoutRevokesRefreshTokenFromCookie(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, err := oidcop.NewProvider()
	if err != nil {
		t.Fatalf("unable to oidcop.NewProvider: e: %v", err)
	}
	defer func() {
		oidcop.PROVIDER = nil
	}()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	client := schema.NewClient(types.WEB_CLIENT_ID, types.CLIENT_TYPE_APP, nil, "localhost")
	_ = schema.UpdateClient(client)

	accessToken, refreshToken := testGenTokensWithRefreshToken(t)

	// only the refresh token is sent (e.g. the access token cookie already expired).
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	req, _ := http.NewRequest("POST", LOGOUT_R, bytes.NewBuffer(nil))
	req.AddCookie(&http.Cookie{Name: types.REFRESH_TOKEN_NAME, Value: refreshToken})
	c.Request = req

	user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST)}
	if _, statusCode, err := Logout("localhost", user, nil, c); err != nil || statusCode != 200 {
		t.Fatalf("Logout(): statusCode: %v e: %v", statusCode, err)
	}

	if _, err = schema.GetOIDCRefreshToken(refreshToken); err == nil {
		t.Errorf("GetOIDCRefreshToken() after logout succeeded unexpectedly")
	}
	if _, _, _, _, _, err = verifyAccessToken(accessToken); !errors.Is(err, ErrInvalidToken) {
		t.Errorf("verifyAccessToken() after logout: e: %v, want ErrInvalidToken", err)
	}
}

// testGenTokensWithRefreshToken creates an access token, and a refresh token bound to it in the token store.
func testGenTokensWithRefreshToken(t *testing.T) (accessToken string, refreshToken string) {
	accessToken, _, _, err := genAccessAndRefreshTokens(context.Background(), "SYSOP", types.WEB_CLIENT_ID, "")
	if err != nil {
		t.Fatalf("unable to genAccessAndRefreshTokens: e: %v", err)
	}
	claim, err := parseAccessToken(accessToken)
	if err != nil {
		t.Fatalf("unable to parseAccessToken: e: %v", err)
	}
	accessToken_db, err := schema.GetOIDCAccessToken(claim.JWTID)
	if err != nil {
		t.Fatalf("unable to GetOIDCAccessToken: e: %v", err)
	}

	refreshToken = "test-refresh-token"
	accessToken_db.RefreshTokenID = refreshToken
	_ = schema.DeleteOIDCAccessToken(claim.JWTID)
	if err = schema.SetOIDCAccessToken(accessToken_db); err != nil {
		t.Fatalf("unable to SetOIDCAccessToken: e: %v", err)
	}
	err = schema.SetOIDCRefreshToken(&schema.OIDCRefreshToken{
		ID:            refreshToken,
		Token:         refreshToken,
		ApplicationID: types.WEB_CLIENT_ID,
		Subject:       "SYSOP",
		AccessToken:   claim.JWTID,
	})
	if err != nil {
		t.Fatalf("unable to SetOIDCRefreshToken: e: %v", err)
	}

	return accessToken, refreshToken
}
