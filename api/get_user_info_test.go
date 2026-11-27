package api

import (
	"sync"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)

	path0 := &GetUserInfoPath{Username: "SYSOP"}

	type args struct {
		remoteAddr string
		userID     bbs.UUserID
		params     any
		path       any
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
			args:               args{remoteAddr: testIP, userID: "SYSOP", params: nil, path: path0, c: nil},
			expectedResult:     testUserInfoResult0,
			expectedStatusCode: 200,
		},
	}
	var wg sync.WaitGroup
	for _, tt := range tests {
		wg.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			defer wg.Done()

			user := &UserInfo{UserID: tt.args.userID, IsOver18: true}
			gotResult, gotStatusCode, err := GetUserInfo(tt.args.remoteAddr, user, tt.args.params, tt.args.path, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			result := gotResult.(*GetUserInfoResult)
			result.UpdateTS = 0
			result.EmailTS = 0
			testutil.TDeepEqual(t, "result", result, tt.expectedResult)

			if gotStatusCode != tt.expectedStatusCode {
				t.Errorf("GetUserInfo() gotStatusCode = %v, want %v", gotStatusCode, tt.expectedStatusCode)
			}
		})
	}
	wg.Wait()
}

func TestGetUserInfo_otherUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)
	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserChhsiao123_b, 123456891000000000)
	_ = schema.SetUserIsGovernmentIDOver18("SYSOP", 123456892000000000)

	// chhsiao123 is a normal user querying SYSOP.
	user := &UserInfo{UserID: "chhsiao123", IsOver18: true}
	gotResult, statusCode, err := GetUserInfo(testIP, user, nil, &GetUserInfoPath{Username: "SYSOP"}, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, statusCode)

	got := gotResult.(*GetUserInfoResult)
	assert.Equal(t, bbs.UUserID("SYSOP"), got.UserID)
	assert.Equal(t, "神", got.Nickname)

	assert.Empty(t, got.Realname)
	assert.False(t, got.IsGovernmentID)
	assert.Empty(t, got.LastIP)
	assert.False(t, got.Over18)
	assert.False(t, got.Invisible)
	assert.Empty(t, got.Career)
	assert.Zero(t, got.Role)
	assert.Zero(t, got.Userlevel)

	// the owner still sees everything.
	self := &UserInfo{UserID: "SYSOP", IsOver18: true}
	gotResult, _, err = GetUserInfo(testIP, self, nil, &GetUserInfoPath{Username: "SYSOP"}, nil)
	require.NoError(t, err)
	gotSelf := gotResult.(*GetUserInfoResult)
	assert.Equal(t, "CodingMan", gotSelf.Realname)
	assert.True(t, gotSelf.Over18)
	assert.True(t, gotSelf.IsGovernmentID)
	assert.Equal(t, "全景軟體", gotSelf.Career)
	assert.NotZero(t, gotSelf.Userlevel)
}

func TestGetUserInfo_notFound(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	_, statusCode, err := GetUserInfo(testIP, user, nil, &GetUserInfoPath{Username: "nosuchuser"}, nil)
	assert.Error(t, err)
	assert.Equal(t, 404, statusCode)
}

func TestGetUserInfo_guestQueriesUsername(t *testing.T) {
	setupTest()
	defer teardownTest()

	_, _ = DeserializeUserDetailAndUpdateDBForTest(testUserSYSOP_b, 123456890000000000)

	// a guest gets the queried user's public info, not the guest's own info.
	guest := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST)}
	gotResult, statusCode, err := GetUserInfo(testIP, guest, nil, &GetUserInfoPath{Username: "SYSOP"}, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, statusCode)

	got := gotResult.(*GetUserInfoResult)
	assert.Equal(t, bbs.UUserID("SYSOP"), got.UserID)
	assert.Equal(t, "神", got.Nickname)
	assert.Empty(t, got.Realname)
	assert.Zero(t, got.Userlevel)
}
