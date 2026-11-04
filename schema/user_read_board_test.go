package schema

import (
	"cmp"
	"reflect"
	"slices"
	"testing"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateUserReadBoard(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer UserBoard_c.Drop()

	userReadBoard := &UserReadBoard{
		UserID:           bbs.UUserID("testuser0"),
		BBoardID:         bbs.BBoardID("testboard0"),
		ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	type args struct {
		userReadBoard *UserReadBoard
	}
	tests := []struct {
		name    string
		args    args
		want    *UserReadBoard
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{userReadBoard},
			want: userReadBoard,
		},
		{
			args: args{userReadBoard},
			want: userReadBoard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadBoard(tt.args.userReadBoard); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadBoard() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := &UserReadBoard{}

			query := bson.M{
				USER_BOARD_USER_ID_b:   tt.args.userReadBoard.UserID,
				USER_BOARD_BBOARD_ID_b: tt.args.userReadBoard.BBoardID,
			}

			_ = UserBoard_c.FindOne(query, got, nil)
			testutil.TDeepEqual(t, "got", got, tt.want)
		})
	}
}

func TestUpdateUserReadBoards(t *testing.T) {
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
	userReadBoards1 := []*UserReadBoard{
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard2"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
		{
			UserID:           bbs.UUserID("testuser0"),
			BBoardID:         bbs.BBoardID("testboard3"),
			ReadUpdateNanoTS: types.NanoTS(1234567890000000000),
		},
	}

	type args struct {
		userReadBoards []*UserReadBoard
		updateNanoTS   types.NanoTS
	}
	tests := []struct {
		name         string
		args         args
		findUserID   bbs.UUserID
		findBoardIDs []bbs.BBoardID
		wnat         []*UserReadBoard
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			args:         args{userReadBoards: userReadBoards, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:   bbs.UUserID("testuser0"),
			findBoardIDs: []bbs.BBoardID{"testboard0", "testboard1", "testboard2"},
			wnat:         userReadBoards,
		},
		{
			args:         args{userReadBoards: userReadBoards, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:   bbs.UUserID("testuser0"),
			findBoardIDs: []bbs.BBoardID{"testboard0", "testboard1", "testboard2"},
			wnat:         userReadBoards,
		},
		{
			args:         args{userReadBoards: userReadBoards1, updateNanoTS: types.NanoTS(1234567890000000000)},
			findUserID:   bbs.UUserID("testuser0"),
			findBoardIDs: []bbs.BBoardID{"testboard2", "testboard3"},
			wnat:         userReadBoards1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUserReadBoards(tt.args.userReadBoards, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserReadBoards() error = %v, wantErr %v", err, tt.wantErr)
				got, err := FindUserReadBoards(tt.findUserID, tt.findBoardIDs)
				if err != nil {
					t.Errorf("UpdateUserReadBoards: e: %v", err)
				}

				slices.SortStableFunc(got, func(a, b *UserReadBoard) int {
					return cmp.Compare(a.BBoardID, b.BBoardID)
				})
				testutil.TDeepEqual(t, "got", got, tt.wnat)

				t.Errorf("false")

			}
		})
	}
}

func TestFindUserReadBoards(t *testing.T) {
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

	type args struct {
		userID   bbs.UUserID
		boardIDs []bbs.BBoardID
	}
	tests := []struct {
		name    string
		args    args
		want    []*UserReadBoard
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				userID:   bbs.UUserID("testuser0"),
				boardIDs: []bbs.BBoardID{"testboard0", "testboard1", "testboard2"},
			},
			want: userReadBoards,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindUserReadBoards(tt.args.userID, tt.args.boardIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserReadBoards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserReadBoards() = %v, want %v", got, tt.want)
			}

			slices.SortStableFunc(got, func(a, b *UserReadBoard) int {
				return cmp.Compare(a.BBoardID, b.BBoardID)
			})
			testutil.TDeepEqual(t, "got", got, tt.want)
		})
	}
}
