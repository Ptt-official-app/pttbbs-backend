package api

import (
	"reflect"
	"sync"
	"testing"
	"time"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestRegisterUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.AccessToken_c.Drop()

	token, err := schema.SetEmailVerification("test@ptt.test", time.Duration(1)*time.Second)
	logrus.Infof("api.TestRegisterUser: after SetEmailVerification: e: %v", err)

	params0 := &RegisterUserParams{
		Token: token,
	}

	expectedDB0 := []*schema.AccessToken{{UserID: "SYSOP"}}

	expected0 := types.INIT_URL

	type args struct {
		remoteAddr string
		params     interface{}
		c          *gin.Context
	}
	tests := []struct {
		name               string
		args               args
		expectedResult     string
		expectedStatusCode int
		expectedDB         []*schema.AccessToken
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			args:               args{remoteAddr: "localhost", params: params0},
			expectedResult:     expected0,
			expectedStatusCode: 303,
			expectedDB:         expectedDB0,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST)}
			gotResult, gotStatusCode, err := RegisterUser(tt.args.remoteAddr, user, tt.args.params, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return

			}

			/*
				query := make(map[string]string)
				query["user_id"] = "testuserid1"

				var ret []*schema.AccessToken
				err = schema.AccessToken_c.Find(query, 0, &ret, nil, nil)
				if err != nil {
					t.Errorf("Login(): unable to find: e: %v", err)
				}
				for _, each := range ret {
					each.UpdateNanoTS = 0
				}
				if len(ret) < 1 {
					t.Errorf("RegisterUser(): unable to find access-token.")
					return
				}
				expected.AccessToken = ret[0].AccessToken
			*/

			if !reflect.DeepEqual(gotResult, tt.expectedResult) {
				t.Errorf("RegisterUser() gotResult = %v, want %v", gotResult, tt.expectedResult)
			}
			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("RegisterUser() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}
