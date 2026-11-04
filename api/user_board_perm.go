package api

import (
	"slices"

	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/Ptt-official-app/pttbbs-backend/schema"
	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/gin-gonic/gin"
)

type UserBoardPermReadable struct {
	IsSYSOP bool

	IsPolice bool

	IsBM bool

	IsBoardFriend bool
}

type CooldownLimit struct {
	NUser    int
	Posttime int
}

var COOLDOWN_LIMIT = []*CooldownLimit{
	{NUser: 4000, Posttime: 1},
	{NUser: 2000, Posttime: 2},
	{NUser: 1000, Posttime: 3},
	{NUser: 0, Posttime: 10},
}

// CheckUserBoardPermCreatable
func CheckUserBoardPermCreatable(user *UserInfo, c *gin.Context) (err error) {
	userPermInfo, err := getUserPermInfo(user.UserID, c)
	if err != nil {
		return err
	}
	if userPermInfo == nil {
		return ErrInvalidUser
	}

	userPermInfo.Over18 = userPermInfo.Over18 && user.IsOver18

	return checkUserBoardPermCreatableCore(userPermInfo)
}

func checkUserBoardPermCreatableCore(userPermInfo *schema.UserPermInfo) (err error) {
	if userPermInfo.Userlevel.HasUserPerm(ptttype.PERM_SYSOP) {
		return nil
	}

	if userPermInfo.Userlevel.HasUserPerm(ptttype.PERM_BOARD) {
		return nil
	}

	return ErrPermBoardCreatePermission
}

// CheckUserBoardPermReadable
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L185
func CheckUserBoardPermReadable(user *UserInfo, boardID bbs.BBoardID, c *gin.Context) (userBoardPerm *UserBoardPermReadable, err error) {
	userID := user.UserID
	userPermInfo, err := getUserPermInfo(userID, c)
	if err != nil {
		return nil, err
	}
	if userPermInfo == nil {
		return nil, ErrNoUser
	}

	userPermInfo.Over18 = userPermInfo.Over18 && user.IsOver18

	boardPermInfo, err := schema.GetBoardPermInfo(boardID)
	if err != nil {
		return nil, err
	}
	if boardPermInfo == nil {
		return nil, ErrNoBoard
	}

	userBoardInfo, err := schema.FindUserBoard(userID, boardID)
	if err != nil {
		return nil, err
	}

	return checkUserBoardPermReadableCore(userID, boardID, userPermInfo, boardPermInfo, userBoardInfo)
}

// checkUserBoardPermReadableCore
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L185
func checkUserBoardPermReadableCore(userID bbs.UUserID, boardID bbs.BBoardID, userPermInfo *schema.UserPermInfo, boardPermInfo *schema.BoardPermInfo, userBoardInfo *schema.UserBoard) (userBoardPerm *UserBoardPermReadable, err error) {
	// 1. if the user is SYSOP: can read, edit, post.
	if userPermInfo.Userlevel.HasUserPerm(ptttype.PERM_SYSOP) {
		return &UserBoardPermReadable{
			IsSYSOP: true,
		}, nil
	}

	// 2. if is POLICE or POLICE_MAN and  board is BM-board: can read, edit, post.
	if boardPermInfo.Level.HasUserPerm(ptttype.PERM_BM) && userPermInfo.Userlevel.HasUserPerm(ptttype.PERM_POLICE|ptttype.PERM_POLICE_MAN) {
		return &UserBoardPermReadable{
			IsPolice: true,
		}, nil
	}

	// 3. if is BM: can read and edit
	isBM, err := checkBoardBM(userID, boardID, boardPermInfo.BMs, boardPermInfo.ParentID)
	if err != nil {
		return nil, err
	}
	if isBM {
		return &UserBoardPermReadable{
			IsBM: true,
		}, nil
	}

	// 4. if is hidden board
	if boardPermInfo.BrdAttr.HasPerm(ptttype.BRD_HIDE) {
		err = checkBoardFriend(userBoardInfo)
		if err != nil {
			return nil, ErrPermBoardReadHidden
		}
	}

	// 5. if user blocks board
	err = checkBoardBlocked(userBoardInfo)
	if err != nil {
		return nil, ErrPermBoardReadBlocked
	}

	// 6. if user reports board
	err = checkBoardReported(userBoardInfo)
	if err != nil {
		return nil, ErrPermBoardReadReported
	}

	// 5. require age over 18
	if boardPermInfo.IsOver18 && !userPermInfo.Over18 {
		return nil, ErrPermBoardReadNotOver18
	}

	// 6. board perm meets user perm
	if boardPermInfo.Level != 0 && !boardPermInfo.BrdAttr.HasPerm(ptttype.BRD_POSTMASK) && !userPermInfo.Userlevel.HasUserPerm(boardPermInfo.Level) {
		return nil, ErrPermBoardReadPermission
	}

	// 7. pass
	return &UserBoardPermReadable{}, nil
}

// CheckUserBoardPermPostable
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/cache.c#L209
func CheckUserBoardPermPostable(user *UserInfo, boardID bbs.BBoardID, c *gin.Context) (err error) {
	userID := user.UserID
	userPermInfo, err := getUserPermInfo(userID, c)
	if err != nil {
		return err
	}
	if userPermInfo == nil {
		return ErrInvalidUser
	}

	userPermInfo.Over18 = userPermInfo.Over18 && user.IsOver18

	boardPermInfo, err := schema.GetBoardPermInfo(boardID)
	if err != nil {
		return err
	}
	if boardPermInfo == nil {
		return ErrNoBoard
	}

	userBoardInfo, err := schema.FindUserBoard(userID, boardID)
	if err != nil {
		return err
	}

	userBoardPermReadable, err := checkUserBoardPermReadableCore(userID, boardID, userPermInfo, boardPermInfo, userBoardInfo)
	if err != nil {
		return err
	}

	return checkUserBoardPermPostableCore(userID, boardID, userPermInfo, boardPermInfo, userBoardInfo, userBoardPermReadable)
}

// checkUserBoardPermPostableCore
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/cache.c#L209
func checkUserBoardPermPostableCore(userID bbs.UUserID, boardID bbs.BBoardID, userPermInfo *schema.UserPermInfo, boardPermInfo *schema.BoardPermInfo, userBoardInfo *schema.UserBoard, userBoardPermReadable *UserBoardPermReadable) (err error) {
	err = checkReadOnlyBoard(boardPermInfo)
	if err != nil {
		return err
	}

	// 1. if the user is SYSOP: can post
	if userBoardPermReadable.IsSYSOP {
		return nil
	}

	// 2. if is banned by board
	err = checkBannedByBoard(userBoardInfo)
	if err != nil {
		return err
	}

	// 3. if is default-board (SYSOP): can post
	if boardID == bbs.BBoardID(ptttype.DEFAULT_BOARD) {
		return nil
	}

	// 4. if guest posttable
	if boardPermInfo.BrdAttr.HasPerm(ptttype.BRD_GUESTPOST) {
		return nil
	}

	// 5. XXX BM posttable?
	/*
		if userBoardPermReadable.IsBM {
			return nil
		}
	*/

	// 6. no post permission
	if !userPermInfo.Userlevel.HasUserPerm(ptttype.PERM_POST) {
		return ErrPermBoardPostPost
	}

	// 7. allow only board-friend
	//
	// 7.1. hidden board already checked in read.
	if !boardPermInfo.BrdAttr.HasPerm(ptttype.BRD_HIDE) && boardPermInfo.BrdAttr.HasPerm(ptttype.BRD_RESTRICTEDPOST) {
		err = checkBoardFriend(userBoardInfo)
		if err != nil {
			return ErrPermBoardPostRestricted
		}
	}

	// 8. violate law
	if userPermInfo.Userlevel.HasUserPerm(ptttype.PERM_VIOLATELAW) {
		if boardPermInfo.Level.HasUserPerm(ptttype.PERM_VIOLATELAW) {
			return nil
		}
		return ErrPermBoardPostViolateLaw
	}

	// 9. board perm meets user perm
	if boardPermInfo.Level != 0 && !boardPermInfo.BrdAttr.HasPerm(ptttype.BRD_POSTMASK) && !userPermInfo.Userlevel.HasUserPerm(boardPermInfo.Level) {
		return ErrPermBoardPostPermission
	}

	// 10. board post restrictions
	err = checkBoardRestriction(userPermInfo, boardPermInfo)
	if err != nil {
		return err
	}

	// 11. check cooldown
	err = checkCooldown(userID, userPermInfo, boardPermInfo)
	if err != nil {
		return err
	}

	// 12. pass
	return nil
}

// CheckUserBoardPermEditable
func CheckUserBoardPermEditable(user *UserInfo, boardID bbs.BBoardID, c *gin.Context) (err error) {
	userID := user.UserID
	userPermInfo, err := getUserPermInfo(userID, c)
	if err != nil {
		return err
	}
	if userPermInfo == nil {
		return ErrInvalidUser
	}

	userPermInfo.Over18 = userPermInfo.Over18 && user.IsOver18

	boardPermInfo, err := schema.GetBoardPermInfo(boardID)
	if err != nil {
		return err
	}
	if boardPermInfo == nil {
		return ErrNoBoard
	}

	userBoardInfo, err := schema.FindUserBoard(userID, boardID)
	if err != nil {
		return err
	}

	userBoardPermReadable, err := checkUserBoardPermReadableCore(userID, boardID, userPermInfo, boardPermInfo, userBoardInfo)
	if err != nil {
		return err
	}

	err = checkUserBoardPermPostableCore(userID, boardID, userPermInfo, boardPermInfo, userBoardInfo, userBoardPermReadable)
	if err != nil {
		return err
	}

	if userBoardPermReadable.IsBM || userBoardPermReadable.IsSYSOP {
		return nil
	}

	return ErrPermBoardEditPermission
}

func checkBoardFriend(userBoardInfo *schema.UserBoard) (err error) {
	if userBoardInfo == nil {
		return ErrNotFriend
	}

	if !userBoardInfo.BoardFriend {
		return ErrNotFriend
	}

	return nil
}

func checkBoardBlocked(userBoardInfo *schema.UserBoard) (err error) {
	if userBoardInfo == nil {
		return nil
	}

	if userBoardInfo.BoardBlocked {
		return ErrBoardBlocked
	}

	return nil
}

func checkBoardReported(userBoardInfo *schema.UserBoard) (err error) {
	if userBoardInfo == nil {
		return nil
	}

	if userBoardInfo.BoardReported {
		return ErrBoardReported
	}

	return nil
}

// checkReadOnlyBoard
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/board.c#L59
func checkReadOnlyBoard(boardPermInfo *schema.BoardPermInfo) (err error) {
	if boardPermInfo.Brdname == ptttype.BN_ALLPOST_s ||
		boardPermInfo.Brdname == ptttype.BN_SECURITY_s ||
		boardPermInfo.Brdname == ptttype.BN_ALLHIDPOST_s {
		return ErrPermPostReadOnly
	}

	return nil
}

// checkBannedByBoard
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/acl.c#L75
func checkBannedByBoard(userBoardInfo *schema.UserBoard) (err error) {
	nowNanoTS := types.NowNanoTS()
	if userBoardInfo.BoardBucketExpireNanoTS > nowNanoTS {
		return ErrBoardBucket
	}

	return nil
}

func checkBoardBM(userID bbs.UUserID, boardID bbs.BBoardID, bms []bbs.UUserID, boardParentID bbs.BBoardID) (ret bool, err error) {
	if slices.Contains(bms, userID) {
		return true, nil
	}

	for boardParentID != "" {
		boardPermInfo, err := schema.GetBoardPermInfo(boardParentID)
		if err != nil {
			return false, err
		}

		if slices.Contains(boardPermInfo.BMs, userID) {
			return true, nil
		}

		boardParentID = boardPermInfo.ParentID
	}

	return false, nil
}

// checkBoardRestriction
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/cal.c#L6
func checkBoardRestriction(userPermInfo *schema.UserPermInfo, boardPermInfo *schema.BoardPermInfo) (err error) {
	if userPermInfo.Numlogindays/10 < boardPermInfo.PostLimitLogins {
		return ErrPermBoardPostLoginDays
	}

	if userPermInfo.BadPost > (255 - boardPermInfo.PostLimitBadpost) {
		return ErrPermBoardPostPostLimit
	}
	return nil
}

// checkCooldown
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/bbs.c#L4244
func checkCooldown(userID bbs.UUserID, userPermInfo *schema.UserPermInfo, boardPermInfo *schema.BoardPermInfo) (err error) {
	nowNanoTS := types.NowNanoTS()
	diffNanoTS := userPermInfo.CooldownNanoTS - nowNanoTS
	if diffNanoTS < 0 { // cooldown expired
		if userPermInfo.Posttime > 0 {
			_ = schema.UpdateUserPosttime(userID, 0)
		}
		return nil
	}

	if userPermInfo.Userlevel.HasUserPerm(ptttype.PERM_SYSOP) {
		return nil
	}

	if boardPermInfo.BrdAttr.HasPerm(ptttype.BRD_COOLDOWN) {
		return ErrBoardCooldown(diffNanoTS)
	}

	if userPermInfo.Posttime == types.POSTTIME_REJECT {
		return ErrBoardPosttime(diffNanoTS)
	}

	for _, each := range COOLDOWN_LIMIT {
		if boardPermInfo.NUser >= each.NUser && userPermInfo.Posttime >= each.Posttime {
			return ErrFloodReject(diffNanoTS)
		}
	}

	return nil
}
