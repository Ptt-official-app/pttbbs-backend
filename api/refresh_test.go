package api

import (
	"net/http"
	"reflect"
	"sync"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"

	"github.com/gin-gonic/gin"
)

func TestRefresh(t *testing.T) {
	setupTest()
	defer teardownTest()

	jwt, _, _ := pttbbsapi.CreateToken("SYSOP", "")
	refreshJwt, _, _ := pttbbsapi.CreateRefreshToken("SYSOP", "")

	req, _ := http.NewRequest("POST", "http://localhost/refresh", nil)
	req.Header = map[string][]string{
		"Authorization": {"bearer " + jwt},
	}
	c0 := &gin.Context{}
	c0.Request = req

	params0 := &RefreshParams{
		ClientID:     "default_client_id",
		ClientSecret: "test_client_secret",
		RefreshToken: refreshJwt,
	}

	expected0 := &RefreshResult{
		Username:  "SYSOP",
		TokenType: "bearer",
	}

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name           string
		args           args
		wantResult     interface{}
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{
			args:           args{params: params0, c: c0},
			wantStatusCode: 200,
			wantResult:     expected0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST), IsOver18: true}
			gotResult, gotStatusCode, err := Refresh(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Refresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			result, _ := gotResult.(*RefreshResult)
			result.AccessToken = ""
			result.RefreshToken = ""
			result.AccessExpireTS = 0
			result.RefreshExpireTS = 0
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Refresh() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("Refresh() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
		wg.Wait()
	}
}
