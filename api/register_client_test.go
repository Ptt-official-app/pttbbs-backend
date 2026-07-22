package api

import (
	"net/http"
	"reflect"
	"sync"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	"github.com/zitadel/oidc/v3/pkg/op"
)

func TestRegisterClient(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.Client_c.Drop()

	params := &RegisterClientParams{
		ClientID: "test_client_id",
	}

	expected := &RegisterClientResult{TokenUser: "SYSOP"}

	expectedDB := []*schema.Client{
		{
			ClientID:           "test_client_id",
			RemoteAddr:         "localhost",
			TheApplicationType: op.ApplicationTypeNative,
			TheAccessTokenType: op.AccessTokenTypeJWT,
			TheResponseTypes:   []oidc.ResponseType{oidc.ResponseTypeCode, oidc.ResponseTypeIDToken, oidc.ResponseTypeIDTokenOnly},
			TheGrantTypes:      []oidc.GrantType{oidc.GrantTypeBearer, oidc.GrantTypeCode, oidc.GrantTypeRefreshToken},
		},
	}

	type args struct {
		remoteAddr string
		params     interface{}
	}
	tests := []struct {
		name       string
		args       args
		expected   interface{}
		expectedDB []*schema.Client
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			args:       args{remoteAddr: "localhost", params: params},
			expected:   expected,
			expectedDB: expectedDB,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &UserInfo{UserID: bbs.UUserID("SYSOP")}
			got, _, err := RegisterClient(tt.args.remoteAddr, user, tt.args.params, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("RegisterClient() = %v, want %v", got, tt.expected)
			}

			query := make(map[string]string)
			query["client_id"] = "test_client_id"

			var ret []*schema.Client
			err = schema.Client_c.Find(query, 0, &ret, nil, nil)
			if err != nil {
				t.Errorf("RegisterClient(): unable to find: e: %v", err)
			}
			for _, each := range ret {
				each.UpdateNanoTS = 0
			}
			if !reflect.DeepEqual(ret, expectedDB) {
				t.Errorf("RegisterClient: find: %v expected: %v", ret, expectedDB)
			}

			testutil.TDeepEqual(t, "ret", ret, expectedDB)
		})
	}
}

func TestRegisterClientWrapper(t *testing.T) {
	setupTest()
	defer teardownTest()

	params0 := &RegisterClientParams{
		ClientID:   "test_client_id_app",
		ClientType: "app",
	}
	params1 := &RegisterClientParams{
		ClientID:   "test_client_id_web",
		ClientType: "web",
	}

	accessTokenSYSOP := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGkiOiJ7XCJjXCI6IFwidGVzdHdlYlwiLCBcInRcIjogMX0iLCJleHAiOjE5NDE5MjYxOTIsInN1YiI6IlNZU09QIn0.cs5tp85Y6ECeysMjW3o5HSX3HkTCVMgSsLQ3EHQuQkI"

	csrfToken, _ := createCSRFToken()

	type args struct {
		params *RegisterClientParams
	}
	tests := []struct {
		name     string
		args     args
		expected types.ClientType
	}{
		// TODO: Add test cases.
		{
			args:     args{params: params0},
			expected: types.CLIENT_TYPE_APP,
		},
		{
			args:     args{params: params1},
			expected: types.CLIENT_TYPE_WEB,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()
			w, c, r := testSetRequest(REGISTER_CLIENT_R, REGISTER_CLIENT_R, tt.args.params, accessTokenSYSOP, csrfToken, nil, "POST", RegisterClientWrapper)
			logrus.Infof("RegisterClientWrapper: remote-addr: %v", c.Request.RemoteAddr)
			r.ServeHTTP(w, c.Request)

			if w.Code != http.StatusOK {
				t.Errorf("code: %v", w.Code)
			}

			client, _ := schema.GetClient(tt.args.params.ClientID)
			if client == nil {
				t.Errorf("RegisterClientWrapper: unable to find client: %v", tt.args.params)
			}
			if !reflect.DeepEqual(client.ClientType, tt.expected) {
				t.Errorf("RegisterClientWrapper: clientType: %v expected: %v", client.ClientType, tt.expected)
			}
		})
	}
	wg.Wait()
}
