package math

import (
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

func Sin(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Sin(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Sinh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Sinh(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Asin(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Asin(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Asinh(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Asinh(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}
