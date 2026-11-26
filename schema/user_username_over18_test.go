package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetUserUsernameOver18ByUsername(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer User_c.Drop()

	_, err := User_c.CreateOnly(bson.M{USER_USER_ID_b: bbs.UUserID("over18user")}, &UserUsernameOver18{
		UserID:   "over18user",
		Username: "over18name",
		Over18:   true,
	})
	require.NoError(t, err)
	_, err = User_c.CreateOnly(bson.M{USER_USER_ID_b: bbs.UUserID("minoruser")}, &UserUsernameOver18{
		UserID:   "minoruser",
		Username: "minorname",
		Over18:   false,
	})
	require.NoError(t, err)

	tests := []struct {
		name     string
		username string
		want     *UserUsernameOver18
	}{
		{name: "over18", username: "over18name", want: &UserUsernameOver18{UserID: "over18user", Username: "over18name", Over18: true}},
		{name: "not over18", username: "minorname", want: &UserUsernameOver18{UserID: "minoruser", Username: "minorname", Over18: false}},
		{name: "not found", username: "nobody", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserUsernameOver18ByUsername(tt.username)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
