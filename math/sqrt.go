package math

import (
	"math"
	"github.com/jamillosantos/functions-for-govaluate/errors"
)

func Sqrt(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Sqrt(v), nil
	}
	return nil, errors.WrongParamsCount
}
