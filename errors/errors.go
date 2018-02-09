package errors

import (
	"github.com/pkg/errors"
	"fmt"
)

var WrongParamsCount = errors.New("wrong params count")

func NewWrongParamType(index int) error {
	return fmt.Errorf("wrong param type at %d", index)
}
