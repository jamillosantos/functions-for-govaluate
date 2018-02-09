package math

import (
	"math"
	"github.com/jamillosantos/functions-for-govaluate/errors"
)

func Cos(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Cos(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Cosh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Cosh(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Acos(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Acos(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Acosh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Acosh(v), nil
	}
	return nil, errors.WrongParamsCount
}
