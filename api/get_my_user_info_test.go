package api

import (
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMyUserInfo(t *testing.T) {
	setupTest()
	defer teardownTest()

	// the username differs from the user-id (genUsername assigns a random username),
	// so /user/me must look up the user by user-id.
	userec := *testUserSYSOP_b
	userec.Username = "codingman.ab12"
	_, err := DeserializeUserDetailAndUpdateDBForTest(pttbbsapi.GetUserResult(&userec), 123456890000000000)
	require.NoError(t, err)

	user := &UserInfo{UserID: "SYSOP", IsOver18: true}
	gotResult, statusCode, err := GetMyUserInfo(testIP, user, &GetUserInfoParams{}, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, statusCode)

	got := gotResult.(*GetUserInfoResult)
	assert.Equal(t, bbs.UUserID("SYSOP"), got.UserID)
	assert.Equal(t, "codingman.ab12", got.Username)
	// the token owner gets the private fields.
	assert.Equal(t, "CodingMan", got.Realname)
	assert.NotZero(t, got.Userlevel)
}

func TestGetMyUserInfo_guest(t *testing.T) {
	setupTest()
	defer teardownTest()

	user := &UserInfo{UserID: bbs.UUserID(pttbbsapi.GUEST)}
	gotResult, statusCode, err := GetMyUserInfo(testIP, user, &GetUserInfoParams{}, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, statusCode)

	got := gotResult.(*GetUserInfoResult)
	assert.Equal(t, bbs.UUserID(pttbbsapi.GUEST), got.UserID)
}
