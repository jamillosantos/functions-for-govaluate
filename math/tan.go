package math

import (
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

func Tan(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Tan(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Tanh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Tanh(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Atan(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Atan(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Atan2(args ...interface{}) (interface{}, error) {
	if len(args) == 2 {
		v1, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		v2, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(1)
		}
		return math.Atan2(v1, v2), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Atanh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Atanh(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}
