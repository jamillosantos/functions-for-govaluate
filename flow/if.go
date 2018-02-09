package flow

import (
	"github.com/jamillosantos/functions-for-govaluate/errors"
)

func If(args ...interface{}) (interface{}, error) {
	if len(args) == 3 {
		condition := false
		if args[0] != nil {
			switch cond := args[0].(type) {
			case bool:
				condition = cond
			case float64, float32, int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
				return IsDefined(args[0])
			default:
				return nil, errors.NewWrongParamType(0)
			}
		}
		if condition {
			return args[1], nil
		} else {
			return args[2], nil
		}
	}
	return nil, errors.WrongParamsCount
}
