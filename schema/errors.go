package schema

import "errors"

var (
	ErrNoLock               = errors.New("no lock")
	ErrNoMatch              = errors.New("no match")
	ErrNoUpdate             = errors.New("no update")
	ErrNoCreate             = errors.New("no create")
	ErrInvalidUserFavorites = errors.New("invalid user favorites")
	ErrRDBAlreadyExists     = errors.New("rdb already exists")

	ErrNotFound = errors.New("not found")

	ErrInvalidPTTBID = errors.New("invalid bid")
)
