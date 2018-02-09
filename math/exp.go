package math

import (
	"math"
	"github.com/jamillosantos/functions-for-govaluate/errors"
)

func Exp(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Exp(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Exp2(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Exp2(v), nil
	}
	return nil, errors.WrongParamsCount
}
