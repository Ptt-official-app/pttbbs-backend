package types

import (
	"errors"
	"fmt"
)

func ErrRecover(err any) error {
	str := fmt.Sprintf("(recover) %v", err)
	return errors.New(str)
}
