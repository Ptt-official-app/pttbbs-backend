package api

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/oidcop"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

func TestRefresh(t *testing.T) {
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

	_, refreshToken, _, err := genAccessAndRefreshTokens(context.Background(), "SYSOP", types.WEB_CLIENT_ID, "")
	if err != nil {
		t.Fatalf("unable to genAccessAndRefreshTokens: e: %v", err)
	}
	if refreshToken == "" {
		t.Fatalf("genAccessAndRefreshTokens() did not issue a refresh token")
	}

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST), IsOver18: true}

	// wrong client.
	_, statusCode, err := Refresh("localhost", user, &RefreshParams{ClientID: "other-client", RefreshToken: refreshToken}, c)
	if statusCode != 401 || !errors.Is(err, ErrInvalidToken) {
		t.Errorf("Refresh() wrong client: statusCode: %v e: %v, want 401", statusCode, err)
	}

	// empty refresh token.
	_, statusCode, _ = Refresh("localhost", user, &RefreshParams{}, c)
	if statusCode != 400 {
		t.Errorf("Refresh() empty token: statusCode: %v, want 400", statusCode)
	}

	// refresh.
	gotResult, statusCode, err := Refresh("localhost", user, &RefreshParams{ClientID: types.WEB_CLIENT_ID, RefreshToken: refreshToken}, c)
	if statusCode != 200 || err != nil {
		t.Fatalf("Refresh(): statusCode: %v e: %v", statusCode, err)
	}
	result, ok := gotResult.(*RefreshResult)
	if !ok {
		t.Fatalf("Refresh() result is not *RefreshResult: %v", gotResult)
	}
	if result.Username != "SYSOP" || result.TokenType != "bearer" || result.AccessToken == "" || result.RefreshToken == "" || result.RefreshToken == refreshToken {
		t.Errorf("Refresh() unexpected result: %+v", result)
	}
	if result.AccessExpireTS == 0 || result.RefreshExpireTS == 0 {
		t.Errorf("Refresh() expire not set: %+v", result)
	}

	userID, _, _, _, _, err := verifyAccessToken(result.AccessToken)
	if err != nil || userID != "SYSOP" {
		t.Errorf("verifyAccessToken() new access token: userID: %v e: %v", userID, err)
	}

	// the old refresh token is rotated out.
	_, statusCode, err = Refresh("localhost", user, &RefreshParams{RefreshToken: refreshToken}, c)
	if statusCode != 401 || !errors.Is(err, ErrInvalidToken) {
		t.Errorf("Refresh() reuse old refresh token: statusCode: %v e: %v, want 401", statusCode, err)
	}

	// the new refresh token works.
	_, statusCode, err = Refresh("localhost", user, &RefreshParams{RefreshToken: result.RefreshToken}, c)
	if statusCode != 200 || err != nil {
		t.Errorf("Refresh() with new refresh token: statusCode: %v e: %v", statusCode, err)
	}
}
