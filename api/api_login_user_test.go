package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/oidcop"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func Test_getLoginUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, err := oidcop.NewProvider()
	if err != nil {
		t.Fatalf("unable to oidcop.NewProvider: e: %v", err)
	}
	defer func() {
		oidcop.PROVIDER = nil
	}()

	origServiceMode := types.SERVICE_MODE
	types.SERVICE_MODE = types.PRODUCTION // check csrf-token.
	defer func() {
		types.SERVICE_MODE = origServiceMode
	}()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	client := schema.NewClient(types.WEB_CLIENT_ID, types.CLIENT_TYPE_APP, nil, "localhost")
	_ = schema.UpdateClient(client)

	accessToken, _, _, err := genAccessAndRefreshTokens(context.Background(), "SYSOP", types.WEB_CLIENT_ID, "")
	if err != nil {
		t.Fatalf("unable to genAccessAndRefreshTokens: e: %v", err)
	}
	csrfToken, _ := createCSRFToken()

	tests := []struct {
		name           string
		accessToken    string
		csrfHeader     string
		csrfCookie     string
		wantUserID     bbs.UUserID
		wantStatusCode int
		wantErr        error
	}{
		{
			name:           "no token: guest",
			wantUserID:     bbs.UUserID(pttbbsapi.GUEST),
			wantStatusCode: 200,
		},
		{
			name:           "invalid token: 401",
			accessToken:    "invalid-token",
			csrfHeader:     csrfToken,
			csrfCookie:     csrfToken,
			wantStatusCode: 401,
			wantErr:        ErrInvalidToken,
		},
		{
			name:           "valid token without csrf-token: guest",
			accessToken:    accessToken,
			wantUserID:     bbs.UUserID(pttbbsapi.GUEST),
			wantStatusCode: 200,
		},
		{
			name:           "valid token with csrf-token: user",
			accessToken:    accessToken,
			csrfHeader:     csrfToken,
			csrfCookie:     csrfToken,
			wantUserID:     "SYSOP",
			wantStatusCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			req, _ := http.NewRequest("GET", "/", nil)
			if tt.accessToken != "" {
				req.AddCookie(&http.Cookie{Name: types.ACCESS_TOKEN_NAME, Value: tt.accessToken})
			}
			if tt.csrfCookie != "" {
				req.AddCookie(&http.Cookie{Name: types.CSRF_TOKEN, Value: tt.csrfCookie})
			}
			if tt.csrfHeader != "" {
				req.Header.Set("X-CSRFToken", tt.csrfHeader)
			}
			c.Request = req

			gotUserID, _, gotStatusCode, err := getLoginUser(c)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("getLoginUser() e: %v, want %v", err, tt.wantErr)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("getLoginUser() statusCode: %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("getLoginUser() userID: %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}
