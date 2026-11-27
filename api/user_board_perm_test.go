package api

import (
	"testing"
	"time"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testPermUserID = bbs.UUserID("testuser")

const testPermBoardID = bbs.BBoardID("10_WhoAmI")

func testPermUser() *schema.UserPermInfo {
	return &schema.UserPermInfo{
		UserID:       testPermUserID,
		Userlevel:    ptttype.PERM_BASIC | ptttype.PERM_POST | ptttype.PERM_LOGINOK,
		Over18:       true,
		Numlogindays: 100,
	}
}

func testPermBoard() *schema.BoardPermInfo {
	return &schema.BoardPermInfo{
		BBoardID: testPermBoardID,
		Brdname:  "WhoAmI",
	}
}

func testFutureNanoTS() types.NanoTS {
	return types.NowNanoTS() + types.NanoTS(time.Hour)
}

func TestCheckUserBoardPermReadableCore(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name      string
		user      func(u *schema.UserPermInfo)
		board     func(b *schema.BoardPermInfo)
		userBoard *schema.UserBoard
		expected  *UserBoardPermReadable
		wantErr   error
	}{
		{name: "normal user", expected: &UserBoardPermReadable{}},
		{
			name:     "sysop reads a hidden board",
			user:     func(u *schema.UserPermInfo) { u.Userlevel |= ptttype.PERM_SYSOP },
			board:    func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_HIDE },
			expected: &UserBoardPermReadable{IsSYSOP: true},
		},
		{
			name:     "police on a BM-level board",
			user:     func(u *schema.UserPermInfo) { u.Userlevel |= ptttype.PERM_POLICE },
			board:    func(b *schema.BoardPermInfo) { b.Level = ptttype.PERM_BM },
			expected: &UserBoardPermReadable{IsPolice: true},
		},
		{
			name:     "police on a normal board is a normal user",
			user:     func(u *schema.UserPermInfo) { u.Userlevel |= ptttype.PERM_POLICE },
			expected: &UserBoardPermReadable{},
		},
		{
			name:     "BM reads a hidden board",
			board:    func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_HIDE; b.BMs = []bbs.UUserID{testPermUserID} },
			expected: &UserBoardPermReadable{IsBM: true},
		},
		{
			name:    "hidden board, not a board friend",
			board:   func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_HIDE },
			wantErr: ErrPermBoardReadHidden,
		},
		{
			name:      "hidden board, board friend",
			board:     func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_HIDE },
			userBoard: &schema.UserBoard{BoardFriend: true},
			expected:  &UserBoardPermReadable{},
		},
		{
			name:      "blocked by the user",
			userBoard: &schema.UserBoard{BoardBlocked: true},
			wantErr:   ErrPermBoardReadBlocked,
		},
		{
			name:      "reported by the user",
			userBoard: &schema.UserBoard{BoardReported: true},
			wantErr:   ErrPermBoardReadReported,
		},
		{
			name:    "over-18 board, user not over 18",
			user:    func(u *schema.UserPermInfo) { u.Over18 = false },
			board:   func(b *schema.BoardPermInfo) { b.IsOver18 = true },
			wantErr: ErrPermBoardReadNotOver18,
		},
		{
			name:    "board level not met",
			board:   func(b *schema.BoardPermInfo) { b.Level = ptttype.PERM_BOARD },
			wantErr: ErrPermBoardReadPermission,
		},
		{
			name:     "board level only limits posting",
			board:    func(b *schema.BoardPermInfo) { b.Level = ptttype.PERM_BOARD; b.BrdAttr = ptttype.BRD_POSTMASK },
			expected: &UserBoardPermReadable{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, board := testPermUser(), testPermBoard()
			if tt.user != nil {
				tt.user(user)
			}
			if tt.board != nil {
				tt.board(board)
			}
			userBoard := tt.userBoard
			if userBoard == nil {
				userBoard = &schema.UserBoard{}
			}

			got, err := checkUserBoardPermReadableCore(testPermUserID, testPermBoardID, user, board, userBoard)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCheckUserBoardPermPostableCore(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name      string
		user      func(u *schema.UserPermInfo)
		board     func(b *schema.BoardPermInfo)
		userBoard *schema.UserBoard
		readable  *UserBoardPermReadable
		wantErr   error
		// errors with the remaining cooldown time in the message
		wantErrPrefix string
	}{
		{name: "normal user"},
		{
			name:     "read-only board, even for sysop",
			board:    func(b *schema.BoardPermInfo) { b.Brdname = ptttype.BN_ALLPOST_s },
			readable: &UserBoardPermReadable{IsSYSOP: true},
			wantErr:  ErrPermPostReadOnly,
		},
		{
			name:      "sysop ignores the bucket",
			userBoard: &schema.UserBoard{BoardBucketExpireNanoTS: testFutureNanoTS()},
			readable:  &UserBoardPermReadable{IsSYSOP: true},
		},
		{
			name:      "in the bucket",
			userBoard: &schema.UserBoard{BoardBucketExpireNanoTS: testFutureNanoTS()},
			wantErr:   ErrBoardBucket,
		},
		{
			name:      "bucket expired",
			userBoard: &schema.UserBoard{BoardBucketExpireNanoTS: types.NowNanoTS() - types.NanoTS(time.Hour)},
		},
		{
			name:  "guest-post board without post permission",
			user:  func(u *schema.UserPermInfo) { u.Userlevel = ptttype.PERM_BASIC },
			board: func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_GUESTPOST },
		},
		{
			name:    "no post permission",
			user:    func(u *schema.UserPermInfo) { u.Userlevel = ptttype.PERM_BASIC },
			wantErr: ErrPermBoardPostPost,
		},
		{
			name:    "restricted post, not a board friend",
			board:   func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_RESTRICTEDPOST },
			wantErr: ErrPermBoardPostRestricted,
		},
		{
			name:      "restricted post, board friend",
			board:     func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_RESTRICTEDPOST },
			userBoard: &schema.UserBoard{BoardFriend: true},
		},
		{
			name:    "violate law on a normal board",
			user:    func(u *schema.UserPermInfo) { u.Userlevel |= ptttype.PERM_VIOLATELAW },
			wantErr: ErrPermBoardPostViolateLaw,
		},
		{
			name:  "violate law on a board for violators",
			user:  func(u *schema.UserPermInfo) { u.Userlevel |= ptttype.PERM_VIOLATELAW },
			board: func(b *schema.BoardPermInfo) { b.Level = ptttype.PERM_VIOLATELAW },
		},
		{
			name:    "board level not met",
			board:   func(b *schema.BoardPermInfo) { b.Level = ptttype.PERM_BOARD },
			wantErr: ErrPermBoardPostPermission,
		},
		{
			name:    "not enough login days",
			user:    func(u *schema.UserPermInfo) { u.Numlogindays = 19 },
			board:   func(b *schema.BoardPermInfo) { b.PostLimitLogins = 2 },
			wantErr: ErrPermBoardPostLoginDays,
		},
		{
			name:  "just enough login days",
			user:  func(u *schema.UserPermInfo) { u.Numlogindays = 20 },
			board: func(b *schema.BoardPermInfo) { b.PostLimitLogins = 2 },
		},
		{
			name:    "too many bad posts",
			user:    func(u *schema.UserPermInfo) { u.BadPost = 6 },
			board:   func(b *schema.BoardPermInfo) { b.PostLimitBadpost = 250 },
			wantErr: ErrPermBoardPostPostLimit,
		},
		{
			name:  "bad posts within the limit",
			user:  func(u *schema.UserPermInfo) { u.BadPost = 5 },
			board: func(b *schema.BoardPermInfo) { b.PostLimitBadpost = 250 },
		},
		{
			name: "cooldown expired",
			user: func(u *schema.UserPermInfo) {
				u.CooldownNanoTS = types.NowNanoTS() - types.NanoTS(time.Hour)
				u.Posttime = types.POSTTIME_REJECT
			},
			board: func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_COOLDOWN },
		},
		{
			name: "sysop level ignores cooldown",
			user: func(u *schema.UserPermInfo) {
				u.Userlevel |= ptttype.PERM_SYSOP
				u.CooldownNanoTS = testFutureNanoTS()
			},
			board: func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_COOLDOWN },
		},
		{
			name:          "cooldown board",
			user:          func(u *schema.UserPermInfo) { u.CooldownNanoTS = testFutureNanoTS() },
			board:         func(b *schema.BoardPermInfo) { b.BrdAttr = ptttype.BRD_COOLDOWN },
			wantErrPrefix: "board cooldown ",
		},
		{
			name: "posttime rejected",
			user: func(u *schema.UserPermInfo) {
				u.CooldownNanoTS = testFutureNanoTS()
				u.Posttime = types.POSTTIME_REJECT
			},
			wantErrPrefix: "board posttime ",
		},
		{
			name: "flood on a popular board",
			user: func(u *schema.UserPermInfo) {
				u.CooldownNanoTS = testFutureNanoTS()
				u.Posttime = 1
			},
			board:         func(b *schema.BoardPermInfo) { b.NUser = 4000 },
			wantErrPrefix: "flood reject ",
		},
		{
			name: "posttime below the limit of a quiet board",
			user: func(u *schema.UserPermInfo) {
				u.CooldownNanoTS = testFutureNanoTS()
				u.Posttime = 9
			},
		},
		{
			name: "flood on a quiet board",
			user: func(u *schema.UserPermInfo) {
				u.CooldownNanoTS = testFutureNanoTS()
				u.Posttime = 10
			},
			wantErrPrefix: "flood reject ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, board := testPermUser(), testPermBoard()
			if tt.user != nil {
				tt.user(user)
			}
			if tt.board != nil {
				tt.board(board)
			}
			userBoard := tt.userBoard
			if userBoard == nil {
				userBoard = &schema.UserBoard{}
			}
			readable := tt.readable
			if readable == nil {
				readable = &UserBoardPermReadable{}
			}

			err := checkUserBoardPermPostableCore(testPermUserID, testPermBoardID, user, board, userBoard, readable)
			switch {
			case tt.wantErrPrefix != "":
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErrPrefix)
			default:
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestCheckUserBoardPermCreatableCore(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name      string
		userlevel ptttype.PERM
		wantErr   error
	}{
		{name: "sysop", userlevel: ptttype.PERM_SYSOP},
		{name: "board admin", userlevel: ptttype.PERM_BOARD},
		{name: "normal user", userlevel: ptttype.PERM_BASIC | ptttype.PERM_POST, wantErr: ErrPermBoardCreatePermission},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUserBoardPermCreatableCore(&schema.UserPermInfo{Userlevel: tt.userlevel})
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestCheckBoardBM(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer schema.Board_c.Drop()

	classID := bbs.BBoardID("2_Class")
	err := schema.UpdateBoardSummaries([]*schema.BoardSummary{
		{BBoardID: classID, Brdname: "Class", BMs: []bbs.UUserID{"classBM"}},
	}, types.NowNanoTS())
	require.NoError(t, err)

	tests := []struct {
		name     string
		userID   bbs.UUserID
		bms      []bbs.UUserID
		parentID bbs.BBoardID
		expected bool
	}{
		{name: "BM of the board", userID: "boardBM", bms: []bbs.UUserID{"boardBM"}, parentID: classID, expected: true},
		{name: "BM of the parent class", userID: "classBM", bms: []bbs.UUserID{"boardBM"}, parentID: classID, expected: true},
		{name: "not a BM", userID: testPermUserID, bms: []bbs.UUserID{"boardBM"}, parentID: classID},
		{name: "parent not in db", userID: testPermUserID, parentID: "3_NotExists"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkBoardBM(tt.userID, testPermBoardID, tt.bms, tt.parentID)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
