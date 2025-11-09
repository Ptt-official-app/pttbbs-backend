package api

import (
	"errors"
	"fmt"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

var (
	ErrInvalidRemoteAddr = errors.New("invalid remote addr")

	ErrInvalidParams            = errors.New("invalid params")
	ErrLoginFailed              = errors.New("login failed")
	ErrNotImplemented           = errors.New("not implemented")
	ErrInvalidPath              = errors.New("invalid path")
	ErrInvalidToken             = errors.New("invalid token")
	ErrInvalidOrigin            = errors.New("invalid origin")
	ErrInvalidBackendStatusCode = errors.New("invalid backend status code")
	ErrNoUser                   = errors.New("no user")
	ErrNoBoard                  = errors.New("no board")
	ErrNoArticle                = errors.New("no article")
	ErrNoUserBoard              = errors.New("no user board")
	ErrInvalidBoardname         = errors.New("invalid boardname")

	ErrAlreadyDeleted = errors.New("already deleted")
	ErrFileNotFound   = errors.New("file not found")

	ErrInvalidUser   = errors.New("invalid user")
	ErrInvalidClient = errors.New("invalid client")

	ErrAlreadyExists = errors.New("already exists")

	ErrInvalidFav = errors.New("invalid fav")

	ErrNotFriend     = errors.New("not friend")
	ErrBoardBlocked  = errors.New("board blocked")
	ErrBoardReported = errors.New("board reported")
	ErrBoardBucket   = errors.New("board bucket")

	ErrBoardNotPopular = errors.New("board not popular")

	ErrPermBoardCreatePermission = errors.New("no board create permission")

	ErrPermBoardReadHidden   = errors.New("hidden board")
	ErrPermBoardReadBlocked  = errors.New("blocked board")
	ErrPermBoardReadReported = errors.New("reported board")

	ErrPermBoardReadNotOver18  = errors.New("not over18")
	ErrPermBoardReadPermission = errors.New("no board read permission")

	ErrPermPostReadOnly        = errors.New("read only")
	ErrPermPostBannedByBoard   = errors.New("banned by board")
	ErrPermBoardPostPost       = errors.New("no user post permission")
	ErrPermBoardPostRestricted = errors.New("only board friends")
	ErrPermBoardPostViolateLaw = errors.New("violate law")
	ErrPermBoardPostPermission = errors.New("no board post permission")

	ErrPermBoardPostLoginDays = errors.New("invalid login days")
	ErrPermBoardPostPostLimit = errors.New("reached post limit")

	ErrPermBoardEditPermission = errors.New("no board edit permission")
)

func ErrBoardCooldown(diffNanoTS types.NanoTS) error {
	diffTS := diffNanoTS.ToTime8()
	diffMin := diffTS / 60
	diffSec := diffTS % 60

	return fmt.Errorf("board cooldown %v:%02d", diffMin, diffSec)
}

func ErrBoardPosttime(diffNanoTS types.NanoTS) error {
	diffTS := diffNanoTS.ToTime8()
	diffMin := diffTS / 60
	diffSec := diffTS % 60

	return fmt.Errorf("board posttime %v:%02d", diffMin, diffSec)
}

func ErrFloodReject(diffNanoTS types.NanoTS) error {
	diffTS := diffNanoTS.ToTime8()
	diffMin := diffTS / 60
	diffSec := diffTS % 60

	return fmt.Errorf("flood reject %v:%02d", diffMin, diffSec)
}
