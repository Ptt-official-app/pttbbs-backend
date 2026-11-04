package schema

import (
	"cmp"
	"reflect"
	"slices"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func TestFindUserBoards(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserBoard_c.Drop()

	userReadBoards := []*UserReadBoard{
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard0"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard1"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard2"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	_ = UpdateUserReadBoards(userReadBoards, types.NanoTS(1234567890000000000))

	userBoards := []*UserBoard{
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard0"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard1"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard2"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	type args struct {
		userID   bbs.UUserID
		boardIDs []bbs.BBoardID
	}
	tests := []struct {
		name    string
		args    args
		want    []*UserBoard
		wantErr bool
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			args: args{
				userID:   bbs.UUserID("testuser0"),
				boardIDs: []bbs.BBoardID{"testboard0", "testboard1", "testboard2"},
			},
			want: userBoards,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUserBoards(tt.args.userID, tt.args.boardIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserBoards() = %v, want %v", got, tt.want)
			}

			slices.SortStableFunc(got, func(a, b *UserBoard) int {
				return cmp.Compare(a.BBoardID, b.BBoardID)
			})
			testutil.TDeepEqual(t, "got", got, tt.want)
		})
	}
}
