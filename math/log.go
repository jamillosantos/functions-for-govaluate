package math

import (
	"math"
	"github.com/jamillosantos/functions-for-govaluate"
)

func Log(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Log(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}

func Log10(args ...interface{}) (interface{}, error) {
	if len(args) == 1 {
		v, ok := args[0].(float64)
		if !ok {
			return math.NaN(), functions_for_govaluate.NewWrongParamType(0)
		}
		return math.Log10(v), nil
	}
	return nil, functions_for_govaluate.WrongParamsCount
}
