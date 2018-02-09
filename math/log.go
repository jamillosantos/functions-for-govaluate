package math

import (
	"math"
	"github.com/jamillosantos/functions-for-govaluate/errors"
)

func Log(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := ToFloat64(args[0])
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Log(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Log10(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := ToFloat64(args[0])
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Log10(v), nil
	}
	return nil, errors.WrongParamsCount
}
