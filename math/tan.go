package math

import (
	"math"
	"github.com/jamillosantos/functions-for-govaluate/errors"
)

func Tan(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Tan(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Tanh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Tanh(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Atan(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Atan(v), nil
	}
	return nil, errors.WrongParamsCount
}

func Atan2(args ...interface{}) (interface{}, error) {
	if len(args) == 2 {
		v1, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		v2, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(1)
		}
		return math.Atan2(v1, v2), nil
	}
	return nil, errors.WrongParamsCount
}

func Atanh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), errors.NewWrongParamType(0)
		}
		return math.Atanh(v), nil
	}
	return nil, errors.WrongParamsCount
}
