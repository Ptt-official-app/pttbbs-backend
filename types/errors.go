package types

import (
	"errors"
	"fmt"
)

func ErrRecover(err interface{}) error {
	str := fmt.Sprintf("(recover) %v", err)
	return errors.New(str)
}
