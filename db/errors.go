package db

import "errors"

var (
	ErrInvalidOp     = errors.New("invalid op")
	ErrEmptyInRemove = errors.New("nil in remove")

	ErrRDBAlreadyExists = errors.New("rdb already exists")
)
